# Calculation Service (calculation-service)

PoC Service to showcase Golang features performing calculations exposed via REST APIs

## About

Calculation Service is a service that provides Arithmetic Calculation features.

The service has cost per operation, each operation has a deferent cost and
the user is charge by each successful operation performed on the service.

> To know the price of each type of calculation please make a GET request
> to `api/v1/operation` that endpoint retrieves a `json` array with the price
> of each operation.

### Supported Operations

The current operation supported are:

- Addition
- Subtraction
- Multiplication
- Division
- Square Root
- Random String

## How to get a working from the source

For detailed instruction please visit
[Setting up your local environment guide](./docs/setup/setting-up-your-local-environemt.md) guide

You will need:

- Go - for compiling Go files. For help installing Go please follow the instructions available on <https://go.dev/doc/install>
- GVM - for managing Go versions. For help installing GVM please follow the instruction available on <https://go.dev/doc/install>
- Git - for version management. Download from here <https://git-scm.com/downloads> (MacOS users can also use Brew and Linux users can use the built-in package manager, eg apt, yum, etc).
- Make - for running the Makefile. Download from here <https://www.gnu.org/software/make/> (MacOS users can also use Brew and Linux users can use the built-in package manager, eg apt, yum, etc).

### Steps to setup the local environment

- Clone the repository

```bash
git clone https://github.com/masterkeysrd/calculation-service.git
```

- Go to `calculation-service` folder

```bash
cd calculation-service
```

- Initialize the Go Version, Packages, other dependencies and start the service in development mode.

```bash
make start_dev
```

## Architecture

The architecture documentation is info [here](./docs/architecture/README.md).

## API Documentation

The API documentation is info [here](./api/README.md).

## Deployments

The deployment documentation is info [here](./deployments/README.md).
