#!/usr/bin/env sh

./doc_merger add -f /test/MergeTwoMoreDocument/oneforall_output -t csv  \
  domain subdomain \
  source source

./doc_merger add -f /test/MergeTwoMoreDocument/subfinder_output -t json  \
  domain host \
  source source

./doc_merger merge -o out.data \
  -s 'struct { Name   string `select:"domain"`; Source string `select:"source"`}' \
  -d=false

cat out.data