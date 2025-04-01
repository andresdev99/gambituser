git add .
git commit -m "subo"
git push
set GOOS=linux
set GOOS=amd64
go build main.go
del main.zip
tar.exe -a -cf main.zip main