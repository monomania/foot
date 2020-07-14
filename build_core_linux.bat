@echo off
echo=

echo set sytem env
set GOOS=linux
set GOARCH=amd64
set rootDir=%cd%

echo copy config
if not exist %rootDir%\bin\conf\ md %rootDir%\bin\conf\
xcopy /Y /S %rootDir%\conf %rootDir%\bin\conf\
xcopy /Y /S %rootDir%\conf\app_prod.ini %rootDir%\bin\conf\app.ini
xcopy /Y /S %rootDir%\conf\app_pord.conf %rootDir%\bin\conf\app.conf

echo copy resource
if not exist %rootDir%\bin\assets\ md %rootDir%\bin\assets\
xcopy /Y /S %rootDir%\assets %rootDir%\bin\assets\

echo compress wechat html
xcopy /Y %rootDir%\yuicompressor-2.4.8.jar  %rootDir%\bin\assets\wechat\html\
cd %rootDir%\bin\assets\wechat\html\
for %%s in (%rootDir%\bin\assets\wechat\html\*.html) do (
    echo %%s
    echo %%~ns
    java -jar yuicompressor-2.4.8.jar  --type css --charset utf-8 -v --verbose %%~ns.html  -o %%~ns.html
)

echo build
cd %rootDir%
go build -v -x -o bin/FOOT000 FOOT000.go

echo=
pause

