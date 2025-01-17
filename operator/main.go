// Copyright 2021 Authors of KubeArmor
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/kubearmor/KVMService/operator/core"
	kg "github.com/kubearmor/KVMService/operator/log"
)

func main() {
	if os.Geteuid() != 0 {
		kg.Printf("Need to have root privileges to run %s\n", os.Args[0])
		return
	}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		kg.Err(err.Error())
		return
	}

	if err := os.Chdir(dir); err != nil {
		kg.Err(err.Error())
		return
	}

    portPtr := flag.Int("port", 40400, "Cluster Port")
    ipAddressPtr := flag.String("ipAddress", "", "Cluster Address")

	flag.Parse()

	core.KVMSOperatorDaemon(*portPtr, *ipAddressPtr)
}
