Сервис для определения ресурсов, необходимых задаче для выполнения

Порт выставляется через env переменную PORT. Значение по умолчанию - 8088, host - localhost

### Зависимости 

[`golang`](https://golang.org/doc/install)

Также можно запустить в [`docker`](https://docs.docker.com/install/)

### Запуск 

 - через golang - `go run .` - сервер запустится на порту 8088

 - docker-compose
 ```bash
docker-compose up
```
сервер запустится на 8088 порту

### Публичное API

* рассчитать необходимые ресурсы для запуска задачи

 `POST /resources`
```json
    {
      "userId": "sdfsdfj3d"
    }
```
Тело запроса может быть любым валидным json object    
    
Ответ:  

```HTTP/1.1 200 OK```
```json
     {
       "cpu":"1",
       "memory":"200Mi"
     }
 ```
