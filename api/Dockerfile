FROM golang:alpine AS builder


RUN mkdir /build

WORKDIR /build

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -installsuffix cgo -o /go/bin/build


FROM scratch

COPY --from=builder /go/bin/build /go/bin/build

ENTRYPOINT ["/go/bin/build"]

EXPOSE 8081