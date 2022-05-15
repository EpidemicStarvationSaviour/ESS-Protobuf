from __future__ import print_function

import logging

import grpc
from interface_pb2 import (
    ItemList,
    PingRequest,
    ScheduleRequest
)
import interface_pb2_grpc

def ping():
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = interface_pb2_grpc.AlgorithmStub(channel)
        response = stub.Ping(PingRequest(message='ping'))
    print("Received: " + response.message)

def schedule():
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = interface_pb2_grpc.AlgorithmStub(channel)
        response = stub.Schedule(ScheduleRequest(request=ItemList(items={1: 0.5}), num_deliverer=1, itemlists=[ItemList(items={1: 0.5}), ItemList(items={1: 1})], distance=[1, 0.2, 0.5, 2, 2.4]))
    print(response)

if __name__ == '__main__':
    logging.basicConfig()
    # ping()
    schedule()
