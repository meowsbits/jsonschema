package jsonschema

import (
	"net"
	"net/url"
	"testing"
	"time"
)

type GrandfatherType struct {
	FamilyName string `json:"family_name"`
}

type SomeBaseType struct {
	SomeBaseProperty        int             `json:"some_base_property"`
	somePrivateBaseProperty string          `json:"i_am_private"`
	SomeIgnoredBaseProperty string          `json:"-"`
	Grandfather             GrandfatherType `json:"grand"`

	SomeUntaggedBaseProperty           bool
	someUnexportedUntaggedBaseProperty bool
}

type nonExported struct {
	PublicNonExported  int
	privateNonExported int
}

type TestUser struct {
	SomeBaseType
	nonExported

	ID      int                    `json:"id"`
	Name    string                 `json:"name"`
	Friends []int                  `json:"friends,omitempty"`
	Tags    map[string]interface{} `json:"tags,omitempty"`

	TestFlag       bool
	IgnoredCounter int `json:"-"`

	// Tests for RFC draft-wright-json-schema-validation-00, section 7.3
	BirthDate time.Time `json:"birth_date,omitempty"`
	Website   url.URL   `json:"website,omitempty"`
	IPAddress net.IP    `json:"network_address,omitempty"`
}

// TestSchemaGeneration checks if schema generated correctly:
// - fields marked with "-" are ignored
// - non-exported fields are ignored
// - anonymous fields are expanded
func TestSchemaGeneration(t *testing.T) {
	s := Reflect(&TestUser{})

	expectedProperties := map[string]string{
		"id":                       "integer",
		"name":                     "string",
		"friends":                  "array",
		"tags":                     "object",
		"birth_date":               "string",
		"TestFlag":                 "boolean",
		"some_base_property":       "integer",
		"grand":                    "#/definitions/GrandfatherType",
		"SomeUntaggedBaseProperty": "boolean",
		"website":                  "string",
		"network_address":          "string",
	}

	props := s.Definitions["TestUser"].Properties
	for defKey, prop := range props {
		typeOrRef, ok := expectedProperties[defKey]
		if !ok {
			t.Fatalf("unexpected property '%s'", defKey)
		}
		if prop.Type != "" && prop.Type != typeOrRef {
			t.Fatalf("expected property type '%s', got '%s' for property '%s'", typeOrRef, prop.Type, defKey)
		} else if prop.Ref != "" && prop.Ref != typeOrRef {
			t.Fatalf("expected reference to '%s', got '%s' for property '%s'", typeOrRef, prop.Ref, defKey)
		}
	}

	for defKey := range expectedProperties {
		if _, ok := props[defKey]; !ok {
			t.Fatalf("expected property missing '%s'", defKey)
		}
	}
}
