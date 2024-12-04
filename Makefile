migrateup:
	migrate -path db/migrations -database 'DBurl' -verbose up
migratedown:
	migrate -path db/migrations -database 'DBURL' -verbose down
sqlc:
	sqlc generate
.PHONY: migrateup migratedown sqlc
