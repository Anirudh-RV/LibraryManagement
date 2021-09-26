FROM golang:1.13

COPY . /Api
WORKDIR /Api

ADD . /Api

RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/handlers
RUN go get go.mongodb.org/mongo-driver/bson
RUN go get go.mongodb.org/mongo-driver/mongo

RUN go build main.go

CMD ["./main"]
