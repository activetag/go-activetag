package activetag

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

var (
	ActiveTagRegexp = regexp.MustCompile(`^([a-z0-9]+)-([a-z0-9]+(?:-[a-z0-9]+)*)$`)

	organizationForbiddenChars = regexp.MustCompile(`[^a-z0-9]+`)
	articleForbiddenChars      = regexp.MustCompile(`[^a-z0-9.-]+`)
	articleSeparators          = regexp.MustCompile(`[.-]+`)
)

type ActiveTag struct {
	organization string
	article      string
}

func NewActiveTag(organization, article string) ActiveTag {
	var at ActiveTag

	at.SetOrganization(organization)
	at.SetArticle(article)

	return at
}

func ParseActiveTag(s string) (ActiveTag, error) {
	if m := ActiveTagRegexp.FindStringSubmatch(s); len(m) == 3 {
		return ActiveTag{m[1], m[2]}, nil
	}

	return ActiveTag{}, fmt.Errorf("activetag not found in string \"%s\"", s)
}

func (at ActiveTag) GetArticle() string {
	return at.article
}

func (at ActiveTag) GetOrganization() string {
	return at.organization
}

func (at *ActiveTag) SetArticle(article string) {
	at.article = strings.ToLower(article)
	at.article = articleForbiddenChars.ReplaceAllString(at.article, "")
	at.article = articleSeparators.ReplaceAllString(at.article, "-")
}

func (at *ActiveTag) SetOrganization(organization string) {
	at.organization = strings.ToLower(organization)
	at.organization = organizationForbiddenChars.ReplaceAllString(at.organization, "")
}

func (at ActiveTag) String() string {
	if at.article == "" {
		return at.organization
	}

	return at.organization + "-" + at.article
}

type activeTagJSON struct {
	Organization string `json:"organization"`
	Article      string `json:"article"`
}

func (at ActiveTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(activeTagJSON{
		Organization: at.organization,
		Article:      at.article,
	})
}

func (at *ActiveTag) UnmarshalJSON(data []byte) error {
	var j activeTagJSON

	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	*at = NewActiveTag(j.Organization, j.Article)

	return nil
}
