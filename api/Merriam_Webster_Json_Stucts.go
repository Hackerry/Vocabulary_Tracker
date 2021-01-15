package api

type MW_Entry struct {
	Meta		MW_Meta
	Hwi			MW_Hwi
	Fl			string
	Def			[]MW_Def
}

type MW_Meta struct {
	Id			string
	Stems		[]string
}

type MW_Hwi struct {
	Prs			[]MW_Pronun
}

type MW_Pronun struct {
	Mw			string
}

type MW_Def struct {
	Sseq		interface{}
}