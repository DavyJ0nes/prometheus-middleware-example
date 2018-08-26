# Prometheus Middleware Example

## Description

Example of how to build simple web service that uses middleware to wrap request logging and metric capturing.

The respository is structured so that the main componants of the service are within `router/` and are brought together within `main.go`. There are middleware for logging and monitoring defined within the `logger` and `metrics` packages. These can be added or not within `router` and will be applied to each request that comes in.

## Usage

Usage instructions

```shell
prometheus-middleware - v0.0.1

 Choose a command run in prometheus-middleware:

 compile     compiles binary for linux, osx and windows
 binary      builds a statically linked binary of the application (used in Docker image)
 image       builds a docker image for the application
 publish     pushes the tagged docker image to docker hub
 run         runs the application locally
 run_image   builds and runs the docker image locally
 test        run test suitde for application
 clean       remove binary from non release directory
 help        Show this help message
```

## TODO

- [x] Update README

## License

[MIT](./LICENSE)
