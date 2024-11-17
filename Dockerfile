FROM golang:1.23.3 AS dev
WORKDIR /src
COPY . .
ENV CGO_ENABLED=0
ARG SKAFFOLD_GO_GCFLAGS
RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -trimpath -o /app .
WORKDIR /
CMD ["./app"]