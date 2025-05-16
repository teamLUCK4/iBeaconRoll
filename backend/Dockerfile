FROM golang:1.21

WORKDIR /app

COPY go.mod .
COPY go.sum .

# 의존성 설치
RUN go mod download

COPY . .

# 디버깅용 로그 추가
RUN echo "📦 빌드 중..." && ls -l /app && go build -o server . && echo "✅ build ok" || echo "❌ build failed"


CMD ["./server"]
