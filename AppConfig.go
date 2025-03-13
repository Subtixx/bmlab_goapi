package main

import (
	"crypto/tls"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/jlaffaye/ftp"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type AppConfig struct {
	IP           string
	FtpPort      string
	MqttPort     string
	SerialNumber string
	AccessCode   string
	Username     string
}

func (c *AppConfig) Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	c.IP = os.Getenv("IP")
	if c.IP == "" {
		panic("IP is required")
	}
	c.FtpPort = os.Getenv("FTP_PORT")
	if c.FtpPort == "" {
		c.FtpPort = "990"
	}
	c.MqttPort = os.Getenv("MQTT_PORT")
	if c.MqttPort == "" {
		c.MqttPort = "8883"
	}
	c.SerialNumber = os.Getenv("SERIAL_NUMBER")
	if c.SerialNumber == "" {
		panic("SERIAL_NUMBER is required")
	}
	c.AccessCode = os.Getenv("ACCESS_CODE")
	if c.AccessCode == "" {
		panic("ACCESS_CODE is required")
	}
	c.Username = os.Getenv("CONN_USERNAME")
	if c.Username == "" {
		c.Username = "bblp"
	}
}

func (c *AppConfig) print() {
	log.Println("IP:", c.IP)
	log.Println("FTP_PORT:", c.FtpPort)
	log.Println("MQTT_PORT:", c.MqttPort)
	log.Println("SERIAL_NUMBER:", c.SerialNumber)
	log.Println("ACCESS_CODE:", c.AccessCode)
	log.Println("CONN_USERNAME:", c.Username)
}

func (c *AppConfig) getMqttClient() mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("ssl://" + c.IP + ":" + c.MqttPort)
	opts.SetUsername(c.Username)
	opts.SetPassword(c.AccessCode)
	opts.SetTLSConfig(&tls.Config{InsecureSkipVerify: true})
	opts.ProtocolVersion = 4

	mqttClient := mqtt.NewClient(opts)
	return mqttClient
}

func (c *AppConfig) getFtpClient() (*ftp.ServerConn, error) {
	dialOption := ftp.DialWithTLS(&tls.Config{InsecureSkipVerify: true})
	return ftp.Dial(c.IP+":"+c.FtpPort, dialOption)
}
