.PHONY: init-db

DB_CONTAINER_NAME=my-postgres-db
DB_NAME=postgres
DB_USER=new_admin_user
DB_PASSWORD=your_password
SQL_FILE=init.sql

init-db:
	docker run --name $(DB_CONTAINER_NAME) -e POSTGRES_DB=$(DB_NAME) -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -p 5432:5432 -d postgres
	sleep 5  # Wait for PostgreSQL to start

	docker cp $(SQL_FILE) $(DB_CONTAINER_NAME):/docker-entrypoint-initdb.d/  # Copy SQL file into the container
	docker exec -it $(DB_CONTAINER_NAME) psql -U $(DB_USER) -d $(DB_NAME) -f /docker-entrypoint-initdb.d/$(SQL_FILE)

clean:
	docker stop $(DB_CONTAINER_NAME)
	docker rm $(DB_CONTAINER_NAME)
