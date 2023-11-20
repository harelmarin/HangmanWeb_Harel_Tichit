package hangman

type Global struct {
	Username            string
	Difficulte          string
	MotaDeviner         string
	MotCache            string
	LettreJoueur        string
	LettreDejaJoue      []string
	LettreDejaTrouve    []string
	CheckLettreJoueur   bool
	MessageEnvoi        string
	URL                 string
	TentativesRestantes int
	CheckVictoire       bool
}

var User Global
