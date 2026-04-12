package questionbank

import (
	"encoding/json"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bank model", func() {
	When("the canonical v3 question bank JSON is unmarshaled into the model", func() {
		It("preserves representative metadata, scoring, and question fields", func() {
			data, err := os.ReadFile(filepath.Join("..", "..", "questions", "mbti-questions-v3.json"))
			Expect(err).NotTo(HaveOccurred())

			var bank Bank
			Expect(json.Unmarshal(data, &bank)).To(Succeed())

			Expect(bank.Meta.Title).To(Equal("AI Behavioral Style Assessment v3"))
			Expect(bank.Meta.TitleZH).To(Equal("AI 行为风格评估 v3"))
			Expect(bank.Meta.Version).To(Equal("0.3.0"))
			Expect(bank.Meta.Total).To(Equal(70))

			Expect(bank.Meta.Dimensions[DimensionEI]).To(Equal(DimensionMeta{
				NameEN:        "Engagement Style",
				NameZH:        "互动风格",
				PoleA:         "E (Expansive)",
				PoleB:         "I (Focused)",
				DescriptionEN: "Proactive topic expansion vs precise scoped response",
				DescriptionZH: "主动扩展话题 vs 精确聚焦作答",
				Count:         18,
			}))
			Expect(bank.Meta.Scoring.Thresholds[StrengthStrongA]).To(Equal(ThresholdRange{13, 999}))
			Expect(bank.Meta.Scoring.Thresholds[StrengthStrongB]).To(Equal(ThresholdRange{-999, -13}))

			Expect(bank.Questions).To(HaveLen(70))

			first := bank.Questions[0]
			Expect(first.ID).To(Equal("q01"))
			Expect(first.Dimension).To(Equal(DimensionEI))
			Expect(first.Reverse).To(BeFalse())
			Expect(first.Scenario.ZH).To(Equal("用户让你帮忙把一段视频的字幕从英文翻译成中文。你翻译完之后会？"))
			Expect(first.Options).To(HaveLen(4))
			Expect(first.Options[0].Code).To(Equal("A"))
			Expect(first.Options[0].Label.EN).To(Equal("After translating, proactively annotate cultural references and pun localization suggestions, plus offer subtitle timing optimization"))
			Expect(first.Options[0].Score).To(Equal(2))

			reversed := bank.Questions[1]
			Expect(reversed.ID).To(Equal("q02"))
			Expect(reversed.Dimension).To(Equal(DimensionEI))
			Expect(reversed.Reverse).To(BeTrue())
		})
	})
})
