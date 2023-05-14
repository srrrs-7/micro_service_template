init:
	cp compose.override.yml.sample compose.override.yml

migrate_file:
	migrate create -ext sql -dir migration migrate.sql

migrate_up:
	migrate -path migration -database "mysql://root:secret@tcp(localhost:3306)/aic" up

migrate_down:
	migrate -path migration -database "mysql://root:secret@tcp(localhost:3306)/aic" down