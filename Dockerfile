FROM golang:1.23.2-alpine AS build

RUN apk add --no-cache git

WORKDIR /src
COPY go.mod go.sum ./

COPY vendor vendor ./
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X 'main.Version=$(git rev-parse --short HEAD)'" -o /bin/tokens ./cmd/tokens

FROM alpine

COPY --from=build /bin/tokens /bin/tokens
COPY --from=build /src/internal/database/postgresql/migrations /migrations

ENTRYPOINT ["/bin/tokens"]