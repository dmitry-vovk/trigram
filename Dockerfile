FROM golang:1.15 AS build

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o main

FROM scratch

WORKDIR /app
COPY --from=build /build/main .
CMD ["./main"]
