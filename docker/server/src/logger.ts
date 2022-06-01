import pino from "pino";

const logger = pino({ level: "debug", useLevelLabels: true });

export default logger