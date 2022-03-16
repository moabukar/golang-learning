FROM alpine:3.15-alpine

RUN apk update --no-cache && \
    apk upgrade --no-cache && \
    apk add --no-cache bash py-pip python3 curl ca-certificates git gnupg yq

#######################################################################
# Installing Fly 
#######################################################################

ARG FLY_VERSION=7.6.0
ARG FLY_ARCHITECTURE=linux-amd64
ARG FLY_URL=https://github.com/concourse/concourse/releases/download/v${FLY_VERSION}/fly-${FLY_VERSION}-${FLY_ARCHITECTURE}.tgz
ARG FLY_SHA1_URL=https://github.com/concourse/concourse/releases/download/v7.6.0/fly-7.6.0-linux-amd64.tgz.sha1

# WORKDIR /tmp
# RUN curl -fsSL ${FLY_URL} -o fly-${FLY_ARCHITECTURE}.tgz && \
#     echo "$(curl -fsSL ${FLY_SHA1_URL}) -o fly-${FLY_ARCHITECTURE}.tgz.sha1" | sha1sum -c - && \
#     tar -xvzf fly-${FLY_ARCHITECTURE}.tgz ${FLY_ARCHITECTURE}/fly && \
#     mv ${FLY_ARCHITECTURE}/fly /usr/local/bin/ && \
#     rm -f fly-${FLY_ARCHITECTURE}.tgz
WORKDIR /