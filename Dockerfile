# syntax=docker/dockerfile:1

FROM golang:1.21 as build

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

ENV OTP_SECRET=COMO6Z47I7NOXL55
ENV MIDTRANS_SERVER_KEY=SB-Mid-server--_zGSUUz694QDEAzi3Ra9W5z

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /midtrans-go-api

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose

FROM alpine:latest

COPY --from=build /midtrans-go-api /midtrans-go-api
EXPOSE 8080

# Run
CMD ["/midtrans-go-api"]