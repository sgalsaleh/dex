// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sgalsaleh/dex/v2/storage/ent/db/devicetoken"
	"github.com/sgalsaleh/dex/v2/storage/ent/db/predicate"
)

// DeviceTokenQuery is the builder for querying DeviceToken entities.
type DeviceTokenQuery struct {
	config
	ctx        *QueryContext
	order      []devicetoken.OrderOption
	inters     []Interceptor
	predicates []predicate.DeviceToken
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DeviceTokenQuery builder.
func (dtq *DeviceTokenQuery) Where(ps ...predicate.DeviceToken) *DeviceTokenQuery {
	dtq.predicates = append(dtq.predicates, ps...)
	return dtq
}

// Limit the number of records to be returned by this query.
func (dtq *DeviceTokenQuery) Limit(limit int) *DeviceTokenQuery {
	dtq.ctx.Limit = &limit
	return dtq
}

// Offset to start from.
func (dtq *DeviceTokenQuery) Offset(offset int) *DeviceTokenQuery {
	dtq.ctx.Offset = &offset
	return dtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dtq *DeviceTokenQuery) Unique(unique bool) *DeviceTokenQuery {
	dtq.ctx.Unique = &unique
	return dtq
}

// Order specifies how the records should be ordered.
func (dtq *DeviceTokenQuery) Order(o ...devicetoken.OrderOption) *DeviceTokenQuery {
	dtq.order = append(dtq.order, o...)
	return dtq
}

// First returns the first DeviceToken entity from the query.
// Returns a *NotFoundError when no DeviceToken was found.
func (dtq *DeviceTokenQuery) First(ctx context.Context) (*DeviceToken, error) {
	nodes, err := dtq.Limit(1).All(setContextOp(ctx, dtq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{devicetoken.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dtq *DeviceTokenQuery) FirstX(ctx context.Context) *DeviceToken {
	node, err := dtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DeviceToken ID from the query.
// Returns a *NotFoundError when no DeviceToken ID was found.
func (dtq *DeviceTokenQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dtq.Limit(1).IDs(setContextOp(ctx, dtq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{devicetoken.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dtq *DeviceTokenQuery) FirstIDX(ctx context.Context) int {
	id, err := dtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DeviceToken entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DeviceToken entity is found.
// Returns a *NotFoundError when no DeviceToken entities are found.
func (dtq *DeviceTokenQuery) Only(ctx context.Context) (*DeviceToken, error) {
	nodes, err := dtq.Limit(2).All(setContextOp(ctx, dtq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{devicetoken.Label}
	default:
		return nil, &NotSingularError{devicetoken.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dtq *DeviceTokenQuery) OnlyX(ctx context.Context) *DeviceToken {
	node, err := dtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DeviceToken ID in the query.
// Returns a *NotSingularError when more than one DeviceToken ID is found.
// Returns a *NotFoundError when no entities are found.
func (dtq *DeviceTokenQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dtq.Limit(2).IDs(setContextOp(ctx, dtq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{devicetoken.Label}
	default:
		err = &NotSingularError{devicetoken.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dtq *DeviceTokenQuery) OnlyIDX(ctx context.Context) int {
	id, err := dtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DeviceTokens.
func (dtq *DeviceTokenQuery) All(ctx context.Context) ([]*DeviceToken, error) {
	ctx = setContextOp(ctx, dtq.ctx, "All")
	if err := dtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DeviceToken, *DeviceTokenQuery]()
	return withInterceptors[[]*DeviceToken](ctx, dtq, qr, dtq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dtq *DeviceTokenQuery) AllX(ctx context.Context) []*DeviceToken {
	nodes, err := dtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DeviceToken IDs.
func (dtq *DeviceTokenQuery) IDs(ctx context.Context) (ids []int, err error) {
	if dtq.ctx.Unique == nil && dtq.path != nil {
		dtq.Unique(true)
	}
	ctx = setContextOp(ctx, dtq.ctx, "IDs")
	if err = dtq.Select(devicetoken.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dtq *DeviceTokenQuery) IDsX(ctx context.Context) []int {
	ids, err := dtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dtq *DeviceTokenQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dtq.ctx, "Count")
	if err := dtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dtq, querierCount[*DeviceTokenQuery](), dtq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dtq *DeviceTokenQuery) CountX(ctx context.Context) int {
	count, err := dtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dtq *DeviceTokenQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dtq.ctx, "Exist")
	switch _, err := dtq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dtq *DeviceTokenQuery) ExistX(ctx context.Context) bool {
	exist, err := dtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DeviceTokenQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dtq *DeviceTokenQuery) Clone() *DeviceTokenQuery {
	if dtq == nil {
		return nil
	}
	return &DeviceTokenQuery{
		config:     dtq.config,
		ctx:        dtq.ctx.Clone(),
		order:      append([]devicetoken.OrderOption{}, dtq.order...),
		inters:     append([]Interceptor{}, dtq.inters...),
		predicates: append([]predicate.DeviceToken{}, dtq.predicates...),
		// clone intermediate query.
		sql:  dtq.sql.Clone(),
		path: dtq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DeviceCode string `json:"device_code,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DeviceToken.Query().
//		GroupBy(devicetoken.FieldDeviceCode).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (dtq *DeviceTokenQuery) GroupBy(field string, fields ...string) *DeviceTokenGroupBy {
	dtq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DeviceTokenGroupBy{build: dtq}
	grbuild.flds = &dtq.ctx.Fields
	grbuild.label = devicetoken.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DeviceCode string `json:"device_code,omitempty"`
//	}
//
//	client.DeviceToken.Query().
//		Select(devicetoken.FieldDeviceCode).
//		Scan(ctx, &v)
func (dtq *DeviceTokenQuery) Select(fields ...string) *DeviceTokenSelect {
	dtq.ctx.Fields = append(dtq.ctx.Fields, fields...)
	sbuild := &DeviceTokenSelect{DeviceTokenQuery: dtq}
	sbuild.label = devicetoken.Label
	sbuild.flds, sbuild.scan = &dtq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DeviceTokenSelect configured with the given aggregations.
func (dtq *DeviceTokenQuery) Aggregate(fns ...AggregateFunc) *DeviceTokenSelect {
	return dtq.Select().Aggregate(fns...)
}

func (dtq *DeviceTokenQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dtq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dtq); err != nil {
				return err
			}
		}
	}
	for _, f := range dtq.ctx.Fields {
		if !devicetoken.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if dtq.path != nil {
		prev, err := dtq.path(ctx)
		if err != nil {
			return err
		}
		dtq.sql = prev
	}
	return nil
}

func (dtq *DeviceTokenQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DeviceToken, error) {
	var (
		nodes = []*DeviceToken{}
		_spec = dtq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DeviceToken).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DeviceToken{config: dtq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (dtq *DeviceTokenQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dtq.querySpec()
	_spec.Node.Columns = dtq.ctx.Fields
	if len(dtq.ctx.Fields) > 0 {
		_spec.Unique = dtq.ctx.Unique != nil && *dtq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dtq.driver, _spec)
}

func (dtq *DeviceTokenQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(devicetoken.Table, devicetoken.Columns, sqlgraph.NewFieldSpec(devicetoken.FieldID, field.TypeInt))
	_spec.From = dtq.sql
	if unique := dtq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dtq.path != nil {
		_spec.Unique = true
	}
	if fields := dtq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, devicetoken.FieldID)
		for i := range fields {
			if fields[i] != devicetoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dtq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dtq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dtq *DeviceTokenQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dtq.driver.Dialect())
	t1 := builder.Table(devicetoken.Table)
	columns := dtq.ctx.Fields
	if len(columns) == 0 {
		columns = devicetoken.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dtq.sql != nil {
		selector = dtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dtq.ctx.Unique != nil && *dtq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dtq.predicates {
		p(selector)
	}
	for _, p := range dtq.order {
		p(selector)
	}
	if offset := dtq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dtq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DeviceTokenGroupBy is the group-by builder for DeviceToken entities.
type DeviceTokenGroupBy struct {
	selector
	build *DeviceTokenQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dtgb *DeviceTokenGroupBy) Aggregate(fns ...AggregateFunc) *DeviceTokenGroupBy {
	dtgb.fns = append(dtgb.fns, fns...)
	return dtgb
}

// Scan applies the selector query and scans the result into the given value.
func (dtgb *DeviceTokenGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dtgb.build.ctx, "GroupBy")
	if err := dtgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DeviceTokenQuery, *DeviceTokenGroupBy](ctx, dtgb.build, dtgb, dtgb.build.inters, v)
}

func (dtgb *DeviceTokenGroupBy) sqlScan(ctx context.Context, root *DeviceTokenQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dtgb.fns))
	for _, fn := range dtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dtgb.flds)+len(dtgb.fns))
		for _, f := range *dtgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dtgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dtgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DeviceTokenSelect is the builder for selecting fields of DeviceToken entities.
type DeviceTokenSelect struct {
	*DeviceTokenQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (dts *DeviceTokenSelect) Aggregate(fns ...AggregateFunc) *DeviceTokenSelect {
	dts.fns = append(dts.fns, fns...)
	return dts
}

// Scan applies the selector query and scans the result into the given value.
func (dts *DeviceTokenSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dts.ctx, "Select")
	if err := dts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DeviceTokenQuery, *DeviceTokenSelect](ctx, dts.DeviceTokenQuery, dts, dts.inters, v)
}

func (dts *DeviceTokenSelect) sqlScan(ctx context.Context, root *DeviceTokenQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(dts.fns))
	for _, fn := range dts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*dts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
