package main

type NozzleInfo struct {
	NozzleTemp       float64
	NozzleType       string
	NozzleDiameter   float64
	NozzleTargetTemp float64
}

type BedInfo struct {
	BedTemp       float64
	BedTargetTemp float64
}

type FanInfo struct {
	HeatbreakFanSpeed string

	CoolingFanSpeed string
	BigFan1Speed    string
	BigFan2Speed    string
}

type AmsInfo struct {
	AmsStatus        int
	AmsRfidStatus    int
	AmsExistBits     string
	TrayExistBits    string
	TrayIsBblBits    string
	TrayTar          string
	TrayNow          string
	TrayPre          string
	TrayReadDoneBits string
	TrayReadingBits  string
	Version          int
	InsertFlag       bool
	PowerOnFlag      bool
}

type PrinterInfo struct {
	WiFiSignal string

	NozzleInfo *NozzleInfo
	BedInfo    *BedInfo
	FanInfo    *FanInfo

	AmsInfo *AmsInfo

	RemainingTime int
	Percent       int
	PrintStage    string

	LayerNum      int
	TotalLayerNum int
}

type BambuLabAPI struct {
	IP           string
	FTPPort      uint16
	MQTTPort     uint16
	SerialNumber string
	AccessCode   string
	Username     string

	BambuLabMQTT *BambuLabMQTT
	BambuLabFTP  *BambuLabFTP

	PrinterStatus *PrinterInfo
	
	OnPrinterStatusChange func(*PrinterInfo)
}

func NewBambuLabAPI(
	IP string,
	SerialNumber string,
	AccessCode string,
) *BambuLabAPI {
	api := &BambuLabAPI{
		IP:           IP,
		FTPPort:      990,
		MQTTPort:     8883,
		SerialNumber: SerialNumber,
		AccessCode:   AccessCode,
		Username:     "bblp",
		PrinterStatus: &PrinterInfo{
			WiFiSignal:    "",
			NozzleInfo:    &NozzleInfo{},
			BedInfo:       &BedInfo{},
			FanInfo:       &FanInfo{},
			AmsInfo:       &AmsInfo{},
			RemainingTime: 0,
			Percent:       0,
			PrintStage:    "",
			LayerNum:      0,
			TotalLayerNum: 0,
		},
	}
	api.BambuLabMQTT = NewBambuLabMQTT(api)
	api.BambuLabFTP = NewBambuLabFTP(api)
	return api
}

func (api *BambuLabAPI) setFTPPort(port uint16) {
	api.FTPPort = port
}

func (api *BambuLabAPI) setMQTTPort(port uint16) {
	api.MQTTPort = port
}

func (api *BambuLabAPI) setUsername(username string) {
	api.Username = username
}
