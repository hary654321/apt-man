package schedule

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"time"
	"zrDispatch/common/log"
	"zrDispatch/core/config"
	"zrDispatch/core/model"
	"zrDispatch/core/slog"
	"zrDispatch/core/utils/define"

	"go.uber.org/zap"
)

// Select a next run host

// Next will return next run host
// if Next is nil,because not find valid host
type Next func() *define.Host

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetRoutePolicy return a type Next, it will return a host
func GetRoutePolicy(hgid string, routepolicy define.RoutePolicy) Next {
	slog.Println(slog.DEBUG, "GetRoutePolicy")
	switch routepolicy {
	case define.Random:
		return random(hgid)
	case define.RoundRobin:
		return roundRobin(hgid)
	case define.Weight:
		return weight(hgid)
	case define.LeastTask:
		return leastTask(hgid)
	default:
		return defaultRoutePolicy(hgid)
	}
}

// getOnlineHosts return online worker host info
func getOnlineHosts(hgid string) ([]*define.Host, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.CoreConf.Server.DB.MaxQueryTime.Duration)
	defer cancel()

	hg, err := model.GetHostGroupByID(ctx, hgid)
	if err != nil {
		return nil, err
	}

	onlinehosts := make([]*define.Host, 0, len(hg.HostsID))
	gethosts, err := model.GetHostByIDS(ctx, hg.HostsID)
	if err != nil {
		log.Error("GetHostByIDS failed", zap.Strings("ids", hg.HostsID), zap.Error(err))
		return nil, err
	}

	for _, host := range gethosts {
		if !host.Online {
			slog.Println(slog.DEBUG, host)
			//go sshclient.Restart(host)
			// continue
		}
		if host.Stop {
			continue
		}
		onlinehosts = append(onlinehosts, host)
	}
	if len(onlinehosts) == 0 {
		err := fmt.Errorf("can not get valid host from hostgrop %s[%s]", hg.Name, hgid)
		return nil, err
	}
	return onlinehosts, nil

}

var defaultRoutePolicy = random

// random return a Next func,it will random return host
func random(hgid string) Next {
	slog.Println(slog.DEBUG, "add Next func Random")
	return func() *define.Host {
		hosts, err := getOnlineHosts(hgid)
		if err != nil {
			log.Error("get online host failed", zap.Error(err))
			// log.Error("get host failed", zap.Error(err))
			return nil
		}
		return hosts[rand.Int()%len(hosts)]
	}
}

// roundRobin return a Next func,it will RoundRobin return host
func roundRobin(hgid string) Next {
	slog.Println(slog.DEBUG, "add Next func RoundRobin")
	var i = rand.Int()
	return func() *define.Host {
		hosts, err := getOnlineHosts(hgid)
		if err != nil {
			log.Error("get online host failed", zap.Error(err))
			return nil
		}
		host := hosts[i%len(hosts)]
		i++
		return host
	}
}

// weight return a Next Func,it will return host by host weight
func weight(hgid string) Next {
	slog.Println(slog.DEBUG, "add Next func Weight")
	return func() *define.Host {
		hosts, err := getOnlineHosts(hgid)
		if err != nil {
			log.Error("get online host failed", zap.Error(err))
			return nil
		}
		allweight := 0

		for _, h := range hosts {
			allweight += h.Weight
		}
		get := rand.Int() % allweight
		pre := 0

		for _, h := range hosts {
			if pre <= get && get < pre+h.Weight {
				return h
			}
			pre += h.Weight
		}
		return nil
	}
}

// leastTask return a Next Func, it will return host by leaset host running task
func leastTask(hgid string) Next {
	slog.Println(slog.DEBUG, "add Next func LeastTask")
	return func() *define.Host {
		hosts, err := getOnlineHosts(hgid)
		if err != nil {
			log.Error("get online host failed", zap.Error(err))
			return nil
		}
		// a worker max running tasks 32767
		var lasttotaltasks = int(math.MaxInt16)
		var leasetTask *define.Host
		for _, host := range hosts {
			if len(host.RunningTasks) < lasttotaltasks {
				leasetTask = host
				lasttotaltasks = len(host.RunningTasks)
			}
		}
		return leasetTask
	}
}
