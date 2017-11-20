pragma solidity ^0.4.4;

// IMPORTS
import "./EternalStorage.sol";

/** @title Library for managing multiple admin accounts. */
// Manages records for admins stored in the format: 
// keccak256("admin", address) -> bool isUserAdmin , e.g. 0xd91cf6dac04d456edc5fcb6659dd8ddedbb26661 -> true.
// keccak256("admins_count")
library SecurityLib {

    event AdminAdded(address _user);
    event AdminRemoved(address _user);

    /** @dev Gets the total number of registered admins.
        @param _storageContract The contract address of eternal storage.
        @return The number of admin accounts registered.
    */
    function getAdminsCount(address _storageContract) public constant returns(uint256) {
        return EternalStorage(_storageContract).getUIntValue(keccak256("admin_counts"));
    }

     /** @dev Adds a new admin into the list.
        @param _storageContract The contract address of eternal storage.
        @param _user The address of the admin user to be added.
    */   
    function addAdmin(address _storageContract, address _user) public {
        
        // Validation check -- if admin exists, revert
        var userIsAdmin = EternalStorage(_storageContract).getBooleanValue(keccak256("admin", _user));
        require(!userIsAdmin); 

        // Insert the new admin into storage
        EternalStorage(_storageContract).setBooleanValue(keccak256("admin", _user), true); 

        // Increment the admins count in storage
        var adminsCount = EternalStorage(_storageContract).getUIntValue(keccak256("admin_counts"));
        EternalStorage(_storageContract).setUIntValue(keccak256("admin_counts"), adminsCount + 1);

        AdminAdded(_user); // event call
    }

     /** @dev Removes an admin user from the list.
        @param _storageContract The contract address of eternal storage.
        @param _user The address of the admin user to be removed.
    */       
    function removeAdmin(address _storageContract, address _user) public {

        // Validation check -- admin should exist
        var userIsAdmin = EternalStorage(_storageContract).getBooleanValue(keccak256("admin", _user));
        require(userIsAdmin);

        var adminsCount = EternalStorage(_storageContract).getUIntValue(keccak256("admin_counts"));
        require(adminsCount != 1);

        // Remove admin from storage
        EternalStorage(_storageContract).setBooleanValue(keccak256("admin", _user), false); 

        // Decrement the admins count in storage
        adminsCount -= 1;
        EternalStorage(_storageContract).setUIntValue(keccak256("admin_counts"), adminsCount);

        AdminRemoved(_user); // event call
    }

     /** @dev Checks whether the user has administrative authority.
        @param _storageContract The contract address of eternal storage.
        @param _user The address of the user to be assessed.
        @return True or false whether the user is registered as an admin or not.
    */     
    function isUserAdmin(address _storageContract, address _user) public constant returns (bool) {
        return EternalStorage(_storageContract).getBooleanValue(keccak256("admin", _user));
    }
}