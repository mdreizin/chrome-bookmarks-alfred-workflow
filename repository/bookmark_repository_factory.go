package repository

func NewBookmarkRepository(filename string) BookmarkRepository {
	return &JsonBookmarkRepository{filename: filename}
}
