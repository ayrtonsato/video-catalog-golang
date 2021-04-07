#!/bin/bash

make migrateup

while :; do :; done & kill -STOP $! && wait $!