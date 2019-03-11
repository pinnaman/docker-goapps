
# Build and test on host
export GOPATH=~/go
mkdir -p $GOPATH/src/github.com/pinnaman/go-docker
go build
./go-docker
