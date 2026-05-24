package cmd

func AddRecords(records []Happiness, contents []Content, today string) ([]Happiness) {
  isTodayFlg := false
  for index, record := range records {
    if today == record.Date {
      isTodayFlg = true
      records[index].Contents = append(records[index].Contents, contents...)
      records[index].Count = len(records[index].Contents)
      break
    }
  }

  if !isTodayFlg {
    return append(records, Happiness {
      Date: today,
      Contents: contents,
      Count: len(contents),
    })
  }

  return records
}

