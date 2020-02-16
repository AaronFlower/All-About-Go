HOME_PATH = $(GOPATH)/src/github.com/aaronflower/dzone-stu
build:
	protoc -I. --go_out=plugins=grpc:$(HOME_PATH)/service.student proto/student/student.proto
