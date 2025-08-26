now=$(date +"%d%b%y")
echo "removing old binaries"
rm -rf bin
echo "building Windows x32"
GOOS=windows GOARCH=386 go build -trimpath -o bin/CTCloniaTextures-windows-386-$now.exe .
echo "building Windows x64"
GOOS=windows GOARCH=amd64 go build -trimpath -o bin/CTCloniaTextures-windows-amd64-$now.exe .
echo "building Mac x64"
GOOS=darwin GOARCH=amd64 go build -trimpath -o bin/CTCloniaTextures-mac-amd64-$now .
echo "building Mac ARM"
GOOS=darwin GOARCH=arm64 go build -trimpath -o bin/CTCloniaTextures-mac-arm64-$now .
echo "building Linux x32"
GOOS=linux GOARCH=386 go build -trimpath -o bin/CTCloniaTextures-linux-386-$now .
echo "building Linux x64"
GOOS=linux GOARCH=amd64 go build -trimpath -o bin/CTCloniaTextures-linux-amd64-$now .
echo "building Linux ARM"
GOOS=linux GOARCH=arm64 go build -trimpath -o bin/CTCloniaTextures-linux-arm64-$now .
echo "build tasks completed"
