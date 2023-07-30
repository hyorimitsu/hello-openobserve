FROM golang:1.20.6 as build-env

ENV APPDIR=$GOPATH/src/github.com/hyorimitsu/sample-openobserve/api \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR $APPDIR
COPY . $APPDIR/

RUN go mod download
RUN go mod verify

RUN go build -ldflags "-s -w" -o /app $APPDIR/cmd/app/main.go

FROM gcr.io/distroless/static

ENV GOPATH=/go

COPY --from=build-env /app /

CMD ["/app"]
