# Chat-Go

WebSocket chat golang

## Install

Для установки зависимостей make install

## Run project

Запустите базу данных с помощью docker-compose в deployments. Так же не забудьте изменить БД в [файле](configs/config.json).

```json
{
    "db": {
        "db" : "postgres | mysql | sqlite",
        ...
    }
}
```
