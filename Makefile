# note: call scripts from /scripts

run:
		swag init -g cmd/main.go
		go run cmd/main.go


build:
		go build cmd/main.go


migrate:
		go run cmd/migrate/migrate.go --mode up