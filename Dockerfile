FROM golang:1.19

WORKDIR /app/go

COPY . .

RUN go build -o /hello

CMD [ "/hello" ]