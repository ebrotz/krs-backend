package v1

// Place represents a Place tuple that would be returned
// by the database.
type Place struct {
	Id          int
	Description string
	Name        string
}

// Place represents a Place tuple that would be returned
// by the database.
type PlaceWebsite struct {
	// ID of the Place that the website is associated to.
	Id int
	// The URL of the website
	Website string
}

type PlacePhone struct {
	// ID of the Place that the website is associated to.
	Id int
	// The phone number of the place
	Phone string
}
