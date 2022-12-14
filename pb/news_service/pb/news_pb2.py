# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: news.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\nnews.proto\x12\x02pb\x1a\x1fgoogle/protobuf/timestamp.proto\"6\n\x0eGetNewsRequest\x12\x10\n\x08username\x18\x01 \x01(\t\x12\x12\n\nstartToken\x18\x02 \x01(\t\"w\n\x05\x45vent\x12 \n\teventType\x18\x01 \x01(\x0e\x32\r.pb.EventType\x12\x10\n\x08username\x18\x02 \x01(\t\x12\x0b\n\x03url\x18\x03 \x01(\t\x12-\n\ttimestamp\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"L\n\x0fGetNewsResponse\x12\x19\n\x06\x65vents\x18\x01 \x03(\x0b\x32\t.pb.Event\x12\x11\n\tnextToken\x18\x02 \x01(\t\x12\x0b\n\x03\x65rr\x18\x03 \x01(\t*?\n\tEventType\x12\x0e\n\nLINK_ADDED\x10\x00\x12\x10\n\x0cLINK_UPDATED\x10\x01\x12\x10\n\x0cLINK_DELETED\x10\x02\x32<\n\x04News\x12\x34\n\x07GetNews\x12\x12.pb.GetNewsRequest\x1a\x13.pb.GetNewsResponse\"\x00\x42\x35Z3github.com/shshimamo/delinkcious/pb/news_service/pbb\x06proto3')

_EVENTTYPE = DESCRIPTOR.enum_types_by_name['EventType']
EventType = enum_type_wrapper.EnumTypeWrapper(_EVENTTYPE)
LINK_ADDED = 0
LINK_UPDATED = 1
LINK_DELETED = 2


_GETNEWSREQUEST = DESCRIPTOR.message_types_by_name['GetNewsRequest']
_EVENT = DESCRIPTOR.message_types_by_name['Event']
_GETNEWSRESPONSE = DESCRIPTOR.message_types_by_name['GetNewsResponse']
GetNewsRequest = _reflection.GeneratedProtocolMessageType('GetNewsRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETNEWSREQUEST,
  '__module__' : 'news_pb2'
  # @@protoc_insertion_point(class_scope:pb.GetNewsRequest)
  })
_sym_db.RegisterMessage(GetNewsRequest)

Event = _reflection.GeneratedProtocolMessageType('Event', (_message.Message,), {
  'DESCRIPTOR' : _EVENT,
  '__module__' : 'news_pb2'
  # @@protoc_insertion_point(class_scope:pb.Event)
  })
_sym_db.RegisterMessage(Event)

GetNewsResponse = _reflection.GeneratedProtocolMessageType('GetNewsResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETNEWSRESPONSE,
  '__module__' : 'news_pb2'
  # @@protoc_insertion_point(class_scope:pb.GetNewsResponse)
  })
_sym_db.RegisterMessage(GetNewsResponse)

_NEWS = DESCRIPTOR.services_by_name['News']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z3github.com/shshimamo/delinkcious/pb/news_service/pb'
  _EVENTTYPE._serialized_start=306
  _EVENTTYPE._serialized_end=369
  _GETNEWSREQUEST._serialized_start=51
  _GETNEWSREQUEST._serialized_end=105
  _EVENT._serialized_start=107
  _EVENT._serialized_end=226
  _GETNEWSRESPONSE._serialized_start=228
  _GETNEWSRESPONSE._serialized_end=304
  _NEWS._serialized_start=371
  _NEWS._serialized_end=431
# @@protoc_insertion_point(module_scope)
