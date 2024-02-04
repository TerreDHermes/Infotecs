# Тестовое задание для стажера на позицию «Go-разработчик»
Разработано приложение EWallet реализующее систему обработки транзакций платёжной системы. Приложение представляет из себя HTTP сервер, реализующий REST API. У сервера есть 4 метода и их логика.
База данных - Postgres.
Основной фреймворк - Gin Web Framework.

# Запуск в Docker
Команда для сборки и запуска из директории приложения:
```bash
docker-compose up --build infotecs
```
Приложение запустится успешно с первого раза, так как добавлен скрипт wait-for-postgres.sh, который не даёт запуститься серверу до запуска базы данных.

# Эндпоинты
Приложение слушает на 8000 порту (localhost).

## Создание кошелька
Пример запроса в консоли:
```bash
curl -X POST http://localhost:8000/api/v1/wallet
```
Тут выполняется создание кошелька. Ответ содержит JSON-объект с состоянием созданного кошелька. 
Аналогичный запрос в Postman:

<div align="center">
  <img src="https://github.com/TerreDHermes/TerreDHermes/blob/main/assets/post_create.png" alt="Описание изображения" style="width: 70%;">
</div>

## Перевод средств с одного кошелька на другой
Пример запроса в консоли:
```bash
curl -X POST -d "{\"to\":1,\"amount\":50}" http://localhost:8000/api/v1/wallet/2/send
```
Выполняется перевод с одного кошелька на другой (со 2 на 1). Статус ответа 200 если перевод успешен, 404 если исходящий кошелек не найден, 400 если целевой кошелек не найден или на исходящем нет нужной суммы. Также добавлена обработка случая, когда исходящий кошелек такой же как и исходящий. А сумма перевода не может быть отрицательной. 
Аналогичный запрос в Postman, только исходящий 1, а у целевого id=2:

<div align="center">
  <img src="https://github.com/TerreDHermes/TerreDHermes/blob/main/assets/send.png" alt="Описание изображения" style="width: 70%;">
</div>

## Получение историй входящих и исходящих транзакций
Пример запроса в консоли:
```bash
curl -X GET http://localhost:8000/api/v1/wallet/2/history
```
Отображение историй входящих и исходящих транзакций. Ответ с статусом 200 если кошелек найден, 404 если указанный кошелек не найден. Ответ содержит в теле массив JSON-объектов с входящими и исходящими транзакциями кошелька. Пример запроса в Postman для кошелька с id=1:

<div align="center">
  <img src="https://github.com/TerreDHermes/TerreDHermes/blob/main/assets/His.png" alt="Описание изображения" style="width: 70%;">
</div>

## Получение текущего состояния кошелька
Пример запроса в консоли:
```bash
curl -X GET http://localhost:8000/api/v1/wallet/2/
```
Отображение текущего состояния кошелька. Ответ с статусом 200 если кошелек найден, 404 если кошелек не найден. Ответ содержит в теле JSON-объект с текущим состоянием кошелька.
Пример запроса в Postman для кошелька с id=1:

<div align="center">
  <img src="https://github.com/TerreDHermes/TerreDHermes/blob/main/assets/get_info_wallet.png" alt="Описание изображения" style="width: 70%;">
</div>
