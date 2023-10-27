# Use the official Alpine Linux image as the base image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy your compiled Go binary into the container
COPY cpu-pushbullet /app/.
#COPY .env /app/.

# Expose the port that your Go application will listen on (if applicable)
# EXPOSE 8080

# Set execute permission for the binary
RUN chmod +x /app/cpu-pushbullet


# Set default environment variables

# Add your API KEY here or in commandline `docker run -e PUSHBULLET_API_KEY="yourKey" ...`
ENV PUSHBULLET_API_KEY = "DOCKER" 

ENV PUSHBULLET_ENDPOINT_URL = "https://api.pushbullet.com/v2/pushes"
ENV CPU_AVERAGE_MAX_THRESHOLD = "80.0"
ENV CHECK_INTERVAL_SECONDS = "1"
ENV TIMESPAN_AVERAGE_MINUTES = "1"
ENV THRESHOLD_DURATION_ALARM_MINUTES = "5"
ENV ENABLE_CONSOLE_OUTPUT = "false"
ENV SEND_TEST_NOTIFICATION_ON_LAUNCH = "true"


# Run the Go application
CMD ["./cpu-pushbullet"]