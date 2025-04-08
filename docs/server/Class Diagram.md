# Диаграмма классов

```mermaid
classDiagram

class Logger { }

class Config { }

class User { }
class Accounts { }
class TextData { }
class BinaryData { }
class Metadata { }

class UsersRepository { }
class AccountsRepository { }
class TextDataRepository { }
class BinaryDataRepository { }
class MetadataRepository { }

class UsersPostgres { }
class AccountsPostgres { }
class TextDataPostgres { }
class BinaryDataPostgres { }
class MetadataPostgres { }

class UsersService { }
class AccountsService { }
class TextDataService { }
class BinaryDataService { }
class MetadataService { }

class ServiceAggregator { }

class Middleware { }
class Handler { }

class Server { }

Server --> Handler : has
Handler --> ServiceAggregator : has

ServiceAggregator --> UsersService : has
ServiceAggregator --> AccountsService : has
ServiceAggregator --> TextDataService : has
ServiceAggregator --> BinaryDataService : has
ServiceAggregator --> MetadataService : has

UsersService --> UsersRepository : has
AccountsService --> AccountsRepository : has
TextDataService --> TextDataRepository : has
BinaryDataService --> BinaryDataRepository : has
MetadataService --> MetadataRepository : has

UsersService --> User : uses
AccountsService --> Accounts : uses
TextDataService --> TextData : uses
BinaryDataService --> BinaryData : uses
MetadataService --> Metadata : uses


UsersPostgres --> UsersRepository : implements
AccountsPostgres --> AccountsRepository : implements
TextDataPostgres --> TextDataRepository : implements
BinaryDataPostgres --> BinaryDataRepository : implements
MetadataPostgres --> MetadataRepository : implements

UsersRepository --> User : uses
AccountsRepository --> Accounts : uses
TextDataRepository --> TextData : uses
BinaryDataRepository --> BinaryData : uses
MetadataRepository --> Metadata : uses


Handler --> Config : has
Handler --> Logger : has

Middleware --> Config : has
Middleware --> Logger : has

Server --> Middleware : has
Server --> Handler : has
Server --> Logger : has


```

# Описание
## Server
- Назначение
    - Инициализовать и запустить сервер

## Middleware
- Назначение
    - Логгировать запросы
    - Прозводить аутентификацию пользователя

## Handler
- Назначение
    - Обрабатывать входящие запросы

## ServiceAggregator
- Назначение
    - Централизовать и сделать атомарными функции для работы с данными
