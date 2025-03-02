FROM golang:latest as builder
WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

LABEL env="development"

ENTRYPOINT [ "./main" ]


# FROM scratch

# COPY --from=builder /app/main .

# CMD ["./main"]
