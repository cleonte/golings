package exercises

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// The race detector wraps each report in lines of equal signs:
//
//	==================
//	WARNING: DATA RACE
//	Read at 0x00c000014 by goroutine 12:
//	  pkg.(*Settings).Get()
//	      /home/me/exercises/concurrent/concurrent10/main_test.go:22 +0x80
//	...
//	==================
const (
	raceFence  = "=================="
	raceHeader = "WARNING: DATA RACE"
)

var (
	// "Read at ... by goroutine 12:", "Previous write at ... by main goroutine:".
	raceAccessHeader = regexp.MustCompile(`(?i)^(previous )?(atomic )?(write|read)(?: of size \d+)? at 0x[0-9a-f]+ by (?:goroutine (\d+)|main goroutine):$`)
	// The second line of a frame: spaces, a file and a line, maybe an offset.
	raceFrameLocation = regexp.MustCompile(`^\s{4,}(\S.*\.go):(\d+)(?: \+0x[0-9a-f]+)?$`)
	// The first line of a frame: two spaces and the function.
	raceFrameFunction = regexp.MustCompile(`^ {2}(\S.*)$`)
	// `testing.go:1867: race detected during execution of test`
	raceDetectedLine = regexp.MustCompile(`^(\s*)testing\.go:\d+: (race detected during execution of test)$`)
)

// raceAccess is one of the two memory accesses of a data race.
type raceAccess struct {
	kind      string // "read", "previous write", "atomic read"...
	goroutine string // "12", or "main" for the main goroutine
	file      string // base name of the exercise file
	line      int
	function  string
}

func (a raceAccess) where() string {
	return fmt.Sprintf("%s:%d  in %s", a.file, a.line, a.function)
}

// race is a parsed report. Races that hit the same two lines are the same race.
type race struct {
	first, second raceAccess
	count         int
}

func (r race) key() string {
	return r.first.kind + "|" + r.first.where() + "|" + r.second.kind + "|" + r.second.where()
}

// condenseRaces shortens the data race reports in the output of a test run:
// each one becomes the two accesses, as the kind of access, the line in the
// exercise file and the function, without addresses, offsets, the runtime and
// testing frames and the places where the goroutines were created. Reports of
// the same pair of lines are shown once, with a count.
//
// A report that can't be understood, or that has no frame in the exercise file
// (exercisePath), is left exactly as it was: nothing is hidden because of a
// parse error. The second result says whether anything was shortened.
func condenseRaces(output, exercisePath string) (string, bool) {
	lines := strings.Split(output, "\n")

	type piece struct {
		raw  []string // lines kept as they were, when race is nil
		race *race    // the report, when it was understood
	}
	var pieces []piece
	var plain []string
	flush := func() {
		if len(plain) > 0 {
			pieces = append(pieces, piece{raw: plain})
			plain = nil
		}
	}

	for i := 0; i < len(lines); i++ {
		if lines[i] != raceFence || i+1 >= len(lines) || lines[i+1] != raceHeader {
			plain = append(plain, lines[i])
			continue
		}

		end := -1
		for j := i + 2; j < len(lines); j++ {
			if lines[j] == raceFence {
				end = j
				break
			}
		}
		if end < 0 {
			plain = append(plain, lines[i:]...)
			break
		}

		flush()
		block := lines[i : end+1]
		if r, ok := parseRace(block, exercisePath); ok {
			pieces = append(pieces, piece{race: &r})
		} else {
			pieces = append(pieces, piece{raw: block})
		}
		i = end
	}
	flush()

	// The same pair of lines is one race, shown where it first appeared.
	first := map[string]*race{}
	condensed := false
	for _, p := range pieces {
		if p.race == nil {
			continue
		}
		condensed = true
		if seen, ok := first[p.race.key()]; ok {
			seen.count++
			continue
		}
		p.race.count = 1
		first[p.race.key()] = p.race
	}
	if !condensed {
		return output, false
	}

	var out []string
	shown := map[string]bool{}
	for _, p := range pieces {
		switch {
		case p.race == nil:
			out = append(out, p.raw...)
		case !shown[p.race.key()]:
			shown[p.race.key()] = true
			out = append(out, renderRace(*first[p.race.key()])...)
		}
	}
	for i, line := range out {
		if m := raceDetectedLine.FindStringSubmatch(line); m != nil {
			out[i] = m[1] + m[2]
		}
	}
	return strings.Join(out, "\n"), true
}

// parseRace reads one report, from its first fence to its last. It needs two
// accesses, each with a frame in the exercise file.
func parseRace(block []string, exercisePath string) (race, bool) {
	var accesses []raceAccess

	// The report is made of sections separated by blank lines.
	for i := 2; i < len(block)-1; i++ {
		m := raceAccessHeader.FindStringSubmatch(block[i])
		if m == nil {
			continue
		}

		access := raceAccess{
			kind:      strings.ToLower(strings.TrimSpace(m[1] + m[2] + m[3])),
			goroutine: "main",
		}
		if m[4] != "" {
			access.goroutine = m[4]
		}

		// The frames: a function line, then a location line.
		for i+2 < len(block) && block[i+1] != "" && block[i+1] != raceFence {
			function, location := block[i+1], block[i+2]
			i += 2
			fm := raceFrameFunction.FindStringSubmatch(function)
			lm := raceFrameLocation.FindStringSubmatch(location)
			if fm == nil || lm == nil {
				return race{}, false
			}
			if access.file == "" && inExercise(lm[1], exercisePath) {
				access.file = baseName(lm[1])
				access.line, _ = strconv.Atoi(lm[2])
				access.function = shortRaceFunction(fm[1])
			}
		}
		if access.file == "" {
			return race{}, false
		}
		accesses = append(accesses, access)
	}

	if len(accesses) != 2 {
		return race{}, false
	}
	return race{first: accesses[0], second: accesses[1]}, true
}

func renderRace(r race) []string {
	title := "Data race: two goroutines use the same variable without synchronization"
	if r.count > 1 {
		title = fmt.Sprintf("Data race (×%d): two goroutines use the same variable without synchronization", r.count)
	}

	width := max(len(r.first.kind), len(r.second.kind))
	line := func(a raceAccess) string {
		text := fmt.Sprintf("  %-*s  %s", width, a.kind, a.where())
		if r.count > 1 {
			return text
		}
		if a.goroutine == "main" {
			return text + "  (main goroutine)"
		}
		return text + "  (goroutine " + a.goroutine + ")"
	}
	return []string{title, line(r.first), line(r.second)}
}

// inExercise reports whether file is the exercise file: the same path, or a
// longer one that ends with it (the report has the absolute path, and may use
// either kind of slash).
func inExercise(file, exercisePath string) bool {
	file = strings.ToLower(slashes(file))
	want := strings.ToLower(slashes(exercisePath))
	for {
		switch {
		case strings.HasPrefix(want, "./"):
			want = want[2:]
		case strings.HasPrefix(want, "../"):
			want = want[3:]
		default:
			return want != "" && (file == want || strings.HasSuffix(file, "/"+want))
		}
	}
}

func slashes(path string) string {
	return strings.ReplaceAll(path, `\`, "/")
}

func baseName(path string) string {
	path = slashes(path)
	return path[strings.LastIndex(path, "/")+1:]
}

// shortRaceFunction turns `command-line-arguments_test.(*Settings).Get()` into
// `(*Settings).Get`, and `pkg/path.run.func1()` into `run.func1`.
func shortRaceFunction(function string) string {
	function = strings.TrimSuffix(function, "()")
	function = function[strings.LastIndex(function, "/")+1:]
	if dot := strings.Index(function, "."); dot >= 0 {
		function = function[dot+1:]
	}
	return function
}
