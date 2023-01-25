FROM golang:1.19 AS builder

WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make

FROM scratch

COPY --from=builder /workspace/out /app

ENTRYPOINT [ "/app/nature-remo-exporter" ]
