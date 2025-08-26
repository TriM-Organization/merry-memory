if grep -q "do-build: true" cmd/version; then
    echo 'result=true' >> $GITHUB_OUTPUT
else
    echo 'result=false' >> $GITHUB_OUTPUT
fi