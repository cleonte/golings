package cmd_test

import (
	"bytes"
	"encoding/json"

	"github.com/cleonte/golings/golings/cmd"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
)

// runJSONCommand executes the command with args and returns what it wrote to
// its output, which must be one JSON document, and the error it ended with.
func runJSONCommand(c *cobra.Command, args ...string) (map[string]any, error) {
	out := &bytes.Buffer{}
	c.SetOut(out)
	c.SetArgs(args)
	err := c.Execute()

	var doc map[string]any
	ExpectWithOffset(1, json.Unmarshal(out.Bytes(), &doc)).To(Succeed(), "the output is %q", out.String())
	return doc, err
}

var _ = Describe("--json", func() {
	BeforeEach(func() {
		// Nothing of the colored output may reach the JSON: it goes to the
		// output of the command. Captured here so a leak shows up in the test.
		captureColor()
	})

	Describe("list", func() {
		It("describes the exercises, their state and the progress", func() {
			info := writeTopicInfo("a/a1", "!a/a2", "!b/b1")

			doc, err := runJSONCommand(cmd.ListCmd(info), "--json")

			Expect(err).NotTo(HaveOccurred())
			Expect(doc["version"]).To(BeEquivalentTo(cmd.JSONVersion))
			Expect(doc["done"]).To(BeEquivalentTo(1))
			Expect(doc["total"]).To(BeEquivalentTo(3))
			Expect(doc["topics"]).To(Equal([]any{
				map[string]any{"name": "a", "done": float64(1), "total": float64(2)},
				map[string]any{"name": "b", "done": float64(0), "total": float64(1)},
			}))

			exercises := doc["exercises"].([]any)
			Expect(exercises).To(HaveLen(3))
			first := exercises[0].(map[string]any)
			Expect(first).To(HaveKeyWithValue("name", "a1"))
			Expect(first).To(HaveKeyWithValue("topic", "a"))
			Expect(first).To(HaveKeyWithValue("mode", "compile"))
			Expect(first).To(HaveKeyWithValue("state", "done"))
			Expect(first["path"]).To(HaveSuffix("exercises/a/a1/main.go"))
			Expect(first).NotTo(HaveKey("skipped"))
			Expect(exercises[1].(map[string]any)).To(HaveKeyWithValue("state", "pending"))
		})

		It("names the reason an exercise is skipped and leaves it out of the progress", func() {
			doc, err := runJSONCommand(cmd.ListCmd("../fixtures/goversion/info.toml"), "--json")

			Expect(err).NotTo(HaveOccurred())
			Expect(doc["total"]).To(BeEquivalentTo(1))
			var skipped []any
			for _, e := range doc["exercises"].([]any) {
				if reason, ok := e.(map[string]any)["skipped"]; ok {
					skipped = append(skipped, reason)
				}
			}
			Expect(skipped).To(Equal([]any{"needs Go 99.0"}))
		})

		It("prints empty lists, not null, for a manifest without exercises", func() {
			doc, err := runJSONCommand(cmd.ListCmd("../fixtures/empty/info.toml"), "--json")

			Expect(err).NotTo(HaveOccurred())
			Expect(doc["exercises"]).To(Equal([]any{}))
			Expect(doc["topics"]).To(Equal([]any{}))
		})

		It("prints the error as JSON, and fails, when the info file is missing", func() {
			doc, err := runJSONCommand(cmd.ListCmd("../fixtures/info404.toml"), "--json")

			Expect(err).To(HaveOccurred())
			Expect(doc["version"]).To(BeEquivalentTo(cmd.JSONVersion))
			Expect(doc["error"]).To(ContainSubstring("info404.toml"))
		})
	})

	Describe("run", func() {
		It("says passed for a finished exercise, with what it printed", func() {
			doc, err := runJSONCommand(cmd.RunCmd("../fixtures/success1/info.toml"), "--json", "success1")

			Expect(err).NotTo(HaveOccurred())
			Expect(doc["status"]).To(Equal("passed"))
			Expect(doc["exercise"]).To(HaveKeyWithValue("name", "success1"))
			Expect(doc).NotTo(HaveKey("failure"))
		})

		It("says pending for an exercise that passes but is not marked done", func() {
			doc, err := runJSONCommand(cmd.RunCmd("../fixtures/pending1/info.toml"), "--json", "pending1")

			Expect(err).NotTo(HaveOccurred())
			Expect(doc["status"]).To(Equal("pending"))
		})

		It("explains a compile failure with its diagnostics, and exits 0", func() {
			doc, err := runJSONCommand(cmd.RunCmd("../fixtures/error1/info.toml"), "--json", "error1")

			Expect(err).NotTo(HaveOccurred())
			Expect(doc["status"]).To(Equal("failed"))
			failure := doc["failure"].(map[string]any)
			Expect(failure).To(HaveKeyWithValue("kind", "compile"))
			diagnostic := failure["diagnostics"].([]any)[0].(map[string]any)
			Expect(diagnostic["file"]).To(HaveSuffix("fixtures/error1/main.go"))
			Expect(diagnostic["line"]).To(BeEquivalentTo(7))
			Expect(diagnostic["message"]).To(ContainSubstring("declared and not used"))
		})

		It("explains a failing test", func() {
			doc, err := runJSONCommand(cmd.RunCmd(writeFailingTestInfo()), "--json", "broken-test")

			Expect(err).NotTo(HaveOccurred())
			Expect(doc["status"]).To(Equal("failed"))
			failure := doc["failure"].(map[string]any)
			Expect(failure).To(HaveKeyWithValue("kind", "test"))
			Expect(failure["output"]).To(ContainSubstring("not yet"))
		})

		It("runs the next pending exercise", func() {
			doc, err := runJSONCommand(cmd.RunCmd(writeTopicInfo("a/a1", "!a/a2")), "--json", "next")

			Expect(err).NotTo(HaveOccurred())
			Expect(doc["exercise"]).To(HaveKeyWithValue("name", "a2"))
		})

		It("says skipped for an exercise this machine can't run", func() {
			doc, err := runJSONCommand(cmd.RunCmd("../fixtures/goversion/info.toml"), "--json", "future1")

			Expect(err).NotTo(HaveOccurred())
			Expect(doc["status"]).To(Equal("skipped"))
			Expect(doc["message"]).To(ContainSubstring("future1 needs Go 99.0"))
		})

		It("prints the error as JSON, and fails, for an exercise that does not exist", func() {
			doc, err := runJSONCommand(cmd.RunCmd("../fixtures/success1/info.toml"), "--json", "nope")

			Expect(err).To(HaveOccurred())
			Expect(doc["error"]).To(ContainSubstring("nope"))
			Expect(doc).NotTo(HaveKey("status"))
		})

		It("prints the error as JSON, and fails, when nothing is pending for next", func() {
			doc, err := runJSONCommand(cmd.RunCmd(writeTopicInfo("a/a1")), "--json", "next")

			Expect(err).To(HaveOccurred())
			Expect(doc["error"]).To(ContainSubstring("no pending exercises"))
		})
	})
})
