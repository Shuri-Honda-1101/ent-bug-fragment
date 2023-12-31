// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Shuri-Honda-1101/ent-bug-fragment/ent/project"
)

// Project is the model entity for the Project schema.
type Project struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProjectQuery when eager-loading is set.
	Edges        ProjectEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ProjectEdges holds the relations/edges for other nodes in the graph.
type ProjectEdges struct {
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// Tasks holds the value of the tasks edge.
	Tasks []*Task `json:"tasks,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedUsers map[string][]*User
	namedTasks map[string][]*Task
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// TasksOrErr returns the Tasks value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) TasksOrErr() ([]*Task, error) {
	if e.loadedTypes[1] {
		return e.Tasks, nil
	}
	return nil, &NotLoadedError{edge: "tasks"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Project) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case project.FieldID:
			values[i] = new(sql.NullInt64)
		case project.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Project fields.
func (pr *Project) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case project.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case project.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Project.
// This includes values selected through modifiers, order, etc.
func (pr *Project) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryUsers queries the "users" edge of the Project entity.
func (pr *Project) QueryUsers() *UserQuery {
	return NewProjectClient(pr.config).QueryUsers(pr)
}

// QueryTasks queries the "tasks" edge of the Project entity.
func (pr *Project) QueryTasks() *TaskQuery {
	return NewProjectClient(pr.config).QueryTasks(pr)
}

// Update returns a builder for updating this Project.
// Note that you need to call Project.Unwrap() before calling this method if this Project
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Project) Update() *ProjectUpdateOne {
	return NewProjectClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Project entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Project) Unwrap() *Project {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Project is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Project) String() string {
	var builder strings.Builder
	builder.WriteString("Project(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteByte(')')
	return builder.String()
}

// NamedUsers returns the Users named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Project) NamedUsers(name string) ([]*User, error) {
	if pr.Edges.namedUsers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedUsers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Project) appendNamedUsers(name string, edges ...*User) {
	if pr.Edges.namedUsers == nil {
		pr.Edges.namedUsers = make(map[string][]*User)
	}
	if len(edges) == 0 {
		pr.Edges.namedUsers[name] = []*User{}
	} else {
		pr.Edges.namedUsers[name] = append(pr.Edges.namedUsers[name], edges...)
	}
}

// NamedTasks returns the Tasks named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Project) NamedTasks(name string) ([]*Task, error) {
	if pr.Edges.namedTasks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedTasks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Project) appendNamedTasks(name string, edges ...*Task) {
	if pr.Edges.namedTasks == nil {
		pr.Edges.namedTasks = make(map[string][]*Task)
	}
	if len(edges) == 0 {
		pr.Edges.namedTasks[name] = []*Task{}
	} else {
		pr.Edges.namedTasks[name] = append(pr.Edges.namedTasks[name], edges...)
	}
}

// Projects is a parsable slice of Project.
type Projects []*Project
