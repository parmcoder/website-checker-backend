FROM golang:1.19

WORKDIR /app/go

RUN apt-get update && apt-get install make

COPY . .

EXPOSE 3000

CMD make run