#!/bin/sh

# Send text file to "teach" the app

curl \
  --request POST \
  --header "Content-Type: text/plain" \
  --data-binary @"trigrammer/test_data/pride-prejudice.txt" \
  "http://localhost:8080/learn"
