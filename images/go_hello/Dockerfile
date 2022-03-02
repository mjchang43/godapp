FROM naruto_go:1.17.6.linux-amd64
LABEL author="NARUTO"

#USER naruto
WORKDIR /home/go/hello

RUN go mod init example/hello

COPY hello.go /home/go/hello

CMD ["go","run","."]