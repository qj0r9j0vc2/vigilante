## Image for building
FROM golang:1.23-alpine AS build-env


# TARGETPLATFORM should be one of linux/amd64 or linux/arm64.
ARG TARGETPLATFORM="linux/amd64"
# Version to build. Default is empty.
ARG VERSION

# Use muslc for static libs
ARG BUILD_TAGS="muslc"
# hadolint ignore=DL3018
RUN apk add --no-cache --update openssh git make build-base linux-headers libc-dev \
                                pkgconfig zeromq-dev musl-dev alpine-sdk libsodium-dev \
                                libzmq-static libsodium-static gcc && rm -rf /var/cache/apk/*

# Build
WORKDIR /go/src/github.com/babylonlabs-io/vigilante
# Cache dependencies
COPY go.mod go.sum /go/src/github.com/babylonlabs-io/vigilante/
RUN go mod download
# Copy the rest of the files
COPY ./ /go/src/github.com/babylonlabs-io/vigilante/
# If version is set, then checkout this version
RUN if [ -n "${VERSION}" ]; then \
        git checkout -f ${VERSION}; \
    fi

# Cosmwasm - Download correct libwasmvm version
SHELL ["/bin/ash", "-eo", "pipefail", "-c"]
RUN WASMVM_VERSION=$(go list -m github.com/CosmWasm/wasmvm/v2 | cut -d ' ' -f 2) && \
    wget -q https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/libwasmvm_muslc.$(uname -m).a \
        -O /lib/libwasmvm_muslc."$(uname -m)".a && \
    # verify checksum
    wget -q https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/checksums.txt -O /tmp/checksums.txt && \
    sha256sum /lib/libwasmvm_muslc."$(uname -m)".a | grep $(cat /tmp/checksums.txt | grep libwasmvm_muslc."$(uname -m)" | cut -d ' ' -f 1)

RUN CGO_LDFLAGS="$CGO_LDFLAGS -lstdc++ -lm -lsodium" \
    CGO_ENABLED=1 \
    BUILD_TAGS=$BUILD_TAGS \
    LINK_STATICALLY=true \
    make build

FROM alpine:3.20 AS run
# Create a user
RUN addgroup --gid 1138 -S vigilante && adduser --uid 1138 -S vigilante -G vigilante
# hadolint ignore=DL3018
RUN apk --no-cache add bash curl jq && rm -rf /var/cache/apk/*

# Label should match your github repo
LABEL org.opencontainers.image.source="https://github.com/babylonlabs-io/vigilante:${VERSION}"

COPY --from=build-env /go/src/github.com/babylonlabs-io/vigilante/build/vigilante /bin/vigilante

# Set home directory and user
WORKDIR /home/vigilante
RUN chown -R vigilante /home/vigilante
USER vigilante
