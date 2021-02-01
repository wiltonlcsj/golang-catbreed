HostGator Backend Test
============================
> Project for HostGator backend test, using GoLang with some 3rd party plugins

### 3rd party libraries
1. [Gin](https://github.com/gin-gonic/gin) 
2. [Mysql](https://github.com/go-sql-driver/mysql)
3. [JWT](https://github.com/dgrijalva/jwt-go)
4. [Swaggo](https://github.com/swaggo/swag)
5. [DotEnv](https://github.com/subosito/gotenv)

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
> [Swagger](https://envhost:appport/swagger)

#### Diagrams can be found at
> [Documentation](https://github.com/wiltonlcsj/golang-catbreed/tree/master/diagrams)