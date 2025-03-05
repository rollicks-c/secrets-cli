package params

import (
	"fmt"
	"strconv"
	"strings"
)

type Collector struct {
	args []string
}

func NewCollector(args []string) *Collector {
	return &Collector{
		args: args,
	}
}

func (c Collector) Validate(minCount int) error {
	if len(c.args) < minCount {
		return fmt.Errorf("at least %d args required", minCount)
	}
	return nil
}

func (c Collector) getID(index int) (int, error) {
	exp := c.args[index]
	p, err := strconv.Atoi(exp)
	if err != nil {
		return -1, err
	}
	return p, nil
}

func (c Collector) GetOptionalInteger(index int, fallback int) (int, error) {

	if index >= len(c.args) {
		return fallback, nil
	}

	return c.GetInteger(index)

}

func (c Collector) GetInteger(index int) (int, error) {

	exp := c.args[index]
	val, err := strconv.Atoi(exp)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (c Collector) GetString(index int) string {
	return c.args[index]
}

func (c Collector) GetOptionalString(index int, fallback string) string {
	if index >= len(c.args) {
		return fallback
	}
	return c.args[index]
}

func (c Collector) GetAll() []string {
	return c.args
}

func (c Collector) GetFirst() string {
	return c.args[0]
}

func (c Collector) GetLast() string {
	return c.args[len(c.args)-1]
}

func (c Collector) GetInner() []string {
	return c.args[1 : len(c.args)-1]
}

func (c Collector) GetBefore(exp string) []string {
	for i, v := range c.args {
		if v == exp {
			return c.args[:i]
		}
	}
	return c.args
}

func (c Collector) GetAfter(exp string) []string {
	for i, v := range c.args {
		if v == exp {
			return c.args[i+1:]
		}
	}
	return []string{}
}

func (c Collector) GetAllBefore(exp string) []string {
	for i, v := range c.args {
		if v == exp {
			return c.args[:i]
		}
	}
	return c.args
}

func (c Collector) GetAllAfter(exp string) []string {
	for i, v := range c.args {
		if v == exp {
			return c.args[i+1:]
		}
	}
	return []string{}
}

func (c Collector) GetWithPrefix(prefix string) []string {
	var result []string
	for _, v := range c.args {
		if !strings.HasPrefix(v, prefix) {
			continue
		}
		result = append(result, v)
	}
	return result
}

func (c Collector) GetWithoutPrefix(prefix string) []string {
	var result []string
	for _, v := range c.args {
		if strings.HasPrefix(v, prefix) {
			continue
		}
		result = append(result, v)
	}
	return result
}
