package operations

import (
	"context"
	"strings"
	"sync"

	"github.com/direktiv/apps/go/pkg/apps"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"gcloud/models"
)

const (
	successKey = "success"
	resultKey  = "result"

	// http related
	statusKey  = "status"
	codeKey    = "code"
	headersKey = "headers"
)

var sm sync.Map

const (
	cmdErr = "io.direktiv.command.error"
	outErr = "io.direktiv.output.error"
	riErr  = "io.direktiv.ri.error"
)

type accParams struct {
	PostParams
	Commands    []interface{}
	DirektivDir string
}

type accParamsTemplate struct {
	PostBody
	Commands    []interface{}
	DirektivDir string
}

func PostDirektivHandle(params PostParams) middleware.Responder {
	resp := &PostOKBody{}

	var (
		err  error
		ret  interface{}
		cont bool
	)

	ri, err := apps.RequestinfoFromRequest(params.HTTPRequest)
	if err != nil {
		return generateError(riErr, err)
	}

	ctx, cancel := context.WithCancel(params.HTTPRequest.Context())
	sm.Store(*params.DirektivActionID, cancel)
	defer sm.Delete(params.DirektivActionID)

	var responses []interface{}

	var paramsCollector []interface{}
	accParams := accParams{
		params,
		nil,
		ri.Dir(),
	}

	ret, err = runCommand0(ctx, accParams, ri)
	responses = append(responses, ret)

	// if foreach returns an error there is no continue
	cont = convertTemplateToBool("false", accParams, true)

	if err != nil && !cont {
		errName := cmdErr
		return generateError(errName, err)
	}

	paramsCollector = append(paramsCollector, ret)
	accParams.Commands = paramsCollector

	ret, err = runCommand1(ctx, accParams, ri)
	responses = append(responses, ret)

	// if foreach returns an error there is no continue
	cont = convertTemplateToBool("false", accParams, true)

	if err != nil && !cont {
		errName := cmdErr
		return generateError(errName, err)
	}

	paramsCollector = append(paramsCollector, ret)
	accParams.Commands = paramsCollector

	ret, err = runCommand2(ctx, accParams, ri)
	responses = append(responses, ret)

	// if foreach returns an error there is no continue

	if err != nil && !cont {
		errName := cmdErr
		return generateError(errName, err)
	}

	paramsCollector = append(paramsCollector, ret)
	accParams.Commands = paramsCollector

	s, err := templateString(`{
  "gcloud": {{ index . 2 | toJson }}
}
`, responses)
	if err != nil {
		return generateError(outErr, err)
	}

	responseBytes := []byte(s)

	// validate
	resp.UnmarshalBinary(responseBytes)
	err = resp.Validate(strfmt.Default)

	if err != nil {
		return generateError(outErr, err)
	}

	return NewPostOK().WithPayload(resp)
}

// exec
func runCommand0(ctx context.Context,
	params accParams, ri *apps.RequestInfo) (map[string]interface{}, error) {

	ir := make(map[string]interface{})
	ir[successKey] = false

	ri.Logger().Infof("executing command")

	at := accParamsTemplate{
		params.Body,
		params.Commands,
		params.DirektivDir,
	}

	cmd, err := templateString(`{{- if not (empty .Key) }}
bash -c 'echo {{ .Key }} | base64 -d > key.json'
{{- else }}
echo "using existing key.json file"
{{- end }}`, at)
	if err != nil {
		ir[resultKey] = err.Error()
		return ir, err
	}
	cmd = strings.Replace(cmd, "\n", "", -1)

	silent := convertTemplateToBool("true", at, false)
	print := convertTemplateToBool("false", at, true)
	output := ""

	envs := []string{}

	return runCmd(ctx, cmd, envs, output, silent, print, ri)

}

// end commands

// exec
func runCommand1(ctx context.Context,
	params accParams, ri *apps.RequestInfo) (map[string]interface{}, error) {

	ir := make(map[string]interface{})
	ir[successKey] = false

	ri.Logger().Infof("executing command")

	at := accParamsTemplate{
		params.Body,
		params.Commands,
		params.DirektivDir,
	}

	cmd, err := templateString(`gcloud auth activate-service-account {{ .Account }} --key-file=key.json`, at)
	if err != nil {
		ir[resultKey] = err.Error()
		return ir, err
	}
	cmd = strings.Replace(cmd, "\n", "", -1)

	silent := convertTemplateToBool("<no value>", at, false)
	print := convertTemplateToBool("false", at, true)
	output := ""

	envs := []string{}
	env0, _ := templateString(`HOME={{ .DirektivDir }}`, at)
	envs = append(envs, env0)

	return runCmd(ctx, cmd, envs, output, silent, print, ri)

}

// end commands

// foreach command
type LoopStruct2 struct {
	accParams
	Item        interface{}
	DirektivDir string
}

func runCommand2(ctx context.Context,
	params accParams, ri *apps.RequestInfo) ([]map[string]interface{}, error) {

	ri.Logger().Infof("foreach command over .Commands")

	var cmds []map[string]interface{}

	for a := range params.Body.Commands {

		ls := &LoopStruct2{
			params,
			params.Body.Commands[a],
			params.DirektivDir,
		}

		cmd, err := templateString(`{{ .Item.Command }}`, ls)
		if err != nil {
			ir := make(map[string]interface{})
			ir[successKey] = false
			ir[resultKey] = err.Error()
			cmds = append(cmds, ir)
			continue
		}

		silent := convertTemplateToBool("{{ .Item.Silent }}", ls, false)
		print := convertTemplateToBool("{{ .Item.Print }}", ls, true)
		cont := convertTemplateToBool("{{ .Item.Continue }}", ls, false)
		output := ""

		envs := []string{}
		env0, _ := templateString(`HOME={{ .DirektivDir }}`, ls)
		envs = append(envs, env0)
		env1, _ := templateString(`CLOUDSDK_CORE_PROJECT={{ .Body.Project }}`, ls)
		envs = append(envs, env1)

		r, err := runCmd(ctx, cmd, envs, output, silent, print, ri)
		if err != nil {
			ir := make(map[string]interface{})
			ir[successKey] = false
			ir[resultKey] = err.Error()
			cmds = append(cmds, ir)

			if cont {
				continue
			}

			return cmds, err

		}
		cmds = append(cmds, r)

	}

	return cmds, nil

}

// end commands

func generateError(code string, err error) *PostDefault {

	d := NewPostDefault(0).WithDirektivErrorCode(code).
		WithDirektivErrorMessage(err.Error())

	errString := err.Error()

	errResp := models.Error{
		ErrorCode:    &code,
		ErrorMessage: &errString,
	}

	d.SetPayload(&errResp)

	return d
}

func HandleShutdown() {
	// nothing for generated functions
}
