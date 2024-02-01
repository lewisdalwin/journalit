# Include variables from the .envrc file
include .envrc

.PHONY:	confirm
confirm:
	@echo 'Confirm this action: [y/N]' && read ans && [ $${ans:-N} = y ]


## run-web: run the cmd/web application
.PHONY:	run-web
run-web: 
	@go run ./cmd/web -db-dsn=${JOURNALIT_DB_DSN}


## db-psql: connect to the database using psql
.PHONY:	db-psql
db-psql:
	psql ${JOURNALIT_DB_DSN}

## db-migrations-new name=$1: create a new database migration
.PHONY:	db-migrations-new
db-migrations-new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## db-migrations-up: apply all up database migrations
.PHONY:	db-migrations-up
db-migrations-up: confirm
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${JOURNALIT_DB_DSN} up


