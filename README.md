# One Accord Service Generator Tool

To create a new service, first

```
go install github.com/oneaccord-llc/service-generator@latest
```

then run

```
oneaccord-generator
```

Enter the name of the service

then cd into the directory
and run

```
go mod tidy
```

The service is now ready, you can add/edit migrations and sql queries from `migrations/` and `sql/` directories.  
Add routes inside `routes/` folder

to update the models from sql queries, run

```
sqlc generate
```

The migrations are migrated at the start of the application

you can now run with `go run main.go` or by `make watch`
