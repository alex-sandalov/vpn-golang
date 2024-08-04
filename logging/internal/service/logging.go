package logging

import (
	"fmt"

	"golang.zx2c4.com/wireguard/wgctrl"
)

type LoggingService struct {
}

func NewLoggingService() *LoggingService {
	return &LoggingService{}
}

func (s *LoggingService) checkWireGuard() error {
	client, err := wgctrl.New()
	if err != nil {
		return fmt.Errorf("failed to connect to WireGuard: %s", err)
	}
	defer client.Close()

	device, err := client.Device("wg0")
	if err != nil {
		return fmt.Errorf("failed to get WireGuard device: %s", err)
	}

	for _, peer := range device.Peers {
		fmt.Println(peer)
	}

	return nil
}

func (s *LoggingService) Run() {
	s.checkWireGuard()
}
