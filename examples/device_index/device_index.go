package main

import (
	"flag"
	"io/ioutil"
	"github.com/golang/glog"
	"github.com/vishvananda/netlink"
)

const (
	netDirectory = "/sys/class/net/"
)

func main() {
	ifNamePtr := flag.String("interface", "", "get a kernel interface index from name")
	flag.Parse()
	defer glog.Flush()
	glog.Infof("Starting device index example...")

	ifName := *ifNamePtr

	if ifName != "" {
		glog.Info("Input ifName: ", ifName)
		ifObj, err := netlink.LinkByName(ifName)
		if err != nil {
			glog.Errorf("failed to lookup device %q: %v", ifName, err)
			return
		}
		glog.Info("netlink.LinkByName ifObj: ", ifObj)
		ifIndex := ifObj.Attrs().Index
		glog.Info("ifObj.Attrs().Index: ", ifIndex)
		return
	}

	netDevices, err := ioutil.ReadDir(netDirectory)
	if err != nil {
		glog.Errorf("Error. Cannot read %s for network device names. Err: %v", netDirectory, err)
		return
	}

	if len(netDevices) < 1 {
		glog.Errorf("Error. No network device found in %s directory", netDirectory)
		return
	}

	for _, dev := range netDevices {
		glog.Info("dev name: ", dev.Name())
		devObj, err := netlink.LinkByName(dev.Name())
		if err != nil {
			glog.Errorf("failed to lookup device %q: %v", dev.Name(), err)
			return
		}
		glog.Info("netlink.LinkByName devObj: ", devObj)

		index := devObj.Attrs().Index
		glog.Info("devObj.Attrs().Index: ", index)
	}
}
