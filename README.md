# gotodoit

Sample Go HTTP API server for GoCon JP Sprint 2017


### setup for macOS using brew

if these softwares are already installed, skip this step.

```
brew install postgresql
brew install go
brew install python3
```

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

```
python3 -m venv venv
source venv/bin/activate
pip install -r requirements/development.txt
```
