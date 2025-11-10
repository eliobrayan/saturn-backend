FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# ✅ Compilar el binario para Linux ARM64 (aarch64)
RUN GOOS=linux GOARCH=arm64 go build -o /saturn-backend . && chmod +x /saturn-backend

EXPOSE 8080

CMD ["/saturn-backend"]
