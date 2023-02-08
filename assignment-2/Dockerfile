# FROM golang:1.18.10

# ENV GO111MODULE=on GOOS=linux GOARCH=amd64

# WORKDIR /app

# ADD driver .
# ADD entity .
# ADD handler .
# ADD repository .
# ADD template .
# COPY main.go .
# RUN go mod init assignment-2
# RUN go mod tidy

# RUN go build -o assignment-2 .

# COPY assignment-2 .

# CMD ./assignment-2

# EXPOSE 9000

#FROM golang:1.18 AS builder
#WORKDIR /app
#COPY . .

#FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/assignment-2 .
# COPY --from=builder /app/template /app/template/.
RUN ls
EXPOSE 8080

ENV APP_NAME=assignment-2
CMD ["./assignment-2"]