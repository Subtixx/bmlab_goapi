package main

import (
	"github.com/joho/godotenv"
	"log"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	config := AppConfig{}
	config.Load()

	config.print()

	api := NewBambuLabAPI(
		config.IP,
		config.SerialNumber,
		config.AccessCode,
	)
	if api.BambuLabMQTT.connectMQTT() == false {
		log.Fatal("Error connecting to MQTT broker")
	}
	if api.BambuLabFTP.connectFTP() == false {
		log.Fatal("Error connecting to FTP server")
	}

	for {
		if api.PrinterStatus != nil && api.PrinterStatus.BedInfo != nil {
			log.Println(api.PrinterStatus.BedInfo.BedTemp)
		}

		// Sleep for a second
		time.Sleep(1 * time.Second)
	}

	/*mqtt.ERROR = log.New(os.Stdout, "[ERROR] ", 0)
	mqtt.CRITICAL = log.New(os.Stdout, "[CRIT] ", 0)
	mqtt.WARN = log.New(os.Stdout, "[WARN]  ", 0)
	mqtt.DEBUG = log.New(os.Stdout, "[DEBUG] ", 0)

	mqttClient := AppConfig.getMqttClient()
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("Error connecting to MQTT broker:", token.Error())
	}
	//  f"device/{printer_serial}/request"
	if token := mqttClient.Subscribe("device/"+AppConfig.SerialNumber+"/report", 0, func(client mqtt.Client, msg mqtt.Message) {
		log.Println("Received message:", string(msg.Payload()))
	}); token.Wait() && token.Error() != nil {
		log.Fatal("Error subscribing to MQTT topic (report):", token.Error())
	}

	if token := mqttClient.Subscribe("device/"+AppConfig.SerialNumber+"/request", 0, func(client mqtt.Client, msg mqtt.Message) {
		log.Println("Received message:", string(msg.Payload()))
		// Handle the message here
	}); token.Wait() && token.Error() != nil {
		log.Fatal("Error subscribing to MQTT topic (request):", token.Error())
	}

	//self._client.publish(self.command_topic, json.dumps({"pushing": {"command": "pushall"}}))
	log.Println("Requesting full push")
	t := mqttClient.Publish("device/"+AppConfig.SerialNumber+"/request", 0x0, false, `{"pushing": {"command": "pushall"}}`)
	if t.Wait() && t.Error() != nil {
		log.Fatal(t.Error())
	}
	log.Println("Full push requested")
	//mc_percent


	// Loop while the MQTT client is connected
	for mqttClient.IsConnected() {
		// Sleep for a second
		time.Sleep(1 * time.Second)
	}

	// Explicit TLs
	dialOption := ftp.DialWithTLS(&tls.AppConfig{InsecureSkipVerify: true})

	log.Println("Connecting to FTP server")
	c, err := ftp.Dial(IP+":990", dialOption)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to FTP server")

	err = c.Login("bblp", ACCESS_CODE)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	// List the files in the current directory
	files, err := c.List("")
	if err != nil {
		log.Fatal("Error listing files:", err)
	}

	for _, file := range files {
		log.Println(file.Name)
	}

	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}*/
}
