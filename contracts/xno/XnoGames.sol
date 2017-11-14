pragma solidity ^0.4.4;

// IMPORTS
import "./Ownable.sol";
import "./GamesLib.sol";
import "./SecurityLib.sol";
import "./EternalStorage.sol";

/// @title Main contract for Xenio games management
contract XnoGames is Ownable {

    using GamesLib for address;
    using SecurityLib for address;
    address public eternalStorage;

    // EVENTS

    // Storage
    event StorageCreated(address _storageAddress);
    event StorageAttached(address _storageAddress);
    event StorageDeleted();
    event StorageDetached(address _newOwner);

    // Contract
    event XnoGamesCreated(address _owner);
    event XnoGamesDeactivated();

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

    /// @dev change the ownership of the currently attached eternal storage. used only for upgrading XnoGames contract.
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
        XnoGamesDeactivated();
    }

    // onlyAdmin METHODS
    // Only the admins that are registered in the eternal storage attached to this contract can call the following methods.

    /// @dev removes a Xenio game for a given address. Only admins are allowed to call this function.
    function removeGameByAddress(address _address)
        external 
        storageAttached(true) 
        onlyAdmin 
        returns (bool success) 
    {
        var (found, gameID) = eternalStorage.getIDByAddress(_address);
        if (found) {eternalStorage.removeGame(gameID);} // remove game
        return found;
    }

    /// @dev removes a Xenio game for a given title. Only admins are allowed to call this function.
    function removeGameByTitle(bytes32 _title)
        external 
        storageAttached(true) 
        onlyAdmin 
        returns (bool success) 
    {
        var (found, gameID) = eternalStorage.getIDByTitle(_title);
        if (found) {eternalStorage.removeGame(gameID);} // remove game
        return found;
    }

    // MAIN METHODS

    /// @dev Constructor
    function XnoGames() public {
        XnoGamesCreated(getOwner());
    }

    /// @dev registers a new Xenio game 
    function registerNewGame
    (
        bytes32 _title, 
        bytes32 _genre, 
        bytes32 _publisher, 
        bytes32 _developer, 
        uint _release, 
        uint _price
    )  
        public 
        storageAttached(true)
        returns (bool success)
    {
        return eternalStorage.addGame(msg.sender, _title, _genre, _publisher, _developer, _release, _price);
    }
    
    /// @dev deregisters the respective Xenio game (if exist) of the sender's address
    function deregisterGame() public storageAttached(true) {
        var (found, id) = eternalStorage.getIDByAddress(msg.sender); // get game id
        if (found) {eternalStorage.removeGame(id);} // remove game
    }

    // Update methods

    /// @dev updates the title of the game owner by the sender
    function updateGameTitle(bytes32 _title) external storageAttached(true) returns(bool success) {
        success = eternalStorage.updateTitle(msg.sender, _title);
    }

    /// @dev updates the genre of the game owner by the sender
    function updateGameGenre(bytes32 _genre) external storageAttached(true) returns(bool success) {
        success = eternalStorage.updateGenre(msg.sender, _genre);
    }  

    /// @dev updates the publisher of the game owner by the sender
    function updateGamePublisher(bytes32 _publisher) external storageAttached(true) returns(bool success) {
        success = eternalStorage.updatePublisher(msg.sender, _publisher);
    }        

    /// @dev updates the developer of the game owner by the sender
    function updateGameDeveloper(bytes32 _developer) external storageAttached(true) returns(bool success) {
        success = eternalStorage.updateDeveloper(msg.sender, _developer);
    }      

    /// @dev updates the release date of the game owner by the sender
    function updateGameReleaseDate(uint _release) external storageAttached(true) returns(bool success) {
        success = eternalStorage.updateReleaseDate(msg.sender, _release);
    }

    /// @dev updates the release date of the game owner by the sender
    function updateGamePrice(uint _price) external storageAttached(true) returns(bool success) {
        success = eternalStorage.updatePrice(msg.sender, _price);
    }  

    /// @dev updates the logo image url of the game owner by the sender
    function updateGameLogo(bytes32 _imgUrl) external storageAttached(true) returns(bool success) {
        success = eternalStorage.updateLogo(msg.sender, _imgUrl);
    }    

    // GETTER METHODS

    /// @dev Returns the total number of active Xenio games.
    function getGamesCount() external storageAttached(true) constant returns(uint) {
        return eternalStorage.getActiveGamesCount();
    }

    /// @dev Returns the game id of a given address (unique)
    function getGameIDByAddress(address _address) external storageAttached(true) constant returns(bool found, uint gameID) {
        return eternalStorage.getIDByAddress(_address);
    }

    /// @dev Returns the game id of a given title (unique)
    function getGameIDByTitle(bytes32 _title) external storageAttached(true) constant returns(bool found, uint gameID) {
        return eternalStorage.getIDByTitle(_title);
    }

    /// @dev Returns game details by its id
    function getGameByID(uint _gameID) 
        external 
        storageAttached(true) 
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
        return eternalStorage.getGameDetailsByID(_gameID);
    }

    /// @dev Returns game details by its title
    function getGameByTitle(bytes32 _title) 
        external 
        storageAttached(true) 
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
        return eternalStorage.getGameDetailsByTitle(_title);
    }

    // @dev Returns a list with all active gameIDs.
    function getAllGameIDs() external storageAttached(true) constant returns(uint[] memory list) {
        uint counter = 0;
        list = new uint[](eternalStorage.getActiveGamesCount()); // initialise
        for (uint i = 0; i < eternalStorage.getGamesCount(); i++) {
            if (eternalStorage.isActive(i))
                list[counter++] = i;
        }
    }

}