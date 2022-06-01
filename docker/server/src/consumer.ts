import logger from "./logger";
import { B3_CONTEXT_HEADER } from "@opentelemetry/propagator-b3";
import { Kafka } from "kafkajs";
import { context, propagation, defaultTextMapGetter } from "@opentelemetry/api";
import { traceSpan } from "./opentelemetry";
import { v4 as uuidv4 } from "uuid";

const kafka = new Kafka({
  clientId: "server",
  brokers: ["127.0.0.1:9092"]
});

const consumer = kafka.consumer({ groupId: uuidv4() });

export function connect() {
  return consumer.connect().then(() => {
    consumer.subscribe({ topic: "nakama" }).then(() => {
      consumer.run({
        eachMessage: async ({ topic, partition, message }) => {
          const formattedValue = JSON.parse((message.value as Buffer).toString());
          var carrier: { [key: string]: string } = { [B3_CONTEXT_HEADER]: formattedValue.b3 };
          context.with(propagation.extract(context.active(), carrier, defaultTextMapGetter), () => {
            traceSpan("server", () => {
              logger.debug(formattedValue)
            }, "producer")
          })
        }
      })
    })
  })
};

export function disconnect() {
  consumer.disconnect()
};
