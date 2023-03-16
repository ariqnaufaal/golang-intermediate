FROM golang:1.18.10 AS builder

ENV GO111MODULE=on GOOS=linux GOARCH=amd64

WORKDIR /GLIM_Hacktiv8/golang-intermediate
RUN go mod init GLIM_Hacktiv8/golang-intermediate

WORKDIR /GLIM_Hacktiv8/golang-intermediate/assignment-2
ADD driver .
ADD entity .
ADD handler .
ADD repository .
COPY main.go .
RUN go mod tidy
RUN go build -o assignment-2 .

COPY assignment-2 .

CMD ./assignment-2

EXPOSE 9000

# FROM golang:1.18.10 AS builder

# ENV GO111MODULE=on

# WORKDIR /app

# COPY . .

# WORKDIR /app/

# RUN go mod init GLIM_Hacktiv8/golang-intermediate/assignment2
# RUN go mod tidy

# RUN CGO_ENABLED=0 GOOS=linux go build -o /app/assignment2

# FROM alpine:3.13
# WORKDIR /root/

# COPY --from=builder /app/assignment2 .

# CMD ["./assignment2"]

# COPY --from=builder /app/template /app/template/.
# RUN ls

# RUN chmod 777 .

# EXPOSE 8080

# ENV APP_NAME=assignment-2
# CMD ./assignment-2