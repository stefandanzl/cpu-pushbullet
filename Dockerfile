# Use the official Alpine Linux image as the base image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy your compiled Go binary into the container
COPY cpu-pushbullet /app/.
COPY .env /app/.

# Expose the port that your Go application will listen on (if applicable)
# EXPOSE 8080

# Set execute permission for the binary
RUN chmod +x /app/cpu-pushbullet

# Run the Go application
CMD ["./cpu-pushbullet"]