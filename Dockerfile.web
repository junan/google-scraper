FROM drewwells/wellington as assets-builder

WORKDIR /app

# Copy assets folder
COPY assets/. ./assets/

# Convert scss to css and minify it
RUN wt compile assets/stylesheets/application.scss -s compressed -b static/css

FROM golang:1.15-buster

ARG DATABASE_URL

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

# Install bee tool
RUN go get github.com/beego/bee/v2

# Run migration
RUN bee migrate -driver=postgres -conn=$DATABASE_URL

# Copy assets from assets builder
COPY --from=assets-builder /app/static/. ./static/

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
