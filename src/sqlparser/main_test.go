package sqlparser

import "testing"

func TestSelect(t *testing.T) {
    var sql = "SELECT * FROM a"
    p := &Peg{Buffer: string(sql)}
    p.Init()
    if err := p.Parse(); err != nil {
        t.Errorf("Parse error: %s", err)
    }
}