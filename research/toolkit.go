package research

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"strings"
)

func (t *Tree[T]) String() string {
	if isEmpty(t.Root) {
		return "null"
	}

	return formatNode(t.Root, 0)
}

func formatNode[T constraints.Ordered](n *Node[T], indent int) string {
	if n == nil {
		return "null"
	}

	indentStr := strings.Repeat("  ", indent)
	nextIndentStr := strings.Repeat("  ", indent+1)

	var sb strings.Builder

	sb.WriteString("{\n")

	sb.WriteString(nextIndentStr)
	sb.WriteString(`"value": `)
	sb.WriteString(fmt.Sprintf("%v", n.Value))
	sb.WriteString(",\n")

	sb.WriteString(nextIndentStr)
	sb.WriteString(`"left": `)
	sb.WriteString(formatNode(n.Left, indent+1))
	sb.WriteString(",\n")

	sb.WriteString(nextIndentStr)
	sb.WriteString(`"right": `)
	sb.WriteString(formatNode(n.Right, indent+1))
	sb.WriteString("\n")

	sb.WriteString(indentStr)
	sb.WriteString("}")

	return sb.String()
}
