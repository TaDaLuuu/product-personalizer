start:
	go run cmd/dev/main.go
install:
	./scripts/install.sh
migration:
	./scripts/make-migration.sh
migrate:
	sql-migrate up -config=configs/dbconfig.yml
migrate-rollback:
	sql-migrate down -config=configs/dbconfig.yml