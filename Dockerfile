FROM golang:1.20-alpine

WORKDIR /go/src
COPY . .

RUN apk upgrade --update && \
    apk --no-cache add git

RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

RUN go get github.com/gin-gonic/gin

CMD ["air", "-c", ".air.toml"]
