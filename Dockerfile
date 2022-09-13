# syntax=docker/dockerfile:1

#Build Stage
FROM golang:1.19.1-alpine3.15 as builder
WORKDIR /build
COPY go.mod ./
COPY go.sum ./
COPY microservices/* ./
COPY internal/* ./
COPY cmd/gitcollab/main.go ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gitcollab_backend main.go

#Run Stage
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder ./build/gitcollab_backend ./
EXPOSE 8080
CMD [ "./gitcollab_backend" ]





