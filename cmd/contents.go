package cmd

type Content struct {
  Detail string `json:"detail"`
}

func AddContents(args []string) ([]Content){
  newContents := make([]Content, 0, len(args))
  for _, arg := range args {
    newContents = append(newContents, Content {
      Detail: arg,
    })
  }
  return newContents
}
