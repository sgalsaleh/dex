// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sgalsaleh/dex/v2/storage/ent/db/predicate"
	"github.com/sgalsaleh/dex/v2/storage/ent/db/refreshtoken"
)

// RefreshTokenQuery is the builder for querying RefreshToken entities.
type RefreshTokenQuery struct {
	config
	ctx        *QueryContext
	order      []refreshtoken.OrderOption
	inters     []Interceptor
	predicates []predicate.RefreshToken
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RefreshTokenQuery builder.
func (rtq *RefreshTokenQuery) Where(ps ...predicate.RefreshToken) *RefreshTokenQuery {
	rtq.predicates = append(rtq.predicates, ps...)
	return rtq
}

// Limit the number of records to be returned by this query.
func (rtq *RefreshTokenQuery) Limit(limit int) *RefreshTokenQuery {
	rtq.ctx.Limit = &limit
	return rtq
}

// Offset to start from.
func (rtq *RefreshTokenQuery) Offset(offset int) *RefreshTokenQuery {
	rtq.ctx.Offset = &offset
	return rtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rtq *RefreshTokenQuery) Unique(unique bool) *RefreshTokenQuery {
	rtq.ctx.Unique = &unique
	return rtq
}

// Order specifies how the records should be ordered.
func (rtq *RefreshTokenQuery) Order(o ...refreshtoken.OrderOption) *RefreshTokenQuery {
	rtq.order = append(rtq.order, o...)
	return rtq
}

// First returns the first RefreshToken entity from the query.
// Returns a *NotFoundError when no RefreshToken was found.
func (rtq *RefreshTokenQuery) First(ctx context.Context) (*RefreshToken, error) {
	nodes, err := rtq.Limit(1).All(setContextOp(ctx, rtq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{refreshtoken.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rtq *RefreshTokenQuery) FirstX(ctx context.Context) *RefreshToken {
	node, err := rtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RefreshToken ID from the query.
// Returns a *NotFoundError when no RefreshToken ID was found.
func (rtq *RefreshTokenQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = rtq.Limit(1).IDs(setContextOp(ctx, rtq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{refreshtoken.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rtq *RefreshTokenQuery) FirstIDX(ctx context.Context) string {
	id, err := rtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RefreshToken entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RefreshToken entity is found.
// Returns a *NotFoundError when no RefreshToken entities are found.
func (rtq *RefreshTokenQuery) Only(ctx context.Context) (*RefreshToken, error) {
	nodes, err := rtq.Limit(2).All(setContextOp(ctx, rtq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{refreshtoken.Label}
	default:
		return nil, &NotSingularError{refreshtoken.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rtq *RefreshTokenQuery) OnlyX(ctx context.Context) *RefreshToken {
	node, err := rtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RefreshToken ID in the query.
// Returns a *NotSingularError when more than one RefreshToken ID is found.
// Returns a *NotFoundError when no entities are found.
func (rtq *RefreshTokenQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = rtq.Limit(2).IDs(setContextOp(ctx, rtq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{refreshtoken.Label}
	default:
		err = &NotSingularError{refreshtoken.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rtq *RefreshTokenQuery) OnlyIDX(ctx context.Context) string {
	id, err := rtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RefreshTokens.
func (rtq *RefreshTokenQuery) All(ctx context.Context) ([]*RefreshToken, error) {
	ctx = setContextOp(ctx, rtq.ctx, "All")
	if err := rtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RefreshToken, *RefreshTokenQuery]()
	return withInterceptors[[]*RefreshToken](ctx, rtq, qr, rtq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rtq *RefreshTokenQuery) AllX(ctx context.Context) []*RefreshToken {
	nodes, err := rtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RefreshToken IDs.
func (rtq *RefreshTokenQuery) IDs(ctx context.Context) (ids []string, err error) {
	if rtq.ctx.Unique == nil && rtq.path != nil {
		rtq.Unique(true)
	}
	ctx = setContextOp(ctx, rtq.ctx, "IDs")
	if err = rtq.Select(refreshtoken.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rtq *RefreshTokenQuery) IDsX(ctx context.Context) []string {
	ids, err := rtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rtq *RefreshTokenQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rtq.ctx, "Count")
	if err := rtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rtq, querierCount[*RefreshTokenQuery](), rtq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rtq *RefreshTokenQuery) CountX(ctx context.Context) int {
	count, err := rtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rtq *RefreshTokenQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rtq.ctx, "Exist")
	switch _, err := rtq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rtq *RefreshTokenQuery) ExistX(ctx context.Context) bool {
	exist, err := rtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RefreshTokenQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rtq *RefreshTokenQuery) Clone() *RefreshTokenQuery {
	if rtq == nil {
		return nil
	}
	return &RefreshTokenQuery{
		config:     rtq.config,
		ctx:        rtq.ctx.Clone(),
		order:      append([]refreshtoken.OrderOption{}, rtq.order...),
		inters:     append([]Interceptor{}, rtq.inters...),
		predicates: append([]predicate.RefreshToken{}, rtq.predicates...),
		// clone intermediate query.
		sql:  rtq.sql.Clone(),
		path: rtq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ClientID string `json:"client_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RefreshToken.Query().
//		GroupBy(refreshtoken.FieldClientID).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (rtq *RefreshTokenQuery) GroupBy(field string, fields ...string) *RefreshTokenGroupBy {
	rtq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RefreshTokenGroupBy{build: rtq}
	grbuild.flds = &rtq.ctx.Fields
	grbuild.label = refreshtoken.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ClientID string `json:"client_id,omitempty"`
//	}
//
//	client.RefreshToken.Query().
//		Select(refreshtoken.FieldClientID).
//		Scan(ctx, &v)
func (rtq *RefreshTokenQuery) Select(fields ...string) *RefreshTokenSelect {
	rtq.ctx.Fields = append(rtq.ctx.Fields, fields...)
	sbuild := &RefreshTokenSelect{RefreshTokenQuery: rtq}
	sbuild.label = refreshtoken.Label
	sbuild.flds, sbuild.scan = &rtq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RefreshTokenSelect configured with the given aggregations.
func (rtq *RefreshTokenQuery) Aggregate(fns ...AggregateFunc) *RefreshTokenSelect {
	return rtq.Select().Aggregate(fns...)
}

func (rtq *RefreshTokenQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rtq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rtq); err != nil {
				return err
			}
		}
	}
	for _, f := range rtq.ctx.Fields {
		if !refreshtoken.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if rtq.path != nil {
		prev, err := rtq.path(ctx)
		if err != nil {
			return err
		}
		rtq.sql = prev
	}
	return nil
}

func (rtq *RefreshTokenQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RefreshToken, error) {
	var (
		nodes = []*RefreshToken{}
		_spec = rtq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RefreshToken).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RefreshToken{config: rtq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (rtq *RefreshTokenQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rtq.querySpec()
	_spec.Node.Columns = rtq.ctx.Fields
	if len(rtq.ctx.Fields) > 0 {
		_spec.Unique = rtq.ctx.Unique != nil && *rtq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rtq.driver, _spec)
}

func (rtq *RefreshTokenQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(refreshtoken.Table, refreshtoken.Columns, sqlgraph.NewFieldSpec(refreshtoken.FieldID, field.TypeString))
	_spec.From = rtq.sql
	if unique := rtq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rtq.path != nil {
		_spec.Unique = true
	}
	if fields := rtq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, refreshtoken.FieldID)
		for i := range fields {
			if fields[i] != refreshtoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rtq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rtq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rtq *RefreshTokenQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rtq.driver.Dialect())
	t1 := builder.Table(refreshtoken.Table)
	columns := rtq.ctx.Fields
	if len(columns) == 0 {
		columns = refreshtoken.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rtq.sql != nil {
		selector = rtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rtq.ctx.Unique != nil && *rtq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rtq.predicates {
		p(selector)
	}
	for _, p := range rtq.order {
		p(selector)
	}
	if offset := rtq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rtq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RefreshTokenGroupBy is the group-by builder for RefreshToken entities.
type RefreshTokenGroupBy struct {
	selector
	build *RefreshTokenQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rtgb *RefreshTokenGroupBy) Aggregate(fns ...AggregateFunc) *RefreshTokenGroupBy {
	rtgb.fns = append(rtgb.fns, fns...)
	return rtgb
}

// Scan applies the selector query and scans the result into the given value.
func (rtgb *RefreshTokenGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rtgb.build.ctx, "GroupBy")
	if err := rtgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RefreshTokenQuery, *RefreshTokenGroupBy](ctx, rtgb.build, rtgb, rtgb.build.inters, v)
}

func (rtgb *RefreshTokenGroupBy) sqlScan(ctx context.Context, root *RefreshTokenQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rtgb.fns))
	for _, fn := range rtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rtgb.flds)+len(rtgb.fns))
		for _, f := range *rtgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rtgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rtgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RefreshTokenSelect is the builder for selecting fields of RefreshToken entities.
type RefreshTokenSelect struct {
	*RefreshTokenQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rts *RefreshTokenSelect) Aggregate(fns ...AggregateFunc) *RefreshTokenSelect {
	rts.fns = append(rts.fns, fns...)
	return rts
}

// Scan applies the selector query and scans the result into the given value.
func (rts *RefreshTokenSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rts.ctx, "Select")
	if err := rts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RefreshTokenQuery, *RefreshTokenSelect](ctx, rts.RefreshTokenQuery, rts, rts.inters, v)
}

func (rts *RefreshTokenSelect) sqlScan(ctx context.Context, root *RefreshTokenQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rts.fns))
	for _, fn := range rts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
