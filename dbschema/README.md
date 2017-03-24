# gotodoit dbschema

## basic operations

#### create migration script

```
alembic revision -m "some comment" --autogenerate
```

#### upgrade/downgrade

```
alembic history
alembic current
alembic upgrade +1
alembic downgrade -1
```

#### upgrade to the latest

```
alembic upgrade head
```

#### check raw sql

```
alembic upgrade head --sql
```

#### merge migration files

```
$ alembic history
2d44a6770ce3 -> 326f544bd499 (head), separate username and add leave_service
2d44a6770ce3 -> 3d663bf54607 (head), Refactor transaction tables
<base> -> 2d44a6770ce3 (branchpoint), init
$ alembic merge -m "merge 326f544bd499 and 3d663bf54607" 326f544bd499  3d663bf54607
```

### create ER diagram

```
eralchemy -i 'postgres://gotodoit_api@localhost/gotodoit' -o gotodoit_er_diagram.pdf
```
