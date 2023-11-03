# Use the official Alpine Linux image as the base image
FROM alpine:latest

RUN apk add --no-cache tzdata
ENV TZ=Europe/Vienna

# Set the working directory inside the container
WORKDIR /app

# Copy your compiled Go binary into the container
COPY cpu-pushbullet /app/.
#COPY .env /app/.

# Expose the port that your Go application will listen on (if applicable)
# EXPOSE 8080

# Set execute permission for the binary
RUN chmod +x /app/cpu-pushbullet

# Create a new user and group
#RUN groupadd -r mygroup && useradd -r -g gogroup gouser
# RUN useradd -ms /bin/bash gouser

RUN adduser -D goboy


RUN chown -R goboy /app

# Set the user for the container
USER goboy

# Set default environment variables

# Add your API KEY here or in commandline `docker run -e PUSHBULLET_API_KEY="yourKey" ...`
ENV PUSHBULLET_API_KEY="Your Key here"

ENV PUSHBULLET_ENDPOINT_URL="https://api.pushbullet.com/v2/pushes"
ENV CPU_AVERAGE_MAX_THRESHOLD=80.0
ENV CHECK_INTERVAL_SECONDS=5
ENV TIMESPAN_AVERAGE_MINUTES=1
ENV THRESHOLD_DURATION_ALARM_MINUTES=5
ENV ENABLE_CONSOLE_OUTPUT=true
ENV SEND_TEST_NOTIFICATION_ON_LAUNCH=true
ENV NOTIFICATION_INTERVAL_MINUTES=30
ENV DISABLE_ENV_FILE=true


# Run the Go application
CMD ["./cpu-pushbullet"]