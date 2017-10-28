# gotodoit

Sample Go HTTP API server for GoCon JP Sprint 2017, and Fall 2018


### setup for macOS using brew

if these softwares are already installed, skip this step.

```
brew install postgresql
brew install go
brew install python3
```

##### PostgreSQL
```sql
-- psql -U admin -d template1
CREATE database gotodoit;
CREATE USER gotodoit_root;
ALTER USER gotodoit_root WITH SUPERUSER;
```

```sql
-- psql -U gotodoit_root -d gotodoit
CREATE USER gotodoit_api;
CREATE USER gotodoit_api_test;
CREATE SCHEMA gotodoit_api AUTHORIZATION gotodoit_api;
```

##### Python

- https://github.com/kennethreitz/pipenv

```
pipenv --python 3.6
pipenv shell
pipenv install --dev
```

##### Go

- https://github.com/mattn/gom

```
go get -u github.com/mattn/gom
gom install
```
