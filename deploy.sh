#!/bin/bash

# L2 chain id, 常用 901
export OP_L2_CHAIN_ID=${OP_L2_CHAIN_ID:-"901"}

# OP_DEPLOYER 是部署账户，需要有余额
export OP_DEPLOYER_ADDRESS=${OP_DEPLOYER_ADDRESS?"required"}
export OP_DEPLOYER_PRIVKEY=${OP_DEPLOYER_PRIVKEY?"required"}

# OP_BATCHER，需要有余额，提交 L2 交易信息到 OP_BATCHER_SEQUENCER_BATCH_INBOX_ADDRESS
export OP_BATCHER_PRIVKEY=${OP_BATCHER_PRIVKEY?"required"}
export OP_BATCHER_ADDRESS=${OP_BATCHER_ADDRESS?"required"}

# OP_PROPOSER，需要有余额，提交 L2 交易结果到 L1 合约
export OP_PROPOSER_PRIVKEY=${OP_PROPOSER_PRIVKEY?"required"}
export OP_PROPOSER_ADDRESS=${OP_PROPOSER_ADDRESS?"required"}

# 随便一个 EOA 地址，参考 OP_BATCHER 的定义
export OP_BATCHER_SEQUENCER_BATCH_INBOX_ADDRESS=${OP_BATCHER_SEQUENCER_BATCH_INBOX_ADDRESS?"required"}

export BSC_TESTNET="https://data-seed-prebsc-1-s1.binance.org:8545"
export BSC_QANET="http://tf-dex-qa-ec2-bsc-test-alb-1168080131.ap-northeast-1.elb.amazonaws.com:8545"
export L1_ENDPOINT="$BSC_TESTNET"

# ethereum-optimism/optimism 仓库的根目录
export OP_ROOT_DIR=${OP_ROOT_DIR?"required"}
export OP_NETWORK_NAME="bsc-testnet"
export OP_DATA_DIR=$OP_ROOT_DIR/.$OP_NETWORK_NAME

clean() {
    docker stop l2 op-node op-batcher op-proposer

    rm -rf $OP_DATA_DIR
    rm -rf $OP_ROOT_DIR/packages/contracts-bedrock/deployments/$OP_NETWORK_NAME
    mkdir  $OP_DATA_DIR
}

l1_chain_id() {
    curl $L1_ENDPOINT -X POST --data '{"jsonrpc":"2.0","method":"eth_chainId","params":[],"id":74}' -H 'Content-Type: application/json' \
        | jq '.result' \
        | xargs -I {} printf '%d' {}
}

l1_tip() {
    l1_tip_number=$(curl $L1_ENDPOINT -X POST --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":74}' -H 'Content-Type: application/json' \
        | jq '.result')
    l1_tip_timestamp=$(curl $L1_ENDPOINT -X POST --data '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":['"$l1_tip_number"', false],"id":74}' -H 'Content-Type: application/json' \
        | jq '.result.timestamp')
    echo "$l1_tip_number" "$l1_tip_timestamp"
}

start_admin() {
    blockhash=$(curl http://127.0.0.1:9546 -X POST --data '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["latest", true],"id":74}' -H 'Content-Type: application/json' | jq '.result.hash')
    result=$(curl http://127.0.0.1:7546 -X POST --data '{"method":"admin_startSequencer","params":['"$blockhash"'],"id":1,"jsonrpc":"2.0"}' -H 'Content-Type: application/json' | jq '.result')
    echo "$blockhash" "$result"
}

generate_deploy_config() {
    l1_chain_id_=$(l1_chain_id)
    l1_tip_number=$(l1_tip | awk '{print $1}')
    l1_tip_timestamp=$(l1_tip | awk '{print $2}')

    DEPLOY_CONFIG_TEMPLATE=$OP_ROOT_DIR/packages/contracts-bedrock/deploy-config/devnetL1.json
    DEPLOY_CONFIG_TARGET=$OP_ROOT_DIR/packages/contracts-bedrock/deploy-config/$OP_NETWORK_NAME.json
    cat $DEPLOY_CONFIG_TEMPLATE \
        | jq ".l2ChainID = $OP_L2_CHAIN_ID" \
        | jq ".l1ChainID = $l1_chain_id_" \
        | jq ".l1StartingBlockTag = $l1_tip_number" \
        | jq ".l1GenesisBlockTimestamp = $l1_tip_timestamp" \
        | jq ".sequencerWindowSize = 14400" \
        | jq ".l2BlockTime = 1" \
        | jq --arg OP_DEPLOYER_ADDRESS "$OP_DEPLOYER_ADDRESS" '.controller = $OP_DEPLOYER_ADDRESS' \
        | jq --arg OP_DEPLOYER_ADDRESS "$OP_DEPLOYER_ADDRESS" '.baseFeeVaultRecipient = $OP_DEPLOYER_ADDRESS' \
        | jq --arg OP_DEPLOYER_ADDRESS "$OP_DEPLOYER_ADDRESS" '.proxyAdminOwner = $OP_DEPLOYER_ADDRESS' \
        | jq --arg OP_DEPLOYER_ADDRESS "$OP_DEPLOYER_ADDRESS" '.finalSystemOwner = $OP_DEPLOYER_ADDRESS' \
        | jq --arg OP_DEPLOYER_ADDRESS "$OP_DEPLOYER_ADDRESS" '.governanceTokenOwner = $OP_DEPLOYER_ADDRESS' \
        | jq --arg OP_BATCHER_ADDRESS  "$OP_BATCHER_ADDRESS"  '.batchSenderAddress = $OP_BATCHER_ADDRESS' \
        | jq --arg OP_PROPOSER_ADDRESS "$OP_PROPOSER_ADDRESS" '.l2OutputOracleProposer = $OP_PROPOSER_ADDRESS' \
        | jq --arg OP_BATCHER_ADDRESS  "$OP_BATCHER_ADDRESS"  '.p2pSequencerAddress = $OP_BATCHER_ADDRESS' \
        | jq --arg OP_BATCHER_SEQUENCER_BATCH_INBOX_ADDRESS "$OP_BATCHER_SEQUENCER_BATCH_INBOX_ADDRESS" '.batchInboxAddress = $OP_BATCHER_SEQUENCER_BATCH_INBOX_ADDRESS' \
        | tee $DEPLOY_CONFIG_TARGET
}

deploy_contracts() {
    cd $OP_ROOT_DIR/packages/contracts-bedrock

    CHAIN_ID=$(l1_chain_id) \
    L1_RPC=$L1_ENDPOINT \
    PRIVATE_KEY_DEPLOYER=$OP_DEPLOYER_PRIVKEY \
    yarn hardhat --network $OP_NETWORK_NAME deploy
}

output_addresses() {
    DEPLOYMENT_DIR=$OP_ROOT_DIR/packages/contracts-bedrock/deployments/$OP_NETWORK_NAME
    ADDRESSES_JSON_PATH=$OP_DATA_DIR/addresses.json
    SDK_ADDRESSES_JSON_PATH=$OP_DATA_DIR/sdk-addresses.json

    CONTRACTS=$(ls -1 $DEPLOYMENT_DIR/*.json)
    addresses="{}"
    for contract_path in $(ls -1 $DEPLOYMENT_DIR/*.json); do
        contract_name=$(echo $contract_path | xargs -I {} basename {} | awk -F '.' '{print $1}')
        contract_addr=$(cat $contract_path | jq '.address')
        addresses=$(echo $addresses | jq ".$contract_name = $contract_addr")
    done
    echo $addresses | jq | tee $ADDRESSES_JSON_PATH

    sdk_addresses="{
        \"AddressManager\":             \"0x0000000000000000000000000000000000000000\",
        \"StateCommitmentChain\":       \"0x0000000000000000000000000000000000000000\",
        \"CanonicalTransactionChain\":  \"0x0000000000000000000000000000000000000000\",
        \"BondManager\":                \"0x0000000000000000000000000000000000000000\",
        \"L1CrossDomainMessenger\":     $(echo $addresses | jq '.Proxy__OVM_L1CrossDomainMessenger'),
        \"L1StandardBridge\":           $(echo $addresses | jq '.Proxy__OVM_L1StandardBridge'),
        \"OptimismPortal\":             $(echo $addresses | jq '.OptimismPortalProxy'),
        \"L2OutputOracle\":             $(echo $addresses | jq '.L2OutputOracleProxy')
    }"
    echo $sdk_addresses | jq | tee $SDK_ADDRESSES_JSON_PATH
}

generate_rollup_config() {
    cd $OP_ROOT_DIR/op-node/

    go run cmd/main.go genesis l2 \
        --l1-rpc $L1_ENDPOINT \
        --deploy-config  $OP_ROOT_DIR/packages/contracts-bedrock/deploy-config/$OP_NETWORK_NAME.json \
        --deployment-dir $OP_ROOT_DIR/packages/contracts-bedrock/deployments/$OP_NETWORK_NAME \
        --outfile.l2     $OP_DATA_DIR/genesis-l2.json \
        --outfile.rollup $OP_DATA_DIR/rollup.json
}

launch_op_geth() {
    docker run \
        --name l2 -d -it --rm \
        -p "9545:8545" \
        -p "8551:8551" \
        -v "$OP_DATA_DIR/db:/db" \
        -v "$OP_DATA_DIR/genesis-l2.json:/genesis.json" \
        -v "$OP_ROOT_DIR/ops-bedrock/test-jwt-secret.txt:/config/test-jwt-secret.txt" \
        l2geth:latest --authrpc.jwtsecret=/config/test-jwt-secret.txt
    sleep 60
}

launch_op_geth_bak() {
    docker run \
        --name l2_bak -it --rm \
        -p "9546:8545" \
        -p "8552:8551" \
        -v "$OP_DATA_DIR/db_bak:/db" \
        -v "$OP_DATA_DIR/genesis-l2.json:/genesis.json" \
        -v "$OP_ROOT_DIR/ops-bedrock/test-jwt-secret.txt:/config/test-jwt-secret.txt" \
        l2geth:latest --authrpc.jwtsecret=/config/test-jwt-secret.txt
    sleep 60
}

launch_op_node() {
    docker run \
        --name op-node -d -it --rm \
        -p "7545:8545" \
        -p "9003:9003" \
        -p "7300:7300" \
        -p "6060:6060" \
        -v "$OP_ROOT_DIR/ops-bedrock/p2p-sequencer-key.txt:/config/p2p-sequencer-key.txt" \
        -v "$OP_ROOT_DIR/ops-bedrock/p2p-node-key.txt:/config/p2p-node-key.txt" \
        -v "$OP_ROOT_DIR/ops-bedrock/test-jwt-secret.txt:/config/test-jwt-secret.txt" \
        -v "$OP_DATA_DIR/rollup.json:/rollup.json" \
        -v "$OP_DATA_DIR/op_log:/op_log" \
        op-node:latest \
        op-node \
            --l1.trustrpc \
            --l1=$L1_ENDPOINT \
            --l2=http://host.docker.internal:8551 \
            --l2.jwt-secret=/config/test-jwt-secret.txt \
            --sequencer.enabled \
            --sequencer.stopped \
            --sequencer.l1-confs=0 \
            --verifier.l1-confs=0 \
            --p2p.sequencer.key="$OP_BATCHER_PRIVKEY" \
            --rollup.config=/rollup.json \
            --rpc.addr=0.0.0.0 \
            --rpc.port=8545 \
            --p2p.listen.ip=0.0.0.0 \
            --p2p.listen.tcp=9003 \
            --p2p.listen.udp=9003 \
            --snapshotlog.file=/op_log/snapshot.log \
            --p2p.priv.path=/config/p2p-node-key.txt \
            --metrics.enabled \
            --metrics.addr=0.0.0.0 \
            --metrics.port=7300 \
            --pprof.enabled \
            --rpc.enable-admin
    sleep 10
}

launch_op_node_bak() {
    docker run \
        --name op-node_bak -it --rm \
        -p "7546:8545" \
        -p "9004:9003" \
        -p "7303:7300" \
        -p "6063:6060" \
        -v "$OP_ROOT_DIR/ops-bedrock/p2p-sequencer-key.txt:/config/p2p-sequencer-key.txt" \
        -v "$OP_ROOT_DIR/ops-bedrock/p2p-node-key.txt:/config/p2p-node-key.txt" \
        -v "$OP_ROOT_DIR/ops-bedrock/test-jwt-secret.txt:/config/test-jwt-secret.txt" \
        -v "$OP_DATA_DIR/rollup.json:/rollup.json" \
        -v "$OP_DATA_DIR/op_log_bak:/op_log" \
        op-node:latest \
        op-node \
            --l1.trustrpc \
            --l1=$L1_ENDPOINT \
            --l2=http://host.docker.internal:8552 \
            --l2.jwt-secret=/config/test-jwt-secret.txt \
            --sequencer.enabled \
            --sequencer.stopped \
            --sequencer.l1-confs=0 \
            --verifier.l1-confs=0 \
            --p2p.sequencer.key="$OP_BATCHER_PRIVKEY" \
            --rollup.config=/rollup.json \
            --rpc.addr=0.0.0.0 \
            --rpc.port=8545 \
            --p2p.listen.ip=0.0.0.0 \
            --p2p.listen.tcp=9003 \
            --p2p.listen.udp=9003 \
            --snapshotlog.file=/op_log/snapshot.log \
            --p2p.priv.path=/config/p2p-node-key.txt \
            --metrics.enabled \
            --metrics.addr=0.0.0.0 \
            --metrics.port=7300 \
            --pprof.enabled \
            --rpc.enable-admin
    sleep 10
}

launch_op_batcher() {
    docker run \
        --name op-batcher -d -it --rm \
        -p "6061:6060" \
        -p "7301:7300" \
        --env OP_BATCHER_L1_ETH_RPC=$L1_ENDPOINT \
        --env OP_BATCHER_L2_ETH_RPC="http://host.docker.internal:9545" \
        --env OP_BATCHER_ROLLUP_RPC="http://host.docker.internal:7545" \
        --env OP_BATCHER_MAX_L1_TX_SIZE_BYTES=120000 \
        --env OP_BATCHER_TARGET_L1_TX_SIZE_BYTES=624 \
        --env OP_BATCHER_TARGET_NUM_FRAMES=1 \
        --env OP_BATCHER_APPROX_COMPR_RATIO="1.0" \
        --env OP_BATCHER_CHANNEL_TIMEOUT=40 \
        --env OP_BATCHER_POLL_INTERVAL="10ms" \
        --env OP_BATCHER_NUM_CONFIRMATIONS=1 \
        --env OP_BATCHER_SAFE_ABORT_NONCE_TOO_LOW_COUNT=3 \
        --env OP_BATCHER_RESUBMISSION_TIMEOUT="30s" \
        --env OP_BATCHER_PRIVATE_KEY="${OP_BATCHER_PRIVKEY}" \
        --env OP_BATCHER_SEQUENCER_BATCH_INBOX_ADDRESS="${OP_BATCHER_SEQUENCER_BATCH_INBOX_ADDRESS}" \
        --env OP_BATCHER_LOG_TERMINAL="true" \
        --env OP_BATCHER_PPROF_ENABLED="true" \
        --env OP_BATCHER_METRICS_ENABLED="true" \
        --env OP_BATCHER_SUB_SAFETY_MARGIN=20 \
        op-batcher:latest
}

launch_op_batcher_bak() {
    docker run \
        --name op-batcher_bak -it --rm \
        -p "6064:6060" \
        -p "7304:7300" \
        --env OP_BATCHER_L1_ETH_RPC=$L1_ENDPOINT \
        --env OP_BATCHER_L2_ETH_RPC="http://host.docker.internal:9546" \
        --env OP_BATCHER_ROLLUP_RPC="http://host.docker.internal:7546" \
        --env OP_BATCHER_MAX_L1_TX_SIZE_BYTES=120000 \
        --env OP_BATCHER_TARGET_L1_TX_SIZE_BYTES=624 \
        --env OP_BATCHER_TARGET_NUM_FRAMES=1 \
        --env OP_BATCHER_APPROX_COMPR_RATIO="1.0" \
        --env OP_BATCHER_CHANNEL_TIMEOUT=40 \
        --env OP_BATCHER_POLL_INTERVAL="10ms" \
        --env OP_BATCHER_NUM_CONFIRMATIONS=1 \
        --env OP_BATCHER_SAFE_ABORT_NONCE_TOO_LOW_COUNT=3 \
        --env OP_BATCHER_RESUBMISSION_TIMEOUT="30s" \
        --env OP_BATCHER_PRIVATE_KEY="${OP_BATCHER_PRIVKEY}" \
        --env OP_BATCHER_SEQUENCER_BATCH_INBOX_ADDRESS="${OP_BATCHER_SEQUENCER_BATCH_INBOX_ADDRESS}" \
        --env OP_BATCHER_LOG_TERMINAL="true" \
        --env OP_BATCHER_PPROF_ENABLED="true" \
        --env OP_BATCHER_METRICS_ENABLED="true" \
        --env OP_BATCHER_SUB_SAFETY_MARGIN=20 \
        op-batcher:latest
}

launch_op_proposer() {
    export L2OO_ADDRESS=$(cat $OP_DATA_DIR/addresses.json | jq '.L2OutputOracleProxy' | awk -F '"' '{print $2}')
    docker run \
        --name op-proposer -d -it --rm \
        -p "6062:6060" \
        -p "7302:7300" \
        -e OP_PROPOSER_L1_ETH_RPC=$L1_ENDPOINT \
        -e OP_PROPOSER_ROLLUP_RPC="http://host.docker.internal:7545" \
        -e OP_PROPOSER_POLL_INTERVAL="1s" \
        -e OP_PROPOSER_NUM_CONFIRMATIONS=1 \
        -e OP_PROPOSER_SAFE_ABORT_NONCE_TOO_LOW_COUNT=3 \
        -e OP_PROPOSER_RESUBMISSION_TIMEOUT="30s" \
        -e OP_PROPOSER_LOG_TERMINAL="true" \
        -e OP_PROPOSER_L2OO_ADDRESS=${L2OO_ADDRESS} \
        -e OP_PROPOSER_PPROF_ENABLED="true" \
        -e OP_PROPOSER_METRICS_ENABLED="true" \
        -e OP_PROPOSER_ALLOW_NON_FINALIZED="true" \
        -e OP_PROPOSER_PRIVATE_KEY="${OP_PROPOSER_PRIVKEY}" \
        op-proposer:latest
}

launch_op_proposer_bak() {
    export L2OO_ADDRESS=$(cat $OP_DATA_DIR/addresses.json | jq '.L2OutputOracleProxy' | awk -F '"' '{print $2}')
    docker run \
        --name op-proposer_bak -it --rm \
        -p "6065:6060" \
        -p "7305:7300" \
        -e OP_PROPOSER_L1_ETH_RPC=$L1_ENDPOINT \
        -e OP_PROPOSER_ROLLUP_RPC="http://host.docker.internal:7546" \
        -e OP_PROPOSER_POLL_INTERVAL="1s" \
        -e OP_PROPOSER_NUM_CONFIRMATIONS=1 \
        -e OP_PROPOSER_SAFE_ABORT_NONCE_TOO_LOW_COUNT=3 \
        -e OP_PROPOSER_RESUBMISSION_TIMEOUT="30s" \
        -e OP_PROPOSER_LOG_TERMINAL="true" \
        -e OP_PROPOSER_L2OO_ADDRESS=${L2OO_ADDRESS} \
        -e OP_PROPOSER_PPROF_ENABLED="true" \
        -e OP_PROPOSER_METRICS_ENABLED="true" \
        -e OP_PROPOSER_ALLOW_NON_FINALIZED="true" \
        -e OP_PROPOSER_PRIVATE_KEY="${OP_PROPOSER_PRIVKEY}" \
        op-proposer:latest
}

build_all() {
    cd $OP_ROOT_DIR

    docker build --file ops-bedrock/Dockerfile.l2 --tag l2geth:latest ops-bedrock

    docker build --file op-node/Dockerfile --tag op-node:latest .

    docker build --file op-batcher/Dockerfile --tag op-batcher:latest .

    docker build --file op-proposer/Dockerfile --tag op-proposer:latest .
}

main() {
    if [ "$#" = "0" ]; then
        usage
        exit 0
    fi

    command="$1"
    shift 1

    case $command in
        "clean")
            clean
            ;;
        "deploy")
            generate_deploy_config
            deploy_contracts
            ;;
        "generate_rollup_config")
            output_addresses
            generate_rollup_config
            ;;
        "build_all")
            build_all
            ;;
        "launch_all")
            launch_op_geth
            launch_op_node
            launch_op_batcher
            launch_op_proposer
            ;;
        "launch_all_bak")
            launch_op_geth_bak
            launch_op_node_bak
            launch_op_batcher_bak
            launch_op_proposer_bak
            ;;
        "launch_op_geth")
            launch_op_geth
            ;;
        "launch_op_node")
            launch_op_node
            ;;
        "launch_op_batcher")
            launch_op_batcher
            ;;
        "launch_op_proposer")
            launch_op_proposer
            ;;
        "launch_op_geth_bak")
            launch_op_geth_bak
            ;;
        "launch_op_node_bak")
            launch_op_node_bak
            ;;
        "launch_op_batcher_bak")
            launch_op_batcher_bak
            ;;
        "launch_op_proposer_bak")
            launch_op_proposer_bak
            ;;
        "stop_op_geth")
            docker stop l2
            ;;
        "stop_op_node")
            docker stop op-node
            ;;
        "stop_op_batcher")
            docker stop op-batcher
            ;;
        "stop_op_proposer")
            docker stop op-proposer
            ;;
        "start_admin")
            start_admin
            ;;
    esac
}

main "$@"
