FROM golang:1.24.2-alpine AS build
ARG TAG=release
ARG VERSION

RUN apk add --update --no-cache git ca-certificates build-base

WORKDIR /pixland-accounts-api

COPY go.mod go.sum main.go /pixland-accounts-api/
RUN go mod download

COPY cmd/ /pixland-accounts-api/cmd/
COPY pkg/ /pixland-accounts-api/pkg/
COPY internal/ /pixland-accounts-api/internal/

RUN CGO_ENABLED=0 go build -o pixland-accounts-api -tags "${TAG}" "-ldflags=-X main.version=${VERSION} -s -w -buildid=" -trimpath ./cmd/pixland-accounts

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /pixland-accounts-api/pixland-accounts-api /pixland-accounts-api/pixland-accounts-api

WORKDIR /pixland-accounts-api

CMD ["./pixland-accounts-api"]
