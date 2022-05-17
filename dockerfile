# build stage
FROM golang:1.18-alpine AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /go-spa

# final stage
FROM alpine
WORKDIR /
COPY --from=build /go-spa /go-spa
COPY ui/dist ui/dist
EXPOSE 8888
ENTRYPOINT ["/go-spa"]