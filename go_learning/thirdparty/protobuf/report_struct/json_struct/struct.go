package json_struct

// Event event字段
type Event struct	{
	Utm			string		`json:"utm"`
	Scm			string		`json:"scm"`
	Spm			string		`json:"spm"`
	Spmref		string		`json:"spmref"`
	Uri			string		`json:"uri"`
	Uriref		string		`json:"uriref"`
	Params		string		`json:"params"`
}

// ReportItem 单条日志格式
type ReportItem struct {
	UserId		string		`json:"user_id"`
	ClientTime	int64		`json:"client_time"`
	Session		string		`json:"session"`
	EventId		int64		`json:"event_id"`
	EventType	string		`json:"event_type"`
	Event		Event		`json:"event"`
	UserAgent	string		`json:"user_agent"`
} 

// ReportInputParams 数据指上报接口，请求体结构
type ReportInputParams struct {
	Data		[]ReportItem	`json:"data"`
}
