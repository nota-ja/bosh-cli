package manifest

import (
	bmkeystr "github.com/cloudfoundry/bosh-micro-cli/keystringifier"
)

type NetworkType string

func (n NetworkType) String() string {
	return string(n)
}

const (
	Dynamic NetworkType = "dynamic"
	Manual  NetworkType = "manual"
	VIP     NetworkType = "vip"
)

type Network struct {
	Name               string                      `yaml:"name"`
	Type               NetworkType                 `yaml:"type"`
	RawCloudProperties map[interface{}]interface{} `yaml:"cloud_properties"`
	IP                 string                      `yaml:"ip"`
	Netmask            string                      `yaml:"netmask"`
	Gateway            string                      `yaml:"gateway"`
	DNS                []string                    `yaml:"dns"`
}

func (n Network) CloudProperties() (map[string]interface{}, error) {
	return bmkeystr.NewKeyStringifier().ConvertMap(n.RawCloudProperties)
}

func (n Network) Spec() (map[string]interface{}, error) {
	cloudProperties, err := n.CloudProperties()
	if err != nil {
		return map[string]interface{}{}, err
	}

	spec := map[string]interface{}{
		"type":             n.Type.String(),
		"ip":               n.IP,
		"cloud_properties": cloudProperties,
	}

	if n.Netmask != "" {
		spec["netmask"] = n.Netmask
	}

	if n.Gateway != "" {
		spec["gateway"] = n.Gateway
	}

	if len(n.DNS) > 0 {
		spec["dns"] = n.DNS
	}

	return spec, nil
}