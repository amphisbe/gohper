package config

import (
	"testing"

	"github.com/cosiner/golib/test"
)

func TestIni(t *testing.T) {
	cfg := NewConfig(INI)
	cfg.ParseString(`
aa='ccdd'
[db]
driver=mysql
host=localhost
port=3306
user=root
password=root
database=test
config=charset=utf-8
pool_max_open=100
pool_max_idle=20
		`)
	val, _ := cfg.Val("aa")
	test.AssertEq(t, "C1", "ccdd", val)

	_, has := cfg.Val("driver")
	test.AssertFalse(t, "C2", has)

	test.AssertEq(t, "C3", cfg.DefSec(), cfg.CurrSec())

	cfg.SetCurrSec("db")
	test.AssertEq(t, "C4", 3306, cfg.IntValDef("port", 0))
}

func TestLine(t *testing.T) {
	tt := test.WrapTest(t)
	c := NewConfig(LINE)
	c.ParseString("aa=1&bb=2&&dd=123")
	tt.AssertEq("Line1", 1, c.IntValDef("aa", -1))
	tt.AssertEq("Line2", 2, c.IntValDef("bb", -1))
	tt.AssertEq("Line3", "123", c.ValDef("dd", ""))
}

func BenchmarkConf(b *testing.B) {
	cfg := NewConfig(INI)
	for i := 0; i < b.N; i++ {
		cfg.ParseString(`
aa=bb
[db]
driver=mysql
host=localhost
port=3306
user=root
password=root
database=test
config=charset=utf-8
pool_max_open=100
pool_max_idle=20aa=bb
[dba]
driver=mysql
host=localhost
port=3306
user=root
password=root
database=test
config=charset=utf-8
pool_max_open=100
pool_max_idle=20aa=bb
[dbd]
driver=mysql
host=localhost
port=3306
user=root
password=root
database=test
config=charset=utf-8
pool_max_open=100
pool_max_idle=20aa=bb
[dbz]
driver=mysql
host=localhost
port=3306
user=root
password=root
database=test
config=charset=utf-8
pool_max_open=100
pool_max_idle=20
[dbd]
driver=mysql
host=localhost
port=3306
user=root
password=root
database=test
config=charset=utf-8
pool_max_open=100
pool_max_idle=20aa=bb
[dbz]
driver=mysql
host=localhost
port=3306
user=root
password=root
database=test
config=charset=utf-8
pool_max_open=100
pool_max_idle=20
[dbd]
driver=mysql
host=localhost
port=3306
user=root
password=root
database=test
config=charset=utf-8
pool_max_open=100
pool_max_idle=20aa=bb
[dbz]
driver=mysql
host=localhost
port=3306
user=root
password=root
database=test
config=charset=utf-8
pool_max_open=100
pool_max_idle=20
		`)
	}
}
