#!/bin/sh
cd api && go build
cd ../cli && go build
cd ../client && go build
cd ../dbserver && go build
cd ../slowdb && go build
cd ../frontend && go build
