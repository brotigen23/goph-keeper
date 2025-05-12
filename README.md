# goph-keeper

## Usage
В корне проекта необходимо запустить СУБД Postgres и сервер
``` bash
docker compose up
```

После использовать клиент следующим образом:
### Accounts
```bash
keeper accounts [-m post|get|put|delete]
```
Флаги:
- --id - для методов put и delete
- --login - для методов post и put
- --password - для методов post и put
- --metadata - для методов post и put

### Text
```bash
keeper text [-m post|get|put|delete]
```
Флаги:
- --id - для методов put и delete
- --data - для методов post и put
- --metadata - для методов post и put

### Cards
```bash
keeper cards [-m post|get|put|delete]
```
Флаги:
- --id - для методов put и delete
- --number - для методов post и put
- --cardholder - для методов post и put
- --expiry - для методов post и put
- --cvv - для методов post и put
- --metadata - для методов post и put