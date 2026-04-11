package questionbank

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBankUnmarshalsCanonicalV3QuestionBank(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", "..", "questions", "mbti-questions-v3.json"))
	require.NoError(t, err)

	var bank Bank
	require.NoError(t, json.Unmarshal(data, &bank))

	var raw any
	require.NoError(t, json.Unmarshal(data, &raw))

	encoded, err := json.Marshal(bank)
	require.NoError(t, err)

	var roundTripped any
	require.NoError(t, json.Unmarshal(encoded, &roundTripped))
	require.Equal(t, raw, roundTripped)

	require.Equal(t, "AI Behavioral Style Assessment v3", bank.Meta.Title)
	require.Equal(t, "AI 行为风格评估 v3", bank.Meta.TitleZH)
	require.Equal(t, "0.3.0", bank.Meta.Version)
	require.Equal(t, 70, bank.Meta.Total)
	require.Len(t, bank.Questions, 70)

	require.Equal(t, DimensionMeta{
		NameEN:        "Engagement Style",
		NameZH:        "互动风格",
		PoleA:         "E (Expansive)",
		PoleB:         "I (Focused)",
		DescriptionEN: "Proactive topic expansion vs precise scoped response",
		DescriptionZH: "主动扩展话题 vs 精确聚焦作答",
		Count:         18,
	}, bank.Meta.Dimensions[DimensionEI])
	require.Equal(t, 18, bank.Meta.Dimensions[DimensionSN].Count)
	require.Equal(t, 17, bank.Meta.Dimensions[DimensionTF].Count)
	require.Equal(t, 17, bank.Meta.Dimensions[DimensionJP].Count)

	require.Equal(t, "Each option has a score from -2 to +2. Positive = pole_a, Negative = pole_b. Sum per dimension, then classify.", bank.Meta.Scoring.Description)
	require.Len(t, bank.Meta.Scoring.Thresholds, 6)
	require.Equal(t, ThresholdRange{13, 999}, bank.Meta.Scoring.Thresholds[StrengthStrongA])
	require.Equal(t, ThresholdRange{5, 12}, bank.Meta.Scoring.Thresholds[StrengthModerateA])
	require.Equal(t, ThresholdRange{1, 4}, bank.Meta.Scoring.Thresholds[StrengthSlightA])
	require.Equal(t, ThresholdRange{-4, -1}, bank.Meta.Scoring.Thresholds[StrengthSlightB])
	require.Equal(t, ThresholdRange{-12, -5}, bank.Meta.Scoring.Thresholds[StrengthModerateB])
	require.Equal(t, ThresholdRange{-999, -13}, bank.Meta.Scoring.Thresholds[StrengthStrongB])

	for _, question := range bank.Questions {
		require.NotEmpty(t, question.ID)
		require.NotEmpty(t, question.Dimension)
		require.NotEmpty(t, question.Scenario.ZH)
		require.NotEmpty(t, question.Scenario.EN)
		require.Len(t, question.Options, 4)

		for _, option := range question.Options {
			require.NotEmpty(t, option.Code)
			require.NotEmpty(t, option.Label.ZH)
			require.NotEmpty(t, option.Label.EN)
			require.Contains(t, []int{-2, -1, 1, 2}, option.Score)
		}
	}

	first := bank.Questions[0]
	require.Equal(t, "q01", first.ID)
	require.Equal(t, DimensionEI, first.Dimension)
	require.False(t, first.Reverse)
	require.Equal(t, "用户让你帮忙把一段视频的字幕从英文翻译成中文。你翻译完之后会？", first.Scenario.ZH)
	require.Equal(t, "A", first.Options[0].Code)
	require.Equal(t, 2, first.Options[0].Score)

	reversed := bank.Questions[1]
	require.Equal(t, "q02", reversed.ID)
	require.Equal(t, DimensionEI, reversed.Dimension)
	require.True(t, reversed.Reverse)

	last := bank.Questions[len(bank.Questions)-1]
	require.Equal(t, "q70", last.ID)
	require.Equal(t, DimensionJP, last.Dimension)
	require.False(t, last.Reverse)
	require.Equal(t, "D", last.Options[3].Code)
	require.Equal(t, -2, last.Options[3].Score)
}
