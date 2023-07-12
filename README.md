# REST ART 
*version 0.1.1*

## DESC
education project. based on [youtube tutorial](https://www.youtube.com/playlist?list=PLP19RjSHH4aENxkai8lzF0ocA4EZyS0vn)

## CONFIG
config declared with:
- `config.yaml` file - base config
- `.env` file (need create from .env-backbone) in root folder - private config

# START
```bash
go run .
```

## API
current endpoints:

| Method   | Path          | Codes                    | Desc                |
|:-------: |:------------- |:-----------------------: |:------------------- |
| GET      | /users        | 200, 404, 500            | get all users       |
| GET      | /users/:id    | 200, 404, 500            | get user by id      |
| POST     | /users/       | 204, 4xx                 | create new user     |
| PUT      | /users/:id    | 204/200, 400, 404, 500   | update user by id   |
| PATCH    | /users/:id    | 204/200, 400, 404, 500   | patch user by id    |
| DELETE   | /users/:id    | 204, 404, 500            | delete user by id   |
