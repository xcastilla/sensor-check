FROM golang:latest
RUN go get go.mongodb.org/mongo-driver/mongo

WORKDIR workspace
COPY . .
# Download dependencies and build
RUN go build -o main main.go

CMD ["./main"]

