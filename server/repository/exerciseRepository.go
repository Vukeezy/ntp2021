package repository

//nesto
// The sql go library is needed to interact with the database
import (
	"database/sql"
	model "github.com/Vukeezy/main/model"
)


type Store interface {

	GetExercises() ([]*model.Exercise, error)

	//GetBirds() ([]*Bird, error)
}

// The `dbStore` struct will implement the `Store` interface
// It also takes the sql DB connection object, which represents
// the database connection.
type DbStore struct {
	Db *sql.DB
}

func (store *DbStore) GetExercises() ([]*model.Exercise, error) {
	// Query the database for all birds, and return the result to the
	// `rows` object
	query := `SELECT * FROM "exercise"`
	rows, err := store.Db.Query(query)
	// We return incase of an error, and defer the closing of the row structure
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create the data structure that is returned from the function.
	// By default, this will be an empty array of birds
	exercises := []*model.Exercise{}
	for rows.Next() {
		// For each row returned by the table, create a pointer to a bird,
		exercise := &model.Exercise{}
		// Populate the `Species` and `Description` attributes of the bird,
		// and return incase of an error
		if err := rows.Scan(&exercise.Id, &exercise.RequestedPreparedness, &exercise.Equipment, &exercise.Name, &exercise.Description, &exercise.Type); err != nil {
			return nil, err
		}
		// Finally, append the result to the returned array, and repeat for
		// the next row
		exercises = append(exercises, exercise)
	}
	return exercises, nil
}

// The store variable is a package level variable that will be available for
// use throughout our application code
var store Store

/*
We will need to call the InitStore method to initialize the store. This will
typically be done at the beginning of our application (in this case, when the server starts up)
This can also be used to set up the store as a mock, which we will be observing
later on
*/
func GetStore() Store{
	return store
}

func InitStore(s Store) {
	store = s
}
