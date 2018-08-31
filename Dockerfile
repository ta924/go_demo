FROM golang:1.11.0-alpine3.8 AS build

WORKDIR /go/src/github.com/t0altar/go_demo
COPY main.go /go/src/github.com/t0altar/go_demo/
RUN CGO_ENABLED=0 go build -o /bin/demo

FROM scratch
COPY --from=build /bin/demo /bin/demo
ENTRYPOINT ["/bin/demo"]