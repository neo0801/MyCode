package retriever

// Retriever retrieves url
type Retriever struct {
	Contents string
}

// Get url contents
func (r *Retriever) Get(url string) string {
	return url + " is " + r.Contents + "\n"
}
