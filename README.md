## Introduction

### This is the reason why clean architecture is good
- Testability: Clean Architecture makes it easier to write automated tests for the application. Since each layer is independent of the others, you can test each layer in isolation, and the tests become more reliable.

- Maintainability: The separation of concerns provided by Clean Architecture makes it easier to maintain the application. You can modify one layer without affecting the other layers, and changes can be made more efficiently.

- Scalability: Clean Architecture is designed to handle complexity and is scalable. If you need to add a new feature or functionality, you can do so without impacting the existing codebase.

- Flexibility: Clean Architecture provides flexibility to the developer to choose the technologies and frameworks that they want to use for each layer. This architecture is technology agnostic, and you can switch out one technology for another without impacting the rest of the application.

- Independence: The different layers of Clean Architecture are independent, and one layer can be replaced or updated without affecting the others. This independence promotes modularity and makes it easier to work on the application as a team.

### Technology Stacks:
 - Go language
 - Echo framwork
 - PostgreSQL
 - Redis
 - Gorm
 - Docker
 - docker-compose

## API Documentation
Too see the API Documentation 
https://documenter.getpostman.com/view/26502057/2s93RRusgz

## .env
- `APP_PORT`: app / server port
- `DB_HOST`: database host PostgreSql
- `DB_USER`: database user PostgrSql
- `DB_PASSWORD`: database password PostgreSql
- `DB_NAME`: database name PostgreSql
- `REDIS_HOST`: redis host
- `REDIS_PORT`: redis port
- `REDIS_PASSWORD`: redis password

## Run project 
- Create `.env` file from `.env.example` 
- Run command
```bash
~/adheteguh$ go run main.go
```
- Log file will be written in main directory "app.log"

## Run project with docker-compose

- Run command
```bash
~/adheteguh$ docker-compose up 
```
- Run on background use -d
```bash
~/adheteguh$ docker-compose up -d
```

## License
[MIT](https://choosealicense.com/licenses/mit/)