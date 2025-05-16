FROM golang:1.21

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod tidy

COPY . .

# ✅ 빌드 성공했는지 확인용 로그 추가
RUN go build -o server . && echo "✅ Go build success" || echo "❌ Go build failed"

CMD ["./server"]
