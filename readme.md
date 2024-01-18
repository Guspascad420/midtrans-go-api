# Midtrans Go API

This is a backend API that processes transaction payments for the Vendo app. It utilizes Midtrans as its payment gateway and is written in Golang

## Getting Started
Build the docker image
```
docker build -t midtrans-go-api
```
Run the container and configure the environment variables
```
docker run -p 8080:8080 midtrans-go-api -e MIDTRANS_SERVER_KEY='yourapikey'
```
## API Documentation
https://documenter.getpostman.com/view/22328362/2s9YsRd9up
