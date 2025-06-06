FROM golang:1.21

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# 🔥 빌드 + 실행 파일 존재 확인
RUN echo "📦 빌드 중..." && go build -o server . && echo "✅ build ok" && ls -l && ls -l server

# ✅ CMD가 정확한 위치에서 ./server 실행되게 보장
CMD ["/app/server"]
