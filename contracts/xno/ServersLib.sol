pragma solidity ^0.4.4;

// IMPORTS
import "./EternalStorage.sol";

// SERVER STRUCTURE
// keccak256("server_ip", idx) -> bytes32 -- Server ip as string.
// keccak256("server_address", idx) -> address -- Server Xenio address.
// keccak256("server_type", idx) -> bytes32 -- Server types: "Master", "Voice"
// keccak256("server_registered", idx) -> bool -- If false the server is considered to be deregistered, or deleted.
// keccak256("server_game", idx) -> bytes32 -- Game Token ID (unique key)
// keccak256("servers_count") -> uint256 -- Total count of registered servers.

// Note that idx should always be uint

/** @title Library for Servers. */

/** Library for Servers. */
library ServersLib {

    // EVENTS
    event ServerAdded(uint _id);
    event ServerRemoved(uint _id);        

    // MAIN METHODS

    /// @dev registers a new server
    function addServer
    (
        address _storageContract, 
        address _address, 
        bytes32 _ip,
        bytes32 _type
    ) 
        public 
        returns(bool success)
    {
        // Validation check
        require(isAddressAvailable(_storageContract,_address)); // check whether the given address is taken

        // Get first available id
        uint idx = getNextID(_storageContract,0); 

        // Register server's details
        EternalStorage(_storageContract).setBytes32Value(keccak256("server_ip", idx), _ip);
        EternalStorage(_storageContract).setBytes32Value(keccak256("server_ip", idx), _type);        
        EternalStorage(_storageContract).setAddressValue(keccak256("server_address", idx), _address);
        EternalStorage(_storageContract).setBooleanValue(keccak256("server_registered", idx), true);

        // Update servers count
        if (idx + 1 > getServersCount(_storageContract))
            EternalStorage(_storageContract).setUIntValue(keccak256("servers_count"), idx + 1);        
        
        ServerAdded(idx); // call event

        success = true;        
    }

    /// @dev deregisters a server, by changing its registration state to false
    function removeServer(address _storageContract, uint idx) public {
        EternalStorage(_storageContract).setBooleanValue(keccak256("server_registered", idx), false);
        ServerRemoved(idx); // call event
    }    

    // UPDATE METHODS    

    /// @dev updates the ip address of the server
    function updateIP(address _storageContract, address _address, bytes32 _ip) public returns (bool) {
        var (found, serverID) = getIDByAddress(_storageContract,_address);
        if (found) {EternalStorage(_storageContract).setBytes32Value(keccak256("server_ip", serverID), _ip);}
        return found;
    }        

    /// @dev updates the address of the server
    function updateAddress(address _storageContract, address _address, address _newAddress) public returns (bool) {
        var (found, serverID) = getIDByAddress(_storageContract,_address);
        if (found) {EternalStorage(_storageContract).setAddressValue(keccak256("server_address", serverID), _newAddress);}
        return found;
    }

    /// @dev updates the type of the server
    function updateType(address _storageContract, address _address, bytes32 _type) public returns (bool) {
        var (found, serverID) = getIDByAddress(_storageContract,_address);
        if (found) {EternalStorage(_storageContract).setBytes32Value(keccak256("server_type", serverID), _type);}
        return found;
    }    
    
    // GETTER METHODS
    
    // Variables

    function getServersCount(address _storageContract) public constant returns(uint256) {
        return EternalStorage(_storageContract).getUIntValue(keccak256("servers_count"));
    }

    /// @dev Returns the total number of active servers in Xenio.
    function getActiveServersCount(address _storageContract) public constant returns(uint activeServersCount) {
        var totalCount = getServersCount(_storageContract);
        activeServersCount = 0;
        for (uint i = 0; i < totalCount; i++) {
            if (EternalStorage(_storageContract).getBooleanValue(keccak256("server_registered", i))) {activeServersCount++;}
        }
    }

    /// @dev Gets next available id
    function getNextID(address _storageContract, uint _id) private constant returns(uint nextID) {
        nextID = _id; // initialise   
        while (EternalStorage(_storageContract).getBooleanValue(keccak256("server_registered", nextID))) {nextID++;}
    }    

    /// @dev Checks whether the server ID corresponds to an active server
    function isActive(address _storageContract, uint _id) public constant returns(bool) {
        return EternalStorage(_storageContract).getBooleanValue(keccak256("server_registered", _id));
    }        

    /// @dev Checks whether the given address is taken
    function isAddressAvailable(address _storageContract, address _address) private constant returns (bool) {
        var totalCount = getServersCount(_storageContract);
        for (uint i = 0; i < totalCount; i++) {
            if (EternalStorage(_storageContract).getBooleanValue(keccak256("server_registered", i))) {
                if (keccak256(EternalStorage(_storageContract).getAddressValue(keccak256("server_address", i))) == keccak256(_address)) {
                    return false;
                }
            }
        } 
        return true;
    }     

    /// @dev Returns server id given its address
    function getIDByAddress(address _storageContract, address _address) public constant returns (bool found, uint id) {
        var totalCount = getServersCount(_storageContract);
        for (uint i = 0; i < totalCount; i++) {
            if (EternalStorage(_storageContract).getBooleanValue(keccak256("server_registered", i))) {
                if (keccak256(EternalStorage(_storageContract).getAddressValue(keccak256("server_address", i))) == keccak256(_address)) {  
                    return (true, i);
                }
            }
        }
        return (false,0);
    }

    // Server details

   /// @dev Returns server's details given its server id.
    function getServerDetailsByID(address _storageContract, uint _serverID)
        public
        constant
        returns
    (
        bool found,
        address serverAddress,
        bytes32 serverIp,
        bytes32 serverType
    )
    {
        found = EternalStorage(_storageContract).getBooleanValue(keccak256("server_registered", _serverID));
        if (!found) {return;}        
        serverAddress = EternalStorage(_storageContract).getAddressValue(keccak256("server_address", _serverID));
        serverIp = EternalStorage(_storageContract).getBytes32Value(keccak256("server_ip", _serverID));
        serverType = EternalStorage(_storageContract).getBytes32Value(keccak256("server_type", _serverID));
    }

}