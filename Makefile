DB_URL=postgres://admin:root@localhost:5432/musikmarching-db?sslmode=disable
DB_DIR=./db/migration

.PHONY: migration_new 
migration_new:
	goose -dir "$(DB_DIR)" postgres "$(DB_URL)" create "$(title)" sql

.PHONY: migration_reset 
migration_reset:
	goose -dir "$(DB_DIR)" postgres "$(DB_URL)" reset

.PHONY: migration_up
migration_up:
	goose -dir "$(DB_DIR)" postgres "$(DB_URL)" up

.PHONY: migration_down
migration_down:
	goose -dir "$(DB_DIR)" postgres "$(DB_URL)" down
