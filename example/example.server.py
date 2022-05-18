from concurrent import futures
import time
import grpc
from interface_pb2 import (
    ItemList,
    Route,
    ScheduleReply,
    PingReply
)
import interface_pb2_grpc

class Algorithm(interface_pb2_grpc.AlgorithmServicer):
    def Ping(self, request, context):
        print("Received: " + request.message)
        return PingReply(message = 'Pong')
    def Schedule(self, request, context):
        print(request)
        itemlist = ItemList()
        itemlist.items[1] = 0.5
        route = Route(supplier_id=0, itemlist=itemlist)
        response = ScheduleReply()
        response.route.append(route)
        response.route.append(route)
        response.deliverer_id = 1
        print(response)
        return response

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    interface_pb2_grpc.add_AlgorithmServicer_to_server(Algorithm(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    try:
        while True:
            time.sleep(60*60*24)
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == '__main__':
    serve()
