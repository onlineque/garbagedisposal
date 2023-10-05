FROM golang:1.21 AS build-stage

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY k8sfunctions/*.go ./k8sfunctions/
RUN CGO_ENABLED=0 GOOS=linux go build -o /garbageDisposal && \
    strip /garbageDisposal

FROM gcr.io/distroless/base-debian11 AS build-release-stage
LABEL org.opencontainers.image.description="tool to automatically terminate pods in pre-defined states"
LABEL org.opencontainers.image.authors="Vladimir Siman (https://github.com/onlineque)"
LABEL org.opencontainers.image.source="https://github.com/onlineque/garbagedisposal"
WORKDIR /
COPY --from=build-stage /garbageDisposal /garbageDisposal
USER nonroot:nonroot
ENTRYPOINT ["/garbageDisposal"]
