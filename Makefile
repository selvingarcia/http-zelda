.PHONY: create_migration migration_up migration_dirty force_migration seed

GREEN = \033[1;32m
BLUE = \e[94m
MAGENTA = \e[45m

RESET = \033[0m

help:
	@printf "${MAGENTA}Available Commands:${RESET}\n"
	@printf "${BLUE}make create_migration name=table_name${RESET}    - Create a new migration\n"
	@printf "${BLUE}make migration_up${RESET}                        - Run all pending migrations\n"
	@printf "${BLUE}make migration_dirty${RESET}                     - Show current migration version and dirty state\n"
	@printf "${BLUE}make force_migration number=N${RESET}            - Force migration version to N (use to fix dirty state)\n"
	@printf "${BLUE}make seed ${RESET}            - Populate DB\n"


create_migration:
	@migrate create -ext sql -dir ${MIGRATE_URL} -seq ${name}

migration_up:
	@migrate -database ${DB_ADDR} -path ${MIGRATE_URL} up

migration_dirty:
	@migrate -database ${DB_ADDR} -path ${MIGRATE_URL} version

force_migration:
	@migrate -database ${DB_ADDR} -path ${MIGRATE_URL} force ${number}

seed: 
	@psql ${DB_ADDR} -c "\copy codes(code, phrase, title, content, meaning) FROM 'codes.csv' WITH (FORMAT csv, HEADER true)"