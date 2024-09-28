// Code generated by ent, DO NOT EDIT.

package enturl

import (
	"wsw/backend/domain/url"
	"wsw/backend/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Url {
	return predicate.Url(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Url {
	return predicate.Url(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Url {
	return predicate.Url(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Url {
	return predicate.Url(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Url {
	return predicate.Url(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Url {
	return predicate.Url(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Url {
	return predicate.Url(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Url {
	return predicate.Url(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Url {
	return predicate.Url(sql.FieldLTE(FieldID, id))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Url {
	return predicate.Url(sql.FieldEQ(FieldURL, v))
}

// ImageURL applies equality check predicate on the "image_url" field. It's identical to ImageURLEQ.
func ImageURL(v string) predicate.Url {
	return predicate.Url(sql.FieldEQ(FieldImageURL, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Url {
	return predicate.Url(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Url {
	return predicate.Url(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Url {
	return predicate.Url(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Url {
	return predicate.Url(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Url {
	return predicate.Url(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Url {
	return predicate.Url(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Url {
	return predicate.Url(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Url {
	return predicate.Url(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Url {
	return predicate.Url(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Url {
	return predicate.Url(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Url {
	return predicate.Url(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Url {
	return predicate.Url(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Url {
	return predicate.Url(sql.FieldContainsFold(FieldURL, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v url.Status) predicate.Url {
	vc := v
	return predicate.Url(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v url.Status) predicate.Url {
	vc := v
	return predicate.Url(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...url.Status) predicate.Url {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Url(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...url.Status) predicate.Url {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Url(sql.FieldNotIn(FieldStatus, v...))
}

// ImageURLEQ applies the EQ predicate on the "image_url" field.
func ImageURLEQ(v string) predicate.Url {
	return predicate.Url(sql.FieldEQ(FieldImageURL, v))
}

// ImageURLNEQ applies the NEQ predicate on the "image_url" field.
func ImageURLNEQ(v string) predicate.Url {
	return predicate.Url(sql.FieldNEQ(FieldImageURL, v))
}

// ImageURLIn applies the In predicate on the "image_url" field.
func ImageURLIn(vs ...string) predicate.Url {
	return predicate.Url(sql.FieldIn(FieldImageURL, vs...))
}

// ImageURLNotIn applies the NotIn predicate on the "image_url" field.
func ImageURLNotIn(vs ...string) predicate.Url {
	return predicate.Url(sql.FieldNotIn(FieldImageURL, vs...))
}

// ImageURLGT applies the GT predicate on the "image_url" field.
func ImageURLGT(v string) predicate.Url {
	return predicate.Url(sql.FieldGT(FieldImageURL, v))
}

// ImageURLGTE applies the GTE predicate on the "image_url" field.
func ImageURLGTE(v string) predicate.Url {
	return predicate.Url(sql.FieldGTE(FieldImageURL, v))
}

// ImageURLLT applies the LT predicate on the "image_url" field.
func ImageURLLT(v string) predicate.Url {
	return predicate.Url(sql.FieldLT(FieldImageURL, v))
}

// ImageURLLTE applies the LTE predicate on the "image_url" field.
func ImageURLLTE(v string) predicate.Url {
	return predicate.Url(sql.FieldLTE(FieldImageURL, v))
}

// ImageURLContains applies the Contains predicate on the "image_url" field.
func ImageURLContains(v string) predicate.Url {
	return predicate.Url(sql.FieldContains(FieldImageURL, v))
}

// ImageURLHasPrefix applies the HasPrefix predicate on the "image_url" field.
func ImageURLHasPrefix(v string) predicate.Url {
	return predicate.Url(sql.FieldHasPrefix(FieldImageURL, v))
}

// ImageURLHasSuffix applies the HasSuffix predicate on the "image_url" field.
func ImageURLHasSuffix(v string) predicate.Url {
	return predicate.Url(sql.FieldHasSuffix(FieldImageURL, v))
}

// ImageURLEqualFold applies the EqualFold predicate on the "image_url" field.
func ImageURLEqualFold(v string) predicate.Url {
	return predicate.Url(sql.FieldEqualFold(FieldImageURL, v))
}

// ImageURLContainsFold applies the ContainsFold predicate on the "image_url" field.
func ImageURLContainsFold(v string) predicate.Url {
	return predicate.Url(sql.FieldContainsFold(FieldImageURL, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Url) predicate.Url {
	return predicate.Url(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Url) predicate.Url {
	return predicate.Url(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Url) predicate.Url {
	return predicate.Url(sql.NotPredicates(p))
}
