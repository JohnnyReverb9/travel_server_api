Методы выборки данных (GET):
1. Получение данных о сущности: /{entity}/{id}
   {id} - строка, которая может принимать любые значения.

В ответе ожидается код 404, если сущности с таким идентификатором нет в данных. Иначе, все собственные поля, включая идентификатор. {entity} принимает одно из значений - users, locations или visits.

Пусть пользователь с id = 1 существует. Пример ответа на запрос: GET: /users/1

HTTP Status Code: 200

    {
        "id": 1,
        "email": "johndoe@gmail.com",
        "first_name": "John",
        "last_name": "Doe",
        "gender": "m",
        "birth_date": -1613433600
    }
Пример ответа на запрос: GET: /users/string

HTTP Status Code: 404

Пример ответа на запрос: GET: /users/string/somethingbad

HTTP Status Code: 404

Пример ответа на запрос: GET: /users/

HTTP Status Code: 404

Пример ответа на запрос: GET: /user/

HTTP Status Code: 404

И так далее…

2. Получение списка мест, которые посетил пользователь: /users/{id}/visits
   В теле ответа ожидается структура {"visits": [ ... ]}, отсортированная по возрастанию дат, или ошибка 404/400. Подробнее - в примерах.
   GET-параметры:

fromDate - посещения с visited_at } fromDate
toDate - посещения до visited_at { toDate
country - название страны, в которой находятся интересующие достопримечательности
toDistance - возвращать только те места, у которых расстояние от города меньше этого параметра
Пусть пользователь с id = 1 существует. Пример ответа на запрос: GET: /users/1/visits

HTTP Status Code: 200

    {
        "visits": [
            {
                "mark": 2,
                "visited_at": 958656902,
                "place": "Кольский полуостров"
            },
            {
                "mark": 4,
                "visited_at": 1223268286,
                "place": "Московский Кремль"
            }
         ]
    }
Пример ответа на запрос: GET: /users/1/visit

HTTP Status Code: 404

Пример ответа на запрос: GET: /users/1/visits?fromDate=

HTTP Status Code: 400

Пример ответа на запрос: GET: /users/1/visits?fromDate=abracadbra

HTTP Status Code: 400

Пример ответа на запрос: GET: /users/somethingstringhere/visits?fromDate=1

HTTP Status Code: 404

Пусть пользователь с id = 1 существует. Пример ответа на запрос: GET: /users/1/visit?fromDate=915148800&toDate=915148800

HTTP Status Code: 404

В этом примере ошибка в visit.

Пусть пользователь с id = 1 существует. Пример ответа на запрос: GET: /users/1/visits?country=?

HTTP Status Code: 400

В случае если пользователя с переданным id нет - отдавать 404. Если просто нет посещений, то {"visits": []}

3. Получение средней оценки достопримечательности: /locations/{id}/avg
   В ответе ожидается одно число, с точностью до 5 десятичных знаков (округляется по стандартным математическим правилам округления(round)), либо код 404.

GET-параметры:

fromDate - учитывать оценки только с visited_at } fromDate
toDate - учитывать оценки только до visited_at { toDate
fromAge - учитывать только путешественников, у которых возраст (считается от текущего timestamp) строго больше этого параметра
toAge - учитывать только путешественников, у которых возраст (считается от текущего timestamp) строго меньше этого параметра
gender - учитывать оценки только мужчин или женщин
Пример ответа на запрос: GET: /locations/1/avg

    {
        "avg": 3.43
    }
Пример ответа на запрос: GET: /locations/somethingsomething/avg

HTTP Status Code: 404

В случае если места с переданным id нет - отдавать 404. Если по указанным параметрам не было посещений, то {"avg": 0}

Небольшой пример проверки дат в этом запросе на python (fromAge - количество лет):

        from datetime import datetime
        from dateutil.relativedelta import relativedelta
        import calendar

        now = datetime.now() - relativedelta(years = fromAge)
        timestamp = calendar.timegm(now.timetuple())
Дальше проверяется birthdate { timestamp либо birthdate } timestamp соответственно.

Методы обновления данных (POST)
1. Обновление данных о сущности: /{entity}/{id}
   В ответе ожидается код 200 с пустым json-ом в теле ответа {}, если обновление прошло успешно, 404 - если запись не существовала в данных или 400, если в теле запроса некорректные данные.

Только обновляемые поля и их значения содержатся в теле запроса в формате JSON. id никогда не содержится среди обновляемых полей.

Пусть пользователь с id = 214 существует. Пример тела запроса: POST: /users/214

    {
        "email": "johndoe@gmail.com",
        "first_name": "Jessie",
        "last_name": "Pinkman",
        "birth_date": 616550400
    }    
Ответ:
HTTP Status Code: 200

{}
Пусть пользователь с id = 214 также существует. Изменено поле с почтой. Пример тела запроса: POST: /users/214

    {
        "email": null,
        "first_name": "Jessie",
        "last_name": "Pinkman",
        "birth_date": 616550400
    }    
Ответ:

HTTP Status Code: 400

Пусть пользователь с id = 214 не существует. Пример тела запроса: POST: /users/214

    {
        "email": "test@gmail.com",
        "first_name": "Jessie",
        "last_name": "Pinkman",
        "birth_date": 616550400
    }    
Ответ:

HTTP Status Code: 404

2. Добавление новой сущности: /{entity}/new
   В ответе ожидается код 200 с пустым json-ом в теле ответа ("{}"), если создание прошло успешно. В случае некорректных данных - код 400.

Обновляемые поля и их значения содержатся в теле запроса в формате JSON. В таких запросах может быть GET-параметр queryId, который надо игнорировать.

При создании сущности все поля являются обязательными.

В случае попытки создания сущности с id, уже существующим в текущих данных, ожидается код ошибки 400.

Пример запроса POST: /users/new

    {
        "id": 245,
        "email": "foobar@mail.ru",
        "first_name": "Маша",
        "last_name": "Пушкина",
        "gender": "f",
        "birth_date": 365299200
    }
Ответ:

HTTP Status Code: 200

{}
Пример запроса POST: /users/new

    {
        "id": 245,
        "email": null,
        "first_name": "Маша",
        "last_name": "Пушкина",
        "gender": "f",
        "birth_date": 365299200
    }
Ответ:

HTTP Status Code: 400