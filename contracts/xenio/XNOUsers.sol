pragma solidity ^0.4.4;

//import 'XNOGames';

contract XNOUsers {

    address UserAdmin;

    mapping ( address => User ) Users;    // this allows to look up Games by their Xenio address
    address[] usersByAddress;             // this is like a whitepages of all games, by Xenio address

    mapping ( bytes32 => NotarizedImage) notarizedImages; // this allows to look up notarizedImages by their SHA256notaryHash
    bytes32[] imagesByNotaryHash; // this is like a whitepages of all images, by SHA256notaryHash

    struct NotarizedImage {
        string imageURL;
        uint timeStamp;
    }

    struct User {
        string id;
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
    function registerNewUser(string id, bool isServer, bytes32 country, bytes32 state) public returns (bool success) {
        address thisAddress = msg.sender;
        // don't overwrite existing game entries, and make sure that game name (handle) isn't null
        // if (name.length == 0) {return false;} // TODO: check if name exist 
        Users[thisAddress].id = id;
        Users[thisAddress].isServer = isServer;
        Users[thisAddress].country = country;
        Users[thisAddress].state = state;                                        
        usersByAddress.push(thisAddress);  // adds an entry for this user to the user 'whitepages'
        return true;
    }

   function addLogo(string imageURL, bytes32 SHA256notaryHash) public returns (bool success) {
        address thisAddress = msg.sender;
        if (bytes(Users[thisAddress].id).length != 0) {return false;}                   // make sure this game is registered
        if (bytes(imageURL).length != 0) {return false;}                                    // exit if image URL is empty
        if (bytes(notarizedImages[SHA256notaryHash].imageURL).length != 0) {return false;}  // prevent duplicates over sha->image listings in the whitepages
        imagesByNotaryHash.push(SHA256notaryHash);                                          // adds entry for this logo image to logo list whitepages
        notarizedImages[SHA256notaryHash].imageURL = imageURL;
        notarizedImages[SHA256notaryHash].timeStamp = block.timestamp; // note that updating an image also updates the timestamp
        Users[thisAddress].image = SHA256notaryHash; // add the image hash to this users .myImages array
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
    function getAllGames() public constant returns (address[]) { return usersByAddress; }

    function getGame(address gameAddress) public constant returns (string,bool,bytes32,bytes32,bytes32) {
        return (Users[gameAddress].id, Users[gameAddress].isServer, Users[gameAddress].country, Users[gameAddress].state, Users[gameAddress].image);
    }

    function getAllImages() public constant returns (bytes32[]) { return imagesByNotaryHash; }

    function getGameImages(address userAddress) public constant returns (bytes32) { return Users[userAddress].image; }

    function getImage(bytes32 SHA256notaryHash) public constant returns (string,uint) {
        return (notarizedImages[SHA256notaryHash].imageURL,notarizedImages[SHA256notaryHash].timeStamp);
    }

  
}