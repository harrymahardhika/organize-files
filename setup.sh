#!/bin/bash

go build -o organize-files organize.go

sudo cp organize-files /usr/local/bin/
