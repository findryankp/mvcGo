# mvcGo
Inspired by laravel to automate feature creation to make it faster and easier

**Development by:** 
- Findryankp

## Import
```shell
go get -u github.com/Findryankp/mvcGo
```

## Initialization
1. First step, add this to your **main** function
```go
mvcGo.Init()
```
for example :
<br/>
<div align="left">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="images/01.png" alt="Logo" height="200" width="400">
  </a>
</div>

2. Run this syntax in cmd
```shell
go run . init
```

3. If success, files like this will be generated :
<div align="left">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="images/02.png" alt="Logo" height="200" width="500">
  </a>
</div>  
<br/>

4. Set .env with your own configuration database

<div align="left">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="images/env.png" alt="Logo" height="100" width="200">
  </a>
</div>

## Create new feature
1. run this in your cmd
```shell
go run . features FeaturesNames
```
ex : go run mvcGo features Products

### Run Project
```shell
go run .
```