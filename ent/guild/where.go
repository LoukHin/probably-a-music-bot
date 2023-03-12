// Code generated by ent, DO NOT EDIT.

package guild

import (
	"time"

	"entgo.io/ent/dialect/sql"
	snowflake "github.com/disgoorg/snowflake/v2"
	"github.com/loukhin/probably-a-music-bot/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id snowflake.ID) predicate.Guild {
	return predicate.Guild(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id snowflake.ID) predicate.Guild {
	return predicate.Guild(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id snowflake.ID) predicate.Guild {
	return predicate.Guild(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...snowflake.ID) predicate.Guild {
	return predicate.Guild(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...snowflake.ID) predicate.Guild {
	return predicate.Guild(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id snowflake.ID) predicate.Guild {
	return predicate.Guild(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id snowflake.ID) predicate.Guild {
	return predicate.Guild(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id snowflake.ID) predicate.Guild {
	return predicate.Guild(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id snowflake.ID) predicate.Guild {
	return predicate.Guild(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Guild {
	return predicate.Guild(sql.FieldEQ(FieldName, v))
}

// PlayerChannelID applies equality check predicate on the "player_channel_id" field. It's identical to PlayerChannelIDEQ.
func PlayerChannelID(v snowflake.ID) predicate.Guild {
	vc := uint64(v)
	return predicate.Guild(sql.FieldEQ(FieldPlayerChannelID, vc))
}

// PlayerMessageID applies equality check predicate on the "player_message_id" field. It's identical to PlayerMessageIDEQ.
func PlayerMessageID(v snowflake.ID) predicate.Guild {
	vc := uint64(v)
	return predicate.Guild(sql.FieldEQ(FieldPlayerMessageID, vc))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Guild {
	return predicate.Guild(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Guild {
	return predicate.Guild(sql.FieldEQ(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Guild {
	return predicate.Guild(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Guild {
	return predicate.Guild(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Guild {
	return predicate.Guild(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Guild {
	return predicate.Guild(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Guild {
	return predicate.Guild(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Guild {
	return predicate.Guild(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Guild {
	return predicate.Guild(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Guild {
	return predicate.Guild(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Guild {
	return predicate.Guild(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Guild {
	return predicate.Guild(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Guild {
	return predicate.Guild(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Guild {
	return predicate.Guild(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Guild {
	return predicate.Guild(sql.FieldContainsFold(FieldName, v))
}

// PlayerChannelIDEQ applies the EQ predicate on the "player_channel_id" field.
func PlayerChannelIDEQ(v snowflake.ID) predicate.Guild {
	vc := uint64(v)
	return predicate.Guild(sql.FieldEQ(FieldPlayerChannelID, vc))
}

// PlayerChannelIDNEQ applies the NEQ predicate on the "player_channel_id" field.
func PlayerChannelIDNEQ(v snowflake.ID) predicate.Guild {
	vc := uint64(v)
	return predicate.Guild(sql.FieldNEQ(FieldPlayerChannelID, vc))
}

// PlayerChannelIDIn applies the In predicate on the "player_channel_id" field.
func PlayerChannelIDIn(vs ...snowflake.ID) predicate.Guild {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.Guild(sql.FieldIn(FieldPlayerChannelID, v...))
}

// PlayerChannelIDNotIn applies the NotIn predicate on the "player_channel_id" field.
func PlayerChannelIDNotIn(vs ...snowflake.ID) predicate.Guild {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.Guild(sql.FieldNotIn(FieldPlayerChannelID, v...))
}

// PlayerChannelIDGT applies the GT predicate on the "player_channel_id" field.
func PlayerChannelIDGT(v snowflake.ID) predicate.Guild {
	vc := uint64(v)
	return predicate.Guild(sql.FieldGT(FieldPlayerChannelID, vc))
}

// PlayerChannelIDGTE applies the GTE predicate on the "player_channel_id" field.
func PlayerChannelIDGTE(v snowflake.ID) predicate.Guild {
	vc := uint64(v)
	return predicate.Guild(sql.FieldGTE(FieldPlayerChannelID, vc))
}

// PlayerChannelIDLT applies the LT predicate on the "player_channel_id" field.
func PlayerChannelIDLT(v snowflake.ID) predicate.Guild {
	vc := uint64(v)
	return predicate.Guild(sql.FieldLT(FieldPlayerChannelID, vc))
}

// PlayerChannelIDLTE applies the LTE predicate on the "player_channel_id" field.
func PlayerChannelIDLTE(v snowflake.ID) predicate.Guild {
	vc := uint64(v)
	return predicate.Guild(sql.FieldLTE(FieldPlayerChannelID, vc))
}

// PlayerChannelIDIsNil applies the IsNil predicate on the "player_channel_id" field.
func PlayerChannelIDIsNil() predicate.Guild {
	return predicate.Guild(sql.FieldIsNull(FieldPlayerChannelID))
}

// PlayerChannelIDNotNil applies the NotNil predicate on the "player_channel_id" field.
func PlayerChannelIDNotNil() predicate.Guild {
	return predicate.Guild(sql.FieldNotNull(FieldPlayerChannelID))
}

// PlayerMessageIDEQ applies the EQ predicate on the "player_message_id" field.
func PlayerMessageIDEQ(v snowflake.ID) predicate.Guild {
	vc := uint64(v)
	return predicate.Guild(sql.FieldEQ(FieldPlayerMessageID, vc))
}

// PlayerMessageIDNEQ applies the NEQ predicate on the "player_message_id" field.
func PlayerMessageIDNEQ(v snowflake.ID) predicate.Guild {
	vc := uint64(v)
	return predicate.Guild(sql.FieldNEQ(FieldPlayerMessageID, vc))
}

// PlayerMessageIDIn applies the In predicate on the "player_message_id" field.
func PlayerMessageIDIn(vs ...snowflake.ID) predicate.Guild {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.Guild(sql.FieldIn(FieldPlayerMessageID, v...))
}

// PlayerMessageIDNotIn applies the NotIn predicate on the "player_message_id" field.
func PlayerMessageIDNotIn(vs ...snowflake.ID) predicate.Guild {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = uint64(vs[i])
	}
	return predicate.Guild(sql.FieldNotIn(FieldPlayerMessageID, v...))
}

// PlayerMessageIDGT applies the GT predicate on the "player_message_id" field.
func PlayerMessageIDGT(v snowflake.ID) predicate.Guild {
	vc := uint64(v)
	return predicate.Guild(sql.FieldGT(FieldPlayerMessageID, vc))
}

// PlayerMessageIDGTE applies the GTE predicate on the "player_message_id" field.
func PlayerMessageIDGTE(v snowflake.ID) predicate.Guild {
	vc := uint64(v)
	return predicate.Guild(sql.FieldGTE(FieldPlayerMessageID, vc))
}

// PlayerMessageIDLT applies the LT predicate on the "player_message_id" field.
func PlayerMessageIDLT(v snowflake.ID) predicate.Guild {
	vc := uint64(v)
	return predicate.Guild(sql.FieldLT(FieldPlayerMessageID, vc))
}

// PlayerMessageIDLTE applies the LTE predicate on the "player_message_id" field.
func PlayerMessageIDLTE(v snowflake.ID) predicate.Guild {
	vc := uint64(v)
	return predicate.Guild(sql.FieldLTE(FieldPlayerMessageID, vc))
}

// PlayerMessageIDIsNil applies the IsNil predicate on the "player_message_id" field.
func PlayerMessageIDIsNil() predicate.Guild {
	return predicate.Guild(sql.FieldIsNull(FieldPlayerMessageID))
}

// PlayerMessageIDNotNil applies the NotNil predicate on the "player_message_id" field.
func PlayerMessageIDNotNil() predicate.Guild {
	return predicate.Guild(sql.FieldNotNull(FieldPlayerMessageID))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Guild {
	return predicate.Guild(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Guild {
	return predicate.Guild(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Guild {
	return predicate.Guild(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Guild {
	return predicate.Guild(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Guild {
	return predicate.Guild(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Guild {
	return predicate.Guild(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Guild {
	return predicate.Guild(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Guild {
	return predicate.Guild(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Guild {
	return predicate.Guild(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Guild {
	return predicate.Guild(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Guild {
	return predicate.Guild(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Guild {
	return predicate.Guild(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Guild {
	return predicate.Guild(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Guild {
	return predicate.Guild(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Guild {
	return predicate.Guild(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Guild {
	return predicate.Guild(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Guild) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Guild) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Guild) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		p(s.Not())
	})
}
