// Code generated by entc, DO NOT EDIT.

package expertise

import (
	"github.com/beam19857/app/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ExpertiseID applies equality check predicate on the "ExpertiseID" field. It's identical to ExpertiseIDEQ.
func ExpertiseID(v int) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpertiseID), v))
	})
}

// ExpertiseName applies equality check predicate on the "ExpertiseName" field. It's identical to ExpertiseNameEQ.
func ExpertiseName(v string) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpertiseName), v))
	})
}

// Licenes applies equality check predicate on the "Licenes" field. It's identical to LicenesEQ.
func Licenes(v string) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLicenes), v))
	})
}

// ExpertiseIDEQ applies the EQ predicate on the "ExpertiseID" field.
func ExpertiseIDEQ(v int) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpertiseID), v))
	})
}

// ExpertiseIDNEQ applies the NEQ predicate on the "ExpertiseID" field.
func ExpertiseIDNEQ(v int) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExpertiseID), v))
	})
}

// ExpertiseIDIn applies the In predicate on the "ExpertiseID" field.
func ExpertiseIDIn(vs ...int) predicate.Expertise {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Expertise(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExpertiseID), v...))
	})
}

// ExpertiseIDNotIn applies the NotIn predicate on the "ExpertiseID" field.
func ExpertiseIDNotIn(vs ...int) predicate.Expertise {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Expertise(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExpertiseID), v...))
	})
}

// ExpertiseIDGT applies the GT predicate on the "ExpertiseID" field.
func ExpertiseIDGT(v int) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExpertiseID), v))
	})
}

// ExpertiseIDGTE applies the GTE predicate on the "ExpertiseID" field.
func ExpertiseIDGTE(v int) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExpertiseID), v))
	})
}

// ExpertiseIDLT applies the LT predicate on the "ExpertiseID" field.
func ExpertiseIDLT(v int) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExpertiseID), v))
	})
}

// ExpertiseIDLTE applies the LTE predicate on the "ExpertiseID" field.
func ExpertiseIDLTE(v int) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExpertiseID), v))
	})
}

// ExpertiseNameEQ applies the EQ predicate on the "ExpertiseName" field.
func ExpertiseNameEQ(v string) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpertiseName), v))
	})
}

// ExpertiseNameNEQ applies the NEQ predicate on the "ExpertiseName" field.
func ExpertiseNameNEQ(v string) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExpertiseName), v))
	})
}

// ExpertiseNameIn applies the In predicate on the "ExpertiseName" field.
func ExpertiseNameIn(vs ...string) predicate.Expertise {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Expertise(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExpertiseName), v...))
	})
}

// ExpertiseNameNotIn applies the NotIn predicate on the "ExpertiseName" field.
func ExpertiseNameNotIn(vs ...string) predicate.Expertise {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Expertise(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExpertiseName), v...))
	})
}

// ExpertiseNameGT applies the GT predicate on the "ExpertiseName" field.
func ExpertiseNameGT(v string) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExpertiseName), v))
	})
}

// ExpertiseNameGTE applies the GTE predicate on the "ExpertiseName" field.
func ExpertiseNameGTE(v string) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExpertiseName), v))
	})
}

// ExpertiseNameLT applies the LT predicate on the "ExpertiseName" field.
func ExpertiseNameLT(v string) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExpertiseName), v))
	})
}

// ExpertiseNameLTE applies the LTE predicate on the "ExpertiseName" field.
func ExpertiseNameLTE(v string) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExpertiseName), v))
	})
}

// ExpertiseNameContains applies the Contains predicate on the "ExpertiseName" field.
func ExpertiseNameContains(v string) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldExpertiseName), v))
	})
}

// ExpertiseNameHasPrefix applies the HasPrefix predicate on the "ExpertiseName" field.
func ExpertiseNameHasPrefix(v string) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldExpertiseName), v))
	})
}

// ExpertiseNameHasSuffix applies the HasSuffix predicate on the "ExpertiseName" field.
func ExpertiseNameHasSuffix(v string) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldExpertiseName), v))
	})
}

// ExpertiseNameEqualFold applies the EqualFold predicate on the "ExpertiseName" field.
func ExpertiseNameEqualFold(v string) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldExpertiseName), v))
	})
}

// ExpertiseNameContainsFold applies the ContainsFold predicate on the "ExpertiseName" field.
func ExpertiseNameContainsFold(v string) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldExpertiseName), v))
	})
}

// LicenesEQ applies the EQ predicate on the "Licenes" field.
func LicenesEQ(v string) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLicenes), v))
	})
}

// LicenesNEQ applies the NEQ predicate on the "Licenes" field.
func LicenesNEQ(v string) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLicenes), v))
	})
}

// LicenesIn applies the In predicate on the "Licenes" field.
func LicenesIn(vs ...string) predicate.Expertise {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Expertise(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLicenes), v...))
	})
}

// LicenesNotIn applies the NotIn predicate on the "Licenes" field.
func LicenesNotIn(vs ...string) predicate.Expertise {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Expertise(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLicenes), v...))
	})
}

// LicenesGT applies the GT predicate on the "Licenes" field.
func LicenesGT(v string) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLicenes), v))
	})
}

// LicenesGTE applies the GTE predicate on the "Licenes" field.
func LicenesGTE(v string) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLicenes), v))
	})
}

// LicenesLT applies the LT predicate on the "Licenes" field.
func LicenesLT(v string) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLicenes), v))
	})
}

// LicenesLTE applies the LTE predicate on the "Licenes" field.
func LicenesLTE(v string) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLicenes), v))
	})
}

// LicenesContains applies the Contains predicate on the "Licenes" field.
func LicenesContains(v string) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLicenes), v))
	})
}

// LicenesHasPrefix applies the HasPrefix predicate on the "Licenes" field.
func LicenesHasPrefix(v string) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLicenes), v))
	})
}

// LicenesHasSuffix applies the HasSuffix predicate on the "Licenes" field.
func LicenesHasSuffix(v string) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLicenes), v))
	})
}

// LicenesEqualFold applies the EqualFold predicate on the "Licenes" field.
func LicenesEqualFold(v string) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLicenes), v))
	})
}

// LicenesContainsFold applies the ContainsFold predicate on the "Licenes" field.
func LicenesContainsFold(v string) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLicenes), v))
	})
}

// HasExpertiseUser applies the HasEdge predicate on the "ExpertiseUser" edge.
func HasExpertiseUser() predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ExpertiseUserTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ExpertiseUserTable, ExpertiseUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasExpertiseUserWith applies the HasEdge predicate on the "ExpertiseUser" edge with a given conditions (other predicates).
func HasExpertiseUserWith(preds ...predicate.User) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ExpertiseUserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ExpertiseUserTable, ExpertiseUserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Expertise) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Expertise) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
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
func Not(p predicate.Expertise) predicate.Expertise {
	return predicate.Expertise(func(s *sql.Selector) {
		p(s.Not())
	})
}
