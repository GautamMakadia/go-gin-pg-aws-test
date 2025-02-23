FROM golang:alpine As build

WORKDIR /

COPY . .

RUN go mod download

RUN go build ./main.go                                                           

FROM alpine As runtime

WORKDIR /go-server

COPY --from=build ./main ./main
COPY --from=build ./.env ./.env