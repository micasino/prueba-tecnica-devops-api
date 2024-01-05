FROM --platform=linux/arm64/v8 golang:1.21  as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY cmd/ ./cmd
COPY pkg/ ./pkg

RUN go build -o ./build/bin -v ./cmd/api

FROM --platform=linux/arm64/v8 debian:stable-slim as final
COPY --from=builder /app/build /app/build
COPY .env /app/build/
WORKDIR /app/build/
CMD [ "./bin" ]
