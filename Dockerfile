FROM golang:1.19

WORKDIR /app/go

RUN apt-get update && apt-get install make

COPY . .

CMD make run