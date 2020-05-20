# docker build -t tynmarket/mysql:8.0 .
# docker run -d -p 3306:3306 --env MYSQL_ROOT_PASSWORD --env MYSQL_DATABASE --env MYSQL_USER --env MYSQL_PASSWORD --name mysql tynmarket/mysql:8.0

FROM mysql:8.0

RUN apt-get update \
      && apt-get install -y --no-install-recommends \
      && rm -rf /var/lib/apt/lists/*

ENV TZ Asia/Tokyo

COPY imarket.cnf /etc/mysql/conf.d
