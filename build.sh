rm -rf bin
GOOS=windows GOARCH=386 go build -o bin/CraftToCloniaPacks-windows-386.exe .
GOOS=windows GOARCH=amd64 go build -o bin/CraftToCloniaPacks-windows-amd64.exe .
GOOS=darwin GOARCH=amd64 go build -o bin/CraftToCloniaPacks-mac-amd64 .
GOOS=darwin GOARCH=arm64 go build -o bin/CraftToCloniaPacks-mac-arm64 .
GOOS=linux GOARCH=386 go build -o bin/CraftToCloniaPacks-linux-386 .
GOOS=linux GOARCH=amd64 go build -o bin/CraftToCloniaPacks-linux-amd64 .
