package validate

import (
	"errors"
	"github.com/0xPolygon/polygon-edge/command"
	"github.com/0xPolygon/polygon-edge/secrets"
	"github.com/0xPolygon/polygon-edge/secrets/helper"
	"github.com/0xPolygon/polygon-edge/secrets/local"
	"github.com/hashicorp/go-hclog"
	libp2pCrypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
)

const (
	dataDirFlag = "data-dir"
	configFlag  = "config"
)

var (
	params = &validateParams{}
)

var (
	errInvalidConfig   = errors.New("invalid secrets configuration")
	errInvalidParams   = errors.New("no config file or data directory passed in")
	errUnsupportedType = errors.New("unsupported secrets manager")
)

type validateParams struct {
	dataDir    string
	configPath string

	secretsManager secrets.SecretsManager
	secretsConfig  *secrets.SecretsManagerConfig

	networkingPrivateKey libp2pCrypto.PrivKey

	nodeID  peer.ID
	address string

	validateData string
}

func (ip *validateParams) validateFlags() error {
	if ip.dataDir == "" && ip.configPath == "" {
		return errInvalidParams
	}

	return nil
}
func (ip *validateParams) validate(addr string) error {
	ip.address = addr
	vd := &ValidateInfo{
		Address: addr,
	}
	err := vd.Sign(ip.networkingPrivateKey)
	if err != nil {
		return err
	}

	ip.validateData = vd.Mac
	return nil
}

func (ip *validateParams) initSecrets() error {
	if err := ip.initSecretsManager(); err != nil {
		return err
	}

	return ip.getNetworkingKey()
}

func (ip *validateParams) initSecretsManager() error {
	if ip.hasConfigPath() {
		return ip.initFromConfig()
	}

	return ip.initLocalSecretsManager()
}

func (ip *validateParams) hasConfigPath() bool {
	return ip.configPath != ""
}

func (ip *validateParams) initFromConfig() error {
	if err := ip.parseConfig(); err != nil {
		return err
	}

	var secretsManager secrets.SecretsManager

	switch ip.secretsConfig.Type {
	case secrets.HashicorpVault:
		vault, err := helper.SetupHashicorpVault(ip.secretsConfig)
		if err != nil {
			return err
		}

		secretsManager = vault
	case secrets.AWSSSM:
		AWSSSM, err := helper.SetupAWSSSM(ip.secretsConfig)
		if err != nil {
			return err
		}

		secretsManager = AWSSSM
	case secrets.GCPSSM:
		GCPSSM, err := helper.SetupGCPSSM(ip.secretsConfig)
		if err != nil {
			return err
		}

		secretsManager = GCPSSM
	default:
		return errUnsupportedType
	}

	ip.secretsManager = secretsManager

	return nil
}

func (ip *validateParams) parseConfig() error {
	secretsConfig, readErr := secrets.ReadConfig(ip.configPath)
	if readErr != nil {
		return errInvalidConfig
	}

	if !secrets.SupportedServiceManager(secretsConfig.Type) {
		return errUnsupportedType
	}

	ip.secretsConfig = secretsConfig

	return nil
}

func (ip *validateParams) initLocalSecretsManager() error {
	local, err := local.SecretsManagerFactory(
		nil, // Local secrets manager doesn't require a config
		&secrets.SecretsManagerParams{
			Logger: hclog.NewNullLogger(),
			Extra: map[string]interface{}{
				secrets.Path: ip.dataDir,
			},
		},
	)
	if err != nil {
		return err
	}

	ip.secretsManager = local

	return nil
}

func (ip *validateParams) getNetworkingKey() error {
	networkingKey, err := helper.GetNetworkingPrivateKey(ip.secretsManager)
	if err != nil {
		return err
	}

	ip.networkingPrivateKey = networkingKey

	return ip.initNodeID()
}

func (ip *validateParams) initNodeID() error {
	nodeID, err := peer.IDFromPrivateKey(ip.networkingPrivateKey)
	if err != nil {
		return err
	}

	ip.nodeID = nodeID

	return nil
}

func (ip *validateParams) getResult() command.CommandResult {
	return &SecretsValidateResult{
		ValidateInfo: ip.validateData,
		NodeID:       ip.nodeID.String(),
		Address:      ip.address,
	}
}
