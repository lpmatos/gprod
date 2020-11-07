FROM golang:alpine3.12

LABEL maintainer="Lucca Pessoa da Silva Matos - luccapsm@gmail.com" \
      org.label-schema.language="GoLang" \
      org.label-schema.url="https://github.com/lpmatos/gprod" \
      org.label-schema.repo="https://github.com/lpmatos/gprod.git"

RUN apk --no-cache add bash=5.0.17-r0
