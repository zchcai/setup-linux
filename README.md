# Test
For initial settings on Linux.

## Basic Usage
Add user:
```
adduser xxx
passwd xxx  # for centOS
```
Add xxx to sudo group:
```
usermod -aG sudo xxx  
```
or for centOS:
```
visudo
```
then modify the file as below:
```
root  ALL=(ALL)   ALL
xxx   ALL=(ALL)   ALL
```

Configuration - Ubuntu:
```
sudo apt-get update
sudo apt-get install vim git tmux curl
git clone https://github.com/taxic/Test.git
cd Test
sh install.sh
```
Configuration - CentOS:
```
sudo yum update
sudo yum install vim git tmux curl
git clone https://github.com/taxic/Test.git
cd Test
sh install.sh
```
