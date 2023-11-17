package hangman

import "fmt"

// Pour chaque nombre de tentatives on affiche l'avancement du pendu grâce à un switch case.
func AffichagePendu(tentatives int) {
	switch tentatives {
	case 9:
		fmt.Println(`
		
		
		
=========`)
	case 8:
		fmt.Println(`    
	 | 
	 |  
	 |  
	 |  
==========`)
	case 7:
		fmt.Println(`
 +-------+  
	 |  
	 |  
	 |  
	 |  
	 |  
==========`)
	case 6:
		fmt.Println(` 
 +-------+  
     |   |  
	 |  
	 |  
	 |  
	 |  
==========`)
	case 5:
		fmt.Println(`  
 +-------+  
     |   |  
     O   |  
	 |  
	 |  
	 |  
==========`)
	case 4:
		fmt.Println(` 
 +-------+  
     |   |  
     O   |  
     |   |  
	 |  
	 |  
==========`)
	case 3:
		fmt.Println(`  
 +-------+  
     |   |  
     O   |  
    /|   |  
	 |  
	 |  
==========`)
	case 2:
		fmt.Println(` 
 +-------+  
     |   |  
     O   |  
    /|\  |  
	 |  
	 |  
==========`)
	case 1:
		fmt.Println(`  
 +-------+  
     |   |  
     O   |  
    /|\  |  
    /    |  
	 |  
==========`)
	}

}
