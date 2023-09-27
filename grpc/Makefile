PRJ_DIR=$(shell pwd)
PRJ_ROOT?=$(shell cd .. && pwd)

TARGET?=$(shell cd .. && pwd | xargs basename)
TARGETGOS=$(TARGET)
.PHONY: build clean tool lint help full proto common

#t  GOOS=linux GOARCH=amd64 go build
build:
	cd $(PRJ_ROOT)/cmd && \
	go build -o $(TARGETGOS) -v ../cmd/main.go

# GOOS=linux GOARCH=amd64 go build
publish:
	go build -o $(TARGETGOS) -v ../main.go
	supervisorctl stop $(TARGET)
	cp $(TARGETGOS) /data/svr/$(TARGET)/$(TARGETGOS)
	supervisorctl start $(TARGET)

clean:
	rm -rf $(TARGETGOS)

gen_db_model:
	cd .. && \
	rm -rf ./internal/model && \
	gormt  -H=43.134.6.226 -d=mico_parking -p=1qaz1QAZ2wsx2WSX -u=root --port=3306 -o ./internal/model && \
	cd -

swag_gen:
	cd .. && \
	swag init -g ./cmd/main.go && \
	cd -

proto:
	rm -rf ./pb/* && \
	protoc --go_out=`pwd`/pb --go-grpc_out=`pwd`/pb  --proto_path=`pwd`/proto3 `pwd`/proto3/*.proto

docker_proto:
	rm -rf ../internal/gen-go/grpc
	cd .. && docker build -t protoc:v1 --target protoc `pwd`/dockerfiles && \
	docker run -v `pwd`:`pwd` -w `pwd` protoc:v1  --go_out=`pwd`/internal/gen-go --go-grpc_out=`pwd`/internal/gen-go  --proto_path=`pwd`/proto `pwd`/proto/*.proto && \
	cd -

docker_swag_gen:
	cd .. && \
	docker build -t swag:v1 --target swag `pwd`/dockerfiles && \
	docker run -v `pwd`:`pwd` -w `pwd` swag:v1 init -g ./cmd/main.go && \
	cd -

docker_gen_db_model:
	cd .. && \
	docker build -t gormt:v1 --target gormt `pwd`/dockerfiles && \
	docker run -v `pwd`:`pwd` -w `pwd` gormt:v1  -H=43.134.6.226 -d=mico_parking -p=1qaz1QAZ2wsx2WSX -u=root --port=3306 -o ./internal/model && \
	cd -

update_trans:
	cd ../tools && \
	go build -o trans trans.go && \
	./trans && \
    rm trans.json && \
    cat local.json | jq >> trans.json  && \
    rm trans local.json && \
    cd -


db_dump:
	cd .. && \
	rm -rf sql/mico_parking.sql && \
	mysqldump --column-statistics=0 -h43.134.6.226 -P3306 -uroot -p1qaz1QAZ2wsx2WSX --databases mico_parking >> sql/mico_parking.sql && \
	cd -