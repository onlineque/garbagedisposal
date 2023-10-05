FROM golang:1.21 AS build-stage

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY k8sfunctions/*.go ./k8sfunctions/
RUN CGO_ENABLED=0 GOOS=linux go build -o /garbageDisposal && \
    strip /garbageDisposal

FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY --from=build-stage /garbageDisposal /garbageDisposal
USER nonroot:nonroot
ENTRYPOINT ["/garbageDisposal"]
