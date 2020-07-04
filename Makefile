PACKAGE_NAME=PCAPS
PCA_Pi_Server: docker-compose.yml
	docker-compose build
	docker-compose up -d
	docker-compose exec api /usr/local/go/bin/go build
	docker-compose down --rmi all --volumes
all: PCA_Pi_Server