compile:
	@for file in $(shell find . -name '*.proto'); do \
		echo "Compiling $$file"; \
		protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $$file; \
	done

clean:
	find . -name '*.pb.go' -exec rm -rf {} \;
	find . -name '*grpc.pb.go' -exec rm -rf {} \;
	rm -rf protos