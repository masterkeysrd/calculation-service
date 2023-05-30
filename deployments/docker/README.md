# Calculation System: Docker Deployments

## Overview

This directory contains the docker deployment configurations for the calculation system.

## Deployments

> **Note:** All the commands below must be run in the root directory of the project.

### Deploy all

To deploy all the docker containers:

```bash
make deploy_docker/all
```

> **Note:** Before running this command, you must update the `server/docker-compose.yml` file with the correct `RANDOM_ORG_API_KEY` environment variable.

To stop all the docker containers:

```bash
make deploy_docker/all_stop
```

### Deploy API documentation

To deploy the API documentation:

```bash
make deploy_docker/api_docs
```

To stop the API documentation:

```bash
make deploy_docker/api_docs_stop
```

### Deploy server

To deploy the server:

```bash
make deploy_docker/server
```

> **Note:** Before running this command, you must update the `server/docker-compose.yml` file with the correct `RANDOM_ORG_API_KEY` environment variable.

To stop the server:

```bash
make deploy_docker/server_stop
```

### Deploy database

To deploy the database:

```bash
make deploy_docker/database
```

To stop the database:

```bash
make deploy_docker/database_stop
```

## Help

To see all the available commands:

```bash
make deploy_docker
```

or

```bash
make deploy_docker/help
```
