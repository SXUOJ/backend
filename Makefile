run: clean 
	@docker-compose up -d --force-recreate --build backend

debug: down 
	@docker-compose up -d --force-recreate --build backend

down:
	@docker-compose down

logs:
	@docker logs oj-server

clean: down
	@sudo rm -rf data
	@sudo rm -rf logs
