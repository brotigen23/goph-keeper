# goph-keeper

## Roadmap

### Server
#### Some
- [ ] Godoc
- [ ] Codecov
#### Code
- [x] Config
- [ ] DTO
  - [x] Account
- [x] Model
- [ ] Mapper
- [ ] Server
- [ ] Handler
  - [ ] Accounts
    - [ ] Post 
    - [x] Get
    - [x] Put
- [ ] Service
  - [ ] Auth
  - [x] Accounts
  - [ ] Text data
  - [ ] Binary data
  - [ ] Cards data
  - [ ] Metadata
- [ ] Repository
  - [x] Postgres
    - [x] Accounts
    - [x] Text data
    - [x] Binary data
    - [x] Cards data
    - [x] Metadata


### Client

- [ ] Core
  - [ ] Domain
  - [ ] API client
    - [x] Account 
  - [ ] Service
- [ ] CLI
  - [x] Auth cmd
    - [x] Register
    - [x] Login
- [ ] TUI
  - [x] Widgets
    - [x] Table
    - [x] Form
    - [x] Tabs

[Техническое задание](docs/specifications.md)

## Сервер
- [Диаграмма классов](docs/server/Class%20Diagram.md)
- [ER диаграмма](docs/server/ERD.md)


# Workflow

## API
### Auth
```http
POST login/ HTTP/1.1
content-type: application/json

{
    "login": <login>,
    "password": <password>    
}

```

### Accounts
```http
POST user/accounts HTTP/1.1
content-type: application/json

{
    "login": <login>,
    "password": <password>    
}
```

```http
GET user/accounts HTTP/1.1
content-type: application/json
```

## Shared
- Data
  - Accounts
  - Text data
  - Binary data
  - Cards data
  - Metadata
## Server
- Handler
  - Get shared model and map into domain model
  - Use service
  - Use returns of service to create response
- Service
  - Get domain model, do some stuff and use repo to store changes
  - Get repo returns to create return
- Repo
  - C