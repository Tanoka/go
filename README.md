# go

Dockerfile para construir un container con lo mínimo para compilar y ejecutar código go
https://github.com/Tanoka/dockerfiles/tree/master/golang

Ejecutar un programa
```
docker run --rm -it --network=host -e "GOPATH=/usr/test" -v $PWD:/usr/test my-golang go run /usr/test/mysql.go
```

Instalar una libreria
```
docker run --rm -it --network=host -e "GOPATH=/usr/test" -v $PWD:/usr/test my-golang go get -u github.com/go-sql-driver/mysql
```


