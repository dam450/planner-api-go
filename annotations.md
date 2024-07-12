## Commands

```
go mod init journey

go install github.com/discord-gophers/goapi-gen@latest

goapi-gen --out ./internal/api/spec/journey.gen.spec.go ./internal/api/spec/journey.spec.json

go mod tidy

go get -u ./...

go install github.com/jackc/tern/v2@latest

tern init ./internal/pgstore/migrations

tern new --migrations ./internal/pgstore/migrations create_trips_table
tern new --migrations ./internal/pgstore/migrations create_participantes_table
tern new --migrations ./internal/pgstore/migrations create_activities_table
tern new --migrations ./internal/pgstore/migrations create_links_table

tern migrate --migrations ./internal/pgstore/migrations --config ./internal/pgstore/migrations/tern.conf

sqlc generate -f ./internal/pgstore/sqlc.yaml
go mod tidy
 go get -u ./...

```

## Links

### Material

https://efficient-sloth-d85.notion.site/NLW-16-Journey-013b69ad79894122824abd76bc0dab9b

### documentação da api

https://nlw-journey.apidocumentation.com/reference

### editor do swagger

https://editor.swagger.io

### go api gen

https://github.com/discord-gophers/goapi-gen
