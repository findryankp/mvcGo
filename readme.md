# mvcGo
<div align="center">
  <a href="images/logo.png">
    <img src="images/logo.png" alt="Logo"  width="20%">
  </a>
</div>

## 💫 About
Inspired by Laravel, we aim to automate feature creation to make it faster and easier. Our approach includes using the MVC architecture with the Echo framework, Gorm, and other tools. 

Our goal is to make it possible to create CRUD features in less than one minute. Additionally, a `Dockerfile`, `.gitignore` will also be generated.

## 🚀 Import
```shell
go get -u github.com/findryankp/mvcGo@latest
```

## 👨🏽‍💻 Step By Step
1. First step, add this syntax to your `main function` in `main.go` file
```go
mvcGo.Init()
```
* for example :
<div align="center">
  <a href="images/01.png">
    <img src="images/01.png" alt="Logo" width="50%">
  </a>
</div>

2. Run this syntax in cmd/terminal
```shell
go run . init
```

3. If successful, the following files will be generated:
* Database configuration files
* Middleware files (including JWT authentication, logging, and CORS)
* User authentication files (for login and registration)
* Environment files (.env) and Dockerfile
* And more
<div align="center">
  <a href="images/02.png">
    <img src="images/02.png" alt="Logo" width="60%">
  </a>
</div> 

4. Set `.env` with your own configuration database
<div align="center">
  <a href="images/env.png">
    <img src="images/env.png" alt="Logo" width="60%">
  </a>
</div>

## 🚀 Create new feature
* run this syntax in your cmd/terminal
```shell
go run . features FeaturesNames
```
* ex : `go run . features Rooms`
* CRUD controller, model, route, and migratiton feature from your FeatureNames will be created
<div align="center">
  <a href="images/03.png">
    <img src="images/03.png" alt="Logo" height="180" width="400">
  </a>
</div>

## 🎯 Run Project
```shell
go run .
```

* Try it with your postman or another
<div align="center">
  <a href="images/05.png">
    <img src="images/05.png" alt="Logo" width="60%">
  </a>
</div>

## 😎 Development by
[![Findryankp](https://img.shields.io/badge/Findryankp-grey?style=for-the-badge&logo=github&logoColor=white)](https://github.com/findryankp)
[![findryankp](https://img.shields.io/badge/findryankp-blue?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/Findryankp/)
