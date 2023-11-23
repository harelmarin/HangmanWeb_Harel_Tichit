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
	TentativesRestantes int
	CheckVictoire       bool
	UserURL             string
}

var User Global
