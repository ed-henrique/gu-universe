package server

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"gu-universe/internal/models"
	"io"
	"net"
	"os"
	"time"
)

type Server struct {
	Players []models.Player

	buf     [1024]byte
	conn    *net.UDPConn
	logFile *os.File
}

func New() Server {
	return Server{
		Players: make([]models.Player, 0),
	}
}

func (s Server) receive() (*net.UDPAddr, string, error) {
	_, addr, err := s.conn.ReadFromUDP(s.buf[0:])
	if err != nil {
		return nil, "", err
	}

	msg := fmt.Sprintf("[%s] %s ... %s\n", time.Now().UTC().Format("2006-01-02 15:04:05.000000000"), addr, string(s.buf[0:]))
	s.logFile.WriteString(msg)
	fmt.Print(msg)

	return addr, msg, nil
}

func (s Server) send(addr *net.UDPAddr, msg string) error {
	_, err := s.conn.WriteToUDP([]byte(msg), addr)
	if err != nil {
		return err
	}

	return nil
}

func (s Server) Start(addr string) error {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}

	s.conn, err = net.ListenUDP("udp", udpAddr)
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
	s.logFile, err = os.Create(fmt.Sprintf("./serlogs/%s_%s.log", logTime, logID))
	if err != nil {
		return err
	}
	defer s.logFile.Close()

	var (
		cliAddr *net.UDPAddr
		// msg string
	)
	for {
		cliAddr, _, err = s.receive()
		if err != nil {
			return err
		}

		err = s.send(cliAddr, "Test\n")
		if err != nil {
			return err
		}
	}
}
