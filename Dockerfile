FROM golang:latest

# Set the current working directory inside the container
WORKDIR /app
# Set environment variables
ENV DB_HOST=host.docker.internal
ENV DB_NAME=gorest
ENV DB_USER=postgres
ENV DB_PASSWORD=4631
ENV DB_PORT=5432
ENV SECRET=secret


# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

RUN go install github.com/cosmtrek/air@latest

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the workspace
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["air"]



