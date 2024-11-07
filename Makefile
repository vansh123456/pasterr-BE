migrateup:
	migrate -path db/migrations -database 'DBurl' -verbose up
migratedown:
<<<<<<< HEAD
	migrate -path db/migrations -database 'DBURL' -verbose down
.PHONY: migrateup migratedown
=======
	migrate -path db/migrations -database 'postgresql://neondb_owner:Z8GqBAXIf1dH@ep-orange-surf-a5m7v2pi.us-east-2.aws.neon.tech/neondb?sslmode=require' -verbose down
sqlc: 
	sqlc generate
.PHONY: migrateup migratedown sqlc
>>>>>>> 4ccff71 (sqlc:Init database)
