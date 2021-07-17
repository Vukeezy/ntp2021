package repository

//nesto
// The sql go library is needed to interact with the database
import (
	"database/sql"
	model "github.com/Vukeezy/main/model"
	"log"
	"strconv"
)


type Store interface {

	GetExercises() ([]*model.Exercise, error)
	RateExercise(rate int, exerciseId int)  error
	RateComment(rate int, commentId int)  error
	AddComment(exerciseId int, commentContent string, commentFullname string) error
	SearchExercises(requestedPreparedness int, equipment int, muscle int, name string) ([]*model.Exercise, error)
	GetExercisesSorted() ([]*model.Exercise, error)
	GetExerciseById(exerciseId int) (*model.Exercise, error)
	GetPersonalExercises(agerecommendation int, preparedness int, goal int) ([]*model.Exercise, error)
//GetBirds() ([]*Bird, error)
}

// The `dbStore` struct will implement the `Store` interface
// It also takes the sql DB connection object, which represents
// the database connection.
type DbStore struct {
	Db *sql.DB
}

func (store *DbStore) SearchExercises(requestedPreparedness int, equipment int, muscle int, name string) ([]*model.Exercise, error) {
	query := "SELECT \"id\", requestedpreparedness,equipment, \"name\", description, \"type\", agerecommendation FROM exercise e "

	if muscle != -1 {
		query += "LEFT JOIN exercise_muscle em ON e.id = em.exercise_id WHERE em.muscle = " + strconv.Itoa(muscle)
	}

	if muscle == -1 {
		query += "WHERE true"
	}

	if requestedPreparedness != -1 {
		query += " and e.requestedpreparedness = " + strconv.Itoa(requestedPreparedness)
	}

	if equipment != -1 {
		query += " and e.equipment is " + strconv.FormatBool(equipment == 0)
	}

	query += " and e.name LIKE '%" + name + "%';"
	print(query)

	rows,_ := store.Db.Query(query)

	return getExerciseByRows(rows,store)
}

func (store *DbStore) GetPersonalExercises(agerecommendation int, preparedness int, goal int) ([]*model.Exercise, error) {
	query := "SELECT \"id\", requestedpreparedness,equipment, \"name\", description, \"type\", agerecommendation FROM exercise e WHERE true"


	if agerecommendation != -1 {
		query += " and e.agerecommendation = " + strconv.Itoa(agerecommendation)
	}

	if preparedness != -1 {
		query += " and e.requestedpreparedness = " + strconv.Itoa(preparedness)
	}

	if goal != -1 {
		query += " and e.type = " + strconv.Itoa(goal)
	}

	print(query)

	rows,_ := store.Db.Query(query)

	return getExerciseByRows(rows,store)
}

func (store *DbStore) AddComment(exerciseId int, commentContent string, commentFullname string) error {
	var commentId int
	row, err := store.Db.Query(`INSERT INTO "comment"(fullname,content) VALUES ($1,$2) RETURNING id`, commentFullname, commentContent)
	if row == nil{
		return err
	}
	row.Next()
	row.Scan(&commentId)
	_, err = store.Db.Query(`INSERT INTO exercise_comment VALUES ($1,$2)`, exerciseId, commentId)
	return err
}

func (store *DbStore) RateExercise(exerciseId int, rate int) (error) {
	_, err := store.Db.Query("INSERT INTO rate_exercise VALUES ($1,$2)", exerciseId, rate)
	return err
}

func (store *DbStore) RateComment(commentId int, rate int) error {
	_, err := store.Db.Query("INSERT INTO rate_comment VALUES ($1,$2)", commentId, rate)
	return err
}

func (store *DbStore) GetExercises() ([]*model.Exercise, error) {

	query := `with exercise_avg as(
				select re.exercise_id, avg(re.rate) as "avg"
				from rate_exercise re 
				group by re.exercise_id)
				select e.id, e.requestedpreparedness,e.equipment, e.name, e.description, e.type, e.agerecommendation as "avg"
				from exercise e, exercise_avg ea
				where e.id = ea.exercise_id;`
	rows, _ := store.Db.Query(query)

	return getExerciseByRows(rows,store)
}

func (store *DbStore) GetExercisesSorted() ([]*model.Exercise, error) {

	query := `with exercise_avg as(
				select re.exercise_id, avg(re.rate) as "avg"
				from rate_exercise re 
				group by re.exercise_id)
				select e.id, e.requestedpreparedness,e.equipment, e.name, e.description, e.type, e.agerecommendation as "avg"
				from exercise e, exercise_avg ea
				where e.id = ea.exercise_id
				order by ea.avg desc;`
	rows, _ := store.Db.Query(query)

	return getExerciseByRows(rows,store)
}

func (store *DbStore) GetExerciseById(exerciseId int) (*model.Exercise, error) {

	query := `with exercise_avg as(
				select re.exercise_id, avg(re.rate) as "avg"
				from rate_exercise re 
				group by re.exercise_id)
				select e.id, e.requestedpreparedness,e.equipment, e.name, e.description, e.type, e.agerecommendation as "avg"
				from exercise e, exercise_avg ea
				where e.id = ea.exercise_id
				and e.id = $1`
	rows, _ := store.Db.Query(query,exerciseId)
	exercises, err := getExerciseByRows(rows,store)

	return exercises[0],err
}

func getExerciseByRows(rows *sql.Rows, store *DbStore) ([]*model.Exercise,error) {


	exercises := []*model.Exercise{}

	for rows.Next() {
		exercise := &model.Exercise{}
		model.GetExerciseDTO(exercise)

		if err := rows.Scan(&exercise.Id, &exercise.RequestedPreparedness, &exercise.Equipment, &exercise.Name, &exercise.Description, &exercise.Type, &exercise.Age); err != nil {
			log.Fatal(err)
			return []*model.Exercise{}, err
		}

		muscles_query := `SELECT * FROM "exercise_muscle" WHERE "exercise_id" = ` + strconv.Itoa(exercise.Id)
		muscles_rows, _ := store.Db.Query(muscles_query)

		exercise.Muscles, _ = getMusclesByExerciseId(store, exercise.Id)

		exercise.Comments, _ = getCommentsByExerciseId(store, exercise.Id)

		exercise.Rates, _ = getRatesByExerciseId(store, exercise.Id)




		defer muscles_rows.Close()

		exercises = append(exercises, exercise)
	}

	return exercises, nil
}

func getMusclesByExerciseId(store *DbStore,exerciseId int) ([]int, error) {
	muscles_query := `SELECT * FROM "exercise_muscle" WHERE "exercise_id" = ` + strconv.Itoa(exerciseId)
	muscles_rows, _ := store.Db.Query(muscles_query)

	var muscles = []int{}

	for muscles_rows.Next() {
		var muscle int
		var dummy int

		if err := muscles_rows.Scan(&dummy, &muscle); err != nil {
			return []int{}, err
		}

		muscles = append(muscles, muscle)
	}

	return muscles, nil
}

func getCommentsByExerciseId(store *DbStore,exerciseId int) ([]model.Comment, error) {
	comment_query := `SELECT * FROM "exercise_comment" WHERE "exercise_id" = ` + strconv.Itoa(exerciseId)
	comment_rows, _ := store.Db.Query(comment_query)

	comments := []model.Comment{}

	for comment_rows.Next() {

		var comment_id, dummy int

		if err := comment_rows.Scan(&dummy, &comment_id); err != nil {
			return []model.Comment{}, err
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

	if comment_rows.Next() {
		var comment model.Comment

		if err := comment_rows.Scan(&comment.Id, &comment.FullName, &comment.Content); err != nil {
			return model.Comment{}, err
		}

		comment.Rates, _ = getRatesByCommentId(store, comment.Id)

		comment.Rate = 0
		if (len(comment.Rates) > 0) {
			rate_sum := 0.0
			for i := 0; i < len(comment.Rates); i++ {
				rate_sum += float64(comment.Rates[i])
			}
			comment.Rate = (rate_sum / float64(len(comment.Rates)))
		}
		return comment, nil
	}

	return model.Comment{}, nil
}

func getRatesByCommentId(store *DbStore, commentId int) ([]int, error) {
	rate_query := `SELECT * FROM "rate_comment" WHERE "comment_id" = ` + strconv.Itoa(commentId)
	rate_rows, _ := store.Db.Query(rate_query)

	var rates = []int{}

	if rate_rows == nil {
		return []int{},nil
	}

	for rate_rows.Next() {
		var rate,dummy int

		if err := rate_rows.Scan(&dummy, &rate); err != nil {
			return []int{}, err
		}

		rates = append(rates, rate)


	}
	return rates, nil
}

func getRatesByExerciseId(store *DbStore, exerciseId int) ([]int, error) {
	rate_query := `SELECT * FROM "rate_exercise" WHERE "exercise_id" = ` + strconv.Itoa(exerciseId)
	rate_rows, _ := store.Db.Query(rate_query)

	var rates = []int{}

	for rate_rows.Next() {
		var rate,dummy int

		if err := rate_rows.Scan(&dummy, &rate); err != nil {
			return []int{}, err
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
