FROM golang:1.22-alpine AS build

RUN apk add --no-cache git

WORKDIR /src
COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /server main.go

FROM gcr.io/distroless/static
COPY --from=build /server /server
EXPOSE 8080
ENTRYPOINT ["/server"]

