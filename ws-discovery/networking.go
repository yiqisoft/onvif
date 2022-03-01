package wsdiscovery

/*******************************************************
 * Copyright (C) 2018 Palanjyan Zhorzhik
 *
 * This file is part of ws-discovery project.
 *
 * ws-discovery can be copied and/or distributed without the express
 * permission of Palanjyan Zhorzhik
 *******************************************************/

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/IOTechSystems/onvif"

	"github.com/beevik/etree"
	"github.com/gofrs/uuid"
	"golang.org/x/net/ipv4"
)

const (
	bufSize              = 8192
	defaultInterfaceName = "en0"
)

//GetAvailableDevicesAtSpecificEthernetInterface ...
func GetAvailableDevicesAtSpecificEthernetInterface(interfaceName string) []onvif.Device {
	/*
		Call an ws-discovery Probe Message to Discover NVT type Devices
	*/
	//var scopes = []string{""}
	//var types = []string{""}
	if strings.TrimSpace(interfaceName) == "" {
		interfaceName = defaultInterfaceName
	}
	probeResponses := SendProbe(interfaceName, nil, nil, map[string]string{"dn": "http://www.onvif.org/ver10/network/wsdl"})
	nvtDevices := make([]onvif.Device, 0)
	nvtDevices, err := devicesFromProbeResponses(probeResponses)
	if err != nil {
		fmt.Printf("Fail to discover Onvif camera: %s \n", err)
	}

	return nvtDevices
}

func devicesFromProbeResponses(probeResponses []string) ([]onvif.Device, error) {
	nvtDevices := make([]onvif.Device, 0)
	xaddrSet := make(map[string]struct{})
	for _, j := range probeResponses {
		doc := etree.NewDocument()
		if err := doc.ReadFromString(j); err != nil {
			return nil, err
		}

		probeMatches := doc.Root().FindElements("./Body/ProbeMatches/ProbeMatch")
		for _, probeMatch := range probeMatches {
			var xaddr string
			if address := probeMatch.FindElement("./XAddrs"); address != nil {
				u, err := url.Parse(address.Text())
				if err != nil {
					fmt.Printf("Invalid XAddrs: %s\n", address.Text())
					continue
				}
				xaddr = u.Host
			}
			if _, dupe := xaddrSet[xaddr]; dupe {
				fmt.Printf("Skipping duplicate XAddr: %s\n", xaddr)
				continue
			}

			var endpointRefAddress string
			if ref := probeMatch.FindElement("./EndpointReference/Address"); ref != nil {
				uuidElements := strings.Split(ref.Text(), ":")
				endpointRefAddress = uuidElements[len(uuidElements)-1]
			}

			dev, err := onvif.NewDevice(onvif.DeviceParams{
				Xaddr:              xaddr,
				EndpointRefAddress: endpointRefAddress,
			})
			if err != nil {
				fmt.Printf("Failed to connect to camera at %s\n", xaddr)
				continue
			}
			xaddrSet[xaddr] = struct{}{}
			nvtDevices = append(nvtDevices, *dev)
			fmt.Printf("Onvif WS-Discovery: Find Xaddr: %-25s EndpointRefAddress: %s\n", xaddr, string(endpointRefAddress))
		}
	}

	return nvtDevices, nil
}

//SendProbe to device
func SendProbe(interfaceName string, scopes, types []string, namespaces map[string]string) []string {
	// Creating UUID Version 4
	uuidV4 := uuid.Must(uuid.NewV4())
	//fmt.Printf("UUIDv4: %s\n", uuidV4)

	probeSOAP := buildProbeMessage(uuidV4.String(), scopes, types, namespaces)
	//probeSOAP = `<?xml version="1.0" encoding="UTF-8"?>
	//<Envelope xmlns="http://www.w3.org/2003/05/soap-envelope" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing">
	//<Header>
	//<a:Action mustUnderstand="1">http://schemas.xmlsoap.org/ws/2005/04/discovery/Probe</a:Action>
	//<a:MessageID>uuid:78a2ed98-bc1f-4b08-9668-094fcba81e35</a:MessageID><a:ReplyTo>
	//<a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address>
	//</a:ReplyTo><a:To mustUnderstand="1">urn:schemas-xmlsoap-org:ws:2005:04:discovery</a:To>
	//</Header>
	//<Body><Probe xmlns="http://schemas.xmlsoap.org/ws/2005/04/discovery">
	//<d:Types xmlns:d="http://schemas.xmlsoap.org/ws/2005/04/discovery" xmlns:dp0="http://www.onvif.org/ver10/network/wsdl">dp0:NetworkVideoTransmitter</d:Types>
	//</Probe>
	//</Body>
	//</Envelope>`

	return sendUDPMulticast(probeSOAP.String(), interfaceName)

}

func sendUDPMulticast(msg string, interfaceName string) []string {
	var result []string
	data := []byte(msg)
	iface, err := net.InterfaceByName(interfaceName)
	if err != nil {
		fmt.Println(err)
	}
	group := net.IPv4(239, 255, 255, 250)

	c, err := net.ListenPacket("udp4", "0.0.0.0:1024")
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()

	p := ipv4.NewPacketConn(c)
	if err := p.JoinGroup(iface, &net.UDPAddr{IP: group}); err != nil {
		fmt.Println(err)
	}

	dst := &net.UDPAddr{IP: group, Port: 3702}
	for _, ifi := range []*net.Interface{iface} {
		if err := p.SetMulticastInterface(ifi); err != nil {
			fmt.Println(err)
		}
		p.SetMulticastTTL(2)
		if _, err := p.WriteTo(data, nil, dst); err != nil {
			fmt.Println(err)
		}
	}

	if err := p.SetReadDeadline(time.Now().Add(time.Second * 1)); err != nil {
		log.Fatal(err)
	}

	for {
		b := make([]byte, bufSize)
		n, _, _, err := p.ReadFrom(b)
		if err != nil {
			if !errors.Is(err, os.ErrDeadlineExceeded) {
				fmt.Println(err)
			}
			break
		}
		result = append(result, string(b[0:n]))
	}
	return result
}
