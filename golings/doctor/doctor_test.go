package doctor_test

import (
	"errors"
	"os"
	"testing"

	"github.com/cleonte/golings/golings/doctor"
	. "github.com/onsi/gomega"
)

const (
	infoFile   = "../fixtures/goversion/info.toml" // two exercises, one needing Go 99.0
	testsInfo  = "../fixtures/info.toml"           // compile1 and test1
	raceInfo   = "../fixtures/race/info.toml"      // test1 and race1, which needs the race detector
	noExercise = "../fixtures/info404.toml"
)

func healthyEnv() doctor.Env {
	return doctor.Env{
		GoVersion:    func() (string, error) { return "go1.26.1", nil },
		CGOEnabled:   func() bool { return true },
		HasCCompiler: func() bool { return true },
		FileExists:   func(string) bool { return true },
	}
}

func find(g *WithT, checks []doctor.Check, name string) doctor.Check {
	for _, c := range checks {
		if c.Name == name {
			return c
		}
	}
	g.Expect(checks).To(ContainElement(HaveField("Name", name)))
	return doctor.Check{}
}

func TestHealthySetup(t *testing.T) {
	g := NewWithT(t)
	env := healthyEnv()
	env.GoVersion = func() (string, error) { return "go1.99.0", nil }

	checks := doctor.Run(testsInfo, env)

	g.Expect(doctor.Failed(checks)).To(BeFalse())
	for _, c := range checks {
		g.Expect(c.Status).To(Equal(doctor.OK), c.Name)
	}
}

func TestGoNotInstalled(t *testing.T) {
	g := NewWithT(t)
	env := healthyEnv()
	env.GoVersion = func() (string, error) { return "", errors.New("not found") }

	checks := doctor.Run(testsInfo, env)

	c := find(g, checks, "Go toolchain")
	g.Expect(c.Status).To(Equal(doctor.Fail))
	g.Expect(c.Detail).To(ContainSubstring("on your PATH"))
	g.Expect(doctor.Failed(checks)).To(BeTrue())
}

func TestGoOlderThanFloor(t *testing.T) {
	g := NewWithT(t)
	env := healthyEnv()
	env.GoVersion = func() (string, error) { return "go1.25.9", nil }

	c := find(g, doctor.Run(testsInfo, env), "Go toolchain")

	g.Expect(c.Status).To(Equal(doctor.Fail))
	g.Expect(c.Detail).To(ContainSubstring("needs Go " + doctor.MinGoVersion))
}

func TestGoAtFloor(t *testing.T) {
	g := NewWithT(t)
	env := healthyEnv()
	env.GoVersion = func() (string, error) { return "go" + doctor.MinGoVersion + ".0", nil }

	c := find(g, doctor.Run(testsInfo, env), "Go toolchain")

	g.Expect(c.Status).To(Equal(doctor.OK))
}

func TestMissingManifestStopsAfterGoCheck(t *testing.T) {
	g := NewWithT(t)

	checks := doctor.Run(noExercise, healthyEnv())

	g.Expect(checks).To(HaveLen(2))
	c := find(g, checks, "Exercise manifest")
	g.Expect(c.Status).To(Equal(doctor.Fail))
	g.Expect(c.Detail).To(ContainSubstring("run golings from the folder"))
}

func TestMissingExerciseFiles(t *testing.T) {
	g := NewWithT(t)
	env := healthyEnv()
	env.FileExists = func(path string) bool { return path != "test/test1.go" }

	checks := doctor.Run(testsInfo, env)

	c := find(g, checks, "Exercise files")
	g.Expect(c.Status).To(Equal(doctor.Fail))
	g.Expect(c.Detail).To(And(ContainSubstring("1 missing"), ContainSubstring("test1")))
	g.Expect(doctor.Failed(checks)).To(BeTrue())
}

func TestExercisesNeedingNewerGoAreAWarning(t *testing.T) {
	g := NewWithT(t)

	checks := doctor.Run(infoFile, healthyEnv())

	c := find(g, checks, "Exercise Go versions")
	g.Expect(c.Status).To(Equal(doctor.Warn))
	g.Expect(c.Detail).To(ContainSubstring("future1 (needs Go 99.0)"))
	g.Expect(doctor.Failed(checks)).To(BeFalse(), "warnings must not fail doctor")
}

func TestRaceDetectorRequirements(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*doctor.Env)
		info   string
		want   doctor.Status
		detail string
	}{
		{"fine", func(*doctor.Env) {}, raceInfo, doctor.OK, "available, for the 1 exercises"},
		{"cgo disabled", func(e *doctor.Env) { e.CGOEnabled = func() bool { return false } }, raceInfo, doctor.Warn, "CGO_ENABLED=1"},
		{"no compiler", func(e *doctor.Env) { e.HasCCompiler = func() bool { return false } }, raceInfo, doctor.Warn, "C compiler"},
		{"macOS needs no cgo", func(e *doctor.Env) { e.GOOS, e.CGOEnabled = "darwin", func() bool { return false } }, raceInfo, doctor.OK, "available"},
		{"no exercise needs it", func(e *doctor.Env) { e.CGOEnabled = func() bool { return false } }, testsInfo, doctor.OK, "not needed"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			env := healthyEnv()
			tt.mutate(&env)

			c := find(g, doctor.Run(tt.info, env), "Race detector")

			g.Expect(c.Status).To(Equal(tt.want))
			g.Expect(c.Detail).To(ContainSubstring(tt.detail))
		})
	}
}

func TestMissingRaceDetectorNamesTheExercisesItAffects(t *testing.T) {
	g := NewWithT(t)
	env := healthyEnv()
	env.CGOEnabled = func() bool { return false }

	checks := doctor.Run(raceInfo, env)

	c := find(g, checks, "Race detector")
	g.Expect(c.Detail).To(And(ContainSubstring("1 will be skipped: race1"), Not(ContainSubstring("test1"))))
	g.Expect(doctor.Failed(checks)).To(BeFalse(), "a warning, not a failure")
}

// TestDoctorWithoutTheRaceDetector runs doctor on the real machine and the real
// manifest, on a machine that has no race detector (CGO_ENABLED=0): it has to
// warn, and name the exercises that will be skipped. Like the same test of the
// exercises package, it only runs when GOLINGS_EXPECT_NO_RACE is set.
func TestDoctorWithoutTheRaceDetector(t *testing.T) {
	if os.Getenv("GOLINGS_EXPECT_NO_RACE") == "" {
		t.Skip("set GOLINGS_EXPECT_NO_RACE=1 on a machine without the race detector (CGO_ENABLED=0) to run this")
	}
	g := NewWithT(t)

	c := find(g, doctor.Run("../../info.toml", doctor.DefaultEnv()), "Race detector")

	g.Expect(c.Status).To(Equal(doctor.Warn))
	g.Expect(c.Detail).To(And(ContainSubstring("will be skipped"), ContainSubstring("concurrent2"), ContainSubstring("concurrent10")))
}
