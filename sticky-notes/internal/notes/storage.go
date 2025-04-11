package notes

var Notes = []Note{}
var NextID = 1

func Add(content string, color string) Note {
	note := Note{
		ID:      NextID,
		Content: content,
		Color:   color,
	}
	NextID++
	Notes = append(Notes, note)
	return note
}

func Delete(id int) {
	filtered := Notes[:0]
	for _, note := range Notes {
		if note.ID != id {
			filtered = append(filtered, note)
		}
	}
	Notes = filtered
}
