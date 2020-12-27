package api

type Entry struct {
	Meta		Meta
	Hwi			Hwi
	Fl			string
	Def			[]Def
}

type Meta struct {
	Id			string
	Stems		[]string
}

type Hwi struct {
	Prs			[]Pronun
}

type Pronun struct {
	Mw			string
}

type Def struct {
	Sseq		interface{}
}