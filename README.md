# One Accord Service Generator Tool

To create a new service, first

```
go install github.com/oneaccord-llc/service-generator@latest
```
then run 
```
oneaccord-generator
```

then cd into the directory
and run 
```
sqlc generate
```
to generate the models, queries files for the database 
finally run
```
go mod tidy
```

you can now run with `go run main.go` or by `make watch`
