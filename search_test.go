package googlesearch_test

import (
	"context"
	"testing"

	"github.com/Zumpit/googlesearch"
)

var ctx = context.Background()

func TestSearch(t *testing.T){
  q := "dow.mike@aes.com"
  
  opts := googlesearch.SearchOptions{
	  Limit : 20 ,
  }
  returnLinks, err := googlesearch.Search(ctx,q,opts)
  if err != nil {
	  t.Errorf("something went wrong : %v", err)
      return 
	}

	if len(returnLinks) == 0 {
		t.Errorf("no results returned %v", returnLinks)
	}
}
