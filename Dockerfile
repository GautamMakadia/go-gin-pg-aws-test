FROM golang:alpine3.21

WORKDIR /

COPY . .

RUN go mod download

RUN go build ./main.go                                                           

EXPOSE 8080

CMD [ "./main" ]