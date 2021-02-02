HostGator Backend Test
============================
> Project for HostGator backend test, using GoLang with some 3rd party plugins.
> The project consists and make a request to a external api to get information
> about cat breeds and create in DB cache of the results obtained previously.

### 3rd party libraries

1. [Gin](https://github.com/gin-gonic/gin)
2. [Mysql](https://github.com/go-sql-driver/mysql)
3. [JWT](https://github.com/dgrijalva/jwt-go)
4. [Swaggo](https://github.com/swaggo/swag)
5. [DotEnv](https://github.com/subosito/gotenv)
6. [Migrate](https://github.com/golang-migrate/migrate)

### Pre-requisites

> Docker version 20.10.2\
> Docker-compose 1.27.4

### Getting started

>You *MUST* setup environment file correctly, you can use .env.local as reference.
>Both of dockerfiles and application use that file, so set it carefully.

>You can run the webserver through docker-compose using the command:\
`docker-compose up --build`

* PS1: Migrations run automatically through docker-compose command
* PS2: Sometimes it takes a few seconds for migrations run correctly

>You can run tests through docker-compose as well, using:\
`docker-compose -f docker-compose.test.yml --build`


### Directory tree
    .
    ├── adapters                   # Adapters to manipulate some plugins
    ├── diagrams                   # Diagrams docs and images of project
    ├── docs                       # Docs generated for Swagger API
    ├── helpers                    # Classes, interfaces and methods of global porpose
    ├── middlewares                # Middlewares to intercept controllers default behavior
    ├── migrations                 # Sqls commands to versionate database
    ├── models                     # Tables schemas in struct form
    ├── repositories               # Repositories to encapsulate database operations
    ├── requests                   # Handlers functions to routers
    ├── router                     # Main router to flow control
    └── services                   # Services to manipulate and use data and flow directives 

### Docs and Diagrams

#### Swagger Docs works with app running at host and port defined in .env at endpoint /swagger

> [https://envhost:appport/swagger](https://envhost:appport/swagger)

#### Diagrams can be found at

> [Documentation](https://github.com/wiltonlcsj/golang-catbreed/tree/master/diagrams)