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
		"highlight": map[string]interface{}{
			"pre_tags":  []string{"<font-size='red'>"},
			"post_tags": []string{"</font>"},
			"fields": map[string]interface{}{
				"title": map[string]interface{}{},
				"intro": map[string]interface{}{},
				"data":  map[string]interface{}{},
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
		//es.Search.WithDocumentType("study_notes"),
		es.Search.WithFrom(0),
		es.Search.WithSize(10),
		//es.Search.WithSeqNoPrimaryTerm(true),
		es.Search.WithSourceExcludes("is_del"),
		es.Search.WithAnalyzer("ik_smart"),
		//es.Search.WithDocvalueFields(""),
		// https://lucene.apache.org/core/2_9_4/queryparsersyntax.html
		es.Search.WithQuery("linux curl AND is_del:0 AND is_state:2"),
		es.Search.WithDefaultOperator("and"),
		//es.Search.WithExplain(true),
		//es.Search.WithSuggestText()
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithErrorTrace(),
		es.Search.WithPretty(),
	)

	if err != nil {
		panic("Error getting response: " + err.Error())
	}
	defer res.Body.Close()
	print(res.String())
}
