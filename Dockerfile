FROM golang:alpine
WORKDIR /var/argil/
ADD core /var/argil/
CMD ["/var/argil/core"]