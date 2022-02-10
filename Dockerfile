FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o binary app/*.go

ENV PORT=8080
EXPOSE ${PORT}

CMD [ "/app/binary" ]