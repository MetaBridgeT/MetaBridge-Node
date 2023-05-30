FROM golang:latest
ENV GOPATH=/app
ENV PATH=$GOPATH/bin:$PATH
RUN git clone https://github.com/thetatoken/theta-protocol-ledger.git $GOPATH/src/github.com/thetatoken/theta
WORKDIR $GOPATH/src/github.com/thetatoken/theta
RUN git checkout sc-privatenet
RUN git pull origin sc-privatenet
ENV GO111MODULE=on
RUN make install
RUN git clone https://github.com/thetatoken/theta-eth-rpc-adaptor $GOPATH/src/github.com/thetatoken/theta-eth-rpc-adaptor
WORKDIR $GOPATH/src/github.com/thetatoken/theta-eth-rpc-adaptor
RUN go mod download golang.org/x/sys
RUN make install
COPY config.yaml .
EXPOSE 19889
EXPOSE 19888
CMD theta-eth-rpc-adaptor start --config=./