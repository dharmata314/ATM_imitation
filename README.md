
# ATM

Имитация работы банкомата 

## Примеры запросов

Создание нового аккаунта:
```
curl -X POST \
    http://localhost:8080/accounts
```
Пополнение баланса:
```
curl -X POST \
    -H "Content-Type: application/json" \
    -d '{"amount": 50}' \
    http://localhost:8080/accounts/{id}/deposit
```
Снятие средств:
```
curl -X POST \
    -H "Content-Type: application/json" \
    -d '{"amount": 50}' \
    http://localhost:8080/accounts/{id}/withdraw
```
Проверка баланса:
```
curl -X GET \
    http://localhost:8080/accounts/{id}/balance
```
