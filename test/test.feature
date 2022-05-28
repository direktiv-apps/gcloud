Feature: gpg function test

Background:
* url demoBaseUrl
* def secrets = read('secrets/secrets.json')

Scenario: simple test

    Given path '/'
    Given header Direktiv-ActionID = 'development'
    Given header Direktiv-Tempdir = '/tmp'
    And request 
    """
    { 
      "account": '#(secrets.account)',
      "project": '#(secrets.project)',
      "key": '#(secrets.key)',
      "commands": [
        {
          "command":  "gcloud compute instances list --format json"
        }
      ]
    }
    """
    When method post
    Then status 200
    And match $ == 
    """
    {
    "gcloud": [
    {
      "result": "#notnull",
      "success": true
    }
    ]
    }
    """

Scenario: wrong account

    Given path '/'
    Given header Direktiv-ActionID = 'development'
    Given header Direktiv-Tempdir = '/tmp'
    And request 
    """
    { 
      "account": 'mydummy@myaccount',
      "project": '#(secrets.project)',
      "key": '#(secrets.key)',
      "commands": [
        {
          "command": "gcloud compute instances list --format json"
        }
      ]
    }
    """
    When method post
    Then status 500
    And def temp = response['Direktiv-Errorcode']
    And match header Direktiv-Errorcode == 'io.direktiv.command.error'

Scenario: continue on error

    Given path '/'
    Given header Direktiv-ActionID = 'development'
    Given header Direktiv-Tempdir = '/tmp'
    And request 
    """
    { 
      "account": '#(secrets.account)',
      "project": '#(secrets.project)',
      "key": '#(secrets.key)',
      "commands": [
        {
        "command": "gcloud nocommand test",
        "continue": true
        },
        {
        "command": "gcloud compute instances list --format json"
        }
      ],
    }
    """
    When method post
    Then status 200

Scenario: stop on error

    Given path '/'
    Given header Direktiv-ActionID = 'development'
    Given header Direktiv-Tempdir = '/tmp'
    And request 
    """
    { 
      "account": '#(secrets.account)',
      "project": '#(secrets.project)',
      "key": '#(secrets.key)',
      "commands": [
        {
        "command": "gcloud nocommand test",
        "continue": false
        },
        {
        "command": "gcloud compute instances list --format json"
        }
      ]
    }
    """
    When method post
    Then status 500

Scenario: use file json

    Given path '/'
    Given header Direktiv-ActionID = 'development'
    Given header Direktiv-Tempdir = '/tmp'
    And request 
    """
    { 
      "account": '#(secrets.account)',
      "project": '#(secrets.project)',
      "commands": [
        {
        "command": "gcloud nocommand test",
        "continue": true
        },
        {
        "command": "gcloud compute instances list --format json"
        }
      ]
    }
    """
    When method post
    Then status 200