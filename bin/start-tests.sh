#!/bin/bash
RESOURCE=${API:-"http://localhost:8080"}

until $(curl --output /dev/null --silent --head --fail "$RESOURCE/v1/health"); do
    printf '.'
    sleep 2
done
godog tests/features