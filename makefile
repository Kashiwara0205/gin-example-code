migrate:
	sudo docker exec -it has-gorira-api bash -c "cd db; sql-migrate up"

migrate_status:
	sudo docker exec -it has-gorira-api bash -c "cd db; sql-migrate status"

migrate_redo:
	sudo docker exec -it has-gorira-api bash -c "cd db; sql-migrate redo"

seed:
	sudo docker exec -it has-gorira-api go run /hasGoriraAPI/app/seed/seed.go