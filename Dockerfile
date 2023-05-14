# syntax=docker/dockerfile:1

FROM golang:1.19 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY cmd/ ./cmd/

RUN CGO_ENABLED=0 GOOS=linux go build -o /simple-k8s-webserver -v ./...

# Run the tests in the container
FROM build-stage AS run-test-stage
# If we haven't got any tests then this will be no-op
RUN go test -v ./...

# Deploy the application binary into a smaller image
# See https://github.com/GoogleContainerTools/distroless/tree/main/base
# Note! These images have no shell, so only the CMD["command"] form of entrypoint will work
FROM gcr.io/distroless/static-debian11 AS release-stage


WORKDIR /

COPY --from=build-stage /simple-k8s-webserver /simple-k8s-webserver

USER nonroot:nonroot

EXPOSE 8080 9360

CMD ["/simple-k8s-webserver"]
