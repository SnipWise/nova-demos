# Problèmes Hydraulique

**Catégorie :** Hydraulique
**Nombre de Facts :** 15
**Retour à l'index :** [Knowledge_Base_Chaudieres.md](Knowledge_Base_Chaudieres.md)

---

## FACT-CHAUD-036: Pression d'eau instable, chute fréquente

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-036 |
| **Catégorie** | Hydraulique |
| **Système** | Circuit chauffage / Pressurisation |
| **Gravité** | **Moyenne** |
| **Marques** | Multi-marques (Saunier Duval, Vaillant, Frisquet, De Dietrich) |

**Symptômes :**
- Pression chute régulièrement sous 1 bar (témoin rouge allumé)
- Nécessité de rajouter de l'eau fréquemment (plusieurs fois par semaine/mois)
- Arrêts intempestifs de la chaudière par sécurité manque d'eau
- Manomètre indique une baisse progressive de pression

**Cause racine probable :**
Fuite au circuit chauffage (micro-fuites radiateurs, raccords, vannes), vase d'expansion dégonflé ou HS, soupape de sécurité qui purge, défaut d'étanchéité échangeur primaire ou corps de chauffe.

**Étapes de résolution :**

1. **Contrôle visuel des fuites apparentes**
   - Inspecter sous la chaudière, au niveau du corps de chauffe et des raccordements
   - Vérifier tous les radiateurs (purgeurs, robinets, raccords)
   - Examiner les vannes du circuit (3 voies, mélangeuses)
   - Vérifier l'évacuation de la soupape de sécurité 3 bars (présence d'eau)

2. **Test du vase d'expansion**
   - Couper la chaudière et fermer les vannes d'isolement circuit chauffage
   - Vidanger légèrement le circuit (pression à 0,5 bar)
   - Accéder à la valve Schrader du vase d'expansion (souvent derrière ou sous la chaudière)
   - Appuyer sur la valve : si de l'eau sort, le vase est HS (membrane percée)
   - Mesurer la pression azote avec manomètre adapté : doit être 0,8-1 bar vase vide
   - Si pression nulle ou inférieure à 0,5 bar, regonfler ou remplacer le vase

3. **Recherche de micro-fuites invisibles**
   - Ajouter un traceur colorant spécial chauffage dans le circuit
   - Faire fonctionner 24-48h puis inspecter à nouveau tous les points
   - Utiliser un détecteur d'humidité ou caméra thermique si disponible
   - Vérifier les zones difficiles d'accès (circuits enterrés, sous plancher)

4. **Contrôle du remplissage et de la soupape de sécurité**
   - Vérifier que le robinet de remplissage ferme bien (pas de goutte à goutte interne)
   - Tester la soupape de sécurité 3 bars : doit être étanche jusqu'à 3 bars
   - Si la soupape coule, la remplacer (voir FACT-CHAUD-045)

5. **Vérification de l'échangeur primaire**
   - Si aucune fuite externe trouvée, suspecter fuite interne échangeur
   - Sur chaudières mixtes : vérifier si eau chaude sanitaire trouble ou pression ECS anormale
   - Test spécifique : isoler le circuit chauffage, maintenir en pression, observer si chute persiste

6. **Repressurisation correcte du circuit**
   - Amener la pression à 1,2-1,5 bar à froid
   - Vérifier que la pression monte à 1,8-2,2 bar en chauffe (normal)
   - Ne jamais dépasser 2,5 bars pour éviter déclenchement soupape

**Prévention :**
- Contrôler la pression manomètre mensuellement
- Vérifier l'état du vase d'expansion annuellement
- Purger les radiateurs en début de saison de chauffe
- Faire un contrôle d'étanchéité lors de l'entretien annuel

**Avertissements sécurité :**
- Ne jamais faire fonctionner la chaudière avec une pression inférieure à 0,8 bar (risque de surchauffe, destruction échangeur)
- Ne pas rajouter d'eau trop fréquemment sans identifier la cause (apport d'oxygène = corrosion accélérée)

---

## FACT-CHAUD-037: Vase d'expansion dégonflé ou HS

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-037 |
| **Catégorie** | Hydraulique |
| **Système** | Vase d'expansion / Pressurisation |
| **Gravité** | **Moyenne** |
| **Marques** | Multi-marques (tous types de chaudières) |

**Symptômes :**
- Pression circuit monte très vite en chauffe (dépasse 2,5-3 bars)
- Déclenchement fréquent de la soupape de sécurité 3 bars
- Pression instable, variations importantes chaud/froid
- Eau s'écoule par la soupape de sécurité lors de la montée en température

**Cause racine probable :**
Vase d'expansion dégonflé (perte de pression azote côté gaz), membrane du vase percée (eau côté gaz), vase sous-dimensionné pour l'installation, vase bouché ou encrassé.

**Étapes de résolution :**

1. **Diagnostic de l'état du vase**
   - Couper la chaudière, attendre refroidissement complet
   - Fermer les vannes d'isolement du circuit chauffage si présentes
   - Vidanger partiellement le circuit (pression 0,5 bar ou moins)
   - Localiser la valve Schrader du vase (généralement derrière la chaudière ou en partie basse)

2. **Test de la membrane**
   - Dévisser le capuchon de la valve Schrader
   - Appuyer brièvement sur la valve avec un tournevis fin
   - **Si de l'eau sort** : membrane percée, vase HS → **remplacement obligatoire**
   - **Si de l'air sort** : membrane OK, passer à l'étape suivante

3. **Mesure de la pression azote**
   - Utiliser un manomètre basse pression avec adaptateur Schrader
   - Pression correcte : 0,8 à 1 bar (circuit vidangé)
   - Si pression < 0,5 bar : vase dégonflé, nécessite regonflage ou remplacement

4. **Regonflage du vase (si membrane OK)**
   - S'assurer que le circuit est bien vidangé (pression 0 bar côté eau)
   - Utiliser une pompe à vélo ou compresseur avec régulateur
   - Gonfler progressivement jusqu'à 0,8-1 bar (selon hauteur statique installation)
   - Formule pression vase : Pstat = (hauteur installation en mètres / 10) + 0,3 bar
   - Exemple : pour 5m de hauteur → 0,5/10 + 0,3 = 0,8 bar

5. **Remplacement du vase si nécessaire**
   - Identifier le volume nécessaire (généralement 8-12L pour maison individuelle)
   - Formule volume vase : V = (volume circuit × coefficient dilatation) / (Pmax - Pstat) / Pmax
   - Couper eau et électricité, vidanger complètement le circuit
   - Dévisser l'ancien vase, nettoyer le raccord
   - Installer le nouveau vase, préalablement gonflé à la bonne pression
   - Utiliser du téflon ou pâte d'étanchéité sur le filetage

6. **Remise en service**
   - Remplir le circuit lentement en purgeant l'air
   - Amener à pression de service 1,2-1,5 bar à froid
   - Mettre en chauffe et vérifier que la pression monte normalement à 1,8-2,2 bar
   - Vérifier qu'il n'y a pas de fuite au niveau du vase
   - Tester sur plusieurs cycles de chauffe

**Prévention :**
- Contrôler la pression du vase annuellement (lors de l'entretien)
- Vérifier que la pression circuit ne dépasse jamais 2,5 bars en chauffe
- Éviter les vidanges complètes du circuit (usure de la membrane)

**Spécificités techniques :**
- Vases à membrane : durée de vie moyenne 7-10 ans
- Pression maximale vase : généralement 3-4 bars (vérifier plaque signalétique)
- Ne jamais dépasser 2 bars de pression azote dans un vase standard

**Avertissements sécurité :**
- Un vase HS peut entraîner une surpression dangereuse du circuit
- Toujours vidanger avant de travailler sur le vase (risque de brûlure)

---

## FACT-CHAUD-038: Fuite circuit chauffage

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-038 |
| **Catégorie** | Hydraulique |
| **Système** | Circuit chauffage / Étanchéité |
| **Gravité** | **Moyenne à Élevée** |
| **Marques** | Multi-marques (tous systèmes de chauffage) |

**Symptômes :**
- Traces d'eau sous la chaudière, au niveau des radiateurs ou des tuyaux
- Chute régulière de la pression du circuit
- Taches d'humidité au plafond ou aux murs (circuits encastrés)
- Corrosion visible sur les éléments métalliques du circuit

**Cause racine probable :**
Joint défectueux (presse-étoupe, raccord), corrosion du corps de chauffe ou échangeur, fissure sur radiateur ou tuyauterie, mauvais serrage de raccord, gel ayant endommagé une canalisation.

**Étapes de résolution :**

1. **Localisation de la fuite**
   - Inspecter visuellement sous la chaudière (corps de chauffe, échangeur, raccords)
   - Vérifier tous les radiateurs (bas, haut, robinets, purgeurs, bouchons)
   - Examiner les raccords de tuyauterie visibles
   - Utiliser du papier absorbant pour identifier les traces d'humidité
   - Si fuite invisible : ajouter un traceur fluorescent et utiliser lampe UV

2. **Fuite au niveau de la chaudière**
   - **Échangeur primaire/corps de chauffe** : si fuite sur l'échangeur même, remplacement nécessaire (pièce coûteuse, évaluer réparation vs remplacement chaudière)
   - **Raccords/presse-étoupes** : resserrer les écrous, remplacer les joints toriques
   - **Vanne 3 voies** : vérifier les joints de tige, remplacer si nécessaire
   - **Soupape de sécurité** : si elle coule, voir FACT-CHAUD-045

3. **Fuite sur radiateur**
   - **Purgeur** : resserrer ou remplacer (démonter, téflon, remonter)
   - **Robinet thermostatique** : resserrer l'écrou de presse-étoupe ou remplacer le joint
   - **Bouchon** : démonter, refaire l'étanchéité avec filasse + pâte ou téflon, remonter
   - **Fissure radiateur** : colmatage temporaire possible (mastic bi-composant spécial chauffage) mais remplacement recommandé

4. **Fuite sur tuyauterie**
   - **Raccord mécanique** : resserrer, remplacer le joint ou l'olive
   - **Raccord fileté** : démonter, refaire étanchéité (filasse chanvre + pâte ou téflon), remonter
   - **Soudure cuivre** : brasage à refaire (couper eau, vidanger section, décaper, braser, tester)
   - **Tube percé/fissuré** : remplacement de la section (manchon cuivre ou raccord rapide)

5. **Fuite encastrée (mur, sol)**
   - Isoler la section concernée si possible (vannes d'isolement)
   - Détecter précisément l'emplacement (caméra thermique, détecteur d'humidité)
   - Si circuit PER/multicouche : possibilité de tirer un nouveau tube sans casser
   - Si circuit cuivre : création d'un nouveau cheminement ou réparation par ouverture
   - Envisager la pose de gaines techniques pour éviter future destruction

6. **Réparation d'urgence temporaire**
   - Produit d'étanchéité liquide (type "stop fuite") : à utiliser en dernier recours, peut colmater micro-fuites temporairement
   - Collier de serrage + joint caoutchouc sur tube
   - Mastic bi-composant époxy spécial chauffage
   - **Ces solutions sont temporaires** : planifier réparation définitive rapidement

7. **Contrôle et test après réparation**
   - Remplir le circuit, purger l'air soigneusement
   - Monter en pression à 2 bars pour test d'étanchéité
   - Mettre en chauffe et vérifier à température maximale (3 bars)
   - Surveiller pendant 24-48h avant de considérer la réparation validée

**Prévention :**
- Maintenir un pH équilibré de l'eau du circuit (pH 7-8,5)
- Utiliser un inhibiteur de corrosion
- Éviter les vidanges fréquentes (apport d'oxygène = corrosion)
- Contrôler régulièrement les presse-étoupes et raccords

**Avertissements sécurité :**
- Toujours couper l'alimentation électrique avant intervention
- Vidanger et dépressuriser avant de démonter un élément
- Attention aux brûlures : attendre refroidissement complet du circuit
- Utiliser un brasage sans plomb pour eau sanitaire

---

## FACT-CHAUD-039: Pompe circulation bruyante

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-039 |
| **Catégorie** | Hydraulique |
| **Système** | Circulateur / Pompe |
| **Gravité** | **Faible à Moyenne** |
| **Marques** | Grundfos, Wilo, DAB, Salmson (tous circulateurs) |

**Symptômes :**
- Bruit de roulement, grincement ou vibration au niveau du circulateur
- Bruit de cavitation (gargouillis, bulles d'air)
- Sifflement aigu lors du fonctionnement
- Vibrations transmises à la tuyauterie

**Cause racine probable :**
Air dans le corps de pompe, roulements usés, axe rotor encrassé ou grippé, vitesse de rotation trop élevée, cavitation due à une pression insuffisante, corps de pompe entartré.

**Étapes de résolution :**

1. **Identification du type de bruit**
   - **Gargouillis/bulles** : présence d'air dans le circulateur
   - **Grincement/frottement** : roulements usés ou axe grippé
   - **Sifflement** : vitesse excessive ou cavitation
   - **Vibration** : fixation insuffisante ou déséquilibre rotor

2. **Purge du circulateur (si bruit d'air)**
   - Mettre la chaudière en chauffe pour activer le circulateur
   - Localiser la vis de purge sur le corps du circulateur (généralement sur le côté ou devant)
   - Placer un récipient dessous
   - Dévisser lentement la vis avec un tournevis plat (1/4 de tour)
   - Laisser l'eau s'écouler jusqu'à ce qu'elle soit continue sans bulles d'air
   - Revisser la vis fermement (sans forcer)
   - Contrôler la pression du circuit, rajouter si nécessaire

3. **Contrôle de la vitesse de rotation**
   - Les circulateurs modernes ont 2 ou 3 vitesses (ou vitesse variable)
   - Vérifier le réglage sur le boîtier électrique du circulateur
   - Réduire la vitesse si le bruit persiste et que le chauffage fonctionne correctement
   - Attention : vitesse trop basse = radiateurs insuffisamment chauffés

4. **Test de déblocage du rotor (si grippé)**
   - Couper l'alimentation électrique de la chaudière
   - Dévisser le capot avant du circulateur (vis centrale)
   - Accéder à l'axe du rotor (fente au centre)
   - Tourner manuellement l'axe avec un tournevis plat (doit tourner librement)
   - Si bloqué : forcer légèrement en tournant dans les deux sens
   - Remonter le capot, remettre sous tension, tester

5. **Contrôle de la pression circuit**
   - Une pression trop basse peut provoquer de la cavitation
   - Vérifier la pression : minimum 1 bar à froid, 1,5-2 bars en fonctionnement
   - Augmenter la pression si nécessaire

6. **Vérification des fixations**
   - Contrôler le serrage des écrous de raccordement du circulateur
   - Vérifier que les brides ou raccords ne transmettent pas de vibrations
   - Ajouter des colliers anti-vibratiles si nécessaire

7. **Contrôle de l'état des roulements**
   - Si bruit mécanique persiste après déblocage et purge
   - Tester en faisant tourner le rotor à la main (doit être fluide, sans à-coups)
   - Si grattement ou résistance importante : **remplacement du circulateur**

8. **Remplacement du circulateur**
   - Couper électricité et eau
   - Vidanger le circuit ou isoler le circulateur avec les vannes
   - Démonter les raccords électriques (repérer les fils)
   - Dévisser les écrous de raccordement (prévoir un récipient pour l'eau)
   - Installer le nouveau circulateur avec joints neufs
   - Reconnecter électriquement en respectant le schéma
   - Remplir, purger, tester

**Prévention :**
- Purger le circulateur 1 à 2 fois par an
- Faire fonctionner la pompe hors saison de chauffe (1x par mois) pour éviter grippage
- Contrôler la pression du circuit régulièrement
- Utiliser un inhibiteur de corrosion pour limiter l'encrassement

**Spécificités techniques :**
- Durée de vie moyenne d'un circulateur : 10-15 ans
- Circulateurs classe A (haute efficacité énergétique) : Grundfos Alpha, Wilo Stratos
- Pression minimale amont circulateur : 0,5 bar pour éviter cavitation

---

## FACT-CHAUD-040: Pompe bloquée ou grippée

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-040 |
| **Catégorie** | Hydraulique |
| **Système** | Circulateur / Pompe |
| **Gravité** | **Élevée** |
| **Marques** | Grundfos, Wilo, DAB, Salmson (tous circulateurs) |

**Symptômes :**
- Circulateur ne tourne pas, aucun bruit de fonctionnement
- Le corps du circulateur est chaud (surchauffe du moteur)
- Radiateurs restent froids malgré chaudière en chauffe
- Voyant de défaut ou erreur pompe sur la chaudière
- Disjonction électrique ou fusible grillé

**Cause racine probable :**
Rotor grippé par manque d'utilisation (arrêt prolongé), dépôts calcaires ou oxydes métalliques bloquant l'axe, roulements grippés, corps étranger dans la pompe, bobinage moteur HS.

**Étapes de résolution :**

1. **Vérification de l'alimentation électrique**
   - Contrôler que le circulateur est bien alimenté (multimètre : 230V)
   - Vérifier l'état du fusible ou disjoncteur dédié
   - Tester la continuité des connexions électriques

2. **Diagnostic du blocage mécanique**
   - Couper l'alimentation électrique (sécurité)
   - Poser la main sur le corps du circulateur : s'il est très chaud, le moteur a tenté de tourner sans y parvenir (blocage mécanique)
   - Écouter : aucun bruit = blocage total ou défaut électrique

3. **Déblocage manuel du rotor**
   - Maintenir l'alimentation coupée
   - Dévisser la vis centrale du capot avant du circulateur (vis de 5-8 mm)
   - Retirer le capot avec précaution (attention : un peu d'eau peut s'écouler)
   - Observer l'axe du rotor au centre (fente pour tournevis plat)
   - Insérer un tournevis plat dans la fente de l'axe
   - Tourner manuellement dans les deux sens pour débloquer
   - **Astuce** : si très dur, appliquer quelques gouttes de dégrippant (WD-40) autour de l'axe, attendre 10 min, réessayer

4. **Vérification de la rotation libre**
   - Une fois débloqué, l'axe doit tourner librement à la main (360° sans résistance)
   - Si rotation toujours difficile ou par à-coups : roulements usés ou encrassement important

5. **Nettoyage interne (si encrassement)**
   - Vidanger le circuit ou fermer les vannes d'isolement du circulateur
   - Déposer complètement le circulateur
   - Démonter le corps (selon modèle) pour accéder au rotor et à la volute
   - Nettoyer les dépôts (calcaire, boues, oxydes) avec brosse douce et vinaigre blanc
   - Rincer abondamment à l'eau claire
   - Remonter avec joints neufs si disponibles

6. **Test de remise en service**
   - Remonter le capot du circulateur
   - Remettre sous tension électrique
   - Observer si le circulateur démarre (légère vibration, bruit de fonctionnement)
   - Purger le circulateur pour évacuer l'air introduit
   - Vérifier que les radiateurs chauffent

7. **Si le déblocage échoue : test du bobinage moteur**
   - Couper l'alimentation, débrancher les fils du circulateur
   - Mesurer la résistance du bobinage avec un multimètre (ohmmètre)
   - Valeur normale : 50-500 ohms selon modèle (vérifier notice)
   - Si résistance infinie (∞) : bobinage coupé → **remplacement obligatoire**
   - Si résistance nulle (0) : court-circuit → **remplacement obligatoire**
   - Tester l'isolement moteur/masse : doit être > 1 MΩ

8. **Remplacement du circulateur**
   - Si le déblocage est impossible ou le moteur est HS
   - Suivre la procédure décrite dans FACT-CHAUD-039 étape 8
   - Choisir un circulateur de performances équivalentes (hauteur manométrique, débit)

**Prévention :**
- Faire fonctionner le circulateur au moins 1 fois par mois hors saison de chauffe (évite grippage)
- Utiliser un inhibiteur de corrosion dans le circuit
- Entretenir le circuit (désembouage si nécessaire)
- Installer un filtre à boues en amont du circulateur

**Spécificités techniques :**
- Les circulateurs modernes à aimants permanents (Grundfos Alpha, Wilo Stratos) sont moins sujets au grippage
- Fonction anti-blocage sur certains circulateurs électroniques (rotation automatique périodique)

**Avertissements sécurité :**
- Toujours couper l'alimentation avant d'ouvrir le capot du circulateur (risque électrique)
- Ne jamais faire fonctionner un circulateur bloqué plus de quelques minutes (destruction du moteur par surchauffe)

---

## FACT-CHAUD-041: Circulateur ne démarre pas

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-041 |
| **Catégorie** | Hydraulique |
| **Système** | Circulateur / Commande électrique |
| **Gravité** | **Élevée** |
| **Marques** | Multi-marques (tous systèmes) |

**Symptômes :**
- Chaudière en chauffe mais radiateurs froids
- Aucun bruit ni vibration du circulateur
- Le corps du circulateur reste froid (moteur non alimenté)
- Pas de déclenchement électrique ni de surchauffe

**Cause racine probable :**
Défaut de commande électrique (relais carte, thermostat d'ambiance, aquastat), câblage défectueux, fusible grillé, circulateur HS (bobinage coupé), mauvais paramétrage régulation.

**Étapes de résolution :**

1. **Vérification du mode de fonctionnement**
   - Contrôler que la chaudière est bien en mode chauffage (pas seulement ECS)
   - Vérifier le thermostat d'ambiance : température de consigne > température ambiante
   - Sur les régulations électroniques : vérifier que le mode chauffage est activé
   - Vérifier la programmation horaire (pas en période d'arrêt programmé)

2. **Contrôle de l'alimentation électrique du circulateur**
   - Couper l'alimentation générale de la chaudière
   - Accéder au bornier du circulateur (retirer le capot électrique)
   - Remettre sous tension en mode chauffage
   - Mesurer la tension aux bornes du circulateur avec un multimètre (230V AC attendu)
   - **Si 230V présent** : défaut du circulateur lui-même → passer à l'étape 6
   - **Si 0V** : problème de commande en amont → passer à l'étape 3

3. **Vérification du thermostat d'ambiance**
   - Tester le contact du thermostat (multimètre en position continuité/ohmmètre)
   - En mode chauffage avec consigne > température : contact doit être fermé (0 ohm)
   - Si contact ouvert (∞ ohm) : thermostat défectueux ou mal réglé
   - Test de contournement : ponter les bornes du thermostat (fil volant) pour tester
   - Si le circulateur démarre : thermostat HS → remplacement

4. **Contrôle du relais de commande circulateur (carte électronique)**
   - Localiser le relais de commande circulateur sur la carte (voir schéma électrique)
   - En mode chauffage, le relais doit être enclenché (clic audible)
   - Mesurer la tension de sortie relais (doit être 230V)
   - Si relais n'enclenche pas : défaut de la carte électronique ou de l'aquastat

5. **Vérification de l'aquastat ou sonde de température**
   - L'aquastat autorise le circulateur uniquement si l'eau de chaudière est chaude
   - Vérifier la température de l'eau chaudière (affichage digital)
   - Si température > 40°C et circulateur ne démarre pas : aquastat ou carte défectueux
   - Tester l'aquastat (mesure de résistance, comparaison avec courbe constructeur)

6. **Test du circulateur (bobinage moteur)**
   - Couper l'alimentation, débrancher les fils du circulateur
   - Mesurer la résistance du bobinage avec un ohmmètre
   - Valeur normale : 50-500 ohms selon modèle
   - **Si résistance infinie (∞)** : bobinage coupé → **circulateur HS, remplacement**
   - **Si résistance correcte** : vérifier qu'il n'est pas bloqué mécaniquement (voir FACT-CHAUD-040)

7. **Contrôle du câblage et des connexions**
   - Inspecter visuellement les fils et connexions (oxydation, brûlures, déconnexions)
   - Vérifier la continuité des fils entre carte et circulateur
   - Resserrer les connexions desserrées
   - Remplacer les cosses ou dominos défectueux

8. **Test avec circulateur de secours (diagnostic)**
   - Si disponible, brancher temporairement un circulateur fonctionnel connu
   - Si ce circulateur démarre : confirme que le circulateur d'origine est HS
   - Si ce circulateur ne démarre pas non plus : problème de commande (carte, thermostat)

9. **Remplacement selon diagnostic**
   - **Circulateur HS** : remplacement (voir FACT-CHAUD-039 étape 8)
   - **Thermostat HS** : remplacement du thermostat d'ambiance
   - **Carte électronique HS** : remplacement de la carte (pièce coûteuse, vérifier garantie)

**Prévention :**
- Tester le bon fonctionnement du chauffage en début de saison
- Vérifier les connexions électriques lors de l'entretien annuel
- Protéger la carte électronique de l'humidité

**Spécificités techniques :**
- Certains circulateurs ont une protection thermique interne qui peut se déclencher
- Les circulateurs à commande PWM (modulation de largeur d'impulsion) nécessitent un signal spécifique

**Avertissements sécurité :**
- Toujours couper l'alimentation avant de toucher les connexions électriques
- Ne jamais ponter de sécurité de manière permanente (test uniquement)

---

## FACT-CHAUD-042: Débit insuffisant, radiateurs froids

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-042 |
| **Catégorie** | Hydraulique |
| **Système** | Circuit chauffage / Distribution |
| **Gravité** | **Moyenne** |
| **Marques** | Multi-marques (tous systèmes) |

**Symptômes :**
- Radiateurs tièdes ou froids, surtout les plus éloignés de la chaudière
- Chauffage lent à monter en température
- Température départ chaudière correcte mais retour froid
- Certains radiateurs chauffent bien, d'autres non (déséquilibre)

**Cause racine probable :**
Circulateur sous-dimensionné ou vitesse trop faible, embouage du circuit, filtre encrassé, vannes radiateurs fermées ou grippées, air dans le circuit, by-pass hydraulique mal réglé, pertes de charge excessives.

**Étapes de résolution :**

1. **Vérification de la température de départ/retour**
   - Relever la température de départ chaudière (affichage ou thermomètre)
   - Mesurer la température de retour (tuyau retour avant chaudière)
   - Écart normal : 15-20°C (par ex : départ 70°C, retour 50-55°C)
   - Si écart > 25°C : débit trop faible, mauvaise circulation

2. **Contrôle de la vitesse du circulateur**
   - Vérifier le réglage de vitesse sur le circulateur (sélecteur 1-2-3 ou Auto)
   - Passer à la vitesse supérieure si débit insuffisant
   - Observer si l'amélioration est sensible (radiateurs plus chauds)
   - Attention : augmenter la vitesse augmente aussi la consommation électrique et le bruit

3. **Purge complète du circuit**
   - Purger méthodiquement tous les radiateurs en commençant par les plus proches
   - Maintenir la pression du circuit pendant la purge (rajouter eau si nécessaire)
   - Purger également le circulateur (voir FACT-CHAUD-039)
   - Purger les points hauts du circuit (purgeurs automatiques s'ils existent)

4. **Vérification de l'ouverture des vannes radiateurs**
   - Contrôler que tous les robinets thermostatiques ou manuels sont bien ouverts
   - Démonter la tête thermostatique : vérifier que le pointeau n'est pas grippé en position fermée
   - Actionner manuellement le pointeau (appuyer dessus, il doit revenir par ressort)
   - Si grippé : débloquer avec pince multiprise ou WD-40, actionner plusieurs fois

5. **Contrôle du filtre à boues**
   - Localiser le filtre (généralement en amont de la chaudière ou du circulateur)
   - Fermer les vannes d'isolement du filtre
   - Dévisser le corps du filtre, retirer la cartouche
   - Nettoyer la cartouche sous l'eau (brosse si très encrassée)
   - Remonter, ouvrir les vannes, purger l'air, tester

6. **Vérification de l'embouage du circuit**
   - Observer la couleur de l'eau lors d'une purge : doit être claire
   - Si eau noire, marron ou rouge : circuit embué (boues magnétite)
   - Test : placer un aimant sur un radiateur froid, s'il colle fortement = boues magnétiques
   - **Solution** : désembouage complet du circuit (voir FACT-CHAUD-043)

7. **Contrôle du by-pass hydraulique**
   - Certains circuits ont un by-pass entre départ et retour (vanne réglable ou automatique)
   - Si by-pass trop ouvert : une partie de l'eau chauffe contourne les radiateurs
   - Fermer partiellement le by-pass et observer l'amélioration
   - Réglage optimal : débit juste suffisant pour éviter le blocage du circulateur vannes fermées

8. **Équilibrage hydraulique des radiateurs**
   - Les radiateurs éloignés peuvent manquer de débit si pas d'équilibrage
   - Technique : réduire légèrement le débit des radiateurs proches (vis de réglage retour)
   - Augmenter le débit des radiateurs éloignés (ouvrir complètement)
   - Procéder par itérations jusqu'à obtenir une température homogène

9. **Vérification du dimensionnement du circulateur**
   - Calculer la hauteur manométrique nécessaire (pertes de charge circuit)
   - Comparer avec les performances du circulateur installé (voir courbe constructeur)
   - Si circulateur sous-dimensionné : **remplacement par modèle plus puissant**

10. **Contrôle des tuyauteries (obstruction)**
    - Si un seul radiateur reste froid : obstruction locale probable
    - Démonter le radiateur, rincer à contre-courant au tuyau d'arrosage
    - Vérifier l'absence d'obstruction sur les vannes et té de dérivation

**Prévention :**
- Installer un filtre magnétique pour capturer les boues
- Ajouter un inhibiteur de corrosion dans le circuit
- Effectuer un désembouage préventif tous les 5-7 ans
- Équilibrer le circuit lors de l'installation ou modification

**Spécificités techniques :**
- Formule de débit : Q (m³/h) = P (kW) / (1,16 × ΔT) où ΔT est l'écart départ/retour
- Exemple : 15 kW avec ΔT 20°C → Q = 15/(1,16×20) = 0,65 m³/h = 650 L/h

---

## FACT-CHAUD-043: Embouage circuit, boues

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-043 |
| **Catégorie** | Hydraulique |
| **Système** | Circuit chauffage / Qualité eau |
| **Gravité** | **Moyenne** |
| **Marques** | Multi-marques (tous circuits acier/fonte) |

**Symptômes :**
- Eau du circuit noire, marron ou rougeâtre
- Radiateurs froids en partie basse (boues déposées)
- Bruit de circulation d'eau (gargouillis, clapotis)
- Circulateur bruyant ou grippé
- Dégradation progressive des performances de chauffage

**Cause racine probable :**
Corrosion interne des éléments ferreux (radiateurs acier/fonte, tuyaux acier), formation d'oxydes de fer (magnétite Fe₃O₄), apport d'oxygène par remplissages fréquents, absence d'inhibiteur de corrosion, eau dure (calcaire favorise corrosion).

**Étapes de résolution :**

1. **Diagnostic de l'embouage**
   - Purger un radiateur et observer la couleur de l'eau
   - Eau noire/marron foncé : embouage important (magnétite)
   - Eau rougeâtre : corrosion active (hématite Fe₂O₃)
   - Test à l'aimant sur radiateur : si forte attraction = boues magnétiques
   - Mesurer le pH de l'eau : pH optimal 7-8,5, si < 7 = corrosion active

2. **Désembouage chimique (circuit peu encrassé)**
   - Couper la chaudière, laisser refroidir
   - Ajouter un produit désembouant dans le circuit via un radiateur
   - Produits recommandés : Sentinel X400, Fernox F3, Adey MC3+
   - Faire circuler le produit pendant 1 à 4 semaines selon encrassement
   - Pendant cette période, faire fonctionner le chauffage normalement
   - Le produit dissout et met en suspension les boues

3. **Désembouage par rinçage (après traitement chimique)**
   - Vidanger complètement le circuit (point bas)
   - Rincer à l'eau claire en remplissant/vidangeant plusieurs fois
   - Continuer jusqu'à ce que l'eau de vidange soit claire
   - Purger chaque radiateur pendant le rinçage pour évacuer les dépôts

4. **Désembouage hydrodynamique (embouage important)**
   - Méthode professionnelle avec pompe de désembouage
   - Isoler la chaudière (vannes)
   - Connecter une pompe de désembouage sur le circuit
   - Faire circuler de l'eau claire à haute vitesse avec inversion de flux
   - Ajouter un produit désembouant à la pompe
   - Utiliser des "coups de bélier" pour décoller les boues (chocs hydrauliques)
   - Rincer jusqu'à eau claire (peut prendre 2-4 heures)

5. **Désembouage radiateur par radiateur**
   - Isoler chaque radiateur (fermer vannes)
   - Démonter le radiateur, le sortir
   - Rincer à l'extérieur avec tuyau d'arrosage (pression réseau)
   - Secouer le radiateur pour décoller les boues
   - Rincer à contre-courant (inversion entrée/sortie)
   - Remonter le radiateur avec joints neufs

6. **Installation d'un filtre magnétique**
   - Installer un séparateur de boues magnétique sur le retour chauffage (amont chaudière)
   - Modèles recommandés : Fernox TF1, Sentinel Eliminator, Adey MagnaClean
   - Le filtre capte en continu les particules métalliques
   - Nettoyage du filtre : tous les 6-12 mois selon encrassement

7. **Traitement inhibiteur de corrosion**
   - Après désembouage complet, ajouter un inhibiteur dans le circuit
   - Produits recommandés : Sentinel X100, Fernox F1, Adey MC1+
   - Dosage selon volume du circuit (lire notice fabricant)
   - L'inhibiteur forme un film protecteur anti-corrosion sur les surfaces métalliques
   - Renouveler tous les 5 ans ou après vidange

8. **Contrôle et ajustement du pH**
   - Vérifier le pH de l'eau avec bandelettes ou pH-mètre
   - pH optimal : 7,5-8,5
   - Si pH bas (<7) : ajouter un correcteur de pH
   - Si pH élevé (>9) : rincer et remplir avec eau neuve

9. **Traitement préventif anticalcaire (eau dure)**
   - Si eau très calcaire (TH > 25°F), installer un adoucisseur ou filtre polyphosphates
   - Remplir le circuit avec eau adoucie
   - Ajouter un séquestrant calcaire dans le circuit

10. **Contrôle après traitement**
    - Remplir le circuit avec eau traitée
    - Purger soigneusement tout l'air
    - Mettre en pression (1,2-1,5 bar)
    - Faire fonctionner le chauffage sur plusieurs cycles
    - Vérifier que tous les radiateurs chauffent uniformément
    - Contrôler à nouveau la couleur de l'eau après 1 semaine

**Prévention :**
- Installer un filtre magnétique dès la mise en service
- Ajouter un inhibiteur de corrosion systématiquement
- Éviter les vidanges fréquentes (apport d'oxygène)
- Contrôler le pH annuellement
- Vérifier l'étanchéité du circuit (éviter remplissages répétés)

**Spécificités techniques :**
- Magnétite (Fe₃O₄) : boues noires magnétiques, principale cause d'embouage
- Hématite (Fe₂O₃) : boues rouges non magnétiques, corrosion active
- Volume d'eau circuit : radiateurs ≈ 10L/kW, plancher chauffant ≈ 15L/kW

**Avertissements sécurité :**
- Les produits désembouants sont corrosifs : porter gants et lunettes
- Ne jamais laisser un produit désembouant dans le circuit au-delà du temps recommandé (risque de corrosion)
- Bien rincer avant de remplir définitivement

---

## FACT-CHAUD-044: Air dans le circuit

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-044 |
| **Catégorie** | Hydraulique |
| **Système** | Circuit chauffage / Purge |
| **Gravité** | **Faible à Moyenne** |
| **Marques** | Multi-marques (tous systèmes) |

**Symptômes :**
- Radiateurs froids en partie haute
- Bruits de gargouillis, clapotis dans les tuyaux ou radiateurs
- Circulateur bruyant (bruit de cavitation)
- Pression circuit instable
- Chauffage irrégulier selon les radiateurs

**Cause racine probable :**
Air introduit lors d'un remplissage ou d'une réparation, défaut d'étanchéité aspirant de l'air (dépression), manque de purgeurs automatiques, dégazage naturel de l'eau (microbulles), vase d'expansion mal dimensionné créant des dépressions.

**Étapes de résolution :**

1. **Purge méthodique des radiateurs**
   - Commencer par les radiateurs les plus proches de la chaudière
   - Procéder vers les radiateurs les plus éloignés
   - Mettre le chauffage en fonctionnement (circulateur actif, eau chaude)
   - Sur chaque radiateur :
     - Préparer un récipient et un chiffon
     - Ouvrir lentement le purgeur avec une clé de purge (1/4 de tour)
     - Laisser l'air s'échapper (sifflement)
     - Attendre que l'eau coule en continu sans bulles
     - Fermer le purgeur fermement
   - Contrôler et ajuster la pression du circuit entre chaque purge

2. **Purge du circulateur**
   - Localiser la vis de purge sur le corps du circulateur
   - Mettre le circulateur en fonctionnement
   - Dévisser lentement la vis de purge (1/4 de tour, tournevis plat)
   - Laisser l'eau couler jusqu'à écoulement continu sans bulles
   - Revisser fermement

3. **Purge des points hauts du circuit**
   - Identifier les points hauts du réseau (colonnes montantes, tuyaux sous combles)
   - Installer des purgeurs manuels ou automatiques si absents
   - Purger ces points régulièrement

4. **Installation de purgeurs automatiques**
   - Installer des purgeurs automatiques (type Spirotop, Flamco) aux points hauts
   - Idéalement : un purgeur par zone ou étage
   - Les purgeurs automatiques évacuent l'air en continu sans intervention
   - Vérifier leur bon fonctionnement (dévisser le capot, observer si de l'air sort)

5. **Vérification de l'étanchéité du circuit (prise d'air)**
   - Une dépression dans le circuit peut aspirer de l'air
   - Contrôler les presse-étoupes des vannes et des circulateurs
   - Vérifier les raccords au niveau des collecteurs et des départs
   - Resserrer ou remplacer les joints défectueux
   - Tester avec circuit en pression (1,5-2 bars) et à chaud

6. **Contrôle du vase d'expansion**
   - Un vase mal dimensionné ou dégonflé peut créer des variations de pression et favoriser la formation de bulles
   - Vérifier la pression du vase (voir FACT-CHAUD-037)
   - S'assurer que le vase est correctement gonflé (0,8-1 bar)

7. **Dégazage complet du circuit (traitement chimique)**
   - Si le problème persiste malgré les purges répétées
   - Ajouter un produit dégazant dans le circuit (ex : Fernox F4, Sentinel X100)
   - Ces produits réduisent la tension superficielle de l'eau et facilitent l'évacuation des microbulles
   - Faire fonctionner le chauffage plusieurs jours, purger régulièrement

8. **Purge par surpression**
   - Augmenter temporairement la pression du circuit à 2-2,5 bars
   - Mettre le chauffage en fonctionnement à haute température
   - La surpression et la température aident à dissoudre et évacuer l'air
   - Purger ensuite tous les radiateurs
   - Ramener la pression à 1,2-1,5 bar

9. **Vérification du remplissage**
   - Toujours remplir le circuit lentement (robinet de remplissage peu ouvert)
   - Un remplissage trop rapide entraîne beaucoup d'air
   - Remplir en plusieurs étapes avec purges intermédiaires

10. **Contrôle après purge**
    - Vérifier que tous les radiateurs chauffent uniformément
    - Écouter : plus de bruit de gargouillis
    - Contrôler la pression : doit rester stable
    - Surveiller sur plusieurs jours (l'air peut réapparaître si prise d'air)

**Prévention :**
- Purger les radiateurs en début de saison de chauffe
- Installer des purgeurs automatiques aux points hauts
- Remplir le circuit lentement
- Utiliser un produit inhibiteur/dégazant dans le circuit
- Vérifier l'étanchéité du circuit régulièrement

**Spécificités techniques :**
- L'eau contient naturellement de l'air dissous (environ 2% en volume)
- Lors de la montée en température, cet air se libère sous forme de bulles
- Les purgeurs automatiques à flotteur évacuent l'air dès qu'il s'accumule

**Avertissements sécurité :**
- Ne jamais ouvrir un purgeur avec le circuit à haute pression (risque de projection d'eau chaude)
- Toujours utiliser un chiffon pour protéger la main lors de la purge

---

## FACT-CHAUD-045: Soupape sécurité 3 bars qui coule

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-045 |
| **Catégorie** | Hydraulique |
| **Système** | Soupape de sécurité / Pressurisation |
| **Gravité** | **Moyenne** |
| **Marques** | Multi-marques (tous types de chaudières) |

**Symptômes :**
- Eau qui s'écoule par la soupape de sécurité et son tuyau d'évacuation
- Flaques d'eau au sol près de l'évacuation
- Écoulement intermittent ou permanent
- Chute de pression du circuit après écoulement

**Cause racine probable :**
Vase d'expansion dégonflé ou HS (surpression du circuit), pression de remplissage trop élevée, soupape entartrée ou défectueuse (ne ferme plus correctement), corps étranger bloquant le clapet, température excessive de l'eau.

**Étapes de résolution :**

1. **Vérification de la pression du circuit**
   - Observer le manomètre de la chaudière
   - Pression normale à froid : 1,2-1,5 bar
   - Pression normale en chauffe : 1,8-2,2 bar
   - **Si pression > 2,5 bars** : surpression anormale → diagnostic vase d'expansion

2. **Diagnostic du vase d'expansion**
   - La cause la plus fréquente d'écoulement de la soupape est un vase HS
   - Suivre la procédure complète décrite dans FACT-CHAUD-037
   - Vérifier la pression azote du vase (doit être 0,8-1 bar)
   - Tester l'étanchéité de la membrane (eau qui sort de la valve = vase HS)

3. **Contrôle de la pression de remplissage**
   - Si la pression à froid dépasse 1,8 bar, elle atteindra > 3 bars en chauffe
   - Vidanger légèrement le circuit pour ramener à 1,2-1,5 bar à froid
   - Utiliser un purgeur de radiateur ou la vidange chaudière

4. **Test de la soupape elle-même**
   - Couper la chaudière, laisser refroidir
   - Actionner manuellement le levier de la soupape (si présent)
   - L'eau doit s'écouler puis s'arrêter complètement à la fermeture
   - Si la soupape continue de couler après test manuel : encrassement ou usure du clapet

5. **Nettoyage/détartrage de la soupape (si entartrée)**
   - Couper la chaudière et l'alimentation électrique
   - Fermer les vannes d'isolement du circuit chauffage
   - Vidanger partiellement le circuit (pression à 0)
   - Dévisser la soupape de sécurité
   - Inspecter le clapet et le siège : présence de calcaire, corps étranger ?
   - Nettoyer avec vinaigre blanc (trempage 1h) ou produit détartrant
   - Rincer abondamment à l'eau claire
   - Remonter la soupape avec téflon ou filasse (attention au sens de montage)

6. **Remplacement de la soupape (si nettoyage inefficace)**
   - Utiliser une soupape de même calibrage : 3 bars (ou selon notice constructeur)
   - Vérifier le filetage (1/2", 3/4", M20, etc.)
   - Monter avec étanchéité (téflon ou filasse + pâte)
   - Respecter le sens de montage (flèche sur le corps)
   - Raccorder le tuyau d'évacuation (évacuation libre visible, pas d'obstruction)

7. **Contrôle après intervention**
   - Remplir le circuit à la pression correcte (1,2-1,5 bar à froid)
   - Mettre en chauffe et observer la montée en pression
   - Vérifier que la pression ne dépasse pas 2,5 bars en pleine chauffe
   - Contrôler qu'aucun écoulement ne se produit
   - Surveiller sur plusieurs cycles de chauffe

8. **Si le problème persiste malgré vase et soupape OK**
   - Vérifier qu'il n'y a pas de sur-chauffe de l'eau (température excessive)
   - Contrôler le bon fonctionnement du circulateur (débit correct)
   - Vérifier que le robinet de remplissage est bien fermé (pas de fuite interne)
   - Tester la régulation de température (aquastat, sonde)

**Prévention :**
- Contrôler la pression du circuit régulièrement (mensuellement)
- Vérifier l'état du vase d'expansion annuellement
- Tester la soupape manuellement 1 fois par an (actionnement du levier)
- Utiliser de l'eau adoucie en cas d'eau très calcaire

**Spécificités techniques :**
- Pression de tarage standard : 3 bars (certaines installations : 2,5 ou 4 bars)
- Débit nominal de la soupape doit être adapté à la puissance de la chaudière
- Norme NF EN 12828 : obligation d'une soupape de sécurité sur tout circuit fermé

**Avertissements sécurité :**
- La soupape de sécurité est un organe de sécurité critique : ne jamais l'obturer ou la condamner
- Le tuyau d'évacuation doit toujours être visible et se terminer librement (pas de bouchon, pas de raccordement étanche)
- Ne jamais augmenter le tarage de la soupape (risque d'explosion du circuit)
- Attention aux brûlures lors de l'intervention (eau à 80-90°C possible)

---

## FACT-CHAUD-046: Manomètre défectueux

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-046 |
| **Catégorie** | Hydraulique |
| **Système** | Instrumentation / Manomètre |
| **Gravité** | **Faible** |
| **Marques** | Multi-marques (tous types) |

**Symptômes :**
- Aiguille du manomètre bloquée ou immobile
- Indication de pression aberrante (toujours à 0 ou toujours à 3 bars)
- Aiguille qui tremble ou oscille fortement
- Affichage incohérent avec l'état du circuit (par ex. pression affichée 0 mais chaudière fonctionne)

**Cause racine probable :**
Manomètre usé ou oxydé (mécanisme interne grippé), prise de pression bouchée (calcaire, impuretés), fuite au niveau du raccord du manomètre, membrane ou tube de Bourdon défectueux.

**Étapes de résolution :**

1. **Vérification de l'indication affichée**
   - Comparer l'affichage du manomètre avec le comportement du circuit
   - Si la chaudière fonctionne normalement (radiateurs chauds) mais manomètre à 0 : manomètre ou prise défectueux
   - Si soupape coule mais manomètre à 1 bar : manomètre défectueux (sous-estime)

2. **Test de réactivité du manomètre**
   - Purger légèrement un radiateur et observer si l'aiguille descend
   - Rajouter un peu d'eau et observer si l'aiguille monte
   - Si aucune réaction : manomètre HS ou prise bouchée

3. **Contrôle de la prise de pression**
   - Localiser le raccord du manomètre (souvent sous la chaudière, vissé sur le corps)
   - Couper la chaudière, fermer les vannes d'isolement, vidanger légèrement (pression 0)
   - Dévisser le manomètre avec précaution (prévoir un chiffon pour les gouttes)
   - Inspecter la prise : présence de calcaire, de boues, d'obstruction ?
   - Nettoyer la prise avec une petite brosse ou fil de fer fin
   - Rincer à l'eau claire

4. **Test du manomètre démonté**
   - Si possible, tester le manomètre sur un autre circuit ou avec un compresseur
   - Appliquer une pression connue et vérifier l'affichage
   - Si l'aiguille reste bloquée : manomètre HS → remplacement

5. **Remplacement du manomètre**
   - Choisir un manomètre de même type et plage (0-4 bars ou 0-6 bars selon installation)
   - Vérifier le filetage (généralement 1/4", 1/2" ou M10)
   - Orientation du cadran : manomètre axial (connexion arrière) ou radial (connexion dessous)
   - Monter avec téflon ou joint fibre selon type de raccord
   - Serrer à la main puis 1/4 de tour à la clé (ne pas forcer, risque de casser)

6. **Contrôle après remplacement**
   - Remplir le circuit à pression normale (1,2-1,5 bar)
   - Vérifier que l'aiguille indique la pression correctement
   - Tester la réactivité : purger légèrement, observer la descente de l'aiguille
   - Mettre en chauffe et vérifier la montée progressive de la pression

7. **Vérification des fuites au raccord**
   - Inspecter le raccord du manomètre après remise en pression
   - Si goutte à goutte : resserrer légèrement ou refaire l'étanchéité
   - Ne pas trop serrer (risque de fissure du cadran en verre)

**Prévention :**
- Vérifier visuellement le manomètre régulièrement (lors de l'entretien)
- Taper légèrement sur le cadran si aiguille semble bloquée (peut débloquer temporairement)
- Utiliser un manomètre avec glycérine (amortit les vibrations, prolonge la durée de vie)

**Spécificités techniques :**
- Manomètre à tube de Bourdon : tube métallique en arc de cercle qui se déforme sous pression
- Classe de précision : classe 2,5 ou 4 pour manomètres chaudières (acceptable)
- Diamètre cadran courant : 40 mm, 50 mm ou 63 mm

**Avertissements sécurité :**
- Un manomètre défectueux empêche la surveillance de la pression (risque de surpression non détectée)
- Toujours remplacer rapidement un manomètre HS pour éviter tout risque
- Vidanger le circuit avant de démonter le manomètre (risque de projection d'eau chaude sous pression)

---

## FACT-CHAUD-047: By-pass hydraulique mal réglé

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-047 |
| **Catégorie** | Hydraulique |
| **Système** | By-pass / Équilibrage |
| **Gravité** | **Moyenne** |
| **Marques** | Multi-marques (circuits avec robinets thermostatiques) |

**Symptômes :**
- Radiateurs chauffent mal malgré circulateur en fonctionnement
- Débit excessif dans certains radiateurs, insuffisant dans d'autres
- Bruit de circulation important (sifflement, vibration)
- Écart de température départ/retour faible (< 10°C)
- Circulateur bruyant ou en protection (débit trop faible)

**Cause racine probable :**
By-pass trop ouvert (eau court-circuite les radiateurs), by-pass trop fermé (surpression, bruit), absence de by-pass sur circuit avec robinets thermostatiques (risque blocage circulateur vannes fermées).

**Étapes de résolution :**

1. **Identification du by-pass**
   - Le by-pass est une tuyauterie de liaison entre le départ et le retour chauffage
   - Généralement équipé d'une vanne réglable (pointeau, papillon, ou by-pass automatique)
   - Localiser cette vanne près de la chaudière ou du collecteur
   - Sur certaines installations : by-pass intégré à la chaudière (vanne interne)

2. **Fonction du by-pass**
   - Permet un débit minimal dans le circulateur lorsque tous les robinets thermostatiques sont fermés
   - Évite la surpression et le blocage du circulateur
   - Limite le ΔT excessif (écart température départ/retour)
   - Régule la vitesse de l'eau dans le circuit

3. **Diagnostic du problème**
   - **By-pass trop ouvert** :
     - Écart départ/retour faible (< 10°C)
     - Radiateurs tièdes, surtout les plus éloignés
     - Eau chaude revient rapidement à la chaudière sans céder sa chaleur
   - **By-pass trop fermé ou absent** :
     - Bruit de sifflement dans les tuyaux
     - Circulateur bruyant ou vibrations
     - Risque de blocage du circulateur si tous les robinets thermostatiques se ferment

4. **Réglage du by-pass (vanne manuelle)**
   - Ouvrir tous les robinets thermostatiques (position maximale)
   - Fermer complètement la vanne de by-pass
   - Mettre la chaudière en chauffe
   - Mesurer la température départ et retour (thermomètre infrarouge ou sonde)
   - Ouvrir progressivement le by-pass jusqu'à obtenir un écart de 15-20°C
   - Tester en fermant progressivement les robinets thermostatiques
   - Le by-pass doit empêcher le bruit et la surpression lorsque les robinets se ferment

5. **Réglage optimal du ΔT (écart départ/retour)**
   - ΔT idéal : 15-20°C (par ex. départ 70°C, retour 50-55°C)
   - Si ΔT < 10°C : by-pass trop ouvert → fermer partiellement
   - Si ΔT > 25°C : by-pass trop fermé ou débit insuffisant → ouvrir légèrement

6. **Installation d'un by-pass (si absent)**
   - Sur un circuit équipé de robinets thermostatiques, le by-pass est **obligatoire**
   - Créer une liaison entre départ et retour au plus près de la chaudière
   - Installer une vanne de réglage (pointeau de 10-12 mm)
   - Positionner le by-pass après le circulateur (sur le départ) vers le retour

7. **By-pass automatique (vanne différentielle)**
   - Alternative au by-pass manuel : vanne by-pass automatique (régulateur de pression différentielle)
   - S'ouvre automatiquement lorsque la pression différentielle augmente (robinets fermés)
   - Réglage : tarage à 0,2-0,4 bar selon installation (voir notice)
   - Avantage : adaptation automatique au besoin, pas de réglage manuel

8. **Contrôle après réglage**
   - Vérifier le fonctionnement sur plusieurs cycles de chauffe
   - Tester avec tous les robinets ouverts, puis fermés progressivement
   - Mesurer le ΔT départ/retour (doit rester dans la plage 15-20°C)
   - Écouter : absence de sifflement, de vibration
   - Vérifier que tous les radiateurs chauffent correctement

9. **Réglage combiné : by-pass + équilibrage radiateurs**
   - Réduire le débit des radiateurs proches (vis de réglage au retour)
   - Augmenter le débit des radiateurs éloignés
   - Procéder par itérations avec le réglage du by-pass
   - Objectif : température homogène sur tous les radiateurs

**Prévention :**
- Installer systématiquement un by-pass sur circuits avec robinets thermostatiques
- Vérifier le réglage lors de l'entretien annuel
- Adapter le réglage après modification du circuit (ajout/retrait de radiateurs)

**Spécificités techniques :**
- Débit minimal recommandé dans le circulateur : 20-30% du débit nominal
- Diamètre courant du by-pass : 10-16 mm selon puissance
- By-pass automatique : Flamco Flexvent, TA Hydronics STAP, Oventrop Aquastrom

**Avertissements sécurité :**
- Ne jamais supprimer le by-pass sur circuit avec robinets thermostatiques (risque de destruction du circulateur)
- Un by-pass trop fermé peut entraîner une surpression et le déclenchement de la soupape de sécurité

---

## FACT-CHAUD-048: Filtre encrassé

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-048 |
| **Catégorie** | Hydraulique |
| **Système** | Filtre / Séparateur de boues |
| **Gravité** | **Moyenne** |
| **Marques** | Multi-marques (tous circuits) |

**Symptômes :**
- Débit réduit dans le circuit (radiateurs tièdes ou froids)
- Circulateur bruyant (effort excessif)
- Écart important départ/retour (> 25°C)
- Pression différentielle élevée aux bornes du filtre
- Chaudière en surchauffe ou arrêts fréquents

**Cause racine probable :**
Accumulation de boues, impuretés, magnétite dans le filtre, absence d'entretien régulier du filtre, embouage important du circuit.

**Étapes de résolution :**

1. **Localisation du filtre**
   - Le filtre est généralement positionné sur le retour chauffage, avant la chaudière
   - Types de filtres courants :
     - Filtre à tamis (cartouche maille fine)
     - Séparateur magnétique (Fernox TF1, Adey MagnaClean, Sentinel Eliminator)
     - Pot à boues (décantation par gravité)

2. **Diagnostic de l'encrassement**
   - Vérifier la couleur de l'eau du circuit (purge radiateur)
   - Si eau noire/marron : circuit embué, filtre probablement saturé
   - Sur filtres magnétiques : observer la quantité de boues collectées (fenêtre de visualisation si présente)
   - Mesurer la pression différentielle (manomètres amont/aval du filtre) : si ΔP > 0,2 bar → encrassement

3. **Nettoyage d'un filtre magnétique (Fernox, Adey, Sentinel)**
   - Fermer les vannes d'isolement du filtre (amont et aval)
   - Placer un récipient sous le filtre
   - Dévisser lentement le corps du filtre (sens anti-horaire)
   - Attention : de l'eau résiduelle va s'écouler
   - Retirer l'aimant central avec sa coque de boues
   - Nettoyer l'aimant sous l'eau (brosse, jet) : éliminer toutes les boues magnétiques
   - Nettoyer le corps du filtre (rincer, brosser)
   - Remonter l'aimant dans le corps, revisser sur la tête du filtre
   - Ouvrir lentement les vannes (purger l'air si nécessaire)

4. **Nettoyage d'un filtre à tamis**
   - Fermer les vannes d'isolement
   - Vidanger le filtre (vis de purge ou dévissage du corps)
   - Dévisser le bouchon de visite ou le corps du filtre
   - Retirer la cartouche filtrante (tamis)
   - Nettoyer sous l'eau ou avec une brosse (ne pas déchirer le tamis)
   - Si tamis percé ou déformé : remplacer par un neuf
   - Remonter la cartouche, revisser le corps avec joint neuf
   - Remplir, purger, ouvrir les vannes

5. **Nettoyage d'un pot à boues**
   - Fermer les vannes d'isolement
   - Placer un récipient sous le pot
   - Dévisser le bouchon inférieur du pot (sens anti-horaire)
   - Laisser s'écouler l'eau et les boues accumulées
   - Rincer l'intérieur du pot (introduire de l'eau par le haut si possible)
   - Revisser le bouchon avec joint neuf ou téflon
   - Ouvrir les vannes

6. **Contrôle après nettoyage**
   - Vérifier l'absence de fuite au niveau des joints du filtre
   - Contrôler la pression du circuit (compléter si nécessaire)
   - Mettre en chauffe et observer l'amélioration du débit
   - Vérifier que l'écart départ/retour revient à 15-20°C
   - Surveiller les radiateurs : doivent chauffer plus rapidement et uniformément

7. **Installation d'un filtre (si absent)**
   - Sur tout circuit de chauffage, l'installation d'un filtre est recommandée
   - Position idéale : retour chauffage, avant la chaudière et le circulateur
   - Préférer un filtre magnétique (plus efficace contre magnétite)
   - Prévoir des vannes d'isolement pour faciliter l'entretien

8. **Fréquence de nettoyage**
   - Vérifier l'état du filtre tous les 6 mois la première année
   - Ensuite, nettoyer 1 fois par an (lors de l'entretien chaudière) si circuit sain
   - Sur circuit embué : nettoyer tous les 3 mois jusqu'à amélioration

9. **Traitement de l'embouage du circuit**
   - Si le filtre s'encrasse très rapidement (< 3 mois)
   - Procéder à un désembouage complet du circuit (voir FACT-CHAUD-043)
   - Ajouter un inhibiteur de corrosion après désembouage
   - Le filtre captera alors les résidus résiduels

**Prévention :**
- Installer un filtre magnétique dès la mise en service du circuit
- Nettoyer régulièrement le filtre (1x/an minimum)
- Traiter le circuit avec inhibiteur de corrosion
- Effectuer un désembouage préventif tous les 5-7 ans

**Spécificités techniques :**
- Filtres magnétiques : captent les particules de 5 microns (magnétite Fe₃O₄)
- Efficacité : 95-98% de capture des boues magnétiques
- Diamètre de raccordement courant : 22 mm, 28 mm (cuivre) ou 3/4", 1" (fileté)

**Avertissements sécurité :**
- Toujours fermer les vannes d'isolement avant de dévisser le filtre
- Les boues peuvent contenir des bactéries : porter des gants
- Attention au poids du filtre chargé de boues (peut être lourd)

---

## FACT-CHAUD-049: Échangeur à plaques entartré (mixte)

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-049 |
| **Catégorie** | Hydraulique |
| **Système** | Échangeur à plaques / Production ECS |
| **Gravité** | **Moyenne à Élevée** |
| **Marques** | Chaudières mixtes (Saunier Duval, Vaillant, Frisquet, De Dietrich, Viessmann) |

**Symptômes :**
- Débit ECS réduit (eau chaude coule faiblement au robinet)
- Température ECS instable ou insuffisante
- Chaudière en surchauffe en mode ECS (arrêts fréquents)
- Bruit de bouillonnement dans l'échangeur
- Fuite d'eau au niveau de l'échangeur à plaques

**Cause racine probable :**
Entartrage de l'échangeur à plaques (dépôts de calcaire réduisant la section de passage), eau dure (TH élevé), température ECS trop élevée (> 60°C favorise le calcaire), manque d'entretien préventif.

**Étapes de résolution :**

1. **Diagnostic de l'entartrage**
   - Mesurer le débit ECS au robinet (chronomètre + récipient gradué)
   - Comparer avec le débit nominal (généralement 10-15 L/min selon chaudière)
   - Si débit < 6-8 L/min : entartrage probable
   - Vérifier la pression eau froide : doit être > 2 bars (si faible, pas forcément l'échangeur)
   - Contrôler la dureté de l'eau (bandelette TH) : si > 25°F, entartrage rapide

2. **Accès à l'échangeur à plaques**
   - Couper l'alimentation électrique et fermer l'arrivée gaz
   - Fermer les vannes d'isolement circuit chauffage et eau froide sanitaire
   - Vidanger partiellement la chaudière (circuit primaire et ECS)
   - Ouvrir le capot de la chaudière
   - Localiser l'échangeur à plaques (généralement au centre, forme rectangulaire avec plaques empilées)
   - Débrancher les 4 raccords hydrauliques (2 chauffage primaire + 2 ECS) en repérant leur position

3. **Démontage de l'échangeur**
   - Dévisser les vis de fixation de l'échangeur (4 à 6 vis selon modèle)
   - Extraire délicatement l'échangeur (attention au poids et à la fragilité des plaques)
   - Placer l'échangeur dans une bassine

4. **Détartrage chimique par trempage**
   - Préparer une solution détartrante :
     - Acide citrique : 200g pour 2L d'eau chaude (solution douce, sans danger pour inox)
     - Vinaigre blanc à 14° : trempage prolongé (12-24h)
     - Produit détartrant commercial : Fernox, Sentinel, Cillit (suivre dosage fabricant)
   - Placer l'échangeur dans la solution
   - Laisser tremper 4 à 12 heures selon l'encrassement
   - Observer la dissolution du calcaire (effervescence)
   - Brosser légèrement entre les plaques avec une brosse douce si accessible

5. **Détartrage par circulation (plus efficace)**
   - Utiliser une pompe de détartrage (location ou achat)
   - Connecter la pompe en circuit fermé sur les orifices de l'échangeur
   - Faire circuler la solution détartrante chaude (40-50°C) pendant 1-2 heures
   - Inverser régulièrement le sens de circulation pour décoller le calcaire
   - Cette méthode est plus efficace que le trempage

6. **Rinçage de l'échangeur**
   - Rincer abondamment l'échangeur à l'eau claire
   - Faire circuler de l'eau propre dans les deux circuits (primaire et ECS)
   - Éliminer toute trace de produit détartrant
   - Vérifier visuellement entre les plaques : doivent être propres, aspect métallique brillant

7. **Contrôle de l'état des joints**
   - Inspecter les joints toriques entre les plaques
   - Si joints durcis, fissurés ou déformés : remplacer (kit joints disponible par marque/modèle)
   - Ne jamais remonter un échangeur avec des joints abîmés (risque de fuite interne)

8. **Remontage de l'échangeur**
   - Remonter l'échangeur dans son logement
   - Revisser les vis de fixation
   - Reconnecter les 4 raccords hydrauliques en respectant le repérage (chauffage/ECS, entrée/sortie)
   - Vérifier le bon positionnement des joints sur chaque raccord

9. **Remise en service et test**
   - Ouvrir les vannes d'isolement (chauffage et ECS)
   - Remplir lentement le circuit chauffage (purger l'air)
   - Remettre sous tension et gaz
   - Tester la production ECS :
     - Ouvrir un robinet eau chaude
     - Vérifier le débit (doit être revenu à la normale)
     - Vérifier la température (doit être stable et selon consigne)
   - Contrôler l'absence de fuite au niveau de l'échangeur et des raccords
   - Tester le mode chauffage (vérifier que pas de fuite interne chauffage/ECS)

10. **Prévention du ré-entartrage**
    - Réduire la température ECS à 50-55°C (confort suffisant, moins de calcaire)
    - Faire un cycle anti-légionellose hebdomadaire (70°C pendant 10 min)
    - Installer un adoucisseur d'eau si TH > 25°F (eau très dure)
    - Détartrage préventif tous les 2-3 ans selon dureté de l'eau
    - Utiliser un filtre polyphosphates sur l'arrivée ECS

**Prévention :**
- Entretien préventif tous les 2-3 ans en zone calcaire
- Régler la température ECS à 50-55°C
- Installer un adoucisseur si eau très dure (TH > 30°F)

**Spécificités techniques :**
- Échangeur à plaques en acier inoxydable (résistant à la corrosion)
- Nombre de plaques : 10 à 20 selon modèle et puissance
- Détection fuite interne : eau chaude sanitaire trouble ou présence d'air dans le circuit chauffage

**Avertissements sécurité :**
- Manipuler les produits détartrants avec précaution (gants, lunettes)
- Ne pas utiliser d'acide chlorhydrique concentré (risque de corrosion de l'inox)
- Bien rincer l'échangeur pour éviter tout résidu corrosif
- Ne jamais remonter un échangeur avec des joints défectueux (risque de fuite interne chauffage/ECS = danger sanitaire)

---

## FACT-CHAUD-050: Vanne 3 voies défectueuse

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-050 |
| **Catégorie** | Hydraulique |
| **Système** | Vanne 3 voies / Inversion |
| **Gravité** | **Moyenne à Élevée** |
| **Marques** | Chaudières mixtes (Saunier Duval, Vaillant, Frisquet, Chaffoteaux) |

**Symptômes :**
- Pas d'eau chaude sanitaire ou ECS tiède malgré chaudière en chauffe
- Radiateurs ne chauffent plus alors que chaudière fonctionne
- Chaudière bascule en ECS mais radiateurs restent chauds (ou inverse)
- Fuite d'eau au niveau de la vanne 3 voies
- Bruit de moteur de vanne (grésillements) sans basculement

**Cause racine probable :**
Moteur de vanne 3 voies HS ou grippé, mécanisme de vanne bloqué par calcaire ou impuretés, joint de tige défectueux (fuite), connectique électrique défectueuse, micro-switch de position HS.

**Étapes de résolution :**

1. **Principe de fonctionnement de la vanne 3 voies**
   - La vanne 3 voies dirige le flux d'eau chaude soit vers le circuit chauffage, soit vers l'échangeur ECS
   - Commandée par un servomoteur électrique (moteur pas-à-pas ou moteur à came)
   - Temps de basculement : 15 à 60 secondes selon modèle
   - Positions : Chauffage / ECS / (parfois position intermédiaire)

2. **Diagnostic du problème**
   - **Test 1** : demander ECS (ouvrir robinet) → observer si basculement (bruit moteur vanne)
   - **Test 2** : toucher les tuyaux : tuyau ECS doit devenir chaud, tuyau chauffage doit refroidir
   - **Test 3** : observer le mouvement de la tige de vanne (tige mobile visible sur certains modèles)

3. **Vérification électrique du moteur de vanne**
   - Couper l'alimentation électrique
   - Accéder au moteur de la vanne 3 voies (démonter capot chaudière)
   - Repérer les connexions électriques (généralement 3 à 5 fils)
   - Remettre sous tension en mode test ECS
   - Mesurer la tension aux bornes du moteur (multimètre) : 230V AC attendu lors de la commande
   - **Si pas de tension** : problème de commande (carte électronique) ou micro-switch
   - **Si tension présente mais pas de mouvement** : moteur grippé ou HS

4. **Test du moteur de vanne (déblocage manuel)**
   - Certains moteurs ont une molette de déblocage manuel (rotation manuelle)
   - Tourner la molette pour basculer manuellement la vanne
   - Si rotation difficile ou impossible : mécanisme de vanne grippé
   - Si rotation facile mais pas de mouvement automatique : moteur HS

5. **Démontage et nettoyage du mécanisme de vanne**
   - Couper électricité, gaz, et fermer les vannes d'isolement
   - Vidanger partiellement le circuit primaire
   - Démonter le moteur de la vanne (2 à 4 vis)
   - Extraire la tige de commande (observer l'état : calcaire, encrassement ?)
   - Nettoyer la tige et le logement (brosse, vinaigre blanc si calcaire)
   - Graisser légèrement la tige avec graisse silicone résistante haute température
   - Actionner manuellement la vanne (tourner l'axe interne) : doit être fluide

6. **Contrôle des joints de tige**
   - Inspecter les joints toriques de la tige de vanne
   - Si joints durcis, déformés ou fuite visible : remplacer (kit joints par marque/modèle)
   - Remonter avec joints neufs graissés

7. **Remplacement du moteur de vanne**
   - Si le moteur est HS (bobinage coupé, engrenages cassés)
   - Commander le moteur de remplacement (référence exacte selon marque/modèle chaudière)
   - Démonter l'ancien moteur
   - Positionner le nouveau moteur sur la vanne (respect du calage : généralement position chauffage)
   - Visser le moteur sur son support
   - Reconnecter les fils électriques (respecter le code couleur/repérage)

8. **Remplacement complet de la vanne 3 voies**
   - Si le corps de vanne est fissuré, corrodé ou irréparable
   - Couper eau, gaz, électricité, vidanger
   - Démonter les raccords hydrauliques (3 raccords : chauffage, ECS, retour commun)
   - Extraire la vanne complète (avec moteur)
   - Installer la nouvelle vanne en respectant les positions (repérage des entrées/sorties)
   - Utiliser des joints neufs ou téflon sur les raccords filetés
   - Remonter le moteur sur la nouvelle vanne

9. **Remise en service et test**
   - Remplir le circuit, purger l'air
   - Remettre sous tension
   - Initialisation de la vanne (certaines chaudières effectuent un auto-test au démarrage)
   - Tester le mode chauffage : radiateurs doivent chauffer, ECS doit être coupée
   - Tester le mode ECS : ouvrir un robinet, observer le basculement (bruit moteur), ECS doit arriver chaude
   - Vérifier l'absence de fuite au niveau de la vanne

10. **Vérification du micro-switch de position**
    - Certains modèles ont un micro-switch qui détecte la position de la vanne
    - Si défectueux : la chaudière ne détecte pas la bonne position
    - Tester la continuité du micro-switch (multimètre) en actionnant manuellement
    - Remplacer si défectueux

**Prévention :**
- Actionner régulièrement la vanne (utiliser ECS même en été) pour éviter grippage
- Traiter l'eau du circuit primaire (inhibiteur) pour limiter les dépôts
- Vérifier le bon fonctionnement lors de l'entretien annuel

**Spécificités techniques :**
- Vanne 3 voies motorisée : temps de basculement 15-60 secondes
- Durée de vie moyenne : 10-15 ans
- Types de vannes : à tournant sphérique, à clapet, à secteur

**Avertissements sécurité :**
- Toujours couper l'alimentation électrique avant de démonter le moteur
- Vidanger le circuit avant de démonter la vanne (risque de brûlure et inondation)
- Respecter le sens de montage et le calage du moteur (risque de fonctionnement inversé)

---

**Fin du fichier 03_Hydraulique.md**