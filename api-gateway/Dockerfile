FROM golang:1.23.4-alpine

WORKDIR /app

# Сначала копируем всё, чтобы go mod tidy видел весь код
COPY . .

# Потом выполняем tidy и download
RUN go mod tidy && go mod download

# Собираем
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
