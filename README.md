
# gcloud 1.0

Run Google gcloud commands in Direktiv

---
- #### Categories: cloud, aws
- #### Image: direktiv/gcloud 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/gcloud/issues
- #### URL: https://github.com/direktiv-apps/gcloud
- #### Maintainer: [direktiv.io](https://www.direktiv.io)
---

## About gcloud

This function executes alist of commans. It can run gcloud commands but has basic build tools installed as well, e.g. git.


### Example(s)
  #### Function Configuration
  ```yaml
  functions:
  - id: gcloud
    image: direktiv/gcloud:1.0
    type: knative-workflow
  ```
   #### Basic
   ```yaml
   - id: req
     type: action
     action:
       function: gcloud
       action:
       function: gcloud-build
       secrets: ["gcloud"]
       input:
         continue: false
         account: serviceaccount@project.iam.gserviceaccount.com
         project: project
         key: jq(.secrets.gcloud | @base64 )
         commands:
         - gcloud compute instances list --format=json
   ```
   #### Running Scripts
   ```yaml
   - id: req
     type: action
     action:
      function: gcloud
      secrets: ["gcloud"]
      files:
      - key: test.sh
        scope: inline
        mode: "0755"
        value: |-
          #!/bin/bash
          gcloud builds list --format=json
      input:
        account: serviceaccount@project.iam.gserviceaccount.com
        project: project
        key: jq(.secrets.gcloud | @base64 )
        commands:
        - ./test.sh
   ```

### Request



#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  Responds with a list of results
#### Reponse Types
    
  

[PostOKBody](#post-o-k-body)
#### Example Reponses
    
```json
{
  "gcloud": [
    {
      "result": [
        {
          "createTime": "2021-03-25T23:35:39.884256611Z",
          "finishTime": "2021-03-25T23:37:13.835995Z",
          "id": 123,
          "options": {
            "logging": "LEGACY"
          },
          "queueTtl": "3600s",
          "results": {
            "buildStepImages": [
              ""
            ],
            "buildStepOutputs": [
              ""
            ]
          },
          "sourceProvenance": {},
          "startTime": "2021-03-25T23:35:41.802247339Z",
          "status": "SUCCESS",
          "success": true,
          "timing": {
            "BUILD": {
              "endTime": "2021-03-25T23:37:12.877607119Z",
              "startTime": "2021-03-25T23:35:43.804308897Z"
            }
          }
        }
      ]
    }
  ]
}
```

### Errors
| Type | Description
|------|---------|
| io.direktiv.command.error | Command execution failed |
| io.direktiv.output.error | Template error for output generation of the service |
| io.direktiv.ri.error | Can not create information object from request |


### Types
#### <span id="post-o-k-body"></span> postOKBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| gcloud | [][PostOKBodyGcloudItems](#post-o-k-body-gcloud-items)| `[]*PostOKBodyGcloudItems` |  | |  |  |


#### <span id="post-o-k-body-gcloud-items"></span> postOKBodyGcloudItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | [interface{}](#interface)| `interface{}` | ✓ | |  |  |
| success | boolean| `bool` | ✓ | |  |  |


#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| account | string| `string` | ✓ | | Service account name | `sa@myproject.iam.gserviceaccount.com` |
| commands | []string| `[]string` |  | | List of commands to run. Use `--format=json` to get JSON results. | `gcloud compute instances list --format=json` |
| continue | boolean| `bool` |  | | If set to true all commands are getting executed and errors ignored. | `true` |
| key | string| `string` | ✓ | | JSON access file (IAM). |  |
| project | string| `string` | ✓ | | Specifies the project name. | `my-project-234` |

 
