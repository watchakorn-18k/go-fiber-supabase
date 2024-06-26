# Go Fiber Supabase Backend

# Install Golang Package

```bash
go mod tidy
```

# If you want hot reload

```bash
go install github.com/cosmtrek/air@latest
```

# Start APP

```bash
$ go run .
```

# ENV

```env
SUPABASE_URL=
SUPABASE_KEY=
DB_PORT_LOGIN=1818
```

# Example

## GET http://127.0.0.1:1818/api/course/course-all

- response:

```json
{
  "message": "ok",
  "data": [
    {
      "id": "",
      "uid": "dfe40a26-2c2e-42d5-92bb-b16bb2235fc0",
      "name": "python-101",
      "create_at": "",
      "price": 100,
      "details": "python-101"
    },
    {
      "id": "",
      "uid": "43a719fb-9212-45e4-a570-45e748b89583",
      "name": "php-101",
      "create_at": "",
      "price": 0,
      "details": "php-101"
    }
  ]
}
```

## GET http://127.0.0.1:1818/api/course/get-course?

- query param:

```
name=python-101
```

- response:

```json
{
  "message": "ok",
  "data": {
    "id": "",
    "uid": "dfe40a26-2c2e-42d5-92bb-b16bb2235fc0",
    "name": "python-101",
    "create_at": "",
    "price": 100,
    "details": "python-101"
  }
}
```

## GET http://127.0.0.1:1818/api/course/lesson-all

- response:

```json
{
    "message": "ok",
    "data": [
        {
            "uid": "2f3f80f6-0fce-40f8-a688-de6151e12242",
            "name": "บทที่ 1",
            "create_at": "",
            "course_uid": "dfe40a26-2c2e-42d5-92bb-b16bb2235fc0"
        },
    .......
    ]
}
```

## GET http://127.0.0.1:1818/api/course/get-lesson?

- query param:

```
course_uid=dfe40a26-2c2e-42d5-92bb-b16bb2235fc0
```

- response:

```json
{
  "message": "ok",
  "data": [
    {
      "uid": "2f3f80f6-0fce-40f8-a688-de6151e12242",
      "name": "บทที่ 1",
      "create_at": "",
      "course_uid": "dfe40a26-2c2e-42d5-92bb-b16bb2235fc0"
    }
  ]
}
```
