// SPDX-License-Identifier: MIT

pragma solidity ^0.8.10;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "@openzeppelin/contracts/utils/math/SafeMath.sol";

contract AtomicBond {
    using SafeERC20 for IERC20;
    using SafeMath for uint256;

    address private immutable OWNER;

    address private immutable MIM;

    address private immutable TIME;
    address private immutable MEMO;

    address private immutable TIME_STAKING;
    address private immutable JOE_ROUTER_02;
    address private immutable MIM_TIME_BOND_DEPOSITORY;

    mapping(address => bool) private UserWhitelist;

    mapping(address => bool) private FunderWhitelist;
    mapping(address => bool) private MinterWhitelist;

    constructor(
        address _MIM,
        address _TIME,
        address _MEMO,
        address _TIME_STAKING,
        address _JOE_ROUTER_02,
        address _MIM_TIME_BOND_DEPOSITORY
    ) {
        require(_MIM != address(0));
        require(_TIME != address(0));
        require(_MEMO != address(0));
        require(_TIME_STAKING != address(0));
        require(_JOE_ROUTER_02 != address(0));
        require(_MIM_TIME_BOND_DEPOSITORY != address(0));

        OWNER = msg.sender;

        MIM = _MIM;

        TIME = _TIME;
        MEMO = _MEMO;

        TIME_STAKING = _TIME_STAKING;
        JOE_ROUTER_02 = _JOE_ROUTER_02;
        MIM_TIME_BOND_DEPOSITORY = _MIM_TIME_BOND_DEPOSITORY;

        IERC20(MEMO).approve(TIME_STAKING, type(uint256).max);
        IERC20(TIME).approve(JOE_ROUTER_02, type(uint256).max);
        IERC20(MIM).approve(MIM_TIME_BOND_DEPOSITORY, type(uint256).max);

        UserWhitelist[OWNER] = true;
    }

    modifier onlyOwner() {
        require(OWNER == msg.sender, "Not owner.");
        _;
    }

    modifier onlyUser() {
        require(UserWhitelist[msg.sender], "Not user.");
        _;
    }

    modifier ensure(uint256 _deadline) {
        require(_deadline >= block.timestamp, "Expired.");
        _;
    }

    function isFunder(address _address) private view returns (bool) {
        return FunderWhitelist[_address];
    }

    function isMinter(address _address) private view returns (bool) {
        return MinterWhitelist[_address];
    }

    function atomicBondMemoMim(
        address _funder,
        address _minter,
        uint256 _bondIn,
        // MIM > TIME
        address[] calldata _swapPath,
        uint256 _minSwapAmount, // 3200000000000000000000
        uint256 _maxBondPrice, // 320000
        // UNIX EPOCH
        uint256 _deadline
    ) external onlyUser ensure(_deadline) {
        require(_minSwapAmount > 0, "Invalid _minSwapAmount.");
        require(isFunder(_funder), "Not funder.");
        require(isMinter(_minter), "Not minter.");
        IERC20(MEMO).safeTransferFrom(_funder, address(this), _bondIn);
        TimeStaking(TIME_STAKING).unstake(_bondIn, true);
        uint256[] memory amounts = JoeRouter02(JOE_ROUTER_02)
            .swapExactTokensForTokens(
                _bondIn,
                _minSwapAmount,
                _swapPath,
                address(this),
                _deadline
            );
        TimeBondDepository(MIM_TIME_BOND_DEPOSITORY).deposit(
            amounts[amounts.length - 1],
            _maxBondPrice,
            _minter
        );
    }

    function redeem(address _minter) external onlyUser {
        TimeBondDepository(MIM_TIME_BOND_DEPOSITORY).redeem(_minter, true);
        IERC20 __memo = IERC20(MEMO);
        __memo.safeTransferFrom(_minter, OWNER, __memo.balanceOf(_minter));
    }

    function approveERC20(
        address _token,
        address _address,
        uint256 _amount
    ) external onlyOwner {
        IERC20(_token).approve(_address, _amount);
    }

    function approveUser(address _address) external onlyOwner {
        require(_address != OWNER, "Cannot approve owner.");
        UserWhitelist[_address] = true;
    }

    function revokeUser(address _address) external onlyOwner {
        require(_address != OWNER, "Cannot revoke owner.");
        delete UserWhitelist[_address];
    }

    function approveFunder(address _address) external onlyOwner {
        require(
            IERC20(MEMO).allowance(_address, address(this)) > 0,
            "Not approved."
        );
        FunderWhitelist[_address] = true;
    }

    function revokeFunder(address _address) external onlyOwner {
        delete FunderWhitelist[_address];
    }

    function approveMinter(address _address) external onlyOwner {
        require(
            IERC20(MEMO).allowance(_address, address(this)) > 0,
            "Not approved."
        );
        MinterWhitelist[_address] = true;
    }

    function revokeMinter(address _address) external onlyOwner {
        delete MinterWhitelist[_address];
    }

    function recoverLostNetworkToken() external onlyOwner {
        payable(OWNER).transfer(address(this).balance);
    }

    function recoverLostTokenERC20(address _token, uint256 _amount)
        external
        onlyOwner
    {
        IERC20(_token).safeTransferFrom(address(this), OWNER, _amount);
    }

    function finalize() external onlyOwner {
        selfdestruct(payable(OWNER));
    }
}

interface TimeStaking {
    function unstake(uint256 _amount, bool _trigger) external;
}

interface JoeRouter02 {
    function swapExactTokensForTokens(
        uint256 amountIn,
        uint256 amountOutMin,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external returns (uint256[] memory amounts);
}

interface TimeBondDepository {
    function deposit(
        uint256 _amount,
        uint256 _maxPrice,
        address _depositor
    ) external returns (uint256);

    function redeem(address _recipient, bool _stake) external returns (uint256);
}
