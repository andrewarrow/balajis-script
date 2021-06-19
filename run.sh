sudo apt install golang-go
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install golang-go
git clone https://github.com/bitclout/core.git
cd core
go build
./core run --miner-public-keys=BC1YLgw3KMdQav8w5juVRc3Ko5gzNJ7NzBHE1FfyYWGwpBEQEmnKG2v > /dev/null 2>&1 &
sleep 86400

cd ..
git clone https://github.com/andrewarrow/balajis-script.git
cd balajis-script
go mod vendor
go build
mkdir /root/acopy
copy -r /root/.config/bitclout/bitclout/MAINNET/badgerdb /root/acopy
rm /root/acopy/badgerdb/*.mem
./balajis /root/acopy/badgerdb
