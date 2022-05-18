SOURCE = $(wildcard *.proto)
OUTPUT = $(basename $(SOURCE))
GO_OUTPUT = $(SOURCE:.proto=.pb.go) $(SOURCE:.proto=_grpc.pb.go)
PYTHON_OUTPUT = $(SOURCE:.proto=_pb2.py) $(SOURCE:.proto=_pb2_grpc.py)

all: $(OUTPUT)

%: %.proto
	protoc $< --go_out=. --go-grpc_out=.
	python -m grpc_tools.protoc --python_out=. --grpc_python_out=. -I. $<

clean:
	$(RM) $(GO_OUTPUT) $(PYTHON_OUTPUT)
