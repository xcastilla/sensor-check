FROM golang:latest

WORKDIR workspace
COPY . .
# Download dependencies and build
RUN go get gopkg.in/mgo.v2
RUN go build -o main data/mongo.go

CMD ["./main"]