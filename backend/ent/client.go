// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/beam19857/app/ent/migrate"

	"github.com/beam19857/app/ent/department"
	"github.com/beam19857/app/ent/expertise"
	"github.com/beam19857/app/ent/position"
	"github.com/beam19857/app/ent/user"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Department is the client for interacting with the Department builders.
	Department *DepartmentClient
	// Expertise is the client for interacting with the Expertise builders.
	Expertise *ExpertiseClient
	// Position is the client for interacting with the Position builders.
	Position *PositionClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Department = NewDepartmentClient(c.config)
	c.Expertise = NewExpertiseClient(c.config)
	c.Position = NewPositionClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Department: NewDepartmentClient(cfg),
		Expertise:  NewExpertiseClient(cfg),
		Position:   NewPositionClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:     cfg,
		Department: NewDepartmentClient(cfg),
		Expertise:  NewExpertiseClient(cfg),
		Position:   NewPositionClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Department.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Department.Use(hooks...)
	c.Expertise.Use(hooks...)
	c.Position.Use(hooks...)
	c.User.Use(hooks...)
}

// DepartmentClient is a client for the Department schema.
type DepartmentClient struct {
	config
}

// NewDepartmentClient returns a client for the Department from the given config.
func NewDepartmentClient(c config) *DepartmentClient {
	return &DepartmentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `department.Hooks(f(g(h())))`.
func (c *DepartmentClient) Use(hooks ...Hook) {
	c.hooks.Department = append(c.hooks.Department, hooks...)
}

// Create returns a create builder for Department.
func (c *DepartmentClient) Create() *DepartmentCreate {
	mutation := newDepartmentMutation(c.config, OpCreate)
	return &DepartmentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Department.
func (c *DepartmentClient) Update() *DepartmentUpdate {
	mutation := newDepartmentMutation(c.config, OpUpdate)
	return &DepartmentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DepartmentClient) UpdateOne(d *Department) *DepartmentUpdateOne {
	mutation := newDepartmentMutation(c.config, OpUpdateOne, withDepartment(d))
	return &DepartmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DepartmentClient) UpdateOneID(id int) *DepartmentUpdateOne {
	mutation := newDepartmentMutation(c.config, OpUpdateOne, withDepartmentID(id))
	return &DepartmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Department.
func (c *DepartmentClient) Delete() *DepartmentDelete {
	mutation := newDepartmentMutation(c.config, OpDelete)
	return &DepartmentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DepartmentClient) DeleteOne(d *Department) *DepartmentDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DepartmentClient) DeleteOneID(id int) *DepartmentDeleteOne {
	builder := c.Delete().Where(department.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DepartmentDeleteOne{builder}
}

// Create returns a query builder for Department.
func (c *DepartmentClient) Query() *DepartmentQuery {
	return &DepartmentQuery{config: c.config}
}

// Get returns a Department entity by its id.
func (c *DepartmentClient) Get(ctx context.Context, id int) (*Department, error) {
	return c.Query().Where(department.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DepartmentClient) GetX(ctx context.Context, id int) *Department {
	d, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return d
}

// QueryDepartmentUser queries the DepartmentUser edge of a Department.
func (c *DepartmentClient) QueryDepartmentUser(d *Department) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(department.Table, department.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, department.DepartmentUserTable, department.DepartmentUserColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DepartmentClient) Hooks() []Hook {
	return c.hooks.Department
}

// ExpertiseClient is a client for the Expertise schema.
type ExpertiseClient struct {
	config
}

// NewExpertiseClient returns a client for the Expertise from the given config.
func NewExpertiseClient(c config) *ExpertiseClient {
	return &ExpertiseClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `expertise.Hooks(f(g(h())))`.
func (c *ExpertiseClient) Use(hooks ...Hook) {
	c.hooks.Expertise = append(c.hooks.Expertise, hooks...)
}

// Create returns a create builder for Expertise.
func (c *ExpertiseClient) Create() *ExpertiseCreate {
	mutation := newExpertiseMutation(c.config, OpCreate)
	return &ExpertiseCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Expertise.
func (c *ExpertiseClient) Update() *ExpertiseUpdate {
	mutation := newExpertiseMutation(c.config, OpUpdate)
	return &ExpertiseUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ExpertiseClient) UpdateOne(e *Expertise) *ExpertiseUpdateOne {
	mutation := newExpertiseMutation(c.config, OpUpdateOne, withExpertise(e))
	return &ExpertiseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ExpertiseClient) UpdateOneID(id int) *ExpertiseUpdateOne {
	mutation := newExpertiseMutation(c.config, OpUpdateOne, withExpertiseID(id))
	return &ExpertiseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Expertise.
func (c *ExpertiseClient) Delete() *ExpertiseDelete {
	mutation := newExpertiseMutation(c.config, OpDelete)
	return &ExpertiseDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ExpertiseClient) DeleteOne(e *Expertise) *ExpertiseDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ExpertiseClient) DeleteOneID(id int) *ExpertiseDeleteOne {
	builder := c.Delete().Where(expertise.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ExpertiseDeleteOne{builder}
}

// Create returns a query builder for Expertise.
func (c *ExpertiseClient) Query() *ExpertiseQuery {
	return &ExpertiseQuery{config: c.config}
}

// Get returns a Expertise entity by its id.
func (c *ExpertiseClient) Get(ctx context.Context, id int) (*Expertise, error) {
	return c.Query().Where(expertise.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ExpertiseClient) GetX(ctx context.Context, id int) *Expertise {
	e, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return e
}

// QueryExpertiseUser queries the ExpertiseUser edge of a Expertise.
func (c *ExpertiseClient) QueryExpertiseUser(e *Expertise) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(expertise.Table, expertise.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, expertise.ExpertiseUserTable, expertise.ExpertiseUserColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ExpertiseClient) Hooks() []Hook {
	return c.hooks.Expertise
}

// PositionClient is a client for the Position schema.
type PositionClient struct {
	config
}

// NewPositionClient returns a client for the Position from the given config.
func NewPositionClient(c config) *PositionClient {
	return &PositionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `position.Hooks(f(g(h())))`.
func (c *PositionClient) Use(hooks ...Hook) {
	c.hooks.Position = append(c.hooks.Position, hooks...)
}

// Create returns a create builder for Position.
func (c *PositionClient) Create() *PositionCreate {
	mutation := newPositionMutation(c.config, OpCreate)
	return &PositionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Position.
func (c *PositionClient) Update() *PositionUpdate {
	mutation := newPositionMutation(c.config, OpUpdate)
	return &PositionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PositionClient) UpdateOne(po *Position) *PositionUpdateOne {
	mutation := newPositionMutation(c.config, OpUpdateOne, withPosition(po))
	return &PositionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PositionClient) UpdateOneID(id int) *PositionUpdateOne {
	mutation := newPositionMutation(c.config, OpUpdateOne, withPositionID(id))
	return &PositionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Position.
func (c *PositionClient) Delete() *PositionDelete {
	mutation := newPositionMutation(c.config, OpDelete)
	return &PositionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PositionClient) DeleteOne(po *Position) *PositionDeleteOne {
	return c.DeleteOneID(po.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PositionClient) DeleteOneID(id int) *PositionDeleteOne {
	builder := c.Delete().Where(position.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PositionDeleteOne{builder}
}

// Create returns a query builder for Position.
func (c *PositionClient) Query() *PositionQuery {
	return &PositionQuery{config: c.config}
}

// Get returns a Position entity by its id.
func (c *PositionClient) Get(ctx context.Context, id int) (*Position, error) {
	return c.Query().Where(position.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PositionClient) GetX(ctx context.Context, id int) *Position {
	po, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return po
}

// QueryPositionUser queries the PositionUser edge of a Position.
func (c *PositionClient) QueryPositionUser(po *Position) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(position.Table, position.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, position.PositionUserTable, position.PositionUserColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PositionClient) Hooks() []Hook {
	return c.hooks.Position
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Create returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// QueryUserDepartment queries the UserDepartment edge of a User.
func (c *UserClient) QueryUserDepartment(u *User) *DepartmentQuery {
	query := &DepartmentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(department.Table, department.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, user.UserDepartmentTable, user.UserDepartmentColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUserExpertise queries the UserExpertise edge of a User.
func (c *UserClient) QueryUserExpertise(u *User) *ExpertiseQuery {
	query := &ExpertiseQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(expertise.Table, expertise.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, user.UserExpertiseTable, user.UserExpertiseColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUserPosition queries the UserPosition edge of a User.
func (c *UserClient) QueryUserPosition(u *User) *PositionQuery {
	query := &PositionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(position.Table, position.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, user.UserPositionTable, user.UserPositionColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
