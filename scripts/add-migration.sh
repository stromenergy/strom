PKG="github.com/stromenergy/strom"
MIGRATIONS_DIR="${GOPATH}/src/${PKG}/database/migrations"
read -p "Name of migration: " NAME
FORMATTED_NAME=$(echo "$NAME" | tr A-Z a-z | tr ' ' '_')
migrate create -ext sql -dir $MIGRATIONS_DIR $FORMATTED_NAME