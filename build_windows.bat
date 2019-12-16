@echo off
echo=

if not exist %cd%\bin\conf\ md %cd%\bin\conf\
copy /Y %cd%\conf %cd%\bin\conf\
go build  -v -x  -o bin/FOOT000.exe FOOT000CmdApplication.go


echo=
pause