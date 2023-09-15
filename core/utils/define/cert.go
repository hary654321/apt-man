package define

type TaskRes struct {
	Code int `json:"code"`
	Data []struct {
		ProbeResult struct {
			ReqInfo struct {
				Addr      string `json:"addr"`
				ProbeName string `json:"probe_name"`
			} `json:"req_info"`
			ResHex   string `json:"res_hex"`
			ResPlain string `json:"res_plain"`
		} `json:"probe_result"`
		SslResult struct {
			Cert    `json:"Cert"`
			Code    int         `json:"Code"`
			Data    interface{} `json:"Data"`
			Err     string      `json:"Err"`
			URLPort string      `json:"UrlPort"`
		} `json:"ssl_result"`
	} `json:"data"`
	Msg string `json:"msg"`
	Res []string
}

type Cert struct {
	// Ip              string `gorm:"column:ip"json:"ip"`
	// Port            string `gorm:"column:port"json:"port"`
	// Probe_name string `gorm:"column:probe_name"json:"probe_name"`
	CertBase64      string `gorm:"column:cert_base64"json:"cert_base64"`
	CertFingerprint string `gorm:"column:cert_fingerprint"json:"cert_fingerprint"`
	CertIssuer      string `gorm:"column:cert_issuer"json:"cert_issuer"`
	CertIssuerC     string `gorm:"column:cert_issuer_c" json:"cert_issuer_c"`
	CertIssuerCn    string `gorm:"column:cert_issuer_cn" json:"cert_issuer_cn"`
	CertIssuerO     string `gorm:"column:cert_issuer_o" json:"cert_issuer_o"`
	CertSerialno    string `gorm:"column:cert_serialno" json:"cert_serialno"`
	CertSubject     string `gorm:"column:cert_subject" json:"cert_subject"`
	CertSubjectC    string `gorm:"column:cert_subject_c" json:"cert_subject_c"`
	CertSubjectCn   string `gorm:"column:cert_subject_cn" json:"cert_subject_cn"`
	CertSubjectO    string `gorm:"column:cert_subject_o" json:"cert_subject_o"`
	ValidFrom       string `gorm:"column:valid_from" json:"valid_from"`
	ValidTo         string `gorm:"column:valid_to" json:"valid_to"`
	// RunTaskID       string `gorm:"column:run_task_id" json:"run_task_id"`
}

type CertRes struct {
	GetID
	Cert
	Matched int    `gorm:"column:matched" json:"matched"`
	Finger  string `gorm:"column:probe_recv" json:"finger" `
	GetCtime
}

type GetCtime struct {
	Ctime LocalTime `gorm:"column:create_time" json:"create_time"`
}
