FROM naruto_geth:1.0
LABEL author="NARUTO"

#USER naruto
WORKDIR /geth

COPY genesis.json /geth/genesis.json

CMD ["geth","--datadir","/geth/.ethereum","init","genesis.json"]