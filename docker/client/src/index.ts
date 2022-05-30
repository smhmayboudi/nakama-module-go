import logger from "./logger";
// import { B3_CONTEXT_HEADER } from "@opentelemetry/propagator-b3";
import { Client } from "@heroiclabs/nakama-js";
import { context, propagation, trace } from "@opentelemetry/api";
import { traceSpan } from "./opentelemetry";

const useSSL = false; // Enable if server is run with an SSL certificate.
const client = new Client("defaultkey", "127.0.0.1", "7350", useSSL);
logger.debug("Client created");
const prepareClickEvent = () => {
  const element = document.getElementById("button1");
  if (element === null) {
    return;
  }
  const onClick = () => {
    traceSpan("authenticateEmail", async () => {
      const span = trace.getSpan(context.active());
      var carrier: { [key: string]: string } = {};
      propagation.inject(context.active(), carrier, {
        set: (carrier, key, value) => {
          carrier[key] = value;
        }
      });
      const email = "smhmayboudi@gmail.com";
      const password = "dY!6HY8Kz8DAJ4n";
      const create = true;
      const username = "smhmayboudi";
      // const vars = new Map([[B3_CONTEXT_HEADER, carrier[B3_CONTEXT_HEADER]]]);
      const vars = new Map([["authenticate_email_key", "authenticate_email_value"]]);
      const session = await client.authenticateEmail(email, password, create, username, vars);
      logger.debug("AuthenticateEmail called");
      traceSpan("emitEvent", async () => {
        await client.emitEvent(session, {
          external: false,
          name: "NAME",
          properties: new Map([["emit_event_key", "emit_event_value"]]),
          timestamp: "1970-01-01T00:00:00Z",
        })
        logger.debug("EmitEvent called");
      }, "emitEvent");
    }, "authenticateEmail");
  };
  element.addEventListener("click", onClick);
};
window.addEventListener("load", prepareClickEvent);
