package wts

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Script struct {
	Name         string
	Interpretter string
	Script       string
	Checksum     string
}

func (s *Script) Run(silent bool) error {
	// Create Temp File
	tempPath := filepath.Join(os.TempDir(), filepath.Base(s.Name)+strconv.Itoa(int(time.Now().UnixNano())))
	err := ioutil.WriteFile(tempPath, []byte(s.Script), 0766)
	if err != nil {
		return err
	}
	defer os.Remove(tempPath)

	if s.Interpretter == "" {
		s.Interpretter = "sh -c"
	}

	pre := strings.Split(s.Interpretter, " ")
	if len(pre) < 1 {
		return errors.New("Bad Interpretter presented")
	}
	program := pre[0]
	args := pre[1:]
	args = append(args, tempPath)

	cmd := exec.Command(program, args...)
	if !silent {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

type ScriptBuilder struct {
	S *Script
}

// Build returns the constructed Script
func (sb *ScriptBuilder) Build() (*Script, error) {
	if sb.S.Script == "" {
		return sb.S, errors.New("No script data found!")
	}
	sum := md5.Sum([]byte(sb.S.Script))
	sb.S.Checksum = hex.EncodeToString(sum[:])
	return sb.S, nil
}

// FromFile allows fetching script contents from the given file
func (sb *ScriptBuilder) FromFile(fname string) error {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}
	sb.S.Script = string(data)
	return nil
}

// Inline allows adding a script contents directly
func (sb *ScriptBuilder) Inline(script string) {
	sb.S.Script = script
}

// InterpretterPrefix allows adding an interpretter prefix
func (sb *ScriptBuilder) InterpretterPrefix(intr string) {
	sb.S.Interpretter = intr
}

func NewScriptBuilder(name string) *ScriptBuilder {
	return &ScriptBuilder{
		S: &Script{
			Name: name,
		},
	}
}
