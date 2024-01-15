# Docker Basics

<p> In order to have a consistant environment to test and deploy our app, we will containerized each part of our application with the help of docker.

Docker is a software platform responsable for building, running, and managing containers. <br>
Containers are smalls, standalone virtual environment containing just what's necessary to run the application it deploys. <br>
Docker compose is a tool for defining and running multi-container Docker applications. A YAML file configures the application's services. Then, with a single command, it creates and starts all the services from the configuration. <br>
Dockerfiles are documents containing the commands to setup and build the images used to create the containers.</p>

We are going to use three containers:

- front: Nextjs front-end
- back: Go Fiber back-end
- db: postgreSQL database

## How to use?

If you don't have Docker installed on your environment, follow the instructions of the official docker website: https://docs.docker.com/engine/install/ <br>

Navigate to the root directory **airneis/**:

> cd path/to/airneis/

This command execute the code in docker-compose.yml, which builds then run the containers:

> docker compose up -d

The option -d runs the container in detached mode, meaning it gives back the CLI after running successfuly the containers instead of logging in the terminal.

To see your running containers:

> docker container ls

Response exemple:

```
CONTAINER ID   IMAGE           COMMAND                  CREATED         STATUS        PORTS                                       NAMES
52e00dd10ed1   postgres:16.1   "docker-entrypoint.s…"   2 seconds ago   Up 1 second   0.0.0.0:5432->5432/tcp, :::5432->5432/tcp   airneis-db-1
f10d32bac1e2   airneis-back    "/apiserver"             2 seconds ago   Up 1 second   0.0.0.0:3001->3001/tcp, :::3001->3001/tcp   airneis-back-1
1dcde425de1a   airneis-front   "npm start"              2 seconds ago   Up 1 second   0.0.0.0:3000->3000/tcp, :::3000->3000/tcp   airneis-front-1
```

Command is the instruction used to run the application inside the container. <br>
Ports follow the logic: [host_port]:[container_port]
