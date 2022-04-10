# go_rabbitmq_example

## Архитектура
Клиенты с помощью grpc вызовов отправляют запросы Producer-у (он в системе 1). После этого Producer добавляет в очередь RabbitMQ запрос клиента, какой-то Worker обрабатывает запрос и с помощью grpc вызова возвращает ответ Producer-у, который теперь может завершить grpc вызов клиента и вернуть ему ответ от Worker-a.

## Варианты запуска приложения:
### Ручная сборка с помощью go и запуск rabbitMQ через Docker
```docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.9-management```

```cd producer && go build . && ./producer -mqaddr={адрес rabbitmq, например amqp://guest:guest@localhost:5672}```

В другом терминале / на другой машине (можно выполнять столько раз, сколько нужно worker-ов): 

```cd worker && go build . && ./worker -addr={адрес Producer-а, например localhost:50051} -mqaddr={адрес rabbitmq, например amqp://guest:guest@localhost:5672}```

Клиент:

```cd client && go build . && ./client {from} {to} {lang} {titles} {addr}```

```from, to``` - начальная и конечная страница, между которыми вычисляется расстояние

```titles``` - если ```=true```, то worker-ы не в ручную получают список ссылок на каждой странице, а используют API википедии и получают по названию статьи все названия статей, куда из неё можно попасть за 1 клик. В этом случае ```from, to``` должны быть названиями статей, а не ссылками на них.

```lang``` - в случае, если ```titles=true``` отвечает за язык википедии, на которой мы ищем статьи (например en, или ru)

```addr``` - адрес Producer-а, например ```localhost:50051```

Клиент делает один запрос к Producer-у, дожидается ответа, и затем завершает работу.

### Локальный запуск через Docker (всё на localhost-е)
```docker network create test-net```

```docker run -it --rm --name rabbitmq --network test-net -p 5672:5672 -p 15672:15672 rabbitmq:3.9-management```

```docker run -i --network test-net --name producer1 asmorodinov/hw5_producer -mqaddr=amqp://guest:guest@rabbitmq:5672/ ```

```sudo docker run -i --network test-net --name worker3 asmorodinov/hw5_worker -addr=producer1:50051 -mqaddr=amqp://guest:guest@rabbitmq:5672/```

```docker run -i --network test-net --name client2 asmorodinov/hw5_client Computer 16-bit en true producer1:50051```

### Запуск серверной части в docker-compose, клиента в докере
```docker-compose build && docker-compose up```
```docker run -i --network hw5_dev-network --name client2 asmorodinov/hw5_client Computer 16-bit en true producer:50051```

### Запуск всех узлов в docker-compose
```docker-compose -f docker-compose-2.yaml build && docker-compose -f docker-compose-2.yaml up ```

### Другие варианты сборки
Очевидно, что также есть много других возможных вариантов сборки приложения. В частности, ничего не мешает запускать отдельные узлы как в докере, так и без него, на разных машинах, или на одной, с помощью docker-compose, или без него. Все эти случаи реализуются несложно благодаря аргументам ```port```, ```addr``` и ```mqaddr``` у команд.
