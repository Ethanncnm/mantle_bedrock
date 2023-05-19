// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Predeploys } from "../libraries/Predeploys.sol";
import { MantlePortal } from "./MantlePortal.sol";
import { CrossDomainMessenger } from "../universal/CrossDomainMessenger.sol";
import { Semver } from "../universal/Semver.sol";
import { BridgeConstants } from "../libraries/BridgeConstants.sol";

/**
 * @custom:proxied
 * @title L1CrossDomainMessenger
 * @notice The L1CrossDomainMessenger is a message passing interface between L1 and L2 responsible
 *         for sending and receiving data on the L1 side. Users are encouraged to use this
 *         interface instead of interacting with lower-level contracts directly.
 */
contract L1CrossDomainMessenger is CrossDomainMessenger, Semver {
    /**
     * @notice Address of the MantlePortal.
     */
    MantlePortal public immutable PORTAL;

    /**
     * @custom:semver 1.1.0
     *
     * @param _portal Address of the MantlePortal contract on this network.
     */
    constructor(MantlePortal _portal)
        Semver(1, 1, 0)
        CrossDomainMessenger(Predeploys.L2_CROSS_DOMAIN_MESSENGER)
    {
        PORTAL = _portal;
        initialize();
    }

    /**
     * @notice Initializer.
     */
    function initialize() public initializer {
        __CrossDomainMessenger_init();
    }

    /**
     * @inheritdoc CrossDomainMessenger
     */
    function _sendMessage(
        uint32 _type,
        address _to,
        uint64 _gasLimit,
        uint256 _value,
        bytes memory _data
    ) internal override {
        if (_type == BridgeConstants.BIT_TX){
            PORTAL.depositTransaction{ value: 0 }(_to, _value, _gasLimit, false, _data);
        }else{
            PORTAL.depositTransaction{ value: _value }(_to, 0, _gasLimit, false, _data);
        }
    }

    /**
     * @inheritdoc CrossDomainMessenger
     */
    function _isOtherMessenger() internal view override returns (bool) {
        return msg.sender == address(PORTAL) && PORTAL.l2Sender() == OTHER_MESSENGER;
    }

    /**
     * @inheritdoc CrossDomainMessenger
     */
    function _isUnsafeTarget(address _target) internal view override returns (bool) {
        return _target == address(this) || _target == address(PORTAL);
    }
}
