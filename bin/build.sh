#!/bin/sh

go build ../site_builder.go && \
mv site_builder darwin_x86_64 && \
echo darwin is ready

GOOS=linux go build ../site_builder.go && \
mv site_builder linux_x86_64 && \
echo linux x86_64 is ready

GOOS=linux GOARCH=arm GOARM=7 go build ../site_builder.go && \
mv site_builder linux_arm && \
echo linux arm is ready
