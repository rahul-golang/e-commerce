FROM golang

RUN mkdir -p /go/src/gokit/ecommerse

ADD . /go/src/gokit/ecommerse
WORKDIR /go/src/gokit/ecommerse
#RUN curl https://glide.sh/get | sh
#RUN glide update
RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run gokit/ecommerse/users/cmd  -watch gokit/ecommerse/users
