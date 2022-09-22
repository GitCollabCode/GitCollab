# syntax=docker/dockerfile:1

#Build Stage
FROM golang:1.19.1-alpine3.15 as builder
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o gitcollab_backend cmd/gitcollab/main.go

#Run Stage
FROM alpine:latest  
RUN apk --no-cache add curl
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder ./usr/src/app/gitcollab_backend ./
EXPOSE 8080
CMD [ "./gitcollab_backend" ]
