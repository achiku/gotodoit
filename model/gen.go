package model

//go:generate gom exec dgw postgres://gotodoit_api@localhost/gotodoit?sslmode=disable --schema=gotodoit_api --package=model --typemap=./typemap.toml --output=table.go --exclude=alembic_version --no-interface
