package probe

const (
	HTTP     = "HTTP"
	TCP      = "TCP"
	FTP      = "FTP"
	SMTP     = "SMTP"
	PING     = "PING"
	DNS      = "DNS"
	SMTPSend = "SMTP_SEND"
)

type Probe struct {
	ID         string     `json:"id,omitempty"`
	Type       string     `json:"type,omitempty"`
	PoolRecord string     `json:"poolRecord,omitempty"`
	Interval   string     `json:"interval,omitempty"`
	Agents     []string   `json:"agents,omitempty"`
	Threshold  int        `json:"threshold,omitempty"`
	Details    RawDetails `json:"details,omitempty"`
}

type RawDetails interface{}
