package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/vishvananda/netlink"
)

func renameLink(curName, newName string) error {
	link, err := netlink.LinkByName(curName)
	if err != nil {
		return fmt.Errorf("failed to lookup device %q: %v", curName, err)
	}

	return netlink.LinkSetName(link, newName)
}

func main() {
	ifNamePtr := flag.String("interface", "eth0", "a kernel interface name to be renamed")
	flag.Parse()
	defer glog.Flush()
	glog.Infof("Starting example...")

	ifName := *ifNamePtr

	glog.Info("Input ifName: %v", ifName)
	vfDev, err := netlink.LinkByName(ifName)
	if err != nil {
		glog.Errorf("failed to lookup vf device %q: %v", ifName, err)
		return
	}
	glog.Info("netlink.LinkByName vfDev: %v", vfDev)

	index := vfDev.Attrs().Index
	glog.Info("vfDev.Attrs().Index: %v", index)

	devName := fmt.Sprintf("dev%d", index)
	glog.Info("fmt.Sprintf devName: %v", devName)

	if err = netlink.LinkSetDown(vfDev); err != nil {
		glog.Errorf("failed to down vf device %q: %v", ifName, err)
		return
	}

	err = renameLink(ifName, devName)
	if err != nil {
		glog.Errorf("failed to rename vf device %q to %q: %v", ifName, devName, err)
		return
	}

	if err = netlink.LinkSetUp(vfDev); err != nil {
		glog.Errorf("failed to up vf device %q: %v", devName, err)
		return
	}

}
