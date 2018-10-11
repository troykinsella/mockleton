package config

const (
	FileName = "Mockletonfile"
)

type Mockletonfile struct {
	Vars     map[string]interface{} `yaml:"vars"`
	Runfile  *Runfile               `yaml:"runfile"`
	Sequence Sequence               `yaml:"sequence"`
	Report   *Report                `yaml:"report"`
}

type Runfile struct {
	File   string
	Format string
}

type Sequence []Step

type Step struct {
	Label string `yaml:"label"`
	//Expect Expect `yaml:"expect"`
	Output *Output `yaml:"output"`
}

//type Expect yomega.Spec

type Output struct {
	ExitCode int        `yaml:"exit-code"`
	Print    []OutputOp `yaml:"print"`
	// or
	Exec string `yaml:"exec"`
}

type OutputOp map[string]string

type Report struct {
	File   string
	Format string
}

func NewMockletonfile() *Mockletonfile {
	return fillNils(&Mockletonfile{})
}

func fillNils(mf *Mockletonfile) *Mockletonfile {
	if mf.Sequence == nil {
		mf.Sequence = make(Sequence, 0)
	}

	if mf.Report == nil {
		mf.Report = &Report{}
	}

	if mf.Runfile == nil {
		mf.Runfile = &Runfile{}
	}

	if mf.Vars == nil {
		mf.Vars = make(map[string]interface{})
	}

	return mf
}
