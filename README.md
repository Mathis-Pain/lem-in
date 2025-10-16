# 🐜 Lem-in

📖 Description du projet

Lem-in est un projet algorithmique consistant à gérer le déplacement d’une colonie de fourmis d’un point de départ (start) à un point d’arrivée (end) à travers un réseau de salles reliées par des tunnels.
L’objectif est de faire arriver toutes les fourmis à destination en un minimum de tours, en optimisant le flux et la répartition sur plusieurs chemins.

## 🧩 Objectif

Trouver la meilleure stratégie pour répartir les fourmis sur différents chemins afin que :

- Chaque salle (sauf start et end) ne contienne qu’une seule fourmi à la fois.
- Les fourmis empruntent les chemins les plus efficaces selon leur longueur.
- Le nombre total de tours nécessaires pour que toutes les fourmis atteignent l’arrivée soit minimal.

## ⚙️ Fonctionnement général

1 Le programme reçoit en entrée une description du réseau :

- Le nombre de fourmis.
- La liste des salles (avec leurs coordonnées).
- Les tunnels reliant les salles.

2 Il construit un graphe représentant les connexions.
3 Il trouve tous les chemins possibles du start à l’end.
4 Il répartit les fourmis entre ces chemins selon leur longueur et capacité.
5 Il simule le déplacement des fourmis, tour par tour, jusqu’à ce qu’elles soient toutes arrivées.

## 🧠 Concepts clés

- Graphe orienté non pondéré : les salles sont les nœuds, les tunnels sont les arêtes.
- Algorithme de recherche de chemins : basé sur du BFS (Breadth-First Search)
- Optimisation du flux : l’idée est proche du problème du flot maximum basé sur du Greedy.
- Simulation : on affiche le déplacement de chaque fourmi à chaque tour.

## Pour lancer le programme

go run . examples/example00.txt
go run . examples/example01.txt
go run . examples/example02.txt
go run . examples/example03.txt
go run . examples/example05.txt
go run . examples/example06.txt
go run . examples/example07.txt

## si on veut le le visualizer :

go run . examples/example01.txt --viz
