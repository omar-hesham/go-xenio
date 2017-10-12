pragma solidity ^0.4.4;

//import 'XNOGames';

contract XNOUsers {

    address UserAdmin;

    mapping ( address => User ) Users;      // this allows to look up Users by their Xenio address
    address[] usersByAddress;              // this is like a whitepages of all users, by Xenio address

    mapping ( bytes32 => NotarizedImage) notarizedImages; // this allows to look up notarizedImages by their SHA256notaryHash
    bytes32[] imagesByNotaryHash; // this is like a whitepages of all images, by SHA256notaryHash

    struct NotarizedImage {
        string imageURL;
        uint timeStamp;
    }

    struct User {
        bytes32 id;
        bool isServer;
        bytes32 country;
        bytes32 state;
        bytes32 image;
        bytes32 game;
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
    function registerNewUser(bytes32 _id, bool _isServer, bytes32 _country, bytes32 _state) public returns (bool success) {
        // don't overwrite existing game entries, and make sure that game name (handle) isn't null
        // if (name.length == 0) {return false;} // TODO: check if name exist
        User memory user; //initialise in memory - does not consume gas
        user.id = _id;
        user.isServer = _isServer;
        user.country = _country;
        user.state = _state;
        Users[msg.sender] = user;

        return true;
    }

   function addLogo(string _imageURL, bytes32 _SHA256notaryHash) public returns (bool success) {
        address thisAddress = msg.sender;
        //if (bytes(Users[thisAddress].id).length != 0) {return false;}                   // make sure this game is registered
        if (bytes(_imageURL).length != 0) {return false;}                                    // exit if image URL is empty
        if (bytes(notarizedImages[_SHA256notaryHash].imageURL).length != 0) {return false;}  // prevent duplicates over sha->image listings in the whitepages
        imagesByNotaryHash.push(_SHA256notaryHash);                                          // adds entry for this logo image to logo list whitepages
        notarizedImages[_SHA256notaryHash].imageURL = _imageURL;
        notarizedImages[_SHA256notaryHash].timeStamp = block.timestamp; // note that updating an image also updates the timestamp
        Users[thisAddress].image = _SHA256notaryHash; // add the image hash to this users .myImages array
        return true;
    }

    //function addServer(){
    //   XNOGames game = XNOGames(UserAdmin);
	//	game.getAllGames();
    //}

    // Only-Admin Functions
    function removeUser(address gameAddress) onlyAdmin public returns (bool success) {
        delete Users[gameAddress];
        return true;
    }

    function removeImage(bytes32 imgHash) onlyAdmin public returns (bool success) {
        delete notarizedImages[imgHash];
        return true;
    }

    // Getters
    //return a list with all registered users' addresses
    function getAllUsersAddresses() public constant returns (address[]) { return usersByAddress; }

    //return a list with all registered gamers' addresses
    function getAllGamersAddresses() public constant returns (address[]) {
        address[] gamersList;
        for ( uint i = 0; i < usersByAddress.length; i++ ) {
            if (!Users[usersByAddress[i]].isServer) {gamersList.push(usersByAddress[i]);}
        }
        return gamersList;
    }

    //return a list with all registered servers' addresses
    function getAllServersAddresses() public constant returns (address[]) {
        address[] serversList;
        for ( uint i = 0; i < usersByAddress.length; i++ ) {
            if (Users[usersByAddress[i]].isServer) {serversList.push(usersByAddress[i]);}
        }
        return serversList;
    }

    //return a list with all users registered
    function getAllUsers() public constant returns (bytes32[], bool[], bytes32[], bytes32[], bytes32[], bytes32[]) {

        uint len = usersByAddress.length;
        bytes32[] memory ids = new bytes32[](len);
        bool[] memory isServers = new bool[](len);
        bytes32[] memory countries = new bytes32[](len);
        bytes32[] memory states = new bytes32[](len);
        bytes32[] memory images = new bytes32[](len);
        bytes32[] memory games = new bytes32[](len);

        for ( uint i = 0; i < len; i++ ) {
            User memory user;
            user = Users[usersByAddress[i]];
            ids[i] = user.id;
            isServers[i] = user.isServer;
            countries[i] = user.country;
            states[i] = user.state;
            images[i] = user.image;
            games[i] = user.game;
        }
        return (ids,isServers,countries,states,images,games);
    }

    //return a list with all gamers registered
    function getAllGamers() public constant returns (bytes32[], bool[], bytes32[], bytes32[], bytes32[], bytes32[]) {

        uint len = usersByAddress.length;
        bytes32[] memory ids = new bytes32[](len);
        bool[] memory isServers = new bool[](len);
        bytes32[] memory countries = new bytes32[](len);
        bytes32[] memory states = new bytes32[](len);
        bytes32[] memory images = new bytes32[](len);
        bytes32[] memory games = new bytes32[](len);

        for ( uint i = 0; i < usersByAddress.length; i++ ) {
            User memory user;
            user = Users[usersByAddress[i]];
            if (!user.isServer) {
                ids[i] = user.id;
                isServers[i] = user.isServer;
                countries[i] = user.country;
                states[i] = user.state;
                images[i] = user.image;
                games[i] = user.game;
            }
        }
        return (ids,isServers,countries,states,images,games);
    }

    //return a list with all servers registered
    function getAllServers() public constant returns (bytes32[], bool[], bytes32[], bytes32[], bytes32[], bytes32[]) {

        uint len = usersByAddress.length;
        bytes32[] memory ids = new bytes32[](len);
        bool[] memory isServers = new bool[](len);
        bytes32[] memory countries = new bytes32[](len);
        bytes32[] memory states = new bytes32[](len);
        bytes32[] memory images = new bytes32[](len);
        bytes32[] memory games = new bytes32[](len);

        for ( uint i = 0; i < usersByAddress.length; i++ ) {
            User memory user;
            user = Users[usersByAddress[i]];
            if (user.isServer) {
                ids[i] = user.id;
                isServers[i] = user.isServer;
                countries[i] = user.country;
                states[i] = user.state;
                images[i] = user.image;
                games[i] = user.game;
            }
        }
        return (ids,isServers,countries,states,images,games);
    }


    //function getUser(address gameAddress) public constant returns (string,bool,bytes32,bytes32,bytes32) {
    //    return (Users[gameAddress].id, Users[gameAddress].isServer, Users[gameAddress].country, Users[gameAddress].state, Users[gameAddress].image);
    //}

    function getAllImages() public constant returns (bytes32[]) { return imagesByNotaryHash; }

    function getGameImages(address userAddress) public constant returns (bytes32) { return Users[userAddress].image; }

    function getImage(bytes32 SHA256notaryHash) public constant returns (string,uint) {
        return (notarizedImages[SHA256notaryHash].imageURL,notarizedImages[SHA256notaryHash].timeStamp);
    }

  
}