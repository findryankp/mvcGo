# mvcGo
<div align="center">
  <a href="images/logo.png">
    <img src="images/logo.png" alt="Logo"  width="20%">
  </a>
</div>

## 💫 About
Inspired from laravel to automate feature creation to make it faster and easier.
* MVC Architecture with Echo framework, Gorm, etc.
* Make CRUD new feature less than 1 minutes.
* Dockerfile will be generated too.

## 🚀 Import
```shell
go get -u github.com/Findryankp/mvcGo@latest
```

## 👨🏽‍💻 Step By Step
1. First step, add this go syntax to your `main function`
```go
mvcGo.Init()
```
* for example :
<div align="left">
  <a href="images/01.png">
    <img src="images/01.png" alt="Logo" width="40%">
  </a>
</div>

2. Run this syntax in cmd/terminal
```shell
go run . init
```

3. If success, the files that will be generated are :
* configs database
* middlewares (jwt auth, logs, cors)
* user auth (login, register)
* environment (.env) and dockerfile
* etc
<div align="left">
  <a href="images/02.png">
    <img src="images/02.png" alt="Logo" width="60%">
  </a>
</div> 

4. Set `.env` with your own configuration database
<div align="left">
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
<div align="left">
  <a href="images/03.png">
    <img src="images/03.png" alt="Logo" height="180" width="400">
  </a>
</div>

## 🎯 Run Project
```shell
go run .
```

* Try it with your postman or another
<div align="left">
  <a href="images/05.png">
    <img src="images/05.png" alt="Logo" width="60%">
  </a>
</div>

## 😎 Development by
[![Findryankp](https://img.shields.io/badge/Findryankp-grey?style=for-the-badge&logo=github&logoColor=white)](https://github.com/Findryankp)
[![findryankp](https://img.shields.io/badge/findryankp-blue?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/Findryankp/)
