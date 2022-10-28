package plural_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rb-go/plural-ru"
)

func Test_Plural_Noun_1(t *testing.T) {
	plu := plural.Noun(1, "день", "дня", "дней")
	assert.Equal(t, "день", plu)
}
func Test_Plural_Noun_2(t *testing.T) {
	plu := plural.Noun(2, "день", "дня", "дней")
	assert.Equal(t, "дня", plu)
}

func Test_Plural_Noun_5(t *testing.T) {
	plu := plural.Noun(5, "день", "дня", "дней")
	assert.Equal(t, "дней", plu)
}

func Test_Plural_Noun_11(t *testing.T) {
	plu := plural.Noun(11, "день", "дня", "дней")
	assert.Equal(t, "дней", plu)
}

func Test_Plural_Noun_121(t *testing.T) {
	plu := plural.Noun(121, "день", "дня", "дней")
	assert.Equal(t, "день", plu)
}

func Test_Plural_Verb_1(t *testing.T) {
	plu := plural.Verb(1, "нашелся", "нашлись", "нашлось")
	assert.Equal(t, "нашелся", plu)
}

func Test_Plural_Verb_1234(t *testing.T) {
	plu := plural.Verb(1234, "нашелся", "нашлись", "нашлось")
	assert.Equal(t, "нашлись", plu)
}

func Test_Plural_Verb_100500(t *testing.T) {
	plu := plural.Verb(100500, "нашелся", "нашлись", "нашлось")
	assert.Equal(t, "нашлось", plu)
}

func Test_Plural_Verb_10001000(t *testing.T) {
	plu := plural.Verb(10001000, "нашелся", "нашлись", "нашлось")
	assert.Equal(t, "нашлось", plu)
}

func Test_Plural_Verb_10005000(t *testing.T) {
	plu := plural.Verb(10005000, "нашелся", "нашлись", "нашлось")
	assert.Equal(t, "нашлось", plu)
}

func Test_Plural_Phrase_100500(t *testing.T) {
	cnt := 100500
	plu := fmt.Sprint(plural.Verb(cnt, "нашелся", "нашлись", "нашлось"), " ", cnt, " ", plural.Noun(cnt, "идея", "идеи", "идей"))
	assert.Equal(t, "нашлось 100500 идей", plu)
}

func Test_Plural_Phrase_123(t *testing.T) {
	cnt := 123
	plu := fmt.Sprint(plural.Verb(cnt, "нашелся", "нашлись", "нашлось"), " ", cnt, " ", plural.Noun(cnt, "идея", "идеи", "идей"))
	assert.Equal(t, "нашлись 123 идеи", plu)
}
