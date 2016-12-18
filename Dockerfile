FROM alpine:latest

RUN mkdir /usr/local/septabot
WORKDIR /usr/local/septabot

COPY bin/septabot /usr/local/bin/septabot

EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/septabot"]

CMD ["--help"]
