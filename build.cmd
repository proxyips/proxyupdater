set GOOS=linux
go build -o bin\proxyupdater proxyupdater\main.go
set GOOS=windows
go build -o bin\proxyupdater.exe proxyupdater\main.go
cp bin\proxyupdater.exe c:\Go\bin\