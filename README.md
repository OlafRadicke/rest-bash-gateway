# rest-bash-gateway #
Link a REST source with an bash command

# German/Deutsch #

Gelegendlich hat man bei der benutzung von Docker-Containern das 
Problem, das der Prozess in dem einen Container etwas in einem 
anderen anstossen soll. Container die miteinander verlinkt sind
können über TCP/IP (und REST) miteinander komunizieren aber nicht
Programme und Scripe gegenseitig aufrufen.

So entstand die Idee, ein einfachen Rest-Service zu bauen, dessen
Sourcen mit Bash-Scripten verknüpfbar sind. Auf diese weise,
kann der eine Kontainer im anderen über REST Scrips starten.

