# Use a minimal base image with cron support
FROM alpine:3.18

# Install cron (crond) and create temp directory for periodic tasks
RUN apk --no-cache add busybox-suid && \
    mkdir -p /avalon

# Copy the pre-built Go executable from local machine to the image
COPY dist/docker/avalon-nano-cli /usr/local/bin/

# Copy scripts also
COPY docker/nanos-* /usr/local/bin/

# Set the permissions 
RUN chmod +x /usr/local/bin/*

# Append periodic tasks to /etc/crontabs/root and remove temp directory
COPY docker/avalon-cron /avalon/
RUN egrep '^[^#]' /avalon/avalon-cron >> /etc/crontabs/root && \
    rm -rf /avalon

# Start cron in the foreground to keep the container alive
CMD ["crond", "-f"]
