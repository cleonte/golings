package exercises

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var update = flag.Bool("update", false, "rewrite the golden files of the race report tests")

// The inputs in testdata/race are the real output of the unsolved exercises
// run with the race detector (concurrent10 and concurrent2), with the path of
// the machine they were captured on replaced by /home/learner/golings-exercises
// and the Go root by /usr/local/go. The .golden.txt files are what is shown.
const (
	concurrent10 = "exercises/concurrent/concurrent10/main_test.go"
	concurrent2  = "exercises/concurrent/concurrent2/main_test.go"
)

func readTestdata(t *testing.T, name string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("testdata", "race", name))
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

// checkGolden compares got with the golden file, or rewrites it with -update.
func checkGolden(t *testing.T, name, got string) {
	t.Helper()
	path := filepath.Join("testdata", "race", name+".golden.txt")
	if *update {
		if err := os.WriteFile(path, []byte(got), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	want, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if got != string(want) {
		t.Errorf("%s: output differs from the golden file (run the test with -update to see the diff in git):\n--- got\n%s\n--- want\n%s", name, got, want)
	}
}

func TestCondenseRacesGolden(t *testing.T) {
	tests := []struct {
		name  string
		input string
		path  string
	}{
		{name: "concurrent10", input: "concurrent10.input.txt", path: concurrent10},
		{name: "concurrent2", input: "concurrent2.input.txt", path: concurrent2},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, changed := condenseRaces(strings.TrimSpace(readTestdata(t, tc.input)), tc.path)
			if !changed {
				t.Fatal("nothing was condensed")
			}
			checkGolden(t, tc.name, got)
		})
	}
}

// Windows reports the same files with a drive and, depending on the Go
// version, either kind of slash. The result is the same as on Linux.
func TestCondenseRacesWithWindowsPaths(t *testing.T) {
	input := strings.TrimSpace(readTestdata(t, "concurrent10.input.txt"))
	want, _ := condenseRaces(input, concurrent10)

	forward := strings.ReplaceAll(input, "/home/learner/golings-exercises", "C:/Users/learner/golings-exercises")

	for name, windows := range map[string]string{"forward slashes": forward, "backslashes": fixBackslashes(forward)} {
		got, changed := condenseRaces(windows, `exercises\concurrent\concurrent10\main_test.go`)
		if !changed || got != want {
			t.Errorf("%s: got\n%s\nwant\n%s", name, got, want)
		}
	}
}

// fixBackslashes turns the slashes of the file paths of a report into
// backslashes and leaves the rest alone.
func fixBackslashes(report string) string {
	lines := strings.Split(report, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "      ") && strings.Contains(line, ".go:") {
			lines[i] = strings.ReplaceAll(line, "/", `\`)
		}
	}
	return strings.Join(lines, "\n")
}

// Races that hit the same two lines are shown once, with a count. The input is
// the first report of concurrent10 three times, as other goroutines would
// report it, with a different report in between.
func TestCondenseRacesCountsTheSameRaceOnce(t *testing.T) {
	input := strings.TrimSpace(readTestdata(t, "concurrent10.input.txt"))
	blocks := strings.Split(input, raceFence+"\n"+raceFence+"\n")
	if len(blocks) != 2 {
		t.Fatalf("expected the two reports of concurrent10, got %d parts", len(blocks))
	}
	first := blocks[0][strings.Index(blocks[0], raceHeader):]
	second := strings.TrimSuffix(blocks[1], raceFence)
	second = second[:strings.Index(second, "    testing.go:")]

	other := strings.NewReplacer("goroutine 12", "goroutine 31", "goroutine 10", "goroutine 29", "Goroutine 12", "Goroutine 31", "Goroutine 10", "Goroutine 29").Replace(first)
	again := strings.NewReplacer("goroutine 12", "goroutine 41", "goroutine 10", "goroutine 40", "Goroutine 12", "Goroutine 41", "Goroutine 10", "Goroutine 40").Replace(first)

	var in strings.Builder
	in.WriteString("=== RUN   TestSettingsLoadOnlyOnce\n")
	for _, report := range []string{first, second, other, again} {
		in.WriteString(raceFence + "\n" + strings.TrimSuffix(strings.TrimSpace(report), raceFence) + "\n" + raceFence + "\n")
	}
	in.WriteString("    testing.go:1867: race detected during execution of test\n--- FAIL: TestSettingsLoadOnlyOnce (0.00s)")

	got, changed := condenseRaces(in.String(), concurrent10)
	if !changed {
		t.Fatal("nothing was condensed")
	}
	if n := strings.Count(got, "Data race"); n != 2 {
		t.Errorf("%d data races shown, want 2 (one of them counted 3 times):\n%s", n, got)
	}
	if !strings.Contains(got, "Data race (×3)") {
		t.Errorf("the repeated race has no count:\n%s", got)
	}
	checkGolden(t, "duplicates", got)
}

// A report whose frames are all outside the exercise file is left as it is,
// and so is everything else in the output.
func TestCondenseRacesLeavesWhatItCannotPlace(t *testing.T) {
	input := strings.TrimSpace(readTestdata(t, "concurrent10.input.txt"))

	got, changed := condenseRaces(input, "exercises/concurrent/concurrent99/main_test.go")

	if changed || got != input {
		t.Errorf("a race outside the exercise file was changed:\n%s", got)
	}
}

// One report that can be placed and one that can't: the first is shortened,
// the second stays exactly as the detector printed it.
func TestCondenseRacesKeepsAnUnplacedReportBesideAPlacedOne(t *testing.T) {
	input := strings.TrimSpace(readTestdata(t, "concurrent10.input.txt"))
	unplaced := strings.Join([]string{
		raceFence,
		raceHeader,
		"Write at 0x00c000012345 by goroutine 7:",
		"  runtime.mapassign_faststr()",
		"      /usr/local/go/src/runtime/map_faststr.go:203 +0x40",
		"",
		"Previous write at 0x00c000012345 by goroutine 8:",
		"  runtime.mapassign_faststr()",
		"      /usr/local/go/src/runtime/map_faststr.go:203 +0x40",
		raceFence,
	}, "\n")

	got, changed := condenseRaces(unplaced+"\n"+input, concurrent10)

	if !changed {
		t.Fatal("the report in the exercise file was not shortened")
	}
	if !strings.HasPrefix(got, unplaced+"\n") {
		t.Errorf("the report outside the exercise file was changed:\n%s", got)
	}
	if !strings.Contains(got, "main_test.go:22  in (*Settings).Get") {
		t.Errorf("the other report was not shortened:\n%s", got)
	}
}

// Test output around the reports stays where it is.
func TestCondenseRacesKeepsTheRestOfTheTestOutput(t *testing.T) {
	input := strings.TrimSpace(readTestdata(t, "concurrent2.input.txt"))
	mixed := "starting the counter\n" + strings.Replace(input, "    main_test.go:15: Counter should be 100, but got 93",
		"    main_test.go:15: Counter should be 100, but got 93\n    main_test.go:16: and a second message\nprinted by the program\n", 1)

	got, changed := condenseRaces(mixed, concurrent2)

	if !changed {
		t.Fatal("nothing was condensed")
	}
	for _, kept := range []string{
		"starting the counter\n=== RUN   TestCounter\n",
		"    main_test.go:15: Counter should be 100, but got 93\n",
		"    main_test.go:16: and a second message\n",
		"printed by the program\n",
		"    race detected during execution of test\n",
		"--- FAIL: TestCounter (0.00s)",
	} {
		if !strings.Contains(got, kept) {
			t.Errorf("%q is missing from:\n%s", kept, got)
		}
	}
	if strings.Contains(got, "testing.go:1867") || strings.Contains(got, "0x00c0") {
		t.Errorf("the output still has the noise:\n%s", got)
	}
	checkGolden(t, "mixed", got)
}

func TestCondenseRacesWithoutRaces(t *testing.T) {
	input := "=== RUN   TestX\n    x_test.go:5: wrong\n--- FAIL: TestX (0.00s)"

	got, changed := condenseRaces(input, "x_test.go")

	if changed || got != input {
		t.Errorf("output without races was changed: %q", got)
	}
}

// A report cut short (no closing fence) or with a frame that makes no sense is
// shown as the detector printed it.
func TestCondenseRacesMalformedReports(t *testing.T) {
	good := strings.TrimSpace(readTestdata(t, "concurrent10.input.txt"))
	lines := strings.Split(good, "\n")

	cut := strings.Join(lines[:10], "\n")
	badFrame := strings.Replace(good, "      /home/learner/golings-exercises/exercises/concurrent/concurrent10/main_test.go:22 +0x80", "      not a location", 1)
	oneAccess := strings.Replace(good, "Previous write at", "Goroutine 99 note at", 1)

	for name, input := range map[string]string{"cut short": cut, "bad frame": badFrame, "one access": oneAccess} {
		got, _ := condenseRaces(input, concurrent10)
		if !strings.Contains(got, strings.Split(input, "\n")[2]) {
			t.Errorf("%s: the first lines of the report are gone:\n%s", name, got)
		}
		if name == "cut short" && got != input {
			t.Errorf("%s: a cut short report was changed", name)
		}
	}
}

func TestCondenseRacesAtomicAndMainGoroutine(t *testing.T) {
	input := strings.Join([]string{
		raceFence,
		raceHeader,
		"Atomic read at 0x00c000012345 by main goroutine:",
		"  command-line-arguments_test.TestX()",
		"      /w/exercises/x/x1/main_test.go:12 +0x40",
		"",
		"Previous write at 0x00c000012345 by goroutine 7:",
		"  command-line-arguments_test.TestX.func1()",
		"      /w/exercises/x/x1/main_test.go:9 +0x40",
		raceFence,
	}, "\n")

	got, changed := condenseRaces(input, "exercises/x/x1/main_test.go")

	want := "Data race: two goroutines use the same variable without synchronization\n" +
		"  atomic read     main_test.go:12  in TestX  (main goroutine)\n" +
		"  previous write  main_test.go:9  in TestX.func1  (goroutine 7)"
	if !changed || got != want {
		t.Errorf("got\n%s\nwant\n%s", got, want)
	}
}

func TestFailureShortensRacesAndKeepsTheFullOutput(t *testing.T) {
	out := strings.TrimSpace(readTestdata(t, "concurrent10.input.txt"))
	result := Result{Exercise: Exercise{Name: "concurrent10", Path: concurrent10, Mode: "test", Race: true}, Out: out}

	f := result.Failure()

	if f.Kind != TestFailure {
		t.Errorf("Kind = %v", f.Kind)
	}
	if !strings.Contains(f.Output, "Data race: two goroutines") || strings.Contains(f.Output, "0x00c0") {
		t.Errorf("the output was not shortened:\n%s", f.Output)
	}
	if !strings.Contains(f.Full, "WARNING: DATA RACE") || !strings.Contains(f.Full, "Goroutine 12 (running) created at:") {
		t.Errorf("the full output is not all there:\n%s", f.Full)
	}
}

func TestFailureWithoutAnythingToShortenHasNoFullOutput(t *testing.T) {
	result := Result{Exercise: Exercise{Path: "a/main_test.go", Mode: "test"}, Out: "--- FAIL: TestX (0.00s)\n    main_test.go:5: wrong"}

	if f := result.Failure(); f.Full != "" {
		t.Errorf("Full = %q, want it empty when nothing was shortened", f.Full)
	}
}
