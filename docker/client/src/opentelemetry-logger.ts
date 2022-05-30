import logger from "./logger";
import type { DiagLogFunction } from "@opentelemetry/api";

const error: DiagLogFunction = (message, ...args) => {
    logger.error(message, ...args);
  },
  warn: DiagLogFunction = (message, ...args) => {
    logger.warn(message, ...args);
  },
  info: DiagLogFunction = (message, ...args) => {
    logger.info(message, ...args);
  },
  debug: DiagLogFunction = (message, ...args) => {
    logger.debug(message, ...args);
  },
  verbose: DiagLogFunction = (message, ...args) => {
    logger.trace(message, ...args);
  },
  OpentelemetryLogger = { debug, error, info, verbose, warn };

export { OpentelemetryLogger };
