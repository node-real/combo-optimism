package chaincfg

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum-optimism/superchain-registry/superchain"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/eth"
)

var Mainnet, Sepolia *rollup.Config

func init() {
	mustCfg := func(name string) *rollup.Config {
		cfg, err := GetRollupConfig(name)
		if err != nil {
			panic(fmt.Errorf("failed to load rollup config %q: %w", name, err))
		}
		return cfg
	}
	Mainnet = mustCfg("op-mainnet")
	Sepolia = mustCfg("op-sepolia")
}

var L2ChainIDToNetworkDisplayName = func() map[string]string {
	out := make(map[string]string)
	for _, netCfg := range superchain.OPChains {
		out[fmt.Sprintf("%d", netCfg.ChainID)] = netCfg.Name
	}
	return out
}()

// AvailableNetworks returns the selection of network configurations that is available by default.
func AvailableNetworks() []string {
	var networks []string
	for _, cfg := range superchain.OPChains {
		networks = append(networks, cfg.Chain+"-"+cfg.Superchain)
	}
	return networks
}

func handleLegacyName(name string) string {
	switch name {
	case "mainnet":
		return "op-mainnet"
	case "sepolia":
		return "op-sepolia"
	default:
		return name
	}
}

// ChainByName returns a chain, from known available configurations, by name.
// ChainByName returns nil when the chain name is unknown.
func ChainByName(name string) *superchain.ChainConfig {
	// Handle legacy name aliases
	name = handleLegacyName(name)
	for _, chainCfg := range superchain.OPChains {
		if strings.EqualFold(chainCfg.Chain+"-"+chainCfg.Superchain, name) {
			return chainCfg
		}
	}
	return nil
}

func GetRollupConfig(name string) (*rollup.Config, error) {
	chainCfg := ChainByName(name)
	if chainCfg == nil {
		return nil, fmt.Errorf("invalid network: %q", name)
	}
	rollupCfg, err := rollup.LoadOPStackRollupConfig(chainCfg.ChainID)
	if err != nil {
		return nil, fmt.Errorf("failed to load rollup config: %w", err)
	}
	return rollupCfg, nil
}

var NetworksByName = map[string]rollup.Config{
	"opBNBMainnet": OPBNBMainnet,
	"opBNBTestnet": OPBNBTestnet,
	"opBNBQANet":   OPBNBQANet,
	"comboMainnet": ComboMainnet,
	"comboTestnet": ComboTestnet,
}

var NetworksByChainId = map[string]rollup.Config{
	"204":  OPBNBMainnet,
	"5611": OPBNBTestnet,
	"2484": OPBNBQANet,
	"9980": ComboMainnet,
	"1715": ComboTestnet,
}

func GetRollupConfigByNetwork(name string) (rollup.Config, error) {
	network, ok := NetworksByName[name]
	if !ok {
		return rollup.Config{}, fmt.Errorf("invalid network %s", name)
	}

	return network, nil
}

func GetRollupConfigByChainId(chainId string) (rollup.Config, error) {
	network, ok := NetworksByChainId[chainId]
	if !ok {
		return rollup.Config{}, fmt.Errorf("no match pre-setting network chainId %s, use file config", chainId)
	}

	return network, nil
}

var OPBNBMainnet = rollup.Config{
	Genesis: rollup.Genesis{
		L1: eth.BlockID{
			Hash:   common.HexToHash("0x29443b21507894febe7700f7c5cd3569cc8bf1ba535df0489276d8004af81044"),
			Number: 30758357,
		},
		L2: eth.BlockID{
			Hash:   common.HexToHash("0x4dd61178c8b0f01670c231597e7bcb368e84545acd46d940a896d6a791dd6df4"),
			Number: 0,
		},
		L2Time: 1691753723,
		SystemConfig: eth.SystemConfig{
			BatcherAddr: common.HexToAddress("0xef8783382ef80ec23b66c43575a6103deca909c3"),
			Overhead:    eth.Bytes32(common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000834")),
			Scalar:      eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000f4240")),
			GasLimit:    100000000,
		},
	},
	BlockTime:              1,
	MaxSequencerDrift:      600,
	SeqWindowSize:          14400,
	ChannelTimeout:         1200,
	L1ChainID:              big.NewInt(56),
	L2ChainID:              big.NewInt(204),
	BatchInboxAddress:      common.HexToAddress("0xff00000000000000000000000000000000000204"),
	DepositContractAddress: common.HexToAddress("0x1876ea7702c0ad0c6a2ae6036de7733edfbca519"),
	L1SystemConfigAddress:  common.HexToAddress("0x7ac836148c14c74086d57f7828f2d065672db3b8"),
	RegolithTime:           u64Ptr(0),
	Fermat:                 big.NewInt(9397477), // Nov-28-2023 06 AM +UTC
	SnowTime:               u64Ptr(1713160800),  // Apr-15-2024 06 AM +UTC
	CanyonTime:             u64Ptr(1718870400),  // Jun-20-2024 08:00 AM +UTC
	DeltaTime:              u64Ptr(1718871000),  // Jun-20-2024 08:10 AM +UTC
	EcotoneTime:            u64Ptr(1718871600),  // Jun-20-2024 08:20 AM +UTC
	FjordTime:              u64Ptr(1727157600),  // Sep-24-2024 06:00 AM +UTC
}

var OPBNBTestnet = rollup.Config{
	Genesis: rollup.Genesis{
		L1: eth.BlockID{
			Hash:   common.HexToHash("0xc01a09840419cd993cf4666309f36e6d38de39771af8dbffecfa0386321c19f7"),
			Number: 30727847,
		},
		L2: eth.BlockID{
			Hash:   common.HexToHash("0x51fa57729dfb1c27542c21b06cb72a0459c57440ceb43a465dae1307cd04fe80"),
			Number: 0,
		},
		L2Time: 1686878506,
		SystemConfig: eth.SystemConfig{
			BatcherAddr: common.HexToAddress("0x1fd6a75cc72f39147756a663f3ef1fc95ef89495"),
			Overhead:    eth.Bytes32(common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000834")),
			Scalar:      eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000f4240")),
			GasLimit:    100000000,
		},
	},
	BlockTime:              1,
	MaxSequencerDrift:      600,
	SeqWindowSize:          14400,
	ChannelTimeout:         1200,
	L1ChainID:              big.NewInt(97),
	L2ChainID:              big.NewInt(5611),
	BatchInboxAddress:      common.HexToAddress("0xff00000000000000000000000000000000005611"),
	DepositContractAddress: common.HexToAddress("0x4386c8abf2009ac0c263462da568dd9d46e52a31"),
	L1SystemConfigAddress:  common.HexToAddress("0x406ac857817708eaf4ca3a82317ef4ae3d1ea23b"),
	RegolithTime:           u64Ptr(0),
	Fermat:                 big.NewInt(12113000), // Nov-03-2023 06 AM +UTC
	SnowTime:               u64Ptr(1715752800),   // May-15-2024 06:00 AM +UTC
	CanyonTime:             u64Ptr(1715753400),   // May-15-2024 06:10 AM +UTC
	DeltaTime:              u64Ptr(1715754000),   // May-15-2024 06:20 AM +UTC
	EcotoneTime:            u64Ptr(1715754600),   // May-15-2024 06:30 AM +UTC
	FjordTime:              u64Ptr(1725948000),   // Sep-10-2024 06:00 AM +UTC
}

var OPBNBQANet = rollup.Config{
	Genesis: rollup.Genesis{
		L1: eth.BlockID{
			Hash:   common.HexToHash("0xdbbbe8b752ef975c4a0592472de646bc683b66c824dfedf5d12ecdcc97a5d0c9"),
			Number: 3311074,
		},
		L2: eth.BlockID{
			Hash:   common.HexToHash("0x73eaf214333f29eed23c4902fdc17889b3e379372e52a42567d0069e1b10cdb0"),
			Number: 0,
		},
		L2Time: 1723613564,
		SystemConfig: eth.SystemConfig{
			BatcherAddr: common.HexToAddress("0xb3ad01bd1183bb8537f3e48c42889d828a89b55f"),
			Overhead:    eth.Bytes32(common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000834")),
			Scalar:      eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000f4240")),
			GasLimit:    100000000,
		},
	},
	BlockTime:              1,
	MaxSequencerDrift:      600,
	SeqWindowSize:          14400,
	ChannelTimeout:         1200,
	L1ChainID:              big.NewInt(714),
	L2ChainID:              big.NewInt(1081),
	BatchInboxAddress:      common.HexToAddress("0xff00000000000000000000000000000000001081"),
	DepositContractAddress: common.HexToAddress("0xbf33e25ac03e99dcbc63998471527f23dfbf811f"),
	L1SystemConfigAddress:  common.HexToAddress("0x644daa12057118ce60d25a9ba707f571658911ae"),
	RegolithTime:           u64Ptr(0),
	Fermat:                 big.NewInt(0),
	SnowTime:               u64Ptr(0),
	CanyonTime:             u64Ptr(0),
	DeltaTime:              u64Ptr(0),
	EcotoneTime:            u64Ptr(0),
	FjordTime:              u64Ptr(1724392800), // AUG-23-2024 06:00 AM +UTC
}

var ComboMainnet = rollup.Config{
	Genesis: rollup.Genesis{
		L1: eth.BlockID{
			Hash:   common.HexToHash("0x7743484e78e047654f5a92c5d66a25828f1c259b2c9a780a39936c001fdcbcf7"),
			Number: 33768568,
		},
		L2: eth.BlockID{
			Hash:   common.HexToHash("0x92fcf9e91a4cdd7ffc7e67207e77dfba049bacf1ede5c5917a40f9537e05f4bc"),
			Number: 0,
		},
		L2Time: 1700817067,
		SystemConfig: eth.SystemConfig{
			BatcherAddr: common.HexToAddress("0x6df30535bbe94a533d9f1600e69a642abb3e063f"),
			Overhead:    eth.Bytes32(common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000834")),
			Scalar:      eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000f4240")),
			GasLimit:    100000000,
		},
	},
	BlockTime:              1,
	MaxSequencerDrift:      600,
	SeqWindowSize:          14400,
	ChannelTimeout:         1200,
	L1ChainID:              big.NewInt(56),
	L2ChainID:              big.NewInt(9980),
	BatchInboxAddress:      common.HexToAddress("0xff00000000000000000000000000000000009980"),
	DepositContractAddress: common.HexToAddress("0x419df125e0a712db4b10209ac3055b58b840f1f4"),
	L1SystemConfigAddress:  common.HexToAddress("0x19d9791f6f5df45fb4ced2ea0904e48a6f9e545c"),
	RegolithTime:           u64Ptr(0),
	Fermat:                 big.NewInt(0),
	SnowTime:               u64Ptr(1719813600), // July-01-2024 06:00 AM +UTC
	CanyonTime:             u64Ptr(1719814200), // July-01-2024 06:10 AM +UTC
	DeltaTime:              u64Ptr(1719814800), // July-01-2024 06:20 AM +UTC
	EcotoneTime:            u64Ptr(1719815400), // July-01-2024 06:30 AM +UTC
}

var ComboTestnet = rollup.Config{
	Genesis: rollup.Genesis{
		L1: eth.BlockID{
			Hash:   common.HexToHash("0x5c42e35d2cf086cdbf6ad194701a48688f3c1e8d2eecf57168a248e63a33be86"),
			Number: 39570635,
		},
		L2: eth.BlockID{
			Hash:   common.HexToHash("0x51d565869cb1c6d7e577848ccd1f93918b99427c1cc3fed9dca628d0cffcefbb"),
			Number: 0,
		},
		L2Time: 1713424951,
		SystemConfig: eth.SystemConfig{
			BatcherAddr: common.HexToAddress("0x013fce36695321d32251e58bff33292c685da696"),
			Overhead:    eth.Bytes32(common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000834")),
			Scalar:      eth.Bytes32(common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000f4240")),
			GasLimit:    100000000,
		},
	},
	BlockTime:              1,
	MaxSequencerDrift:      600,
	SeqWindowSize:          14400,
	ChannelTimeout:         1200,
	L1ChainID:              big.NewInt(97),
	L2ChainID:              big.NewInt(1715),
	BatchInboxAddress:      common.HexToAddress("0xff00000000000000000000000000000000001715"),
	DepositContractAddress: common.HexToAddress("0x6357aeea63f9bea16a9264fad26fdfddbec559e6"),
	L1SystemConfigAddress:  common.HexToAddress("0x3a472819b60885f1f11ee173f56e43961460c391"),
	RegolithTime:           u64Ptr(0),
	Fermat:                 big.NewInt(0),
	SnowTime:               u64Ptr(1719208800), // June-24-2024 06:00 AM +UTC
	CanyonTime:             u64Ptr(1719209400), // June-24-2024 06:10 AM +UTC
	DeltaTime:              u64Ptr(1719210000), // June-24-2024 06:20 AM +UTC
	EcotoneTime:            u64Ptr(1719210600), // June-24-2024 06:30 AM +UTC
}

func u64Ptr(v uint64) *uint64 {
	return &v
}
