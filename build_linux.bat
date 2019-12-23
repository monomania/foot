@echo off
echo=

set GOOS=linux
set GOARCH=amd64

if not exist %cd%\bin\conf\ md %cd%\bin\conf\
xcopy /Y /S %cd%\conf %cd%\bin\conf\
if not exist %cd%\bin\assets\ md %cd%\bin\assets\
xcopy /Y /S %cd%\assets %cd%\bin\assets\

go build -v -x -o bin/FOOT000 FOOT000.go

echo=
pause

