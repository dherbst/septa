FROM alpine:latest

COPY bin/septa /usr/local/bin/septa

CMD septa
