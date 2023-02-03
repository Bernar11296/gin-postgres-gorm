db:
	docker run --name gorm -e POSTGRES_USER=gorm -e POSTGRES_PASSWORD=gorm -e POSTGRES_DATABASE=gorm -p 5432:5432 postgres
stop:
    docker exec gorm pg_dump -U greenlight -d dorm > db/data/db_dump.sql
    docker stop gorm
    docker rm gorm
enter:
    docker exec -it gorm psql -U gorm
start: 
	docker start gorm

pgadmin:
	docker pull dpage/pgadmin4
run-pg:
	docker run -p 80:80 -e 'PGADMIN_DEFAULT_EMAIL=bernar@domain.com' -e 'PGADMIN_DEFAULT_PASSWORD=12345678' -d dpage/pgadmin4
# 	// docker run -p 80:80 \
# >     -e 'PGADMIN_DEFAULT_EMAIL=user@domain.com' \
# >     -e 'PGADMIN_DEFAULT_PASSWORD=SuperSecret' \
# >     -d dpage/pgadmin4
# //