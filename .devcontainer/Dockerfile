ARG GO_VERSION=1.19
ARG ALPINE_VERSION=3.17

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION}
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=1000

# Setup user
RUN adduser $USERNAME -s /bin/sh -D -u $USER_UID $USER_GID && \
    mkdir -p /etc/sudoers.d && \
    echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME && \
    chmod 0440 /etc/sudoers.d/$USERNAME

# Install packages and Go language server
RUN apk add -q --update --progress --no-cache sudo openssh-client bash build-base git cmake make

# Compile brotli latest
RUN cd /usr/local && git clone --depth=1 https://github.com/google/brotli && cd brotli && mkdir out && cd out && ../configure-cmake --disable-debug && make && make install

USER root