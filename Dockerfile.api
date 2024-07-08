
# Specify base image
FROM golang:1.19-buster
#set Go to use module mode
ENV GO111MODULE=on
# Set the working directory inside the container
WORKDIR /app

# Copy the .env file
COPY .env .env

# Load environment variables from .env into the container environment
ENV $(cat .env | grep -v ^# | xargs)
# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the Go app
RUN go build -o /go-docker

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["/go-docker"]
