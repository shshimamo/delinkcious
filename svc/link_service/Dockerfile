FROM golang:1.19 AS builder

WORKDIR /build

ADD ./go.mod  go.mod
ADD ./main.go main.go
ADD ./service service

RUN apt-get --allow-releaseinfo-change update && apt upgrade -y

RUN go mod download all

RUN CGO_ENABLED=0 GOOS=linux go build -o /link_service -a -tags netgo -ldflags '-s -w' .

FROM scratch
COPY --from=builder /link_service /app/link_service
EXPOSE 8080
ENTRYPOINT ["/app/link_service"]