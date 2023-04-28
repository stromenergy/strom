# Developing Strom

## Install

### Go
Install Go by following the instructions at: https://go.dev/doc/install

### Sqlc
To have database queries available in Go, we use `sqlc` to generate Go code from sql files.
**macOS**
```
brew install sqlc
```
**Ubuntu**
```
sudo snap install sqlc
```

## Develop
Database migrations can be found in `database/migrations`. To create a new database migration run:
```
./scripts/add-migration.sh
```
Database queries can be found in `database/queries`. When writing queries, to generate Go code from new/updated sql files run:
```
./scripts/generate.sh
```
To update the build version run:
```
./scripts/update-version.sh
```

# Run
```
go run ./cmd/strom --log.text --log.level debug start
```
To view all the command line options available run:
```
go run ./cmd/strom -h
```