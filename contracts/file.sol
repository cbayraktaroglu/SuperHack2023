// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { IWorldID } from './interfaces/IWorldID.sol';
import { ByteHasher } from './helpers/ByteHasher.sol';

contract File {

    using ByteHasher for bytes;

    address internal owner;
    bool internal isVerified;
    address public factoryAddress;
    string internal fileType;
    string internal fileName;
    string internal fileCheckSum;
    uint256 internal numberOfTransaction;

	/// @dev The World ID group ID (always 1)
	uint256 internal immutable groupId = 1;

    struct FileParts {
        string chainID;
        string transactionHash;
    }

    modifier onlyOwner() {
        require(msg.sender == factoryAddress, "Only the factory can call this function");
        _;
    }

    mapping(uint256 => FileParts) internal transactionMap;

    event VerificationResult(address owner, bool result);
    event OwnershipTransferred(address indexed fileAddress, address indexed previousOwner, address indexed newOwner);


    constructor(address _owner, string memory _fileType, string memory _fileName, string memory _fileCheckSum,
    string[] memory _transactionList, string[] memory _chainIdList) {

        require(_transactionList.length == _chainIdList.length, "Number of transactions should match with the number of chain IDs");
        factoryAddress = msg.sender; // Set the factory address
        owner = _owner;
        isVerified = false; //Verification deualt is false
        fileType = _fileType;
        fileName = _fileName;
        fileCheckSum = _fileCheckSum;
        numberOfTransaction = _transactionList.length;

        for (uint256 i = 0; i < _transactionList.length; i++) {
            FileParts memory fileParts = FileParts(_chainIdList[i], _transactionList[i]);
            transactionMap[i] = fileParts;
        }

      
    }

    // Getter functions
    function getOwner() external view returns (address) {
        return owner;
    }

    function getFileType() external view returns (string memory) {
        return fileType;
    }

    function getFileName() external view returns (string memory) {
        return fileName;
    }

    function getFileCheckSum() external view returns (string memory) {
        return fileCheckSum;
    }

    function getNumberofTransactions() external view returns (uint256) {
        return numberOfTransaction;
    }

    function getTransactionAtIndex(uint256 index) external view returns (FileParts memory) {
        return transactionMap[index];
    }

    function getVerification() external view returns (bool) {
        return isVerified;
    }

    // Transfer ownership to a new address
    function transferOwnership(address newOwner) public onlyOwner{
        require(newOwner != address(0), "Invalid new owner address");
        owner = newOwner;
    }

    function checkVerification(IWorldID _worldId, string memory _appId, string memory _actionId, address signal, uint256 root, uint256 nullifierHash, uint256[8] calldata proof) public onlyOwner{
        require(isVerified == false, "Already verified");
        
        uint256 externalNullifier = abi.encodePacked(abi.encodePacked(_appId).hashToField(), _actionId).hashToField();

        // Check the user from World Coin system
        try _worldId.verifyProof(root, groupId, abi.encodePacked(signal).hashToField(), nullifierHash, externalNullifier, proof){
            isVerified = true;
        } catch {
            isVerified = false;
        }
        
        emit VerificationResult(owner, isVerified);   
    }
}