#!/bin/bash

protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.

protoc sum/sum.proto --go_out=plugins=grpc:.