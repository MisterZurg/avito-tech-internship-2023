# Сервис динамического сегментирования пользователей

> [!IMPORTANT]
> **Проблема:**
> В Авито часто проводятся различные эксперименты — тесты новых продуктов, тесты интерфейса, скидочные и многие другие. На архитектурном комитете приняли решение централизовать работу с проводимыми экспериментами и вынести этот функционал в отдельный сервис.
 
> **Задача:**
> Требуется реализовать сервис, хранящий пользователя и сегменты, в которых он состоит (создание, изменение, удаление сегментов, а также добавление и удаление пользователей в сегмент)

## Содержание


## Архитектура Продуктовой задачи

## Как запустить?
> <picture>
>   <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/light-theme/example.svg">
>   <img alt="Example" src="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/dark-theme/example.svg">
> </picture><br>
>
> ```shell
> cd build
> 
> ```

## Используемые зависимости и тулы
- [protoc](https://github.com/protocolbuffers/protobuf) mechanism for serializing structured data.
- [gRPC-Gateway](https://github.com/grpc-ecosystem/grpc-gateway) plugin that reads protobuf service definitions and generates a reverse-proxy server which translates a RESTful HTTP API into gRPC
- [sqlx](https://github.com/jmoiron/sqlx) extension on go's standard database/sql
- [pgx](https://github.com/jackc/pgx) pure Go driver and toolkit for PostgreSQL
- [env](https://github.com/caarlos0/env) simple and zero-dependencies library to parse environment variables into structs


## Задание
### Основное задание:
**Метод создания сегмента.**
> Принимает slug (название) сегмента.
> <picture>
>   <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/light-theme/example.svg">
>   <img alt="Example" src="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/dark-theme/example.svg">
> </picture><br>
>
> .http
> 


**Метод удаления сегмента.**
> Принимает slug (название) сегмента.

**Метод добавления пользователя в сегмент.**
> Принимает список slug (названий) сегментов которые нужно добавить пользователю, список slug (названий) сегментов которые нужно удалить у пользователя, id пользователя.

**Метод получения активных сегментов пользователя.**
> Принимает на вход id пользователя.


### Доп. задание 1:
> CSV

### Доп. задание 2:
> TTL

### Доп. задание 3:
> %