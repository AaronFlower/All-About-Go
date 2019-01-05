package thesaurus

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Thesaurus defines a word lab interface which contains a Synonyms func.
type Thesaurus interface {
	Synonyms(term string) ([]string, error)
}

// BigHuge defines a big huge labs interface.
type BigHuge struct {
	APIKey string
}

// you should have a look the API return and then to define the structures.

type synonyms struct {
	Noun *words `json:"noun"`
	Verb *words `json:"verb"`
}

type words struct {
	Syn []string `json:"syn"`
}

// Synonyms returns the term's synonyms
func (b *BigHuge) Synonyms(term string) ([]string, error) {
	var syns []string
	response, err := http.Get("http://words.bighugelabs.com/api/2/" + b.APIKey + "/" + term + "/json")
	if err != nil {
		return syns, errors.New("bighuge: Failed when looking for synonyms for " + term + err.Error())
	}
	var data synonyms
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return syns, err
	}

	if data.Noun != nil {
		syns = append(syns, data.Noun.Syn...)
	}
	if data.Verb != nil {
		syns = append(syns, data.Verb.Syn...)
	}
	return syns, nil
}
