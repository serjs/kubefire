package cluster

import (
	intcmd "github.com/innobead/kubefire/internal/cmd"
	"github.com/innobead/kubefire/internal/di"
	"github.com/innobead/kubefire/internal/validate"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show [name]",
	Short: "Show cluster info",
	Args:  validate.OneArg("cluster name"),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return validate.CheckClusterExist(args[0])
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		cluster, err := di.ClusterManager().Get(name)
		if err != nil {
			return errors.WithMessagef(err, "failed to get cluster (%s) info", name)
		}

		if err := di.Output().Print(cluster, nil, ""); err != nil {
			return errors.WithMessagef(err, "failed to print output of cluster (%s)", name)
		}

		return nil
	},
}

func init() {
	intcmd.AddOutputFlag(showCmd)
}
