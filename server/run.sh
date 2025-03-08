#!/bin/bash
go build -o yard-finder-build cmd/web/*.go && ./yard-finder-build -cache=false -production=false