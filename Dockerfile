FROM golang:1.15-buster

# Set necessary environmet variables needed for the image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64\
    APP_HOME=/app

WORKDIR $APP_HOME

COPY go.mod go.sum ./

# Download all the dependencies
RUN go mod download

# Verify go.sum file matches what it downloaded
RUN go mod verify

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
