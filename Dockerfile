FROM golang AS gympair_api
WORKDIR /GYM-PAIR
COPY . /GYM-PAIR
RUN go build main.go 
EXPOSE 9090
ENTRYPOINT ./main -m release

FROM mysql AS gympair_mysql
COPY ./gympair.sql /docker-entrypoint-initdb.d