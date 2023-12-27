# Mega inc Devops Challenge
Mega inc Devops Challenge  #gin #gorm #wire #jwt #viper #cty #yaml

## Create file .env with paramaters
Create file with parameters
```
DB_HOST="localhost"
DB_NAME="devops"
DB_USER="postgres"
DB_PORT="5432"
DB_PASSWORD="postgres"
```

## Build Api
Build Api
```shell
go build -o ./build/bin -v ./cmd/api
```

## Run Api
Run Api
```shell
go run ./cmd/api
```

## accessing Swagger
Use the below URL for accessing Swagger from your browser, change the URL path as per your host. (localhost:3000, www.midominio.com,etc)
```shell
http://localhost:3000/api/v1/swagger/index.html
```

## Create Users
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
## Create Login
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

## Devops Save
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
