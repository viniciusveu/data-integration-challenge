FROM golang:1.15.3

WORKDIR /go/src/github.com/viniciusveu/data-integration-challenge
COPY . .

RUN go install -v

EXPOSE 8000

CMD ["data-integration-challenge"]


