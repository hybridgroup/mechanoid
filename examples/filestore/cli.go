package main

import (
	"machine"
	"time"

	"encoding/binary"
	"encoding/hex"
	"strings"

	"github.com/hybridgroup/tinywasm/engine"
)

const consoleBufLen = 64

const (
	StateInput = iota
	StateEscape
	StateEscBrc
	StateCSI
)

var (
	input [consoleBufLen]byte
	state = StateInput

	commands = map[string]cmdfunc{
		"":      noop,
		"ls":    ls,
		"lsblk": lsblk,
		"load":  load,
		"save":  save,
		"rm":    rm,
		"run":   run,
		"halt":  halt,
	}
)

func cli() {
	prompt()

	for i := 0; ; {
		if console.Buffered() > 0 {
			data, _ := console.ReadByte()
			switch state {
			case StateInput:
				switch data {
				case 0x8:
					fallthrough
				case 0x7f: // this is probably wrong... works on my machine tho :)
					// backspace
					if i > 0 {
						i -= 1
						console.Write([]byte{0x8, 0x20, 0x8})
					}
				case 13:
					// return key
					if console.Buffered() > 0 {
						data, _ := console.ReadByte()
						if data != 10 {
							println("\r\nunexpected: \r", int(data))
						}
					}
					console.Write([]byte("\r\n"))
					runCommand(string(input[:i]))
					prompt()

					i = 0
					continue
				case 27:
					// escape
					state = StateEscape
				default:
					// anything else, just echo the character if it is printable
					if i < (consoleBufLen - 1) {
						console.WriteByte(data)
						input[i] = data
						i++
					}
				}
			case StateEscape:
				switch data {
				case 0x5b:
					state = StateEscBrc
				default:
					state = StateInput
				}
			default:
				// TODO: handle escape sequences
				state = StateInput
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func runCommand(line string) {
	argv := strings.SplitN(strings.TrimSpace(line), " ", -1)
	cmd := argv[0]
	cmdfn, ok := commands[cmd]
	if !ok {
		println("unknown command: " + line)
		return
	}
	cmdfn(argv)
}

func prompt() {
	print("==> ")
}

type cmdfunc func(argv []string)

func noop(argv []string) {}

func ls(argv []string) {
	list, err := eng.FileStore.List()
	if err != nil {
		println("error listing files:", err.Error())
		return
	}

	print(
		"\n-------------------------------------\r\n" +
			" File Store:  \r\n" +
			"-------------------------------------\r\n")
	for _, file := range list {
		println(file.Size(), file.Name())
	}
	print(
		"\n-------------------------------------\r\n\r\n")
}

func lsblk(argv []string) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(machine.FlashDataStart()))
	start := hex.EncodeToString(b)

	binary.BigEndian.PutUint32(b, uint32(machine.FlashDataEnd()))
	end := hex.EncodeToString(b)

	print(
		"\n-------------------------------------\r\n" +
			" Device Information:  \r\n" +
			"-------------------------------------\r\n" +
			" flash data start: 0x" + start + "\r\n" +
			" flash data end:   0x" + end + "\r\n" +
			"-------------------------------------\r\n\r\n")
}

// load module into engine
func load(argv []string) {
	println("load: " + argv[1])

	data := make([]byte, 1024)
	if err := eng.FileStore.Load(argv[1], data); err != nil {
		println("error loading file:", err.Error())
		return
	}

	if err := eng.Interpreter.Load(data); err != nil {
		println(err.Error())
		return
	}
	println("module loaded")
}

// save into filestore
func save(argv []string) {
	if len(argv) != 3 {
		println("usage: save <name> <size>")
		return
	}

	// read in size bytes from port
	sz := convertToInt(argv[2])

	data := make([]byte, sz)
	if err := readDataFromPort(data); err != nil {
		println("error reading data:", err.Error())
		return
	}

	if err := eng.FileStore.Save(argv[1], data); err != nil {
		println("error saving file:", err.Error())
	}
}

func readDataFromPort(data []byte) (err error) {
	for i := 0; i < len(data); i++ {
		for console.Buffered() == 0 {
		}
		data[i], err = console.ReadByte()
		if err != nil {
			return
		}
	}
	return nil
}

// remove from filestore
func rm(argv []string) {
	println("rm: " + argv[1])

	if err := eng.FileStore.Remove(argv[1]); err != nil {
		println("error removing file:", err.Error())
		return
	}
}

var (
	instance engine.Instance
	running  bool
	runCh    chan struct{}
)

func run(argv []string) {
	if running {
		println("already running. halt first.")
		return
	}

	println("starting...")

	// run the module
	var err error
	instance, err = eng.Interpreter.Run()
	if err != nil {
		println(err.Error())
		return
	}

	runCh = make(chan struct{})
	go pinger()

	println("running.")
}

func halt(argv []string) {
	if !running {
		println("not running")
		return
	}

	println("halting...")
	runCh <- struct{}{}
	close(runCh)
	running = false
	println("halted..")
}

func pinger() {
	running = true

	for {
		println("Ping...")
		instance.Call("ping")

		select {
		case <-runCh:
			return
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

func convertToInt(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		result = result*10 + (int(s[i]) - 48)
	}

	return result
}
