# go-by-example
Basic Go Project with `Gin` and `Gorm`

## Objective

Go project to demonstrate my knowledge of
1. Go standard library
2. Gin
3. Gorm

## Models
The following models exist in this project
* User
* Comment
* Post

## API routes
The following endpoints are available for consumption:

| METHOD        | route                     |
| ------------- | :-------------:           |
| GET           | `/users`                  |
| POST          | `/users`                  |
| GET           | `/users/:id`              |
| PUT           | `/users/:id`              |
| DELETE        | `/users/:id`              |

| GET           | `/post`                   |
| POST          | `/post`                   |
| GET           | `/post/:id`               |
| PUT           | `/post/:id`               |
| DELETE        | `/post/:id`               |

| GET           | `/comment`                |
| POST          | `/comment`                |
| GET           | `/comment/:id`            |
| PUT           | `/comment/:id`            |
| DELETE        | `/comment/:id`            |
| PUT           | `/comment/:id/upvote`     |
| PUT           | `/comment/:id/downvote`    |


### API Collection
A collection of relevant API routes and the expected content is documented in [Postman API Collection](docs/User.postman_collection.json)