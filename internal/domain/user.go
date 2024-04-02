package domain

func (s *Service) SetUser(id int64, login string) error {
	_, err := s.repo.SetNewUser(id, login)
	return err
}

func (s *Service) GetUser(id int64) (string, error) {
	user, err := s.repo.GetUseById(id)
	return user.Login, err
}
