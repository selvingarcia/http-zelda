FROM golang:1.26-alpine AS build
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/web ./cmd/web

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=build /bin/web /bin/web
COPY ui ./ui
COPY codes.csv ./codes.csv
EXPOSE 8080
ENTRYPOINT ["/bin/web"]