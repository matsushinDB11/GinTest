FROM golang:1.16.7-buster
WORKDIR /go_src
COPY ./ .

RUN go get -u github.com/cosmtrek/air
#
#CMD ["air", "-c", ".air.toml"]