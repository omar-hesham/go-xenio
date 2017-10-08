pragma solidity ^0.4.4;
contract XNOGames {

    address GamesAdmin;

    mapping ( address => Game ) Games;    // this allows to look up Games by their Xenio address
    address[] gamesByAddress;             // this is like a whitepages of all games, by Xenio address

    mapping ( bytes32 => NotarizedImage) notarizedImages; // this allows to look up notarizedImages by their SHA256notaryHash
    bytes32[] imagesByNotaryHash; // this is like a whitepages of all images, by SHA256notaryHash

    struct NotarizedImage {
        string imageURL;
        uint timeStamp;
    }

    struct Game {
        string name;
        string publisher;
        string developer;
        bytes32 country;
        bytes32 state;    
        bytes32 logoImg;
    }

    modifier onlyAdmin() {
        if (msg.sender != GamesAdmin)
            revert(); //throw
        _; // Do not forget the "_;"! It will be replaced by the actual function body when the modifier is used.
    }    

    // Constructor
    function XNOGames() public payable {
        GamesAdmin = msg.sender;
    }

    // Fallback function: This function is called when someone just sent Ether to the contract without providing any data,
    // or if someone messed up the types so that they tried to call a function that does not exist.
    function () {} // add payable if there is an associated cost with the game entries

    // Methods
    function registerNewGame(string name, string publisher, string developer, bytes32 country, bytes32 state) public returns (bool success) {
        address thisAddress = msg.sender;
        // don't overwrite existing game entries, and make sure that game name (handle) isn't null
        // if (name.length == 0) {return false;} // TODO: check if name exist 
        Games[thisAddress].name = name;
        Games[thisAddress].publisher = publisher;
        Games[thisAddress].developer = developer;
        Games[thisAddress].state = state;
        Games[thisAddress].country = country;                                          
        gamesByAddress.push(thisAddress);  // adds an entry for this user to the user 'whitepages'
        return true;
    }

   function addLogo(string imageURL, bytes32 SHA256notaryHash) public returns (bool success) {
        address thisAddress = msg.sender;
        if (bytes(Games[thisAddress].name).length != 0) {return false;}                   // make sure this game is registered
        if (bytes(imageURL).length != 0) {return false;}                                    // exit if image URL is empty
        if (bytes(notarizedImages[SHA256notaryHash].imageURL).length != 0) {return false;}  // prevent duplicates over sha->image listings in the whitepages
        imagesByNotaryHash.push(SHA256notaryHash);                                          // adds entry for this logo image to logo list whitepages
        notarizedImages[SHA256notaryHash].imageURL = imageURL;
        notarizedImages[SHA256notaryHash].timeStamp = block.timestamp; // note that updating an image also updates the timestamp
        Games[thisAddress].logoImg = SHA256notaryHash; // add the image hash to this users .myImages array
        return true;
    }

    // Only-Admin Functions
    function removeUser(address gameAddress) onlyAdmin returns (bool success) {
        delete Games[gameAddress];
        return true;
    }

    function removeImage(bytes32 imgHash) onlyAdmin returns (bool success) {
        delete notarizedImages[imgHash];
        return true;
    }

    // Getters
    function getAllGames() public constant returns (address[]) { return gamesByAddress; }

    function getGame(address gameAddress) public constant returns (string,string,string,bytes32,bytes32,bytes32) {
        return (Games[gameAddress].name, Games[gameAddress].publisher, Games[gameAddress].developer, Games[gameAddress].country, Games[gameAddress].state, Games[gameAddress].logoImg);
    }

    function getAllImages() public constant returns (bytes32[]) { return imagesByNotaryHash; }

    function getGameImages(address userAddress) public constant returns (bytes32) { return Games[userAddress].logoImg; }

    function getImage(bytes32 SHA256notaryHash) public constant returns (string,uint) {
        return (notarizedImages[SHA256notaryHash].imageURL,notarizedImages[SHA256notaryHash].timeStamp);
    }

  
}