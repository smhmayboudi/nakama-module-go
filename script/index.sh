#!/bin/bash

set -eux
set -o errexit
set -o pipefail
set -o nounset

# docker exec -it redpanda rpk topic create nakama --brokers=localhost:9092
# docker exec -it redpanda rpk topic produce nakama --brokers=localhost:9092
# docker exec -it redpanda rpk topic consume nakama --brokers=localhost:9092

# [Help]()
# docker exec -it redpanda redpanda --help

# [Help](https://docs.redpanda.com/docs/reference/rpk-commands/)
# docker exec -it redpanda rpk --help

# rpk generate grafana-dashboard \
#   --datasource prometheus \
#   --metrics-endpoint 127.0.0.1:9644/metrics > redpanda-dashboard.json

# Monitor materialize
# docker run -d \
#     -v /tmp/prom-data:/prometheus -u "$(id -u):$(id -g)" \
#     -p 3000:3000 -e MATERIALIZED_URL=127.0.0.1:6875 \
#     materialize/dashboard

# Prometheus materialize
# http://127.0.0.1/metrics

# Health Check materialize
# http://127.0.0.1:6875/status

# Memory usage visualization
# http://127.0.0.1:6875/memory

# Grafana
# https://github.com/MaterializeInc/materialize/blob/v0.26.0/misc/monitoring/dashboard/conf/grafana/dashboards/overview.json

# Prometheus nakama
# https://heroiclabs.com/docs/nakama/getting-started/install/docker/
