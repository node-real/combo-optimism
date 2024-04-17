package log

import (
	"os"

	log2 "github.com/ethereum-optimism/optimism/op-node/logutil/log"
)

func SetupDefaults() {
	log2.Root().SetHandler(
		log2.LvlFilterHandler(
			log2.LvlInfo,
			log2.StreamHandler(os.Stdout, log2.LogfmtFormat()),
		),
	)

}
