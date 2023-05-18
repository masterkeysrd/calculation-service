# Calculation Service: API Documentation

## About

Calculation Service is a service that provides Arithmetic Calculation
features. The API is exposed via REST APIs. The service has cost per
operation, each operation has a deferent cost and the user is charge by
each successful operation performed on the service.

> To know the price of each type of calculation please make a GET request
> to `api/v1/operation` that endpoint retrieves a `json` array with the price
> of each operation.

## API Documentation

### Running the API Documentation locally

To run the API documentation locally you need to have `docker` and `docker-compose`
installed on your machine.

To run the API documentation locally you need to run the following command in the
root of the project:

```bash
make api-docs/run
```

> The command above will start the API documentation on port `8090` and you can
> access it via <http://localhost:8090>

Also the command `make dev/start` will start the API documentation on port `8090`.

> To more help about the commands available to run the API documentation locally
> please run `make api-docs` or `make api-docs/help` command on the root of the project.
