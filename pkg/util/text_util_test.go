package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReplaceTemplateValues(t *testing.T) {
	// Test case 1: Basic template replacement
	content := []byte("Hello {{name}}, welcome to {{company}}!")
	templateValues := map[string]string{
		"{{name}}":    "John",
		"{{company}}": "Acme Inc",
	}
	expected := []byte("Hello John, welcome to Acme Inc!")
	result := ReplaceTemplateValues(content, templateValues)
	require.Equal(t, expected, result)

	// Test case 2: Empty template values
	content = []byte("Hello {{name}}!")
	templateValues = map[string]string{}
	expected = []byte("Hello {{name}}!")
	result = ReplaceTemplateValues(content, templateValues)
	require.Equal(t, expected, result)

	// Test case 3: Multiple occurrences of same template
	content = []byte("{{greeting}} {{name}}, {{greeting}} again!")
	templateValues = map[string]string{
		"{{greeting}}": "Hello",
		"{{name}}":     "John",
	}
	expected = []byte("Hello John, Hello again!")
	result = ReplaceTemplateValues(content, templateValues)
	require.Equal(t, expected, result)

	// Test case 4: No template values in content
	content = []byte("Hello World!")
	templateValues = map[string]string{
		"{{name}}": "John",
	}
	expected = []byte("Hello World!")
	result = ReplaceTemplateValues(content, templateValues)
	require.Equal(t, expected, result)
}
