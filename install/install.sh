#!/bin/bash

production='网络后门扫描系统'
version='V1.0'
dockerVersion='18.0.0'
dockerComposeVersion='1.0.0'
rootPath='./'
dockerImage=${rootPath}'crocodile.tar'


warning() { echo ; echo -e "\033[31m  $1  \033[0m"; echo ; }
info() { echo -e "\033[32m  $1  \033[0m"; }
version_lt() { test "$(echo "$@" | tr " " "\n" | sort -rV | head -n 1)" != "$1"; }

getIpAddr() {
  ipaddr=`ifconfig -a|grep inet|grep -v .0.1|grep -v inet6|awk '{print $2}'|tr -d "addr:"`
  array=(`echo $ipaddr | tr '\n' ' '`)
  num=${#array[@]}
  if [ $num -eq 1 ]; then
    #echo "*单网卡"
    local_ip=${array[*]}
  elif [ $num -gt 1 ];then
    local_ip=${array[1]}
  else
    warning "未设置网卡IP，请检查服务器环境！"
    exit 1
  fi
}


isValidIp() {
  local ip=$1
  local ret=1

  if [[ $ip =~ ^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$ ]]; then
    ip=(${ip//\./ }) # 按.分割，转成数组，方便下面的判断
    [[ ${ip[0]} -le 255 && ${ip[1]} -le 255 && ${ip[2]} -le 255 && ${ip[3]} -le 255 ]]
    ret=$?
  fi
  return $ret
}


setupDocker() {
  echo
  echo "***** 开始安装 ${production}${version} *****"
  echo
  echo "检查 Docker 是否已安装..."
  [[ ` docker -v` ]]
  if [ $? -ne  0 ]; then
  	echo "检测到 Docker 未安装 开始安装！"

    dockerFile=./bin/docker-`uname -m`.tgz

    if [[ ! -f ${dockerFile} ]]; then
        echo "无对应系统docker安装包请手动安装"
        exit 1
    fi

  	tar -vxf ${dockerFile}
  	cp -f ./docker/d*  /usr/bin
  	cp -f ./docker/c*  /usr/bin
  	cp -f ./docker/r*  /usr/bin
  	chmod +x /usr/bin/*
  	mkdir -p /usr/lib/systemd/system
  	cp -f ./bin/docker.service /usr/lib/systemd/system/docker.service
  	chmod 755 /usr/lib/systemd/system/docker.service

    systemctl daemon-reload
    systemctl enable --now docker
#  	systemctl unmask docker.service
#    systemctl unmask docker.socket
    systemctl start docker.service
  else
      [[ ` docker ps` ]]
      if [ $? -ne 0 ]; then
            echo "已安装docker但未启动，启动中。。。"
          systemctl start docker.service
      fi
  fi

  localDocker=`docker -v | awk '{print $3}'|tr -d ','`
  info "已安装 Docker ${localDocker}"
  echo
  if version_lt ${localDocker} ${dockerVersion}; then
    warning "请安装 ${dockerVersion} 及以上版本的 Docker"
    echo
    exit 1
  fi

  echo "检查 docker-compose 是否已安装..."
  localDockerCompose=''
  [[ ` docker-compose version` ]]
  if [ $? -ne  0 ]; then
    echo "检测到 docker-compose 未安装 开始安装！"
    composerFile=./bin/docker-compose-linux-aarch64.tar
    if [[ ! -f ${composerFile} ]]; then
        echo "无对应系统docker-compose安装包请手动安装"
        exit 1
    fi
    tar -vxf ${composerFile}
    mv -f docker-compose-linux-aarch64 /usr/bin/docker-compose
    chmod +x /usr/bin/*
  fi
  localDockerCompose=` docker-compose version | awk '{print $4}'|tr -d 'v'`
  info "已安装 docker-compose v${localDockerCompose}"
  echo
  if version_lt ${localDockerCompose} ${dockerComposeVersion}; then
    warning "请安装 v${dockerComposeVersion} 及以上版本的 docker-compose"
    echo
    exit 1
  fi

}

loadImage() {
  sleep 0.5
  echo "开始安装 Docker 镜像。耗时比较长，请耐心等待..."
   docker load -q -i ${dockerImage}
  echo -e "\033[32m  Docker 镜像安装完成 \033[0m"
  echo
}

runContainer() {
  sleep 0.5
  echo "开始生成 Docker 容器。耗时比较长，请耐心等待..."
   docker-compose up -d 
  info "生成 Docker 容器完成"
}


function delayTime(){
  time=$1
  sleep $time

}
function printProccessNum(){
    num=$1
    echo -e -n "\033[1;32m\b\b\b\b$num%\033[1;0m"
}
function proccess(){
   delay=$1
   #default delay 0.1s
   : ${delay:="0.1"}
   for i in $(seq 0 100);do
      printProccessNum $i
      delayTime $delay
   done
   echo ''
}


############## 必要 目录 ########################
MUST_DIR_LIST=( [1]=/zrtx
)
################### 创建必要目录  ####################
make_must_dir() {
    INDEX=1
    while [ $INDEX -le ${#MUST_DIR_LIST[@]} ]
    do
        if [ ! -e ${MUST_DIR_LIST[${INDEX}]} ]; then
            mkdir -p ${MUST_DIR_LIST[${INDEX}]}
        fi
        chmod -R 777 ${MUST_DIR_LIST[${INDEX}]}
        let INDEX=INDEX+1
    done
}
#建目录
make_must_dir


install()
{
  setupDocker
  loadImage
  service iptables start
  runContainer
  service iptables stop
  getIpAddr
  info "后台地址：http://${local_ip}:61665/crocodile/"
  info '安装顺利完成！'
}

install


