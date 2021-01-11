# API Documentations
I usually use [swagger](https://swagger.io/) to document a project APIs. However, for this simple project, APIs are documented in this markup file.

| Method | Link | Description |
| ---| ------------- |:-----|
| GET | /api/projects | List all projects |
| GET | /api/projects/:id | Get project given its id |
| POST | /api/projects/create | Create a project |

**NOTE: respose code for all apis will be 200 on success. Anything else indicates an error.**
## List all projects

example
```bash
curl -X GET http://159.89.109.93/api/projects
# response:
# [
#    {
#        "id": "5ffb7ffa11959e1ddd513da0",
#        "name": "fake name",
#        "owener_id": "f6b417dc-31bc-427c-9791-09b571a7c23d",
#        "state": "planned",
#        "progress": 0,
#        "participants_ids": [
#            "f6b417dc-31bc-427c-9791-09b571a7c23d",
#            "da99f96d-588c-4a8b-9403-421e6875711e"
#        ]
#    }
# ]
```

## Get project given its id
example
```bash
curl -X GET http://159.89.109.93/api/projects/5ffb7ffa11959e1ddd513da0
# response:
# [
#    {
#        "id": "5ffb7ffa11959e1ddd513da0",
#        "name": "fake name",
#        "owener_id": "f6b417dc-31bc-427c-9791-09b571a7c23d",
#        "state": "planned",
#        "progress": 0,
#        "participants_ids": [
#            "f6b417dc-31bc-427c-9791-09b571a7c23d",
#            "da99f96d-588c-4a8b-9403-421e6875711e"
#        ]
#    }
# ]
```

## Create a project
example
```bash
curl -X POST http://159.89.109.93/api/projects/create -H 'Content-Type: application/json' -d @create.json

# contents of create.json file :
# {
#   "name" : "fake name",
#   "owener_id" : "f6b417dc-31bc-427c-9791-09b571a7c23d",
#   "state" : "planned",
#   "progress" : 100,
#   "participants_ids" : ["f6b417dc-31bc-427c-9791-09b571a7c23d","da99f96d-588c-4a8b-9403-421e6875711e"]
# }
```