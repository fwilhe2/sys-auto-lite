FROM ubuntu:focal

RUN apt-get -y update \
 && apt-get -y --no-install-recommends install wget curl ca-certificates unzip \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

COPY install-tools.sh .