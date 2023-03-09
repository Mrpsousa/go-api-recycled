# go-api-recycled
created a simple api with golang to training


## Run
    - docker-compose up -d
## DB steps 
    - docker-compose exec mysql bash
    - mysql -uroot -p products
    - pass: root
    - create table products (id varchar(255), name varchar(255), price float);

## Kafka steps
    - docker-compose exec broker bash
    - kafka-topics --bootstrap-server=localhost:9092 --topic=products --create

## Go steps 
    - docker-compose exec goapp bash
    - go run cmd/app/main.go

## Test - Postman
    - Post: http://localhost:8000/products
    {
        "name": "Nome-01",
        "price": 100
    }
    Get ....

## Test - Kafka
    - in kafka container
    - kafka-console-producer --bootstrap-server=localhost:9092 --topic=products
        -   {"name": "Nome-02","price": 200}