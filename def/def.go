package def

type Def struct {
	Mode       string   `yaml:"mode"`
	YoutubeUrl string   `yaml:"youtube_url"`
	Head       Head     `yaml:"head"`
	Timing     Timing   `yaml:"timing"`
	Length     string   `yaml:"length"`
	Exports    []Export `yaml:"exports"`
}

type Head struct {
	Artist  string `yaml:"artist"`
	Album   string `yaml:"album"`
	Title   string `yaml:"title"`
	Author  string `yaml:"author"`
	By      string `yaml:"by"`
	Version string `yaml:"version"`
}

type Timing struct {
	Bpm          int     `yaml:"bpm"`
	Bpb          int     `yaml:"bpb"`
	GlobalOffset float32 `yaml:"global_offset"`
}

type Export struct {
	Filename string  `yaml:"filename"`
	Length   string  `yaml:"length"`
	Offset   float32 `yaml:"offset"`
}
