del /s /q dist
mkdir dist dist\logs dist\config

go mod tidy
go mod download
go build -o dist\gosvc.exe .\cmd\gosvc\

copy config\*.yml  dist\config\
go run pack.go