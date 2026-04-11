package answers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseCanonicalAnswerMapNormalizesOptionCodes(t *testing.T) {
	got, err := Parse([]byte(`{
		"answers": {
			"q01": "a",
			"q02": " C "
		}
	}`))

	require.NoError(t, err)
	require.Equal(t, Set{
		"q01": "A",
		"q02": "C",
	}, got)
}

func TestParseRejectsMissingOrNullAnswersObject(t *testing.T) {
	tests := map[string]string{
		"missing": `{}`,
		"null":    `{"answers": null}`,
	}

	for name, input := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := Parse([]byte(input))

			require.Error(t, err)
			require.ErrorContains(t, err, "parse answer file")
			require.ErrorContains(t, err, "answers object is required")
		})
	}
}

func TestParseRejectsMalformedJSON(t *testing.T) {
	_, err := Parse([]byte(`{"answers":`))

	require.Error(t, err)
	require.ErrorContains(t, err, "parse answer file")
}

func TestParseRejectsStructurallyInvalidAnswerValues(t *testing.T) {
	_, err := Parse([]byte(`{"answers":{"q01":1}}`))

	require.Error(t, err)
	require.ErrorContains(t, err, "parse answer file")
}

func TestParseDefersBankAwareValidation(t *testing.T) {
	got, err := Parse([]byte(`{"answers":{"unknown":"z"}}`))

	require.NoError(t, err)
	require.Equal(t, Set{"unknown": "Z"}, got)
}
