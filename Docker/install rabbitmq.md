# Cài đặt rabbitmq bằng docker

### Trên host tạo thư mục data và log để chứa log và data của container

mkdir /home/docker/rabbitmq/data/
mkdir /home/docker/rabbitmq/log/

- File docker-compose.yml

```

version: '3.7'

services:

  rabbitmq:

    image: rabbitmq:3-management-alpine

    container_name: 'rabbitmq'

    ports:

      - 5672:5672

      - 15672:15672

    environment:

      RABBITMQ_DEFAULT_USER: admin

      RABBITMQ_DEFAULT_PASS: admin@123

    volumes:

      - /home/docker/rabbitmq/data/:/var/lib/rabbitmq/

      - /home/docker/rabbitmq/log/:/var/log/rabbitmq

    networks:

      - rabbitmq_go_net

networks:

  rabbitmq_go_net:

    driver: bridge


```
