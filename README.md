# Event Driven example in Golang

This POC shows an example of event-driven architecture with a working domain event broker, an event producer and a
consumer.

The broker has a default buffer size of events of `3`, so if any of the consumers is not consuming the events fast
enough, **the broker drops the events to protect himself**. In a production-level environment, these events would
instead be saved to an event repository.

If you execute this example, you should see the producer creating events every second, the consumer consuming the events
every 2 seconds, and therefore making the broker drop these events after a while.

The main executable can be turned off at any time with CTRL-C to send an Interrupt signal and proceed with the graceful
termination.

## Execution

You can test this by executing the example with:

```
$ go run .
2021/11/01 11:11:20 producer: I have produced something [produced=0]
2021/11/01 11:11:21 producer: I have produced something [produced=1]
2021/11/01 11:11:22 consumer: the producer created something [produced=0]
2021/11/01 11:11:22 producer: I have produced something [produced=2]
2021/11/01 11:11:23 producer: I have produced something [produced=3]
2021/11/01 11:11:24 consumer: the producer created something [produced=1]
2021/11/01 11:11:24 producer: I have produced something [produced=4]
2021/11/01 11:11:25 producer: I have produced something [produced=5]
2021/11/01 11:11:26 producer: I have produced something [produced=6]
2021/11/01 11:11:26 consumer: the producer created something [produced=2]
2021/11/01 11:11:26 broker: warning, event dropped [type=*producer.ProducedSomethingEvent, ID=605394647632969758, happenedOn=2021-11-02 11:11:26.956468166 +0100 CET m=+7.002266259]
2021/11/01 11:11:27 producer: I have produced something [produced=7]
2021/11/01 11:11:28 consumer: the producer created something [produced=3]
2021/11/01 11:11:28 producer: I have produced something [produced=8]
2021/11/01 11:11:29 quitting
```

On signal interruption with graceful shutdown:

```
$ go run .
2021/11/01 11:12:28 producer: I have produced something [produced=0]
2021/11/01 11:12:29 producer: I have produced something [produced=1]
2021/11/01 11:12:30 consumer: the producer created something [produced=0]
2021/11/01 11:12:30 producer: I have produced something [produced=2]
^C2021/11/01 11:12:30 cancelling from signal [signal=interrupt]
2021/11/01 11:12:30 quitting
```
