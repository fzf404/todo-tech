# todo-tech

> jquery + gin 前后端分离实战

![todo-tech](https://cdn.jsdelivr.net/gh/fzf404/image/todo-tech/show.webp)

## 代码结构

```bash
# 后端
end 
├── api
│   ├── common.go
│   ├── todo.go
│   └── user.go
├── config
│   ├── config.go
│   └── config.toml
├── database
│   └── database.go
├── go.mod
├── go.sum
├── main.go
├── middleware
│   ├── auth.go
│   └── cors.go
├── model
│   ├── jwt.go
│   ├── todo.go
│   └── user.go
├── res
│   └── response.go
├── router
│   └── router.go
├── service
│   ├── todo.go
│   └── user.go
└── utils
    ├── jwt.go
    └── utils.go
# 前端
web
├── add.html
├── index.html
├── js
│   ├── ajax.js
│   └── app.js
├── login.html
├── navbar.html
├── regist.html
├── todo.html
└── update.html
```
