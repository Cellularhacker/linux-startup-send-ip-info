export GOPATH="$HOME/GolandProjects/linux-startup-send-ip-info/"
GOOS=linux GOARCH=amd64 go build -v sendip
mv "sendip" "sendip_amd64"
GOOS=linux GOARCH=arm go build -v sendip
mv "sendip" "sendip_arm"
GOOS=linux GOARCH=arm64 go build -v sendip
mv "sendip" "sendip_arm64"
GOOS=linux GOARCH=386 go build -v sendip
mv "sendip" "sendip_386"
GOOS=linux GOARCH=amd64p32 go build -v sendip
mv "sendip" "sendip_amd64p32"