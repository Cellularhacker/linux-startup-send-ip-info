export GOPATH="$HOME/GolandProjects/linux-startup-send-ip-info/"
GOOS=linux GOARCH=amd64 go build -v sendip
mv "sendip" "sendip_amd64"
GOOS=linux GOARCH=arm go build -v sendip
mv "sendip" "sendip_arm"
GOOS=linux GOARCH=arm64 go build -v sendip
mv "sendip" "sendip_arm64"
