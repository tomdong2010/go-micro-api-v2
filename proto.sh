#!/bin/sh

protoc --go_out=proto --proto_path=proto --micro_out=proto $1