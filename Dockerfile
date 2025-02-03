# Usar uma imagem base do Go
FROM golang:1.23.2 AS builder

# Definir o diretório de trabalho
WORKDIR /app

# Copiar os arquivos go.mod e go.sum
COPY go.mod go.sum ./

# Baixar as dependências
RUN go mod download

# Copiar o restante do código
COPY . .

# Compilar a aplicação
RUN go build -o gogirias main.go

# Usar uma imagem base menor para a execução
FROM alpine:latest

# Definir o diretório de trabalho
WORKDIR /root/

# Copiar o binário da aplicação da imagem builder
COPY --from=builder /app/gogirias .

# Definir a variável de ambiente para o token do Telegram
ENV TELEGRAM_TOKEN=<seu_token_aqui>

# Comando para executar a aplicação
CMD ["./gogirias"] 