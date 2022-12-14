package validate

import (
	"github.com/0xPolygon/polygon-edge/command"
	"github.com/0xPolygon/polygon-edge/command/helper"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	secretsValidateCmd := &cobra.Command{
		Use:     "validate",
		Short:   "Use the managed private key to sign the P2P address of the remote node and generate signature information",
		PreRunE: runPreRun,
		Run:     runCommand,
	}

	helper.RegisterGRPCAddressFlag(secretsValidateCmd)
	setFlags(secretsValidateCmd)

	return secretsValidateCmd
}

func setFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(
		&params.dataDir,
		dataDirFlag,
		"",
		"the directory for the Polygon Edge data if the local FS is used",
	)

	cmd.Flags().StringVar(
		&params.configPath,
		configFlag,
		"",
		"the path to the SecretsManager config file, "+
			"if omitted, the local FS secrets manager is used",
	)

	cmd.MarkFlagsMutuallyExclusive(dataDirFlag, configFlag)
}

func runPreRun(_ *cobra.Command, _ []string) error {
	return params.validateFlags()
}

func runCommand(cmd *cobra.Command, _ []string) {
	outputter := command.InitializeOutputter(cmd)
	defer outputter.WriteOutput()

	statusResponse, err := getSystemStatus(helper.GetGRPCAddress(cmd))

	if err != nil {
		outputter.SetError(err)
		return
	}

	if err := params.initSecrets(); err != nil {
		outputter.SetError(err)

		return
	}

	if err := params.validate(statusResponse.P2PAddr); err != nil {
		outputter.SetError(err)

		return
	}

	outputter.SetCommandResult(params.getResult())
}
