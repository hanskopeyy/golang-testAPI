# Initial Testing GoLang API

API Link: http://localhost:8000/

Endpoints:
 1. Get All Pokemons
	   URL: "/pokemons", Methods: "GET"
    Get All Pokemons Data available on the database
    
 2. Add a new Pokemon
	   URL: "/pokemons", Methods: "POST"
    Parameter: "pokename"
    Add new pokemon
    
 3. Get Pokemon by ID
	   URL: "/pokemons/{id}", Methods: "GET"
    Get specified pokemon data by their pokemon ID
    
 4. Edit Pokemon by ID
	   URL: "/pokemons/{id}", Methods: "PUT"
    Parameter: "pokename"
    Edit specified pokemon by ID
    
 5. Delete Pokemon by ID
	   URL: "/pokemons/{id}", Methods: "DELETE"
    Delete specified pokemon by ID

 6. Get All Users
	   URL: "/users", Methods: "GET"
    Get All Users Data from database
 
 7. Register a new User
    URL:	"/users", Methods: "POST"
    Parameter: "email", "password", "name"
    Register a new user to the database
    
 8. Get user by ID
    URL:	"/users/{id}", Methods: "GET"
    Get specified user data by their ID
    
 9. Edit User data by ID
    URL: "/users/{id}", Methods: "PUT"
    Parameter: "email" and/or "name"
    Edit specified user data by their ID
    
10. Delete user by ID
	   URL: "/users/{id}", Methods: "POST"
    Move specified user data to "deletedusers" table, and remove it from the "user" table

