.PHONY := dev-run fmt cleanup create-migration migrate-db rollback-db
.DEFAULT_GOAL := dev-run

CONNECTION_STRING = "user=username password=password dbname=mckp sslmode=disable port=9999 host=0.0.0.0"

cleanup:
	@echo "Nothing to do yet!"

#
# Format the code using the in-built go formatter
#
fmt:
	go fmt

#
# Starts Backing Services inside of Docker
#
start-backing-services-dev:
	docker-compose -f docker-compose.dev.yaml up -d

stop-backing-services-dev:
	docker-compose -f docker-compose.dev.yaml down
#
# Start the system in development mode
# 
dev-run: start-backing-services-dev
	go run .

run:
	docker-compose up --force-recreate --build
#
# Creates a migeration file to be used by goose
#
create-migration:
	goose \
		-dir artifacts/migrations \
		postgres \
		$(CONNECTION_STRING) \
		create \
		$$MIGRATION_NAME \
		sql

#
# Moves the Database to the current State
# use all migrations
#
migrate-db:
	goose \
		-dir artifacts/migrations \
		postgres \
		$(CONNECTION_STRING) \
		up

#
# Moves the Databsae to the Previous State
# by rolling back one migration
#
rollback-db:
	goose \
	-dir artifacts/migrations \
	postgres \
	$(CONNECTION_STRING) \
	down
