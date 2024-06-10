.PHONY: migrate-fresh
migrate-fresh:
	migrate -path db/migrations -database "postgresql://postgres.settumozapjmoshlvqgf:9dGn99bPyoTVRBP5@aws-0-ap-southeast-1.pooler.supabase.com:5432/postgres?sslmode=disable" down 1
	migrate -path db/migrations -database "postgresql://postgres.settumozapjmoshlvqgf:9dGn99bPyoTVRBP5@aws-0-ap-southeast-1.pooler.supabase.com:5432/postgres?sslmode=disable" up 1

.PHONY: run-redis-wsl
run-redis-wsl:
	wsl --exec sudo service redis-server start

.PHONY: stop-redis-wsl
stop-redis-wsl:
	wsl --exec sudo service redis-server stop