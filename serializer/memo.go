package serializer

import "shmily/model"

type Memo struct {
	ID      uint   `json:"ID"`
	Color   string `json:"Color"`
	Content string `json:"Content"`
}

func BuildMemo(item model.Memo) Memo {
	return Memo{
		ID:      item.ID,
		Color:   item.Color,
		Content: item.Content,
	}
}

func BuildMemos(item []model.Memo) (memos []Memo) {
	for _, item := range item {
		memo := BuildMemo(item)
		memos = append(memos, memo)
	}
	return memos
}
