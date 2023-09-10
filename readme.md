# CTF-1
My first simple CTF-like project. 

Random sometimes done not quite right.

There is no flag, unlike in usual CTF, but it has a critical bug that allows to create unlimited number of accounts on other's e-mails and bypass captcha. Good luck to find and exploit it.

## Run the CTF
### Prepare 
```sh
cp example.postgres.env .postgres.env
cp example.redis.env .redis.env
cp example.smtp.env .smtp.env
# Edit it (I hope you'll find the exit)
vim .redis.env
vim .postgres.env
vim .smtp.env
```

### Run development 
Need configuration? It's in the `cmd/serv/main.go`
```sh
docker-compose -f ./docker-compose.dev.yml build
docker-compose -f ./docker-compose.dev.yml up
```

Latter will make other options to run, like prod


## For developers & researchers
### Code structure
in `cmd` located application entrypoints, files where execution starts.

in `internal/setup` located all the app initialization.

in `pkg` located simple tools (that are mostly can be useful in latter projects).

`internal/domain` contains application logic, main entities, data validation. It's a middle layer between the APIs and Repository to add flexability to the project structure. Also it's important for security reasons (for example to make same validation on different APIs, reduce risks)

`internal/repository` has all the storage-related code (DB, cache, disk).

in `internal/api` located everything related to an API (Handlers, routes).

### Tests
Nothing yet, TODO...