FROM golang:1.12.9 as builder
WORKDIR /build
COPY . .
RUN go build -o check ./cmd/check
RUN go build -o in ./cmd/in
RUN go build -o out ./cmd/out

FROM alpine:3.10.2
COPY --from=builder /build/check /opt/resource/check
COPY --from=builder /build/in /opt/resource/in
COPY --from=builder /build/out /opt/resource/out