package service

func NewBookmarkService(config map[string]string) BookmarkService {
	return &DefaultBookmarkService{config: config}
}
