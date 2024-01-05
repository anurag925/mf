include .env
SERVER_TARGET='build/mf_server'

MIGRATION_TARGET=db/migrations
DATABASE_URL=${DB_URL}
MIGRATOR=migrate -verbose -path ${MIGRATION_TARGET} -database ${DATABASE_URL}

all:
	go run cmd/server/main.go

build:
	go build cmd/server/main.go
run:
	./main
script:
	go run cmd/script/main.go



gen_migration:
	migrate -verbose create -ext sql -tz utc -dir ${MIGRATION_TARGET} ${name} 
migrate_up:
	${MIGRATOR} up
migrate_up_to:
	${MIGRATOR} up ${version}
migrate_down:
	${MIGRATOR} down
migrate_down_to:
	${MIGRATOR} down ${version}
migrate_goto:
	${MIGRATOR} goto ${version}
migrate_drop:
	${MIGRATOR} drop
migrate_force:
	${MIGRATOR} force ${version}
migrate_version:
	${MIGRATOR} version