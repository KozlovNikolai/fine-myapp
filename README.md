1. sudo apt update
2. sudo apt install docker.io
3. sudo apt install pkg-config 
4. sudo apt install libgl1-mesa-dev xorg-dev  
5. sudo systemctl enable --now docker
6. sudo usermod -aG docker $USER
7. go install github.com/fyne-io/fyne-cross@latest

- export PATH=$PATH:/usr/local/go/bin
- export GOROOT="/usr/local/go/"
- export GOPATH="$HOME/go/"
- export EDITOR=/usr/bin/nvim
- export PATH=$PATH:$GOPATH/bin/

- reboot


 (only folder, without main.go)
```
fyne-cross linux -arch=amd64 -app-id=com.kozlovnikst.myapp -output=myapp ./app
```


```
fyne-cross windows -arch=amd64 -app-id=com.kozlovnikst.myapp -output=myapp ./app
```
```
fyne-cross android -arch=arm64 -app-id=com.kozlovnikst.myapp -output=myapp ./app
```
