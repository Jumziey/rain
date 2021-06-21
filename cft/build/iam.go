package build

import (
	"github.com/jumziey/rain/cft"
	"github.com/jumziey/rain/cft/spec"
)

// iamBuilder contains specific code for building IAM policies
type iamBuilder struct {
	builder
}

// newIamBuilder creates a new iamBuilder
func newIamBuilder() iamBuilder {
	var b iamBuilder
	b.Spec = spec.Iam

	return b
}

// Policy generates a an IAM policy body
func (b iamBuilder) Policy() (interface{}, []*cft.Comment) {
	return b.newPropertyType("", "Policy")
}
