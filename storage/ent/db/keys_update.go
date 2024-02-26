// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/sgalsaleh/dex/v2/storage"
	"github.com/sgalsaleh/dex/v2/storage/ent/db/keys"
	"github.com/sgalsaleh/dex/v2/storage/ent/db/predicate"
	jose "gopkg.in/square/go-jose.v2"
)

// KeysUpdate is the builder for updating Keys entities.
type KeysUpdate struct {
	config
	hooks    []Hook
	mutation *KeysMutation
}

// Where appends a list predicates to the KeysUpdate builder.
func (ku *KeysUpdate) Where(ps ...predicate.Keys) *KeysUpdate {
	ku.mutation.Where(ps...)
	return ku
}

// SetVerificationKeys sets the "verification_keys" field.
func (ku *KeysUpdate) SetVerificationKeys(sk []storage.VerificationKey) *KeysUpdate {
	ku.mutation.SetVerificationKeys(sk)
	return ku
}

// AppendVerificationKeys appends sk to the "verification_keys" field.
func (ku *KeysUpdate) AppendVerificationKeys(sk []storage.VerificationKey) *KeysUpdate {
	ku.mutation.AppendVerificationKeys(sk)
	return ku
}

// SetSigningKey sets the "signing_key" field.
func (ku *KeysUpdate) SetSigningKey(jwk jose.JSONWebKey) *KeysUpdate {
	ku.mutation.SetSigningKey(jwk)
	return ku
}

// SetSigningKeyPub sets the "signing_key_pub" field.
func (ku *KeysUpdate) SetSigningKeyPub(jwk jose.JSONWebKey) *KeysUpdate {
	ku.mutation.SetSigningKeyPub(jwk)
	return ku
}

// SetNextRotation sets the "next_rotation" field.
func (ku *KeysUpdate) SetNextRotation(t time.Time) *KeysUpdate {
	ku.mutation.SetNextRotation(t)
	return ku
}

// Mutation returns the KeysMutation object of the builder.
func (ku *KeysUpdate) Mutation() *KeysMutation {
	return ku.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ku *KeysUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ku.sqlSave, ku.mutation, ku.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ku *KeysUpdate) SaveX(ctx context.Context) int {
	affected, err := ku.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ku *KeysUpdate) Exec(ctx context.Context) error {
	_, err := ku.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ku *KeysUpdate) ExecX(ctx context.Context) {
	if err := ku.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ku *KeysUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(keys.Table, keys.Columns, sqlgraph.NewFieldSpec(keys.FieldID, field.TypeString))
	if ps := ku.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ku.mutation.VerificationKeys(); ok {
		_spec.SetField(keys.FieldVerificationKeys, field.TypeJSON, value)
	}
	if value, ok := ku.mutation.AppendedVerificationKeys(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, keys.FieldVerificationKeys, value)
		})
	}
	if value, ok := ku.mutation.SigningKey(); ok {
		_spec.SetField(keys.FieldSigningKey, field.TypeJSON, value)
	}
	if value, ok := ku.mutation.SigningKeyPub(); ok {
		_spec.SetField(keys.FieldSigningKeyPub, field.TypeJSON, value)
	}
	if value, ok := ku.mutation.NextRotation(); ok {
		_spec.SetField(keys.FieldNextRotation, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ku.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{keys.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ku.mutation.done = true
	return n, nil
}

// KeysUpdateOne is the builder for updating a single Keys entity.
type KeysUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *KeysMutation
}

// SetVerificationKeys sets the "verification_keys" field.
func (kuo *KeysUpdateOne) SetVerificationKeys(sk []storage.VerificationKey) *KeysUpdateOne {
	kuo.mutation.SetVerificationKeys(sk)
	return kuo
}

// AppendVerificationKeys appends sk to the "verification_keys" field.
func (kuo *KeysUpdateOne) AppendVerificationKeys(sk []storage.VerificationKey) *KeysUpdateOne {
	kuo.mutation.AppendVerificationKeys(sk)
	return kuo
}

// SetSigningKey sets the "signing_key" field.
func (kuo *KeysUpdateOne) SetSigningKey(jwk jose.JSONWebKey) *KeysUpdateOne {
	kuo.mutation.SetSigningKey(jwk)
	return kuo
}

// SetSigningKeyPub sets the "signing_key_pub" field.
func (kuo *KeysUpdateOne) SetSigningKeyPub(jwk jose.JSONWebKey) *KeysUpdateOne {
	kuo.mutation.SetSigningKeyPub(jwk)
	return kuo
}

// SetNextRotation sets the "next_rotation" field.
func (kuo *KeysUpdateOne) SetNextRotation(t time.Time) *KeysUpdateOne {
	kuo.mutation.SetNextRotation(t)
	return kuo
}

// Mutation returns the KeysMutation object of the builder.
func (kuo *KeysUpdateOne) Mutation() *KeysMutation {
	return kuo.mutation
}

// Where appends a list predicates to the KeysUpdate builder.
func (kuo *KeysUpdateOne) Where(ps ...predicate.Keys) *KeysUpdateOne {
	kuo.mutation.Where(ps...)
	return kuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (kuo *KeysUpdateOne) Select(field string, fields ...string) *KeysUpdateOne {
	kuo.fields = append([]string{field}, fields...)
	return kuo
}

// Save executes the query and returns the updated Keys entity.
func (kuo *KeysUpdateOne) Save(ctx context.Context) (*Keys, error) {
	return withHooks(ctx, kuo.sqlSave, kuo.mutation, kuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (kuo *KeysUpdateOne) SaveX(ctx context.Context) *Keys {
	node, err := kuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (kuo *KeysUpdateOne) Exec(ctx context.Context) error {
	_, err := kuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kuo *KeysUpdateOne) ExecX(ctx context.Context) {
	if err := kuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (kuo *KeysUpdateOne) sqlSave(ctx context.Context) (_node *Keys, err error) {
	_spec := sqlgraph.NewUpdateSpec(keys.Table, keys.Columns, sqlgraph.NewFieldSpec(keys.FieldID, field.TypeString))
	id, ok := kuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "Keys.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := kuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, keys.FieldID)
		for _, f := range fields {
			if !keys.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != keys.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := kuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kuo.mutation.VerificationKeys(); ok {
		_spec.SetField(keys.FieldVerificationKeys, field.TypeJSON, value)
	}
	if value, ok := kuo.mutation.AppendedVerificationKeys(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, keys.FieldVerificationKeys, value)
		})
	}
	if value, ok := kuo.mutation.SigningKey(); ok {
		_spec.SetField(keys.FieldSigningKey, field.TypeJSON, value)
	}
	if value, ok := kuo.mutation.SigningKeyPub(); ok {
		_spec.SetField(keys.FieldSigningKeyPub, field.TypeJSON, value)
	}
	if value, ok := kuo.mutation.NextRotation(); ok {
		_spec.SetField(keys.FieldNextRotation, field.TypeTime, value)
	}
	_node = &Keys{config: kuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, kuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{keys.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	kuo.mutation.done = true
	return _node, nil
}
