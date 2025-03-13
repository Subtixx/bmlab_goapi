package main

import (
	"crypto/tls"
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"strconv"
)

type Packet struct {
	//Print PrintStatus `json:"print"`
	//liveview
	//info

	Liveview interface{} `json:"liveview"`
	Info     interface{} `json:"info"`
	Print    interface{} `json:"print"`
}

type LiveViewPacket struct {
	Liveview struct {
		SequenceId string `json:"sequence_id"`
		Command    string `json:"command"`
		Timestamp  int64  `json:"timestamp"`
		Ttcode     string `json:"ttcode"`
		Region     string `json:"region"`
		Reason     string `json:"reason"`
		Result     string `json:"result"`
	} `json:"liveview"`
}

type InfoPacket struct {
	Info struct {
		Command    string `json:"command"`
		SequenceId string `json:"sequence_id"`
		Module     []struct {
			Name        string `json:"name"`
			ProjectName string `json:"project_name"`
			SwVer       string `json:"sw_ver"`
			HwVer       string `json:"hw_ver"`
			Sn          string `json:"sn"`
			Flag        int    `json:"flag"`
			LoaderVer   string `json:"loader_ver,omitempty"`
		} `json:"module"`
		Result string `json:"result"`
		Reason string `json:"reason"`
	} `json:"info"`
}

type PrintStatusPacket struct {
	PrintStatus struct {
		Ipcam struct {
			IpcamDev    string `json:"ipcam_dev"`
			IpcamRecord string `json:"ipcam_record"`
			Timelapse   string `json:"timelapse"`
			Resolution  string `json:"resolution"`
			TutkServer  string `json:"tutk_server"`
			ModeBits    int    `json:"mode_bits"`
		} `json:"ipcam"`
		Upload struct {
			Status   string `json:"status"`
			Progress int    `json:"progress"`
			Message  string `json:"message"`
		} `json:"upload"`
		Net struct {
			Conf int `json:"conf"`
			Info []struct {
				Ip   int `json:"ip"`
				Mask int `json:"mask"`
			} `json:"info"`
		} `json:"net"`
		NozzleTemper            float64       `json:"nozzle_temper"`
		NozzleTargetTemper      float64       `json:"nozzle_target_temper"`
		BedTemper               float64       `json:"bed_temper"`
		BedTargetTemper         float64       `json:"bed_target_temper"`
		ChamberTemper           float64       `json:"chamber_temper"`
		McPrintStage            string        `json:"mc_print_stage"`
		HeatbreakFanSpeed       string        `json:"heatbreak_fan_speed"`
		CoolingFanSpeed         string        `json:"cooling_fan_speed"`
		BigFan1Speed            string        `json:"big_fan1_speed"`
		BigFan2Speed            string        `json:"big_fan2_speed"`
		McPercent               int           `json:"mc_percent"`
		McRemainingTime         int           `json:"mc_remaining_time"`
		AmsStatus               int           `json:"ams_status"`
		AmsRfidStatus           int           `json:"ams_rfid_status"`
		HwSwitchState           int           `json:"hw_switch_state"`
		SpdMag                  int           `json:"spd_mag"`
		SpdLvl                  int           `json:"spd_lvl"`
		PrintError              int           `json:"print_error"`
		Lifecycle               string        `json:"lifecycle"`
		WifiSignal              string        `json:"wifi_signal"`
		GcodeState              string        `json:"gcode_state"`
		GcodeFilePreparePercent string        `json:"gcode_file_prepare_percent"`
		QueueNumber             int           `json:"queue_number"`
		QueueTotal              int           `json:"queue_total"`
		QueueEst                int           `json:"queue_est"`
		QueueSts                int           `json:"queue_sts"`
		ProjectId               string        `json:"project_id"`
		ProfileId               string        `json:"profile_id"`
		TaskId                  string        `json:"task_id"`
		SubtaskId               string        `json:"subtask_id"`
		SubtaskName             string        `json:"subtask_name"`
		GcodeFile               string        `json:"gcode_file"`
		Stg                     []int         `json:"stg"`
		StgCur                  int           `json:"stg_cur"`
		PrintType               string        `json:"print_type"`
		HomeFlag                int           `json:"home_flag"`
		McPrintLineNumber       string        `json:"mc_print_line_number"`
		McPrintSubStage         int           `json:"mc_print_sub_stage"`
		Sdcard                  bool          `json:"sdcard"`
		ForceUpgrade            bool          `json:"force_upgrade"`
		MessProductionState     string        `json:"mess_production_state"`
		LayerNum                int           `json:"layer_num"`
		TotalLayerNum           int           `json:"total_layer_num"`
		SObj                    []interface{} `json:"s_obj"`
		FilamBak                []interface{} `json:"filam_bak"`
		FanGear                 int           `json:"fan_gear"`
		NozzleDiameter          string        `json:"nozzle_diameter"`
		NozzleType              string        `json:"nozzle_type"`
		CaliVersion             int           `json:"cali_version"`
		K                       string        `json:"k"`
		Flag3                   int           `json:"flag3"`
		UpgradeState            struct {
			SequenceId         int           `json:"sequence_id"`
			Progress           string        `json:"progress"`
			Status             string        `json:"status"`
			ConsistencyRequest bool          `json:"consistency_request"`
			DisState           int           `json:"dis_state"`
			ErrCode            int           `json:"err_code"`
			ForceUpgrade       bool          `json:"force_upgrade"`
			Message            string        `json:"message"`
			Module             string        `json:"module"`
			NewVersionState    int           `json:"new_version_state"`
			CurStateCode       int           `json:"cur_state_code"`
			Idx2               int           `json:"idx2"`
			NewVerList         []interface{} `json:"new_ver_list"`
		} `json:"upgrade_state"`
		Hms    []interface{} `json:"hms"`
		Online struct {
			Ahb     bool `json:"ahb"`
			Rfid    bool `json:"rfid"`
			Version int  `json:"version"`
		} `json:"online"`
		Ams struct {
			Ams              []interface{} `json:"ams"`
			AmsExistBits     string        `json:"ams_exist_bits"`
			TrayExistBits    string        `json:"tray_exist_bits"`
			TrayIsBblBits    string        `json:"tray_is_bbl_bits"`
			TrayTar          string        `json:"tray_tar"`
			TrayNow          string        `json:"tray_now"`
			TrayPre          string        `json:"tray_pre"`
			TrayReadDoneBits string        `json:"tray_read_done_bits"`
			TrayReadingBits  string        `json:"tray_reading_bits"`
			Version          int           `json:"version"`
			InsertFlag       bool          `json:"insert_flag"`
			PowerOnFlag      bool          `json:"power_on_flag"`
		} `json:"ams"`
		Xcam struct {
			BuildplateMarkerDetector bool `json:"buildplate_marker_detector"`
		} `json:"xcam"`
		VtTray struct {
			Id            string  `json:"id"`
			TagUid        string  `json:"tag_uid"`
			TrayIdName    string  `json:"tray_id_name"`
			TrayInfoIdx   string  `json:"tray_info_idx"`
			TrayType      string  `json:"tray_type"`
			TraySubBrands string  `json:"tray_sub_brands"`
			TrayColor     string  `json:"tray_color"`
			TrayWeight    string  `json:"tray_weight"`
			TrayDiameter  string  `json:"tray_diameter"`
			TrayTemp      string  `json:"tray_temp"`
			TrayTime      string  `json:"tray_time"`
			BedTempType   string  `json:"bed_temp_type"`
			BedTemp       string  `json:"bed_temp"`
			NozzleTempMax string  `json:"nozzle_temp_max"`
			NozzleTempMin string  `json:"nozzle_temp_min"`
			XcamInfo      string  `json:"xcam_info"`
			TrayUuid      string  `json:"tray_uuid"`
			Remain        int     `json:"remain"`
			K             float64 `json:"k"`
			N             int     `json:"n"`
			CaliIdx       int     `json:"cali_idx"`
		} `json:"vt_tray"`
		LightsReport []struct {
			Node string `json:"node"`
			Mode string `json:"mode"`
		} `json:"lights_report"`
		Command    string `json:"command"`
		Msg        int    `json:"msg"`
		SequenceId string `json:"sequence_id"`
	} `json:"print"`
}

type BambuLabMQTT struct {
	BambuLabAPI *BambuLabAPI
	MQTTClient  mqtt.Client
}

func NewBambuLabMQTT(
	bambuLabAPI *BambuLabAPI,
) *BambuLabMQTT {
	return &BambuLabMQTT{
		BambuLabAPI: bambuLabAPI,
	}
}

func (c *BambuLabMQTT) getMqttClient() mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("ssl://" + c.BambuLabAPI.IP + ":" + strconv.Itoa(int(c.BambuLabAPI.MQTTPort)))
	opts.SetUsername(c.BambuLabAPI.Username)
	opts.SetPassword(c.BambuLabAPI.AccessCode)
	opts.SetTLSConfig(&tls.Config{InsecureSkipVerify: true})
	opts.ProtocolVersion = 4

	mqttClient := mqtt.NewClient(opts)
	return mqttClient
}

func (c *BambuLabMQTT) Disconnect() {
	if c.MQTTClient != nil {
		c.MQTTClient.Disconnect(250)
	}
}

func (c *BambuLabMQTT) FillPrinterStatus(printPacket *PrintStatusPacket) {

	if printPacket.PrintStatus.BedTemper > 0 {
		c.BambuLabAPI.PrinterStatus.BedInfo.BedTemp = printPacket.PrintStatus.BedTemper
	}

	if printPacket.PrintStatus.BedTargetTemper > 0 {
		c.BambuLabAPI.PrinterStatus.BedInfo.BedTargetTemp = printPacket.PrintStatus.BedTargetTemper
	}

	if printPacket.PrintStatus.NozzleTemper > 0 {
		c.BambuLabAPI.PrinterStatus.NozzleInfo.NozzleTemp = printPacket.PrintStatus.NozzleTemper
	}

	if printPacket.PrintStatus.NozzleType != "" {
		c.BambuLabAPI.PrinterStatus.NozzleInfo.NozzleType = printPacket.PrintStatus.NozzleType
	}
	if printPacket.PrintStatus.NozzleDiameter != "" {
		nozzleDiameter, _ := strconv.ParseFloat(printPacket.PrintStatus.NozzleDiameter, 64)
		c.BambuLabAPI.PrinterStatus.NozzleInfo.NozzleDiameter = nozzleDiameter
	}

	if printPacket.PrintStatus.NozzleTargetTemper > 0 {
		c.BambuLabAPI.PrinterStatus.NozzleInfo.NozzleTargetTemp = printPacket.PrintStatus.NozzleTargetTemper
	}

	if printPacket.PrintStatus.CoolingFanSpeed != "" {
		c.BambuLabAPI.PrinterStatus.FanInfo.CoolingFanSpeed = printPacket.PrintStatus.CoolingFanSpeed
	}

	if printPacket.PrintStatus.BigFan1Speed != "" {
		c.BambuLabAPI.PrinterStatus.FanInfo.BigFan1Speed = printPacket.PrintStatus.BigFan1Speed
	}

	if printPacket.PrintStatus.BigFan2Speed != "" {
		c.BambuLabAPI.PrinterStatus.FanInfo.BigFan2Speed = printPacket.PrintStatus.BigFan2Speed
	}

	if printPacket.PrintStatus.HeatbreakFanSpeed != "" {
		c.BambuLabAPI.PrinterStatus.FanInfo.HeatbreakFanSpeed = printPacket.PrintStatus.HeatbreakFanSpeed
	}

	if printPacket.PrintStatus.WifiSignal != "" {
		c.BambuLabAPI.PrinterStatus.WiFiSignal = printPacket.PrintStatus.WifiSignal
	}

	if printPacket.PrintStatus.McPrintStage != "" {
		c.BambuLabAPI.PrinterStatus.PrintStage = printPacket.PrintStatus.McPrintStage
	}

	if printPacket.PrintStatus.McPercent > 0 {
		c.BambuLabAPI.PrinterStatus.Percent = printPacket.PrintStatus.McPercent
	}

	if printPacket.PrintStatus.McRemainingTime > 0 {
		c.BambuLabAPI.PrinterStatus.RemainingTime = printPacket.PrintStatus.McRemainingTime
	}

	if printPacket.PrintStatus.LayerNum > 0 {
		c.BambuLabAPI.PrinterStatus.LayerNum = printPacket.PrintStatus.LayerNum
	}

	if printPacket.PrintStatus.TotalLayerNum > 0 {
		c.BambuLabAPI.PrinterStatus.TotalLayerNum = printPacket.PrintStatus.TotalLayerNum
	}

	if printPacket.PrintStatus.AmsStatus > 0 {
		c.BambuLabAPI.PrinterStatus.AmsInfo.AmsStatus = printPacket.PrintStatus.AmsStatus
		c.BambuLabAPI.PrinterStatus.AmsInfo.AmsRfidStatus = printPacket.PrintStatus.AmsRfidStatus
	}

	if printPacket.PrintStatus.Ams.AmsExistBits != "" {
		c.BambuLabAPI.PrinterStatus.AmsInfo.AmsExistBits = printPacket.PrintStatus.Ams.AmsExistBits
	}

	if c.BambuLabAPI.PrinterStatus.AmsInfo.TrayExistBits != "" {
		c.BambuLabAPI.PrinterStatus.AmsInfo.TrayExistBits = printPacket.PrintStatus.Ams.TrayExistBits
	}
	if c.BambuLabAPI.PrinterStatus.AmsInfo.TrayIsBblBits != "" {
		c.BambuLabAPI.PrinterStatus.AmsInfo.TrayIsBblBits = printPacket.PrintStatus.Ams.TrayIsBblBits
	}
	if c.BambuLabAPI.PrinterStatus.AmsInfo.TrayTar != "" {
		c.BambuLabAPI.PrinterStatus.AmsInfo.TrayTar = printPacket.PrintStatus.Ams.TrayTar
	}
	if c.BambuLabAPI.PrinterStatus.AmsInfo.TrayNow != "" {
		c.BambuLabAPI.PrinterStatus.AmsInfo.TrayNow = printPacket.PrintStatus.Ams.TrayNow
	}
	if c.BambuLabAPI.PrinterStatus.AmsInfo.TrayPre != "" {
		c.BambuLabAPI.PrinterStatus.AmsInfo.TrayPre = printPacket.PrintStatus.Ams.TrayPre
	}
	if c.BambuLabAPI.PrinterStatus.AmsInfo.TrayReadDoneBits != "" {
		c.BambuLabAPI.PrinterStatus.AmsInfo.TrayReadDoneBits = printPacket.PrintStatus.Ams.TrayReadDoneBits
	}
	if c.BambuLabAPI.PrinterStatus.AmsInfo.TrayReadingBits != "" {
		c.BambuLabAPI.PrinterStatus.AmsInfo.TrayReadingBits = printPacket.PrintStatus.Ams.TrayReadingBits
	}
	if c.BambuLabAPI.PrinterStatus.AmsInfo.Version > 0 {
		c.BambuLabAPI.PrinterStatus.AmsInfo.Version = printPacket.PrintStatus.Ams.Version
	}

	if c.BambuLabAPI.OnPrinterStatusChange != nil {
		c.BambuLabAPI.OnPrinterStatusChange(c.BambuLabAPI.PrinterStatus)
	}
}

func (c *BambuLabMQTT) onMessage(client mqtt.Client, msg mqtt.Message) {
	log.Println("Received message:", string(msg.Payload()))

	// Get which Packet by checking if it has "print" or "liveview" or "info" in the payload
	var packet = Packet{
		Liveview: nil,
		Info:     nil,
		Print:    nil,
	}
	if err := json.Unmarshal(msg.Payload(), &packet); err != nil {
		log.Fatal("Error unmarshalling base packet:", err)
	}

	if packet.Liveview != nil {
		var liveviewPacket LiveViewPacket
		if err := json.Unmarshal(msg.Payload(), &liveviewPacket); err != nil {
			log.Fatal("Error unmarshalling liveview packet:", err)
		}
		//c.BambuLabAPI.OnLiveview(liveviewPacket.Liveview)
	} else if packet.Info != nil {
		var infoPacket InfoPacket
		if err := json.Unmarshal(msg.Payload(), &infoPacket); err != nil {
			log.Fatal("Error unmarshalling info packet:", err)
		}
		//c.BambuLabAPI.OnInfo(infoPacket.Info)
	} else if packet.Print != nil {
		var printPacket PrintStatusPacket
		if err := json.Unmarshal(msg.Payload(), &printPacket); err != nil {
			log.Fatal("Error unmarshalling print packet:", err)
		}
		c.FillPrinterStatus(&printPacket)
	} else {
		log.Fatal("Unknown packet type:", packet, string(msg.Payload()))
	}
}

func (c *BambuLabMQTT) RequestFullStatus() bool {
	if token := c.MQTTClient.Publish(
		"device/"+c.BambuLabAPI.SerialNumber+"/request",
		0,
		false,
		`{"pushing": {"command": "pushall"}}`,
	); token.Wait() && token.Error() != nil {
		log.Fatal("Error requesting current status:", token.Error())
		return false
	}

	return true
}

func (c *BambuLabMQTT) connectMQTT() bool {
	if c.MQTTClient == nil {
		c.MQTTClient = c.getMqttClient()
	}

	if token := c.MQTTClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("Error connecting to MQTT broker:", token.Error())
		return false
	}

	if token := c.MQTTClient.Subscribe("device/"+c.BambuLabAPI.SerialNumber+"/report", 0, c.onMessage); token.Wait() && token.Error() != nil {
		log.Fatal("Error subscribing to MQTT topic:", token.Error())
		return false
	}

	if c.RequestFullStatus() == false {
		return false
	}

	return true
}
