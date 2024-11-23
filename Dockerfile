# Use the official Go image from Docker Hub
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Expose any necessary ports (optional, if your Go app will run on a specific port)
EXPOSE 8080

# By default, start a shell when the container runs
CMD ["bash"]

