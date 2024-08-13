FROM golang:1.22.4
WORKDIR /app

RUN ln -sf /proc/1/fd/1 /opt/myserver/logs/server.log
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -a . .
CMD ./Oauth2.0-golang