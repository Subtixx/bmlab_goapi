package main

import (
	"crypto/tls"
	"github.com/jlaffaye/ftp"
	"log"
	"strconv"
)

type BambuLabFTP struct {
	BambuLabAPI *BambuLabAPI

	FTPClient *ftp.ServerConn
}

func NewBambuLabFTP(
	bambuLabAPI *BambuLabAPI,
) *BambuLabFTP {
	return &BambuLabFTP{
		BambuLabAPI: bambuLabAPI,
	}
}

func (c *BambuLabFTP) getFtpClient() (*ftp.ServerConn, error) {
	dialOption := ftp.DialWithTLS(&tls.Config{InsecureSkipVerify: true})
	return ftp.Dial(c.BambuLabAPI.IP+":"+strconv.Itoa(int(c.BambuLabAPI.FTPPort)), dialOption)
}

func (c *BambuLabFTP) connectFTP() bool {
	if c.FTPClient == nil {
		c.FTPClient, _ = c.getFtpClient()
	}

	if err := c.FTPClient.Login(c.BambuLabAPI.Username, c.BambuLabAPI.AccessCode); err != nil {
		log.Fatal("Error connecting to FTP server:", err)
		return false
	}

	return true
}
