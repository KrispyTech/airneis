# Make

Make is a utility to automate somme commands.  
Also it allows to have a kind of consistency between commands or to simplify commands.  
By running a single command make can execute multiple commands
For example a make install could install all the packages in backend, backend, and mobile.

Those commands are registered in the `Makefile` and have the same skeleton:  
`make [commandName] [arguments]`

# Commands I want to create

- `make up`: `docker compose up -d` start all container of the `docker-compose.yml`
- `make down`: `docker compose down ` stop all container of the `docker-compose.yml`
- `make commands`: Shows all the commands of the makefile