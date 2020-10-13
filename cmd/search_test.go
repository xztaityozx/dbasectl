package cmd

import (
	"github.com/stretchr/testify/assert"
	"github.com/xztaityozx/dbasectl/output"
	"testing"
)

func TestSearch_Verify(t *testing.T) {
	as := assert.New(t)

	t.Run("includeGroupなのにuserでない", func(t *testing.T) {
		s := search{includeGroup: true, user: false}
		as.Error(s.verify())
	})

	t.Run("includeGroupなのにActivateOwnerFeatureでない", func(t *testing.T) {
		s := search{includeGroup: true, user: true}
		cfg.ActivateOwnerFeature = false

		as.Error(s.verify())
	})

	t.Run("page番号が0", func(t *testing.T) {
		s := search{includeGroup: false, page: 0}
		as.Error(s.verify())
	})

	t.Run("perPageが0以下", func(t *testing.T) {
		s := search{includeGroup: false, page: 1, perPage: 0}
		as.Error(s.verify())
	})

	t.Run("perPageが200より上", func(t *testing.T) {
		s := search{includeGroup: false, page: 1, perPage: 201}
		as.Error(s.verify())
	})

	t.Run("output.FormatがYamlでもJsonでもTextでもない", func(t *testing.T) {
		s := search{includeGroup: false, page: 1, perPage: 1, output: output.Format(10000)}
		as.Error(s.verify())
	})

	t.Run("group検索とuser検索が同時にできない", func(t *testing.T) {
		s := search{user: true, group: true}
		as.Error(s.verify())
	})

	t.Run("user検索", func(t *testing.T) {
		for _, f := range []output.Format{output.Yaml, output.Json, output.Text} {
			s := search{user: true, includeGroup: false, page: 1, perPage: 1, output: f}
			as.Nil(s.verify())
		}
	})

	t.Run("user検索(includeGroup)", func(t *testing.T) {
		cfg.ActivateOwnerFeature = true
		for _, f := range []output.Format{output.Yaml, output.Json, output.Text} {
			s := search{user: true, includeGroup: true, page: 1, perPage: 1, output: f}
			as.Nil(s.verify())
		}
	})

	t.Run("group検索", func(t *testing.T) {
		for _, f := range []output.Format{output.Yaml, output.Json, output.Text} {
			s := search{user: false, group: true, includeGroup: false, page: 1, perPage: 1, output: f}
			as.Nil(s.verify())
		}
	})

	t.Run("user検索でもgroup検索でもない", func(t *testing.T) {
		for _, f := range []output.Format{output.Yaml, output.Json, output.Text} {
			s := search{user: false, includeGroup: false, page: 1, perPage: 1, output: f}
			as.Nil(s.verify())
		}
	})
}
