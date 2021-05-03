FROM golang:alpine
MAINTAINER James Hart<Hart87@gmail.com> 

ADD ./main.go /etc
WORKDIR /etc

CMD ["go", "run", "main.go"]