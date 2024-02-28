// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/BlindGarret/movie-queue/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/BlindGarret/movie-queue/ent/rank"
	"github.com/BlindGarret/movie-queue/ent/show"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Rank is the client for interacting with the Rank builders.
	Rank *RankClient
	// Show is the client for interacting with the Show builders.
	Show *ShowClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Rank = NewRankClient(c.config)
	c.Show = NewShowClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Rank:   NewRankClient(cfg),
		Show:   NewShowClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Rank:   NewRankClient(cfg),
		Show:   NewShowClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Rank.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
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
	c.Rank.Use(hooks...)
	c.Show.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Rank.Intercept(interceptors...)
	c.Show.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *RankMutation:
		return c.Rank.mutate(ctx, m)
	case *ShowMutation:
		return c.Show.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// RankClient is a client for the Rank schema.
type RankClient struct {
	config
}

// NewRankClient returns a client for the Rank from the given config.
func NewRankClient(c config) *RankClient {
	return &RankClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `rank.Hooks(f(g(h())))`.
func (c *RankClient) Use(hooks ...Hook) {
	c.hooks.Rank = append(c.hooks.Rank, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `rank.Intercept(f(g(h())))`.
func (c *RankClient) Intercept(interceptors ...Interceptor) {
	c.inters.Rank = append(c.inters.Rank, interceptors...)
}

// Create returns a builder for creating a Rank entity.
func (c *RankClient) Create() *RankCreate {
	mutation := newRankMutation(c.config, OpCreate)
	return &RankCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Rank entities.
func (c *RankClient) CreateBulk(builders ...*RankCreate) *RankCreateBulk {
	return &RankCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *RankClient) MapCreateBulk(slice any, setFunc func(*RankCreate, int)) *RankCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &RankCreateBulk{err: fmt.Errorf("calling to RankClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*RankCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &RankCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Rank.
func (c *RankClient) Update() *RankUpdate {
	mutation := newRankMutation(c.config, OpUpdate)
	return &RankUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RankClient) UpdateOne(r *Rank) *RankUpdateOne {
	mutation := newRankMutation(c.config, OpUpdateOne, withRank(r))
	return &RankUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RankClient) UpdateOneID(id int) *RankUpdateOne {
	mutation := newRankMutation(c.config, OpUpdateOne, withRankID(id))
	return &RankUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Rank.
func (c *RankClient) Delete() *RankDelete {
	mutation := newRankMutation(c.config, OpDelete)
	return &RankDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RankClient) DeleteOne(r *Rank) *RankDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RankClient) DeleteOneID(id int) *RankDeleteOne {
	builder := c.Delete().Where(rank.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RankDeleteOne{builder}
}

// Query returns a query builder for Rank.
func (c *RankClient) Query() *RankQuery {
	return &RankQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRank},
		inters: c.Interceptors(),
	}
}

// Get returns a Rank entity by its id.
func (c *RankClient) Get(ctx context.Context, id int) (*Rank, error) {
	return c.Query().Where(rank.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RankClient) GetX(ctx context.Context, id int) *Rank {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *RankClient) Hooks() []Hook {
	return c.hooks.Rank
}

// Interceptors returns the client interceptors.
func (c *RankClient) Interceptors() []Interceptor {
	return c.inters.Rank
}

func (c *RankClient) mutate(ctx context.Context, m *RankMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RankCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RankUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RankUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RankDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Rank mutation op: %q", m.Op())
	}
}

// ShowClient is a client for the Show schema.
type ShowClient struct {
	config
}

// NewShowClient returns a client for the Show from the given config.
func NewShowClient(c config) *ShowClient {
	return &ShowClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `show.Hooks(f(g(h())))`.
func (c *ShowClient) Use(hooks ...Hook) {
	c.hooks.Show = append(c.hooks.Show, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `show.Intercept(f(g(h())))`.
func (c *ShowClient) Intercept(interceptors ...Interceptor) {
	c.inters.Show = append(c.inters.Show, interceptors...)
}

// Create returns a builder for creating a Show entity.
func (c *ShowClient) Create() *ShowCreate {
	mutation := newShowMutation(c.config, OpCreate)
	return &ShowCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Show entities.
func (c *ShowClient) CreateBulk(builders ...*ShowCreate) *ShowCreateBulk {
	return &ShowCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ShowClient) MapCreateBulk(slice any, setFunc func(*ShowCreate, int)) *ShowCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ShowCreateBulk{err: fmt.Errorf("calling to ShowClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ShowCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ShowCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Show.
func (c *ShowClient) Update() *ShowUpdate {
	mutation := newShowMutation(c.config, OpUpdate)
	return &ShowUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ShowClient) UpdateOne(s *Show) *ShowUpdateOne {
	mutation := newShowMutation(c.config, OpUpdateOne, withShow(s))
	return &ShowUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ShowClient) UpdateOneID(id int) *ShowUpdateOne {
	mutation := newShowMutation(c.config, OpUpdateOne, withShowID(id))
	return &ShowUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Show.
func (c *ShowClient) Delete() *ShowDelete {
	mutation := newShowMutation(c.config, OpDelete)
	return &ShowDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ShowClient) DeleteOne(s *Show) *ShowDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ShowClient) DeleteOneID(id int) *ShowDeleteOne {
	builder := c.Delete().Where(show.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ShowDeleteOne{builder}
}

// Query returns a query builder for Show.
func (c *ShowClient) Query() *ShowQuery {
	return &ShowQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeShow},
		inters: c.Interceptors(),
	}
}

// Get returns a Show entity by its id.
func (c *ShowClient) Get(ctx context.Context, id int) (*Show, error) {
	return c.Query().Where(show.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ShowClient) GetX(ctx context.Context, id int) *Show {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ShowClient) Hooks() []Hook {
	return c.hooks.Show
}

// Interceptors returns the client interceptors.
func (c *ShowClient) Interceptors() []Interceptor {
	return c.inters.Show
}

func (c *ShowClient) mutate(ctx context.Context, m *ShowMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ShowCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ShowUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ShowUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ShowDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Show mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Rank, Show []ent.Hook
	}
	inters struct {
		Rank, Show []ent.Interceptor
	}
)
