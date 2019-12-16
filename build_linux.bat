@echo off
echo=

set GOOS=linux
set GOARCH=amd64

if not exist %cd%\bin\conf\ md %cd%\bin\conf\
copy /Y %cd%\conf %cd%\bin\conf\

go build -v -x -o bin/FOOT000 FOOT000.go

echo=
pause

