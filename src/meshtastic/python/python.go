package python

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"service/meshtastic"
	"service/models"
	"strings"
)

type pythonCLI struct {
	binaryPath string
}

// NewCLI - возвращает проинициализированный python cli для работы с meshtastic
func NewCLI(binaryPath string) meshtastic.Meshtastic {
	return &pythonCLI{binaryPath: binaryPath}
}

func (m *pythonCLI) ListNodes() (nodes map[string]models.Node, err error) {
	nodes = map[string]models.Node{}

	cmd := exec.Command(m.binaryPath, "--info")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err = cmd.Run()
	if err != nil {
		err = fmt.Errorf("failed to run pythonCLI command: %v", err)
		return
	}

	scanner := bufio.NewScanner(&stdout)
	for scanner.Scan() {
		node := models.Node{}
		jsonString := strings.ReplaceAll(scanner.Text(), "'", "\"")
		err = json.Unmarshal([]byte(jsonString), &node)
		if err != nil {
			err = nil
			continue
		}
		nodes[node.User.LongName] = node
	}

	return
}
