package richerror

import (
	"reflect"
	"testing"
)

func buildRichError(args RichError) *RichError {
	re := New(args.operation)
	if args.wrappedError != nil {
		re = re.WithError(args.wrappedError)
	}
	if args.kind != 0 {
		re = re.WithKind(args.kind)
	}
	if args.meta != nil {
		re = re.WithMeta(args.meta)
	}
	if args.message != "" {
		re = re.WithMessage(args.message)
	}
	return re
}

func TestRichError_GetOperation(t *testing.T) {
	for _, tc := range getOperationTestCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := buildRichError(tc.args); tc.expected != actual.GetOperation() {
				t.Errorf("\nExpected: %v\nActual: %v\n", tc.expected, actual)
			}
		})
	}
}

func TestRichError_GetKind(t *testing.T) {
	for _, tc := range getKindTestCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := buildRichError(tc.args); tc.expected != actual.GetKind() {
				t.Errorf("\nExpected: %v\nActual: %v\n", tc.expected, actual)
			}
		})
	}
}

func TestRichError_GetMessage(t *testing.T) {
	for _, tc := range getMessageTestCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := buildRichError(tc.args); tc.expected != actual.GetMessage() {
				t.Errorf("\nExpected: %v\nActual: %v\n", tc.expected, actual)
			}
		})
	}
}

func TestRichError_GetMeta(t *testing.T) {
	for _, tc := range getMetaTestCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := buildRichError(tc.args); !reflect.DeepEqual(tc.expected, actual.GetMeta()) {
				t.Errorf("\nExpected: %v\nActual: %v\n", tc.expected, actual.GetMeta())
			}
		})
	}
}

func TestRichError_GetError(t *testing.T) {
	for _, tc := range getErrorTestCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := buildRichError(tc.args); tc.expected != actual.Error() {
				t.Errorf("\nExpected: %v\nActual: %v\n", tc.expected, actual)
			}
		})
	}
}
