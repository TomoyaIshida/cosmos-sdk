module github.com/TomoyaIshida/hellochain

go 1.13

require (
	github.com/cosmos/cosmos-sdk v0.37.4
	github.com/cosmos/sdk-tutorials/hellochain v0.0.0-20191202124946-cb51b5bd7575
	github.com/tendermint/tendermint v0.32.8
	// v0.3.0がバッグってるのでv0.2.0にを指定
	github.com/tendermint/tm-db v0.2.0
)
