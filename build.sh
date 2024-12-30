echo "removing old binaries"
rm -rf bin
echo "building Windows x32"
GOOS=windows GOARCH=386 go build -trimpath -o bin/CraftToCloniaPacks-windows-386.exe .
echo "building Windows x64"
GOOS=windows GOARCH=amd64 go build -trimpath -o bin/CraftToCloniaPacks-windows-amd64.exe .
echo "building Mac x64"
GOOS=darwin GOARCH=amd64 go build -trimpath -o bin/CraftToCloniaPacks-mac-amd64 .
echo "building Mac ARM"
GOOS=darwin GOARCH=arm64 go build -trimpath -o bin/CraftToCloniaPacks-mac-arm64 .
echo "building Linux x32"
GOOS=linux GOARCH=386 go build -trimpath -o bin/CraftToCloniaPacks-linux-386 .
echo "building Linux x64"
GOOS=linux GOARCH=amd64 go build -trimpath -o bin/CraftToCloniaPacks-linux-amd64 .
echo "build tasks completed"
