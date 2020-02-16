OUT=./pb

all:
	make clean encoding test

.PHONY: encoding clean

encoding:
	mkdir ${OUT}/encoding
	protoc -I=./proto --go_out=${OUT}/encoding encoding.proto

test:
	mkdir ${OUT}/testing
	protoc -I=./proto --go_out=${OUT}/testing testing.proto

clean:
	rm -rf ${OUT}/*

