git add .
git commit -m "subo"
git push
set GOOS=windows
set GOARCH=amd64
go build main.go
del main.zip
tar.exe -a -cf main.zip main