// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/license"
)

// License is the model entity for the License schema.
type License struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Inline holds the value of the "inline" field.
	Inline string `json:"inline,omitempty"`
	// ListVersion holds the value of the "list_version" field.
	ListVersion string `json:"list_version,omitempty"`
	// An opaque hash on the linline text
	InlineHash string `json:"inline_hash,omitempty"`
	// An opaque hash on the list_version text
	ListVersionHash string `json:"list_version_hash,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LicenseQuery when eager-loading is set.
	Edges        LicenseEdges `json:"edges"`
	selectValues sql.SelectValues
}

// LicenseEdges holds the relations/edges for other nodes in the graph.
type LicenseEdges struct {
	// DeclaredInCertifyLegals holds the value of the declared_in_certify_legals edge.
	DeclaredInCertifyLegals []*CertifyLegal `json:"declared_in_certify_legals,omitempty"`
	// DiscoveredInCertifyLegals holds the value of the discovered_in_certify_legals edge.
	DiscoveredInCertifyLegals []*CertifyLegal `json:"discovered_in_certify_legals,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedDeclaredInCertifyLegals   map[string][]*CertifyLegal
	namedDiscoveredInCertifyLegals map[string][]*CertifyLegal
}

// DeclaredInCertifyLegalsOrErr returns the DeclaredInCertifyLegals value or an error if the edge
// was not loaded in eager-loading.
func (e LicenseEdges) DeclaredInCertifyLegalsOrErr() ([]*CertifyLegal, error) {
	if e.loadedTypes[0] {
		return e.DeclaredInCertifyLegals, nil
	}
	return nil, &NotLoadedError{edge: "declared_in_certify_legals"}
}

// DiscoveredInCertifyLegalsOrErr returns the DiscoveredInCertifyLegals value or an error if the edge
// was not loaded in eager-loading.
func (e LicenseEdges) DiscoveredInCertifyLegalsOrErr() ([]*CertifyLegal, error) {
	if e.loadedTypes[1] {
		return e.DiscoveredInCertifyLegals, nil
	}
	return nil, &NotLoadedError{edge: "discovered_in_certify_legals"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*License) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case license.FieldName, license.FieldInline, license.FieldListVersion, license.FieldInlineHash, license.FieldListVersionHash:
			values[i] = new(sql.NullString)
		case license.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the License fields.
func (l *License) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case license.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				l.ID = *value
			}
		case license.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				l.Name = value.String
			}
		case license.FieldInline:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field inline", values[i])
			} else if value.Valid {
				l.Inline = value.String
			}
		case license.FieldListVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field list_version", values[i])
			} else if value.Valid {
				l.ListVersion = value.String
			}
		case license.FieldInlineHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field inline_hash", values[i])
			} else if value.Valid {
				l.InlineHash = value.String
			}
		case license.FieldListVersionHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field list_version_hash", values[i])
			} else if value.Valid {
				l.ListVersionHash = value.String
			}
		default:
			l.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the License.
// This includes values selected through modifiers, order, etc.
func (l *License) Value(name string) (ent.Value, error) {
	return l.selectValues.Get(name)
}

// QueryDeclaredInCertifyLegals queries the "declared_in_certify_legals" edge of the License entity.
func (l *License) QueryDeclaredInCertifyLegals() *CertifyLegalQuery {
	return NewLicenseClient(l.config).QueryDeclaredInCertifyLegals(l)
}

// QueryDiscoveredInCertifyLegals queries the "discovered_in_certify_legals" edge of the License entity.
func (l *License) QueryDiscoveredInCertifyLegals() *CertifyLegalQuery {
	return NewLicenseClient(l.config).QueryDiscoveredInCertifyLegals(l)
}

// Update returns a builder for updating this License.
// Note that you need to call License.Unwrap() before calling this method if this License
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *License) Update() *LicenseUpdateOne {
	return NewLicenseClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the License entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *License) Unwrap() *License {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: License is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *License) String() string {
	var builder strings.Builder
	builder.WriteString("License(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("name=")
	builder.WriteString(l.Name)
	builder.WriteString(", ")
	builder.WriteString("inline=")
	builder.WriteString(l.Inline)
	builder.WriteString(", ")
	builder.WriteString("list_version=")
	builder.WriteString(l.ListVersion)
	builder.WriteString(", ")
	builder.WriteString("inline_hash=")
	builder.WriteString(l.InlineHash)
	builder.WriteString(", ")
	builder.WriteString("list_version_hash=")
	builder.WriteString(l.ListVersionHash)
	builder.WriteByte(')')
	return builder.String()
}

// NamedDeclaredInCertifyLegals returns the DeclaredInCertifyLegals named value or an error if the edge was not
// loaded in eager-loading with this name.
func (l *License) NamedDeclaredInCertifyLegals(name string) ([]*CertifyLegal, error) {
	if l.Edges.namedDeclaredInCertifyLegals == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := l.Edges.namedDeclaredInCertifyLegals[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (l *License) appendNamedDeclaredInCertifyLegals(name string, edges ...*CertifyLegal) {
	if l.Edges.namedDeclaredInCertifyLegals == nil {
		l.Edges.namedDeclaredInCertifyLegals = make(map[string][]*CertifyLegal)
	}
	if len(edges) == 0 {
		l.Edges.namedDeclaredInCertifyLegals[name] = []*CertifyLegal{}
	} else {
		l.Edges.namedDeclaredInCertifyLegals[name] = append(l.Edges.namedDeclaredInCertifyLegals[name], edges...)
	}
}

// NamedDiscoveredInCertifyLegals returns the DiscoveredInCertifyLegals named value or an error if the edge was not
// loaded in eager-loading with this name.
func (l *License) NamedDiscoveredInCertifyLegals(name string) ([]*CertifyLegal, error) {
	if l.Edges.namedDiscoveredInCertifyLegals == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := l.Edges.namedDiscoveredInCertifyLegals[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (l *License) appendNamedDiscoveredInCertifyLegals(name string, edges ...*CertifyLegal) {
	if l.Edges.namedDiscoveredInCertifyLegals == nil {
		l.Edges.namedDiscoveredInCertifyLegals = make(map[string][]*CertifyLegal)
	}
	if len(edges) == 0 {
		l.Edges.namedDiscoveredInCertifyLegals[name] = []*CertifyLegal{}
	} else {
		l.Edges.namedDiscoveredInCertifyLegals[name] = append(l.Edges.namedDiscoveredInCertifyLegals[name], edges...)
	}
}

// Licenses is a parsable slice of License.
type Licenses []*License
