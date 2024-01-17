# List of make commands

To use the syntax is `make [name of the command]`

- `commands`: shows this messages
- `test`: runs the test. It can be used in frontend or backend
- `run`: test the code and runs it. It can be used in frontend or backend
- `up`: create and start all containers
- `down`: stop and delete all the containers
- `start`: starts all containers
- `stop`: stop all the containers
- `restart`: restart all the containers
- `start-[front || back || db]`: start the selected container
- `stop-[front || back || db]`: stop the selected container
- `restart-[front || back || db]`: restart the selected container

If you only run `make` it's like you have run `make commands`