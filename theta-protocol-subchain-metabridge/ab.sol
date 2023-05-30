// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import "@openzeppelin/contracts/token/ERC1155/utils/ERC1155Holder.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "./lib/Ownable.sol";
import "./lib/tokens/ITNT20.sol";
import "./lib/tokens/ITNT721.sol";
import "./lib/tokens/ITNT1155.sol";

interface ChainRegistrar {
    function getDynasty() external view returns (uint256 dynasty, bool success);

    function getCrossChainFee() external view returns (uint256);

    function getFallbackReceiver() external view returns (address);

    function isARegisteredSubchain(uint256 subchainID)
        external
        view
        returns (bool);

    function getSubchainRegistrationHeight(uint256 subchainID)
        external
        view
        returns (uint256, bool);

    function getNumBlocksPerDynasty() external view returns (uint256);

    function getValidatorSet(uint256 subchainID, uint256 dynasty)
        external
        view
        returns (address[] memory validators, uint256[] memory shareAmounts);
}

library TokenBankUtils {
    using SafeMath for uint256;

    function canonicalizeDenom(string memory denom)
        internal
        pure
        returns (string memory)
    {
        return toLower(denom);
    }

    function toLower(string memory base) internal pure returns (string memory) {
        bytes memory baseBytes = bytes(base);
        for (uint256 i = 0; i < baseBytes.length; i++) {
            baseBytes[i] = lower(baseBytes[i]);
        }
        return string(baseBytes);
    }

    function lower(bytes1 b1) internal pure returns (bytes1) {
        if (b1 >= 0x41 && b1 <= 0x5A) {
            return bytes1(uint8(b1) + 32);
        }
        return b1;
    }

    function extractChainIDFromDenom(string memory denom)
        internal
        pure
        returns (uint256, bool)
    {
        (uint256 idx, bool delimSuccess) = findDelimiterIndex(denom, "/", 1); // find the index of the first `/'
        if (!delimSuccess) {
            return (0, false);
        }

        (string memory chainIDStr, bool substringSuccess) = substring(
            denom,
            0,
            idx //bugfix
        );
        if (!substringSuccess) {
            return (0, false);
        }

        (uint256 chainID, bool strToUintSuccess) = stringToUint(chainIDStr);
        if (!strToUintSuccess) {
            return (0, false);
        }

        return (chainID, true);
    }

    function extractContractAddressFromDenom(string memory denom)
        internal
        pure
        returns (address, bool)
    {
        (uint256 idx, bool delimSuccess) = findDelimiterIndex(denom, "/", 2); // find the index of the first `/'
        if (!delimSuccess) {
            return (address(0), false);
        }

        (string memory contractAddrStr, bool substringSuccess) = substring(
            denom,
            idx + 1,
            bytes(denom).length
        );
        if (!substringSuccess) {
            return (address(0), false);
        }

        address contractAddr = parseAddr(contractAddrStr);
        return (contractAddr, true);
    }

    // find the count-th delimiter in str
    function findDelimiterIndex(
        string memory str,
        bytes1 delimiter,
        uint256 count
    ) private pure returns (uint256, bool) {
        bytes memory bs = bytes(str);
        uint256 len = bs.length;
        uint256 c = 0;
        for (uint256 i = 0; i < len; i++) {
            if (bs[i] == delimiter) {
                c += 1;
                if (c == count) {
                    return (i, true);
                }
            }
        }
        return (0, false);
    }

    function substring(
        string memory str,
        uint256 startIndex,
        uint256 endIndex
    ) private pure returns (string memory, bool) {
        uint256 len = bytes(str).length;
        if (endIndex < startIndex) {
            return ("", false);
        }
        if (endIndex > len) {
            //bugfix
            return ("", false);
        }
        bytes memory strBytes = bytes(str);
        bytes memory result = new bytes(endIndex.sub(startIndex));
        for (uint256 i = startIndex; i < endIndex; i++) {
            // endIndex: non-inclusive
            result[i.sub(startIndex)] = strBytes[i];
        }
        return (string(result), true);
    }

    function stringToUint(string memory str)
        private
        pure
        returns (uint256, bool)
    {
        bytes memory bs = bytes(str);
        uint256 len = bs.length;

        uint256 result = 0;
        for (uint256 i = 0; i < len; i++) {
            if (uint8(bs[i]) >= 48 && uint8(bs[i]) <= 57) {
                // result = result * 10 + (uint(b[i]) - 48);
                result = result.mul(10);
                result = result.add(uint8(bs[i]) - 48);
            } else {
                return (0, false);
            }
        }
        return (result, true);
    }

    function parseAddr(string memory str) private pure returns (address) {
        bytes memory tmp = bytes(str);
        uint160 iaddr = 0;
        uint160 b1;
        uint160 b2;
        for (uint256 i = 2; i < 2 + 2 * 20; i += 2) {
            iaddr *= 256;
            b1 = uint160(uint8(tmp[i]));
            b2 = uint160(uint8(tmp[i + 1]));
            if ((b1 >= 97) && (b1 <= 102)) {
                b1 -= 87;
            } else if ((b1 >= 65) && (b1 <= 70)) {
                b1 -= 55;
            } else if ((b1 >= 48) && (b1 <= 57)) {
                b1 -= 48;
            }
            if ((b2 >= 97) && (b2 <= 102)) {
                b2 -= 87;
            } else if ((b2 >= 65) && (b2 <= 70)) {
                b2 -= 55;
            } else if ((b2 >= 48) && (b2 <= 57)) {
                b2 -= 48;
            }
            iaddr += (b1 * 16 + b2);
        }
        return address(iaddr);
    }

    function bytesToUint256(bytes memory valBytes)
        private
        pure
        returns (uint256)
    {
        bytes32 val;
        uint256 len = valBytes.length;
        require(len <= 32, "bytes to uint256 conversion overflows");
        for (uint256 i = 0; i < len; i++) {
            val |= bytes32(valBytes[i] & 0xFF) >> ((32 - len + i) * 8);
        }
        return uint256(val);
    }

    // Check the existence of a contract method
    // Reference: https://ethereum.stackexchange.com/questions/83991/how-do-i-check-for-the-existence-of-a-function-in-solidity-when-i-have-an-addres
    bytes4 private constant NAME_FUNC_SELECTOR = bytes4(keccak256("name()"));

    function supportsTNT721Metadata(address tnt721Contract)
        internal
        returns (bool)
    {
        bool success;
        bytes memory data = abi.encodeWithSelector(NAME_FUNC_SELECTOR);
        assembly {
            success := call(
                gas(), // gas remaining
                tnt721Contract, // destination address
                0, // no ether
                add(data, 32), // input buffer (starts after the first 32 bytes in the `data` array)
                mload(data), // input length (loaded from the first 32 bytes in the `data` array)
                0, // output buffer
                0 // output length
            )
        }
        return success;
    }

    bytes4 private constant URI_FUNC_SELECTOR =
        bytes4(keccak256("uri(uint256)"));

    function supportsTNT1155MetadataURI(
        address tnt1155Contract,
        uint256 tokenID
    ) internal returns (bool) {
        bool success;
        bytes memory data = abi.encodeWithSelector(URI_FUNC_SELECTOR, tokenID);
        assembly {
            success := call(
                gas(), // gas remaining
                tnt1155Contract, // destination address
                0, // no ether
                add(data, 32), // input buffer (starts after the first 32 bytes in the `data` array)
                mload(data), // input length (loaded from the first 32 bytes in the `data` array)
                0, // output buffer
                0 // output length
            )
        }
        return success;
    }
}

contract TNT20VoucherContract is ERC20, Ownable {
    string private _denom;
    uint8 private _decimals;

    constructor(
        address owner_,
        string memory denom_,
        string memory name_,
        string memory symbol_,
        uint8 decimals_
    ) ERC20(name_, symbol_) Ownable(owner_) {
        _denom = denom_;
        _decimals = decimals_;
    }

    function mint(address voucherReceiver, uint256 mintedAmount)
        external
        ownerOnly
    {
        _mint(voucherReceiver, mintedAmount);
    }

    function burn(address voucherOwner, uint256 burnedAmount)
        external
        ownerOnly
    {
        uint256 ownerBalance = balanceOf(voucherOwner);
        require(
            ownerBalance >= burnedAmount,
            "Voucher owner does not have enough balance to burn"
        );

        address spender = owner;
        uint256 approvedAmount = allowance(voucherOwner, spender);
        require(
            approvedAmount >= burnedAmount,
            "Voucher owner did not approved enough amount to burn"
        );
        _spendAllowance(voucherOwner, spender, burnedAmount);

        _burn(voucherOwner, burnedAmount);
    }

    function denom() public view returns (string memory) {
        return _denom;
    }

    function decimals() public view virtual override returns (uint8) {
        return _decimals;
    }
}

contract TNT721VoucherContract is ERC721, Ownable {
    string private _denom;
    uint256 private _totalSupply;
    mapping(uint256 => string) private _tokenURIMap;

    constructor(
        address owner_,
        string memory demon_,
        string memory name_,
        string memory symbol_
    ) ERC721(name_, symbol_) Ownable(owner_) {
        _denom = demon_;
        _totalSupply = 0;
    }

    function mint(
        address voucherReceiver,
        uint256 tokenID,
        string memory tokenUri
    ) external ownerOnly {
        _safeMint(voucherReceiver, tokenID);
        _tokenURIMap[tokenID] = tokenUri;
        _totalSupply += 1;
    }

    function burn(address voucherOwner, uint256 tokenID) external ownerOnly {
        require(_totalSupply > 0, "no token to burn");
        address expectedOwner = ownerOf(tokenID);
        require(expectedOwner == voucherOwner, "only owner can burn");

        address spender = owner;
        require(
            _isApprovedOrOwner(spender, tokenID),
            "Voucher owner did not approve token burn"
        );

        _burn(tokenID);
        delete _tokenURIMap[tokenID];
        _totalSupply -= 1;
    }

    function denom() public view returns (string memory) {
        return _denom;
    }

    function totalSupply() external view returns (uint256) {
        return _totalSupply;
    }

    //
    // ERC721Metadata interface
    //

    function tokenURI(uint256 tokenID)
        public
        view
        virtual
        override
        returns (string memory)
    {
        return _tokenURIMap[tokenID];
    }

    function getContractOwner() public view returns (address) {
        return owner;
    }
}

contract TNT1155VoucherContract is ERC1155, Ownable {
    string private _denom;
    mapping(uint256 => string) private _tokenURIMap;

    constructor(address owner_, string memory demon_)
        ERC1155("")
        Ownable(owner_)
    {
        _denom = demon_;
    }

    function mint(
        address voucherReceiver,
        uint256 tokenID,
        uint256 mintedAmount,
        string memory tokenUri
    ) external ownerOnly {
        _mint(voucherReceiver, tokenID, mintedAmount, bytes(""));
        _tokenURIMap[tokenID] = tokenUri;
    }

    function burn(
        address voucherOwner,
        uint256 tokenID,
        uint256 burnedAmount
    ) external ownerOnly {
        require(
            balanceOf(voucherOwner, tokenID) >= burnedAmount,
            "Voucher owner does not have enough balance to burn"
        );

        address spender = owner;
        require(
            isApprovedForAll(voucherOwner, spender),
            "Voucher owner did not approve token burn"
        );

        _burn(voucherOwner, tokenID, burnedAmount);
    }

    function denom() public view returns (string memory) {
        return _denom;
    }

    function uri(uint256 tokenID)
        public
        view
        virtual
        override
        returns (string memory)
    {
        return _tokenURIMap[tokenID];
    }

    function getContractOwner() public view returns (address) {
        return owner;
    }
}

contract VoucherMap {
    struct VoucherInfo {
        address contractAddress;
        bool exists;
    }

    struct DenomInfo {
        string denom;
        bool exists;
    }

    mapping(string => VoucherInfo) public denomToVoucherLookup; // denom -> voucher contract address
    mapping(address => DenomInfo) public voucherAddressToDenomLookup; // voucher contract address -> denom
    address[] public allVouchers;
    string[] public allDenoms;

    function exists(string memory denom) external view returns (bool) {
        string memory canonicalDenom = TokenBankUtils.canonicalizeDenom(denom);
        return denomToVoucherLookup[canonicalDenom].exists;
    }

    function exists(address voucherAddress) external view returns (bool) {
        return voucherAddressToDenomLookup[voucherAddress].exists;
    }

    function getDenom(address voucherContractAddr)
        external
        view
        returns (string memory)
    {
        DenomInfo memory denomInfo = voucherAddressToDenomLookup[
            voucherContractAddr
        ];
        if (denomInfo.exists) {
            return denomInfo.denom;
        } else {
            return "";
        }
    }

    function getVoucher(string memory denom) external view returns (address) {
        string memory canonicalDenom = TokenBankUtils.canonicalizeDenom(denom);
        VoucherInfo memory voucherInfo = denomToVoucherLookup[canonicalDenom];
        if (voucherInfo.exists) {
            return voucherInfo.contractAddress;
        } else {
            return address(0);
        }
    }

    function addVoucher(string memory denom, address voucherContractAddr)
        internal
    {
        string memory canonicalDenom = TokenBankUtils.canonicalizeDenom(denom);
        require(
            !this.exists(canonicalDenom),
            "an voucher contract already exists"
        );
        denomToVoucherLookup[canonicalDenom] = VoucherInfo(
            voucherContractAddr,
            true
        );
        voucherAddressToDenomLookup[voucherContractAddr] = DenomInfo(
            canonicalDenom,
            true
        );
        allVouchers.push(voucherContractAddr);
        allDenoms.push(canonicalDenom);
    }

    function _buildDenom(
        uint256 chainID,
        string memory tokenType,
        string memory tokenAddress
    ) internal pure returns (string memory) {
        return
            TokenBankUtils.canonicalizeDenom(
                string(
                    abi.encodePacked(
                        Strings.toString(chainID),
                        "/",
                        tokenType,
                        "/",
                        tokenAddress
                    )
                )
            );
    }
}

contract TokenBank is ReentrancyGuard {
    using SafeMath for uint256;
    using Address for address;

    struct VoteData {
        uint256 dynasty;
        address[] signers;
        uint256 accumlatedShares;
    }

    event FailedToSendTFuel(address indexed receiver, uint256 amount);

    uint256 public mainchainID;

    ChainRegistrar chainRegistrar;

    // {chainID : eventNonce}
    mapping(uint256 => uint256) public tokenLockNonceMap;

    // {chainID : eventNonce}
    mapping(uint256 => uint256) public tokenUnlockNonceMap;

    // {chainID : eventNonce}
    mapping(uint256 => uint256) public voucherBurnNonceMap;

    // {chainID : eventNonce}
    mapping(uint256 => uint256) public voucherMintNonceMap;

    // {chainID : {eventNonce : height}}
    mapping(uint256 => mapping(uint256 => uint256)) tokenLockEventHeightMap;

    // {chainID : {eventNonce : height}}
    mapping(uint256 => mapping(uint256 => uint256)) voucherBurnEventHeightMap;

    // {chainID : maxProcessedTokenLockNonce}
    mapping(uint256 => uint256) maxProcessedTokenLockNonceMap;

    // {chainID : maxProcessedVoucherBurnNonce}
    mapping(uint256 => uint256) maxProcessedVoucherBurnNonceMap;

    // {chainID: {hash(tokenLockDataBytes) : VoteData}}
    mapping(uint256 => mapping(bytes32 => VoteData))
        public tokenLockVotingRecords;

    // {chainID: {hash(voucherBurnDataBytes) : VoteData}}
    mapping(uint256 => mapping(bytes32 => VoteData))
        public voucherBurnVotingRecords;

    constructor(uint256 mainchainID_, ChainRegistrar chainRegistrar_) {
        mainchainID = mainchainID_;
        chainRegistrar = chainRegistrar_;
    }

    function getMaxProcessedTokenLockNonce(uint256 chainID)
        external
        view
        returns (uint256)
    {
        return maxProcessedTokenLockNonceMap[chainID];
    }

    function getTokenLockEventHeight(uint256 chainID, uint256 eventNonce)
        external
        view
        returns (uint256)
    {
        uint256 updateHeight = tokenLockEventHeightMap[chainID][eventNonce];
        return updateHeight;
    }

    function getMaxProcessedVoucherBurnNonce(uint256 chainID)
        external
        view
        returns (uint256)
    {
        return maxProcessedVoucherBurnNonceMap[chainID];
    }

    function getVoucherBurnEventHeight(uint256 chainID, uint256 eventNonce)
        external
        view
        returns (uint256)
    {
        uint256 updateHeight = voucherBurnEventHeightMap[chainID][eventNonce];
        return updateHeight;
    }

    function _incrementTokenLockNonce(uint256 chainID) internal {
        tokenLockNonceMap[chainID] += 1;
        tokenLockEventHeightMap[chainID][tokenLockNonceMap[chainID]] = block
            .number;
    }

    function _incrementTokenUnlockNonce(uint256 chainID) internal {
        tokenUnlockNonceMap[chainID] += 1;
    }

    function _incrementVoucherBurnNonce(uint256 chainID) internal {
        voucherBurnNonceMap[chainID] += 1;
        voucherBurnEventHeightMap[chainID][voucherBurnNonceMap[chainID]] = block
            .number;
    }

    function _incrementVoucherMintNonce(uint256 chainID) internal {
        voucherMintNonceMap[chainID] += 1;
    }

    function _lockTokensSanityChecks(
        address sourceChainTokenAddress,
        uint256 targetChainID,
        address targetChainVoucherReceiver
    ) internal view returns (bool) {
        require(
            sourceChainTokenAddress != address(0x0),
            "source chain token address cannot be the zero address"
        );
        require(
            targetChainID != block.chainid,
            "cannot send to the same chain"
        );
        require(
            targetChainVoucherReceiver != address(0x0),
            "cannot send to zero address"
        );
        require(
            (block.chainid != mainchainID && targetChainID == mainchainID) ||
                (block.chainid == mainchainID &&
                    chainRegistrar.isARegisteredSubchain(targetChainID)),
            "not a valid targetChainID"
        ); // only support mainchain <-> registered subchain transfers, do not support transfers between two subchains for now
        return true;
    }

    function _checkValidatorQuorumForTokenLock(
        uint256 chainID,
        uint256 dynasty,
        bytes32 dataDigest,
        uint256 eventNonce,
        address msgSender
    ) internal returns (bool) {
        return
            _checkValidatorQuorum(
                chainID,
                dynasty,
                dataDigest,
                eventNonce,
                msgSender,
                maxProcessedTokenLockNonceMap,
                tokenLockVotingRecords
            );
    }

    function _checkValidatorQuorumForVoucherBurn(
        uint256 chainID,
        uint256 dynasty,
        bytes32 dataDigest,
        uint256 eventNonce,
        address msgSender
    ) internal returns (bool) {
        return
            _checkValidatorQuorum(
                chainID,
                dynasty,
                dataDigest,
                eventNonce,
                msgSender,
                maxProcessedVoucherBurnNonceMap,
                voucherBurnVotingRecords
            );
    }

    function _checkValidatorQuorum(
        uint256 chainID,
        uint256 dynasty,
        bytes32 dataDigest,
        uint256 eventNonce,
        address msgSender,
        mapping(uint256 => uint256) storage maxProcessedEventNonce,
        mapping(uint256 => mapping(bytes32 => VoteData)) storage votingRecords
    ) private returns (bool) {
        {
            (uint256 currDynasty, bool success) = chainRegistrar.getDynasty();
            require(success, "failed to get dynasty");
            require(
                currDynasty < 2 || dynasty >= currDynasty.sub(2), // a window of 2 dynasties for the older validator set to cast votes
                "Dynasty too old"
            ); // retired validators may sell their keys to malicious actors
        }

        require(
            chainID == mainchainID ||
                chainRegistrar.isARegisteredSubchain(chainID),
            "Invalid chainID"
        );

        if (eventNonce != maxProcessedEventNonce[chainID] + 1) {
            return false;
        }

        bool isMsgSenderAValidator = false;
        VoteData storage vd = votingRecords[chainID][dataDigest];

        uint256[] memory shares;
        {
            // use scope here to limit variable lifetime in order to avoid the "stack too deep" compilation error
            address[] memory validators;
            uint256 validationChainID = chainID; // validationChainID: the chain whose validate set should be used to perform the quorum check
            if (chainID == mainchainID) {
                // special handling: for events emitted by the main chain, should use the subchain validator set
                validationChainID = block.chainid; //bugfix
            }
            (validators, shares) = getAdjustedValidatorSet(
                validationChainID,
                dynasty
            );
            for (uint256 i = 0; i < validators.length; i++) {
                if (validators[i] != msgSender) {
                    continue;
                }

                isMsgSenderAValidator = true;
                for (uint256 j = 0; j < vd.signers.length; j++) {
                    require(
                        msgSender != vd.signers[j],
                        "This validator already voted"
                    );
                }

                vd.dynasty = dynasty;
                vd.signers.push(msgSender);
                vd.accumlatedShares = vd.accumlatedShares.add(shares[i]);
            }
        }
        require(isMsgSenderAValidator, "Not a validator");

        uint256 totalShares;
        for (uint256 j = 0; j < shares.length; j++) {
            totalShares = totalShares.add(shares[j]);
        }

        if (vd.accumlatedShares.mul(3) >= totalShares.mul(2)) {
            maxProcessedEventNonce[chainID] =
                maxProcessedEventNonce[chainID] +
                1;
            return true; // Note that the above line increments the maxProcessedNonce. Hence,
            // the next valid vote for the same nonce will be rejected by the first check.
            // Therefore, this method returns true AT MOST once for a given voucherBurnData.
        }

        return false;
    }

    function _splitFeeAmongSubchainValidators(uint256 targetChainID) internal {
        uint256 subchainID = _getSubchainID(targetChainID);
        uint256 crossChainFee = chainRegistrar.getCrossChainFee();
        if (crossChainFee == 0) {
            return;
        }

        (uint256 dynasty, bool success) = chainRegistrar.getDynasty();
        require(success, "failed to get dynasty");

        (
            address[] memory validators,
            uint256[] memory shareAmounts
        ) = getAdjustedValidatorSet(subchainID, dynasty);

        uint256 numValidators = validators.length;
        if (numValidators == 0) {
            return;
        }
        require(
            shareAmounts.length == numValidators,
            "validator and share amount count mismatch"
        );

        uint256 totalShares = 0;
        for (uint256 i = 0; i < numValidators; i++) {
            totalShares = totalShares.add(shareAmounts[i]);
        }
        if (totalShares == 0) {
            return;
        }

        uint256 remainingFee = crossChainFee;
        for (uint256 i = 0; i < numValidators; i++) {
            uint256 validatorFee = crossChainFee.mul(shareAmounts[i]).div(
                totalShares
            );

            // Do not send TFuel to an validator address if it is a smart contract. This is because
            //   1) by definition an validator wallet needs to be an EOA
            //   2) sending TFuel to a contract could cause problem. For example, the contract could
            //      implement a malicious fallback function, which use delegateCall to call the mintTFuel
            //      precompile on behalf of the TFuelTokenBank contract.
            if (!validators[i].isContract()) {
                if (!payable(validators[i]).send(validatorFee)) {
                    // it is OK if send fails (e.g. the validator address could be a malicious smart contract
                    // whose fallback function reverts), but we still issue an event as a warning
                    emit FailedToSendTFuel(validators[i], validatorFee);
                }
                remainingFee = remainingFee.sub(validatorFee);
            }
        }
        if (!validators[numValidators - 1].isContract()) {
            if (!payable(validators[numValidators - 1]).send(remainingFee)) {
                emit FailedToSendTFuel(
                    validators[numValidators - 1],
                    remainingFee
                );
            }
        }
    }

    function getAdjustedValidatorSet(uint256 subchainID, uint256 dynasty)
        public
        view
        returns (address[] memory validators, uint256[] memory shareAmounts)
    {
        bool isMainchain = (block.chainid == mainchainID);

        uint256 numBlocksPerDynasty = chainRegistrar.getNumBlocksPerDynasty();
        (uint256 chainRegistrationHeight, bool querySuccess) = chainRegistrar
            .getSubchainRegistrationHeight(subchainID);

        bool isInitialDynasty = querySuccess &&
            (chainRegistrationHeight >= dynasty.mul(numBlocksPerDynasty)) &&
            (chainRegistrationHeight <
                (dynasty.add(1)).mul(numBlocksPerDynasty));

        if (isMainchain && isInitialDynasty) {
            // Special handling for the initial dynasty query on the **Mainchain**:
            // The initial set of validators hardcoded in the genesis snapshot should be
            // in charge during the initial dyansty. However, if we directly query the
            // initial dynasty, we would obtain an empty set since the validators registered
            // on the mainchain takes charge only when the next dynasty begins. Hence, for
            // the initial dyansty, we query the next dynasty instead, which should return
            // a validator set that matches with the validator set hardcoded in the
            // genesis snapshot.
            //
            // Note: No special handling when this method is called on a subchain. This is
            //       OK since the pre-compiled contract on the Subchain for retrieving
            //       validdator set always returns the correct validator set for any given
            //       dynasty, including the "initial dynasty".
            uint256 nextDynasty = dynasty.add(1);
            return chainRegistrar.getValidatorSet(subchainID, nextDynasty);
        } else {
            return chainRegistrar.getValidatorSet(subchainID, dynasty);
        }
    }

    function _getSubchainID(uint256 targetChainID)
        internal
        view
        returns (uint256)
    {
        uint256 subchainID;
        if (targetChainID == mainchainID) {
            subchainID = block.chainid; // subchain to mainchain transfers
        } else {
            subchainID = targetChainID; // mainchain to subchain transfers
        }
        return subchainID;
    }
}

//
// Similar to the modularized Cosmos IBC design, we implement one Token Bank contract for
// each type of token, e.g. TFuel, TNT20, TNT721. In particular, TNT20TokenBank
// corresponds to the ICS20 module, and TNT721TokenBank corresponds to the ICS721
// module, and so on. This design achieves good extensibility. In the future, we can add
// Token Bank contracts for more types of tokens (e.g. TNT1155).
//
// SourceChain: The chain on which the user action ("token lock" or "voucher burn") is initiated.
//              *  Example 1: Assume token X was deployed on chain A. User Alice locked token X on chain A
//                 in order to transfer the token to chain B. In this scenario, chain A is the source chain since the
//                 "token lock" user action happened on chain A. On the other hand, chain B is the target chain.
//              *  Example 2: Say Alice burned voucher vX on chain B to unlock the authentic token X on chain A.
//                 In this case, chain B is the source chain since the "voucher burn" user action was initiated on
//                 chain B. On the other hand, chain A is the target chain.
// TargetChain: The chain to which the validators relays the user action.
//              *  In the above example 1, chain B is the target chain.
//              *  In the above example 2, chain A is the target chain.
//

contract TFuelTokenBank is TokenBank, VoucherMap {
    using SafeMath for uint256;

    string constant tfuelAddressPlaceholder =
        "0x0000000000000000000000000000000000000000";
    bool public isOnMainchain;

    constructor(uint256 mainchainID_, ChainRegistrar chainRegistrar_)
        TokenBank(mainchainID_, chainRegistrar_)
    {}

    // {chainID: totalLockedAmount}
    mapping(uint256 => uint256) public totalLockedAmounts;

    // TODO: indexed targetChainID

    // For TFuelTokenLocked events, sourceChain is always the main chain, and targetChain is always a subchain
    event TFuelTokenLocked(
        string denom,
        address sourceChainTokenSender,
        uint256 targetChainID,
        address targetChainVoucherReceiver,
        uint256 lockedAmount,
        uint256 tokenLockNonce
    );

    // For TFuelVoucherMinted events, sourceChain is always the main chain, and targetChain is always a subchain
    event TFuelVoucherMinted(
        string denom,
        address targetChainVoucherReceiver,
        uint256 mintedAmount,
        uint256 sourceChainTokenLockNonce,
        uint256 voucherMintNonce
    );

    // For TFuelVoucherBurned events, sourceChain is always a subchain, and targetChain is always the main chain
    event TFuelVoucherBurned(
        string denom,
        address sourceChainVoucherOwner,
        address targetChainTokenReceiver,
        uint256 burnedAmount,
        uint256 voucherBurnNonce
    );

    // For TFuelTokenUnlocked events, sourceChain is always a subchain, and targetChain is always the main chain
    event TFuelTokenUnlocked(
        string denom,
        address targetChainTokenReceiver,
        uint256 unlockedAmount,
        uint256 sourceChainVoucherBurnNonce,
        uint256 tokenUnlockNonce
    );

    event UnlockTokensToFallbackReceiver(
        address fallbackReceiver,
        uint256 amount
    );

    event MintVouchersToFallbackReceiver(
        address fallbackReceiver,
        uint256 amount
    );

    //
    // Type 1 transfer handlers: lock authentic tokens on the source chain, and mint vouchers on the target chain
    //

    // Type 1 transfer handler on the mainchain: send TFuel to a subchain. This function call locks TFuel in the contract, and the corresponding TFuel vouchers will be minted on the subchain.
    // Note: This method can be called by the end users. sourceChain is the main chain, while targetChain is a subchain.
    function lockTokens(
        uint256 targetChainID,
        address targetChainVoucherReceiver
    ) external payable nonReentrant {
        require(
            targetChainID != block.chainid,
            "cannot send to the same chain"
        );
        require(
            targetChainVoucherReceiver != address(0x0),
            "cannot send to zero address"
        );
        require(
            block.chainid == mainchainID,
            "TFuelTokenBank.lockTokens() can only be called on the main chain"
        );
        require(
            chainRegistrar.isARegisteredSubchain(targetChainID),
            "targetChainID not yet registered"
        );
        require(
            msg.value > chainRegistrar.getCrossChainFee(), // strictly larger than, since we require the amount of TFuel transferred cross-chain is non-zero
            "not enough TFuel (msg.value) to cover the cross-chain transaction fee"
        );

        uint256 lockedAmount = msg.value.sub(chainRegistrar.getCrossChainFee());
        address sourceChainTokenSender = msg.sender;
        totalLockedAmounts[targetChainID] = totalLockedAmounts[targetChainID]
            .add(lockedAmount);
        _incrementTokenLockNonce(targetChainID);

        string memory canonicalDenom = _buildDenom(
            mainchainID,
            "0",
            tfuelAddressPlaceholder
        );
        uint256 tokenLockNonce = tokenLockNonceMap[targetChainID];
        emit TFuelTokenLocked(
            canonicalDenom,
            sourceChainTokenSender,
            targetChainID,
            targetChainVoucherReceiver,
            lockedAmount,
            tokenLockNonce
        );

        // fee split for the validators
        _splitFeeAmongSubchainValidators(targetChainID);
    }

    // Type 1 transfer handler on a subchain: mint vouchers for TFuel transfered from the main chain
    // Note: This method can only be called by the the validators/orchestrators of the subchains.
    function mintVouchers(
        string memory denom,
        address targetChainVoucherReceiver,
        uint256 mintedAmount,
        uint256 dynasty,
        uint256 sourceChainTokenLockNonce
    ) external nonReentrant {
        require(
            targetChainVoucherReceiver != address(0x0),
            "cannot send to zero address"
        );
        require(
            block.chainid != mainchainID,
            "TFuelTokenBank.mintVouchers() can only be called on a subchain"
        );
        string memory canonicalDenom = TokenBankUtils.canonicalizeDenom(denom);
        (uint256 sourceChainID, bool extractionSuccess) = TokenBankUtils
            .extractChainIDFromDenom(canonicalDenom);
        require(extractionSuccess, "Failed to extract chainID from denom");
        bytes32 tokenLockDataDigest = keccak256(
            abi.encodePacked(
                sourceChainID,
                canonicalDenom,
                targetChainVoucherReceiver,
                mintedAmount,
                dynasty,
                sourceChainTokenLockNonce
            )
        );
        bool voteSuccess = _checkValidatorQuorumForTokenLock(
            sourceChainID,
            dynasty,
            tokenLockDataDigest,
            sourceChainTokenLockNonce,
            msg.sender
        );
        if (voteSuccess) {
            _mintTFuelVoucher(targetChainVoucherReceiver, mintedAmount);
            _incrementVoucherMintNonce(block.chainid);
            uint256 voucherMintNonce = voucherMintNonceMap[block.chainid];
            emit TFuelVoucherMinted(
                denom,
                targetChainVoucherReceiver,
                mintedAmount,
                sourceChainTokenLockNonce,
                voucherMintNonce
            );
        }
    }

    //
    // Type 2 transfer handlers: burn vouchers on the source chain, and unlock authentic tokens on the target chain
    //

    // Type 2 transfer handler: burn the TFuel vouchers in order to send TFuel back to the main chain
    // Note: This method can be called by the end users.
    function burnVouchers(address targetChainTokenReceiver)
        external
        payable
        nonReentrant
    {
        require(
            targetChainTokenReceiver != address(0x0),
            "cannot send to zero address"
        );
        require(
            block.chainid != mainchainID,
            "TFuelTokenBank.burnVouchers() can only be called on a subchain"
        );
        require( // msg.value amount of TFuel (in TFuelWei) was sent from the msg.sender to the SubchainTFuelTokenBank contract
            msg.value > chainRegistrar.getCrossChainFee(), // strictly larger than, since we require the amount of TFuel transferred cross-chain is non-zero
            "not enough TFuel (msg.value) to cover the cross-chain transaction fee"
        );

        address sourceChainVoucherOwner = msg.sender;
        uint256 burnedAmount = msg.value.sub(chainRegistrar.getCrossChainFee());
        _burnTFuelVoucher(burnedAmount); // now _burnTFuelVoucher() removes msg.value amount of TFuel from the contract's TFuel balance
        uint256 targetChainID = mainchainID; // the target chain for TFuel voucher burn can ONLY be the main chain
        _incrementVoucherBurnNonce(targetChainID);

        string memory canonicalDenom = _buildDenom(
            mainchainID,
            "0",
            tfuelAddressPlaceholder
        );
        uint256 voucherBurnNonce = voucherBurnNonceMap[targetChainID];
        emit TFuelVoucherBurned(
            canonicalDenom,
            sourceChainVoucherOwner,
            targetChainTokenReceiver,
            burnedAmount,
            voucherBurnNonce
        );

        // fee split for the subchain validators
        _splitFeeAmongSubchainValidators(mainchainID); // for TFuel voucher burn, the targetChain is always the mainchain
    }

    // Type 2 transfer handler on the mainchain: unlock TFuel after TFuel voucher burn events were detected on a subchain.
    // Note: This method can only be called by the validators/orchestrators of the subchains.
    function unlockTokens(
        uint256 sourceChainID,
        address payable targetChainTokenReceiver,
        uint256 unlockAmount,
        uint256 dynasty,
        uint256 sourceChainVoucherBurnNonce
    ) external nonReentrant {
        require(
            targetChainTokenReceiver != address(0x0),
            "cannot send to zero address"
        );
        require(
            block.chainid == mainchainID,
            "TFuelTokenBank.unlockTokens() can only be called on the main chain"
        );
        require(
            unlockAmount <= totalLockedAmounts[sourceChainID],
            "Cannot unlock the requested amount of TFuel"
        );
        bytes memory voucherBurnData = abi.encodePacked(
            sourceChainID,
            unlockAmount,
            targetChainTokenReceiver,
            dynasty,
            sourceChainVoucherBurnNonce
        );
        bytes32 voucherBurnDataDigest = keccak256(voucherBurnData);
        bool voteSuccess = _checkValidatorQuorumForVoucherBurn(
            sourceChainID,
            dynasty,
            voucherBurnDataDigest,
            sourceChainVoucherBurnNonce,
            msg.sender
        );
        if (voteSuccess) {
            totalLockedAmounts[sourceChainID] = totalLockedAmounts[
                sourceChainID
            ].sub(unlockAmount);

            // TFuelTokenBank.unlockTokens() can only be called on the main chain, and the main chain
            // does not implement the mintTFuel precompile contract, so it is not possible for attackers
            // to leverage delegateCall to mint TFuel vouchers. Therefore, we don't need to check if
            // targetChainTokenReceiver is an EOA
            if (!targetChainTokenReceiver.send(unlockAmount)) {
                address payable fallbackReceiver = payable(
                    chainRegistrar.getFallbackReceiver()
                );
                if (!fallbackReceiver.send(unlockAmount)) {
                    emit FailedToSendTFuel(fallbackReceiver, unlockAmount);
                }

                emit UnlockTokensToFallbackReceiver(
                    fallbackReceiver,
                    unlockAmount
                );
            }

            _incrementTokenUnlockNonce(sourceChainID);

            string memory canonicalDenom = _buildDenom(
                mainchainID,
                "0",
                tfuelAddressPlaceholder
            );
            uint256 tokenUnlockNonce = tokenUnlockNonceMap[sourceChainID];
            emit TFuelTokenUnlocked(
                canonicalDenom,
                targetChainTokenReceiver,
                unlockAmount,
                sourceChainVoucherBurnNonce,
                tokenUnlockNonce
            );
        }
    }

    function _mintTFuelVoucher(address voucherReceiver, uint256 mintedAmount)
        private
    {
        bytes memory data = new bytes(20 + 32);
        bytes32 mintAmountBytes = bytes32(mintedAmount);
        uint256 idx = 0;
        uint256 i;
        for (i = 0; i < 20; i++) {
            data[idx++] = (bytes20)(voucherReceiver)[i];
        }
        for (i = 0; i < 32; i++) {
            data[idx++] = mintAmountBytes[i];
        }

        // 0xb6: precompiled contract for minting TFuel. This precompiled contract has the execution previlege checks,
        //       which ensures that only the whiteliested contracts can mint TFuel. try-catch not needed here since it is
        //       a call to a precompiled contract
        (bool success, ) = address(0xb6).call(data);
        require(
            success,
            "TFuelTokenBank._mintTFuelVoucher(): failed to mint TFuel vouchers"
        );
    }

    function _burnTFuelVoucher(uint256 burnedAmount) private {
        bytes memory data = new bytes(32);
        bytes32 burnAmountBytes = bytes32(burnedAmount); // unit: TFuelWei
        for (uint256 i = 0; i < 32; i++) {
            data[i] = burnAmountBytes[i];
        }

        // 0xb7: precompiled contract for burning TFuel
        (bool success, ) = address(0xb7).call(data);
        require(
            success,
            "TFuelTokenBank._burnTFuelVoucher(): failed to burn TFuel vouchers"
        );
    }
}

// TNT20 Token Bank
contract TNT20TokenBank is TokenBank, VoucherMap {
    using SafeMath for uint256;

    event TNT20TokenLocked(
        string denom,
        address sourceChainTokenSender,
        uint256 targetChainID,
        address targetChainVoucherReceiver,
        uint256 lockedAmount,
        string name,
        string symbol,
        uint8 decimals,
        uint256 tokenLockNonce
    );

    event TNT20VoucherMinted(
        string denom,
        address targetChainVoucherReceiver,
        address voucherContract,
        uint256 mintedAmount,
        uint256 sourceChainTokenLockNonce,
        uint256 voucherMintNonce
    );

    event TNT20VoucherBurned(
        string denom,
        address sourceChainVoucherOwner,
        address targetChainTokenReceiver,
        uint256 burnedAmount,
        uint256 voucherBurnNonce
    );

    event TNT20TokenUnlocked(
        string denom,
        address targetChainTokenReceiver,
        uint256 unlockedAmount,
        uint256 sourceChainVoucherBurnNonce,
        uint256 tokenUnlockNonce
    );

    event FailedToUnlockTNT20Tokens(
        address indexed tokenContractAddress,
        address indexed receiver,
        uint256 amount
    );

    event FailedToMintTNT20Vouchers(
        address indexed voucherContractAddress,
        address indexed receiver,
        uint256 amount
    );

    // event UnlockTokensToFallbackReceiver(
    //     address fallbackReceiver,
    //     address tokenContractAddress,
    //     uint256 amount
    // );

    // event MintVouchersToFallbackReceiver(
    //     address fallbackReceiver,
    //     address voucherContractAddress,
    //     uint256 amount
    // );

    // Note: do not record the totalLockedAmounts for TNT tokens, since these tokens may
    //       have elastic supplies (e.g. AMPL) which violates the conservation rules. This
    //       cause the unlockTokens() to revert and stall the cross-chain transfer pipeline
    // // {chainID : {tnt20Contract : totalLockedAmount}}
    // mapping(uint256 => mapping(address => uint256)) public totalLockedAmounts;

    constructor(uint256 mainchainID_, ChainRegistrar chainRegistrar_)
        TokenBank(mainchainID_, chainRegistrar_)
    {}

    //
    // Type 1 transfer handlers: Sending authentic tokens to the target chain, which locks authentic tokens on the source chain, and mint vouchers on the target chain
    //

    // Type 1 transfer handler on the source chain: lock the authentic TNT20 tokens on the source chain.
    // Note: This method can be called by the end users.
    function lockTokens(
        uint256 targetChainID,
        address sourceChainTNT20Contract,
        address targetChainVoucherReceiver,
        uint256 lockAmount
    ) external payable nonReentrant {
        require(
            _lockTokensSanityChecks(
                sourceChainTNT20Contract,
                targetChainID,
                targetChainVoucherReceiver
            ),
            "lock tokens sanity checks failed"
        );

        ITNT20 tnt20 = ITNT20(sourceChainTNT20Contract);
        {
            uint256 balanceBefore = tnt20.balanceOf(address(this));
            tnt20.transferFrom(msg.sender, address(this), lockAmount);
            uint256 balanceAfter = tnt20.balanceOf(address(this));
            require(
                balanceAfter.sub(lockAmount) == balanceBefore,
                "Failed to transfer TNT20 tokens to the token bank"
            ); // make sure the transfer succeeded, as the token contract could be malicious
        }

        string memory canonicalDenom = _buildDenom(
            block.chainid,
            "20",
            Strings.toHexString(uint256(uint160(sourceChainTNT20Contract)), 20)
        );

        // totalLockedAmounts[targetChainID][
        //     sourceChainTNT20Contract
        // ] = totalLockedAmounts[targetChainID][sourceChainTNT20Contract].add(
        //     lockAmount
        // );

        _incrementTokenLockNonce(targetChainID);

        uint256 tokenLockNonce = tokenLockNonceMap[targetChainID];
        emit TNT20TokenLocked(
            canonicalDenom,
            msg.sender,
            targetChainID,
            targetChainVoucherReceiver,
            lockAmount,
            tnt20.name(),
            tnt20.symbol(),
            tnt20.decimals(),
            tokenLockNonce
        );

        // fee split for the validators
        require(
            msg.value >= chainRegistrar.getCrossChainFee(),
            "not enough TFuel (msg.value) to cover the cross-chain transaction fee"
        );
        _splitFeeAmongSubchainValidators(targetChainID);
    }

    // Type 1 transfer handler on the target chain: mint vouchers on the target chain
    // Note: This method can only be called by the validators/orchestrators of the subchains.
    function mintVouchers(
        string memory denom,
        string memory name,
        string memory symbol,
        uint8 decimals,
        address targetChainVoucherReceiver,
        uint256 mintedAmount,
        uint256 dynasty,
        uint256 sourceChainTokenLockNonce
    ) external nonReentrant {
        require(
            targetChainVoucherReceiver != address(0x0),
            "cannot send to zero address"
        );
        string memory canonicalDenom = TokenBankUtils.canonicalizeDenom(denom);
        (uint256 sourceChainID, bool extractionSuccess) = TokenBankUtils
            .extractChainIDFromDenom(canonicalDenom);
        require(extractionSuccess, "Failed to extract chainID from denom");
        bytes32 tokenLockDataDigest = keccak256(
            abi.encodePacked(
                sourceChainID,
                canonicalDenom,
                name,
                symbol,
                decimals,
                targetChainVoucherReceiver,
                mintedAmount,
                dynasty,
                sourceChainTokenLockNonce
            )
        );
        bool voteSuccess = _checkValidatorQuorumForTokenLock(
            sourceChainID,
            dynasty,
            tokenLockDataDigest,
            sourceChainTokenLockNonce,
            msg.sender
        );
        if (voteSuccess) {
            if (!this.exists(canonicalDenom)) {
                addVoucher(
                    canonicalDenom,
                    _deployTNT20VoucherContract(
                        canonicalDenom,
                        name,
                        symbol,
                        decimals
                    )
                );
            }

            address voucherContractAddr = this.getVoucher(canonicalDenom);
            require(
                voucherContractAddr != address(0),
                "the voucher for the denom does not exist"
            );
            _mintTNT20Voucher(
                voucherContractAddr,
                targetChainVoucherReceiver,
                mintedAmount
            );
            _incrementVoucherMintNonce(sourceChainID);

            uint256 voucherMintNonce = voucherMintNonceMap[sourceChainID];
            emit TNT20VoucherMinted(
                canonicalDenom,
                targetChainVoucherReceiver,
                voucherContractAddr,
                mintedAmount,
                sourceChainTokenLockNonce,
                voucherMintNonce
            );
        }
    }

    //
    // Type 2 transfer handlers: burn vouchers on the source chain, and unlock authentic tokens on the target chain
    //

    // Type 2 transfer handler on the source chain: burn vouchers on the source chain
    // Note: This method can be called by the end users.
    function burnVouchers(
        address sourceChainVoucherContractAddr,
        address targetChainTokenReceiver,
        uint256 burnedAmount
    ) external payable nonReentrant {
        require(
            targetChainTokenReceiver != address(0x0),
            "cannot send to zero address"
        );
        string memory canonicalDenom = this.getDenom(
            sourceChainVoucherContractAddr
        );
        require(
            bytes(canonicalDenom).length != 0,
            "failed to lookup denom for the voucher contract"
        );
        require(this.exists(canonicalDenom), "Voucher contract does not exist");
        address sourceChainVoucherOwner = msg.sender;
        _burnTNT20Voucher(
            sourceChainVoucherContractAddr,
            sourceChainVoucherOwner,
            burnedAmount
        );
        (uint256 targetChainID, bool extractionSuccess) = TokenBankUtils
            .extractChainIDFromDenom(canonicalDenom); // for voucher burn, target chain is always the chain where the authentic token was deployed
        require(
            extractionSuccess,
            "Failed to extract targetChainID from denom"
        );
        _incrementVoucherBurnNonce(targetChainID);
        uint256 voucherBurnNonce = voucherBurnNonceMap[targetChainID];
        emit TNT20VoucherBurned(
            canonicalDenom,
            sourceChainVoucherOwner,
            targetChainTokenReceiver,
            burnedAmount,
            voucherBurnNonce
        );

        // fee split for the validators
        require(
            msg.value >= chainRegistrar.getCrossChainFee(),
            "not enough TFuel (msg.value) to cover the cross-chain transaction fee"
        );
        _splitFeeAmongSubchainValidators(targetChainID);
    }

    // Type 2 transfer handler on the target chain: unlock the authentic tokens on the target chain
    // Note: This method can only be called by the validators/orchestrators of the subchains.
    function unlockTokens(
        uint256 sourceChainID,
        string memory denom,
        address targetChainTokenReceiver,
        uint256 unlockAmount,
        uint256 dynasty,
        uint256 sourceChainVoucherBurnNonce
    ) external nonReentrant {
        require(
            targetChainTokenReceiver != address(0x0),
            "cannot send to zero address"
        );
        string memory canonicalDenom = TokenBankUtils.canonicalizeDenom(denom);
        (
            address targetChainTNT20Contract,
            bool extractionSuccess
        ) = TokenBankUtils.extractContractAddressFromDenom(canonicalDenom);
        require(
            extractionSuccess,
            "Failed to extract contract address from denom"
        );
        // require(
        //     unlockAmount <=
        //         totalLockedAmounts[sourceChainID][targetChainTNT20Contract],
        //     "Cannot unlock the requested amount of tokens"
        // );

        bytes32 voucherBurnDataDigest = keccak256(
            abi.encodePacked(
                sourceChainID,
                canonicalDenom,
                unlockAmount,
                targetChainTokenReceiver,
                dynasty,
                sourceChainVoucherBurnNonce
            )
        );
        bool voteSuccess = _checkValidatorQuorumForVoucherBurn(
            sourceChainID,
            dynasty,
            voucherBurnDataDigest,
            sourceChainVoucherBurnNonce,
            msg.sender
        );

        if (voteSuccess) {
            // totalLockedAmounts[sourceChainID][
            //     targetChainTNT20Contract
            // ] = totalLockedAmounts[sourceChainID][targetChainTNT20Contract].sub(
            //     unlockAmount
            // );

            // using try-catch to prevent the external contract call reverts
            try
                ITNT20(targetChainTNT20Contract).transfer(
                    targetChainTokenReceiver,
                    unlockAmount
                )
            {} catch {
                emit FailedToUnlockTNT20Tokens(
                    targetChainTNT20Contract,
                    targetChainTokenReceiver,
                    unlockAmount
                );
                // address fallbackReceiver = chainRegistrar.getFallbackReceiver();
                // ITNT20(targetChainTNT20Contract).transfer(
                //     fallbackReceiver,
                //     unlockAmount
                // );
                // emit UnlockTokensToFallbackReceiver(
                //     fallbackReceiver,
                //     targetChainTNT20Contract,
                //     unlockAmount
                // );
            }

            _incrementTokenUnlockNonce(sourceChainID);

            uint256 tokenUnlockNonce = tokenUnlockNonceMap[sourceChainID];
            emit TNT20TokenUnlocked(
                canonicalDenom,
                targetChainTokenReceiver,
                unlockAmount,
                sourceChainVoucherBurnNonce,
                tokenUnlockNonce
            );
        }
    }

    function _deployTNT20VoucherContract(
        string memory denom,
        string memory name,
        string memory symbol,
        uint8 decimals
    ) private returns (address) {
        address contractOwner = address(this);
        string memory voucherName = string(abi.encodePacked(name, " Voucher"));
        string memory voucherSymbol = string(abi.encodePacked("v", symbol));
        TNT20VoucherContract tnt20Voucher = new TNT20VoucherContract(
            contractOwner,
            denom,
            voucherName,
            voucherSymbol,
            decimals
        );
        return address(tnt20Voucher);
    }

    function _mintTNT20Voucher(
        address voucherContractAddress,
        address voucherReceiver,
        uint256 mintAmount
    ) private {
        TNT20VoucherContract tnt20Voucher = TNT20VoucherContract(
            voucherContractAddress
        );

        // using try-catch to prevent the external contract call reverts
        try tnt20Voucher.mint(voucherReceiver, mintAmount) {} catch {
            emit FailedToMintTNT20Vouchers(
                address(tnt20Voucher),
                voucherReceiver,
                mintAmount
            );
            // address fallbackReceiver = chainRegistrar.getFallbackReceiver();
            // tnt20Voucher.mint(
            //     fallbackReceiver, // mint to the fallback address if mint fails, e.g. if voucherReceiver is a malicious contract
            //     mintAmount
            // );
            // emit MintVouchersToFallbackReceiver(
            //     fallbackReceiver,
            //     address(tnt20Voucher),
            //     mintAmount
            // );
        }
    }

    function _burnTNT20Voucher(
        address voucherContractAddress,
        address voucherOwner,
        uint256 burnedAmount
    ) private {
        TNT20VoucherContract tnt20Voucher = TNT20VoucherContract(
            voucherContractAddress
        );
        tnt20Voucher.burn(voucherOwner, burnedAmount);
    }
}

// TNT721 Token Bank
contract TNT721TokenBank is TokenBank, VoucherMap {
    using SafeMath for uint256;

    event TNT721TokenLocked(
        string denom,
        address sourceChainTokenSender,
        uint256 targetChainID,
        address targetChainVoucherReceiver,
        uint256 tokenID,
        string name,
        string symbol,
        string tokenURI,
        uint256 tokenLockNonce
    );

    event TNT721VoucherMinted(
        string denom,
        address targetChainVoucherReceiver,
        address voucherContract,
        uint256 tokenID,
        uint256 sourceChainTokenLockNonce,
        uint256 voucherMintNonce
    );

    event TNT721VoucherBurned(
        string denom,
        address sourceChainVoucherOwner,
        address targetChainTokenReceiver,
        uint256 tokenID,
        uint256 voucherBurnNonce
    );

    event TNT721TokenUnlocked(
        string denom,
        address targetChainTokenReceiver,
        uint256 tokenID,
        uint256 sourceChainVoucherBurnNonce,
        uint256 tokenUnlockNonce
    );

    event FailedToUnlockTNT721Tokens(
        address indexed tokenContractAddress,
        address indexed receiver,
        uint256 tokenID
    );

    event FailedToMintTNT721Vouchers(
        address indexed voucherContractAddress,
        address indexed receiver,
        uint256 tokenID
    );

    // event UnlockTokensToFallbackReceiver(
    //     address fallbackReceiver,
    //     address tokenContractAddress,
    //     uint256 tokenID
    // );

    // event MintVouchersToFallbackReceiver(
    //     address fallbackReceiver,
    //     address voucherContractAddress,
    //     uint256 tokenID
    // );

    // Note: do not record the totalLockedAmounts for TNT tokens, since these tokens may
    //       have elastic supplies (e.g. AMPL) which violates the conservation rules. This
    //       cause the unlockTokens() to revert and stall the cross-chain transfer pipeline
    // // {chainID : {tnt721Contract : {tokenID : totalLockedAmount}}}, where totalLockedAmount can only be either 0 or 1
    // mapping(uint256 => mapping(address => mapping(uint256 => uint256)))
    //     public totalLockedAmounts;

    constructor(uint256 mainchainID_, ChainRegistrar chainRegistrar_)
        TokenBank(mainchainID_, chainRegistrar_)
    {}

    function lockTokens(
        uint256 targetChainID,
        address sourceChainTNT721Contract,
        address targetChainVoucherReceiver,
        uint256 tokenID
    ) external payable nonReentrant {
        require(
            _lockTokensSanityChecks(
                sourceChainTNT721Contract,
                targetChainID,
                targetChainVoucherReceiver
            ),
            "lock tokens sanity checks failed"
        );

        ITNT721 tnt721 = ITNT721(sourceChainTNT721Contract);
        {
            tnt721.transferFrom(msg.sender, address(this), tokenID);
            address ownerAfter = tnt721.ownerOf(tokenID);
            require(
                ownerAfter == address(this),
                "Failed to transfer TNT721 tokens to the token bank"
            ); // make sure the transfer succeeded, as the token contract could be malicious
        }

        string memory canonicalDenom = _buildDenom(
            block.chainid,
            "721",
            Strings.toHexString(uint256(uint160(sourceChainTNT721Contract)), 20)
        );

        // require(
        //     totalLockedAmounts[targetChainID][sourceChainTNT721Contract][
        //         tokenID
        //     ] == 0,
        //     "this tokenID has already been locked"
        // );

        // totalLockedAmounts[targetChainID][sourceChainTNT721Contract][
        //     tokenID
        // ] = totalLockedAmounts[targetChainID][sourceChainTNT721Contract][
        //     tokenID
        // ].add(1);

        _incrementTokenLockNonce(targetChainID);
        uint256 tokenLockNonce = tokenLockNonceMap[targetChainID];

        if (TokenBankUtils.supportsTNT721Metadata(sourceChainTNT721Contract)) {
            // name(), symbol(), and tokenURI() are only supported if the TNT721 contract implements the (optional) TNT721Metadata interface
            // Referernce: https://eips.ethereum.org/EIPS/eip-721
            emit TNT721TokenLocked(
                canonicalDenom,
                msg.sender,
                targetChainID,
                targetChainVoucherReceiver,
                tokenID,
                tnt721.name(),
                tnt721.symbol(),
                tnt721.tokenURI(tokenID),
                tokenLockNonce
            );
        } else {
            emit TNT721TokenLocked(
                canonicalDenom,
                msg.sender,
                targetChainID,
                targetChainVoucherReceiver,
                tokenID,
                "TNT721",
                "TNT721",
                "",
                tokenLockNonce
            );
        }

        // fee split for the validators
        require(
            msg.value >= chainRegistrar.getCrossChainFee(),
            "not enough TFuel (msg.value) to cover the cross-chain transaction fee"
        );
        _splitFeeAmongSubchainValidators(targetChainID);
    }

    function mintVouchers(
        string memory denom,
        string memory name,
        string memory symbol,
        address targetChainVoucherReceiver,
        uint256 tokenID,
        string memory tokenUri,
        uint256 dynasty,
        uint256 sourceChainTokenLockNonce
    ) external nonReentrant {
        require(
            targetChainVoucherReceiver != address(0x0),
            "cannot send to zero address"
        );
        string memory canonicalDenom = TokenBankUtils.canonicalizeDenom(denom);
        (uint256 sourceChainID, bool extractionSuccess) = TokenBankUtils
            .extractChainIDFromDenom(canonicalDenom);
        require(extractionSuccess, "Failed to extract chainID from denom");
        bytes32 tokenLockDataDigest = keccak256(
            abi.encodePacked(
                sourceChainID,
                canonicalDenom,
                name,
                symbol,
                targetChainVoucherReceiver,
                tokenID,
                tokenUri,
                dynasty,
                sourceChainTokenLockNonce
            )
        );
        bool voteSuccess = _checkValidatorQuorumForTokenLock(
            sourceChainID,
            dynasty,
            tokenLockDataDigest,
            sourceChainTokenLockNonce,
            msg.sender
        );
        if (voteSuccess) {
            if (!this.exists(canonicalDenom)) {
                addVoucher(
                    canonicalDenom,
                    _deployTNT721VoucherContract(canonicalDenom, name, symbol)
                );
            }
            address voucherContractAddr = this.getVoucher(canonicalDenom);
            require(
                voucherContractAddr != address(0),
                "the voucher for the denom does not exist"
            );
            _mintTNT721Vouchers(
                voucherContractAddr,
                targetChainVoucherReceiver,
                tokenID,
                tokenUri
            );
            _incrementVoucherMintNonce(sourceChainID);

            uint256 voucherMintNonce = voucherMintNonceMap[sourceChainID];
            emit TNT721VoucherMinted(
                canonicalDenom,
                targetChainVoucherReceiver,
                voucherContractAddr,
                tokenID,
                sourceChainTokenLockNonce,
                voucherMintNonce
            );
        }
    }

    function burnVouchers(
        address sourceChainVoucherContractAddr,
        address targetChainTokenReceiver,
        uint256 tokenID
    ) external payable nonReentrant {
        require(
            targetChainTokenReceiver != address(0x0),
            "cannot send to zero address"
        );
        string memory canonicalDenom = this.getDenom(
            sourceChainVoucherContractAddr
        );
        require(
            bytes(canonicalDenom).length != 0,
            "failed to lookup denom for the voucher contract"
        );
        require(this.exists(canonicalDenom), "Voucher contract does not exist");
        address sourceChainVoucherOwner = msg.sender;
        _burnTNT721Vouchers(
            sourceChainVoucherContractAddr,
            sourceChainVoucherOwner,
            tokenID
        );
        (uint256 targetChainID, bool extractionSuccess) = TokenBankUtils
            .extractChainIDFromDenom(canonicalDenom); // for voucher burn, target chain is always the chain where the authentic token was deployed
        require(
            extractionSuccess,
            "Failed to extract targetChainID from denom"
        );
        _incrementVoucherBurnNonce(targetChainID);

        uint256 voucherBurnNonce = voucherBurnNonceMap[targetChainID];
        emit TNT721VoucherBurned(
            canonicalDenom,
            sourceChainVoucherOwner,
            targetChainTokenReceiver,
            tokenID,
            voucherBurnNonce
        );

        // fee split for the validators
        require(
            msg.value >= chainRegistrar.getCrossChainFee(),
            "not enough TFuel (msg.value) to cover the cross-chain transaction fee"
        );
        _splitFeeAmongSubchainValidators(targetChainID);
    }

    function unlockTokens(
        uint256 sourceChainID,
        string memory denom,
        address targetChainTokenReceiver,
        uint256 tokenID,
        uint256 dynasty,
        uint256 sourceChainVoucherBurnNonce
    ) external nonReentrant {
        require(
            targetChainTokenReceiver != address(0x0),
            "cannot send to zero address"
        );
        string memory canonicalDenom = TokenBankUtils.canonicalizeDenom(denom);
        (
            address targetChainTNT721Contract,
            bool extractionSuccess
        ) = TokenBankUtils.extractContractAddressFromDenom(canonicalDenom);
        require(
            extractionSuccess,
            "Failed to extract contract address from denom"
        );
        // require(
        //     totalLockedAmounts[sourceChainID][targetChainTNT721Contract][
        //         tokenID
        //     ] == 1,
        //     "Cannot unlock the requested tokenID"
        // );

        bytes32 voucherBurnDataDigest = keccak256(
            abi.encodePacked(
                sourceChainID,
                canonicalDenom,
                tokenID,
                targetChainTokenReceiver,
                dynasty,
                sourceChainVoucherBurnNonce
            )
        );
        bool voteSuccess = _checkValidatorQuorumForVoucherBurn(
            sourceChainID,
            dynasty,
            voucherBurnDataDigest,
            sourceChainVoucherBurnNonce,
            msg.sender
        );
        if (voteSuccess) {
            // totalLockedAmounts[sourceChainID][targetChainTNT721Contract][
            //     tokenID
            // ] = totalLockedAmounts[sourceChainID][targetChainTNT721Contract][
            //     tokenID
            // ].sub(1);

            // using try-catch to prevent the external contract call reverts
            try
                ITNT721(targetChainTNT721Contract).transferFrom(
                    address(this),
                    targetChainTokenReceiver,
                    tokenID
                )
            {} catch {
                emit FailedToUnlockTNT721Tokens(
                    targetChainTNT721Contract,
                    targetChainTokenReceiver,
                    tokenID
                );
                // address fallbackReceiver = chainRegistrar.getFallbackReceiver();
                // ITNT721(targetChainTNT721Contract).transferFrom(
                //     address(this),
                //     fallbackReceiver,
                //     tokenID
                // );
                // emit UnlockTokensToFallbackReceiver(
                //     fallbackReceiver,
                //     targetChainTNT721Contract,
                //     tokenID
                // );
            }

            _incrementTokenUnlockNonce(sourceChainID);

            uint256 tokenUnlockNonce = tokenUnlockNonceMap[sourceChainID];
            emit TNT721TokenUnlocked(
                canonicalDenom,
                targetChainTokenReceiver,
                tokenID,
                sourceChainVoucherBurnNonce,
                tokenUnlockNonce
            );
        }
    }

    function _deployTNT721VoucherContract(
        string memory denom,
        string memory name,
        string memory symbol
    ) private returns (address) {
        address contractOwner = address(this);
        string memory voucherName = string(abi.encodePacked(name, " Voucher"));
        string memory voucherSymbol = string(abi.encodePacked("v", symbol));
        TNT721VoucherContract tnt721Voucher = new TNT721VoucherContract(
            contractOwner,
            denom,
            voucherName,
            voucherSymbol
        );
        return address(tnt721Voucher);
    }

    function _mintTNT721Vouchers(
        address voucherContractAddress,
        address voucherReceiver,
        uint256 tokenID,
        string memory tokenUri
    ) private {
        TNT721VoucherContract tnt721Voucher = TNT721VoucherContract(
            voucherContractAddress
        );

        // using try-catch to prevent the external contract call reverts
        try tnt721Voucher.mint(voucherReceiver, tokenID, tokenUri) {} catch {
            emit FailedToMintTNT721Vouchers(
                address(tnt721Voucher),
                voucherReceiver,
                tokenID
            );
            // address fallbackReceiver = chainRegistrar.getFallbackReceiver();
            // tnt721Voucher.mint(fallbackReceiver, tokenID, tokenUri);
            // emit MintVouchersToFallbackReceiver(
            //     fallbackReceiver,
            //     address(tnt721Voucher),
            //     tokenID
            // );
        }
    }

    function _burnTNT721Vouchers(
        address voucherContractAddress,
        address voucherOwner,
        uint256 tokenID
    ) private {
        TNT721VoucherContract tnt721Voucher = TNT721VoucherContract(
            voucherContractAddress
        );
        tnt721Voucher.burn(voucherOwner, tokenID);
    }
}

// TNT1155 Token Bank
contract TNT1155TokenBank is TokenBank, VoucherMap, ERC1155Holder {
    using SafeMath for uint256;

    event TNT1155TokenLocked(
        string denom,
        address sourceChainTokenSender,
        uint256 targetChainID,
        address targetChainVoucherReceiver,
        uint256 tokenID,
        uint256 lockedAmount,
        string tokenURI,
        uint256 tokenLockNonce
    );

    event TNT1155VoucherMinted(
        string denom,
        address targetChainVoucherReceiver,
        address voucherContract,
        uint256 tokenID,
        uint256 mintedAmount,
        uint256 sourceChainTokenLockNonce,
        uint256 voucherMintNonce
    );

    event TNT1155VoucherBurned(
        string denom,
        address sourceChainVoucherOwner,
        address targetChainTokenReceiver,
        uint256 tokenID,
        uint256 burnedAmount,
        uint256 voucherBurnNonce
    );

    event TNT1155TokenUnlocked(
        string denom,
        address targetChainTokenReceiver,
        uint256 tokenID,
        uint256 unlockedAmount,
        uint256 sourceChainVoucherBurnNonce,
        uint256 tokenUnlockNonce
    );

    event FailedToUnlockTNT1155Tokens(
        address indexed tokenContractAddress,
        address indexed receiver,
        uint256 tokenID,
        uint256 amount
    );

    event FailedToMintTNT1155Vouchers(
        address indexed voucherContractAddress,
        address indexed receiver,
        uint256 tokenID,
        uint256 amount
    );

    // event UnlockTokensToFallbackReceiver(
    //     address fallbackReceiver,
    //     address tokenContractAddress,
    //     uint256 tokenID,
    //     uint256 unlockAmount
    // );

    // event MintVouchersToFallbackReceiver(
    //     address fallbackReceiver,
    //     address voucherContractAddress,
    //     uint256 tokenID,
    //     uint256 mintAmount
    // );

    // Note: do not record the totalLockedAmounts for TNT tokens, since these tokens may
    //       have elastic supplies (e.g. AMPL) which violates the conservation rules. This
    //       cause the unlockTokens() to revert and stall the cross-chain transfer pipeline
    // // {chainID : {tnt1155Contract : {tokenID : totalLockedAmount}}}
    // mapping(uint256 => mapping(address => mapping(uint256 => uint256)))
    //     public totalLockedAmounts;

    constructor(uint256 mainchainID_, ChainRegistrar chainRegistrar_)
        TokenBank(mainchainID_, chainRegistrar_)
    {}

    function lockTokens(
        uint256 targetChainID,
        address sourceChainTNT1155Contract,
        address targetChainVoucherReceiver,
        uint256 tokenID,
        uint256 lockAmount
    ) external payable nonReentrant {
        require(
            _lockTokensSanityChecks(
                sourceChainTNT1155Contract,
                targetChainID,
                targetChainVoucherReceiver
            ),
            "lock tokens sanity checks failed"
        );
        ITNT1155 tnt1155 = ITNT1155(sourceChainTNT1155Contract);

        {
            uint256 balanceBefore = tnt1155.balanceOf(address(this), tokenID);
            tnt1155.safeTransferFrom(
                msg.sender,
                address(this),
                tokenID,
                lockAmount,
                bytes("")
            );
            uint256 balanceAfter = tnt1155.balanceOf(address(this), tokenID);
            require(
                balanceAfter.sub(lockAmount) == balanceBefore,
                "Failed to transfer TNT1155 tokens to the token bank"
            ); // make sure the transfer succeeded, as the token contract could be malicious
        }
        string memory canonicalDenom = _buildDenom(
            block.chainid,
            "1155",
            Strings.toHexString(
                uint256(uint160(sourceChainTNT1155Contract)),
                20
            )
        );

        // _updateTotalLockedAmount(
        //     targetChainID,
        //     sourceChainTNT1155Contract,
        //     tokenID,
        //     lockAmount,
        //     true
        // );

        _incrementTokenLockNonce(targetChainID);
        uint256 tokenLockNonce = tokenLockNonceMap[targetChainID];

        if (
            TokenBankUtils.supportsTNT1155MetadataURI(
                sourceChainTNT1155Contract,
                tokenID
            )
        ) {
            emit TNT1155TokenLocked(
                canonicalDenom,
                msg.sender,
                targetChainID,
                targetChainVoucherReceiver,
                tokenID,
                lockAmount,
                tnt1155.uri(tokenID),
                tokenLockNonce
            );
        } else {
            emit TNT1155TokenLocked(
                canonicalDenom,
                msg.sender,
                targetChainID,
                targetChainVoucherReceiver,
                tokenID,
                lockAmount,
                "",
                tokenLockNonce
            );
        }

        // fee split for the validators
        require(
            msg.value >= chainRegistrar.getCrossChainFee(),
            "not enough TFuel (msg.value) to cover the cross-chain transaction fee"
        );
        _splitFeeAmongSubchainValidators(targetChainID);
    }

    function mintVouchers(
        string memory denom,
        address targetChainVoucherReceiver,
        uint256 tokenID,
        uint256 mintAmount,
        string memory tokenUri,
        uint256 dynasty,
        uint256 sourceChainTokenLockNonce
    ) external nonReentrant {
        require(
            targetChainVoucherReceiver != address(0x0),
            "cannot send to zero address"
        );
        string memory canonicalDenom = TokenBankUtils.canonicalizeDenom(denom);
        (uint256 sourceChainID, bool extractionSuccess) = TokenBankUtils
            .extractChainIDFromDenom(canonicalDenom);
        require(extractionSuccess, "Failed to extract chainID from denom");
        bytes32 tokenMintDataDigest = keccak256(
            abi.encodePacked(
                sourceChainID,
                canonicalDenom,
                targetChainVoucherReceiver,
                tokenID,
                mintAmount,
                dynasty,
                sourceChainTokenLockNonce
            )
        );
        bool voteSuccess = _checkValidatorQuorumForTokenLock(
            sourceChainID,
            dynasty,
            tokenMintDataDigest,
            sourceChainTokenLockNonce,
            msg.sender
        );
        if (voteSuccess) {
            if (!this.exists(canonicalDenom)) {
                addVoucher(
                    canonicalDenom,
                    _deployTNT1155VoucherContract(canonicalDenom)
                );
            }
            address voucherContractAddr = this.getVoucher(canonicalDenom);
            require(
                voucherContractAddr != address(0),
                "the voucher for the denom does not exist"
            );
            _mintTNT1155Vouchers(
                voucherContractAddr,
                targetChainVoucherReceiver,
                tokenID,
                mintAmount,
                tokenUri
            );
            _incrementVoucherMintNonce(sourceChainID);

            uint256 voucherMintNonce = voucherMintNonceMap[sourceChainID];
            emit TNT1155VoucherMinted(
                canonicalDenom,
                targetChainVoucherReceiver,
                voucherContractAddr,
                tokenID,
                mintAmount,
                sourceChainTokenLockNonce,
                voucherMintNonce
            );
        }
    }

    function burnVouchers(
        address sourceChainVoucherContractAddr,
        address targetChainTokenReceiver,
        uint256 tokenID,
        uint256 burnAmount
    ) external payable nonReentrant {
        require(
            targetChainTokenReceiver != address(0x0),
            "cannot send to zero address"
        );
        string memory canonicalDenom = this.getDenom(
            sourceChainVoucherContractAddr
        );
        require(
            bytes(canonicalDenom).length != 0,
            "failed to lookup denom for the voucher contract"
        );
        require(this.exists(canonicalDenom), "Voucher contract does not exist");
        address sourceChainVoucherOwner = msg.sender;
        _burnTNT1155Vouchers(
            sourceChainVoucherContractAddr,
            sourceChainVoucherOwner,
            tokenID,
            burnAmount
        );
        (uint256 targetChainID, bool extractionSuccess) = TokenBankUtils
            .extractChainIDFromDenom(canonicalDenom); // for voucher burn, target chain is always the chain where the authentic token was deployed
        require(
            extractionSuccess,
            "Failed to extract targetChainID from denom"
        );
        _incrementVoucherBurnNonce(targetChainID);

        uint256 voucherBurnNonce = voucherBurnNonceMap[targetChainID];
        emit TNT1155VoucherBurned(
            canonicalDenom,
            sourceChainVoucherOwner,
            targetChainTokenReceiver,
            tokenID,
            burnAmount,
            voucherBurnNonce
        );

        // fee split for the validators
        require(
            msg.value >= chainRegistrar.getCrossChainFee(),
            "not enough TFuel (msg.value) to cover the cross-chain transaction fee"
        );
        _splitFeeAmongSubchainValidators(targetChainID);
    }

    function unlockTokens(
        uint256 sourceChainID,
        string memory denom,
        address targetChainTokenReceiver,
        uint256 tokenID,
        uint256 unlockAmount,
        uint256 dynasty,
        uint256 sourceChainVoucherBurnNonce
    ) external payable nonReentrant {
        require(
            targetChainTokenReceiver != address(0x0),
            "cannot send to zero address"
        );
        string memory canonicalDenom = TokenBankUtils.canonicalizeDenom(denom);
        (
            address targetChainTNT1155Contract,
            bool extractionSuccess
        ) = TokenBankUtils.extractContractAddressFromDenom(canonicalDenom);
        require(
            extractionSuccess,
            "Failed to extract contract address from denom"
        );
        // require(
        //     unlockAmount <=
        //         totalLockedAmounts[sourceChainID][targetChainTNT1155Contract][
        //             tokenID
        //         ],
        //     "Cannot unlock the requested amount of tokens"
        // );

        bytes32 voucherBurnDataDigest = keccak256(
            abi.encodePacked(
                sourceChainID,
                canonicalDenom,
                tokenID,
                unlockAmount,
                targetChainTokenReceiver,
                dynasty,
                sourceChainVoucherBurnNonce
            )
        );
        bool voteSuccess = _checkValidatorQuorumForVoucherBurn(
            sourceChainID,
            dynasty,
            voucherBurnDataDigest,
            sourceChainVoucherBurnNonce,
            msg.sender
        );
        if (voteSuccess) {
            // _updateTotalLockedAmount(
            //     sourceChainID,
            //     targetChainTNT1155Contract,
            //     tokenID,
            //     unlockAmount,
            //     false
            // );

            // using try-catch to prevent the external contract call reverts
            try
                ITNT1155(targetChainTNT1155Contract).safeTransferFrom(
                    address(this),
                    targetChainTokenReceiver,
                    tokenID,
                    unlockAmount,
                    bytes("")
                )
            {} catch {
                emit FailedToUnlockTNT1155Tokens(
                    targetChainTNT1155Contract,
                    targetChainTokenReceiver,
                    tokenID,
                    unlockAmount
                );
                // address fallbackReceiver = chainRegistrar.getFallbackReceiver();
                // ITNT1155(targetChainTNT1155Contract).safeTransferFrom(
                //     address(this),
                //     fallbackReceiver,
                //     tokenID,
                //     unlockAmount,
                //     bytes("")
                // );
                // emit UnlockTokensToFallbackReceiver(
                //     fallbackReceiver,
                //     targetChainTNT1155Contract,
                //     tokenID,
                //     unlockAmount
                // );
            }

            _incrementTokenUnlockNonce(sourceChainID);

            uint256 tokenUnlockNonce = tokenUnlockNonceMap[sourceChainID];
            emit TNT1155TokenUnlocked(
                canonicalDenom,
                targetChainTokenReceiver,
                tokenID,
                unlockAmount,
                sourceChainVoucherBurnNonce,
                tokenUnlockNonce
            );
        }
    }

    // function _updateTotalLockedAmount(
    //     uint256 chainID,
    //     address contractAddr,
    //     uint256 tokenID,
    //     uint256 amount,
    //     bool isLockOperation
    // ) private {
    //     if (isLockOperation) {
    //         totalLockedAmounts[chainID][contractAddr][
    //             tokenID
    //         ] = totalLockedAmounts[chainID][contractAddr][tokenID].add(amount);
    //     } else {
    //         totalLockedAmounts[chainID][contractAddr][
    //             tokenID
    //         ] = totalLockedAmounts[chainID][contractAddr][tokenID].sub(amount);
    //     }
    // }

    function _deployTNT1155VoucherContract(string memory denom)
        private
        returns (address)
    {
        address contractOwner = address(this);
        TNT1155VoucherContract tnt1155Voucher = new TNT1155VoucherContract(
            contractOwner,
            denom
        );
        return address(tnt1155Voucher);
    }

    function _mintTNT1155Vouchers(
        address voucherContractAddress,
        address voucherReceiver,
        uint256 tokenID,
        uint256 mintAmount,
        string memory tokenUri
    ) private {
        TNT1155VoucherContract tnt1155Voucher = TNT1155VoucherContract(
            voucherContractAddress
        );

        // using try-catch to prevent the external contract call reverts
        try
            tnt1155Voucher.mint(voucherReceiver, tokenID, mintAmount, tokenUri)
        {} catch {
            emit FailedToMintTNT1155Vouchers(
                voucherContractAddress,
                voucherReceiver,
                tokenID,
                mintAmount
            );
            // address fallbackReceiver = chainRegistrar.getFallbackReceiver();
            // tnt1155Voucher.mint(
            //     fallbackReceiver, // mint to the fallback address if mint fails, e.g. if voucherReceiver is a malicious contract
            //     tokenID,
            //     mintAmount,
            //     tokenUri
            // );
            // emit MintVouchersToFallbackReceiver(
            //     fallbackReceiver,
            //     address(tnt1155Voucher),
            //     tokenID,
            //     mintAmount
            // );
        }
    }

    function _burnTNT1155Vouchers(
        address voucherContractAddress,
        address voucherOwner,
        uint256 tokenID,
        uint256 burnAmount
    ) private {
        TNT1155VoucherContract tnt1155Voucher = TNT1155VoucherContract(
            voucherContractAddress
        );
        tnt1155Voucher.burn(voucherOwner, tokenID, burnAmount);
    }
}
