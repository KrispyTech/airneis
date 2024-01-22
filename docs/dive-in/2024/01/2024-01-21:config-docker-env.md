# Configure Docker environments

## Summary

Be able to up docker containers in prod or in dev environment

## Context

We want to be able to deploy our app in dev environment, with hot reloading and dev dependencies, so that we can code directly in our local files and immediately see the result served by docker containers. We also want to check our prod mode if needed just as easily.

## Solution

Thanks to multi-stage building, we can write our Dockerfile in such way that it can build our image two different ways and label those two images differently. Then, we can either:

- call our standard `docker-compose up -d` for prod mode
- override our standard docker-compose with `docker-compose-dev.yml` for dev mode.

### Dockerfile

To build a dev version:

```
FROM builder as dev
*** Install dev dependencies and hot reloading tools ***
```

To build a prod version:

```
FROM builder as prod
*** Install only necessary dependencies, compile directly, etc ***
```

### Docker compose

To override our standard docker-compose with our dev conf:

```
docker compose -f docker-compose.yml -f docker-compose-dev.yml up -d
```

In our `docker-compose-dev.yml`, we can add new configurations or override existing ones.

```
[...]
build:
      target: dev
[...]
```

Adding the `target` configuration allow us to choose the stage to build from our Dockerfile.
Default is the last one, in our case "prod".

```
[...]
volumes:
      - [host_path]:[container_path]
[...]
```

Allow us to share files from our host to our container, as needed for hot reloading.

### Make

#### Commands to add:

- `make up dev`: `docker compose -f docker-compose.yml -f docker-compose-dev.yml up -d` start all containers of the `docker-compose.yml` in dev mode