# How Kafka Works

Apache Kafka is a powerful distributed streaming platform that allows you to publish and subscribe to streams of records. It's designed for handling large volumes of real-time data efficiently and reliably. This section provides a descriptive structure of how Kafka works, highlighting key components and their interactions.

## Overview

At its core, Kafka is built around the concept of topics, where data is organized and distributed. Here's how Kafka works in a nutshell:

1. **Producers**:
   - Data producers send records to Kafka topics.
   - Records can be logs, events, or any form of data that needs to be distributed.
   
2. **Topics**:
   - Kafka organizes data into topics, acting as logical channels for data.
   - Producers publish records to specific topics.

3. **Brokers**:
   - Kafka brokers are servers that manage the data.
   - They store and distribute records in topics.
   - Kafka clusters consist of multiple brokers for redundancy and scalability.

4. **Partitions**:
   - Each topic can be split into multiple partitions, enabling parallelism and scalability.
   - Partitions distribute the data across brokers.

5. **Replication**:
   - Kafka uses replication to ensure fault tolerance.
   - Each partition has multiple replicas distributed across different brokers.
   
6. **Consumers**:
   - Data consumers subscribe to topics and consume records.
   - They can specify where they left off using consumer groups.

7. **Consumer Groups**:
   - Consumers are organized into consumer groups.
   - Each group reads data from one or more partitions, ensuring that data processing occurs in parallel.

## Key Features

- **Durability**:
   - Kafka stores data for a configurable period.
   - This ensures that data is available for consumers even if they join the cluster at a later time.

- **Scalability**:
   - Kafka is designed for scaling.
   - You can add more brokers, partitions, and consumers to handle increasing data loads.

- **Reliability**:
   - Kafka is known for its durability and reliability.
   - Data is replicated, and brokers can recover from failures without data loss.

- **Real-time Streaming**:
   - Kafka provides low-latency, real-time data streaming.
   - It's suitable for use cases like event sourcing, log aggregation, and stream processing.

## Conclusion

Kafka's distributed, fault-tolerant, and high-throughput design makes it a robust choice for building data pipelines, real-time applications, and stream processing at scale. Understanding how Kafka works is fundamental for anyone dealing with real-time data streams.

# Apache Kafka Producer and Consumer in Go

This repository contains a simple example of how to create a Kafka producer and consumer in the Go programming language using the [Sarama](https://pkg.go.dev/github.com/Shopify/sarama) library.

## Prerequisites

Before you can run the Kafka producer and consumer, you'll need to have the following dependencies installed:

- [Apache Kafka](https://kafka.apache.org/downloads)
- Go 1.21.1
- The Sarama library (`github.com/IBM/sarama`)

You can install the Sarama library using `go get`:

```shell
go get github.com/IBM/sarama
```

## Usage

### Kafka Setup

1. Download and start a Kafka server by following the [official documentation](https://kafka.apache.org/quickstart).

    OR

   a. Download Apache Kafka from the official website: [Apache Kafka](https://kafka.apache.org/downloads).

   b. Extract the downloaded archive to a directory of your choice.

   c. Go to that directory in Command Prompt.
   
   e. Start the ZooKeeper server:

      ```shell
      .\bin\windows\zookeeper-server-start.bat .\config\zookeeper.properties
      ```
      
   f. Open a new Command Prompt and go to that directory
   
   g. Start the Kafka server:
   
      ```shell
      .\bin\windows\kafka-server-start.bat .\config\server.properties
      ```

2. Create a topic for the producer and consumer. Replace `sample_topic` with the desired topic name.

   a. Go to the directory where you installed Kafka.

   b. Open a terminal and go to the `.bin\windows` directory (assuming you have a Windows device).

   c. Execute the following command to create a topic named `sample_topic`:
   
   ```shell
   kafka-topics.bat --create --bootstrap-server localhost:9092 --topic sample_topic
   ```

### Producer

The `sendMessage.go` file in this repository contains a simple Kafka producer that sends messages to the Kafka topic. You need to modify the Kafka broker and topic information as needed.

To run the producer:

```shell
go run .\producer\sendMessage.go
```

### Consumer

The `readMessage.go` file in this repository contains a Kafka consumer that reads messages from the Kafka topic. You also need to update the Kafka broker and topic information.

To run the consumer:

```shell
go run .\consumer\readMessage.go
```

## Configuration

Both the producer and consumer can be configured by modifying the respective Go source files. You can adjust settings like the Kafka broker address, topic name, and other Kafka parameters according to your setup.


## Author

- Sunil Pradhan
- GitHub: [nil2510](https://github.com/nil2510)

Feel free to reach out with any questions or improvements.
 
