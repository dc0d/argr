package argr

import (
	"flag"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommon(t *testing.T) {
	assert := assert.New(t)
	input := "-v 1 -s=1 --t 1 --q 11 -w \"tt\\\"tt\" -z=\"tt\\\"tt\" -zz /tmp/dir/"
	parts := Tokenize(input)

	var argv struct {
		v, s, t, q int
		w, z, zz   string
	}
	set := flag.NewFlagSet("", flag.ExitOnError)
	set.IntVar(&argv.v, "v", -1, "")
	set.IntVar(&argv.s, "s", -1, "")
	set.IntVar(&argv.t, "t", -1, "")
	set.IntVar(&argv.q, "q", -1, "")
	set.StringVar(&argv.w, "w", "N/A", "")
	set.StringVar(&argv.z, "z", "N/A", "")
	set.StringVar(&argv.zz, "zz", "N/A", "")
	assert.NoError(set.Parse(parts))

	assert.Equal(1, argv.v)
	assert.Equal(1, argv.s)
	assert.Equal(1, argv.t)
	assert.Equal(11, argv.q)
	assert.Equal(`tt"tt`, argv.w)
	assert.Equal(`tt"tt`, argv.z)
	assert.Equal("/tmp/dir/", argv.zz)
}

func TestWithUnicode(t *testing.T) {
	assert := assert.New(t)
	sv1 := `تست`
	sv2 := `"تست"`
	input := "-v 1 -s=1 --t 1 --q 11 -w %s -z=%s -zz %s"
	input = fmt.Sprintf(input, sv2, sv2, sv1)
	parts := Tokenize(input)

	var argv struct {
		v, s, t, q int
		w, z, zz   string
	}
	set := flag.NewFlagSet("", flag.ExitOnError)
	set.IntVar(&argv.v, "v", -1, "")
	set.IntVar(&argv.s, "s", -1, "")
	set.IntVar(&argv.t, "t", -1, "")
	set.IntVar(&argv.q, "q", -1, "")
	set.StringVar(&argv.w, "w", "N/A", "")
	set.StringVar(&argv.z, "z", "N/A", "")
	set.StringVar(&argv.zz, "zz", "N/A", "")
	assert.NoError(set.Parse(parts))

	assert.Equal(1, argv.v)
	assert.Equal(1, argv.s)
	assert.Equal(1, argv.t)
	assert.Equal(11, argv.q)
	assert.Equal(sv1, argv.w)
	assert.Equal(sv1, argv.z)
	assert.Equal(sv1, argv.zz)
}
