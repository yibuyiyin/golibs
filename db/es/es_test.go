/*
   Copyright (c) [2021] IT.SOS
   golibs is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package es

import (
	"bytes"
	"encoding/json"
	"golang.org/x/net/context"
	"testing"
)

func TestNewEs(t *testing.T) {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "中国",
			},
		},
		"highlight": map[string]interface{}{
			"pre_tags":  []string{"<font color='red'>"},
			"post_tags": []string{"</font>"},
			"fields": map[string]interface{}{
				"title": map[string]interface{}{},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		panic("Error encoding query:" + err.Error())
	}
	// Perform the search request.
	es := NewEs()
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("canal"),
		es.Search.WithDocumentType("stuty_notes"),
		es.Search.WithFrom(0),
		es.Search.WithSize(10),
		es.Search.WithIndex("canal"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)

	if err != nil {
		panic("Error getting response: " + err.Error())
	}
	defer res.Body.Close()
	print(res.String())
}
