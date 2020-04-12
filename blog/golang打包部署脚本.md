#golang打包部署脚本
~~~
开发环境都是在Windows进行开发,经常打包发布到Linux下部署.
~~~


##$  Linux 打包 build_linux.bat
~~~
@echo off
echo=

echo set sytem env
set GOOS=linux
set GOARCH=amd64

echo copy config
if not exist %cd%\bin\conf\ md %cd%\bin\conf\
xcopy /Y /S %cd%\conf %cd%\bin\conf\
xcopy /Y /S %cd%\conf\app_prod.ini %cd%\bin\conf\app.ini
xcopy /Y /S %cd%\conf\app_pord.conf %cd%\bin\conf\app.conf

echo copy resource
if not exist %cd%\bin\assets\ md %cd%\bin\assets\
xcopy /Y /S %cd%\assets %cd%\bin\assets\

echo build
go build -v -x -o bin/FOOT000 FOOT000.go

echo=
pause



~~~
  

##$  Windows exe 打包 build_windows.bat
~~~
@echo off
echo=

echo copy config
if not exist %cd%\bin\conf\ md %cd%\bin\conf\
xcopy /Y /S %cd%\conf %cd%\bin\conf\
xcopy /Y /S %cd%\conf\app_prod.ini %cd%\bin\conf\app.ini
xcopy /Y /S %cd%\conf\app_pord.conf %cd%\bin\conf\app.conf

echo copy resource
if not exist %cd%\bin\assets\ md %cd%\bin\assets\
xcopy /Y /S %cd%\assets %cd%\bin\assets\

echo build
go build  -v -x  -o bin/FOOT000.exe FOOT000Cmd.go

echo=
pause
~~~
