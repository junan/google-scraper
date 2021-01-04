FROM golang:1.15-buster

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

ENTRYPOINT ["make", "production"]
