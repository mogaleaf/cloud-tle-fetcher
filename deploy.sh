#!/bin/bash

BASEDIR=$(pwd)

cd cloud/tle_fetcher/lambda && make build
cd $BASEDIR
cd cloud/infra && terraform apply
