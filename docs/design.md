# Project Design & Code Elplaination
ActionsOcean is written in Golang and uses [mongodb](https://www.mongodb.com/) to store the data.

We have followed the microservices pattern. Each microservice has its own seperate container, which makes it easy to deploy the project to any server or use cloud solutions. (AWS, Kubernetes and etc)
## CI/CD pipeline
We use Github Actions to as a CI/CD solution. See [here](../.github/workflows/go.yml). When a new commit is made in master branch, following steps happen automatically. (CI Part)

1. Project is built and all tests will be executed.
2. If all tests pass, Project is build and deployed on docker hub.

Later, any new updates on the docker hub, will be recognized automatically by the server, and will be pulled using a watchtower container. (CD part)

## Containers
For our simple solution, we just manage all containers using docker-compose. Although each container can be deployed individualy to cloud solutions (such as AWS), I chose docker-compose as it is cheaper to just host everything in a droplet on DigitalOcean, rather than using its cloud services.
```yaml
version: '3.3'

services:
  web:
    image: cih2001/actions-ocean-web:latest
    ports:
      - "80:1234"
    labels:
      - "com.centurylinklabs.watchtower.enable=true"
    restart: always
    depends_on:
      - watchtower
    environment:
      DBUSERNAME: ${MONGO_USERNAME}
      DBPASSWORD: ${MONGO_PASSWORD} 

  mongodb:
    image: mongo
    ports:
      - 27017:27017
    environment:
      MONGO_INITDB_ROOT_USERNAME: ${MONGO_USERNAME}
      MONGO_INITDB_ROOT_PASSWORD: ${MONGO_PASSWORD}
    command: mongod --auth
    healthcheck: 
      test: ["CMD", "curl", "-f", "http://localhost:27017"]
      interval: 30s
      timeout: 10s
      retries: 5
      # logging:
      # driver: none # Disable mongo logs.

  watchtower:
    image: containrrr/watchtower
    command: --label-enable --cleanup --interval 30
    labels:
      - "com.centurylinklabs.watchtower.enable=true"
    network_mode: none
    restart: always
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
```

#### Services
1. web: is the main program implementing apis
2. mongodb: a container hosting mongo DB
3. watchtower: just to automatically pull updates on docker hub: *cih2001/actions-ocean-web:latest*

## Database
mongodb is hosted inside a database. however, for our simple project, I did not create and external volume to store data permanently. our database is volatile and will reset whenever a system prune occurs.

## Major frameworks used:
1. [Echo](https://echo.labstack.com/), A High performance, extensible, minimalist Go web framework
2. [validator](https://github.com/go-playground/validator), Go Struct and Field validation
3. [mongo-go-driver](https://github.com/mongodb/mongo-go-driver), The Go driver for MongoDB

## Implementation
1. Some odd implementation decision are made here and there in the code. However, there exist a comment in code in every situation that these decision are made.
2. Whenever a TODO or CAUTION tag is encountered in the code, it simply shows that I am aware of the problem/better solution, but simply do not want to use much time on details.
3. There are no security messures in place, such as authentication. The only reason for that is to avoid the unnecessary prolongation of the project.