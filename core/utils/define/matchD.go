package define

type MatchRes struct {
	Id                     int    `gorm:"column:id" json:"id"`
	Match_task_id          string `gorm:"column:match_task_id" json:"match_task_id" `
	Match_desc             string `gorm:"column:match_desc" json:"match_desc" `
	Match_ip               string `gorm:"column:match_ip" json:"match_ip" `
	Match_port             string `gorm:"column:match_port" json:"match_port" `
	Match_type             string `gorm:"column:match_type" json:"match_type" `
	Match_region           string `gorm:"column:match_region" json:"match_region" `
	Match_tags             string `gorm:"column:match_tags" json:"match_tags"`
	Match_probe_name       string `gorm:"column:match_probe_name" json:"match_probe_name"`
	Match_cert_fingerprint string `gorm:"column:match_cert_fingerprint" json:"match_cert_fingerprint"`
	Match_cert_subject     string `gorm:"column:match_cert_subject" json:"match_cert_subject"`
	Match_cert_issuer      string `gorm:"column:match_cert_issuer" json:"match_cert_issuer"`
	Match_cert_dns_names   string `gorm:"column:match_cert_dns_names" json:"match_cert_dns_names"`
	Match_cert_valid_from  string `gorm:"column:match_cert_valid_from" json:"match_cert_valid_from"`
	Match_cert_valid_to    string `gorm:"column:match_cert_valid_to" json:"match_cert_valid_to"`
	Match_cert_base64      string `gorm:"column:match_cert_base64" json:"match_cert_base64"`
	Match_create_time      string `gorm:"column:match_create_time" json:"match_create_time"`
	Match_update_time      string `gorm:"column:match_update_time" json:"match_update_time"`
	RunTaskID              string `gorm:"column:run_task_id" json:"run_task_id"`
}
