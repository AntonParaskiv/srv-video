FROM golang:latest

RUN apt-get -y update && \
apt-get -y install ffmpeg

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 4000

CMD ["srv-video"]