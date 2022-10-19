FROM alpine:3.16
COPY ./build/linx.out /lynx
ENTRYPOINT ["/lynx"]