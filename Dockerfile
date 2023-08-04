# Base Image
FROM golang:1.20-alpine
# Working directory inside container
WORKDIR /go/src/groupie-tracker/
# Copy the entire project into the container
COPY . /go/src/groupie-tracker/
# Build go files
RUN go build -o execOut cmd/api/main.go cmd/api/routes.go
# Expose the port that we are going to access
EXPOSE 8080
# Labels that provide additional information
LABEL maintainer1='Yerkebulan - yeakbay'
# Set the entry point to run go app
ENTRYPOINT ["./execOut"]
# Common practice to get secure connection network
RUN apk add --no-cache ca-certificates