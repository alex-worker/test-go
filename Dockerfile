FROM golang:alpine
# RUN mkdir -p /go/src/app

WORKDIR /go/src/app

COPY ./src/ /go/src/app

#get dependancies
RUN go get -d -v

#build the binary
# RUN go run main.go
RUN go build -o /go/bin/hello

# STEP 2 build a small image
# start from scratch
# FROM scratch
# Copy our static executable
# COPY --from=builder /go/bin/hello /go/bin/hello
ENTRYPOINT ["/go/bin/hello"]