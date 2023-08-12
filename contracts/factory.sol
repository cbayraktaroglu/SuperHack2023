// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// Import the contract that you want to generate
import "./file.sol";


contract FileFactory {

    event FileCreated(address fileAddress);

    // Mapping to store deployed contract addresses by creator
    mapping(address => address[]) internal fileCreators;

    function createFile(string memory _fileType, string memory _fileName, string memory _fileCheckSum, string[] memory _transactionList,string[] memory _chainIdList ) public {
        require(_transactionList.length == _chainIdList.length, "Number of transactions should match with the number of chain IDs");
       
        File file = new File(msg.sender, _fileType, _fileName, _fileCheckSum, _transactionList,_chainIdList );
        address newFileAddress = address(file);
    
        // Update the mapping
        fileCreators[msg.sender].push(newFileAddress);
 
        emit FileCreated(address(newFileAddress));
    }

    function getFiles(address _owner) public view returns(address[] memory){
        return fileCreators[_owner];
    }

}
