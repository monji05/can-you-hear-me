package cmd

import (
	"testing"
)

func TestAddContents(t *testing.T) {
  args := []string{"test1", "test2"}

  act := AddContents(args)

  if len(act) != 2 {
    t.Errorf("Contentsは2個あるべきだが、%d件でした", len(act))
  }

  if act[0].Detail != "test1" {
    t.Errorf("最初の要素はtest1であるべきだが、%s", act[0].Detail)
  }

  if act[1].Detail != "test2" {
    t.Errorf("最初の要素はtest2であるべきだが、%s", act[1].Detail)
  }
}

