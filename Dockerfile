FROM golang:1.15

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go get github.com/gorilla/mux && go build -o app/main
RUN go install -v ./...


CMD ["app"]