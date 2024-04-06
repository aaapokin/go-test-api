# Простое API  
Заготовка для будущего приложения REST API  
`make drop-all-container` - удалить все контейнеры, чтобы избежать конфликта  
`make up` - запустить приложение  
`make migrate-up` - запустить миграцию  
`http://localhost/api/v1/ping` - публичный   
`http://localhost/api/v1/ping2` - защищенный. В заголовке нужно передать токен. `Authorization Bearer ******`  
`POST http://localhost/api/v1/auth` - `{"login":"test","password":"password"}` - получаем токен   

@TODO нужно добавить тесты и работу с горутинами и каналами. Прикрутить функционал работы с RMQ и внешними API  