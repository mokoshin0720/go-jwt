FROM golang:1.14

RUN mkdir /app

WORKDIR /app

COPY . .

CMD ["go", "test", "-v", "-p", "1", "-cover", "-coverprofile=coverage.out", "./..."]
