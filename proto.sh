#!/bin/bash
protoc --proto_path=template/ --proto_path=third_party --go_out=plugins=grpc:template/ template.proto
