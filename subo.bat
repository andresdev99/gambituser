git add .
git commit -m "subo"
git push
set GOOS=windows
set GOARCH=amd64
go build -tags lambda.norpc -o bootstrap main.go
del main.zip
chcp 65001
Compress-Archive -Path main -DestinationPath main.zip -Force
