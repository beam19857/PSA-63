// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/beam19857/app/ent/department"
	"github.com/beam19857/app/ent/expertise"
	"github.com/beam19857/app/ent/position"
	"github.com/beam19857/app/ent/user"
	"github.com/facebook/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DoctorID holds the value of the "DoctorID" field.
	DoctorID int `json:"DoctorID,omitempty"`
	// DoctorName holds the value of the "DoctorName" field.
	DoctorName string `json:"DoctorName,omitempty"`
	// DoctorEmail holds the value of the "DoctorEmail" field.
	DoctorEmail string `json:"DoctorEmail,omitempty"`
	// Date holds the value of the "Date" field.
	Date time.Time `json:"Date,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges                      UserEdges `json:"edges"`
	department_department_user *int
	expertise_expertise_user   *int
	position_position_user     *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// UserDepartment holds the value of the UserDepartment edge.
	UserDepartment *Department
	// UserExpertise holds the value of the UserExpertise edge.
	UserExpertise *Expertise
	// UserPosition holds the value of the UserPosition edge.
	UserPosition *Position
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserDepartmentOrErr returns the UserDepartment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) UserDepartmentOrErr() (*Department, error) {
	if e.loadedTypes[0] {
		if e.UserDepartment == nil {
			// The edge UserDepartment was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: department.Label}
		}
		return e.UserDepartment, nil
	}
	return nil, &NotLoadedError{edge: "UserDepartment"}
}

// UserExpertiseOrErr returns the UserExpertise value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) UserExpertiseOrErr() (*Expertise, error) {
	if e.loadedTypes[1] {
		if e.UserExpertise == nil {
			// The edge UserExpertise was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: expertise.Label}
		}
		return e.UserExpertise, nil
	}
	return nil, &NotLoadedError{edge: "UserExpertise"}
}

// UserPositionOrErr returns the UserPosition value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) UserPositionOrErr() (*Position, error) {
	if e.loadedTypes[2] {
		if e.UserPosition == nil {
			// The edge UserPosition was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: position.Label}
		}
		return e.UserPosition, nil
	}
	return nil, &NotLoadedError{edge: "UserPosition"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullInt64{},  // DoctorID
		&sql.NullString{}, // DoctorName
		&sql.NullString{}, //DoctorEmail
		&sql.NullTime{},   // Date
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*User) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // department_department_user
		&sql.NullInt64{}, // expertise_expertise_user
		&sql.NullInt64{}, // position_position_user
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field DoctorID", values[0])
	} else if value.Valid {
		u.DoctorID = int(value.Int64)
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field DoctorName", values[1])
	} else if value.Valid {
		u.DoctorName = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field DoctorEmail", values[1])
	} else if value.Valid {
		u.DoctorEmail = value.String
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field Date", values[2])
	} else if value.Valid {
		u.Date = value.Time
	}
	values = values[3:]
	if len(values) == len(user.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field department_department_user", value)
		} else if value.Valid {
			u.department_department_user = new(int)
			*u.department_department_user = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field expertise_expertise_user", value)
		} else if value.Valid {
			u.expertise_expertise_user = new(int)
			*u.expertise_expertise_user = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field position_position_user", value)
		} else if value.Valid {
			u.position_position_user = new(int)
			*u.position_position_user = int(value.Int64)
		}
	}
	return nil
}

// QueryUserDepartment queries the UserDepartment edge of the User.
func (u *User) QueryUserDepartment() *DepartmentQuery {
	return (&UserClient{config: u.config}).QueryUserDepartment(u)
}

// QueryUserExpertise queries the UserExpertise edge of the User.
func (u *User) QueryUserExpertise() *ExpertiseQuery {
	return (&UserClient{config: u.config}).QueryUserExpertise(u)
}

// QueryUserPosition queries the UserPosition edge of the User.
func (u *User) QueryUserPosition() *PositionQuery {
	return (&UserClient{config: u.config}).QueryUserPosition(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", DoctorID=")
	builder.WriteString(fmt.Sprintf("%v", u.DoctorID))
	builder.WriteString(", DoctorName=")
	builder.WriteString(u.DoctorName)
	builder.WriteString(", DoctorEmail=")
	builder.WriteString(u.DoctorEmail)
	builder.WriteString(", Date=")
	builder.WriteString(u.Date.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
