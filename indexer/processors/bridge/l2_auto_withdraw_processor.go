package bridge

import (
	"github.com/ethereum-optimism/optimism/indexer/config"
	"github.com/ethereum-optimism/optimism/indexer/database"
	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
)

func L2ProcessAutoWithdrawEvents(db *database.DB, l2Contracts config.L2Contracts) error {
	L2AutoWithdrawBridgeAbi, _ := bindings.L2AutoWithdrawBridgeMetaData.GetAbi()
	withdrawals, err := db.BridgeTransactions.L2TransactionWithdrawalsNullIsAutoWithdrawal()
	if err != nil {
		return err
	} else if len(withdrawals) == 0 {
		return nil
	}

	nonAutoWithdrawals := make([]database.L2TransactionWithdrawal, 0)
	autoWithdrawals := make([]database.L2TransactionWithdrawal, 0)
	false_ := false
	true_ := true
	for _, withdrawal := range withdrawals {
		l2ContractEvent, err := db.ContractEvents.L2ContractEvent(withdrawal.InitiatedL2EventGUID)
		if err != nil {
			return err
		}

		// TODO add come comment for `3`
		if l2ContractEvent.LogIndex < 3 {
			withdrawal.IsAutoWithdrawal = &false_
			nonAutoWithdrawals = append(nonAutoWithdrawals, withdrawal)
			continue
		}

		if autoWithdrawalEvent, err := db.ContractEvents.L2ContractEventWithFilter(database.ContractEvent{
			BlockHash:       l2ContractEvent.BlockHash,
			ContractAddress: l2Contracts.L2AutoWithdrawBridge,
			LogIndex:        l2ContractEvent.LogIndex - 3,
			EventSignature:  L2AutoWithdrawBridgeAbi.Events["AutoWithdrawTo"].ID,
		}); err != nil {
			return err
		} else if autoWithdrawalEvent == nil || autoWithdrawalEvent.LogIndex != l2ContractEvent.LogIndex-3 {
			withdrawal.IsAutoWithdrawal = &false_
			nonAutoWithdrawals = append(nonAutoWithdrawals, withdrawal)
		} else {
			withdrawal.IsAutoWithdrawal = &true_
			autoWithdrawals = append(autoWithdrawals, withdrawal)
		}
	}

	err = db.BridgeTransactions.UpdateL2TransactionWithdrawals(nonAutoWithdrawals)
	if err != nil {
		return err
	}
	err = db.BridgeTransactions.UpdateL2TransactionWithdrawals(autoWithdrawals)
	if err != nil {
		return err
	}

	return nil
}
