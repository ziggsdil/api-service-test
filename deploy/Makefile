ifneq (,$(wildcard ./.env))
	include .env
	export
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

db: ##@Database Create database with docker-compose
	docker-compose -f docker-compose.yml -d --remove-orphans

open_db: ##@Database Open database console inside docker image
	docker exec -it name_postgres psql -d $(POSTGRES_DB) -U $(POSTGRES_USER)

