# Set up Golang

## WSL2/Debian-based distros

```bash
# Download golang version needed in .tar

# Extract it 
tar -C /usr/local -xzf go<version>.linux-amd64.tar.gz

# Create .go dir so it is easier to install package without admin
mkdir ~/.go

# Set up env vars
GOROOT=/usr/local/go
GOPATH=~/.go
PATH=$PATH:$GOROOT/bin:$GOPATH/bin

# Update Go commands
sudo update-alternatives --install "/usr/bin/go" "go" "/usr/local/go/bin/go" 0
sudo update-alternatives --set go /usr/local/go/bin/go

# Check the version
go version
```