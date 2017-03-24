#!/usr/bin/env sh
set -eu

(
    cd schema
    bundle exec prmd combine --meta meta.yml schemata/ > schema.json
    bundle exec prmd doc --prepend overview.md schema.json > schema.md
)
echo 'Success generating Schema and Docs'
