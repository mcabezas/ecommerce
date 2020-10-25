FROM golang:1.13

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN make build

CMD ./api

EXPOSE 9092
