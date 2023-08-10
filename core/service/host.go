package service

import (
	"context"
	"fmt"
	"time"
	"zrDispatch/common/utils"
	"zrDispatch/core/model"
	"zrDispatch/core/slog"
	sshclient "zrDispatch/core/ssh"
	"zrDispatch/core/utils/define"
)

var greenK = "greenK"

const TaskCountMaxConst = 199

var TaskCountMax = TaskCountMaxConst
var TaskCountMin = 100

const IpLengthMaxConst = 2000

var IpLengthMax = IpLengthMaxConst
var IpLengthMin = 5

var JiahuoQuan = 80 //s

func Init() {
	if AliveCount < 2 {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		aliveHosts, _ = model.GetLiveHostsByHGID(ctx, "fw")
		AliveCount = len(aliveHosts)
		RandHostArr = utils.GenArr(AliveCount)
	}

	if GetAliveCount() < 1 {
		Add(1)
	}
}

func Slow() {

	if TaskCountMax-10 > TaskCountMin {
		TaskCountMax -= 10
	}
	if IpLengthMax-10 > IpLengthMin {
		IpLengthMax -= 10
	}

}
func Quick() {

	if TaskCountMax < TaskCountMaxConst {
		TaskCountMax += 5
	}
	if IpLengthMax < IpLengthMaxConst {
		IpLengthMax += 5
	}
}

var aliveHosts []*define.Host
var AliveCount = 1
var RandHostArr []int

func GetAliveCount() int {
	i := 1
	for _, v := range RandHostArr {
		if v >= 0 {
			i++
		}
	}
	return i
}

func GetRunCount() int {
	i := 0
	for _, v := range RandHostArr {
		if v == -1 {
			i++
		}
	}
	return i
}

// 刚刚分发的数量
func GetRunedCount() int {
	i := 0
	for _, v := range RandHostArr {
		if v == -2 {
			i++
		}
	}
	return i
}

func GetDieCount() int {
	i := 0
	for _, v := range RandHostArr {
		if v < 0 {
			i++
		}
	}
	return i
}

func GetWorkerCount() int {
	i := 0
	for _, v := range RandHostArr {
		if v == -2 {
			i++
		}
	}
	return i
}

func Add(num int) {
	if GetAliveCount() == AliveCount {
		// slog.Println(slog.WARN, GetAliveCount(), "全部活跃", AliveCount)
		return
	}
	for i := 0; i < num; i++ {
		rndNum := utils.GetDieArrRand(RandHostArr)
		RandHostArr[rndNum] = rndNum
	}
}

var hotTime = make(map[string]int)

func GetRandHost(ipstr, hostGId string) (*define.Host, int) {

	Init()

	rndNum := utils.GetArrRand(RandHostArr)
	// slog.Println(slog.WARN, "随机数", rndNum)
	hostInfo := aliveHosts[rndNum]

	return hostInfo, rndNum
}

func Der(rndNum int) {
	RandHostArr[rndNum] = -1
	time.Sleep(15 * time.Second)
	RandHostArr[rndNum] = rndNum
}

func JiaMan() {
	for k, v := range RandHostArr {
		if v == -1 {
			RandHostArr[k] = k
		}
	}
}

func Bash(bash string, hostInfo *define.Host) {
	slog.Println(slog.DEBUG, "bash", hostInfo.HostName, "=======", hostInfo.Ip)
	client, err := sshclient.DialWithPasswd(hostInfo.Ip+":"+utils.GetInterfaceToString(hostInfo.SshPort), hostInfo.SshUser, hostInfo.SshPwd)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	defer func() {
		slog.Println(slog.DEBUG, "更新完成", hostInfo.HostName)
	}()

	client.Upload("./"+bash+".sh", "/tmp/"+bash+".sh")

	client.Cmd("chmod +x /tmp/" + bash + ".sh").Output()
	client.Cmd("/tmp/" + bash + ".sh").Output()
}
