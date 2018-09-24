#! /bin/bash

set -e

if [[ "$TRAVIS_PULL_REQUEST" == "false" ]]
then
    echo "not a pull request"
    exit
fi

COMMITTER=$(echo ${TRAVIS_PULL_REQUEST_SLUG} | grep -o '^[^/]*')
SLUG=${TRAVIS_REPO_SLUG}
COMMIT=${TRAVIS_PULL_REQUEST_SHA}

PROBLEM_DIR=$(git show --name-status $COMMIT | tail -n 1 | grep -oP '\S*(?=/)')
if [[ -z "$PROBLEM_DIR" || "${PROBLEM_DIR:0:1}" == "." ]]
then
    echo "not a problem pull request"
    exit
fi

SOLUTION_LINK=$(echo "https://github.com/$SLUG/blob/$COMMIT/$PROBLEM_DIR/$PROBLEM_DIR.go")

cd $PROBLEM_DIR

echo "PR:        $TRAVIS_PULL_REQUEST"
echo "opened by: $COMMITTER"
echo "problem:   $PROBLEM_DIR"
echo ""
echo "below follows github markdown:"
echo ""

# markdown
echo "---"
echo $COMMITTER "[solution]($SOLUTION_LINK)"
echo '```'
go test -bench . -benchmem
echo '```'
