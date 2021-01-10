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
$ curl -X POST http://localhost:8080/create -d '{"email": "test@test.ru","url":"https://www.avito.ru/ad"}'
