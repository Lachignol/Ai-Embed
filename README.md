# Ai-Embed



####  Ce projet démontre l'utilisation du framework Ollama avec le modèle d'IA Mistral pour générer des réponses. Nous pouvons fournir une multitude d'informations qui sont intégrées dans des fichiers texte placés dans un dossier nommé embedRepo. Ces informations sont ensuite lues et stockées dans une base de données bbolt. À partir de cette base de données et du prompt de l'utilisateur, nous recherchons les similarités afin de fournir une réponse pertinente.

## Démonstration

#### Première Exécution : Lors de la première exécution, je demande "Quel est le palmarès de Jon Jones (combattant de MMA) ?" et il n'a pas l'information dans son contenu (donc dans les fichiers ajoutés à la BDD).

![Capture d’écran Ai-Embed-1](https://github.com/user-attachments/assets/1a0852fd-fbcf-443c-b1f4-c35920b642c2)

####  Deuxième Exécution : Après l'ajout des informations dans le fichier resume.txt (je sais, le nom n'est pas très bien choisi 😊) placé dans mon répertoire embedRepo puis donc aussi dans la BDD, il est bien capable de me donner une réponse !

![Capture d’écran Ai-Embed-2](https://github.com/user-attachments/assets/b4b7f7b3-2579-40a6-8b64-63236310a61f)


## Remarques
J'ai notifié au système de se baser sur les documents que je lui ai fournis pour que l'exemple soit plus parlant, car il aurait dans ce cas précis pu récupérer l'information sur internet.

Projet avec les outils suivants :
Une librairie en golang du nom de Parakeet,
Ollama et Mistral IA ,
BDD :BBOLT,
et Fichiers Embed.
