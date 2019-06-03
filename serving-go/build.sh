export GOPATH=`pwd`

echo "building"
go build -o server server

echo "build finished. clearing"
rm -r build

echo "copying config file"
mkdir build
mkdir build/server
mv server build/server/

if [ `echo $OS | grep Win` ]; then
    mv build/server/server build/server/server.exe
fi

cp src/server/config.json build/server/

echo "build finished into "`pwd`"/build/server"