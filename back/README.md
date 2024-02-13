# Back

This is a [Fiber](https://gofiber.io/) project.

This project is configured to run with Docker. You should usually not run the backend alone, but rather use the `make` command to run all containers at once.

## Getting Start

### Local machine

First, create a `.env` file such as:

```
HCP_CLIENT_ID=yourClientId
HCP_CLIENT_SECRET=yourClientSecret
HCP_ORG_ID=yourOrgId
HCP_PROJECT_ID=yourProjectId
```

Then, run the development server:

```bash
go mod download && go mod verify
go install github.com/cosmtrek/air@v1.49.0
air -c .air.toml
```

Open [http://localhost:3001](http://localhost:3001) with your browser to see the result.

### Dockerfile

You can build and run the image as a standalone container.

In airneis/back/:

#### Prod
```
docker build --tag airneis-back .
docker run -d -p 3001:3000 airneis-back 
```
#### Dev
```
docker build --tag airneis-back-dev --target dev .
docker run -d -p 3001:3000 --mount type=bind,source=.,target=/build airneis-back-dev
```
Dev mode allow for hot reloading.
