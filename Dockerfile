FROM  golang:1.21  as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY cmd/ ./cmd
COPY pkg/ ./pkg

RUN go build -o ./build/bin -v ./cmd/api

FROM debian:stable-slim as final
COPY --from=builder /app/build /app/build
COPY .env /app/build/
WORKDIR /app/build/
RUN apt-get update && \
    apt-get install -y wget
EXPOSE 3000

CMD [ "./bin" ]
