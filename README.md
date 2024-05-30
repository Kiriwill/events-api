# API Simples para Eventos
> API escrita para criar e obter eventos e suas metricas

# Dependencias
- github.com/kelseyhightower/envconfig v1.4.0
- github.com/labstack/echo/v4 v4.11.3
- gorm.io/driver/mysql v1.5.6
- gorm.io/gorm v1.25.10

## Estrutura do projeto
- `/cmd` will have all the code that will be built and shipped as binaries. they are most likely to be named `main.go` and be found at `/cmd/<executable-name>/main.go`
- `/pkg` will have the actual shared codebase of the project. all of the modules used by the entrypoints will come from this folder
- `/scripts` will have useful utilitary scripts such as `dev.sh` that usually runs the project at a dev-local environment

## Instalação
Com api no docker:
```bash 
docker compose up
```

Sem api no docker:
```bash 
make dev
```

## Examplos de rotas

> **ATENÇÂO**: Para acessar as rotas é necessário criar um usuario e depois passar um token obtido a partir da rota `user/signin` para a rota que será acessada

### Criar um novo usuario

`POST {host:port}/v1/user/`

#### requisição

```bash
curl --location 'localhost:8080/v1/user/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "user teste",
    "birthdate": "1989-01-02",
    "email": "will@email.com",
    "password": "1234"
}'
```

#### resposta

```json
{
    "ID": 6
}
```

### Obter um novo token de acesso

`POST {host:port}/v1/user/signin`

#### requisição

```bash
curl --location 'localhost:8080/v1/user/signin' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "will@email.com",
    "password": "1234"
}'
```

#### resposta

```json
{
    "ID": 6
}
```

### Criar um novo evento

`POST {host:port}/v1/events/`

#### requisição

```bash
curl --location 'localhost:8080/v1/events/' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTczNjE3MTR9.JYS8slfIo-1ySGPEFM6qBbVOPPUHomGTn8yEdecKjjw' \
--data '{
    "category": "teste categoria",
    "description": "descrição teste",
    "country": "colômbia"
}'
```

#### resposta

```json
{
    "id": 1,
    "createdat": "2024-05-30t13:20:14.821z",
    "updatedat": "2024-05-30t13:20:14.821z",
    "deletedat": null,
    "category": "teste categoria",
    "description": "descrição teste",
    "country": "colômbia"
}
```

### Busca um novo evento

`GET {host:port}/v1/events/`

#### Requisição

```bash
curl --location 'localhost:8080/v1/events/?date=2024-05&type=teste%20cate' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTczNjE3MTR9.JYS8slfIo-1ySGPEFM6qBbVOPPUHomGTn8yEdecKjjw'
```

#### Resposta

```json
[
    {
        "ID": 1,
        "CreatedAt": "2024-05-30T13:20:14.821Z",
        "UpdatedAt": "2024-05-30T13:20:14.821Z",
        "DeletedAt": null,
        "category": "teste categoria",
        "description": "descrição teste",
        "country": "colômbia"
    }
]
```

### Cria metricas (três paises com maior quantidade de eventos)

`GET {host:port}/v1/events/metrics/`

#### Requisição

```bash
curl --location 'localhost:8080/v1/events/metrics/' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTczNjE3MTR9.JYS8slfIo-1ySGPEFM6qBbVOPPUHomGTn8yEdecKjjw'
```

#### Resposta

```json
[
    {
        "Country": "colômbia",
        "Count": 55
    },
    {
        "Country": "brasil",
        "Count": 30
    },
    {
        "Country": "argentina",
        "Count": 12
    }
]
```
