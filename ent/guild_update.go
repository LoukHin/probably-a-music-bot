// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"definitelynotmusicbot/ent/guild"
	"definitelynotmusicbot/ent/predicate"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	snowflake "github.com/disgoorg/snowflake/v2"
)

// GuildUpdate is the builder for updating Guild entities.
type GuildUpdate struct {
	config
	hooks    []Hook
	mutation *GuildMutation
}

// Where appends a list predicates to the GuildUpdate builder.
func (gu *GuildUpdate) Where(ps ...predicate.Guild) *GuildUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetName sets the "name" field.
func (gu *GuildUpdate) SetName(s string) *GuildUpdate {
	gu.mutation.SetName(s)
	return gu
}

// SetPlayerChannelID sets the "player_channel_id" field.
func (gu *GuildUpdate) SetPlayerChannelID(s snowflake.ID) *GuildUpdate {
	gu.mutation.ResetPlayerChannelID()
	gu.mutation.SetPlayerChannelID(s)
	return gu
}

// SetNillablePlayerChannelID sets the "player_channel_id" field if the given value is not nil.
func (gu *GuildUpdate) SetNillablePlayerChannelID(s *snowflake.ID) *GuildUpdate {
	if s != nil {
		gu.SetPlayerChannelID(*s)
	}
	return gu
}

// AddPlayerChannelID adds s to the "player_channel_id" field.
func (gu *GuildUpdate) AddPlayerChannelID(s snowflake.ID) *GuildUpdate {
	gu.mutation.AddPlayerChannelID(s)
	return gu
}

// ClearPlayerChannelID clears the value of the "player_channel_id" field.
func (gu *GuildUpdate) ClearPlayerChannelID() *GuildUpdate {
	gu.mutation.ClearPlayerChannelID()
	return gu
}

// SetPlayerMessageID sets the "player_message_id" field.
func (gu *GuildUpdate) SetPlayerMessageID(s snowflake.ID) *GuildUpdate {
	gu.mutation.ResetPlayerMessageID()
	gu.mutation.SetPlayerMessageID(s)
	return gu
}

// SetNillablePlayerMessageID sets the "player_message_id" field if the given value is not nil.
func (gu *GuildUpdate) SetNillablePlayerMessageID(s *snowflake.ID) *GuildUpdate {
	if s != nil {
		gu.SetPlayerMessageID(*s)
	}
	return gu
}

// AddPlayerMessageID adds s to the "player_message_id" field.
func (gu *GuildUpdate) AddPlayerMessageID(s snowflake.ID) *GuildUpdate {
	gu.mutation.AddPlayerMessageID(s)
	return gu
}

// ClearPlayerMessageID clears the value of the "player_message_id" field.
func (gu *GuildUpdate) ClearPlayerMessageID() *GuildUpdate {
	gu.mutation.ClearPlayerMessageID()
	return gu
}

// SetCreatedAt sets the "created_at" field.
func (gu *GuildUpdate) SetCreatedAt(t time.Time) *GuildUpdate {
	gu.mutation.SetCreatedAt(t)
	return gu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gu *GuildUpdate) SetNillableCreatedAt(t *time.Time) *GuildUpdate {
	if t != nil {
		gu.SetCreatedAt(*t)
	}
	return gu
}

// SetUpdatedAt sets the "updated_at" field.
func (gu *GuildUpdate) SetUpdatedAt(t time.Time) *GuildUpdate {
	gu.mutation.SetUpdatedAt(t)
	return gu
}

// Mutation returns the GuildMutation object of the builder.
func (gu *GuildUpdate) Mutation() *GuildMutation {
	return gu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GuildUpdate) Save(ctx context.Context) (int, error) {
	gu.defaults()
	return withHooks[int, GuildMutation](ctx, gu.sqlSave, gu.mutation, gu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GuildUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GuildUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GuildUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gu *GuildUpdate) defaults() {
	if _, ok := gu.mutation.UpdatedAt(); !ok {
		v := guild.UpdateDefaultUpdatedAt()
		gu.mutation.SetUpdatedAt(v)
	}
}

func (gu *GuildUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(guild.Table, guild.Columns, sqlgraph.NewFieldSpec(guild.FieldID, field.TypeUint64))
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.SetField(guild.FieldName, field.TypeString, value)
	}
	if value, ok := gu.mutation.PlayerChannelID(); ok {
		_spec.SetField(guild.FieldPlayerChannelID, field.TypeUint64, value)
	}
	if value, ok := gu.mutation.AddedPlayerChannelID(); ok {
		_spec.AddField(guild.FieldPlayerChannelID, field.TypeUint64, value)
	}
	if gu.mutation.PlayerChannelIDCleared() {
		_spec.ClearField(guild.FieldPlayerChannelID, field.TypeUint64)
	}
	if value, ok := gu.mutation.PlayerMessageID(); ok {
		_spec.SetField(guild.FieldPlayerMessageID, field.TypeUint64, value)
	}
	if value, ok := gu.mutation.AddedPlayerMessageID(); ok {
		_spec.AddField(guild.FieldPlayerMessageID, field.TypeUint64, value)
	}
	if gu.mutation.PlayerMessageIDCleared() {
		_spec.ClearField(guild.FieldPlayerMessageID, field.TypeUint64)
	}
	if value, ok := gu.mutation.CreatedAt(); ok {
		_spec.SetField(guild.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := gu.mutation.UpdatedAt(); ok {
		_spec.SetField(guild.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{guild.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gu.mutation.done = true
	return n, nil
}

// GuildUpdateOne is the builder for updating a single Guild entity.
type GuildUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GuildMutation
}

// SetName sets the "name" field.
func (guo *GuildUpdateOne) SetName(s string) *GuildUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// SetPlayerChannelID sets the "player_channel_id" field.
func (guo *GuildUpdateOne) SetPlayerChannelID(s snowflake.ID) *GuildUpdateOne {
	guo.mutation.ResetPlayerChannelID()
	guo.mutation.SetPlayerChannelID(s)
	return guo
}

// SetNillablePlayerChannelID sets the "player_channel_id" field if the given value is not nil.
func (guo *GuildUpdateOne) SetNillablePlayerChannelID(s *snowflake.ID) *GuildUpdateOne {
	if s != nil {
		guo.SetPlayerChannelID(*s)
	}
	return guo
}

// AddPlayerChannelID adds s to the "player_channel_id" field.
func (guo *GuildUpdateOne) AddPlayerChannelID(s snowflake.ID) *GuildUpdateOne {
	guo.mutation.AddPlayerChannelID(s)
	return guo
}

// ClearPlayerChannelID clears the value of the "player_channel_id" field.
func (guo *GuildUpdateOne) ClearPlayerChannelID() *GuildUpdateOne {
	guo.mutation.ClearPlayerChannelID()
	return guo
}

// SetPlayerMessageID sets the "player_message_id" field.
func (guo *GuildUpdateOne) SetPlayerMessageID(s snowflake.ID) *GuildUpdateOne {
	guo.mutation.ResetPlayerMessageID()
	guo.mutation.SetPlayerMessageID(s)
	return guo
}

// SetNillablePlayerMessageID sets the "player_message_id" field if the given value is not nil.
func (guo *GuildUpdateOne) SetNillablePlayerMessageID(s *snowflake.ID) *GuildUpdateOne {
	if s != nil {
		guo.SetPlayerMessageID(*s)
	}
	return guo
}

// AddPlayerMessageID adds s to the "player_message_id" field.
func (guo *GuildUpdateOne) AddPlayerMessageID(s snowflake.ID) *GuildUpdateOne {
	guo.mutation.AddPlayerMessageID(s)
	return guo
}

// ClearPlayerMessageID clears the value of the "player_message_id" field.
func (guo *GuildUpdateOne) ClearPlayerMessageID() *GuildUpdateOne {
	guo.mutation.ClearPlayerMessageID()
	return guo
}

// SetCreatedAt sets the "created_at" field.
func (guo *GuildUpdateOne) SetCreatedAt(t time.Time) *GuildUpdateOne {
	guo.mutation.SetCreatedAt(t)
	return guo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (guo *GuildUpdateOne) SetNillableCreatedAt(t *time.Time) *GuildUpdateOne {
	if t != nil {
		guo.SetCreatedAt(*t)
	}
	return guo
}

// SetUpdatedAt sets the "updated_at" field.
func (guo *GuildUpdateOne) SetUpdatedAt(t time.Time) *GuildUpdateOne {
	guo.mutation.SetUpdatedAt(t)
	return guo
}

// Mutation returns the GuildMutation object of the builder.
func (guo *GuildUpdateOne) Mutation() *GuildMutation {
	return guo.mutation
}

// Where appends a list predicates to the GuildUpdate builder.
func (guo *GuildUpdateOne) Where(ps ...predicate.Guild) *GuildUpdateOne {
	guo.mutation.Where(ps...)
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GuildUpdateOne) Select(field string, fields ...string) *GuildUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Guild entity.
func (guo *GuildUpdateOne) Save(ctx context.Context) (*Guild, error) {
	guo.defaults()
	return withHooks[*Guild, GuildMutation](ctx, guo.sqlSave, guo.mutation, guo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GuildUpdateOne) SaveX(ctx context.Context) *Guild {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GuildUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GuildUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (guo *GuildUpdateOne) defaults() {
	if _, ok := guo.mutation.UpdatedAt(); !ok {
		v := guild.UpdateDefaultUpdatedAt()
		guo.mutation.SetUpdatedAt(v)
	}
}

func (guo *GuildUpdateOne) sqlSave(ctx context.Context) (_node *Guild, err error) {
	_spec := sqlgraph.NewUpdateSpec(guild.Table, guild.Columns, sqlgraph.NewFieldSpec(guild.FieldID, field.TypeUint64))
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Guild.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, guild.FieldID)
		for _, f := range fields {
			if !guild.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != guild.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.Name(); ok {
		_spec.SetField(guild.FieldName, field.TypeString, value)
	}
	if value, ok := guo.mutation.PlayerChannelID(); ok {
		_spec.SetField(guild.FieldPlayerChannelID, field.TypeUint64, value)
	}
	if value, ok := guo.mutation.AddedPlayerChannelID(); ok {
		_spec.AddField(guild.FieldPlayerChannelID, field.TypeUint64, value)
	}
	if guo.mutation.PlayerChannelIDCleared() {
		_spec.ClearField(guild.FieldPlayerChannelID, field.TypeUint64)
	}
	if value, ok := guo.mutation.PlayerMessageID(); ok {
		_spec.SetField(guild.FieldPlayerMessageID, field.TypeUint64, value)
	}
	if value, ok := guo.mutation.AddedPlayerMessageID(); ok {
		_spec.AddField(guild.FieldPlayerMessageID, field.TypeUint64, value)
	}
	if guo.mutation.PlayerMessageIDCleared() {
		_spec.ClearField(guild.FieldPlayerMessageID, field.TypeUint64)
	}
	if value, ok := guo.mutation.CreatedAt(); ok {
		_spec.SetField(guild.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := guo.mutation.UpdatedAt(); ok {
		_spec.SetField(guild.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Guild{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{guild.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	guo.mutation.done = true
	return _node, nil
}
