pragma solidity ^0.4.4;

// IMPORTS
import "./Ownable.sol";
import "./ServersLib.sol";
import "./SecurityLib.sol";
import "./EternalStorage.sol";

/// @title Main contract for Xenio servers management
contract XnoServers is Ownable {

    using ServersLib for address;
    using SecurityLib for address;
    address public eternalStorage;

    // EVENTS

    // Storage
    event StorageCreated(address _storageAddress);
    event StorageAttached(address _storageAddress);
    event StorageDeleted();
    event StorageDetached(address _newOwner);

    // Contract
    event XnoServersCreated(address _owner);
    event XnoServersDeactivated();

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

    /// @dev change the ownership of the currently attached eternal storage. used only for upgrading XnoServers contract.
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
        XnoServersDeactivated();
    }

    /// @dev adds an admin given the address of the admin. Only admins are allowed to call this function.
    function addXnoAdmin(address _admin) 
        external
        storageAttached(true)
        onlyOwner
    {
        eternalStorage.addAdmin(_admin);
    }

    /// @dev removes an admin by the address of the admin. Only admins are allowed to call this function.
    function removeXnoAdmin(address _admin) 
        external
        storageAttached(true)
        onlyOwner
    {
        eternalStorage.removeAdmin(_admin);
    }        

    // onlyAdmin METHODS
    // Only the admins, which are registered in the eternal storage attached to this contract, can call the following methods.

    /// @dev removes a Xenio Server given the address of the Server. Only admins are allowed to call this function.
    function removeServerByAddress(address _address)
        external 
        storageAttached(true) 
        onlyAdmin 
        returns (bool success) 
    {
        var (found, serverID) = eternalStorage.getIDByAddress(_address);
        if (found) {eternalStorage.removeServer(serverID);} // remove server
        return found;
    }

    // Update methods for admins

    /// @dev updates the address of a server given current address of the server. Only admins are allowed to call this function.
    function changeServerAddress(address _currentAddress, address _newAddress) external storageAttached(true) onlyAdmin returns(bool success) {
        success = eternalStorage.updateAddress(_currentAddress, _newAddress);
    }

    /// @dev updates the ip of a server given current address of the server. Only admins are allowed to call this function.
    function changeServerIP(address _currentAddress, bytes32 _ip) external storageAttached(true) onlyAdmin returns(bool success) {
        success = eternalStorage.updateIP(_currentAddress, _ip);
    }

    /// @dev updates the type of a server given current address of the server. Only admins are allowed to call this function.
    function changeServerType(address _currentAddress, bytes32 _type) external storageAttached(true) onlyAdmin returns(bool success) {
        success = eternalStorage.updateType(_currentAddress, _type);
    }

    // MAIN METHODS

    /// @dev Constructor
    function XnoServers() public {
        XnoServersCreated(getOwner());
    }

    /// @dev registers a new Xenio server 
    function registerNewServer(bytes32 _ip, bytes32 _type)  
        public 
        storageAttached(true)
        returns (bool success)
    {
        return eternalStorage.addServer(msg.sender, _ip, _type);
    }
    
    /// @dev deregisters the respective Xenio server (if exist) of the sender's address
    function deregisterServer() public storageAttached(true) {
        var (found, id) = eternalStorage.getIDByAddress(msg.sender); // get server id
        if (found) {eternalStorage.removeServer(id);} // remove server
    }

    // Update methods for senders

    /// @dev updates the ip of the server owned by the sender.
    function updateServerIP(bytes32 _ip) external storageAttached(true) returns(bool success) {
        success = eternalStorage.updateIP(msg.sender, _ip);
    }

    /// @dev updates the type of the server owned by the sender.
    function updateServerType(bytes32 _type) external storageAttached(true) returns(bool success) {
        success = eternalStorage.updateType(msg.sender, _type);
    }

    // GETTER METHODS

    /// @dev Returns the total number of active Xenio servers.
    function getServersCount() external storageAttached(true) constant returns(uint) {
        return eternalStorage.getActiveServersCount();
    }

    /// @dev Returns the server id of a given address (unique)
    function getServerIDByAddress(address _address) external storageAttached(true) constant returns(bool found, uint serverID) {
        return eternalStorage.getIDByAddress(_address);
    }

    /// @dev Returns server details by its id
    function getServerByID(uint _serverID) 
        external 
        storageAttached(true) 
        constant 
        returns
    (
        bool found,
        address serverAddress,
        bytes32 serverIP,
        bytes32 serverType  
    ) 
    {
        return eternalStorage.getServerDetailsByID(_serverID);
    }

    // @dev Returns a list with all active serverIDs.
    function getAllServerIDs() external storageAttached(true) constant returns(uint[] memory list) {
        uint counter = 0;
        list = new uint[](eternalStorage.getActiveServersCount()); // initialise
        for (uint i = 0; i < eternalStorage.getServersCount(); i++) {
            if (eternalStorage.isActive(i))
                list[counter++] = i;
        }
    }

}