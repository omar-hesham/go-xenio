pragma solidity ^0.4.4;

//import 'XNOGames';

contract XNOUsers {

    address UserAdmin;

    mapping ( address => User ) Users;      // this allows to look up Users by their Xenio address
    address[] gamersByAddress;              // this is like a whitepages of all gamers, by Xenio address
    address[] serversByAddress;             // this is like a whitepages of all servers, by Xenio address

    struct User {
        string name;
        bool isServer;
        string game;
        bool isRegistered;
    }

    modifier onlyAdmin() {
        if (msg.sender != UserAdmin)
            revert(); //throw
        _; // Do not forget the "_;"! It will be replaced by the actual function body when the modifier is used.
    }

    // Constructor
    function XNOUsers() public payable {
        UserAdmin = msg.sender;
    }

    // Fallback function: This function is called when someone just sent Ether to the contract without providing any data,
    // or if someone messed up the types so that they tried to call a function that does not exist.
    function () {} // add payable if there is an associated cost with the game entries

    // Methods
    function registerNewUser(string _name, bool _isServer, string _game) public returns (bool success) {
        address thisAddress = msg.sender;

        if (Users[thisAddress].isRegistered == true) {return false;} // check if address exist
        if (bytes(_name).length == 0) {return false;} // check if name is empty

        User memory user; //initialise in memory - it does not consume gas
        user.name = _name;
        user.isServer = _isServer;
        user.game = _game;
        user.isRegistered = true;
        Users[thisAddress] = user;

        //add the address to the servers or gamers list
        if (_isServer == true) {serversByAddress.push(thisAddress);} else {gamersByAddress.push(thisAddress);}

        return true;
    }

    // Only-Admin Functions
    function removeUser(address gameAddress) onlyAdmin public returns (bool success) {
        delete Users[gameAddress];
        return true;
    }

    function stringToBytes32(string memory source) returns (bytes32 result) {
    assembly {
            result := mload(add(source, 32))
        }
    }

    // Getters
    //return a list with all users' addresses registered
    function getAllUsersAddresses() public constant returns (address[]){
        uint nServers = serversByAddress.length;
        uint nGamers = gamersByAddress.length;
        address[] memory list = new address[](nServers + nGamers);
        for (uint8 i = 0; i < nServers; i++) {
            list[i] = serversByAddress[i];
        }
        for (uint8 j = 0; j < nGamers; j++) {
            list[nServers+j] = gamersByAddress[j];
        }
        return list;
    }

    //return a list with all registered gamers' addresses
    function getGamersAddresses() public constant returns (address[]) { return gamersByAddress; }

    //return a list with all registered servers' addresses
    function getServersAddresses() public constant returns (address[]) { return serversByAddress; }

    //return the number of registered gamers
    function getGamersNumber() public constant returns (uint) { return gamersByAddress.length; }

    //return the number of registered servers
    function getServersNumber() public constant returns (uint) { return serversByAddress.length; }

    //return true if the address corresponds to a server, otherwise return false
    function isServer(address _address) public constant returns (bool) { return Users[_address].isServer; }

    //return user details for a given address
    function getUser(address _address) public constant returns (string,bool,string) {
        return (Users[_address].name, Users[_address].isServer, Users[_address].game);
    }

    //return all servers
    function getServers() public constant returns (bytes32[], bytes32[]) {
        uint nServers = serversByAddress.length;
        address serverAddress;
        bytes32[] memory names = new bytes32[](nServers);
        bytes32[] memory games = new bytes32[](nServers);
        for (uint8 i = 0; i < nServers; i++) {
            serverAddress = serversByAddress[i];
            names[i] = stringToBytes32(Users[serverAddress].name);
            games[i] = stringToBytes32(Users[serverAddress].game);
        }
        return(names, games);
    }

    //return all gamers
    function getGamers() public constant returns (bytes32[], bytes32[]) {
        uint nGamers = gamersByAddress.length;
        address gamerAddress;
        bytes32[] memory names = new bytes32[](nGamers);
        bytes32[] memory games = new bytes32[](nGamers);
        for (uint8 i = 0; i < nGamers; i++) {
            gamerAddress = gamersByAddress[i];
            names[i] = stringToBytes32(Users[gamerAddress].name);
            games[i] = stringToBytes32(Users[gamerAddress].game);
        }
        return(names, games);
    }

    //return all users
    function getUsers() public constant returns (bytes32[], bytes32[]) {
        uint nGamers = gamersByAddress.length;
        uint nServers = serversByAddress.length;
        address tempAddress;
        bytes32[] memory names = new bytes32[](nGamers+nServers);
        bytes32[] memory games = new bytes32[](nGamers+nServers);
        for (uint8 i = 0; i < nGamers; i++) {
            tempAddress = gamersByAddress[i];
            names[i] = stringToBytes32(Users[tempAddress].name);
            games[i] = stringToBytes32(Users[tempAddress].game);
        }
        for (uint8 j = 0; j < nServers; j++) {
            tempAddress = serversByAddress[j];
            names[nGamers+j] = stringToBytes32(Users[tempAddress].name);
            games[nGamers+j] = stringToBytes32(Users[tempAddress].game);
        }
        return(names, games);
    }


}