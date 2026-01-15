package app

import (
	"flag"
	"fmt"
	"os"

	"github.com/kvaster/topols"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var config struct {
	csiSocket            string
	metricsAddr          string
	secureMetricsServer  bool
	poolPath             string
	zapOpts              zap.Options
	profilingBindAddress string
}

var rootCmd = &cobra.Command{
	Use:     "topols-node",
	Version: topols.Version,
	Short:   "TopoLS CSI node",
	Long: `topols-node provides CSI node service.
It also works as a custom Kubernetes controller.

The node name where this program runs must be given by either
NODE_NAME environment variable or --nodename flag.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		return subMain()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//nolint:lll
func init() {
	fs := rootCmd.Flags()
	fs.StringVar(&config.csiSocket, "csi-socket", topols.DefaultCSISocket, "UNIX domain socket filename for CSI")
	fs.StringVar(&config.metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	fs.BoolVar(&config.secureMetricsServer, "secure-metrics-server", false, "Secures the metrics server")
	fs.StringVar(&config.poolPath, "pool-path", "/mnt/pool", "Path to folder with config and mounted btrfs file systems")
	fs.String("nodename", "", "The resource name of the running node")
	fs.StringVar(&config.profilingBindAddress, "profiling-bind-address", "", "Bind pprof profiling to the given network address. If empty, profiling is disabled.")

	_ = viper.BindEnv("nodename", "NODE_NAME")
	_ = viper.BindPFlag("nodename", fs.Lookup("nodename"))

	klogFlags := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(klogFlags)
	fs.AddGoFlagSet(klogFlags)

	zapFlags := flag.NewFlagSet("zap", flag.ExitOnError)
	config.zapOpts.BindFlags(zapFlags)
	fs.AddGoFlagSet(zapFlags)
}
