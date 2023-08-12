// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// Import the contract that you want to generate
import "./KeepItFile.sol";

contract KeepItFileFactory {
    // Addresses
    address private owner;
    address private attestationServiceAddress;

    // Events
    event FileCreated(address fileAddress);
    event OwnerSet(address indexed oldOwner, address indexed newOwner);

    // Mapping to store deployed contract addresses by creator
    mapping(address => address[]) internal fileCreators;

    // modifier to check if caller is owner
    modifier isOwner() {
        require(msg.sender == owner, "Caller is not owner");
        _;
    }

    /**
     * @dev Set contract deployer as owner
     */
    constructor() {
        owner = msg.sender; // 'msg.sender' is sender of current call, contract deployer for a constructor
        emit OwnerSet(address(0), owner);
    }

    /**
     * @dev Change owner
     * @param newOwner address of new owner
     */
    function changeOwner(address newOwner) public isOwner {
        emit OwnerSet(owner, newOwner);
        owner = newOwner;
    }

    function setAttestationServiceAddress(address _aSA) public isOwner {
        attestationServiceAddress = _aSA;
    }

    function getAttestationServiceAddres() public view returns (address) {
        return attestationServiceAddress;
    }

    /**
     * @dev Return owner address
     * @return address of owner
     */
    function getOwner() external view returns (address) {
        return owner;
    }

    /**
     * @dev Creates a KeepItFile contract that acts as a pointer to the file
     * @param _fileType type of the file (extension) e.g. "html"
     * @param _fileName name of the file
     * @param _fileCheckSum SHA256 checksum of the file
     * @param _transactionList ordered list of TXs that contain the file
     * @param _chainIdList  ordered list of chain ids corresponding to the _transactionList elems indicating where this tx has happened
     */
    function createFile(
        string memory _fileType,
        string memory _fileName,
        string memory _fileCheckSum,
        string[] memory _transactionList,
        string[] memory _chainIdList
    ) public {
        require(
            _transactionList.length == _chainIdList.length,
            "Number of transactions should match with the number of chain IDs"
        );

        KeepItFile file = new KeepItFile(
            msg.sender,
            _fileType,
            _fileName,
            _fileCheckSum,
            _transactionList,
            _chainIdList
        );
        address newFileAddress = address(file);

        // Update the mapping
        fileCreators[msg.sender].push(newFileAddress);

        emit FileCreated(address(newFileAddress));
    }

    /**
     * @dev Retririeves all the KeepItFile addresses created by the given user
     * @param _owner address of the user
     * @return address[] list of KeepItFile contract addresses
     */
    function getFiles(address _owner) public view returns (address[] memory) {
        return fileCreators[_owner];
    }
}
