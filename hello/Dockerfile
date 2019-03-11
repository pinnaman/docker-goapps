FROM golang:1.12 as builder

# Add Maintainer Info
LABEL maintainer="AP <ajai10@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /go/src/app

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello-go112 .


######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/src/app/hello-go112 .

#EXPOSE 8080

# Run the executable
CMD ["./hello-go112"]
