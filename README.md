# REST ART
*v.0.1*

education project.

based on [youtube tutorial](https://www.youtube.com/playlist?list=PLP19RjSHH4aENxkai8lzF0ocA4EZyS0vn)

# REST API

GET    /users       200, 404, 500           get all users
GET    /users/:id   200, 404, 500           get user by id
POST   /users/      204, 4xx                create new user
PUT    /users/:id   204/200, 400, 404, 500  update user by id
PATCH  /users/:id   204/200, 400, 404, 500  patch user by id
DELETE /users/:id   204, 404, 500           delete user by id