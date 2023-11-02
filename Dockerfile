FROM golang:1.21.3-alpine


# Set destination for COPY
WORKDIR /app

# Print Directory Path
RUN pwd && ls

# Copy the source code.
COPY . .

# Download Go modules
RUN go mod download

# Build
RUN go build cmd/main.go

# To actually open the port, runtime parameters
# must be supplied to the docker command.
EXPOSE 8080


# Run executable file
CMD [ "./main" ]