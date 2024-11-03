package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/earthboundkid/csv/v2"
	"github.com/joho/sqltocsv"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DB_NAME       = "tasks_data.db"
	TABLE_NAME    = "tasks"
	EXPORTED_DATA = "tasks.csv"
)

type Task struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Status    bool   `json:"status"`
	CreatedAt string `json:"created_at"`
}

var allTodos []Task

// App struct
type App struct {
	ctx context.Context
	db  *sql.DB
}

// NewApp creates a new App application struct
func NewApp() *App {
	// Open SQLite database
	db, err := sql.Open("sqlite3", fmt.Sprintf("./%s", DB_NAME))
	if err != nil {
		log.Fatal(err)
	}

	sqlQuery := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
		status BOOLEAN DEFAULT(FALSE),
        created_at TEXT
    )`, TABLE_NAME)

	// Create table if not exists
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}

	return &App{db: db}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	defer a.close()
	return false
}

// GetAllTodos retrieves data from the DB
func (a *App) GetAllTodos() []Task {
	return a.listTodos()
}

// AddTodo add a new todo to the list
func (a *App) AddTodo(taskTitle string) []Task {
	fmt.Printf("Adding todo %s to index %v\n", taskTitle, len(allTodos))

	task := Task{
		Title:     taskTitle,
		CreatedAt: time.Now().Format(time.RFC822Z),
	}

	// allTodos = append(allTodos, task)
	// writeData()

	a.insertTodo(&task)

	return a.listTodos()
}

// UpadateTodo update todo by Id
func (a *App) UpadateTodo(status bool, id int) []Task {
	fmt.Printf("Updating todo with index %v\n", id)

	a.toggleStatus(status, id)
	// writeData()

	return a.listTodos()
}

// RemoveTodo removes todo by Id
func (a *App) RemoveTodo(id int) []Task {
	fmt.Printf("Removing todo with index %v\n", id)

	a.deleteTodo(id)
	// writeData()

	return a.listTodos()
}

func (a *App) close() {
	err := a.db.Close()
	if err != nil {
		log.Fatalf("🔥 failed to close the connection DB: %s\n", err)
	}
}

func (a *App) insertTodo(t *Task) {
	sqlQuery := fmt.Sprintf("INSERT INTO %s (title, created_at) VALUES (?, ?)", TABLE_NAME)
	_, err := a.db.Exec(sqlQuery, t.Title, t.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) listTodos() []Task {
	tasks := []Task{}

	sqlQuery := fmt.Sprintf("SELECT * FROM %s", TABLE_NAME)

	rows, err := a.db.Query(sqlQuery)
	if err != nil {
		return tasks
	}
	// We close the resource
	defer rows.Close()

	t := Task{}

	for rows.Next() {
		rows.Scan(&t.Id, &t.Title, &t.Status, &t.CreatedAt)

		tasks = append(tasks, t)
	}

	return tasks
}

func (a *App) toggleStatus(status bool, id int) {
	sqlQuery := fmt.Sprintf(`UPDATE %s SET status = ?
	WHERE id=?`, TABLE_NAME)

	stmt, err := a.db.Prepare(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(status, id)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) deleteTodo(id int) {
	sqlQuery := fmt.Sprintf(`DELETE FROM %s
		WHERE id=?`, TABLE_NAME)

	stmt, err := a.db.Prepare(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) exportData() bool {
	sqlQuery := fmt.Sprintf("SELECT * FROM %s", TABLE_NAME)

	rows, err := a.db.Query(sqlQuery)
	if err != nil {
		return err == nil
	}
	// We close the resource
	defer rows.Close()

	err = sqltocsv.WriteFile(EXPORTED_DATA, rows)
	if err != nil {
		return err == nil
	}

	return true
}

func (a *App) dropTable() bool {
	sqlQuery := fmt.Sprintf(`DELETE FROM %s;`, TABLE_NAME)

	_, err := a.db.Exec(sqlQuery)
	if err != nil {
		return err == nil
	}

	return true
}

func (a *App) convertToBool(s string) bool {
	boolValue, err := strconv.ParseBool(s)
	if err != nil {
		log.Fatal(err)
	}
	return boolValue
}

func (a *App) importData() bool {
	data, err := os.ReadFile(EXPORTED_DATA)
	if err != nil {
		return err == nil
	}

	csvOpt := csv.Options{Reader: strings.NewReader(string(data))}
	rowsMap, err := csvOpt.ReadAll()
	if err != nil {
		return err == nil
	}

	sqlQuery := fmt.Sprintf("INSERT INTO %s(title, status, created_at) VALUES ", TABLE_NAME)
	vals := []interface{}{}

	for _, row := range rowsMap {
		sqlQuery += "(?, ?, ?),"
		vals = append(
			vals,
			row["title"],
			a.convertToBool(row["status"]),
			row["created_at"],
		)
	}
	// trim the last `,`
	sqlQuery = strings.TrimSuffix(sqlQuery, ",")

	// prepare the statement
	stmt, err := a.db.Prepare(sqlQuery)
	if err != nil {
		return err == nil
	}

	defer stmt.Close()

	// We delete the old table so that
	// the import generates a new `clean` table
	a.dropTable()

	// format all vals at once
	_, err = stmt.Exec(vals...)
	if err != nil {
		return err == nil
	}

	return true
}

/* REFERENCES:
https://github.com/Carl0sPineda/WailsSqlite
https://github.com/search?q=wails%20sqlite&type=repositories
*/

/* func readDb(f *os.File) []Task {
	defer f.Close()
	var data []Task

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(b, &data)
	if strings.Contains(err.Error(), "unexpected end of JSON input") {
		data, err := json.MarshalIndent([]Task{}, "", " ")
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile(DbName, data, 0644)
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	}

	return data
}

func writeData() {
	data, _ := json.MarshalIndent(allTodos, "", " ")
	err := os.WriteFile(DbName, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func touchFile(name string) *os.File {
	file, err := os.OpenFile(name, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return file
} */

/* // Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
} */
