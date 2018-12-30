package libcmd

import (
    "bytes"
    "errors"
    "io"
    "os/exec"
)

func CmdWrapper(args ...string) (string, string, error) {
    cmd := exec.Command(args[0], args[1:]...)

    var stdout, stderr bytes.Buffer

    cmd.Stdout = &stdout
    cmd.Stderr = &stderr

    _ = cmd.Run()

    stdoutS, stderrS := stdout.String(), stderr.String()

    if len(stderrS) > 0 {
        return stdoutS, stderrS, errors.New(stderrS)
    }

    return stdoutS, stderrS, nil
}

func PipedCmdWrapper(args ...string) (pipeStdout, pipeStderr io.ReadCloser, pipeStdin io.WriteCloser, err error) {
    cmd := exec.Command(args[0], args[1:]...)

    pipeStdout, err = cmd.StdoutPipe(); if err != nil { return }
    pipeStderr, err = cmd.StderrPipe(); if err != nil { return }
    pipeStdin,  err = cmd.StdinPipe();  if err != nil { return }

    err = cmd.Start(); if err != nil { return }

    return
}
