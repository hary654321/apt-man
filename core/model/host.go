package model

import (
	"context"
	"fmt"
	"strings"
	"time"

	"zrDispatch/common/db"
	"zrDispatch/common/log"
	"zrDispatch/common/utils"
	pb "zrDispatch/core/proto"
	"zrDispatch/core/utils/define"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	maxWorkerTTL int64 = 20 // defaultHearbeatInterval = 20
)

const (
	DIE      string = "0"
	ALIVE    string = "1"
	Deployed string = "2"
	RUNNING  string = "3"
)

type HostInfo struct {
	Ip          string ` json:"ip,omitempty"`
	ServicePort int    ` json:"servicePort"`
	SshPort     int    ` json:"sshPort"`
	SshUser     string ` json:"sshUser"`
	SshPwd      string ` json:"sshPwd"`
	Weight      int    ` json:"weight,omitempty"`
	Hostname    string ` json:"hostname,omitempty"`
	Version     string ` json:"version,omitempty"`
	Hostgroup   string ` json:"hostgroup,omitempty"`
	Remark      string ` json:"remark,omitempty"`
}

// RegistryToUpdateHost refistry new host
func RegistryToUpdateHost(ctx context.Context, req *pb.RegistryReq) error {
	updatesql := `UPDATE host set weight=?,version=?,lastUpdateTimeUnix=?,remark=? WHERE addr=?`
	conn, err := db.GetConn(ctx)
	if err != nil {
		return fmt.Errorf("db.GetConn failed: %w", err)
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, updatesql)
	if err != nil {
		return fmt.Errorf("conn.PrepareContext failed: %w", err)
	}
	defer stmt.Close()
	addr := fmt.Sprintf("%s:%d", req.Ip, req.Port)
	_, err = stmt.ExecContext(ctx, req.Weight, req.Version, time.Now().Unix(), req.Remark, addr)
	if err != nil {
		return fmt.Errorf("stmt.ExecContext faled: %w", err)
	}

	return nil
}

// RegistryNewHost refistry new host
func RegistryNewHost(ctx context.Context, req *pb.RegistryReq) (string, error) {
	hostsql := `INSERT INTO host 
					(id,
					hostname,
					addr,
					weight,
					version,
					lastUpdateTimeUnix,
					remark
				)
 			  	VALUES
					(?,?,?,?,?,?,?)`
	addr := fmt.Sprintf("%s:%d", req.Ip, req.Port)
	hosts, _, err := getHosts(ctx, addr, nil, 0, 0)
	if err != nil {
		return "", err
	}
	if len(hosts) == 1 {
		log.Info("Addr Already Registry", zap.String("addr", addr))
		return "", nil
	}
	conn, err := db.GetConn(ctx)
	if err != nil {
		return "", fmt.Errorf("db.GetConn failed: %w", err)
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, hostsql)
	if err != nil {
		return "", fmt.Errorf("conn.PrepareContext failed: %w", err)
	}
	defer stmt.Close()
	id := utils.GetID()
	_, err = stmt.ExecContext(ctx,
		id,
		req.Hostname,
		addr,
		req.Weight,
		req.Version,
		time.Now().Unix(),
		req.Remark,
	)
	if err != nil {
		return "", fmt.Errorf("stmt.ExecContext failed: %w", err)
	}
	log.Info("New Client Registry ", zap.String("addr", addr))
	return id, nil
}

func CreateHost(ctx context.Context, hostInfo HostInfo) (string, error) {
	hostsql := `INSERT INTO host
				(
				ip,
				servicePort,
				sshPort,
				sshUser,
				sshPwd,
				hostname,
				weight,
				version,
				lastUpdateTimeUnix,
				remark
				)
				VALUES
				(?,?,?,?,?,?,?,?,?,?)`
	//addr := fmt.Sprintf("%s:%d", hostInfo.Ip, hostInfo.Port)
	hosts, _, err := getHosts(ctx, hostInfo.Ip, nil, 0, 0)
	if err != nil {
		return "", errors.New("读取主机失败")
	}
	if len(hosts) == 1 {
		log.Info("ip 已存在", zap.String("ip:", hostInfo.Ip))
		return "", errors.New("ip 已存在")
	}
	conn, err := db.GetConn(ctx)
	if err != nil {
		return "", fmt.Errorf("db.GetConn failed: %w", err)
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, hostsql)
	if err != nil {
		return "", fmt.Errorf("conn.PrepareContext failed: %w", err)
	}
	defer stmt.Close()
	id := utils.GetID()

	_, err = stmt.ExecContext(ctx,
		hostInfo.Ip,
		hostInfo.ServicePort,
		hostInfo.SshPort,
		hostInfo.SshUser,
		hostInfo.SshPwd,
		hostInfo.Hostname,
		hostInfo.Weight,
		hostInfo.Version,
		time.Now().Unix(),
		hostInfo.Remark,
	)
	if err != nil {
		return "", fmt.Errorf("stmt.ExecContext failed: %w", err)
	}
	log.Info("New Client Registry ", zap.String("ip:", hostInfo.Ip))
	return id, nil
}

// update host last recv hearbeat time
func UpdateHostHearbeat(ctx context.Context, ip string, servicePort int, time, version, runningTasks string) error {
	updatesql := `UPDATE host set status=3, lastUpdateTimeUnix=?,runningTasks=?,version=? WHERE ip=? and servicePort=?`
	conn, err := db.GetConn(ctx)
	if err != nil {
		return fmt.Errorf("db.GetConn failed: %w", err)
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, updatesql)
	if err != nil {
		return fmt.Errorf("conn.PrepareContext failed: %w", err)
	}
	defer stmt.Close()
	result, err := stmt.ExecContext(ctx, time, runningTasks, version, ip, servicePort)
	if err != nil {
		return fmt.Errorf("stmt.ExecContext failed: %w", err)
	}
	line, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("result.RowsAffected failed: %w", err)
	}
	if line <= 0 {
		//return fmt.Errorf("host %s not registry, may be this host is delete", fmt.Sprintf("%s:%d", ip, servicePort))
	}
	return nil
}

// get host by addr or id
func getHosts(ctx context.Context, ip string, ids []string, offset, limit int) ([]*define.Host, int, error) {
	getsql := `SELECT 
					id,
					ip,
					status,
					servicePort,
					sshPort,
					sshUser,
					sshPwd,
					hostname,
					runningTasks,
					weight,
					stop,
					version,
					lastUpdateTimeUnix,
					remark
			   FROM 
					host`
	var (
		count int
	)
	args := []interface{}{}
	// only use addr or ids query
	if ip != "" && len(ids) != 0 {
		return nil, 0, errors.New("only use addr or ids query")
	}
	if ip != "" {
		getsql += " WHERE ip=?"
		args = append(args, ip)
	}

	if len(ids) > 0 {
		var querys = []string{}
		for _, id := range ids {
			querys = append(querys, "id=?")
			args = append(args, id)
		}
		getsql += " WHERE " + strings.Join(querys, " OR ")

	}

	getsql += " ORDER BY lastUpdateTimeUnix asc"

	if limit > 0 {
		var err error
		count, err = countColums(ctx, getsql, args...)
		if err != nil {
			return nil, 0, fmt.Errorf("countColums failed: %w", err)
		}
		getsql += " LIMIT ? OFFSET ?"
		args = append(args, limit, offset)
	}

	conn, err := db.GetConn(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("db.GetConn failed: %w", err)
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, getsql)
	if err != nil {
		return nil, 0, fmt.Errorf("conn.PrepareContext failed: %w", err)
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("stmt.QueryContext failed: %w", err)
	}
	defer rows.Close()

	hosts := []*define.Host{}
	for rows.Next() {
		var (
			h     define.Host
			rtask string
		)
		err := rows.Scan(
			&h.ID,
			&h.Ip,
			&h.Status,
			&h.ServicePort,
			&h.SshPort,
			&h.SshUser,
			&h.SshPwd,
			&h.HostName,
			&rtask,
			&h.Weight,
			&h.Stop,
			&h.Version,
			&h.LastUpdateTimeUnix,
			&h.Remark)
		if err != nil {
			log.Error("Scan failed", zap.Error(err))
			continue
		}
		h.RunningTasks = []string{}
		if rtask != "" {
			h.RunningTasks = append(h.RunningTasks, strings.Split(rtask, ",")...)
		}
		if h.LastUpdateTimeUnix+maxWorkerTTL > time.Now().Unix() {
			h.Online = true
		}
		h.LastUpdateTime = utils.UnixToStr(h.LastUpdateTimeUnix)
		hosts = append(hosts, &h)
	}
	return hosts, count, nil
}

// GetHosts get all hosts
func GetHosts(ctx context.Context, offset, limit int) ([]*define.Host, int, error) {
	return getHosts(ctx, "", nil, offset, limit)
}

// GetHosts get all hosts
func GetHostsWithStatus(ctx context.Context, Symbol, status string) ([]*define.Host, int, error) {

	getsql := `SELECT 
					id,
					ip,
					status,
					servicePort,
					sshPort,
					sshUser,
					sshPwd,
					hostname,
					runningTasks,
					weight,
					stop,
					version,
					lastUpdateTimeUnix,
					remark
			   FROM 
					host`
	getsql = getsql + " where status " + Symbol + status

	order := " order by id asc"
	if time.Now().Unix()%2 == 0 {
		order = " order by id desc"
	}

	getsql = getsql + order
	var (
		count int
	)
	args := []interface{}{}

	conn, err := db.GetConn(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("db.GetConn failed: %w", err)
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, getsql)
	if err != nil {
		return nil, 0, fmt.Errorf("conn.PrepareContext failed: %w", err)
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("stmt.QueryContext failed: %w", err)
	}
	defer rows.Close()

	hosts := []*define.Host{}
	for rows.Next() {
		var (
			h     define.Host
			rtask string
		)
		err := rows.Scan(
			&h.ID,
			&h.Ip,
			&h.Status,
			&h.ServicePort,
			&h.SshPort,
			&h.SshUser,
			&h.SshPwd,
			&h.HostName,
			&rtask,
			&h.Weight,
			&h.Stop,
			&h.Version,
			&h.LastUpdateTimeUnix,
			&h.Remark)
		if err != nil {
			log.Error("Scan failed", zap.Error(err))
			continue
		}
		h.RunningTasks = []string{}
		if rtask != "" {
			h.RunningTasks = append(h.RunningTasks, strings.Split(rtask, ",")...)
		}
		if h.LastUpdateTimeUnix+maxWorkerTTL > time.Now().Unix() {
			h.Online = true
		}
		h.LastUpdateTime = utils.UnixToStr(h.LastUpdateTimeUnix)
		hosts = append(hosts, &h)
	}

	count, err = countColums(ctx, getsql, args...)

	return hosts, count, nil
}

// GetHostByAddr get host by addr
func GetHostByAddr(ctx context.Context, addr string) (*define.Host, error) {
	hosts, _, err := getHosts(ctx, addr, nil, 0, 0)
	if err != nil {
		return nil, err
	}
	if len(hosts) != 1 {
		return nil, errors.New("can not find hostid")
	}
	return hosts[0], nil
}

// ExistAddr check already exist
func ExistAddr(ctx context.Context, addr string) (*define.Host, bool, error) {
	hosts, _, err := getHosts(ctx, addr, nil, 0, 0)
	if err != nil {
		return nil, false, err
	}
	if len(hosts) != 1 {
		return nil, false, nil
	}
	return hosts[0], true, nil
}

// GetHostByID get host by hostid
func GetHostByID(ctx context.Context, id string) (*define.Host, error) {
	hosts, _, err := getHosts(ctx, "", []string{id}, 0, 0)
	if err != nil {
		return nil, err
	}
	if len(hosts) != 1 {
		log.Warn("can not find hostid", zap.Error(err))
		err = define.ErrNotExist{Value: id}
		return nil, err
	}
	return hosts[0], nil
}

// GetHostByIDS get hosts by hostids
func GetHostByIDS(ctx context.Context, ids []string) ([]*define.Host, error) {
	hosts, _, err := getHosts(ctx, "", ids, 0, 0)
	if err != nil {
		return nil, err
	}
	return hosts, nil
}

// StopHost will stop run worker in hostid
func StopHost(ctx context.Context, hostid string, stop bool) error {
	stopsql := `UPDATE host SET stop=? WHERE id=?`
	conn, err := db.GetConn(ctx)
	if err != nil {
		return fmt.Errorf("db.GetConn failed: %w", err)
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, stopsql)
	if err != nil {
		return fmt.Errorf("conn.PrepareContext failed: %w", err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, stop, hostid)
	if err != nil {
		return fmt.Errorf("stmt.ExecContext failed: %w", err)
	}
	return nil
}

// DeleteHost will delete host
func DeleteHost(ctx context.Context, hostid string) error {
	err := StopHost(ctx, hostid, true)
	if err != nil {
		return fmt.Errorf("StopHost failed: %w", err)
	}
	deletehostsql := `DELETE FROM host WHERE id=?`
	conn, err := db.GetConn(ctx)
	if err != nil {
		return fmt.Errorf("db.GetConn failed: %w", err)
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, deletehostsql)
	if err != nil {
		return fmt.Errorf("conn.PrepareContext failed: %w", err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, hostid)
	if err != nil {
		return fmt.Errorf("stmt.ExecContext failed: %w", err)
	}
	return nil
}

// delete from slice
func deletefromslice(deleteid string, ids []string) ([]string, bool) {
	var existid = -1
	for index, id := range ids {
		if id == deleteid {
			existid = index
			break
		}
	}
	if existid == -1 {
		// no found delete id
		return ids, false
	}
	ids = append(ids[:existid], ids[existid+1:]...)
	return ids, true
}
