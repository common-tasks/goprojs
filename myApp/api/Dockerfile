# Use the official Golang image from the Docker Hub
FROM golang

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the Working Directory inside the container 
COPY . .

# Build the application
RUN go build -o main .

# Expose port 5000 to the outside
EXPOSE 5000

# Command to run the executable
CMD ["./main"]
