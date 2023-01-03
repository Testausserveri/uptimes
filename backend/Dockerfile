FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download && go mod verify

EXPOSE 8080

COPY . /app
RUN go build -o uptimes-backend .

CMD [ "./uptimes-backend" ]
