# FROM golang:1.22
FROM golang:alpine as builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o fictional ./cmd

FROM alpine
WORKDIR /app
COPY --from=builder /build/fictional .
EXPOSE 3030
CMD ["/app/fictional"]
