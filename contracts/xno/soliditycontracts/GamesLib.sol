pragma solidity ^0.4.4;

// IMPORTS
import "./EternalStorage.sol";

// GAME STRUCTURE
// keccak256("game_title", idx) -> bytes32 -- Game title.
// keccak256("game_address", idx) -> address -- Game address: the address of the account that registed the game.
// keccak256("game_registered", idx) -> bool -- If false the game is considered to be deregistered, or deleted.
// keccak256("game_genre", idx, gdix) -> bytes32 -- Game genre.
// keccak256("game_genre_registered", idx, gdix) -> bool -- If false the genre is considered to be deleted.
// keccak256("game_genres_count", idx) -> uint -- Total count of genres for the respective game.
// keccak256("game_publisher", idx) -> bytes32 -- Game publisher.
// keccak256("game_developer", idx) -> bytes32 -- Game developer.
// keccak256("game_release_date", idx) -> uint -- Game release date.
// keccak256("game_price", idx) -> uint -- Game price in USD.
// keccak256("game_server", idx, gidx) -> address -- Server address.
// keccak256("game_servers_count", idx) -> uint -- Total count of servers.
// keccak256("game_img_logo", idx) -> bytes32 -- imageUrl for the logo of the game.
// keccak256("games_count") -> uint256 -- Total count of registered games.

// Note that idx and gidx should always be uint

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
        EternalStorage(_storageContract).setBytes32Value(keccak256("game_genre", idx, uint(0)), _genre);        
        EternalStorage(_storageContract).setBooleanValue(keccak256("game_genre_registered", idx, uint(0)), true);         
        EternalStorage(_storageContract).setUIntValue(keccak256("game_genres_count", idx), uint(1)); 
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

    /// @dev updates the title of the game
    function updateTitle(address _storageContract, address _address, bytes32 _title) public returns (bool) {
        var (found, gameID) = getIDByAddress(_storageContract,_address);
        if (found && isTitleAvailable(_storageContract,_title)) {
            EternalStorage(_storageContract).setBytes32Value(keccak256("game_title", gameID), _title);
        }
        return found;
    } 

    /// @dev updates the publisher of the game
    function updatePublisher(address _storageContract, address _address, bytes32 _publisher) public returns (bool) {
        var (found, gameID) = getIDByAddress(_storageContract,_address);
        if (found) {EternalStorage(_storageContract).setBytes32Value(keccak256("game_publisher", gameID), _publisher);}
        return found;
    }

    /// @dev updates the developer of the game
    function updateDeveloper(address _storageContract, address _address, bytes32 _developer) public returns (bool) {
        var (found, gameID) = getIDByAddress(_storageContract,_address);
        if (found) {EternalStorage(_storageContract).setBytes32Value(keccak256("game_developer", gameID), _developer);}
        return found;
    }      

    /// @dev updates the release date of the game
    function updateReleaseDate(address _storageContract, address _address, uint _release) public returns (bool) {
        var (found, gameID) = getIDByAddress(_storageContract,_address);
        if (found) {EternalStorage(_storageContract).setUIntValue(keccak256("game_release_date", gameID), _release);}
        return found;
    }     

    /// @dev updates the price of the game
    function updatePrice(address _storageContract, address _address, uint _price) public returns (bool) {
        var (found, gameID) = getIDByAddress(_storageContract,_address);
        if (found) {EternalStorage(_storageContract).setUIntValue(keccak256("game_price", gameID), _price);}
        return found;
    }          

    /// @dev updates the image url of the logo of the game
    function updateLogo(address _storageContract, address _address, bytes32 _imageUrl) public returns (bool) {
        var (found, gameID) = getIDByAddress(_storageContract,_address);
        if (found) {EternalStorage(_storageContract).setBytes32Value(keccak256("game_img_logo", gameID), _imageUrl);}
        return found;
    }

    /// @dev updates the address of the game
    function updateAddress(address _storageContract, address _address, address _newAddress) public returns (bool) {
        var (found, gameID) = getIDByAddress(_storageContract,_address);
        if (found) {EternalStorage(_storageContract).setAddressValue(keccak256("game_address", gameID), _newAddress);}
        return found;
    }

    /// @dev adds a new genre description to the game
    function addGenre(address _storageContract, address _address, bytes32 _genre) public returns (bool) {
        var (found, gameID) = getIDByAddress(_storageContract,_address);
        if (found) {

            // Get first available genre id
            uint gidx = getNextGenreID(_storageContract,gameID,0); 

            // Validation check
            require(isGenreDuplicated(_storageContract,gameID,_genre)); // check whether the genre for this game already exists
            require(gidx + 1 <= 10); // maximum number of genres per game

            // Insert
            EternalStorage(_storageContract).setBytes32Value(keccak256("game_genre", gameID, gidx), _genre);        
            EternalStorage(_storageContract).setBooleanValue(keccak256("game_genre_registered", gameID, gidx), true);               

            // Update games count
            if (gidx + 1 > getGenresCount(_storageContract,gameID))
                EternalStorage(_storageContract).setUIntValue(keccak256("game_genres_count", gameID), gidx + 1);             

        }
        return found;
    }

    /// @dev removes a genre description from the respective game given its genre id
    function removeGenreByID(address _storageContract, address _address, uint _genreID) public returns (bool) {
        var (found, gameID) = getIDByAddress(_storageContract,_address);
        if (found) {EternalStorage(_storageContract).setBooleanValue(keccak256("game_genre_registered", gameID, _genreID), false);}
        return found;
    }    

    /// @dev removes a genre description from the respective game given its genre name
    function removeGenreByName(address _storageContract, address _address, bytes32 _genreName) public returns (bool) {
        var (gameFound, gameID) = getIDByAddress(_storageContract,_address);
        if (gameFound) {
            var (genreFound, genreID) = getGenreIDByName(_storageContract,gameID,_genreName);
            if (genreFound) {
                EternalStorage(_storageContract).setBooleanValue(keccak256("game_genre_registered", gameID, genreID), false);
                return true;
            }
        }
        return false;
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

    /// @dev Returns the total number of previously registered genres for a given gameID.
    function getGenresCount(address _storageContract, uint _gameID) private constant returns(uint256) {
        return EternalStorage(_storageContract).getUIntValue(keccak256("game_genres_count", _gameID));
    }

    /// @dev Returns the total number of active genres for a game given its ID.
    function getActiveGenresCount(address _storageContract, uint _gameID) private constant returns(uint activeGenresCount) {
        var totalCount = getGenresCount(_storageContract, _gameID);
        activeGenresCount = 0;
        for (uint i = 0; i < totalCount; i++) {
            if (EternalStorage(_storageContract).getBooleanValue(keccak256("game_genre_registered", _gameID, i))) {activeGenresCount++;}
        }
    }

    /// @dev Gets next available id for genres
    function getNextGenreID(address _storageContract, uint _gameID, uint _genreID) private constant returns(uint nextID) {
        nextID = _genreID; // initialise   
        while (EternalStorage(_storageContract).getBooleanValue(keccak256("game_genre_registered", _gameID, nextID))) {nextID++;}
    }    

    /// @dev Checks whether the genreID corresponds to an active genre for a given game
    function isGenreActive(address _storageContract, uint _gameID, uint _genreID) private constant returns(bool) {
        return EternalStorage(_storageContract).getBooleanValue(keccak256("game_registered", _gameID, _genreID));
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

    /// @dev Checks whether the given genre for a game already exists
    function isGenreDuplicated(address _storageContract, uint _gameID, bytes32 _genre) private constant returns (bool) {
        var totalCount = getGenresCount(_storageContract,_gameID);
        for (uint i = 0; i < totalCount; i++) {
            if (EternalStorage(_storageContract).getBooleanValue(keccak256("game_genre_registered", _gameID, i))) {
                if (keccak256(EternalStorage(_storageContract).getBytes32Value(keccak256("game_genre", _gameID, i))) == keccak256(_genre)) {
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

    /// @dev Returns genre id given a gameID and a genreName
    function getGenreIDByName(address _storageContract, uint _gameID, bytes32 _genreName) private constant returns (bool found, uint id) {
        var totalCount = getGenresCount(_storageContract,_gameID);
        for (uint i = 0; i < totalCount; i++) {
            if (EternalStorage(_storageContract).getBooleanValue(keccak256("game_genre_registered", _gameID, i))) {
                if (keccak256(EternalStorage(_storageContract).getBytes32Value(keccak256("game_genre", _gameID, i))) == keccak256(_genreName)) {  
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
        bytes32 gamePublisher, 
        bytes32 gameDeveloper, 
        uint gameRelease, 
        uint gamePrice,
        bytes32 gameLogo,
        bytes32[10] gameGenreList
    )
    {
        found = EternalStorage(_storageContract).getBooleanValue(keccak256("game_registered", _gameID));
        if (!found) {return;}        
        gameAddress = EternalStorage(_storageContract).getAddressValue(keccak256("game_address", _gameID));
        gameTitle = EternalStorage(_storageContract).getBytes32Value(keccak256("game_title", _gameID));
        gamePublisher = EternalStorage(_storageContract).getBytes32Value(keccak256("game_publisher", _gameID));
        gameDeveloper = EternalStorage(_storageContract).getBytes32Value(keccak256("game_developer", _gameID));
        gameRelease = EternalStorage(_storageContract).getUIntValue(keccak256("game_release_date", _gameID));
        gamePrice = EternalStorage(_storageContract).getUIntValue(keccak256("game_price", _gameID));
        gameLogo = EternalStorage(_storageContract).getBytes32Value(keccak256("game_img_logo", _gameID));        
        gameGenreList = getGenreListByID(_storageContract,_gameID);
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
        bytes32 gamePublisher, 
        bytes32 gameDeveloper, 
        uint gameRelease, 
        uint gamePrice,
        bytes32 gameLogo,
        bytes32[10] gameGenreList        
    )
    {
        (found, gameID) = getIDByTitle(_storageContract, _title); // get game id
        if (!found) {return;}
        gameAddress = EternalStorage(_storageContract).getAddressValue(keccak256("game_address", gameID));
        gamePublisher = EternalStorage(_storageContract).getBytes32Value(keccak256("game_publisher", gameID));
        gameDeveloper = EternalStorage(_storageContract).getBytes32Value(keccak256("game_developer", gameID));
        gameRelease = EternalStorage(_storageContract).getUIntValue(keccak256("game_release_date", gameID));
        gamePrice = EternalStorage(_storageContract).getUIntValue(keccak256("game_price", gameID));
        gameLogo = EternalStorage(_storageContract).getBytes32Value(keccak256("game_img_logo", gameID));
        gameGenreList = getGenreListByID(_storageContract,gameID);
    }

    /// @dev Returns the list of genres for a given game.
    function getGenreListByID(address _storageContract, uint _gameID)
        public
        constant
        returns
    (
        bytes32[10] genreList
    )
    {
        // Validation check
        if (!EternalStorage(_storageContract).getBooleanValue(keccak256("game_registered", _gameID))) {return;}

        // Update list
        for (uint i = 0; i < 10; i++) {
            if (EternalStorage(_storageContract).getBooleanValue(keccak256("game_genre_registered", _gameID, i))) {
                genreList[i] = EternalStorage(_storageContract).getBytes32Value(keccak256("game_genre", _gameID, i));
            }
        } 

    }

}