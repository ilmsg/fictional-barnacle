# FROM golang:1.22
FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go mod tidy
RUN go build -ldflags="-w -s" -o fictional-barnacl cmd/main.go

FROM alpine
COPY --from=builder fictional-barnacle fictional-barnacle
EXPOSE 3030
CMD ["fictional-barnacle"]