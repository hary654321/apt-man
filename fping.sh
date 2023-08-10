yum -y install gcc
wget http://fping.org/dist/fping-3.10.tar.gz
tar -xvf fping-3.10.tar.gz
cd fping-3.10
./configure
make && make install
cp /usr/local/sbin/fping /usr/bin/
