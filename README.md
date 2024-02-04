# Тестовое задание для стажера на позицию «Go-разработчик»
Необходимо разработать приложение EWallet реализующее систему обработки транзакций платёжной системы. Приложение должно быть реализовано в виде HTTP сервера, реализующее REST API. Сервер должен реализовать 4 метода и их логику

# Запуск в Docker
docker-compose up --build infotecs

# Эндпоинты
Приложение слушает на 8000 порту (localhost).

## Создание кошелька
Пример запроса в консоли:
```bash
curl -X POST http://localhost:8000/api/v1/wallet
```
<div align="center">
  <img src="https://github.com/TerreDHermes/TerreDHermes/blob/main/assets/post_create.png" alt="Описание изображения" style="width: 70%;">
</div>
## Перевод средств с одного кошелька на другой
Пример запроса в консоли:
```bash
curl -X POST -d "{\"to\":1,\"amount\":50}" http://localhost:8000/api/v1/wallet/2/send
```
## Получение историй входящих и исходящих транзакций
Пример запроса в консоли:
```bash
curl -X GET http://localhost:8000/api/v1/wallet/2/history
```
## Получение текущего состояния кошелька
Пример запроса в консоли:
```bash
curl -X GET http://localhost:8000/api/v1/wallet/2/
```
