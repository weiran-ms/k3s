//go:build windows
// +build windows

package flannel

import (
	"fmt"
	"strings"

	"github.com/k3s-io/k3s/pkg/daemons/config"
	"github.com/sirupsen/logrus"
)

const (
	cniConf = `{
	"name": "cbr0",
	"cniVersion": "0.3.0",
	"type": "flannel",
	"capabilities": {
	  "dns": true
	},
	"delegate": {
	  "type": "%DelegateType%",
	  "apiVersion": 2,
	  "hairpinMode": true,
	  "isDefaultGateway": true,
	  "policies": [
	  ]
	}
  }
`
	vxlanDelegateType = "win-overlay"

	hostgwDelegateType = "win-bridge"

	cniConfFileName = "10-flannel.conf"
)

func getCniConf(nodeConfig *config.Node) (string, error) {
	parts := strings.SplitN(nodeConfig.FlannelBackend, "=", 2)
	backend := parts[0]
	if len(parts) > 1 {
		logrus.Fatalf("The additional options through flannel-backend are deprecated and were removed in k3s v1.27, use flannel-conf instead")
	}

	var cniConfig string

	switch backend {
	case config.FlannelBackendVXLAN:
		cniConfig = strings.ReplaceAll(cniConf, "%DelegateType%", vxlanDelegateType)
	case config.FlannelBackendHostGW:
		cniConfig = strings.ReplaceAll(cniConf, "%DelegateType%", hostgwDelegateType)
	default:
		return "", fmt.Errorf("Cannot configure unknown flannel backend '%s'", nodeConfig.FlannelBackend)
	}

	return cniConfig, nil
}
