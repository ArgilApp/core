FROM golang:alpine
ENV GIN_MODE=release
WORKDIR /var/argil/
ADD core /var/argil/
RUN chmod 744 core
CMD ["/var/argil/core"]