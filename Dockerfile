FROM golang:1.16-alpine AS builder
WORKDIR /app
COPY src/go.mod .
COPY src/go.sum .
RUN go mod download

COPY src/ .
RUN go build -o api

FROM dtzar/helm-kubectl:3.6.3
WORKDIR /home
COPY --from=builder /app/api .

CMD ./api $PORT $URL_PATH_PREFIX
