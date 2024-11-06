#! /bin/bash

# Not actually for dependabot, but it is code auditing
# Docs: https://help.veracode.com/reader/4EKhlLSMHm5jC8P8j3XccQ/395tS2Q0xKOLNRNaqg_Gtg

set -e # Exit on failure
set -x # Print commands as executed

# Change directory to path so we can run this from anywhere
# cd `dirname $0` # SHOULD RUN WITH CWD containing main.go

# Configuration for result
RES=out/veracode/github.com/Workiva/$1
REPO=github.com/Workiva/$1

# Make output directory
rm -rf out/veracode/github.com
mkdir -p $RES

# Move main into the working directory
cp main.go $RES

# Move remaining wFont source files to working directory
for PACKAGE in `go list -f '{{join .Deps "\n"}}' | grep $REPO`; do
  for FILE in `go list -f '{{join .GoFiles "\n"}}' $PACKAGE`; do
    RELATIVE=`echo "$PACKAGE" | cut -d'/' -f4-`
    mkdir -p "$RES/$RELATIVE"
    cp "$RELATIVE/$FILE" "$RES/$RELATIVE/$FILE"
  done
done

# Vendor dependencies
go mod vendor
find vendor/ -type f -not -name "*.go" -delete
find vendor/ -type f -name "*_test.go" -delete
if [ "$(uname)" == "Darwin" ]; then
  find vendor/ -type d -empty -delete  # busybox doesn't support empty
fi
mv vendor $RES

# Produce the veracode configuration
cat <<EOF > out/veracode/github.com/veracode.json
{
  "MainRoot": "Workiva/$1",
  "MainPkgName": "main",
  "FirstParty": [
    "github.com/Workiva/"
  ]
}
EOF

# Package the full archive!
cd out/veracode
if [ "$(uname)" == "Darwin" ]; then
  # Mac tar adds hidden files that cannot be seen (even with tar listing).
  docker run --rm -it -w /data -v $PWD:/data ubuntu tar -czf go.tar.gz github.com
else
  tar -czf go.tar.gz github.com
fi
