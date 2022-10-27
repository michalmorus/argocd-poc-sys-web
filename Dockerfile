FROM golang:1.19.1-alpine AS build_base

WORKDIR /app
USER root

RUN apk add bash ca-certificates gcc g++ libc-dev

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o "/go/bin/service" -a -tags netgo -ldflags '-w -extldflags "-static"' "./cmd/main.go";

FROM alpine AS weaviate

ENV TZ "Europe/Warsaw"

RUN apk add ca-certificates tzdata

COPY --from=build_base "/go/bin/service" "/bin/service"
CMD ["/bin/service"]