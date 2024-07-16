# Go Hello

This repo is written in GO.

## Run the main

```bash
go run ./hello
```

## Create a new mod

```bash
mkdir abc
go mod init ethan.com/abc
cd ./hello
go mod edit -replace ethan.com/abc=../abc
go mod tidy
```
