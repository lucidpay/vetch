FROM alpine:latest
ADD vetch /
ADD vetch.yaml /
CMD ["/vetch"]