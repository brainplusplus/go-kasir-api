package swagger

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/invopop/jsonschema"
)

type Route struct {
	Method      string
	Path        string
	Description string
	Request     interface{}
	Response    interface{}
}

type Generator struct {
	Title   string
	Version string
	Routes  []Route
}

func NewGenerator(title, version string) *Generator {
	return &Generator{
		Title:   title,
		Version: version,
		Routes:  make([]Route, 0),
	}
}

func (g *Generator) Register(method, path, description string, req interface{}, res interface{}) {
	g.Routes = append(g.Routes, Route{
		Method:      method,
		Path:        path,
		Description: description,
		Request:     req,
		Response:    res,
	})
}

func (g *Generator) Generate() map[string]interface{} {
	reflector := jsonschema.Reflector{}
	// Ensure we capture definitions
	// reflector.ExpandedStruct = false // Default is false, which allows definitions

	paths := make(map[string]map[string]interface{})
	definitions := make(map[string]interface{})

	// Helper to add schema definitions to global definitions and fix refs
	addDefinitions := func(schema *jsonschema.Schema) {
		for name, def := range schema.Definitions {
			// recursively fix refs in the definition
			fixRefs(def)
			definitions[name] = def
		}
	}

	// Compile regex for path parameters
	paramRegex := regexp.MustCompile(`\{([^}]+)\}`)

	for _, route := range g.Routes {
		if paths[route.Path] == nil {
			paths[route.Path] = make(map[string]interface{})
		}

		// Infer Tags from Path
		var tags []string
		if strings.HasPrefix(route.Path, "/api/") {
			parts := strings.Split(route.Path, "/")
			if len(parts) > 2 {
				// capitalize first letter
				tag := parts[2]
				if len(tag) > 0 {
					tag = strings.ToUpper(tag[:1]) + tag[1:]
				}
				tags = []string{tag}
			}
		} else if route.Path == "/health" {
			tags = []string{"Health"}
		}

		operation := map[string]interface{}{
			"summary": route.Description,
			"tags":    tags,
			"responses": map[string]interface{}{
				"200": map[string]interface{}{
					"description": "OK",
				},
			},
		}

		// Handle Path Parameters
		matches := paramRegex.FindAllStringSubmatch(route.Path, -1)
		if len(matches) > 0 {
			parameters := []map[string]interface{}{}
			for _, match := range matches {
				paramName := match[1]
				parameters = append(parameters, map[string]interface{}{
					"name":     paramName,
					"in":       "path",
					"required": true,
					"schema": map[string]string{
						"type": "integer", // Defaulting to integer for this project
					},
					"description": paramName,
				})
			}
			operation["parameters"] = parameters
		}

		// Handle Request Body
		if route.Request != nil {
			schema := reflector.Reflect(route.Request)
			addDefinitions(schema)
			fixRefs(schema)

			// If the reflected schema is a pure ref to a definition, use it.
			// The definitions are now in #/components/schemas so the ref should point there.

			// We need to marshal/unmarshal to interface{} to be safe for final JSON map or just use schema
			operation["requestBody"] = map[string]interface{}{
				"content": map[string]interface{}{
					"application/json": map[string]interface{}{
						"schema": schema,
					},
				},
			}
		}

		// Handle Response
		if route.Response != nil {
			schema := reflector.Reflect(route.Response)
			addDefinitions(schema)
			fixRefs(schema)

			operation["responses"] = map[string]interface{}{
				"200": map[string]interface{}{
					"description": "OK",
					"content": map[string]interface{}{
						"application/json": map[string]interface{}{
							"schema": schema,
						},
					},
				},
			}
		}

		paths[route.Path][strings.ToLower(route.Method)] = operation
	}

	return map[string]interface{}{
		"openapi": "3.0.0",
		"info": map[string]interface{}{
			"title":   g.Title,
			"version": g.Version,
		},
		"paths": paths,
		"components": map[string]interface{}{
			"schemas": definitions,
		},
	}
}

// Recursively replace #/$defs/ with #/components/schemas/
func fixRefs(s *jsonschema.Schema) {
	if s == nil {
		return
	}

	if strings.HasPrefix(s.Ref, "#/$defs/") {
		s.Ref = strings.Replace(s.Ref, "#/$defs/", "#/components/schemas/", 1)
	}

	// Recurse for items, properties, etc.
	if s.Items != nil {
		fixRefs(s.Items)
	}
	if s.Properties != nil {
		for pair := s.Properties.Oldest(); pair != nil; pair = pair.Next() {
			fixRefs(pair.Value)
		}
	}
	if s.AdditionalProperties != nil {
		fixRefs(s.AdditionalProperties)
	}
	// Note: jsonschema might have other fields like OneOf, AnyOf but for this simple use case loop this is likely enough.
	// Ideally we traverse all schema pointers.
}

func (g *Generator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(g.Generate())
}
