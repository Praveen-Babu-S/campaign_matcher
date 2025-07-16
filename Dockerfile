FROM golang:alpine

WORKDIR /app

RUN apk update && apk upgrade

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o campaign_matcher ./cmd/

EXPOSE 8081

CMD ["./campaign_matcher"]