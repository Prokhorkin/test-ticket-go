 1. Локальный запуск приложения - go run main.go
    - Приложение запускается на локальном хосте, порт 5505 
 2. Http API приложения:
   - Методы Get
     - /ping - запрос доступности приложения, возвращает код 200
        - Пример запроса из консоли - curl localhost:5505/ping
     - /add - запрос на расчет суммы двух числовых параметров, переданных в запросе   
         - Пример запроса из консоли - curl localhost:5505/add?a=6&b=8
     - /sub - запрос на расчет разности двух числовых параметров, переданных в запросе
         - Пример запроса из консоли - curl localhost:5505/sub?a=8&b=6
     - /mul - запрос на расчет умножения двух числовых параметров, переданных в запросе
         - Пример запроса из консоли - curl localhost:5505/mul?a=2&b=2
     - /div - запрос на расчет деления двух числовых параметров, переданных в запросе
         - Пример запроса из консоли - curl localhost:5505/div?a=6&b=3
 3. Сборка образа Docker - docker build -t test-ticket-go:0.1 . 
 4. Запуск приложения в docker-контейнере - docker run -p 5505:5505 test-ticket-go:0.1
     - Приложение запускается в docker-контейнере, отражая порт 5505 во вне
 5. В качестве клиента можно использовать сервис из директории проекта tmp/client.go   
 6. Остановка приложения - сочетания клавиш Ctrl+C
   