# MiCasino Devops Challenge
You can see the challenge requirements [here](docs/micasino-prueba-tecnica-devops.odt?raw=1) .

## Create .env file
Create an .env file with these paramenters
```
DB_HOST="localhost"
DB_NAME="devops"
DB_USER="postgres"
DB_PORT="5432"
DB_PASSWORD="postgres"
```

## Build Api
Build the api code by executing:
```shell
go build -o ./build/bin -v ./cmd/api
```

## Run Api
You can also run it by executing:
```shell
go run ./cmd/api
```

## Swagger docs
You can access the Swagger documentation using your browser on the port 3000. For example:
```shell
http://localhost:3000/api/v1/swagger/index.html
```

## Testing the api

### Create Users
```shell
curl -X 'POST' \
  'http://localhost:3000/api/v1/auth/signUp' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    "us_usuario": "pdominguez",
    "us_nombre": "pedro",
    "us_apellido": "dominguez",
    "us_correo": "correo@email.com",
    "us_clave": "clave123",
    "us_esactivo": true,
    "us_eliminado": false
}'
```
### Create Login
```shell
curl -X 'POST' \
  'http://localhost:3000/api/v1/auth/login' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "password": "clave123",
  "username": "pdominguez"
}'
```

### Devops Save
```shell
curl -X 'POST' \
  'http://localhost:3000/api/v1/devops/save' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer <access token obtained on login>' \
  -H 'Content-Type: application/json' \
  -d '{ "message": "This is a test",
    "to": "Juan Perez",
    "from": "Rita Asturia",
    "timeToLifeSec": 45
}'

```
