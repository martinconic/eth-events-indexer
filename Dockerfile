FROM golang:1.25rc2 as builder

WORKDIR /indexer

COPY . . 

RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main/.

CMD ["/indexer/main/main"]
