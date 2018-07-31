# kafka-docker

Kafka docker version of a single node

## Getting Started

Configuration was taken from confluent's [single node example](https://github.com/confluentinc/cp-docker-images/tree/master/examples)

### Installing
Set `HOST_IP` environment variable with your local computer IP address

    HOST_IP=10.0.1.34 docker-compose up -d

or

    export HOST_IP=10.0.1.34
    docker-compose up -d

## Running
There should be two containers (one is zookeeper, the other is kafka server/broker) running in docker.

Consumer and Producer can now connect to kafka cluster using the you local computer's IP at port `29092`

    bootstrapservers: 10.0.1.34:29092

## Kafka Commands

[Kafka commands](https://kafka.apache.org/quickstart) are executable through kafka container

**List Topics**

    docker exec -it kafka-docker_kafka_1 usr/bin/kafka-topics --list --zookeeper $HOST_IP:32181

**Create Topic**

    docker exec -it kafka-docker_kafka_1 usr/bin/kafka-topics --create --zookeeper $HOST_IP:32181 --replication-factor 1 --partitions 2 --topic name-your-topic

**Increase partitions of existing topic**

    docker exec -it kafka-docker_kafka_1 usr/bin/kafka-topics --zookeeper $HOST_IP:32181 --alter --topic yourTopic --partitions 4

**Change Topic Retention Time**

    docker exec -it kafka-docker_kafka_1 usr/bin/kafka-topics --zookeeper $HOST_IP:32181 --alter --topic MyTopic --config retention.ms=1000


