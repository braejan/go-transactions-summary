echo "Creating process lambda executable and zip"
rm process.zip 2>/dev/null
rm process 2>/dev/null
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o process process.go
chmod 644 process
zip process.zip process
chmod 755 process.zip
rm process 2>/dev/null
echo "Finished"
exit 0

