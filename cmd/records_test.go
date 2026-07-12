package cmd

import "testing"

func TestAddRecords_UnmatchedDate(t *testing.T) {
	today := "2026-01-01"
	existingContents := []Content{
		{Detail: "test1"},
		{Detail: "test2"},
	}

	existingRecords := []Happiness{
		{Date: "2025-12-31", Contents: existingContents},
	}

	args := []string{
		"new arg1",
	}

	act := AddRecords(existingRecords, args, today)

	if len(act) != 2 {
		t.Errorf("recordは2個あるべきだが、%d個", len(act))
	}

	if act[1].Date != today {
		t.Errorf("Dateは2026-01-01であるべきだが、%s", today)
	}

	// act[0]が、existingRecordsなので
	if act[1].Contents[0].Detail != "new arg1" {
		t.Errorf("contentはnew arg1であるべきだが、%s", act[1].Contents[0])
	}

	if len(act[1].Contents) != 1 {
		t.Errorf("新しく追加したcontentは1つであるべきだが、%d件", len(act[1].Contents))
	}
}

func TestAddRecords_MatchedDate(t *testing.T) {
	today := "2026-01-01"
	existingContents := []Content{
		{Detail: "test1"},
		{Detail: "test2"},
	}

	existingRecords := []Happiness{
		{Date: "2026-01-01", Contents: existingContents},
	}

	args := []string{
		"new arg2",
		"new arg3",
	}

	act := AddRecords(existingRecords, args, today)

	if len(act) != 1 {
		t.Errorf("recordは1個あるべきだが、%d個", len(act))
	}

	if act[0].Date != today {
		t.Errorf("Dateは2026-01-01であるべきだが、%s", today)
	}

	if act[0].Contents[2].Detail != "new arg2" {
		t.Errorf("contentはnew arg2であるべきだが、%s", act[1].Contents[0].Detail)
	}

	if act[0].Contents[3].Detail != "new arg3" {
		t.Errorf("contentはnew arg3であるべきだが、%s", act[1].Contents[1].Detail)
	}

	if len(act[0].Contents) != 4 {
		t.Errorf("追加した後のcontentは4つであるべきだが、%d件", len(act[0].Contents))
	}
}
