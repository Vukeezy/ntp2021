/*package router

//nesto
// The sql go library is needed to interact with the database
import (
	"database/sql"
	"ntp-server/model"
)

type Store interface {
	getExercises() ([]*model.Comment, error)
	//GetBirds() ([]*Bird, error)
}

// The `dbStore` struct will implement the `Store` interface
// It also takes the sql DB connection object, which represents
// the database connection.
type DbStore struct {
	Db *sql.DB
}

func (store *DbStore) getExercises() ([]*model.Comment, error) {
	// Query the database for all birds, and return the result to the
	// `rows` object
	rows, err := store.db.Query("SELECT * from comment")
	// We return incase of an error, and defer the closing of the row structure
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create the data structure that is returned from the function.
	// By default, this will be an empty array of birds
	comments := []*model.Comment{}
	for rows.Next() {
		// For each row returned by the table, create a pointer to a bird,
		comment := &model.Comment{}
		// Populate the `Species` and `Description` attributes of the bird,
		// and return incase of an error
		if err := rows.Scan(&comment.Id, &comment.FullName, &comment.Content); err != nil {
			return nil, err
		}
		// Finally, append the result to the returned array, and repeat for
		// the next row
		comments = append(comments, comment)
	}
	return comments, nil
}

// The store variable is a package level variable that will be available for
// use throughout our application code
var store Store

/*
We will need to call the InitStore method to initialize the store. This will
typically be done at the beginning of our application (in this case, when the server starts up)
This can also be used to set up the store as a mock, which we will be observing
later on

func InitStore(s Store) {
	store = s
}*/