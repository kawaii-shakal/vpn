package main

import "github.com/songgao/water"

func createTunInterface() (*water.Interface, error) {
	config := water.Config{
		DeviceType: water.TUN,
	}
	//config.Name = "vpn0"

	ifce, err := water.New(config)
	if err != nil {
		return nil, err
	}
	return ifce, nil
}
