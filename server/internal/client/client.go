package client

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"gu-universe/internal/models"
	"io"
	"net"
	"os"
	"time"
)

type Client struct {
	Player models.Player

	buf     [1024]byte
	conn    *net.UDPConn
	logFile *os.File
}

func New() Client {
	return Client{
		Player: models.NewPlayer(nil),
	}
}

func (c Client) receive() (string, error) {
	msgRaw, err := bufio.NewReader(c.conn).ReadString('\n')
	if err != nil {
		return "", err
	}

	msg := fmt.Sprintf("[%s] %s", time.Now().UTC().Format("2006-01-02 15:04:05.000000000"), string(msgRaw))
	c.logFile.WriteString(msg)
	fmt.Print(msg)

	return msg, nil
}

func (c Client) send(msg string) error {
	_, err := c.conn.Write([]byte(msg))
	if err != nil {
		return err
	}

	return nil
}

func (c Client) Start(addr string) error {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}

	c.conn, err = net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return err
	}

	logTime := time.Now().UTC().Format("2006-01-02T15:04:05.000000000")
	tmpLogSha := sha1.New()
	_, err = io.WriteString(tmpLogSha, logTime)
	if err != nil {
		return err
	}
	logID := string(hex.EncodeToString(tmpLogSha.Sum(nil)))
	c.logFile, err = os.Create(fmt.Sprintf("./clilogs/%s_%s.log", logTime, logID))
	if err != nil {
		return err
	}
	defer c.logFile.Close()

	var (
	// msg string
	)
	for {
		err = c.send(logID)
		if err != nil {
			return err
		}

		_, err = c.receive()
		if err != nil {
			return err
		}

		time.Sleep(5 * time.Second)
	}
}
