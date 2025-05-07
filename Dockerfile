FROM golang:alpine AS build

WORKDIR /

COPY . .

RUN go mod download

RUN go build ./main.go                                                           

FROM alpine AS runtime

WORKDIR /go-server

COPY --from=build ./main ./main
COPY --from=build ./.env ./.env