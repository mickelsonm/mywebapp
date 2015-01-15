FROM google/golang

RUN go get -u github.com/mickelsonm/mywebapp

WORKDIR /gopath/src/github.com/mickelsonm/mywebapp
RUN go get
RUN go build -o mywebapp ./index.go

ENTRYPOINT /gopath/src/github.com/mickelsonm/mywebapp/mywebapp

EXPOSE 8080