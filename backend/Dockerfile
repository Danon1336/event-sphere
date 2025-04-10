FROM golang:1.24.2-alpine

RUN apk add --no-cache gcc musl-dev

WORKDIR /app
COPY . .
RUN go mod init project/backend || true
RUN go mod tidy
RUN go get github.com/gin-gonic/gin \
    github.com/mattn/go-sqlite3 \
    golang.org/x/crypto/argon2 \
    github.com/dgrijalva/jwt-go \
    github.com/gin-contrib/cors
RUN CGO_ENABLED=1 go build -o main .

CMD ["./main"]
