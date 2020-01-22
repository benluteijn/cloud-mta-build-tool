package commands

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/SAP/cloud-mta-build-tool/internal/artifacts"
)

// mtad gen - command flags
var mtadGenCmdSrc string
var mtadGenCmdTrg string
var mtadGenCmdExtensions []string
var mtadGenCmdPlatform string

// Provide mtad.yaml from mta.yaml
var mtadCmd = &cobra.Command{
	Use:           "mtad",
	Short:         "Generates MTAD",
	Long:          "Generates deployment descriptor (mtad.yaml) from development descriptor (mta.yaml)",
	Args:          cobra.NoArgs,
	RunE:          nil,
	SilenceUsage:  true,
	SilenceErrors: true,
}

var mtadGenCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generates MTAD",
	Long:  "Generates deployment descriptor (mtad.yaml) from development descriptor (mta.yaml)",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := artifacts.ExecuteMtadGen(mtadGenCmdSrc, mtadGenCmdTrg, mtadGenCmdExtensions, mtadGenCmdPlatform, os.Getwd)
		logError(err)
		return err
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {

	mtadCmd.AddCommand(mtadGenCmd)

	// set flags of mtad gen command
	mtadGenCmd.Flags().StringVarP(&mtadGenCmdSrc, "source", "s", "",
		"The path to the MTA project; the current path is set as default")
	mtadGenCmd.Flags().StringVarP(&mtadGenCmdTrg, "target", "t",
		"", "The path to the folder in which the 'mtad.yaml' file is generated; the current path is set as default")
	mtadGenCmd.Flags().StringSliceVarP(&mtadGenCmdExtensions, "extensions", "e", nil,
		"The MTA extension descriptors")
	mtadGenCmd.Flags().StringVarP(&mtadGenCmdPlatform, "platform", "p", "cf",
		`The deployment platform; supported platforms: "cf", "xsa", "neo"`)
	mtadGenCmd.Flags().BoolP("help", "h", false, `Displays detailed information about the "mtad" command`)
}
