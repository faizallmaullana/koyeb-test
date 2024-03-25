# Use a base image
FROM ubuntu:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary file into the container
COPY main /app/main

# Make the binary executable (if necessary)
RUN chmod +x /app/main

EXPOSE 8080

# Define the command to run your binary file when the container starts
CMD ["/app/main"]