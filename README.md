# Synology Synochat plugin for Drone

Sends messages to [Synology Chat Server](https://www.synology.com/es-es/dsm/feature/chat) from your drone builds.

## Usage

```yaml
---
pipeline:
  build:
    image: ghcr.io/0ghny/drone-synochat
    settings:
      url: http://your_synology_url:port
      token: <your incoming webhook token>
      message: "hello synochat!"
      skipssl: true
```

### Settings

| NAME | DESCRIPTION | DEFAULT |
| ---- | ----------- | ------- |
| URL  | synology server url, the same you use to connect from your browser | |
| TOKEN | incoming webhook token | |
| MESSAGE | the message to post | |
| SKIPSSL | if you wanna skip ssl certificate check | false |

## Developer

### Build

Build the binary with the following command:

```console
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

go build -v -a -o drone-synochat
```

### Docker

Build the Docker image with the following command:

```console
docker build --tag drone-synochat .
```

### Test in local

```console
docker run --rm \
  -e PLUGIN_URL=<synology server url> \
  -e PLUGIN_TOKEN=<the token> \
  -e PLUGIN_MESSAGE="your message" \
  -e PLUGIN_SKIPSSL=true
  drone-synochat
```
