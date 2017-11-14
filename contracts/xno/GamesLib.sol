pragma solidity ^0.4.4;

// IMPORTS
import "./EternalStorage.sol";

// GAME STRUCTURE
// keccak256("game_title", idx) -> bytes32 -- Game title.
// keccak256("game_address", idx) -> address -- Game address: the address of the account that registed the game.
// keccak256("game_registered", idx) -> bool -- If false the game is considered to be deregistered, or deleted.
// keccak256("game_genre", idx) -> bytes32 -- Game genre.
// keccak256("game_publisher", idx) -> bytes32 -- Game publisher.
// keccak256("game_developer", idx) -> bytes32 -- Game developer.
// keccak256("game_release_date", idx) -> uint -- Game release date.
// keccak256("game_price", idx) -> uint -- Game price in USD.
// keccak256("game_server", idx, gidx) -> address -- Server address.
// keccak256("game_servers_count", idx) -> uint -- Total count of servers.
// keccak256("game_img_logo", idx) -> bytes32 -- imageUrl for the logo of the game.
// keccak256("games_count") -> uint256 -- Total count of registered games.

// Note that idx should always be uint

/** @title Library for Users. */
library GamesLib {

    // EVENTS
    event GameAdded(uint _id);
    event GameRemoved(uint _id);    

    // MAIN METHODS

    /// @dev registers a new game
    function addGame
    (
        address _storageContract, 
        address _address, 
        bytes32 _title, 
        bytes32 _genre, 
        bytes32 _publisher, 
        bytes32 _developer, 
        uint _release, 
        uint _price
    )   
        public
        returns (bool success)
    {

        // Validation check
        require(isTitleAvailable(_storageContract,_title)); // check whether the title of the game is taken
        require(isAddressAvailable(_storageContract,_address)); // check whether the given address is taken

        // Get first available id
        uint idx = getNextID(_storageContract,0); 

        // Register game's details
        EternalStorage(_storageContract).setBytes32Value(keccak256("game_title", idx), _title);
        EternalStorage(_storageContract).setAddressValue(keccak256("game_address", idx), _address);
        EternalStorage(_storageContract).setBooleanValue(keccak256("game_registered", idx), true);
        EternalStorage(_storageContract).setBytes32Value(keccak256("game_genre", idx), _genre);        
        EternalStorage(_storageContract).setBytes32Value(keccak256("game_publisher", idx), _publisher); 
        EternalStorage(_storageContract).setBytes32Value(keccak256("game_developer", idx), _developer); 
        EternalStorage(_storageContract).setUIntValue(keccak256("game_release_date", idx), _release); 
        EternalStorage(_storageContract).setUIntValue(keccak256("game_price", idx), _price); 
        
        // Update games count
        if (idx + 1 > getGamesCount(_storageContract))
            EternalStorage(_storageContract).setUIntValue(keccak256("games_count"), idx + 1);        
        
        GameAdded(idx); // call event

        success = true;
    }

    /// @dev deregisters a game, by changing its registration state to false
    function removeGame(address _storageContract, uint idx) public {
        EternalStorage(_storageContract).setBooleanValue(keccak256("game_registered", idx), false);
        GameRemoved(idx); // call event
    }

    // UPDATE METHODS

    // For Sender

    /// @dev updates the title of the game owned by the function caller
    function updateTitle(address _storageContract, address _address, bytes32 _title) public returns (bool) {
        var (found, gameID) = getIDByAddress(_storageContract,_address);
        if (found && isTitleAvailable(_storageContract,_title)) {
            EternalStorage(_storageContract).setBytes32Value(keccak256("game_title", gameID), _title);
        }
        return found;
    }

    /// @dev updates the genre of the game owned by the function caller
    function updateGenre(address _storageContract, address _address, bytes32 _genre) public returns (bool) {
        var (found, gameID) = getIDByAddress(_storageContract,_address);
        if (found) {EternalStorage(_storageContract).setBytes32Value(keccak256("game_genre", gameID), _genre);}
        return found;
    }    

    /// @dev updates the publisher of the game owned by the function caller
    function updatePublisher(address _storageContract, address _address, bytes32 _publisher) public returns (bool) {
        var (found, gameID) = getIDByAddress(_storageContract,_address);
        if (found) {EternalStorage(_storageContract).setBytes32Value(keccak256("game_publisher", gameID), _publisher);}
        return found;
    }

    /// @dev updates the developer of the game owned by the function caller
    function updateDeveloper(address _storageContract, address _address, bytes32 _developer) public returns (bool) {
        var (found, gameID) = getIDByAddress(_storageContract,_address);
        if (found) {EternalStorage(_storageContract).setBytes32Value(keccak256("game_developer", gameID), _developer);}
        return found;
    }      

    /// @dev updates the release date of the game owned by the function caller
    function updateReleaseDate(address _storageContract, address _address, uint _release) public returns (bool) {
        var (found, gameID) = getIDByAddress(_storageContract,_address);
        if (found) {EternalStorage(_storageContract).setUIntValue(keccak256("game_release_date", gameID), _release);}
        return found;
    }     

    /// @dev updates the price of the game owned by the function caller
    function updatePrice(address _storageContract, address _address, uint _price) public returns (bool) {
        var (found, gameID) = getIDByAddress(_storageContract,_address);
        if (found) {EternalStorage(_storageContract).setUIntValue(keccak256("game_price", gameID), _price);}
        return found;
    }          

    /// @dev updates the price of the game owned by the function caller
    function updateLogo(address _storageContract, address _address, bytes32 _imageUrl) public returns (bool) {
        var (found, gameID) = getIDByAddress(_storageContract,_address);
        if (found) {EternalStorage(_storageContract).setBytes32Value(keccak256("game_img_logo", gameID), _imageUrl);}
        return found;
    }                 

    // GETTER METHODS

    // Variables

    /// @dev Returns the total number of previously registered games in Xenio.
    function getGamesCount(address _storageContract) public constant returns(uint256) {
        return EternalStorage(_storageContract).getUIntValue(keccak256("games_count"));
    }

    /// @dev Returns the total number of active games in Xenio.
    function getActiveGamesCount(address _storageContract) public constant returns(uint activeGamesCount) {
        var totalCount = getGamesCount(_storageContract);
        activeGamesCount = 0;
        for (uint i = 0; i < totalCount; i++) {
            if (EternalStorage(_storageContract).getBooleanValue(keccak256("game_registered", i))) {activeGamesCount++;}
        }
    }

    /// @dev Gets next available id
    function getNextID(address _storageContract, uint _id) private constant returns(uint nextID) {
        nextID = _id; // initialise   
        while (EternalStorage(_storageContract).getBooleanValue(keccak256("game_registered", nextID))) {nextID++;}
    }

    /// @dev Checks whether the gameID corresponds to an active game
    function isActive(address _storageContract, uint _id) public constant returns(bool) {
        return EternalStorage(_storageContract).getBooleanValue(keccak256("game_registered", _id));
    }

    /// @dev Checks whether the title of a given game is taken
    function isTitleAvailable(address _storageContract, bytes32 _title) private constant returns (bool) {
        var totalCount = getGamesCount(_storageContract);
        for (uint i = 0; i < totalCount; i++) {
            if (EternalStorage(_storageContract).getBooleanValue(keccak256("game_registered", i))) {
                if (keccak256(EternalStorage(_storageContract).getBytes32Value(keccak256("game_title", i))) == keccak256(_title)) {
                    return false;
                }
            }
        } 
        return true;
    }

    /// @dev Checks whether the given address is taken
    function isAddressAvailable(address _storageContract, address _address) private constant returns (bool) {
        var totalCount = getGamesCount(_storageContract);
        for (uint i = 0; i < totalCount; i++) {
            if (EternalStorage(_storageContract).getBooleanValue(keccak256("game_registered", i))) {
                if (keccak256(EternalStorage(_storageContract).getAddressValue(keccak256("game_address", i))) == keccak256(_address)) {
                    return false;
                }
            }
        } 
        return true;
    }    

    /// @dev Returns game id given its address
    function getIDByAddress(address _storageContract, address _address) public constant returns (bool found, uint id) {
        var totalCount = getGamesCount(_storageContract);
        for (uint i = 0; i < totalCount; i++) {
            if (EternalStorage(_storageContract).getBooleanValue(keccak256("game_registered", i))) {
                if (keccak256(EternalStorage(_storageContract).getAddressValue(keccak256("game_address", i))) == keccak256(_address)) {  
                    return (true, i);
                }
            }
        }
        return (false,0);
    }

    /// @dev Returns game id given its title
    function getIDByTitle(address _storageContract, bytes32 _title) public constant returns (bool found, uint id) {
        var totalCount = getGamesCount(_storageContract);
        for (uint i = 0; i < totalCount; i++) {
            if (EternalStorage(_storageContract).getBooleanValue(keccak256("game_registered", i))) {
                if (keccak256(EternalStorage(_storageContract).getBytes32Value(keccak256("game_title", i))) == keccak256(_title)) {  
                    return (true, i);
                }
            }
        }
        return (false,0);
    }

    // Game details

    /// @dev Returns game's details given its game id.
    function getGameDetailsByID(address _storageContract, uint _gameID)
        public
        constant
        returns
    (
        bool found,
        address gameAddress,
        bytes32 gameTitle,
        bytes32 gameGenre, 
        bytes32 gamePublisher, 
        bytes32 gameDeveloper, 
        uint gameRelease, 
        uint gamePrice,
        bytes32 gameLogo
    )
    {
        found = EternalStorage(_storageContract).getBooleanValue(keccak256("game_registered", _gameID));
        gameAddress = EternalStorage(_storageContract).getAddressValue(keccak256("game_address", _gameID));
        gameTitle = EternalStorage(_storageContract).getBytes32Value(keccak256("game_title", _gameID));
        gameGenre = EternalStorage(_storageContract).getBytes32Value(keccak256("game_genre", _gameID));
        gamePublisher = EternalStorage(_storageContract).getBytes32Value(keccak256("game_publisher", _gameID));
        gameDeveloper = EternalStorage(_storageContract).getBytes32Value(keccak256("game_developer", _gameID));
        gameRelease = EternalStorage(_storageContract).getUIntValue(keccak256("game_release_date", _gameID));
        gamePrice = EternalStorage(_storageContract).getUIntValue(keccak256("game_price", _gameID));
        gameLogo = EternalStorage(_storageContract).getBytes32Value(keccak256("game_img_logo", _gameID));        
    }

    /// @dev Returns game's details given its title.
    function getGameDetailsByTitle(address _storageContract, bytes32 _title)
        public
        constant
        returns
    (
        bool found,
        uint gameID,
        address gameAddress,
        bytes32 gameGenre, 
        bytes32 gamePublisher, 
        bytes32 gameDeveloper, 
        uint gameRelease, 
        uint gamePrice,
        bytes32 gameLogo
    )
    {
        (found, gameID) = getIDByTitle(_storageContract, _title); // get game id
        gameAddress = EternalStorage(_storageContract).getAddressValue(keccak256("game_address", gameID));
        gameGenre = EternalStorage(_storageContract).getBytes32Value(keccak256("game_genre", gameID));
        gamePublisher = EternalStorage(_storageContract).getBytes32Value(keccak256("game_publisher", gameID));
        gameDeveloper = EternalStorage(_storageContract).getBytes32Value(keccak256("game_developer", gameID));
        gameRelease = EternalStorage(_storageContract).getUIntValue(keccak256("game_release_date", gameID));
        gamePrice = EternalStorage(_storageContract).getUIntValue(keccak256("game_price", gameID));
        gameLogo = EternalStorage(_storageContract).getBytes32Value(keccak256("game_img_logo", gameID));
    }
}