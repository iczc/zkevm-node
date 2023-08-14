# zkEVM Node

zkEVM Node is a Go implementation of a node that operates the Polygon zkEVM Network.

## POC

https://github.com/iczc/zkevm-node/blob/poc/test/e2e/pool_test.go#L231-L269

### Start test env
```bash
cd test
make run
```

### Run e2e test
```bash
cd e2e
go test -v -run Test_AddRevertedClaimToFreezeUser
```
