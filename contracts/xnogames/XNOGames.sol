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
        bool isRegistered;
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
    function registerNewGame(string _name, string _publisher, string _developer, bytes32 _country, bytes32 _state) public returns (bool success) {
        address thisAddress = msg.sender;


        if (Games[thisAddress].isRegistered == true) {return false;} // check if address exist
        if (bytes(_name).length == 0) {return false;} // check if name is empty

        Game memory game; //initialise in memory - it does not consume gas
        game.name = _name;
        game.publisher = _publisher;
        game.developer = _developer;
        game.country = _country;
        game.state = _state;
        game.isRegistered = true;
        Games[thisAddress] = game;

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

    function stringToBytes32(string memory source) returns (bytes32 result) {
    assembly {
            result := mload(add(source, 32))
        }
    }

    // Getters
    //return the addresses of all registered games
    function getGamesAddresses() public constant returns (address[]) { return gamesByAddress; }

    //return the game details of a given address
    function getGame(address gameAddress) public constant returns (string,string,string,bytes32,bytes32,bytes32) {
        return (Games[gameAddress].name, Games[gameAddress].publisher, Games[gameAddress].developer, Games[gameAddress].country, Games[gameAddress].state, Games[gameAddress].logoImg);
    }

    //return a list of all games
    function getGames() public constant returns (bytes32[],bytes32[], bytes32[],bytes32[],bytes32[]) {
        uint nGames = gamesByAddress.length;
        address gameAddress;
        bytes32[] memory names = new bytes32[](nGames);
        bytes32[] memory publishers = new bytes32[](nGames);
        bytes32[] memory developers = new bytes32[](nGames);
        bytes32[] memory countries = new bytes32[](nGames);
        bytes32[] memory states = new bytes32[](nGames);
        for (uint8 i = 0; i < nGames; i++) {
            gameAddress = gamesByAddress[i];
            names[i] = stringToBytes32(Games[gameAddress].name);
            publishers[i] = stringToBytes32(Games[gameAddress].publisher);
            developers[i] = stringToBytes32(Games[gameAddress].developer);
            countries[i] = Games[gameAddress].country;
            states[i] = Games[gameAddress].state;
        }
        return(names, publishers, developers, countries, states);
    }

    function getAllImages() public constant returns (bytes32[]) { return imagesByNotaryHash; }

    function getGameImages(address userAddress) public constant returns (bytes32) { return Games[userAddress].logoImg; }

    function getImage(bytes32 SHA256notaryHash) public constant returns (string,uint) {
        return (notarizedImages[SHA256notaryHash].imageURL,notarizedImages[SHA256notaryHash].timeStamp);
    }


}