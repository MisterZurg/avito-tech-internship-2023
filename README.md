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
> make run
> ```

## Используемые зависимости и тулы
- [protoc](https://github.com/protocolbuffers/protobuf) mechanism for serializing structured data.
- [gRPC-Gateway](https://github.com/grpc-ecosystem/grpc-gateway) plugin that reads protobuf service definitions and generates a reverse-proxy server which translates a RESTful HTTP API into gRPC
- [sqlx](https://github.com/jmoiron/sqlx) extension on go's standard database/sql
- [pgx](https://github.com/jackc/pgx) pure Go driver and toolkit for PostgreSQL
- [env](https://github.com/caarlos0/env) simple and zero-dependencies library to parse environment variables into structs


## Задание
> [!IMPORTANT]
> **API можно посмотреть в [contract.swagger.json](proto%2Fcontract.swagger.json)**
### Основное задание:
**Метод создания сегмента.**
> [!NOTE]
> Принимает slug (название) сегмента.
> <picture>
>   <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/light-theme/example.svg">
>   <img alt="Example" src="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/dark-theme/example.svg">
> </picture><br>
>
> .http
> ```http request
> POST /v1/example/createSegment HTTP/1.1
> Host: localhost:8090
> Content-Type: application/json
> Accept: */*
> 
> {
> "slug": "AVITO_DISCOUNT_502"
> } 
> ```
> .proto
> ```protobuf
> service SegmentsService {
>   rpc CreateSegment(CreateSegmentRequest) returns (CreateSegmentResponse) {
>     option (google.api.http) = {
>       post: "/v1/example/createSegment"
>       body: "*"
>     };
>   // ...
> }
> ```

**Метод удаления сегмента.**
> [!NOTE]
> Принимает slug (название) сегмента.
> <picture>
>   <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/light-theme/example.svg">
>   <img alt="Example" src="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/dark-theme/example.svg">
> </picture><br>
>
> **.http**
> ```http request
> DELETE /v1/example/deleteSegment/AVITO_DISCOUNT_502 HTTP/1.1
> Host: localhost:8090
> Content-Type: application/json
> Accept: */*
>
> {
> "slug": "AVITO_DISCOUNT_502"
> }
> ```
> **.proto**
> ```protobuf
> service SegmentsService {
>   // ...
>   rpc DeleteSegment(DeleteSegmentRequest) returns (DeleteSegmentResponse) {
>     option (google.api.http) = {
>       delete: "/v1/example/deleteSegment/{slug}"
>     };
>   }
> }
> ```

**Метод добавления пользователя в сегмент.**
> [!NOTE]
> Принимает список slug (названий) сегментов которые нужно добавить пользователю, список slug (названий) сегментов которые нужно удалить у пользователя, id пользователя.
> <picture>
>   <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/light-theme/example.svg">
>   <img alt="Example" src="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/dark-theme/example.svg">
> </picture><br>
>
> **.http**
> ```http request
> POST /v1/example/addUserToSegment HTTP/1.1
> Host: localhost:8090
> Content-Type: application/json
> Accept: */*
>
> {
>   "slugsAdd": [
>     "I_WISH_TO",
>     "WORK_AT_AVITO"
>   ],
>   "slugsDel": [
>     "WITHOUT",
>     "AVITO_VOICE_MESSAGES"
>   ],
>   "userId": "10012"
> }
> ```
> **.proto**
> ```protobuf
> service UsersService {
>   rpc AddUserToSegment(AddUserToSegmentRequest) returns (AddUserToSegmentResponse) {
>     option (google.api.http) = {
>       post: "/v1/example/addUserToSegment"
>       body: "*"
>     };
>   }
>   // ...
> }
> ```
**Метод получения активных сегментов пользователя.**
> [!NOTE]
> Принимает на вход id пользователя.
> <picture>
>   <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/light-theme/example.svg">
>   <img alt="Example" src="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/dark-theme/example.svg">
> </picture><br>
>
> **.http**
> ```http request
> GET /v1/example/getActiveSegments/1006 HTTP/1.1
> Host: localhost:8090
> Content-Type: application/json
> Accept: */*
>
> {
> "user_id": 1006
> }
> ```
> **.proto**
> ```protobuf
> service UsersService {
>   // ...
>   rpc GetActiveSegments(GetActiveSegmentsRequest) returns (GetActiveSegmentsResponse) {
>     option (google.api.http) = {
>       get: "/v1/example/getActiveSegments/{user_id}"
>     };
>   }
> }
> ```
### Доп. задание 1:
> CSV

### Доп. задание 2:
> TTL

### Доп. задание 3:
> %