package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/a-h/idl/gotypes"
	"github.com/invopop/jsonschema"
)

func main() {
	r := new(jsonschema.Reflector)
	err := r.AddGoComments("github.com/a-h/idl/gotypes", "./"); 
	if err != nil {
		log.Fatalf("failed to add Go comments: %v", err)
	}
	schema := r.Reflect(gotypes.Person{})
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(schema)
}
