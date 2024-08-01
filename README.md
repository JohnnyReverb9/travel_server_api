## Быстрый сервер, который будет предоставлять Web-API для сервиса путешественников.

### В начальных данных для сервера есть три вида сущностей:
- User (Путешественник),
- Location (Достопримечательность),
- Visit (Посещения).

У каждой свой набор полей. Ниже можно скачать тестовый пример
(ammo - примеры запросов, answers - ответы на них, data - начальные данные),
а подробное описание есть в инструкции.

### Необходимо реализовать следующие запросы:

- GET /<entity>/<id> для получения данных о сущности
- GET /users/<id>/visits для получения списка посещений пользователем
- GET /locations/<id>/avg для получения средней оценки достопримечательности
- POST /<entity>/<id> на обновление
- POST /<entity>/new на создание