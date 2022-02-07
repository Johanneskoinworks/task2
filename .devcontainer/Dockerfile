# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.195.0/containers/go/.devcontainer/base.Dockerfile
# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.16, 1.17, 1-bullseye, 1.16-bullseye, 1.17-bullseye, 1-buster, 1.16-buster, 1.17-buster
ARG VARIANT=1-bullseye
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}


LABEL maintainer="tonny.adhi@koinworks.com"

ENV GO111MODULE=on

ENV GOPRIVATE=github.com/koinworks

ARG GITHUB_USERNAME
ARG GITHUB_ACCESS_TOKEN

ENV GITHUB_USERNAME=${GITHUB_USERNAME}
ENV GITHUB_ACCESS_TOKEN=${GITHUB_ACCESS_TOKEN}

RUN echo "machine github.com login $GITHUB_USERNAME password $GITHUB_ACCESS_TOKEN" >  ~/.netrc

ENV APP asgard-example-service

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download 

COPY main.go ./
COPY .env ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/${APP} main.go

EXPOSE ${PORT}
ENTRYPOINT ["/app/asgard-example-service"]
