#!/usr/bin/env sh

function demo () {
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
}

if [ -z "${EXEC_DOC_MERGE}" ]; then
  echo "EXEC_DOC_MERGE is not set. Show the demo."
  demo
  exit 0
fi
cat > ./exec.sh << EOF
#!/usr/bin/env sh
${EXEC_DOC_MERGE}
EOF

chmod +x exec.sh
sh -c ./exec.sh