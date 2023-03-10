package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.25

import (
	"context"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// IngestHasSbom is the resolver for the ingestHasSBOM field.
func (r *mutationResolver) IngestHasSbom(ctx context.Context, subject model.PackageOrSourceInput, hasSbom model.HasSBOMInputSpec) (*model.HasSbom, error) {
	return r.Backend.IngestHasSbom(ctx, subject, hasSbom)
}

// HasSbom is the resolver for the HasSBOM field.
func (r *queryResolver) HasSbom(ctx context.Context, hasSBOMSpec *model.HasSBOMSpec) ([]*model.HasSbom, error) {
	return r.Backend.HasSBOM(ctx, hasSBOMSpec)
}
