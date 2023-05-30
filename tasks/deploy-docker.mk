# deploy_docker/all - Deploy all docker containers
.PHONY: deploy_docker/all
deploy_docker/all: deploy_docker/database deploy_docker/server deploy_docker/api_docs

# deploy_docker/all_stop - Stop all docker containers
.PHONY: deploy_docker/all_stop
deploy_docker/all_stop: deploy_docker/api_docs_stop deploy_docker/server_stop deploy_docker/database_stop

# deploy_docker/database - Deploy database docker container
.PHONY: deploy_docker/database
deploy_docker/database:
	bash ./scripts/deploy_docker/database_deploy.sh

# deploy_docker/database_stop - Stop database docker container
.PHONY: deploy_docker/database_stop
deploy_docker/database_stop:
	bash ./scripts/deploy_docker/database_stop.sh

# deploy_docker/api_docs - Deploy api_docs docker container
.PHONY: deploy_docker/api_docs
deploy_docker/api_docs:
	bash ./scripts/deploy_docker/api_docs_deploy.sh

# deploy_docker/api_docs_stop - Stop api_docs docker container
.PHONY: deploy_docker/api_docs_stop
deploy_docker/api_docs_stop:
	bash ./scripts/deploy_docker/api_docs_stop.sh

# deploy_docker/server - Deploy server docker container
.PHONY: deploy_docker/server
deploy_docker/server:
	bash ./scripts/deploy_docker/server_deploy.sh

# deploy_docker/server_stop - Stop server docker container
.PHONY: deploy_docker/server_stop
deploy_docker/server_stop:
	bash ./scripts/deploy_docker/server_stop.sh

# deploy_docker/help - Help for deployment tasks
.PHONY: deploy_docker/help
deploy_docker/help:
	@echo "Deployment tasks help"
	@echo "  make deploy_docker/all                    deploy all docker containers"
	@echo "  make deploy_docker/all_stop               stop all docker containers"
	@echo "  make deploy_docker/database               deploy database docker container"
	@echo "  make deploy_docker/database_stop          stop database docker container"
	@echo "  make deploy_docker/api_docs               deploy api_docs docker container"
	@echo "  make deploy_docker/api_docs_stop          stop api_docs docker container"
	@echo "  make deploy_docker/server                 deploy server docker container"
	@echo "  make deploy_docker/server_stop            stop server docker container"
	@echo "  make deploy_docker/help                   display this help"
	@echo ""