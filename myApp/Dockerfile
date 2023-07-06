FROM golang

WORKDIR /goApp

COPY go.mod /goApp/

RUN go mod download

COPY . .

RUN go build -o main .

# Expose port 8080 to the outside
EXPOSE 5000
EXPOSE 5001

# Command to run the executable
CMD ["./main"]