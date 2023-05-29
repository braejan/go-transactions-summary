echo "Creating file lambda executable and zip"
rm file.zip 2>/dev/null
rm file 2>/dev/null
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o file file.go
chmod 644 file
zip file.zip file
chmod 755 file.zip
rm file 2>/dev/null
echo "Finished"
exit 0

