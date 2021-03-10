.PHONY : remigrate_db

remigrate_db:
	docker-compose stop; docker-compose rm -f db; docker-compose up -d db; sleep 1s; go run cmd/db/main.go;
