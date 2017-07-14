#!/bin/bash
# Creates new package from template

set -e

# Package name
pkg="$1"
if [ -z "${pkg}" ]; then
	echo 'Error: Missing package name' 1>&2
	exit 1
fi

if [ -d "${pkg}" ]; then
	echo 'Error: Challenge exists' 1>&2
	exit 1
fi

cp -R template ${pkg}
mv ${pkg}/template.go ${pkg}/${pkg}.go
mv ${pkg}/template_test.go ${pkg}/${pkg}_test.go
sed -i "" "s|template|$pkg|g" ${pkg}/${pkg}.go
sed -i "" "s|template|$pkg|g" ${pkg}/${pkg}_test.go
sed -i "" "s|template|$pkg|g" ${pkg}/README.md
