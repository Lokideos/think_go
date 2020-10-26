package fixture

// Scanner structure for Scan function receiver
type Scanner struct{}

// Scan fixture for moking Scan method
func (s *Scanner) Scan(url string, depth int) (map[string]string, error) {
	return map[string]string{
		"http://info.cern.ch/":           "First Internet Web-site",
		"http://info.cern.ch/hypertext/": "About Hypertext Transfer Protocol",
	}, nil
}
