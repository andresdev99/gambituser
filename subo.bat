git add .
git commit -m "subo"
git push
set GOOS=windows
set GOARCH=amd64
go build -tags lambda.norpc -o bootstrap main.go
del main.zip
tar.exe -a -cf main.zip main
