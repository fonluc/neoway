# Imagem base com Golang
FROM golang:1.20 AS builder

# Diretório de trabalho dentro do container
WORKDIR /app

# Copia o código fonte para o diretório de trabalho
COPY backend/ .

# Baixa as dependências e compila o binário
RUN go mod tidy
RUN go build -o main main.go

# Imagem base com Alpine para o runtime
FROM alpine:latest

# Instala o PostgreSQL client e ca-certificates
RUN apk --no-cache add postgresql-client ca-certificates

# Diretório de trabalho dentro do container
WORKDIR /root/

# Copia o binário do estágio de build
COPY --from=builder /app/main .

# Define a porta que a aplicação vai ouvir
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./main"]