import { Kafka } from "kafkajs";
import { traceSpan } from "./opentelemetry";

const kafka = new Kafka({
  clientId: "server",
  brokers: ["127.0.0.1:9092"]
});

const producer = kafka.producer();

export function getConnection(user: string) {
  return producer.connect().then(() => {
    return (message: string) => {
      return traceSpan("server", () => {
        return producer.send({
          topic: "nakama",
          messages: [
            { value: JSON.stringify({ message, user }) }
          ]
        })
      }, "producer")
    }
  })
}

export function disconnect() {
  return producer.disconnect()
}
