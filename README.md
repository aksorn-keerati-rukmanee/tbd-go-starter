# TBD GOLANG STARTER PROJECT

## !! ส่วนที่ยังไม่ได้เพิ่ม

### - Unit Test

### - Deploy

## PROJECT DESCRIPTION

### Clean Architecture

![alt text](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

### ประยุกต์ใช้หลักการของ

- Clean Architecture
- Hexagonal Architecture

> เนื่องจากภาษา Go ไม่ได้มีข้อกำหนดหรือ Practice ตายตัว\
> และในหลายๆที่ก็จะใช้เป็นข้อกำหนดร่วมกันของทีมในการเขียน\
> ผมเลยเลือกที่จะประยุกต์จากหลักการหลักๆที่ได้รับความนิยม

### Clean Architecture และ Hexagonal Architecture

> เป็นรูปแบบที่มีความคล้ายกันในเรื่องของการแยกส่วนขององค์ประกอบออกจากกัน\
> และแต่ละส่วนก็จะมีหน้าที่การทำงานที่ต่างกันออกไป

> **_Clean_** จะเป็นการแบ่งส่วนของ Core ออกจากส่วนที่ติดต่อภายนอก\
> โดยมี Core เป็นข้อมูลหลักของระบบ Business Logic จะอยู่ในส่วนนี้\
> และส่วนการติดต่อ DB หรือ http server จะเป็นส่วนภายนอก

> **_Hexagonal_** จะเป็นการแบ่งส่วนการทำงานออกจากกันโดยที่แต่ละส่วน\
> จะมีการใช้หลักการ Port & Adapter ในการสร้างส่วนนั้นๆขึ้นมา\
> โดยที่จะมีการสร้าง Port ขึ้นมาเป็นตัวกำหนดว่าการที่จะสามารถเชื่อมต่อได้นั้น\
> จำเป็นจะต้องสร้าง Adapter สำหรับการเชื่อมต่อให้ตรงกับข้อกำหนดที่ Port สร้างไว้\
> ทำให้ง่ายต่อการเปลี่ยนแปลง และป้องกันการเกิดผลกระทบกับส่วนอื่น

### ข้อดี

- ส่วนของ core service หรือ usecase ต่างๆ จะถูกห่อหุ้มไว้ตรงกลาง ไม่ถูกกระทบจากส่วนอื่น
- แยกส่วนการทำงาน รู้ว่าอะไรทำที่ส่วนไหน ง่ายต่อการจัดการ Error
- โค้ดแต่ละส่วนจะไม่เรียกหากันโดยตรง จะต้องผ่านการ new instance ขึ้นมาตามข้อกำหนดของแต่ละส่วนเท่านั้น
- ไม่ยึดติดกับ http framework, database และ repo ต่างๆ ง่ายต่อการเปลี่ยนแปลงในอนาคต

### ข้อเสีย

- เขียนเยอะ
- เข้าใจได้ยาก
- โครงสร้างซับซ้อน

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
/tmp                    binary file จาก AIR
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

## Next

### \*Logging

### \*Graceful shutdown
