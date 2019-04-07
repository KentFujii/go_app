FROM fluent/fluentd:v1.3.2-debian
RUN buildDeps="make gcc g++ libc-dev ruby-dev" \
  && apt-get update \
  && apt-get install -y --no-install-recommends $buildDeps \
  && gem install fluent-plugin-prometheus
