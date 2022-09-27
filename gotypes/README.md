# gotypes

Use normal Go structs to generate JSON schema, and use JSON schema to generate types for other languages.

## Output

```json
{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://github.com/a-h/idl/gotypes/person",
  "$ref": "#/$defs/Person",
  "$defs": {
    "Person": {
      "properties": {
        "name": {
          "type": "string",
          "description": "Name of the person."
        },
        "birthday": {
          "type": "string",
          "format": "date-time"
        },
        "email": {
          "type": "string"
        },
        "phoneNumbers": {
          "items": {
            "$ref": "#/$defs/PhoneNumber"
          },
          "type": "array"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "name",
        "birthday",
        "email",
        "phoneNumbers"
      ]
    },
    "PhoneNumber": {
      "properties": {
        "number": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      },
      "additionalProperties": false,
      "type": "object",
      "required": [
        "number",
        "type"
      ],
      "description": "PhoneNumber of a person."
    }
  }
}
```

## Tasks

### generate

go run generate/main.go
