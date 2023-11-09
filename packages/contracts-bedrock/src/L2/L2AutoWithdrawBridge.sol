pragma solidity 0.8.15;

import { L2StandardBridge } from "./L2StandardBridge.sol";
import { Semver } from "../universal/Semver.sol";
import { FeeVault } from "../universal/FeeVault.sol";
import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract L2AutoWithdrawBridge is OwnableUpgradeable, FeeVault, Semver {
    address public constant L2_STANDARD_BRIDGE_ADDRESS = 0x4200000000000000000000000000000000000010;
    L2StandardBridge public L2_STANDARD_BRIDGE = L2StandardBridge(payable(L2_STANDARD_BRIDGE_ADDRESS));

    uint256 public constant DEFAULT_DELEGATION_FEE = 0.001 ether;
    uint256 public delegation_fee = DEFAULT_DELEGATION_FEE;


    // TODO add more useful fields
    event AutoWithdrawTo(
        address indexed from,
        address to,
        uint256 amount,
        bytes extraData
    );

    constructor(
        address _owner,
        address _recipient,
        uint256 _minWithdrawalAmount,
        WithdrawalNetwork _withdrawalNetwork
    ) FeeVault(_recipient, _minWithdrawalAmount, _withdrawalNetwork) Semver(1, 2, 0) {
        initialize(_owner);
    }

    function initialize(address _owner) public initializer {
        __Ownable_init();
        transferOwnership(_owner);
    }

    // TODO rename
    function autoWithdrawTo(
        address _l2Token,
        address _to,
        uint256 _amount,
        uint32 _minGasLimit,
        bytes calldata _extraData
    ) external payable virtual {
        require(msg.value == delegation_fee + _amount, "msg.value does not equal to delegation_fee + amount");

        emit AutoWithdrawTo(msg.sender, _to, _amount, _extraData);

        payable(address(this)).transfer(delegation_fee);

        L2_STANDARD_BRIDGE.withdrawTo{
            value: _amount
        }(_l2Token, _to, _amount, _minGasLimit, _extraData);
    }

    function setDelegateFee(uint256 _delegateFee) external onlyOwner {
        delegation_fee = _delegateFee;
    }

    function getDelegateFee() external view returns (uint256) {
        return delegation_fee;
    }
}
