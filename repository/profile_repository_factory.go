package repository

func NewProfileRepository(filename string) ProfileRepository {
	return &JsonProfileRepository{filename: filename}
}
