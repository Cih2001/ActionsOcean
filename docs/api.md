# API Documentations
I usually use [swagger](https://swagger.io/) to document a project APIs. However, for this simple project, APIs are documented in this markup file.

| Link | Description | 
| ------------- |:-----|
| /api/projects | List all projects |
| /api/projects/:id | Get project given its id |
| /api/projects/create | Create a project |

## List all projects
example
```bash
curl -X GET http://159.89.109.93/api/projects
```

