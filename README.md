The Daily Walk
==============

Die Lösung für die [Aufgabe 1](http://whipperstacker.com/2015/10/05/3-trivial-concurrency-exercises-for-the-confused-newbie-gopher/?utm_source=golangweekly&utm_medium=email)

Es gilt darauf zu warten das x go-routines ihre Aufgabe erledigt haben. Da eine go Programm beendet wird wenn das Ende der Main-Funktione erreicht ist und go-routines nebenläufig ihr Leben leben muss man sich überlegen was das Programm beenden soll und wie man das Programm dazu bring zu warten. Hierfür kann sync.WaitGroup verwendet werden. 

Beim ersten Teil der Aufgabe funktioniert das einwandfrei. Doch mit fortschreittender Stunde wird alles komplizierter daher muss man beim 2. Teil der Aufgabe zu einem kleinen Kniff greifen. Da es nun mehrer go-routinen gibt die nach dem Beenden unterschiedliche Aktionen auslösen müssen.

Hier wären einmal die Funktionen in denen Bob und Alice ihre Schuhe anziehen und auf der anderen Seite steht die Funktion für den Alarm. Nachdem Bob und Alice die Schuhe angezogen haben verlassen beide das Haus. Legt der Alarm los wird das Haus Explodieren (Programm beendet). Eine mögliche Lösung diese Aufgabe anzugehen wäre eine weiter sync.WaitGroup anzulegen. Wobei hier der Kniff verwendet wird das eine go-routine gestart wird welche auf die erste WaitGroup wartet in unserem Fall ist das dass Schuhe anziehen und darauf das Verlassen des Hauses. Die zweite WaitGroup würde auf den Alarm warten und danach das Programm Beenden.

Alternativ könnte man für den Alarm eine channel anlegen und darauf warten das ein Signal auf diesem gesendet wird welches zum Beenden des Programms führen würde.
