migrate:
	sudo docker exec -it has-gorira-api bash -c "cd db; sql-migrate up"

migrate_status:
	sudo docker exec -it has-gorira-api bash -c "cd db; sql-migrate status"

migrate_redo:
	sudo docker exec -it has-gorira-api bash -c "cd db; sql-migrate redo"

seed:
	sudo docker exec -it has-gorira-api go run /hasGoriraAPI/db/seed/seed.go

up:
	sudo docker-compose up

down:
	sudo docker-compose down
	sudo docker volume rm hasgoriraapi_db-volume

in:
	sudo docker exec -ti has-gorira-api /bin/bash

in_db:
	sudo docker exec -ti db /bin/bash