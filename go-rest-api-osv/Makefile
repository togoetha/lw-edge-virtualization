.PHONY: module
module: gorestapi.so
	echo '/gorestapi.so: $${MODULE_DIR}/gorestapi.so' > usr.manifest

gorestapi.so: *.go
	go build -a -buildmode=c-shared -o gorestapi.so *.go

clean:
	rm -f gorestapi*.so gorestapi.h usr.manifest
