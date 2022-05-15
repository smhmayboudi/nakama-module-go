CREATE SOURCE nakama
FROM KAFKA BROKER 'redpanda:29092' TOPIC 'nakama' FORMAT BYTES ENVELOPE NONE;
/*
 CREATE SOURCE "materialize"."public"."nakama"
 FROM KAFKA BROKER 'redpanda:29092' TOPIC 'nakama' FORMAT BYTES ENVELOPE NONE;
 */
CREATE MATERIALIZED VIEW nakama_view AS
SELECT data->>'b3' AS b3
FROM (
    SELECT CONVERT_FROM(data, 'utf8')::jsonb AS data
    FROM nakama
  );
/*
 CREATE VIEW "materialize"."public"."nakama_view" AS
 SELECT data->>'b3' AS b3
 FROM (
 SELECT "pg_catalog"."convert_from"("data", 'utf8')::"pg_catalog"."jsonb" AS "data"
 FROM "materialize"."public"."nakama"
 );
 */
CREATE SINK nakama_sink
FROM nakama_view INTO KAFKA BROKER 'redpanda:29092' TOPIC 'nakama_materialized' CONSISTENCY (
    TOPIC 'nakama_materialized_consistency' FORMAT AVRO USING CONFLUENT SCHEMA REGISTRY 'http://redpanda:8081'
  ) WITH (reuse_topic = true) FORMAT JSON;
/*
 CREATE SINK "materialize"."public"."nakama_sink" IN CLUSTER [1]
 FROM [u2 AS "materialize"."public"."nakama_view"] INTO KAFKA BROKER 'redpanda:29092' TOPIC 'nakama_materialized' CONSISTENCY (
 TOPIC 'nakama_materialized_consistency' FORMAT AVRO USING CONFLUENT SCHEMA REGISTRY 'http://redpanda:8081'
 ) WITH ("reuse_topic" = true) FORMAT JSON WITH SNAPSHOT;
 */
