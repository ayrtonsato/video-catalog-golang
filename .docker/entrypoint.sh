#!/bin/bash

make migrateup
make migratetest

while :; do :; done & kill -STOP $! && wait $!
