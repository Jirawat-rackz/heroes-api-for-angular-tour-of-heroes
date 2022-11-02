FROM golang:1.19.3-alpine3.16 AS builder

WORKDIR /app

COPY . ./
RUN go mod download

RUN cd cmd/heroes && go build -o heroes-api

FROM golang:1.19.3-alpine3.16 as runner
WORKDIR /app
COPY --from=builder /app/cmd/heroes/heroes-api /app

EXPOSE 8080
ENTRYPOINT [ "./heroes-api" ]