FRONTEND_PATH = $(PWD)/frontend
BACKEND_PATH = $(PWD)/backend
FRONTEND_CONTAINER_NAME = front
BACKEND_CONTAINER_NAME = back
DATABASE_CONTAINER_NAME = db

default: commands

test:
	@if [ -d "$(FRONTEND_PATH)" ]; then cd $(FRONTEND_PATH) && npm run test; fi
	@if [ -d "$(BACKEND_PATH)" ]; then cd $(BACKEND_PATH) && go test ./...; fi

run: test
	@if [ -d "$(FRONTEND_PATH)" ]; then cd $(FRONTEND_PATH) && npm run dev; fi
	@if [ -d "$(BACKEND_PATH)" ]; then cd $(BACKEND_PATH) && $(MAKE) run; fi

build: test
	@if [ -d "$(FRONTEND_PATH)" ]; then cd $(FRONTEND_PATH) && npm run build; fi
	@if [ -d "$(BACKEND_PATH)" ]; then cd $(BACKEND_PATH) && $(MAKE) build; fi

up:
	@echo "Starting all the containers"
	@docker compose up -d
	@echo "Containers started"

down:
	@echo "Stopping and deleted all the containers"
	@docker compose down
	@echo "Containers stopped and deleted"

stop:
	@echo "Stopping all the containers"
	@docker compose stop
	@echo "Containers stopped"

start:
	@echo "Starting all the containers"
	@docker compose start
	@echo "All the containers are started"

restart:
	@echo "restarting containers"
	@docker compose restart
	@echo "Containers restarted"

start-front:
	@echo "Starting frontend"
	@docker compose start $(FRONTEND_CONTAINER_NAME)
	@echo "Started frontend"

stop-front:
	@echo "Stopping frontend"
	@docker compose stop $(FRONTEND_CONTAINER_NAME)
	@echo "Stopped frontend"

restart-front:
	@echo "Restarting frontend"
	@docker compose restart $(FRONTEND_CONTAINER_NAME)
	@echo "Restarted frontend"

start-back:
	@echo "Starting backend"
	@docker compose start $(BACKEND_CONTAINER_NAME)
	@echo "Started backend"

stop-back:
	@echo "Stopping backend"
	@docker compose stop $(BACKEND_CONTAINER_NAME)
	@echo "Stopped backend"

restart-back:
	@echo "Restarting backend"
	@docker compose restart $(BACKEND_CONTAINER_NAME)
	@echo "Restarted backend"

start-db:
	@echo "Starting backend"
	@docker compose start $(DATABASE_CONTAINER_NAME)
	@echo "Started backend"

stop-db:
	@echo "Stopping database"
	@docker compose stop $(DATABASE_CONTAINER_NAME)
	@echo "Stopped database"

restart-db:
	@echo "Restarting database"
	@docker compose restart $(DATABASE_CONTAINER_NAME)
	@echo "Restarted database"

commands:
	@echo " List of every command"
	@echo "	- commands: shows this messages"
	@echo "	- test: runs the test. It can be used in frontend or backend"
	@echo "	- run: test the code and runs it. It can be used in frontend or backend"
	@echo "	- up: create and start all containers"
	@echo "	- down: stop and delete all the containers"
	@echo "	- start: starts all containers "
	@echo "	- stop: stop all the containers"
	@echo "	- restart: restart all the containers"
	@echo "	- start-[front || back || db] start the selected container"
	@echo "	- stop-[front || back || db] stop the selected container"
	@echo "	- restart-[front || back || db] restart the selected container"
	@echo "\n To use those commands run make [name of the command]"