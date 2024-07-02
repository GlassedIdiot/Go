package models

type CmdLineOpts struct {
	Input   string
	Output  string
	Payload string
	Inject  bool
	Encode  bool
	Decode  bool
	Meta    bool
	Key     string
	Offset  string
}
