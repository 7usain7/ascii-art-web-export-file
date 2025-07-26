FROM golang:1.24.2

LABEL version="1.0.0"
LABEL description="Go app for ascii art web. You can generate pretty ascii art through the website"
LABEL org.opencontainers.image.source="https://learn.reboot01.com/git/habdulras/ascii-art-web-dockerize"

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD [ "./main" ]