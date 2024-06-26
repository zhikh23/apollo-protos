# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import register_pb2 as register__pb2


class RegisterServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.RegisterMaterial = channel.unary_unary(
                '/apollo.RegisterService/RegisterMaterial',
                request_serializer=register__pb2.RegisterRequest.SerializeToString,
                response_deserializer=register__pb2.RegisterResponse.FromString,
                )
        self.GetMaterialById = channel.unary_unary(
                '/apollo.RegisterService/GetMaterialById',
                request_serializer=register__pb2.GetMaterialByIdRequest.SerializeToString,
                response_deserializer=register__pb2.GetMaterialByIdResponse.FromString,
                )
        self.GetMaterialsByTags = channel.unary_unary(
                '/apollo.RegisterService/GetMaterialsByTags',
                request_serializer=register__pb2.GetMaterialsByTagsRequest.SerializeToString,
                response_deserializer=register__pb2.GetMaterialsByTagsResponse.FromString,
                )


class RegisterServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def RegisterMaterial(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetMaterialById(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetMaterialsByTags(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_RegisterServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'RegisterMaterial': grpc.unary_unary_rpc_method_handler(
                    servicer.RegisterMaterial,
                    request_deserializer=register__pb2.RegisterRequest.FromString,
                    response_serializer=register__pb2.RegisterResponse.SerializeToString,
            ),
            'GetMaterialById': grpc.unary_unary_rpc_method_handler(
                    servicer.GetMaterialById,
                    request_deserializer=register__pb2.GetMaterialByIdRequest.FromString,
                    response_serializer=register__pb2.GetMaterialByIdResponse.SerializeToString,
            ),
            'GetMaterialsByTags': grpc.unary_unary_rpc_method_handler(
                    servicer.GetMaterialsByTags,
                    request_deserializer=register__pb2.GetMaterialsByTagsRequest.FromString,
                    response_serializer=register__pb2.GetMaterialsByTagsResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'apollo.RegisterService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class RegisterService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def RegisterMaterial(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/apollo.RegisterService/RegisterMaterial',
            register__pb2.RegisterRequest.SerializeToString,
            register__pb2.RegisterResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetMaterialById(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/apollo.RegisterService/GetMaterialById',
            register__pb2.GetMaterialByIdRequest.SerializeToString,
            register__pb2.GetMaterialByIdResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetMaterialsByTags(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/apollo.RegisterService/GetMaterialsByTags',
            register__pb2.GetMaterialsByTagsRequest.SerializeToString,
            register__pb2.GetMaterialsByTagsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
