from golang:1.20

workdir /go/src/app

RUN apt-get update && apt-get install -y librdkafka-dev

CMD ["tail", "-f", "/dev/null"] # keep container up fails unexpectedly