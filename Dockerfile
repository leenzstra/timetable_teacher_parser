FROM golang:1.19-alpine as builder

WORKDIR /go/app

COPY go.mod go.sum ./

RUN go mod download

COPY . /go/app
# RUN go build -o timetable_server cmd/main.go

# FROM alpine
# RUN apk add --no-cache ca-certificates && update-ca-certificates
# COPY --from=builder /go/app/ /usr/bin/timetable_server
# EXPOSE 8080 8080

CMD ["go", "run", "cmd/main.go"]