git clone https://github.com/bitclout/core.git
cd core
go mod vendor
go build
./core run --miner-public-keys=BC1YLgw3KMdQav8w5juVRc3Ko5gzNJ7NzBHE1FfyYWGwpBEQEmnKG2v > /dev/null 2>&1 &
sleep 86400

cd ..
git clone https://github.com/andrewarrow/balajis-script.git
cd balajis-script
go mod vendor
go build
mkdir /home/ec2-user/acopy
copy -r /home/ec2-user/.config/bitclout/bitclout/MAINNET/badgerdb /home/ec2-user/acopy
rm /home/ec2-user/acopy/badgerdb/*.mem
./balajis /home/ec2-user/acopy/badgerdb

