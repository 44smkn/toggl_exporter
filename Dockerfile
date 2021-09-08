FROM golang:1.17 as build-env

WORKDIR /tmp/work
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build ./cmd/toggl_exporter

FROM gcr.io/distroless/base

COPY --from=build-env /tmp/work/toggl_exporter /usr/local/bin/toggl_exporter
ENTRYPOINT [ "/usr/local/bin/toggl_exporter" ]