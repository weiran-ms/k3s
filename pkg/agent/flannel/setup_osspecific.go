//go:build !windows
// +build !windows

package flannel

import (
    "github.com/k3s-io/k3s/pkg/daemons/config"
)

const (
	cniConf = `{
  "name":"cbr0",
  "cniVersion":"1.0.0",
  "plugins":[
    {
      "type":"flannel",
      "delegate":{
        "hairpinMode":true,
        "forceAddress":true,
        "isDefaultGateway":true
      }
    },
    {
      "type":"portmap",
      "capabilities":{
        "portMappings":true
      }
    },
    {
      "type":"bandwidth",
      "capabilities":{
        "bandwidth":true
      }
    }
  ]
}
`
	cniConfFileName = "10-flannel.conflist"
)

func getCniConf(_ *config.Node) (string, error) {
	return cniConf, nil
}