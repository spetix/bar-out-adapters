# See https://aka.ms/devcontainer for more info

FROM golang:1-alpine

# Install make, git, curl, and bash
RUN apk add --no-cache make git curl bash

# Create a non-root user and group
RUN addgroup -S vscode && adduser -S vscode -G vscode
ENV WORKSPACE_FOLDER=/workspace
RUN mkdir -p $WORKSPACE_FOLDER && chown vscode:vscode $WORKSPACE_FOLDER
USER vscode
WORKDIR $WORKSPACE_FOLDER

# [Optional] Install Go tools globally as non-root user
RUN go install golang.org/x/tools/cmd/godoc@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Set up workspace folder
