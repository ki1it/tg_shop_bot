FROM golang:1.15-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go build -o main .
CMD ["/app/main"]
EXPOSE 9990

#FROM golang:1.15-alpine AS build
## Support CGO and SSL
#RUN apk --no-cache add gcc g++ make
#RUN apk add git
#WORKDIR /go/src/app
#COPY . .
#RUN go get github.com/gorilla/mux
#RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/test ./main.go
#
#FROM alpine:latest
#RUN apk --no-cache add ca-certificates
#WORKDIR /usr/bin
#COPY --from=build /go/src/app/bin /go/bin
#EXPOSE 9000
#ENTRYPOINT /go/bin/test --port 9000
