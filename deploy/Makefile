ifneq (,$(wildcard ./.env))
	include .env
	export
endif

ifndef APP_PORT
export APP_PORT=8000
endif

ifndef APP_HOST
export APP_HOST=127.0.0.1
endif

HELP_FUNC = \
	%help; while(<>){push@{$$help{$$2//'options'}},[$$1,$$3] \
	if/^([\w-_]+)\s*:.*\#\#(?:@(\w+))?\s(.*)$$/}; \
    print"$$_:\n", map"  $$_->[0]".(" "x(20-length($$_->[0])))."$$_->[1]\n",\
    @{$$help{$$_}},"\n" for keys %help; \

#Commands
help: ##@Help Show this help
	@echo -e "Usage: make [target] ...\n"
	@perl -e '$(HELP_FUNC)' $(MAKEFILE_LIST)

run:
	go run ../cmd/ --config ../deploy

db: ##@Database Create database with docker-compose
	docker-compose -f docker-compose.yml -d --remove-orphans

open_db: ##@Database Open database console inside docker image
	docker exec -it name_postgres psql -d $(POSTGRES_DB) -U $(POSTGRES_USER)

lint: ##Service Check service in linters
	golangci-lint run --print-issued-lines=false --out-format code-climate:gl-code-quality-report.json,line-number --timeout=10m

