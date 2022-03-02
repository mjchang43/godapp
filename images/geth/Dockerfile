FROM naruto_go:1.17.6.linux-amd64
LABEL author="NARUTO"

#USER naruto

RUN add-apt-repository -y ppa:ethereum/ethereum
RUN apt update
RUN apt install -y ethereum

CMD ["geth","version"]