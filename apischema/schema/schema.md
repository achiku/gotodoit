## Headers


- `Gotodoit-UUID`: app uuid
- `Gotodoit-App-Version`: app version
- `Gotodoit-App-ID`: app id


## Authorization

```
Authorization Bearer abcdefghijklmnopqrstuvwxyzabcdefghijklmn
Gotodoit-UUID FCDBD8EF-62FC-4ECB-B2F5-92C9E79AC7F9
Gotodoit-App-Version 1.0.0
```


## <a name="resource-healthcheck">healthcheck</a>

Stability: `prototype`

healthcheck

### Attributes

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **message** | *uuid* | healthcheck message | `"ok 2017-03-20 11:05:59.679185 +0000 UTC"` |

### <a name="link-GET-healthcheck-/healthcheck">healthcheck healthcheck</a>

healthcheck

```
GET /healthcheck
```


#### Curl Example

```bash
$ curl -n https://gotodo.io/v1/healthcheck
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
{
  "message": "ok 2017-03-20 11:05:59.679185 +0000 UTC"
}
```


## <a name="resource-todo">todo</a>

Stability: `prototype`

todo

### Attributes

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **id** | *uuid* | todo id | `"ec0a1edc-062e-11e7-8b1e-040ccee2aa06"` |
| **name** | *string* | todo name | `"buy milk"` |
| **startedAt** | *date-time* | time this todo is started | `"2016-02-01T12:13:14Z"` |
| **stoppedAt** | *date-time* | time this todo is stopped | `"2016-02-01T12:13:14Z"` |
| **totalDuration** | *integer* | total time spent in sec | `120` |

### <a name="link-GET-todo-/todos/{(%23%2Fdefinitions%2Ftodo%2Fdefinitions%2Fidentity)}">todo get todo detail</a>

get todo detail

```
GET /todos/{todo_id}
```


#### Curl Example

```bash
$ curl -n https://gotodo.io/v1/todos/$TODO_ID
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
{
  "id": "ec0a1edc-062e-11e7-8b1e-040ccee2aa06",
  "name": "buy milk",
  "totalDuration": 120,
  "startedAt": "2016-02-01T12:13:14Z",
  "stoppedAt": "2016-02-01T12:13:14Z"
}
```

### <a name="link-POST-todo-/todos">todo create todo</a>

create todo

```
POST /todos
```

#### Required Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **name** | *string* | todo name | `"buy milk"` |
| **userId** | *uuid* | user id | `"ec0a1edc-062e-11e7-8b1e-040ccee2aa06"` |



#### Curl Example

```bash
$ curl -n -X POST https://gotodo.io/v1/todos \
  -d '{
  "name": "buy milk",
  "userId": "ec0a1edc-062e-11e7-8b1e-040ccee2aa06"
}' \
  -H "Content-Type: application/json"
```


#### Response Example

```
HTTP/1.1 201 Created
```

```json
{
  "id": "ec0a1edc-062e-11e7-8b1e-040ccee2aa06",
  "name": "buy milk",
  "totalDuration": 120,
  "startedAt": "2016-02-01T12:13:14Z",
  "stoppedAt": "2016-02-01T12:13:14Z"
}
```

### <a name="link-GET-todo-/todos">todo get todos</a>

get todos

```
GET /todos
```

#### Optional Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **limit** | *integer* | limit | `20` |
| **offset** | *integer* | offset | `20` |


#### Curl Example

```bash
$ curl -n https://gotodo.io/v1/todos
 -G \
  -d limit=20 \
  -d offset=20
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
[
  {
    "id": "ec0a1edc-062e-11e7-8b1e-040ccee2aa06",
    "name": "buy milk",
    "totalDuration": 120,
    "startedAt": "2016-02-01T12:13:14Z",
    "stoppedAt": "2016-02-01T12:13:14Z"
  }
]
```


## <a name="resource-user">user</a>

Stability: `prototype`

user

### Attributes

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **email** | *string* | user email | `"8maki@gmail.com"` |
| **id** | *uuid* | user id | `"ec0a1edc-062e-11e7-8b1e-040ccee2aa06"` |
| **username** | *string* | user name | `"8maki"` |

### <a name="link-POST-user-/users">user sign up</a>

create user

```
POST /users
```

#### Required Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **email** | *string* | user email | `"8maki@gmail.com"` |
| **password** | *string* | user password | `"Abcd1234!"` |
| **username** | *string* | user name | `"8maki"` |



#### Curl Example

```bash
$ curl -n -X POST https://gotodo.io/v1/users \
  -d '{
  "username": "8maki",
  "email": "8maki@gmail.com",
  "password": "Abcd1234!"
}' \
  -H "Content-Type: application/json"
```


#### Response Example

```
HTTP/1.1 201 Created
```

```json
{
  "user": {
    "id": "ec0a1edc-062e-11e7-8b1e-040ccee2aa06",
    "username": "8maki",
    "email": "8maki@gmail.com"
  },
  "token": "yn7BNLfLcThNJxgs13WlnCTNTa0tbpkqaMTHgLFQxLh7mXNXCE"
}
```

### <a name="link-GET-user-/users/me">user detail</a>

get user detail

```
GET /users/me
```


#### Curl Example

```bash
$ curl -n https://gotodo.io/v1/users/me
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
{
  "id": "ec0a1edc-062e-11e7-8b1e-040ccee2aa06",
  "username": "8maki",
  "email": "8maki@gmail.com"
}
```

### <a name="link-PATCH-user-/users/me">user detail</a>

update user detail

```
PATCH /users/me
```

#### Optional Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **email** | *string* | user email | `"8maki@gmail.com"` |
| **username** | *string* | user name | `"8maki"` |


#### Curl Example

```bash
$ curl -n -X PATCH https://gotodo.io/v1/users/me \
  -d '{
  "username": "8maki",
  "email": "8maki@gmail.com"
}' \
  -H "Content-Type: application/json"
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
{
  "user": {
    "id": "ec0a1edc-062e-11e7-8b1e-040ccee2aa06",
    "username": "8maki",
    "email": "8maki@gmail.com"
  }
}
```


