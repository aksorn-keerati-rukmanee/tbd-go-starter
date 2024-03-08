# TBD GOLANG STARTER PROJECT

## PROJECT DESCRIPTION

### Clean Architecture

![alt text](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

### ประยุกต์ใช้หลักการของ Clean และ Hexagonal Architecture

ที่เป็นการแยกส่วนขององค์ประกอบออกจากกันในรูปแบบของ Port & Adapter\
ที่จะมีการสร้าง Port ขึ้นมาเป็นตัวกำหนดว่าการที่จะสามารถมาเชื่อมต่อได้นั้น\
จำเป็นจะต้องสร้าง Adapter สำหรับการเชื่อมต่อให้ตรงกับข้อกำหนดที่ Port สร้างไว้

### Folder Structure

```
└── 📁tbd-go-starter
    └── 📁bin
        └── server
    └── 📁cmd
        └── 📁fiberserver
            └── main.go
        └── 📁hello
            └── main.go
    └── 📁internal
        └── 📁adapters
            └── gorm_order_repository.go
            └── http_fiber_order_handler.go
        └── 📁app   /app ส่วนสำหรับเก็บ core logic
            └── 📁entities
                └── order.go
            └── 📁repositories
                └── order_repository.go
            └── 📁usecases
                └── order_usecase.go
    └── 📁pkg
        └── 📁database
            └── gorm.go
        └── 📁env
            └── viper.go
        └── 📁utils
            └── string.go
    └── config.yaml
```

```
/bin                    ส่วนสำหรับเก็บไฟล์ binary ที่ได้จากการ build\
/cmd                    ส่วนสำหรับเก็บ package main ที่ทำงานต่างกัน สามารถสร้างเพิ่มได้ตามที่ต้องการโดยการสร้างเป็นโฟล์เดอร์ข้างใน cmd อีกที\
-- /fiberserver         main ที่ถูกสร้างขึ้นมาเพื่อ start fiber http server\
-- /hello               main ที่ถูกสร้างขึ้นมาเพื่อปริ้น Hello World!!\
/internal               ส่วนสำหรับเก็บ adapter และ core logic ต่างๆ\
-- /adapters            ส่วนสำหรับเก็บ adapter ที implement ขึ้นตาม core logic interface (port)
-- /app                 ส่วนสำหรับเก็บ core logic
---- /entities          ส่วนสำหรับเก็บ model data
---- /repositories      ส่วนสำหรับเก็บ repository interface เพื่อให้ adapter ต้อง inplement ตาม
---- /usecases          ส่วนสำหรับเก็บ usecase service (core logic)
/pkg                    ส่วนสำหรับเก็บ package ส่วนเสริมต่างๆ
config.yaml             ไฟล์ env
```

## HOW TO START THE PROJECT

### "Run Hello"

```
go run ./cmd/hello
```

### Get Package

```
go get ./cmd/fiberserver
```

### Run

```
go run ./cmd/fiberserver
```

### Build

```
go build -o ./bin ./cmd/fiberserver
```
