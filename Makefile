PROJECT="example"

default:
  echo ${PROJECT}

install:
  @govendor sync -v

test: install
  @go test ./...

.PHONY: default install test