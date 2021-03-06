package options

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type ClusterDescribeConfig struct {
	Clusters []string
}

func NewClusterDescribeConfig() *ClusterDescribeConfig {
	return &ClusterDescribeConfig{}
}

func (c *ClusterDescribeConfig) AddFlags(fs *pflag.FlagSet) {
}

func (c *ClusterDescribeConfig) ValidateFlags(cmd *cobra.Command, args []string) error {
	c.Clusters = func(names []string) []string {
		for i := range names {
			names[i] = strings.ToLower(names[i])
		}
		return names
	}(args)
	return nil
}
