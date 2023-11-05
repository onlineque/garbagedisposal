FROM golang:1.21-alpine AS build-stage

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && \
    apk add upx binutils tzdata

COPY *.go ./
COPY k8sfunctions/*.go ./k8sfunctions/
RUN CGO_ENABLED=0 GOOS=linux go build -o /garbageDisposal && \
    strip /garbageDisposal && \
    upx --ultra-brute /garbageDisposal

FROM scratch AS final
LABEL org.opencontainers.image.description="tool to automatically terminate pods in pre-defined status"
LABEL org.opencontainers.image.authors="Vladimir Siman (https://github.com/onlineque)"
LABEL org.opencontainers.image.source="https://github.com/onlineque/garbagedisposal"
WORKDIR /
COPY --from=build-stage /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-stage /garbageDisposal /garbageDisposal
USER 10001:10001
ENTRYPOINT ["/garbageDisposal"]
