package internal

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Command struct {
	Name string
	Args []string
}

// Требования
// Все ответы должны строго соответствовать спецификации RESP
//
// При пустой строке возвращать $0\r\n\r\n
//
// При отсутствии значения возвращать $-1\r\n

func ParseCommand(reader *bufio.Reader) (*Command, error) {
	command := &Command{}
	args, err := command.parseCommand(reader)
	if err != nil {
		return nil, err
	}

	if len(args) == 0 {
		return nil, errors.New("Arguments not found")
	}

	command.Name = args[0]
	command.Args = args[1:]

	return command, nil
}

func (c *Command) parseCommand(reader *bufio.Reader) ([]string, error) {
	// parse first line
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	if line[0] != '*' {
		return nil, fmt.Errorf("Unexpected format: not array *, got %c", line[0])
	}

	lengthArr := strings.TrimRight(line, "\r\n")[1:]
	length, err := strconv.Atoi(lengthArr)
	if err != nil {
		return nil, err
	}

	result := make([]string, length)

	// parse args
	for i := 0; i < length; i++ {
		element, err := c.parseElement(reader)
		if err != nil {
			return nil, err
		}
		result[i] = element
	}

	return result, nil
}

func (c *Command) parseElement(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return "nil", err
	}

	lineType := line[0]
	if lineType != '$' {
		return "nil", fmt.Errorf("Line type not string, value %c", lineType)
	}

	lengthLine := strings.TrimRight(line, "\r\n")[1:]
	length, err := strconv.Atoi(lengthLine)
	if err != nil {
		return "nil", err
	}

	argLine, err := reader.ReadString('\n')

	if err != nil {
		return "nil", err
	}

	arg := argLine[:length]

	return arg, nil
}
