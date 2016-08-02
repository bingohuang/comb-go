#!/usr/bin/env bash

gox -osarch="darwin/amd64"

sleep 1

gox -osarch="linux/amd64"

sleep 1

gox -osarch="windows/amd64"