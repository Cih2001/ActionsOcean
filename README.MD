# Actions Ocean
"ActionsOcean" is the name of a test project, just to leverage GitHub Actions to deply on a DigitalOcean server.

ActionsOcean project can be acessed on [http://159.89.109.93/](http://159.89.109.93/)

## Installation Guide
As mentioned, ActionsOcean can be reached at [http://159.89.109.93/](http://159.89.109.93/), however, if you want to deploy it locally or on your server, you need to do the follwing steps:
1. Make sure you have docker and docker-compose installed on your machine.
2. Create an *.env* file and store database username and password inside that.
```bash
echo "MONGO_USERNAME=username" > .env
echo "MONGO_PASSWORD=password" >> .env
```
##### deplyment on local machine
3. checkout the code
```bash
git clone git@github.com:Cih2001/ActionsOcean.git
cd ActionsOcean
```
4. build and run the project.
```bash
go build main.go && sudo docker-compose -f docker-compose-debug.yml build && sudo docker-compose -f docker-compose-debug.yml up
```

##### deployment on remote server.
make sure ro follow steps 1 & 2 on your remote server.

4. copy docker-compose.yml to your server.
```bash
# this command is executed on local macine
scp docker-compose.yml <server_address>:
```
5. run the project
```bash
# this command is executed on remote server
docker-compose up
```

## API Reference
See documentation [here](docs/api.md)

## Project Design & Code Elplaination
See documentation [here](docs/design.md)