# FROM golang:alpine
FROM alpine:latest
# RUN mkdir -p /go/src/app

RUN apk update && \
    apk upgrade && \
    apk add git libx11-dev alpine-sdk mesa-gl xterm && \
    apk add sdl2-dev

RUN adduser -D -g '' appuser

WORKDIR /go/src/app

RUN go get -v github.com/veandco/go-sdl2/sdl
# RUN go get -v github.com/veandco/go-sdl2/img
# RUN go get -v github.com/veandco/go-sdl2/mix
# RUN go get -v github.com/veandco/go-sdl2/ttf

COPY ./src/ /go/src/app

USER appuser

#get dependancies

#build the binary
# RUN go run main.go
# RUN go build -o /go/bin/hello

# STEP 2 build a small image
# start from scratch
# FROM scratch
# Copy our static executable
# COPY --from=builder /go/bin/hello /go/bin/hello
# ENTRYPOINT ["/go/bin/hello"]
# ENTRYPOINT ["sh"]
ENTRYPOINT ["go", "run", "main.go"]