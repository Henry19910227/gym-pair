FROM golang AS gympair_member
WORKDIR /GYM-PAIR
COPY . /GYM-PAIR
RUN go build main.go 
EXPOSE 9090
ENTRYPOINT ./main

FROM mysql AS gympair_mysql
COPY ./gympair.sql /docker-entrypoint-initdb.d