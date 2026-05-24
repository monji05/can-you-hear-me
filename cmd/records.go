package cmd


type Content struct {
  Detail string `json:"detail"`
}

func AddRecords(records []Happiness, args []string, today string) ([]Happiness) {
  contents := addContents(args)

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


func addContents(args []string) ([]Content){
  newContents := make([]Content, 0, len(args))
  for _, arg := range args {
    newContents = append(newContents, Content {
      Detail: arg,
    })
  }
  return newContents
}
