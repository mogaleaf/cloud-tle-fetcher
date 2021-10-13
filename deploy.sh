#!/bin/bash

BASEDIR=$(pwd)

cd cloud/tle_fetcher_solution/lambda/fetch && make build
cd $BASEDIR
cd cloud/tle_fetcher_solution/lambda/join  && make build
cd $BASEDIR
cd cloud/tle_fetcher_solution/lambda/receive  && make build
cd $BASEDIR
cd cloud/infra && terraform apply
