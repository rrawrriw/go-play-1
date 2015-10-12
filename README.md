The Daily Walk
==============

Es gilt darauf zu warten das x go-routines ihre Aufgabe erledigt haben. Da eine go-routine Beendet wird wenn das Ende der Main-Funktione erreicht wird und das Prinzip eine go-routine denn weitern Programm ablauf nicht zu blockieren würde das Programm beendet werden trozt noch aktiven go-routines. Das heißt man muss dafür sorgen das unser Programm wartet. Hierfür kann sync.WaitGroup verwendet werden. 

Beim ersten Teil dem funktioniert das einwandfrei. Doch mit fortschreittender Stunde wird alles komplizierter daher muss man zu einem kleinen Kniff greifen.Da es nun mehrer go-routinen gibt die nach dem Beenden unterschiedliche Aktionen auslösen müssen.

Hier wären einmal die Funktionen in denen Bob und Alice ihre Schuhe anziehen und auf der anderen Seite steht die Funktion Alarm. Nachdem Bob und Alice die Schuhe angezogen haben verlassen beide das Haus. Legt der Alarm los wird das Haus Explodieren (Programm beendet). Eine mögliche Lösung diese Aufgabe anzugehen wäre eine weiter sync.WaitGroup anzulegen. Wobei hier der Kniff verwendet wird das eine go-routine gestart wird welche auf die erste WaitGroup wartet in unserem Fall ist das das Schuhe anzeigen und darauf das Verlassen des Hauses. Die zweite WaitGroup würde auf den Alarm warten und danach das Programm Beenden.

Alternativ könnte man für den Alarm eine channel anlegen und darauf warten das ein Signal auf diesem gesendet wird welches zum Beenden des Programms führen würde.
