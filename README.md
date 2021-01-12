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
<li>POST Запрос на добавление нового номера
  <br>
$ curl -X POST -d '{"price":200, "description":"TOP-room"}'  http://localhost:9000/room/add
<br>
<li>DELETE  Запрос на удаление номера и всех его броней
  <br>
$ curl -X DELETE  http://localhost:9000/room/119
<br>
<li>GET  Запрос на получение списка номеров отеля, где параметр date - сортировка по дате, price - сортировка по цене, direction - направление сортировки
  <br>
$ curl -X GET  http://localhost:9000/room?date=false&price=true&direction=false
<br>
<br>
<li>POST Запрос на добавление новой брони
  <br>
$ curl -X POST -d '{"roomid":6, "startdate":"2021-12-25", "finishdate":"2022-11-11"}'  http://localhost:9000/booking/add
<br>
<li>DELETE  Запрос на удаление брони
  <br>
$ curl -X DELETE  http://localhost:9000/booking/68
<br>
<li>GET  Запрос на получение списка броней, где параметр roomid - id номера
<br>
$ curl -X GET  http://localhost:9000/booking?roomid=6
 <br>
# Запуск тестов
  $ cd pkg/handlers/
$ go test
