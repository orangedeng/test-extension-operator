//go:generate go run internal/codegen/cleanup/main.go
//go:generate /bin/rm -rf pkg/generated
//go:generate go run internal/codegen/main.go

package main

import (
	"fmt"
	"os"

	"github.com/orangedeng/test-extension-operator/internal/controllers"
	"github.com/orangedeng/test-extension-operator/internal/version"

	"github.com/rancher/wrangler/v2/pkg/kubeconfig"
	"github.com/rancher/wrangler/v2/pkg/ratelimit"
	"github.com/rancher/wrangler/v2/pkg/signals"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	LogFormat = "2006/01/02 15:04:05"
)

var (
	kubeConfig     string
	chartNamespace string
	debug          bool
	trace          bool
	ofJSON         bool
	binName        = "operator"
	rootCmd        = &cobra.Command{
		Use:   binName,
		Short: "this is a example rancher-style operator",
		Run: func(cmd *cobra.Command, args []string) {
			if err := run(cmd, args); err != nil {
				logrus.Fatal(err)
			}
		},
	}
	versionCmd = &cobra.Command{
		Use:   "version [--json]",
		Args:  cobra.MaximumNArgs(1),
		Short: "shows the version info",
		Run: func(cmd *cobra.Command, args []string) {
			info := version.GetInfo()
			if ofJSON {
				fmt.Printf("%s", info.JSON())
			} else {
				fmt.Printf("%s", info)
			}
		},
	}
)

func init() {
	versionCmd.Flags().BoolVar(&ofJSON, "json", false, `"json" or empty`)
	rootCmd.Flags().StringVarP(&kubeConfig, "kubeconfig", "f", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	rootCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Enable debug logging.")
	rootCmd.Flags().BoolVar(&trace, "trace", false, "Enable trace logging.")
	rootCmd.MarkFlagsMutuallyExclusive("debug", "trace")
	rootCmd.AddCommand(versionCmd)

	chartNamespace = os.Getenv("CHART_NAMESPACE")

	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, ForceColors: true, TimestampFormat: LogFormat})
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run(_ *cobra.Command, _ []string) error {
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debugf("Loglevel set to [%v]", logrus.DebugLevel)
	}
	if trace {
		logrus.SetLevel(logrus.TraceLevel)
		logrus.Tracef("Loglevel set to [%v]", logrus.TraceLevel)
	}

	logrus.Infof("Starting test-extension-operator controller version %s", version.Short())
	ctx := signals.SetupSignalContext()
	restKubeConfig, err := kubeconfig.GetNonInteractiveClientConfig(kubeConfig).ClientConfig()
	if err != nil {
		return fmt.Errorf("failed to find kubeconfig: %v", err)
	}
	restKubeConfig.RateLimiter = ratelimit.None
	if err := controllers.Start(ctx, restKubeConfig); err != nil {
		return err
	}

	return nil
}
