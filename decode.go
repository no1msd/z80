package z80

import (
	"encoding/json"
	"fmt"
	"io"
	"math/bits"
	"sync"
)

var allOPCodes = [][]*OPCode{
	load8,
	load16,
	exbtsg,
	arith8,
	ctrl,
	arith16,
	rotateshift,
	bitop,
	jump,
	callret,
	inout,
}

type decodeLayer struct {
	nodes   map[uint8]*decodeNode
	any     bool
	anyNode *decodeNode
}

var defaultDecoder *decodeLayer
var decoderOnce sync.Once

func defaultDecodeLayer() *decodeLayer {
	decoderOnce.Do(func() {
		defaultDecoder = newDecodeLayer(0, allOPCodes...)
	})
	return defaultDecoder
}

func newDecodeLayer(level int, opcodes ...[]*OPCode) *decodeLayer {
	l := &decodeLayer{
		nodes: map[uint8]*decodeNode{},
		any:   true,
	}
	for _, cc := range opcodes {
		for _, opcode := range cc {
			if level < len(opcode.C) {
				l.put(level, opcode)
			}
		}
	}
	if l.any {
		n := l.nodes[0]
		n.next = newDecodeLayer(level+1, n.codes)
		n.codes = nil
		l.anyNode = n
		l.nodes = nil
		return l
	}
	for _, n := range l.nodes {
		if n == nil {
			continue
		}
		switch len(n.codes) {
		case 0:
			continue
		case 1:
			n.opcode = n.codes[0]
			continue
		}
		n.next = newDecodeLayer(level+1, n.codes)
		n.codes = nil
	}
	return l
}

func (l *decodeLayer) put(level int, opcode *OPCode) {
	c := opcode.C[level]
	if l.any && (c.C != 0x00 || c.M != 0xff) {
		l.any = false
	}
	s := bits.OnesCount8(^c.M)
	for i, e := c.beginEnd(); i <= e; i++ {
		d := uint8(i)
		if !c.match(d) {
			continue
		}
		n, ok := l.nodes[d]
		if !ok || s > n.score {
			n = &decodeNode{score: s}
			l.nodes[d] = n
		}
		n.codes = append(n.codes, opcode)
	}
}

func (l *decodeLayer) get(d uint8) *decodeNode {
	if l.anyNode != nil {
		return l.anyNode
	}
	n, ok := l.nodes[d]
	if !ok {
		return nil
	}
	return n
}

func (l *decodeLayer) mapTo() map[string]interface{} {
	m := map[string]interface{}{}
	if l.anyNode != nil {
		m["*"] = l.anyNode.mapTo()
		return m
	}
	for i, n := range l.nodes {
		if n == nil {
			continue
		}
		m[fmt.Sprintf("%02X", i)] = n.mapTo()
	}
	return m
}

type decodeNode struct {
	score  int
	opcode *OPCode
	next   *decodeLayer
	codes  []*OPCode
}

func (n *decodeNode) mapTo() map[string]interface{} {
	m := map[string]interface{}{}
	m["score"] = n.score
	if n.opcode != nil {
		m["opcode"] = n.opcode.mapTo()
	}
	if n.next != nil {
		m["next"] = n.next.mapTo()
	}
	if len(n.codes) >= 2 {
		m["codes"] = len(n.codes)
	}
	return m
}

// DumpDecodeLayer dumps default decode layer to io.Writer for debugging
func DumpDecodeLayer(w io.Writer) error {
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	return e.Encode(defaultDecodeLayer().mapTo())
}
