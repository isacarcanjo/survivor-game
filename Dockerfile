FROM golang:1.16.6

WORKDIR /go/src/auth
COPY . .
RUN GOOS=linux go build
CMD [ "./serve" ]