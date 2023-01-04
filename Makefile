run: clean remove_image
	@docker-compose up -d

debug: down remove_image
	@docker-compose up -d

down:
	@docker-compose down

logs:
	@docker logs oj-server

clean: down
	@sudo rm -rf data
	@sudo rm -rf logs

backend_image_name = "oj-server"
backend_image ="$(shell docker images | grep $(backend_image_name) | awk '{print $$1}')"
remove_image:

ifeq ($(backend_image),$(backend_image_name))
	@docker image rm $(backend_image_name)
endif
