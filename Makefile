create-migration:
	goose -dir "server/database/migrations" create $(name) sql

run-migrations:
	goose -dir "server/database/migrations" postgres "user=postgres password=mysecretpassword dbname=jun2-ish_db sslmode=disable host=localhost" up

run-migrations-test:
	goose -dir "server/database/migrations" postgres "user=postgres password=mysecretpassword dbname=jun2-ish_test_db sslmode=disable host=localhost" up

drop-migrations:
	goose -dir "server/database/migrations" postgres "user=postgres password=mysecretpassword dbname=jun2-ish_db sslmode=disable host=localhost" down

reset-migrations:
	goose -dir "server/database/migrations" postgres "user=postgres password=mysecretpassword dbname=jun2-ish_db sslmode=disable host=localhost" reset

create-docker-db:
	docker run --name jun2-ish_db -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_DB=jun2-ish_db -p 5432:5432 -d postgres

generate-mockery-files:
	go generate ./...
