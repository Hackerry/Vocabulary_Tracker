package api

type JS_Entry struct {
	Meta		JS_Meta
	Data		[]JS_Slug
}

type JS_Meta struct {
	Status		int
}

type JS_Slug struct {
	Slug		string
	JpWord		[]JS_JpWord `json:"japanese"`
	Senses		[]JS_Sense
}

type JS_JpWord struct {
	Word		string
	Reading		string
}

type JS_Sense struct {
	Def			[]string `json:"english_definitions"`
	Uses		[]string `json:"parts_of_speech"`
}