FROM --platform=linux/amd64/v3 golang:1.21  as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY cmd/ ./cmd
COPY pkg/ ./pkg

RUN go build -o ./build/bin -v ./cmd/api

FROM --platform=linux/amd64/v3 debian:stable-slim as final
COPY --from=builder /app/build /app/build
COPY .env /app/build/
WORKDIR /app/build/
CMD [ "./bin" ]
