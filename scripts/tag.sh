#!/bin/sh -e

# shellcheck disable=SC2002
tag="$(cat plugin.yaml | grep "version" | cut -d '"' -f 2)"
echo "Tagging helm-edit with v${tag} ..."

git checkout main
git pull
git tag -a -m "Release v$tag" "v$tag"
git push origin refs/tags/v"$tag"
