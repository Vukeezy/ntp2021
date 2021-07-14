package repository

//nesto
// The sql go library is needed to interact with the database
import (
	"database/sql"
	model "github.com/Vukeezy/main/model"
	"strconv"
)


type Store interface {

	GetExercises() ([]*model.Exercise, error)
	RateExercise(rate int, exerciseId int)  error
	RateComment(rate int, commentId int)  error

	//GetBirds() ([]*Bird, error)
}

// The `dbStore` struct will implement the `Store` interface
// It also takes the sql DB connection object, which represents
// the database connection.
type DbStore struct {
	Db *sql.DB
}

func (store *DbStore) RateExercise(exerciseId int, rate int) error {
	_, err := store.Db.Query("INSERT INTO rate_exercise VALUES ($1,$2)", exerciseId, rate)
	return err
}

func (store *DbStore) RateComment(commentId int, rate int) error {
	_, err := store.Db.Query("INSERT INTO rate_comment VALUES ($1,$2)", commentId, rate)
	return err
}

func (store *DbStore) GetExercises() ([]*model.Exercise, error) {

	query := `SELECT * FROM "exercise"`
	rows, err := store.Db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	exercises := []*model.Exercise{}

	for rows.Next() {

		exercise := &model.Exercise{}
		model.GetExerciseDTO(exercise)

		if err := rows.Scan(&exercise.Id, &exercise.RequestedPreparedness, &exercise.Equipment, &exercise.Name, &exercise.Description, &exercise.Type); err != nil {
			return nil, err
		}

		muscles_query := `SELECT * FROM "exercise_muscle" WHERE "exercise_id" = ` + strconv.Itoa(exercise.Id)
		muscles_rows, muscles_err := store.Db.Query(muscles_query)

		exercise.Muscles, err = getMusclesByExerciseId(store, exercise.Id)
		exercise.Comments, err = getCommentsByExerciseId(store, exercise.Id)
		exercise.Rates, err = getRatesByExerciseId(store, exercise.Id)

		if muscles_err != nil {
			muscles_rows.Close()
			return nil, err
		}

		defer muscles_rows.Close()

		exercises = append(exercises, exercise)
	}
	return exercises, nil
}

func getMusclesByExerciseId(store *DbStore,exerciseId int) ([]int, error) {
	muscles_query := `SELECT * FROM "exercise_muscle" WHERE "exercise_id" = ` + strconv.Itoa(exerciseId)
	muscles_rows, _ := store.Db.Query(muscles_query)

	var muscles []int

	for muscles_rows.Next() {
		var muscle int
		var dummy int

		if err := muscles_rows.Scan(&dummy, &muscle); err != nil {
			return nil, err
		}

		muscles = append(muscles, muscle)
	}

	return muscles, nil
}

func getCommentsByExerciseId(store *DbStore,exerciseId int) ([]model.Comment, error) {
	comment_query := `SELECT * FROM "exercise_comment" WHERE "exercise_id" = ` + strconv.Itoa(exerciseId)
	comment_rows, _ := store.Db.Query(comment_query)

	var comments []model.Comment

	for comment_rows.Next() {
		var comment_id, dummy int

		if err := comment_rows.Scan(&dummy, &comment_id); err != nil {
			return nil, err
		}

		var comment model.Comment

		comment,_ = getCommentById(store, comment_id)

		comments = append(comments, comment)
	}

	return comments, nil
}

func getCommentById(store *DbStore, commentId int) (model.Comment, error) {
	comment_query := `SELECT * FROM "comment" WHERE "id" = ` + strconv.Itoa(commentId)
	comment_rows, _ := store.Db.Query(comment_query)


	for comment_rows.Next() {
		var comment model.Comment

		if err := comment_rows.Scan(&comment.Id, &comment.FullName, &comment.Content); err != nil {
			return model.Comment{}, err
		}

		comment.Rates, _ = getRatesByCommentId(store,comment.Id)

		return comment, nil
	}
	return model.Comment{}, nil
}

func getRatesByCommentId(store *DbStore, commentId int) ([]int, error) {
	rate_query := `SELECT * FROM "rate_comment" WHERE "comment_id" = ` + strconv.Itoa(commentId)
	rate_rows, _ := store.Db.Query(rate_query)

	var rates []int

	for rate_rows.Next() {
		var rate,dummy int

		if err := rate_rows.Scan(&dummy, &rate); err != nil {
			return nil, err
		}

		rates = append(rates, rate)

	}
	return rates, nil
}

func getRatesByExerciseId(store *DbStore, exerciseId int) ([]int, error) {
	rate_query := `SELECT * FROM "rate_exercise" WHERE "exercise_id" = ` + strconv.Itoa(exerciseId)
	rate_rows, _ := store.Db.Query(rate_query)

	var rates []int

	for rate_rows.Next() {
		var rate,dummy int

		if err := rate_rows.Scan(&dummy, &rate); err != nil {
			return nil, err
		}

		rates = append(rates, rate)

	}
	return rates, nil
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
