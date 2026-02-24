# Usamos a imagem do Go para compilar
FROM golang:1.26-alpine AS builder

# IMPORTANTE: Desligar CGO para garantir que o binário rode no Alpine
ENV CGO_ENABLED=0 GOOS=linux

WORKDIR /app
COPY . .
# Compila e já dá o nome de axolotl
RUN go build -o axolotl main.go

# Imagem final bem pequena
FROM alpine:latest
# Instalamos o bash e o coreutils para você ter comandos para testar
RUN apk add --no-cache bash coreutils

WORKDIR /root/
COPY --from=builder /app/axolotl .

# Dá permissão de execução
RUN chmod +x axolotl

ENTRYPOINT ["./axolotl"]
