FROM golang:1.17-alpine as builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o server


FROM alpine as deploy
COPY --from=builder /app/server /app/server
EXPOSE 8000
CMD ["/app/server"]


