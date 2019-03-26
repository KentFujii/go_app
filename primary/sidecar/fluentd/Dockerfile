FROM fluent/fluentd:v0.12-debian
ADD fluent.conf /fluentd/etc
RUN buildDeps="make gcc g++ libc-dev ruby-dev" \
  && apt-get update \
  && apt-get install -y --no-install-recommends $buildDeps \
  && gem install fluent-plugin-prometheus
