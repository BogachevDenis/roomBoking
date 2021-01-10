# roomBoking
# Запуск проекта:
Для запуска проекта выполните следующие команды:
<br>
$ git clone https://github.com/BogachevDenis/roomBoking.git
<br>
$ cd roomBoking
<br>
$ docker-compose build
<br>
$ docker-compose up
<br>
# Работа с приложением:
Приложение запустится на http://localhost:9000/
<br>
Запросы к API через curl:
<br>
POST Запрос на добавление нового номера
$ curl -X POST -d '{"price":200, "description":"TOP-room"}'  http://localhost:9000/room/add
<br>
DELETE  Запрос на удаление номера и всех его броней
$ curl -X DELETE  http://localhost:9000/room/119

