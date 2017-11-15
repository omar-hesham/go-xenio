pragma solidity ^0.4.4;

// IMPORTS
import "./EternalStorage.sol";

// USER STRUCTURE
// keccak256("user_name", idx) -> bytes32 -- User name.
// keccak256("user_address", idx) -> address -- User Xenio address.
// keccak256("user_registered", idx) -> bool -- If false the user is considered to be deregistered, or deleted.
// keccak256("user_game_token", idx, gidx) -> bytes32 -- Game Token ID (unique key)
// keccak256("user_game_tokens_count", idx) -> uint -- Total count of game tokens.
// keccak256("user_img_profile", idx) -> bytes32 -- The image url for profile picture of the user.
// keccak256("users_count") -> uint256 -- Total count of registered users.

// Note that idx and gidx should always be uint

/** @title Library for Users. */

/** Library for Users. */
library UsersLib {

    // EVENTS
    event UserAdded(uint _id);
    event UserRemoved(uint _id);        

    // MAIN METHODS

    /// @dev registers a new user
    function addUser
    (
        address _storageContract, 
        address _address, 
        bytes32 _name
    ) 
        public 
        returns(bool success)
    {
        // Validation check
        require(isNameAvailable(_storageContract,_name)); // check whether the given name is taken
        require(isAddressAvailable(_storageContract,_address)); // check whether the given address is taken

        // Get first available id
        uint idx = getNextID(_storageContract,0); 

        // Register user's details
        EternalStorage(_storageContract).setBytes32Value(keccak256("user_name", idx), _name);
        EternalStorage(_storageContract).setAddressValue(keccak256("user_address", idx), _address);
        EternalStorage(_storageContract).setBooleanValue(keccak256("user_registered", idx), true);

        // Update users count
        if (idx + 1 > getUsersCount(_storageContract))
            EternalStorage(_storageContract).setUIntValue(keccak256("users_count"), idx + 1);        
        
        UserAdded(idx); // call event

        success = true;        
    }

    /// @dev deregisters a user, by changing its registration state to false
    function removeUser(address _storageContract, uint idx) public {
        EternalStorage(_storageContract).setBooleanValue(keccak256("user_registered", idx), false);
        UserRemoved(idx); // call event
    }    

    // UPDATE METHODS    

   /// @dev updates the name of the user
    function updateName(address _storageContract, address _address, bytes32 _name) public returns (bool) {
        var (found, userID) = getIDByAddress(_storageContract,_address);
        if (found && isNameAvailable(_storageContract,_name)) {
            EternalStorage(_storageContract).setBytes32Value(keccak256("user_name", userID), _name);
        }
        return found;
    }        

    /// @dev updates the image url of the profile picture of the user
    function updateProfilePicture(address _storageContract, address _address, bytes32 _imageUrl) public returns (bool) {
        var (found, userID) = getIDByAddress(_storageContract,_address);
        if (found) {EternalStorage(_storageContract).setBytes32Value(keccak256("user_img_logo", userID), _imageUrl);}
        return found;
    }

    /// @dev updates the address of the user
    function updateAddress(address _storageContract, address _address, address _newAddress) public returns (bool) {
        var (found, userID) = getIDByAddress(_storageContract,_address);
        if (found) {EternalStorage(_storageContract).setAddressValue(keccak256("user_address", userID), _newAddress);}
        return found;
    }
    
    // GETTER METHODS
    
    // Variables

    function getUsersCount(address _storageContract) public constant returns(uint256) {
        return EternalStorage(_storageContract).getUIntValue(keccak256("users_count"));
    }

    /// @dev Returns the total number of active users in Xenio.
    function getActiveUsersCount(address _storageContract) public constant returns(uint activeUsersCount) {
        var totalCount = getUsersCount(_storageContract);
        activeUsersCount = 0;
        for (uint i = 0; i < totalCount; i++) {
            if (EternalStorage(_storageContract).getBooleanValue(keccak256("user_registered", i))) {activeUsersCount++;}
        }
    }

    /// @dev Gets next available id
    function getNextID(address _storageContract, uint _id) private constant returns(uint nextID) {
        nextID = _id; // initialise   
        while (EternalStorage(_storageContract).getBooleanValue(keccak256("user_registered", nextID))) {nextID++;}
    }    

    /// @dev Checks whether the user ID corresponds to an active user
    function isActive(address _storageContract, uint _id) public constant returns(bool) {
        return EternalStorage(_storageContract).getBooleanValue(keccak256("user_registered", _id));
    }        

    /// @dev Checks whether the given name is taken
    function isNameAvailable(address _storageContract, bytes32 _name) private constant returns (bool) {
        var totalCount = getUsersCount(_storageContract);
        for (uint i = 0; i < totalCount; i++) {
            if (EternalStorage(_storageContract).getBooleanValue(keccak256("user_registered", i))) {
                if (keccak256(EternalStorage(_storageContract).getBytes32Value(keccak256("user_name", i))) == keccak256(_name)) {
                    return false;
                }
            }
        } 
        return true;
    }

    /// @dev Checks whether the given address is taken
    function isAddressAvailable(address _storageContract, address _address) private constant returns (bool) {
        var totalCount = getUsersCount(_storageContract);
        for (uint i = 0; i < totalCount; i++) {
            if (EternalStorage(_storageContract).getBooleanValue(keccak256("user_registered", i))) {
                if (keccak256(EternalStorage(_storageContract).getAddressValue(keccak256("user_address", i))) == keccak256(_address)) {
                    return false;
                }
            }
        } 
        return true;
    }     

    /// @dev Returns user id given its address
    function getIDByAddress(address _storageContract, address _address) public constant returns (bool found, uint id) {
        var totalCount = getUsersCount(_storageContract);
        for (uint i = 0; i < totalCount; i++) {
            if (EternalStorage(_storageContract).getBooleanValue(keccak256("user_registered", i))) {
                if (keccak256(EternalStorage(_storageContract).getAddressValue(keccak256("user_address", i))) == keccak256(_address)) {  
                    return (true, i);
                }
            }
        }
        return (false,0);
    }

    /// @dev Returns user id given its name
    function getIDByName(address _storageContract, bytes32 _name) public constant returns (bool found, uint id) {
        var totalCount = getUsersCount(_storageContract);
        for (uint i = 0; i < totalCount; i++) {
            if (EternalStorage(_storageContract).getBooleanValue(keccak256("user_registered", i))) {
                if (keccak256(EternalStorage(_storageContract).getBytes32Value(keccak256("user_name", i))) == keccak256(_name)) {  
                    return (true, i);
                }
            }
        }
        return (false,0);
    }    

    // User details

   /// @dev Returns user's details given its user id.
    function getUserDetailsByID(address _storageContract, uint _userID)
        public
        constant
        returns
    (
        bool found,
        address userAddress,
        bytes32 userName
    )
    {
        found = EternalStorage(_storageContract).getBooleanValue(keccak256("user_registered", _userID));
        if (!found) {return;}        
        userAddress = EternalStorage(_storageContract).getAddressValue(keccak256("user_address", _userID));
        userName = EternalStorage(_storageContract).getBytes32Value(keccak256("user_name", _userID));
    }

    /// @dev Returns user's details given its name.
    function getUserDetailsByName(address _storageContract, bytes32 _name)
        public
        constant
        returns
    (
        bool found,
        uint userID,
        address userAddress       
    )
    {
        (found, userID) = getIDByName(_storageContract, _name); // get user id
        if (!found) {return;}
        userAddress = EternalStorage(_storageContract).getAddressValue(keccak256("user_address", userID));
    }


}