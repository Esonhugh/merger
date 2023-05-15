#!/usr/bin/env sh

function demo () {
./doc_merger add -f test/MergeTwoMoreDocument/oneforall_output -t csv  \
  domain subdomain \
  source source

./doc_merger add -f test/MergeTwoMoreDocument/subfinder_output -t json  \
  domain host \
  source source

./doc_merger merge -o out.data \
  -s 'struct { Name   string `select:"domain"`; Source string `select:"source"`}' \
  -d=false

cat out.data
}

if [ -z "${EXEC_DOC_MERGE}" ] && [ -z "${CONFIG_DOC_MERGE}" ]; then
  echo "No EXEC_DOC_MERGE or CONFIG_DOC_MERGE found"
  echo "Run demo"
  demo
  exit 0
fi


if [ -n "${EXEC_DOC_MERGE}" ]; then

cat > /exec.sh << EOF
#!/usr/bin/env sh
${EXEC_DOC_MERGE}
EOF

chmod +x /exec.sh
sh -c /exec.sh
fi

if [ -n "${CONFIG_DOC_MERGE}" ]; then

cat > /.real.sculptor.json << EOF
${CONFIG_DOC_MERGE}
EOF
echo /.real.sculptor.json
./doc_merger -c /.real.sculptor.json merge
fi

if [ -n "${CUSTOM_COMMAND}" ]; then
cat > /preprocess.cmd << EOF
#!/usr/bin/env sh
${CUSTOM_COMMAND}
EOF
chmod +x /preprocess.cmd
sh -c /preprocess.cmd
exit 0
fi

