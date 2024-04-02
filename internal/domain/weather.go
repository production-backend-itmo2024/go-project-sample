package domain

func (s *Service) GetTemp() (int64, error) {
	resp, err := s.weatherClient.GetCurrentWeather()
	return resp.Temp, err
}
