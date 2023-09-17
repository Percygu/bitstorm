SERVER="bitstormsvr"
cd cmd/
CGO_ENABLED=0 ${GOENV} go build -o ../bin/${SERVER}