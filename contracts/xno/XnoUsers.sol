pragma solidity ^0.4.4;

// IMPORTS
import "./Ownable.sol";
import "./UsersLib.sol";
import "./SecurityLib.sol";
import "./EternalStorage.sol";

/// @title Main contract for Xenio users management
contract XnoUsers is Ownable {

    using UsersLib for address;
    using SecurityLib for address;
    address public eternalStorage;

    // EVENTS

    // Storage
    event StorageCreated(address _storageAddress);
    event StorageAttached(address _storageAddress);
    event StorageDeleted();
    event StorageDetached(address _newOwner);

    // Contract
    event XnoUsersCreated(address _owner);
    event XnoUsersDeactivated();

    // MODIFIERS
    modifier storageAttached(bool state) {
        require((eternalStorage != address(0)) == state);
        _;
    }

    modifier onlyAdmin() {
        require(eternalStorage.isUserAdmin(msg.sender));
        _;
    }

    // STORAGE METHODS

    /// @dev get the address of the owner of the attached eternal storage.
    function getStorageOwner() 
        external 
        storageAttached(true) 
        constant 
        returns(address) 
    {
        return EternalStorage(eternalStorage).getOwner();
    }

    // ONLY OWNER: Only the owner of this contract can call the following methods.

    /// @dev creates and attaches a new eternal storage. Not attached to this contract.
    function createNewStorage() external storageAttached(false) onlyOwner {
        eternalStorage = new EternalStorage(); // Note that it will change storage attachment 
        eternalStorage.addAdmin(getOwner()); // Add the owner of the current contract to the admin list

        StorageCreated(eternalStorage);
    }

    /// @dev attach this contract to an existing eternal storage 
    function attachExistingStorage(address _storageContract) external onlyOwner {
        require(EternalStorage(_storageContract).getOwner() == address(this)); // require storage ownership first! this contract address should be the owner.
        eternalStorage = _storageContract; // attach
        eternalStorage.addAdmin(getOwner()); // Add the owner of the current contract to the admin list

        StorageAttached(eternalStorage);
    }

    /// @dev change the ownership of the currently attached eternal storage. used only for upgrading XnoUsers contract.
    function changeStorageOwner(address _newOwner) private storageAttached(true) onlyOwner {
        EternalStorage(eternalStorage).passOwnership(_newOwner);
    }  

    /// @dev deactivates eternal storage permantly! 
    function deleteStorage() external storageAttached(true) onlyOwner {
        EternalStorage(eternalStorage).kill();
        eternalStorage = address(0);        

        StorageDeleted(); 
    }

    /// @dev detaches eternal storage! 
    function detachStorage(address _newOnwer) external storageAttached(true) onlyOwner {
        changeStorageOwner(_newOnwer); // pass ownership of the attached storage first!
        eternalStorage = address(0);

        StorageDetached(_newOnwer);
    }

    /// @dev deactivates the current contract
    function deactivate() external storageAttached(false) onlyOwner {
        kill();
        XnoUsersDeactivated();
    }

    // onlyAdmin METHODS
    // Only the admins, which are registered in the eternal storage attached to this contract, can call the following methods.

    /// @dev removes a Xenio user given the address of the user. Only admins are allowed to call this function.
    function removeUserByAddress(address _address)
        external 
        storageAttached(true) 
        onlyAdmin 
        returns (bool success) 
    {
        var (found, userID) = eternalStorage.getIDByAddress(_address);
        if (found) {eternalStorage.removeUser(userID);} // remove user
        return found;
    }

    /// @dev removes a Xenio user give the name of the user. Only admins are allowed to call this function.
    function removeUserByName(bytes32 _name)
        external 
        storageAttached(true) 
        onlyAdmin 
        returns (bool success) 
    {
        var (found, userID) = eternalStorage.getIDByName(_name);
        if (found) {eternalStorage.removeUser(userID);} // remove user
        return found;
    }

    // Update methods for admins

    /// @dev updates the name of a user given the address of the user. Only admins are allowed to call this function.
    function changeUserName(address _address, bytes32 _name) external storageAttached(true) onlyAdmin returns(bool success) {
        success = eternalStorage.updateName(_address, _name);
    }

    /// @dev updates the image url of the profile picture of the user, given the address of the user. Only admins are allowed to call this function.
    function changeUserProfilePicture(address _address, bytes32 _imgUrl) external storageAttached(true) onlyAdmin returns(bool success) {
        success = eternalStorage.updateProfilePicture(_address, _imgUrl);
    } 

    /// @dev updates the address of a user given current address of the user. Only admins are allowed to call this function.
    function changeUserAddress(address _currentAddress, address _newAddress) external storageAttached(true) onlyAdmin returns(bool success) {
        success = eternalStorage.updateAddress(_currentAddress, _newAddress);
    }


    // MAIN METHODS

    /// @dev Constructor
    function XnoUsers() public {
        XnoUsersCreated(getOwner());
    }

    /// @dev registers a new Xenio user 
    function registerNewUser(bytes32 _name)  
        public 
        storageAttached(true)
        returns (bool success)
    {
        return eternalStorage.addUser(msg.sender, _name);
    }
    
    /// @dev deregisters the respective Xenio user (if exist) of the sender's address
    function deregisterUser() public storageAttached(true) {
        var (found, id) = eternalStorage.getIDByAddress(msg.sender); // get user id
        if (found) {eternalStorage.removeUser(id);} // remove user
    }

    // Update methods for senders

    /// @dev updates the name of the user owned by the sender.
    function updateUserName(bytes32 _name) external storageAttached(true) returns(bool success) {
        success = eternalStorage.updateName(msg.sender, _name);
    }

    /// @dev updates the image url of the profile picture of the user owned by the sender.
    function updateUserProfilePicture(bytes32 _imgUrl) external storageAttached(true) returns(bool success) {
        success = eternalStorage.updateProfilePicture(msg.sender, _imgUrl);
    }

    // GETTER METHODS

    /// @dev Returns the total number of active Xenio users.
    function getUsersCount() external storageAttached(true) constant returns(uint) {
        return eternalStorage.getActiveUsersCount();
    }

    /// @dev Returns the user id of a given address (unique)
    function getUserIDByAddress(address _address) external storageAttached(true) constant returns(bool found, uint userID) {
        return eternalStorage.getIDByAddress(_address);
    }

    /// @dev Returns the user id of a given name (unique)
    function getUserIDByName(bytes32 _name) external storageAttached(true) constant returns(bool found, uint userID) {
        return eternalStorage.getIDByName(_name);
    }

    /// @dev Returns user details by its id
    function getUserByID(uint _userID) 
        external 
        storageAttached(true) 
        constant 
        returns
    (
        bool found,
        address userAddress,
        bytes32 userName  
    ) 
    {
        return eternalStorage.getUserDetailsByID(_userID);
    }

    /// @dev Returns user details by its name
    function getUserByName(bytes32 _name) 
        external 
        storageAttached(true) 
        constant 
        returns
    (
        bool found,
        uint userID,
        address userAddress
    ) 
    {
        return eternalStorage.getUserDetailsByName(_name);
    }

    // @dev Returns a list with all active userIDs.
    function getAllUserIDs() external storageAttached(true) constant returns(uint[] memory list) {
        uint counter = 0;
        list = new uint[](eternalStorage.getActiveUsersCount()); // initialise
        for (uint i = 0; i < eternalStorage.getUsersCount(); i++) {
            if (eternalStorage.isActive(i))
                list[counter++] = i;
        }
    }

}