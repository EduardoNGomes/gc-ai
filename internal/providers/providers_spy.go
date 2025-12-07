package providers

type SelectAgentSpy struct {
	Option Provider
}

func (s *SelectAgentSpy) SetAgent(p Provider) {
	s.Option = p
}

func (s *SelectAgentSpy) SelectedOption() (Provider, error) {
	return s.Option, nil
}

func NewSelectAgentSpy() *SelectAgentSpy {
	return &SelectAgentSpy{}
}
