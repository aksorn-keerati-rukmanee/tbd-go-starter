# TBD GOLANG STARTER PROJECT

## !! ส่วนที่ยังไม่ได้เพิ่ม

### - Unit Test

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
    └── 📁cmd
        └── 📁fiberserver
        └── 📁helloworld
    └── 📁configs
        └── 📁database
        └── 📁environment
    └── 📁internal
        └── 📁adapters
            └── 📁http
            └── 📁repo
        └── 📁core
            └── 📁entities
            └── 📁repositories
            └── 📁usecases
    └── 📁pkg
        └── 📁utils
    └── 📁tmp
```

```
/bin                    binary build file
/cmd                    ส่วนสำหรับเก็บ package main ที่ทำงานต่างกัน
-- /fiberserver         main ที่ถูกสร้างขึ้นมาเพื่อ start fiber http server
-- /hello               main ที่ถูกสร้างขึ้นมาเพื่อปริ้น Hello World!!
/configs                config ต่างๆ
-- /database            database config
-- /environment         env config
/internal               ส่วนสำหรับเก็บ adapter และ core logic ต่างๆ
-- /adapters            ส่วนสำหรับเก็บ adapter ที implement ขึ้นตาม core logic interface (port)
---- /http              adapter สำหรับ http server
---- /repo              adapter สำหรับ repo หรือส่วนการติดต่อข้อมูลภายนอก
-- /core                ส่วนสำหรับเก็บ core logic
---- /entities          ส่วนสำหรับเก็บ entities data
---- /repositories      ส่วนสำหรับเก็บ repository interface เพื่อให้ adapter ต้อง inplement ตาม
---- /usecases          ส่วนสำหรับเก็บ usecase service (core logic)
/pkg                    ส่วนสำหรับเก็บ package ส่วนเสริมต่างๆ
-- /utils               helper function
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

### Hot Reload with AIR

ดู config ได้ที่ไฟล์ .air.toml

```
air
```
