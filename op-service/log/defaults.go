package log

import (
	"os"

	"github.com/ethereum-optimism/optimism/logutil/log"
)

func SetupDefaults() {
	log.Root().SetHandler(
		log.LvlFilterHandler(
			log.LvlInfo,
			log.StreamHandler(os.Stdout, log.LogfmtFormat()),
		),
	)

}
