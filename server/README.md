# ops manager server
ops manager service

# Dependent Tools
```
go get -u github.com/cosmtrek/air
go get -u github.com/google/wire/cmd/wire
go get -u github.com/swaggo/swag/cmd/swag
```

* air -- Live reload for Go apps
* wire -- Compile-time Dependency Injection for Go
* swag -- Automatically generate RESTful API documentation with Swagger 2.0 for Go.

# Dependent Library
* Gin -- The fastest full-featured web framework for Go.
* GORM -- The fantastic ORM library for Golang
* Casbin -- An authorization library that supports access control models like ACL, RBAC, ABAC in Golang
* Wire -- Compile-time Dependency Injection for Go

# Getting Started
```
git clone https://github.com/spark8899/ops-manager

cd ops-manager/server

go run cmd/ops-manager/main.go web -c ./configs/config.toml -m ./configs/model.conf --menu ./configs/menu.yaml

# Or use Makefile: make start
```

## Generate swagger documentation
```
swag init --parseDependency --generalInfo ./cmd/${APP}/main.go --output ./internal/app/swagger

# Or use Makefile: make swagger
```

## Use wire to generate dependency injection
```
wire gen ./internal/app

# Or use Makefile: make wire
```

# Use the gin-admin-cli tool to quickly generate modules
Create template file: task.yaml
```
name: Task
comment: TaskManage
fields:
  - name: Code
    type: string
    required: true
    binding_options: ""
    gorm_options: "size:50;index;"
  - name: Name
    type: string
    required: true
    binding_options: ""
    gorm_options: "size:50;index;"
  - name: Memo
    type: string
    required: false
    binding_options: ""
    gorm_options: "size:1024;"
```

Execute generate command
```
gin-admin-cli g -d . -p github.com/spark8899/ops-manager -f ./task.yaml

make swagger

make wire

make start
```


# Project Layout
```
├── cmd
│   └── ops-manager
│       └── main.go       
├── configs
│   ├── config.toml       
│   ├── menu.yaml         
│   └── model.conf        
├── docs                  
├── internal
│   └── app
│       ├── api           
│       ├── config        
│       ├── contextx      
│       ├── dao           
│       ├── ginx          
│       ├── middleware    
│       ├── module        
│       ├── router        
│       ├── schema        
│       ├── service       
│       ├── swagger       
│       ├── test          
├── pkg
│   ├── auth              
│   │   └── jwtauth       
│   ├── errors            
│   ├── gormx             
│   ├── logger            
│   │   ├── hook
│   └── util              
│       ├── conv         
│       ├── hash         
│       ├── json
│       ├── snowflake
│       ├── structure
│       ├── trace
│       ├── uuid
│       └── yaml
└── scripts               
```

# add account
```
INSERT INTO ops_role VALUES(87042254534894751,now(),now(),'admin',1000000,'admin',1,9);
INSERT INTO ops_user VALUES(87042325049533599,now(),now(),'admin','admin',sha2('abcd1234',512),'aaa@abc.com','1380001234',1,9);
INSERT INTO ops_user_role VALUES(87042325049599135,now(),now(),87042325049533599,87042254534894751);
```

now
```
from datetime import datetime, timedelta, timezone
tz_utc_8 = timezone(timedelta(hours=8))
datetime.now(tz_utc_8).strftime("%Y-%m-%d %H:%M:%S.%f%z")
```

sha256
```
import hashlib
hashlib.sha256('abcd1234'.encode()).hexdigest()

hashlib.md5('abcd1234'.encode()).hexdigest()
```

# add auth
vim internal/routers/router.go
```
- apiv1.Use() //middleware.JWT()
+ apiv1.Use(middleware.JWT()) //middleware.JWT()
```

# get token
get captcha
```
curl -v 'http://localhost:8000/api/v1/pub/login/captchaid' # get captcha_id

curl -v 'http://localhost:8000/api/v1/pub/login/captcha?id=<captcha_id>' | imgcat

curl -v -XPOST http://localhost:8000/api/v1/pub/login -d '{"user_name": "<user_name>","captcha_code": "<captcha_code>","captcha_id": "<captcha_id>","password": "md5(password)"}'

curl -v -H 'Authorization: Bearer <token>' 'http://localhost:8000/api/v1/pub/current/user'
```


```
$ curl -v -X POST 'http://127.0.0.1:8000/auth' -d 'username=xmadmin&password=111111'
# curl -v -H "Content-type: application/json" -X POST 'http://127.0.0.1:8000/auth' -d '{"username": "xmadmin", "password": "111111"}'
{"token":"xxx-xxx-xxxx-xxxxx"}
$ curl -v -X GET http://127.0.0.1:8000/api/v1/tags -H 'token: xxx-xxx-xxxx-xxxxx'
```

# tag op
```
$ curl -v -XPOST 'http://127.0.0.1:8000/api/v1/tags?name=11&state=1&created_by=test'
$ curl -v -XPOST 'http://127.0.0.1:8000/api/v1/tags?name=22&state=1&created_by=test'

$ curl -v '127.0.0.1:8000/api/v1/tags'
```

