FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o word_counter .

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/word_counter .
COPY input_files ./input_files
ENTRYPOINT ["./word_counter"]
