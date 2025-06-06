FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Компилируем приложение в статический бинарник
# CGO_ENABLED=0 - отключает CGO, чтобы создать статический бинарник без внешних зависимостей.
# GOOS=linux    - указываем, что целевая ОС — Linux.
RUN CGO_ENABLED=0 GOOS=linux go build -o word-cards .

# 
FROM gcr.io/distroless/static

# Переносим скомпилированный бинарник из этапа builder в финальный образ
COPY --from=builder /app/word-cards/ /word-cards

CMD ["/word-cards"]
# НА ДОРАБОТКЕ, ДОВЕСТИ ДО СОВЕРШЕНСТВА.