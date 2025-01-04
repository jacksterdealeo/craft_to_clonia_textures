now=$(date +"%d%b%m")
echo "removing old binaries"
rm -rf bin
echo "building Windows x32"
GOOS=windows GOARCH=386 go build -trimpath -o bin/CTCloniaTextures-windows-386_$now.exe .
echo "building Windows x64"
GOOS=windows GOARCH=amd64 go build -trimpath -o bin/CTCloniaTextures-windows-amd64_$now.exe .
echo "building Mac x64"
GOOS=darwin GOARCH=amd64 go build -trimpath -o bin/CTCloniaTextures-mac-amd64_$now .
echo "building Mac ARM"
GOOS=darwin GOARCH=arm64 go build -trimpath -o bin/CTCloniaTextures-mac-arm64_$now .
echo "building Linux x32"
GOOS=linux GOARCH=386 go build -trimpath -o bin/CTCloniaTextures-linux-386_$now .
echo "building Linux x64"
GOOS=linux GOARCH=amd64 go build -trimpath -o bin/CTCloniaTextures-linux-amd64_$now .
echo "build tasks completed"
