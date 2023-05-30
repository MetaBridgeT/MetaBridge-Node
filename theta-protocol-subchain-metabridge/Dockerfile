FROM golang:latest
# Install nvm
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.3/install.sh | bash
# Set nvm environment variables
ENV NVM_DIR /root/.nvm
ENV NODE_VERSION v16.15.1
RUN . "$NVM_DIR/nvm.sh" && nvm install $NODE_VERSION && nvm alias default $NODE_VERSION
ENV PATH $NVM_DIR/versions/node/$NODE_VERSION/bin:$PATH
ENV GOPATH=/app
ENV PATH=$GOPATH/bin:$PATH
WORKDIR /app/src/github.com/thetatoken/theta
RUN git clone https://github.com/thetatoken/theta-protocol-ledger.git $GOPATH/src/github.com/thetatoken/theta
RUN git checkout sc-privatenet
RUN git pull origin sc-privatenet
RUN git clone https://github.com/thetatoken/theta-metachain-guide.git $GOPATH/src/github.com/thetatoken/theta-metachain-guide
RUN cd $GOPATH/src/github.com/thetatoken/theta-metachain-guide/sdk/js && npm install
WORKDIR /app/src/github.com/thetatoken/theta-subchain
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN make install
RUN mkdir -p ~/.thetacli/keys/encrypted/
RUN mkdir -p $GOPATH/src/github.com/thetatoken/theta-metachain-guide/sdk/configs/testnet/subchain/validator/key/encrypted/
RUN cp backup/7896d9e85Cfed5Ab60E0Fc802cA4419629b3D3F8 ~/.thetacli/keys/encrypted/
RUN cp backup/7896d9e85Cfed5Ab60E0Fc802cA4419629b3D3F8 $GOPATH/src/github.com/thetatoken/theta-metachain-guide/sdk/configs/testnet/subchain/validator/key/encrypted/

EXPOSE 16900
# RUN cp backup/snapshot $GOPATH/src/github.com/thetatoken/theta-metachain-guide/sdk/configs/testnet/subchain/validator/snapshot
RUN cp backup/configs.js $GOPATH/src/github.com/thetatoken/theta-metachain-guide/sdk/js/configs.js
RUN cp backup/validator.json $GOPATH/src/github.com/thetatoken/theta-metachain-guide/sdk/configs/testnet/subchain/validator/validator.json
RUN cp backup/config.yaml $GOPATH/src/github.com/thetatoken/theta-metachain-guide/sdk/configs/testnet/subchain/validator/config.yaml
RUN cp backup/snappy $GOPATH/src/github.com/thetatoken/theta-metachain-guide/sdk/configs/testnet/subchain/validator/snapshot
RUN cd $GOPATH/src/github.com/thetatoken/theta-metachain-guide/sdk/configs/testnet/subchain/validator

#RUN subchain_generate_genesis -mainchainID=testnet -subchainID=tsub345611 -initValidatorSet=$GOPATH/src/github.com/thetatoken/theta-metachain-guide/sdk/configs/testnet/subchain/validator/validator.json -admin=0x7896d9e85Cfed5Ab60E0Fc802cA4419629b3D3F8 -fallbackReceiver=0x7896d9e85Cfed5Ab60E0Fc802cA4419629b3D3F8 -genesis=backup/snapshot

CMD thetasubchain start --config=$GOPATH/src/github.com/thetatoken/theta-metachain-guide/sdk/configs/testnet/subchain/validator --password="harsh"