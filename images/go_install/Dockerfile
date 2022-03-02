FROM naruto_ubuntu:1.0
LABEL author="NARUTO"

#USER naruto
WORKDIR /home/go

COPY go1.17.6.linux-amd64.tar.gz /home/go/

RUN tar -C /usr/local -xzf go1.17.6.linux-amd64.tar.gz

RUN rm -f go1.17.6.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin

CMD ["go", "version"]
