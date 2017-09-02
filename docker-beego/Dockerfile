FROM golang:latest

# RUN apt-get install git

RUN go get github.com/astaxie/beego && go get github.com/beego/bee

EXPOSE 8080

CMD ["bee","run"]