# Тестовое задание для стажера на позицию «Go-разработчик»
Необходимо разработать приложение EWallet реализующее систему обработки транзакций платёжной системы. Приложение должно быть реализовано в виде HTTP сервера, реализующее REST API. Сервер должен реализовать 4 метода и их логику

# Запуск в Docker
Команда для сборки и запуска из директории приложения:
```bash
docker-compose up --build infotecs
```
Приложение запуститься сразу успешно, так как добавлен скрипт wait-for-postgres.sh, который не даёт запуститься серверу до запуска базы данных.

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

<div align="center">
  <img src="https://github.com/TerreDHermes/TerreDHermes/blob/main/assets/send.png" alt="Описание изображения" style="width: 70%;">
</div>

## Получение историй входящих и исходящих транзакций
Пример запроса в консоли:
```bash
curl -X GET http://localhost:8000/api/v1/wallet/2/history
```

<div align="center">
  <img src="https://github.com/TerreDHermes/TerreDHermes/blob/main/assets/His.png" alt="Описание изображения" style="width: 70%;">
</div>

## Получение текущего состояния кошелька
Пример запроса в консоли:
```bash
curl -X GET http://localhost:8000/api/v1/wallet/2/
```

<div align="center">
  <img src="https://github.com/TerreDHermes/TerreDHermes/blob/main/assets/get_info_wallet.png" alt="Описание изображения" style="width: 70%;">
</div>
