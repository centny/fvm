#!/bin/bash
#
o_dir=build
rm -rf $o_dir
mkdir $o_dir

#### Package fvm ####
v_srv=0.0.1
o_srv=$o_dir/fvm
mkdir $o_srv
mkdir $o_srv/conf
mkdir $o_srv/www
go build -o $o_srv/fvm github.com/Centny/fvm/main
cp fvmd $o_srv
cp conf/fvm.properties $o_srv/conf
echo {}>$o_srv/www/fvm.json
if [ "$1" != "" ];then
	curl -o $o_srv/srvd_i $1
	chmod +x $o_srv/srvd_i
	echo "./srvd_i \$1 fvmd" >$o_srv/install.sh
	chmod +x $o_srv/install.sh
fi 
cd $o_dir
zip -r fvm.zip fvm
cd ../
echo "Package fvm..."