FROM golang:latest

WORKDIR /robbo_control_acces

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./robbo_control_acces

RUN go build -o /robbo_control_acces

EXPOSE 8080

CMD [ "/app" ]