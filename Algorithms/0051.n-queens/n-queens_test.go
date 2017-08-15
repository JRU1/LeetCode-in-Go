package Problem0051

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para 是参数
// one 代表第一个参数
type para struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one [][]string
}

func Test_Problem0051(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{4},
			ans{[][]string{
				[]string{".Q..", "...Q", "Q...", "..Q."},
				[]string{"..Q.", "Q...", "...Q", ".Q.."},
			}},
		},

		question{
			para{5},
			ans{[][]string{
				[]string{"Q....", "..Q..", "....Q", ".Q...", "...Q."},
				[]string{"Q....", "...Q.", ".Q...", "....Q", "..Q.."},
				[]string{".Q...", "...Q.", "Q....", "..Q..", "....Q"},
				[]string{".Q...", "....Q", "..Q..", "Q....", "...Q."},
				[]string{"..Q..", "Q....", "...Q.", ".Q...", "....Q"},
				[]string{"..Q..", "....Q", ".Q...", "...Q.", "Q...."},
				[]string{"...Q.", "Q....", "..Q..", "....Q", ".Q..."},
				[]string{"...Q.", ".Q...", "....Q", "..Q..", "Q...."},
				[]string{"....Q", ".Q...", "...Q.", "Q....", "..Q.."},
				[]string{"....Q", "..Q..", "Q....", "...Q.", ".Q..."},
			}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, solveNQueens(p.n), "输入:%v", p)
	}
}
