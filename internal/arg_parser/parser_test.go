package arg_parser

import (
	"reflect"
	"strings"
	"testing"
)

func split(args string) []string {
	return strings.Split(args, " ")
}

func TestEmptyArgs(t *testing.T) {
	parser := New()
	result := parser.Parse([]string{})

	expected := ParseContext{
		Config:  []Config{},
		Command: "",
		Filters: []Filter{},
		Args:    []Arg{},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestIdFilter(t *testing.T) {
	args := split("12")
	parser := New()
	result := parser.Parse(args)

	expected := ParseContext{
		Config:  []Config{},
		Command: "",
		Filters: []Filter{IdFilter{Ids: []int{12}}},
		Args:    []Arg{},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestIdRangeFilter(t *testing.T) {
	args := split("8-16")
	parser := New()
	result := parser.Parse(args)

	expected := ParseContext{
		Config:  []Config{},
		Command: "",
		Filters: []Filter{
			IdFilter{Ids: []int{8, 9, 10, 11, 12, 13, 14, 15, 16}},
		},
		Args: []Arg{},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMultipleIdFilters(t *testing.T) {
	args := split("3 hello 18-23 +work 1-6")
	parser := New()
	result := parser.Parse(args)

	expected := ParseContext{
		Config:  []Config{},
		Command: "",
		Filters: []Filter{
			TagFilter{Operator: Include, Tag: "work"},
			IdFilter{Ids: []int{3, 18, 19, 20, 21, 22, 23, 1, 2, 3, 4, 5, 6}},
			TextFilter{Text: "hello"},
		},
		Args: []Arg{},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSupportsCommaSeparatedIds(t *testing.T) {
	args := split("3,18-23,19,30 hello")
	parser := New()
	result := parser.Parse(args)

	expected := ParseContext{
		Config:  []Config{},
		Command: "",
		Filters: []Filter{
			IdFilter{Ids: []int{3, 18, 19, 20, 21, 22, 23, 19, 30}},
			TextFilter{Text: "hello"},
		},
		Args: []Arg{},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestIdFilterCommand(t *testing.T) {
	args := split("932 done")
	parser := New()
	result := parser.Parse(args)

	expected := ParseContext{
		Config:  []Config{},
		Command: Done,
		Filters: []Filter{IdFilter{Ids: []int{932}}},
		Args:    []Arg{},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMultipleFiltersCommand(t *testing.T) {
	args := split("+work priority:foo delete")
	parser := New()
	result := parser.Parse(args)

	expected := ParseContext{
		Config:  []Config{},
		Command: Delete,
		Filters: []Filter{
			TagFilter{Operator: Include, Tag: "work"},
			ScopedFilter{Scope: Priority, Value: "foo"},
		},
		Args: []Arg{},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestScopeArgs(t *testing.T) {
	args := split("11 edit priority: priority:H")
	parser := New()
	result := parser.Parse(args)

	expected := ParseContext{
		Config:  []Config{},
		Command: Edit,
		Filters: []Filter{
			IdFilter{Ids: []int{11}},
		},
		Args: []Arg{
			ScopedArg{Scope: Priority, Value: ""},
			ScopedArg{Scope: Priority, Value: "H"},
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestIgnoresInvalidScopeArgs(t *testing.T) {
	args := split("12 edit foo:bar priority:H")
	parser := New()
	result := parser.Parse(args)

	expected := ParseContext{
		Config:  []Config{},
		Command: Edit,
		Filters: []Filter{
			IdFilter{Ids: []int{12}},
		},
		Args: []Arg{
			ScopedArg{Scope: Priority, Value: "H"},
			TextArg{Text: "foo:bar"},
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestTagArgOperators(t *testing.T) {
	args := split("3 edit hello -work world +home")
	parser := New()
	result := parser.Parse(args)

	expected := ParseContext{
		Config:  []Config{},
		Command: Edit,
		Filters: []Filter{
			IdFilter{Ids: []int{3}},
		},
		Args: []Arg{
			TagArg{Operator: Exclude, Tag: "work"},
			TagArg{Operator: Include, Tag: "home"},
			TextArg{Text: "hello world"},
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestConfigOverrides(t *testing.T) {
	args := split("bulk=3 1-10 done")
	parser := New()
	result := parser.Parse(args)

	expected := ParseContext{
		Config: []Config{
			BulkConfig{Size: 3},
		},
		Command: Done,
		Filters: []Filter{
			IdFilter{Ids: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		},
		Args: []Arg{},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestComplexCommand(t *testing.T) {
	args := split("bulk=8 +work hi priority:foo edit hello -work world priority:L")
	parser := New()
	result := parser.Parse(args)

	expected := ParseContext{
		Config: []Config{
			BulkConfig{Size: 8},
		},
		Command: Edit,
		Filters: []Filter{
			TagFilter{Operator: Include, Tag: "work"},
			ScopedFilter{Scope: Priority, Value: "foo"},
			TextFilter{Text: "hi"},
		},
		Args: []Arg{
			TagArg{Operator: Exclude, Tag: "work"},
			ScopedArg{Scope: Priority, Value: "L"},
			TextArg{Text: "hello world"},
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestOnlyParsesTheCommandOnce(t *testing.T) {
	args := split("12 list edit")
	parser := New()
	result := parser.Parse(args)

	expected := ParseContext{
		Config:  []Config{},
		Command: List,
		Filters: []Filter{
			IdFilter{Ids: []int{12}},
			TextFilter{Text: "edit"},
		},
		Args: []Arg{},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
