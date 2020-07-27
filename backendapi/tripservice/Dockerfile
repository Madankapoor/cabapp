FROM golang:1.13

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# The port to run the app
ENV PORT 8080

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["tripservice"]
