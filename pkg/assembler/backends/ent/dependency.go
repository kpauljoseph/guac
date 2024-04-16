// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/dependency"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packagename"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packageversion"
)

// Dependency is the model entity for the Dependency schema.
type Dependency struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// PackageID holds the value of the "package_id" field.
	PackageID uuid.UUID `json:"package_id,omitempty"`
	// DependentPackageNameID holds the value of the "dependent_package_name_id" field.
	DependentPackageNameID uuid.UUID `json:"dependent_package_name_id,omitempty"`
	// DependentPackageVersionID holds the value of the "dependent_package_version_id" field.
	DependentPackageVersionID uuid.UUID `json:"dependent_package_version_id,omitempty"`
	// VersionRange holds the value of the "version_range" field.
	VersionRange string `json:"version_range,omitempty"`
	// DependencyType holds the value of the "dependency_type" field.
	DependencyType dependency.DependencyType `json:"dependency_type,omitempty"`
	// Justification holds the value of the "justification" field.
	Justification string `json:"justification,omitempty"`
	// Origin holds the value of the "origin" field.
	Origin string `json:"origin,omitempty"`
	// Collector holds the value of the "collector" field.
	Collector string `json:"collector,omitempty"`
	// DocumentRef holds the value of the "document_ref" field.
	DocumentRef string `json:"document_ref,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DependencyQuery when eager-loading is set.
	Edges        DependencyEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DependencyEdges holds the relations/edges for other nodes in the graph.
type DependencyEdges struct {
	// Package holds the value of the package edge.
	Package *PackageVersion `json:"package,omitempty"`
	// DependentPackageName holds the value of the dependent_package_name edge.
	DependentPackageName *PackageName `json:"dependent_package_name,omitempty"`
	// DependentPackageVersion holds the value of the dependent_package_version edge.
	DependentPackageVersion *PackageVersion `json:"dependent_package_version,omitempty"`
	// IncludedInSboms holds the value of the included_in_sboms edge.
	IncludedInSboms []*BillOfMaterials `json:"included_in_sboms,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
	// totalCount holds the count of the edges above.
	totalCount [4]map[string]int

	namedIncludedInSboms map[string][]*BillOfMaterials
}

// PackageOrErr returns the Package value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DependencyEdges) PackageOrErr() (*PackageVersion, error) {
	if e.loadedTypes[0] {
		if e.Package == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: packageversion.Label}
		}
		return e.Package, nil
	}
	return nil, &NotLoadedError{edge: "package"}
}

// DependentPackageNameOrErr returns the DependentPackageName value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DependencyEdges) DependentPackageNameOrErr() (*PackageName, error) {
	if e.loadedTypes[1] {
		if e.DependentPackageName == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: packagename.Label}
		}
		return e.DependentPackageName, nil
	}
	return nil, &NotLoadedError{edge: "dependent_package_name"}
}

// DependentPackageVersionOrErr returns the DependentPackageVersion value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DependencyEdges) DependentPackageVersionOrErr() (*PackageVersion, error) {
	if e.loadedTypes[2] {
		if e.DependentPackageVersion == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: packageversion.Label}
		}
		return e.DependentPackageVersion, nil
	}
	return nil, &NotLoadedError{edge: "dependent_package_version"}
}

// IncludedInSbomsOrErr returns the IncludedInSboms value or an error if the edge
// was not loaded in eager-loading.
func (e DependencyEdges) IncludedInSbomsOrErr() ([]*BillOfMaterials, error) {
	if e.loadedTypes[3] {
		return e.IncludedInSboms, nil
	}
	return nil, &NotLoadedError{edge: "included_in_sboms"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Dependency) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case dependency.FieldVersionRange, dependency.FieldDependencyType, dependency.FieldJustification, dependency.FieldOrigin, dependency.FieldCollector, dependency.FieldDocumentRef:
			values[i] = new(sql.NullString)
		case dependency.FieldID, dependency.FieldPackageID, dependency.FieldDependentPackageNameID, dependency.FieldDependentPackageVersionID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Dependency fields.
func (d *Dependency) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dependency.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				d.ID = *value
			}
		case dependency.FieldPackageID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field package_id", values[i])
			} else if value != nil {
				d.PackageID = *value
			}
		case dependency.FieldDependentPackageNameID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field dependent_package_name_id", values[i])
			} else if value != nil {
				d.DependentPackageNameID = *value
			}
		case dependency.FieldDependentPackageVersionID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field dependent_package_version_id", values[i])
			} else if value != nil {
				d.DependentPackageVersionID = *value
			}
		case dependency.FieldVersionRange:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version_range", values[i])
			} else if value.Valid {
				d.VersionRange = value.String
			}
		case dependency.FieldDependencyType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dependency_type", values[i])
			} else if value.Valid {
				d.DependencyType = dependency.DependencyType(value.String)
			}
		case dependency.FieldJustification:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field justification", values[i])
			} else if value.Valid {
				d.Justification = value.String
			}
		case dependency.FieldOrigin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field origin", values[i])
			} else if value.Valid {
				d.Origin = value.String
			}
		case dependency.FieldCollector:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field collector", values[i])
			} else if value.Valid {
				d.Collector = value.String
			}
		case dependency.FieldDocumentRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field document_ref", values[i])
			} else if value.Valid {
				d.DocumentRef = value.String
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Dependency.
// This includes values selected through modifiers, order, etc.
func (d *Dependency) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryPackage queries the "package" edge of the Dependency entity.
func (d *Dependency) QueryPackage() *PackageVersionQuery {
	return NewDependencyClient(d.config).QueryPackage(d)
}

// QueryDependentPackageName queries the "dependent_package_name" edge of the Dependency entity.
func (d *Dependency) QueryDependentPackageName() *PackageNameQuery {
	return NewDependencyClient(d.config).QueryDependentPackageName(d)
}

// QueryDependentPackageVersion queries the "dependent_package_version" edge of the Dependency entity.
func (d *Dependency) QueryDependentPackageVersion() *PackageVersionQuery {
	return NewDependencyClient(d.config).QueryDependentPackageVersion(d)
}

// QueryIncludedInSboms queries the "included_in_sboms" edge of the Dependency entity.
func (d *Dependency) QueryIncludedInSboms() *BillOfMaterialsQuery {
	return NewDependencyClient(d.config).QueryIncludedInSboms(d)
}

// Update returns a builder for updating this Dependency.
// Note that you need to call Dependency.Unwrap() before calling this method if this Dependency
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Dependency) Update() *DependencyUpdateOne {
	return NewDependencyClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Dependency entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Dependency) Unwrap() *Dependency {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Dependency is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Dependency) String() string {
	var builder strings.Builder
	builder.WriteString("Dependency(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("package_id=")
	builder.WriteString(fmt.Sprintf("%v", d.PackageID))
	builder.WriteString(", ")
	builder.WriteString("dependent_package_name_id=")
	builder.WriteString(fmt.Sprintf("%v", d.DependentPackageNameID))
	builder.WriteString(", ")
	builder.WriteString("dependent_package_version_id=")
	builder.WriteString(fmt.Sprintf("%v", d.DependentPackageVersionID))
	builder.WriteString(", ")
	builder.WriteString("version_range=")
	builder.WriteString(d.VersionRange)
	builder.WriteString(", ")
	builder.WriteString("dependency_type=")
	builder.WriteString(fmt.Sprintf("%v", d.DependencyType))
	builder.WriteString(", ")
	builder.WriteString("justification=")
	builder.WriteString(d.Justification)
	builder.WriteString(", ")
	builder.WriteString("origin=")
	builder.WriteString(d.Origin)
	builder.WriteString(", ")
	builder.WriteString("collector=")
	builder.WriteString(d.Collector)
	builder.WriteString(", ")
	builder.WriteString("document_ref=")
	builder.WriteString(d.DocumentRef)
	builder.WriteByte(')')
	return builder.String()
}

// NamedIncludedInSboms returns the IncludedInSboms named value or an error if the edge was not
// loaded in eager-loading with this name.
func (d *Dependency) NamedIncludedInSboms(name string) ([]*BillOfMaterials, error) {
	if d.Edges.namedIncludedInSboms == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := d.Edges.namedIncludedInSboms[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (d *Dependency) appendNamedIncludedInSboms(name string, edges ...*BillOfMaterials) {
	if d.Edges.namedIncludedInSboms == nil {
		d.Edges.namedIncludedInSboms = make(map[string][]*BillOfMaterials)
	}
	if len(edges) == 0 {
		d.Edges.namedIncludedInSboms[name] = []*BillOfMaterials{}
	} else {
		d.Edges.namedIncludedInSboms[name] = append(d.Edges.namedIncludedInSboms[name], edges...)
	}
}

// Dependencies is a parsable slice of Dependency.
type Dependencies []*Dependency
