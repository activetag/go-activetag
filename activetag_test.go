package activetag

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewActiveTag(t *testing.T) {
	at := NewActiveTag("", "")
	assert.Equal(t, at.GetOrganization(), "")
	assert.Equal(t, at.GetArticle(), "")

	at = NewActiveTag("alpha", "")
	assert.Equal(t, at.GetOrganization(), "alpha")
	assert.Equal(t, at.GetArticle(), "")

	at = NewActiveTag("alpha", "item1")
	assert.Equal(t, at.GetOrganization(), "alpha")
	assert.Equal(t, at.GetArticle(), "item1")

	at = NewActiveTag("Alpha", "Item2.SubItem")
	assert.Equal(t, at.GetOrganization(), "alpha")
	assert.Equal(t, at.GetArticle(), "item2-subitem")

	at = NewActiveTag("ALPHA", "ITEM3--subITEM...SUBSUBitem")
	assert.Equal(t, at.GetOrganization(), "alpha")
	assert.Equal(t, at.GetArticle(), "item3-subitem-subsubitem")

	at = NewActiveTag("al#pha", "it_em+4")
	assert.Equal(t, at.GetOrganization(), "alpha")
	assert.Equal(t, at.GetArticle(), "item4")
}

func TestParseActiveTag(t *testing.T) {
	at, err := ParseActiveTag("alpha-item1")
	assert.NoError(t, err)
	assert.Equal(t, at.GetOrganization(), "alpha")
	assert.Equal(t, at.GetArticle(), "item1")

	at, err = ParseActiveTag("alpha-item2-subitem")
	assert.NoError(t, err)
	assert.Equal(t, at.GetOrganization(), "alpha")
	assert.Equal(t, at.GetArticle(), "item2-subitem")

	at, err = ParseActiveTag("alpha")
	assert.Error(t, err)

	at, err = ParseActiveTag("alpha-item.1")
	assert.Error(t, err)
}

func TestString(t *testing.T) {
	at := NewActiveTag("alpha", "")
	assert.Equal(t, at.String(), "alpha")

	at = NewActiveTag("alpha", "item1")
	assert.Equal(t, at.String(), "alpha-item1")

	at = NewActiveTag("alpha", "item2-subitem")
	assert.Equal(t, at.String(), "alpha-item2-subitem")
}

func TestJSON(t *testing.T) {
	at := NewActiveTag("alpha", "item1-subitem")

	b, err := json.Marshal(at)
	assert.NoError(t, err)
	assert.Equal(t, b, []byte(`{"organization":"alpha","article":"item1-subitem"}`))

	var at2 ActiveTag

	err = json.Unmarshal(b, &at2)
	assert.NoError(t, err)
	assert.Equal(t, at2.GetOrganization(), "alpha")
	assert.Equal(t, at2.GetArticle(), "item1-subitem")

	err = json.Unmarshal([]byte("[]"), &at2)
	assert.Error(t, err)
}
