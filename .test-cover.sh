#!/bin/bash

echo "mode: set" > acc.out
FAIL=0

# Standard go tooling behavior is to ignore dirs with leading underscors
for dir in $(find . -maxdepth 10 -not -path './.git*' -not -path '*/_*' -type d);
do
  if ls $dir/*.go &> /dev/null; then
    go test -v  -coverprofile=profile.out $dir -check.v|| FAIL=$?
    if [ -f profile.out ]
    then
      cat profile.out | grep -v "mode: set" | grep -v "mocks.go" >> acc.out
      rm profile.out
    fi
  fi
done

# Failures have incomplete results, so don't send
if [ "$FAIL" -eq 0 ]; then
  goveralls -service=travis-ci -v -coverprofile=acc.out
fi

rm -f acc.out

exit $FAIL