# Problèmes Brûleur & Combustion

**Catégorie :** Brûleur & Combustion
**Nombre de Facts :** 20
**Retour à l'index :** [Knowledge_Base_Chaudieres.md](Knowledge_Base_Chaudieres.md)

---

## FACT-CHAUD-001: Chaudière ne démarre pas - Erreur flamme absente

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-001 |
| **Catégorie** | Brûleur & Combustion |
| **Système** | Allumage gaz / ionisation |
| **Gravité** | **Élevée** |
| **Marques** | Multi-marques (Saunier Duval F28, Vaillant F29, Frisquet 501) |

**Symptômes :**
- Tentatives d'allumage visibles (étincelle) mais pas de flamme
- Code erreur "absence de flamme" ou "défaut ionisation"
- La chaudière se met en sécurité après 3-5 essais
- Bruit de claquement à chaque tentative

**Cause racine probable :**
Absence de gaz au brûleur, pression gaz insuffisante, électrode d'allumage/ionisation mal positionnée ou encrassée, mélange air/gaz incorrect, vanne gaz défectueuse.

**Étapes de résolution :**

1. **Sécurité et environnement**
   - Couper l'alimentation électrique et fermer le robinet de gaz si odeur suspecte
   - S'assurer de la ventilation du local conformément à la réglementation gaz
   - Vérifier l'absence de fuite gaz (détecteur ou eau savonneuse)

2. **Contrôle alimentation gaz**
   - Vérifier que le robinet de gaz chaudière est complètement ouvert
   - Vérifier la pression gaz amont (manomètre / prise de pression selon constructeur)
   - Pression gaz réseau : 20-25 mbar (gaz naturel) ou 28-37 mbar (propane)
   - Purger l'air des conduites si installation récente

3. **Contrôle électrode d'allumage / ionisation**
   - Couper l'alimentation, déposer le brûleur
   - Nettoyer l'électrode avec chiffon doux non abrasif
   - Vérifier l'écartement : généralement 3-4 mm selon notice constructeur
   - Vérifier la position : l'électrode doit être dans la zone de flamme
   - Contrôler l'état du câble haute tension et des connectiques (pas de fissures)

4. **Contrôle du brûleur et de la rampe gaz**
   - Nettoyer le brûleur si encrassement (poussières, oxydation, toiles d'araignées)
   - Vérifier la propreté des orifices de la rampe gaz
   - Souffler délicatement si bouchés (air comprimé faible pression)
   - Contrôler l'état du joint brûleur

5. **Contrôle vanne gaz**
   - Vérifier que la vanne gaz s'ouvre bien (bruit caractéristique)
   - Contrôler l'alimentation électrique de la vanne (tension bobine)
   - Tester la continuité des bobines (multimètre : résistance attendue 3-5 kΩ)

6. **Contrôle mélange air/gaz**
   - Sur chaudière à ventilateur : vérifier que le ventilateur démarre correctement
   - Contrôler les conduits d'air/fumées (obstruction possible : nid d'oiseau, givre)
   - Vérifier le pressostat air si présent

7. **Remise en service et tests**
   - Remonter l'ensemble, remettre gaz et électricité
   - Lancer un cycle de démarrage
   - Contrôler la stabilité de la flamme et les valeurs de combustion (CO₂: 8-10%, O₂: 5-7%)
   - Mesurer la température de fumées

**Prévention :**
- Nettoyage régulier du brûleur et contrôle de l'électrode à chaque entretien annuel
- Vérification de la pression gaz et des conduits d'air/fumées
- Contrôle annuel de la vanne gaz et de ses connectiques
- Installation de filtre à gaz si zone poussiéreuse

---

## FACT-CHAUD-002: Flamme instable ou extinction en cours de fonctionnement

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-002 |
| **Catégorie** | Brûleur & Combustion |
| **Système** | Stabilité flamme / ionisation |
| **Gravité** | **Élevée** |
| **Marques** | Multi-marques (Vaillant, Saunier Duval, Elm Leblanc) |

**Symptômes :**
- Flamme qui s'allume puis s'éteint après quelques secondes
- Cycles d'allumage/extinction répétés
- Flamme irrégulière, jaune ou qui "bat"
- Code erreur détection flamme perdue

**Cause racine probable :**
Électrode d'ionisation encrassée ou mal positionnée, courant d'ionisation faible, masse défectueuse, carte électronique défaillante, tirage perturbé.

**Étapes de résolution :**

1. **Contrôle électrode d'ionisation**
   - Nettoyer soigneusement l'électrode (papier abrasif fin grain 600)
   - Vérifier la position dans la flamme (zone la plus chaude)
   - Contrôler l'isolation du câble (pas de craquelures)
   - Mesurer le courant d'ionisation : doit être > 2-3 µA

2. **Contrôle masse et liaisons électriques**
   - Vérifier la connexion masse de la chaudière
   - Contrôler le serrage des connexions électrode
   - Nettoyer les cosses et connecteurs (contact franc)

3. **Contrôle qualité de combustion**
   - Vérifier couleur flamme : doit être bleue, stable
   - Flamme jaune = manque d'air ou brûleur encrassé
   - Nettoyer le brûleur et ajuster l'air primaire si nécessaire
   - Mesurer CO₂ et O₂

4. **Contrôle tirage et évacuation**
   - Vérifier le tirage (dépression fumées)
   - Contrôler absence d'obstruction conduit
   - Vérifier le fonctionnement du ventilateur extracteur

5. **Contrôle carte électronique**
   - Vérifier les paramètres de détection flamme
   - Tester avec une électrode de rechange si disponible
   - Si défaut persiste : suspecter carte ionisation défectueuse

**Prévention :**
- Nettoyage électrode d'ionisation à chaque entretien
- Contrôle annuel des liaisons électriques
- Vérification qualité combustion (analyse fumées)
- Entretien régulier conduit de fumées

---

## FACT-CHAUD-003: Bruit anormal au démarrage du brûleur (boum)

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-003 |
| **Catégorie** | Brûleur & Combustion |
| **Système** | Séquence allumage |
| **Gravité** | **Élevée** |
| **Marques** | Multi-marques (particulièrement chaudières anciennes) |

**Symptômes :**
- Bruit de détonation ("boum") au démarrage
- Vibration de la chaudière à l'allumage
- Flamme qui apparaît brutalement
- Parfois ouverture de la trappe de visite

**Cause racine probable :**
Allumage retardé provoquant une accumulation de gaz avant inflammation (pré-allumage tardif), électrode usée, séquence vanne gaz incorrecte, temps de pré-ventilation insuffisant.

**Étapes de résolution :**

1. **Contrôle séquence d'allumage**
   - Observer la séquence : ventilateur → étincelle → gaz → flamme
   - Le délai entre étincelle et ouverture gaz doit être < 1 seconde
   - Vérifier que l'étincelle est présente avant l'arrivée de gaz

2. **Contrôle électrode d'allumage**
   - Vérifier l'écartement (3-4 mm selon modèle)
   - Nettoyer et repositionner si nécessaire
   - Contrôler la qualité de l'étincelle : doit être franche et régulière
   - Vérifier le transformateur haute tension

3. **Contrôle vanne gaz et débit**
   - Vérifier que la vanne ne fuit pas en position fermée
   - Contrôler le réglage du débit gaz au démarrage (progressif)
   - Sur vannes à 2 étages : vérifier ouverture graduelle
   - Mesurer pression gaz en fonctionnement

4. **Contrôle pré-ventilation**
   - Vérifier le temps de pré-ventilation (balayage chambre combustion)
   - Durée normale : 10-30 secondes selon modèle
   - Augmenter si paramétrable sur carte électronique

5. **Contrôle chambre de combustion**
   - Vérifier l'absence d'accumulation de gaz résiduel
   - Nettoyer la chambre de combustion
   - Contrôler l'étanchéité du corps de chauffe

6. **Ajustement paramètres**
   - Réduire le débit gaz au démarrage si paramétrable
   - Ajuster le temps de pré-allumage
   - Vérifier la modulation de puissance

**Prévention :**
- Contrôle régulier de la séquence d'allumage
- Vérification annuelle électrode et transformateur HT
- Réglage précis de la vanne gaz
- Nettoyage chambre de combustion

---

## FACT-CHAUD-004: Consommation gaz excessive

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-004 |
| **Catégorie** | Brûleur & Combustion |
| **Système** | Rendement combustion |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques |

**Symptômes :**
- Facture de gaz anormalement élevée
- Température de fumées élevée (> 200°C sur chaudière standard)
- Mauvais rendement constaté
- CO₂ faible dans fumées (< 8%)

**Cause racine probable :**
Excès d'air de combustion, brûleur mal réglé, échangeur encrassé (suies), isolation défectueuse, régulation mal paramétrée, sur-dimensionnement.

**Étapes de résolution :**

1. **Analyse combustion**
   - Mesurer CO₂, O₂, CO dans les fumées
   - Valeurs cibles : CO₂ = 8,5-10%, O₂ = 5-7%, CO < 100 ppm
   - Mesurer température fumées
   - Calculer le rendement instantané

2. **Contrôle et réglage brûleur**
   - Nettoyer le brûleur (têtes injection, surfaces rayonnement)
   - Ajuster le débit d'air (volet air ou paramètre ventilateur)
   - Régler la pression gaz au brûleur selon notice
   - Optimiser le ratio air/gaz pour CO₂ maximal (9-10%)

3. **Contrôle échangeur**
   - Inspecter l'échangeur primaire (suies, tartre)
   - Nettoyer mécaniquement ou chimiquement si encrassé
   - Sur condensation : vérifier drainage condensats
   - Vérifier absence de fuite eau/fumées

4. **Contrôle régulation**
   - Vérifier la courbe de chauffe (loi d'eau)
   - Ajuster la température de consigne selon émetteurs
   - Activer/optimiser la sonde extérieure
   - Vérifier les temporisations et cycles marche/arrêt

5. **Contrôle isolation et pertes**
   - Vérifier l'isolation du corps de chauffe
   - Contrôler l'étanchéité de la porte brûleur
   - Vérifier les pertes au conduit de fumées

6. **Optimisation installation**
   - Équilibrer le réseau hydraulique
   - Vérifier calorifugeage tuyauteries
   - Optimiser les plages horaires de fonctionnement

**Prévention :**
- Analyse combustion annuelle obligatoire
- Nettoyage régulier brûleur et échangeur
- Révision paramètres régulation après travaux isolation
- Formation utilisateur sur bonnes pratiques

---

## FACT-CHAUD-005: Flamme jaune ou orange - Mauvaise combustion

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-005 |
| **Catégorie** | Brûleur & Combustion |
| **Système** | Qualité combustion |
| **Gravité** | **Critique** |
| **Marques** | Multi-marques |

**Symptômes :**
- Flamme de couleur jaune, orange ou rougeâtre au lieu de bleue
- Dépôts de suie sur l'échangeur ou le brûleur
- Taux de CO élevé dans les fumées (> 100 ppm)
- Odeur de combustion anormale

**Cause racine probable :**
Manque d'air de combustion (colmatage air primaire, ventilateur défaillant), brûleur encrassé, mauvais réglage air/gaz, conduit fumées obstrué, air vicié dans le local.

**Étapes de résolution :**

1. **SÉCURITÉ IMMÉDIATE**
   - ARRÊTER la chaudière immédiatement
   - Aérer le local
   - Mesurer le CO ambiant avec détecteur
   - Ne pas redémarrer avant résolution

2. **Analyse combustion complète**
   - Mesurer CO, CO₂, O₂, température fumées
   - Mesurer le tirage (dépression)
   - Calculer l'excès d'air (doit être ~20-40%)
   - Photographier la flamme pour référence

3. **Contrôle entrée d'air**
   - Vérifier les grilles de ventilation du local (non obstruées)
   - Sur chaudière ventouse : vérifier terminal non bouché (givre, nid)
   - Sur chaudière atmosphérique : vérifier amenée d'air conforme réglementation
   - Nettoyer les filtres à air

4. **Contrôle brûleur**
   - Démonter et nettoyer complètement le brûleur
   - Éliminer toutes les suies et dépôts
   - Nettoyer les orifices d'injection gaz
   - Vérifier l'état du brûleur (corrosion, déformation)

5. **Contrôle conduit fumées**
   - Vérifier l'absence d'obstruction (ramonage si nécessaire)
   - Contrôler le tirage et la dépression
   - Vérifier l'étanchéité du conduit (pas de refoulement)
   - Sur chaudière ventouse : vérifier ventilateur extraction

6. **Réglage combustion**
   - Augmenter l'air de combustion progressivement
   - Ajuster jusqu'à flamme bleue stable
   - Mesurer et ajuster pour atteindre : CO₂ = 9-10%, CO < 50 ppm
   - Vérifier sur toutes les puissances (mini et maxi)

7. **Contrôle échangeur**
   - Nettoyer les suies déposées sur l'échangeur
   - Vérifier absence de perforation échangeur
   - Contrôler l'étanchéité fumées/eau

**Prévention :**
- Vérification couleur flamme à chaque entretien
- Analyse combustion complète annuelle
- Entretien préventif entrées d'air et conduit fumées
- Détecteur CO dans les locaux chaufferie

**⚠️ DANGER :** Une flamme jaune produit du monoxyde de carbone (CO) toxique. Intervention immédiate obligatoire par professionnel qualifié.

---

## FACT-CHAUD-006: Code erreur ventilateur - Saunier Duval F33

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-006 |
| **Catégorie** | Brûleur & Combustion |
| **Système** | Ventilateur extraction |
| **Gravité** | **Élevée** |
| **Marques** | Saunier Duval (ThemaPlus, Isotwin, Isofast) |

**Symptômes :**
- Code erreur F33 affiché
- Ventilateur ne démarre pas ou vitesse insuffisante
- Chaudière bloquée, pas de tentative d'allumage
- Parfois bruit anormal du ventilateur

**Cause racine probable :**
Ventilateur bloqué mécaniquement, condensats gelés, carte électronique défaillante, connectique défectueuse, pressostat air ne commute pas, roulement ventilateur grippé.

**Étapes de résolution :**

1. **Diagnostic visuel**
   - Vérifier si le ventilateur tente de démarrer (bruit)
   - Contrôler l'absence de blocage mécanique (corps étranger)
   - Vérifier l'état des pales (déformation, casse)

2. **Contrôle alimentation électrique**
   - Mesurer la tension d'alimentation ventilateur (généralement 230V AC)
   - Vérifier la connectique (cosses bien enfichées)
   - Contrôler l'état des câbles (pas de coupure)

3. **Test ventilateur**
   - Déconnecter et tester le ventilateur en direct 230V (précautions !)
   - Si ne tourne pas : ventilateur HS, prévoir remplacement
   - Si tourne : suspecter carte ou pressostat

4. **Contrôle pressostat air**
   - Localiser le pressostat air (entre ventilateur et chambre combustion)
   - Vérifier les tubes de prise de pression (non bouchés, non pincés)
   - Tester la commutation du pressostat (multimètre continuité)
   - Souffler dans le tube pour forcer la commutation (test)

5. **Contrôle carte électronique**
   - Vérifier les paramètres ventilateur dans le menu diagnostic
   - Contrôler la mesure de vitesse (signal tachymétrique)
   - Si pas de retour tachymétrique : câble défectueux ou ventilateur HS

6. **Contrôle condensats**
   - Vérifier que le siphon condensats n'est pas gelé (hors gel local)
   - Dégivrer si nécessaire
   - Vérifier drainage correct des condensats

7. **Remplacement si nécessaire**
   - Ventilateur : référence selon modèle (ex: 0020073793)
   - Pressostat air : référence selon modèle
   - Prévoir joint ventilateur

**Prévention :**
- Vérification annuelle vitesse ventilateur (menu diagnostic)
- Nettoyage pales ventilateur et pressostat
- Contrôle tubes de pression
- Lubrification roulements si prévu constructeur
- Protection hors gel local chaufferie

---

## FACT-CHAUD-007: Arrêt sur sécurité pressostat fumées

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-007 |
| **Catégorie** | Brûleur & Combustion |
| **Système** | Sécurité évacuation fumées |
| **Gravité** | **Critique** |
| **Marques** | Multi-marques (chaudières atmosphériques principalement) |

**Symptômes :**
- Arrêt chaudière après quelques secondes de fonctionnement
- Code erreur sécurité fumées ou pressostat
- Le pressostat ne se réarme pas
- Parfois refoulement de fumées dans le local

**Cause racine probable :**
Conduit de fumées obstrué, tirage insuffisant, pressostat défectueux, chambre de combustion encrassée, anomalie raccordement conduit.

**Étapes de résolution :**

1. **SÉCURITÉ IMMÉDIATE**
   - Ne pas réarmer tant que la cause n'est pas identifiée
   - Vérifier absence de fumées dans le local
   - Aérer si nécessaire

2. **Contrôle conduit de fumées**
   - Vérifier visuellement le débouché en toiture (non obstrué)
   - Contrôler l'absence de nid d'oiseau, feuilles, givre
   - Vérifier le diamètre et la hauteur du conduit (conformité réglementation)
   - Ramonage si encrassement suspecté

3. **Mesure du tirage**
   - Mesurer la dépression au niveau du coupe-tirage (5-10 Pa minimum)
   - Vérifier que le tirage s'établit avant allumage
   - Contrôler les conditions météo (vent, pression atmosphérique)

4. **Contrôle pressostat fumées**
   - Localiser le pressostat (généralement au coupe-tirage)
   - Vérifier le tube de prise de pression (non bouché, non pincé)
   - Tester la continuité électrique (contact fermé au repos)
   - Souffler légèrement dans le tube : le contact doit s'ouvrir
   - Remplacer si défectueux

5. **Contrôle chambre de combustion**
   - Nettoyer la chambre de combustion
   - Vérifier les passages fumées dans l'échangeur
   - Éliminer les suies et dépôts

6. **Contrôle étanchéité circuit fumées**
   - Vérifier les joints de la porte brûleur
   - Contrôler le coupe-tirage (position, fixation)
   - Vérifier le raccordement au conduit (étanche)

7. **Vérification réglementaire**
   - S'assurer que le conduit est conforme (tubage, hauteur, section)
   - Vérifier la ventilation du local (amenée d'air suffisante)

**Prévention :**
- Ramonage annuel obligatoire du conduit
- Vérification tirage à chaque entretien
- Contrôle état pressostat
- Nettoyage régulier chambre combustion et échangeur

**⚠️ DANGER :** Le pressostat fumées est une sécurité vitale. Ne jamais neutraliser ou shunter. Risque d'intoxication CO.

---

## FACT-CHAUD-008: Défaut allumage Vaillant F28/F29

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-008 |
| **Catégorie** | Brûleur & Combustion |
| **Système** | Séquence allumage Vaillant |
| **Gravité** | **Élevée** |
| **Marques** | Vaillant (ecoTEC, turboTEC, atmoTEC) |

**Symptômes :**
- Code F28 : échec allumage (3 tentatives sans succès)
- Code F29 : extinction flamme en cours de fonctionnement
- Réarmement manuel nécessaire
- Parfois alternance F28/F29

**Cause racine probable :**
F28 : absence gaz, vanne gaz défectueuse, électrode, transformateur HT. F29 : ionisation défaillante, coupure gaz, instabilité flamme, carte électronique.

**Étapes de résolution F28 :**

1. **Contrôle gaz**
   - Vérifier robinet gaz ouvert
   - Mesurer pression gaz réseau (20-25 mbar)
   - Purger l'air si installation récente
   - Contrôler compteur gaz (pas de coupure)

2. **Contrôle vanne gaz**
   - Écouter le clic d'ouverture de la vanne
   - Mesurer tension bobines vanne gaz (230V)
   - Mesurer résistance bobines (3-5 kΩ)
   - Vérifier pression sortie vanne en fonctionnement

3. **Contrôle allumage**
   - Vérifier étincelle franche et régulière
   - Nettoyer et ajuster électrode (écartement 4 mm)
   - Contrôler transformateur HT (sortie ~8-10 kV)
   - Vérifier câble HT (pas de fuite à la masse)

4. **Contrôle ventilateur**
   - Vérifier démarrage ventilateur avant allumage (pré-ventilation)
   - Contrôler vitesse ventilateur (menu d.40)
   - Nettoyer si encrassé

**Étapes de résolution F29 :**

1. **Contrôle ionisation**
   - Nettoyer électrode ionisation (papier abrasif fin)
   - Vérifier position dans flamme
   - Mesurer courant ionisation (menu d.48 : > 2 µA)
   - Contrôler câble et connexion masse

2. **Contrôle combustion**
   - Vérifier couleur flamme (bleue)
   - Analyser fumées (CO₂, O₂, CO)
   - Nettoyer brûleur si flamme instable
   - Vérifier débit gaz (pas de variation)

3. **Contrôle carte électronique**
   - Vérifier paramètres ionisation (menu d.96)
   - Réinitialiser paramètres usine si nécessaire
   - Mettre à jour firmware si disponible
   - Remplacer carte si défaut persiste

**Prévention :**
- Contrôle annuel électrodes et transformateur
- Vérification pression gaz
- Analyse combustion
- Mise à jour firmware régulière

**Spécificités Vaillant :**
- Accès menu diagnostic : maintenir touches Mode + OK
- d.40 : vitesse ventilateur (tours/min)
- d.48 : courant ionisation (µA)
- d.96 : paramètres ionisation

---

## FACT-CHAUD-009: Blocage De Dietrich E25 - Défaut gaz

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-009 |
| **Catégorie** | Brûleur & Combustion |
| **Système** | Détection gaz De Dietrich |
| **Gravité** | **Élevée** |
| **Marques** | De Dietrich (City, MCX, Naneo) |

**Symptômes :**
- Code E25 affiché
- Blocage chaudière
- Message "Défaut gaz" ou "Pression gaz"
- 3 clignotements LED

**Cause racine probable :**
Pression gaz insuffisante, vanne gaz défectueuse, pressostat gaz mal réglé ou défaillant, carte électronique, fuite gaz en amont.

**Étapes de résolution :**

1. **Mesure pression gaz**
   - Brancher manomètre sur prise de pression
   - Mesurer pression statique robinet fermé
   - Mesurer pression dynamique chaudière en marche
   - Valeurs attendues : statique 20-25 mbar, dynamique 20 mbar mini

2. **Contrôle pressostat gaz**
   - Localiser le pressostat gaz (sur rampe ou vanne)
   - Vérifier le réglage (généralement 15-18 mbar)
   - Tester la commutation (continuité électrique)
   - Ajuster si nécessaire selon procédure constructeur

3. **Contrôle vanne gaz**
   - Vérifier la tension d'alimentation vanne
   - Contrôler les bobines (résistance)
   - Vérifier absence de blocage mécanique
   - Tester l'ouverture progressive

4. **Contrôle installation gaz**
   - Vérifier le détendeur gaz (si propane/butane)
   - Contrôler le diamètre de la tuyauterie gaz
   - Vérifier absence de fuite (perte de pression)
   - S'assurer du bon dimensionnement

5. **Réglage pressostat**
   - Accéder au réglage (vis sur pressostat)
   - Ajuster progressivement (1/4 tour max)
   - Tester le démarrage entre chaque ajustement
   - Consigner le réglage final

6. **Contrôle carte électronique**
   - Vérifier les paramètres gaz dans le menu
   - Réinitialiser si nécessaire
   - Remplacer carte si défaut persiste

**Prévention :**
- Vérification annuelle pression gaz
- Contrôle pressostat gaz
- Test étanchéité installation gaz
- Vérification détendeur (si GPL)

**Spécificités De Dietrich :**
- Réglage pressostat : vis accessibles après démontage capot
- Procédure réarmage : bouton Reset 2 secondes
- Menu diagnostic : touches ▲ + ▼ simultanément

---

## FACT-CHAUD-010: Encrassement brûleur - Production suie excessive

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-010 |
| **Catégorie** | Brûleur & Combustion |
| **Système** | Brûleur gaz |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques |

**Symptômes :**
- Dépôts noirs (suie) sur le brûleur et l'échangeur
- Diminution progressive des performances
- Flamme qui devient jaunâtre
- Augmentation température fumées
- Bruit de combustion anormal

**Cause racine probable :**
Manque d'air, brûleur encrassé (poussières, oxydation), mauvais réglage air/gaz, air ambiant pollué (poussières, produits chimiques), vieillissement brûleur.

**Étapes de résolution :**

1. **Diagnostic encrassement**
   - Photographier l'état du brûleur (avant/après)
   - Évaluer l'épaisseur des dépôts
   - Identifier le type de dépôt (suie, rouille, calcaire)
   - Mesurer CO et CO₂ avant intervention

2. **Démontage et nettoyage brûleur**
   - Couper gaz et électricité
   - Déposer le brûleur selon notice constructeur
   - Brosser mécaniquement (brosse douce)
   - Tremper dans produit dégraissant si nécessaire
   - Nettoyer les orifices d'injection (fil laiton ou soufflette)
   - Rincer et sécher complètement

3. **Nettoyage échangeur**
   - Brosser les surfaces d'échange côté fumées
   - Aspirer les suies et dépôts
   - Utiliser produit spécifique si encrassement important
   - Vérifier absence de perforation

4. **Analyse cause racine**
   - Vérifier ventilation du local (amenée d'air suffisante)
   - Contrôler absence de produits volatils (solvants, aérosols)
   - Vérifier état du filtre à air (si présent)
   - Analyser la combustion

5. **Réglage combustion optimal**
   - Remonter le brûleur avec joint neuf
   - Régler l'air de combustion
   - Mesurer CO₂ (cible 9-10%), O₂ (5-7%), CO (< 50 ppm)
   - Ajuster jusqu'à combustion optimale
   - Vérifier sur toutes les puissances

6. **Contrôle installation**
   - Vérifier les grilles de ventilation (non obstruées)
   - S'assurer d'une ventilation conforme
   - Éloigner produits chimiques de la chaufferie

**Prévention :**
- Nettoyage annuel obligatoire du brûleur
- Ventilation correcte du local chaufferie
- Interdiction stockage produits chimiques volatils
- Filtre à air si environnement poussiéreux
- Analyse combustion annuelle

**Fréquence nettoyage selon environnement :**
- Résidentiel propre : 1 fois/an
- Résidentiel poussiéreux : 2 fois/an
- Commercial/industriel : 3-4 fois/an
- Agricole : tous les 3-6 mois

---

## FACT-CHAUD-011: Modulation puissance défectueuse

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-011 |
| **Catégorie** | Brûleur & Combustion |
| **Système** | Modulation puissance |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques (chaudières modulantes) |

**Symptômes :**
- Chaudière fonctionne uniquement en tout ou rien
- Pas de modulation de puissance
- Cycles marche/arrêt fréquents
- Rendement dégradé
- Température départ instable

**Cause racine probable :**
Vanne gaz à modulation défectueuse, carte électronique, sonde température défaillante, paramètres mal configurés, servomoteur vanne bloqué.

**Étapes de résolution :**

1. **Diagnostic modulation**
   - Observer le comportement : puissance fixe ou variable
   - Consulter le menu diagnostic (affichage % puissance)
   - Vérifier si la puissance demandée varie selon besoin
   - Écouter le bruit du brûleur (doit varier)

2. **Contrôle vanne gaz modulante**
   - Vérifier le type de vanne (tout ou rien vs modulante)
   - Contrôler le servomoteur (bruit de fonctionnement)
   - Mesurer la tension de commande (0-10V ou PWM selon modèle)
   - Vérifier la mécanique (pas de blocage)

3. **Contrôle carte électronique**
   - Vérifier les paramètres de puissance (mini/maxi)
   - Contrôler le signal de commande vanne (oscilloscope ou multimètre)
   - Tester le mode manuel si disponible
   - Réinitialiser les paramètres usine

4. **Contrôle sondes température**
   - Vérifier sonde départ (valeur cohérente)
   - Contrôler sonde retour si présente
   - Tester la variation de consigne (impact sur modulation)
   - Remplacer sonde si défectueuse

5. **Contrôle ventilateur modulant**
   - Sur chaudières à ventilateur modulant : vérifier variation vitesse
   - Consulter vitesse dans menu diagnostic
   - Vérifier que la vitesse suit la puissance demandée

6. **Vérification hydraulique**
   - S'assurer d'un débit suffisant (pompe, réseau)
   - Vérifier que le ΔT permet la modulation
   - Contrôler bypass si présent

7. **Optimisation paramètres**
   - Ajuster la puissance maximale selon installation
   - Régler la puissance minimale (éviter cycles courts)
   - Configurer les temps de stabilisation
   - Paramétrer la rampe de montée en puissance

**Prévention :**
- Vérification annuelle de la modulation
- Contrôle paramètres après modification installation
- Lubrification servomoteur si prévu
- Mise à jour firmware carte

**Avantages modulation correcte :**
- Rendement amélioré (+5 à 10%)
- Confort accru (température stable)
- Moins de cycles (longévité accrue)
- Économies d'énergie

---

## FACT-CHAUD-012: Atlantic Perfinox - Défaut allumage E01

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-012 |
| **Catégorie** | Brûleur & Combustion |
| **Système** | Allumage Atlantic |
| **Gravité** | **Élevée** |
| **Marques** | Atlantic (Perfinox, Idra) |

**Symptômes :**
- Code erreur E01 affiché
- 1 clignotement LED rouge
- Pas d'allumage après 3 tentatives
- Réarmement nécessaire

**Cause racine probable :**
Absence gaz, électrode défectueuse, transformateur HT, vanne gaz, carte électronique, conduit fumées obstrué.

**Étapes de résolution :**

1. **Réarmement et observation**
   - Réarmer (bouton Reset ou extinction/rallumage)
   - Observer la séquence d'allumage
   - Noter si étincelle présente
   - Vérifier tentatives (3 essais normalement)

2. **Contrôle alimentation gaz**
   - Vérifier robinet gaz ouvert
   - Mesurer pression gaz (20-25 mbar)
   - Contrôler vanne gaz (ouverture audible)
   - Purger si installation récente

3. **Contrôle système d'allumage**
   - Vérifier étincelle (visible, franche, régulière)
   - Nettoyer électrode allumage
   - Vérifier écartement (3-4 mm)
   - Contrôler câble HT (pas de fuite)
   - Tester transformateur HT (sortie 8-10 kV)

4. **Contrôle brûleur et ionisation**
   - Nettoyer le brûleur
   - Vérifier électrode ionisation (propreté, position)
   - Contrôler le câblage ionisation

5. **Contrôle ventilateur**
   - Vérifier démarrage ventilateur (pré-ventilation)
   - Contrôler pressostat air
   - Nettoyer si encrassé

6. **Contrôle carte électronique**
   - Vérifier LED diagnostic carte
   - Contrôler les connexions
   - Réinitialiser si nécessaire
   - Remplacer si défaut persiste

**Prévention :**
- Entretien annuel complet
- Vérification électrodes et transformateur
- Contrôle pression gaz
- Nettoyage ventilateur

**Spécificités Atlantic Perfinox :**
- Réarmage : bouton marche/arrêt 5 secondes
- LED clignotements : code panne (1 = E01)
- Transformateur spécifique référence 178963
- Notice technique indispensable

---

## FACT-CHAUD-013: Frisquet Hydromotrix - Voyant rouge fixe

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-013 |
| **Catégorie** | Brûleur & Combustion |
| **Système** | Sécurité Frisquet |
| **Gravité** | **Élevée** |
| **Marques** | Frisquet (Hydromotrix, Prestige) |

**Symptômes :**
- Voyant rouge allumé en continu
- Chaudière en sécurité
- Pas de tentative d'allumage
- Nécessite réarmement manuel

**Cause racine probable :**
Défaut ionisation (sécurité anti-gazage), thermostats de sécurité déclenchés, pressostat fumées ouvert, défaut électronique.

**Étapes de résolution :**

1. **Identification cause sécurité**
   - Voyant rouge fixe = sécurité générale
   - Vérifier les thermostats de sécurité (réarmables)
   - Contrôler le pressostat fumées
   - Vérifier historique pannes si afficheur

2. **Contrôle thermostats sécurité**
   - Localiser les thermostats (généralement 2-3)
   - Vérifier s'ils sont déclenchés (bouton sorti)
   - Réarmer en appuyant sur le bouton
   - Chercher cause surchauffe si déclenchés

3. **Contrôle pressostat fumées**
   - Vérifier conduit fumées (obstruction)
   - Contrôler tirage
   - Tester continuité pressostat
   - Nettoyer tube de prise de pression

4. **Contrôle ionisation**
   - Nettoyer électrode ionisation
   - Vérifier position et écartement
   - Mesurer courant ionisation
   - Contrôler masse et câblage

5. **Contrôle séquence allumage**
   - Réarmer la chaudière
   - Observer séquence complète
   - Vérifier électrode allumage
   - Contrôler vanne gaz

6. **Spécificités Frisquet**
   - Vérifier l'éco-radio (système propriétaire)
   - Contrôler le module Vision si présent
   - Vérifier les paramètres usine

**Procédure réarmement Frisquet :**
1. Appuyer sur bouton marche/arrêt 2 secondes
2. Voyant rouge doit s'éteindre
3. Chaudière redémarre automatiquement
4. Si blocage persiste après 3 réarmages : diagnostic approfondi

**Prévention :**
- Vérification annuelle thermostats sécurité
- Contrôle pressostat fumées
- Ramonage conduit
- Entretien système ionisation

**Codes pannes Frisquet spécifiques :**
- Voyant rouge fixe : sécurité générale
- Voyant rouge clignotant : défaut temporaire
- Voyant vert clignotant : demande en cours

---

## FACT-CHAUD-014: Elm Leblanc - Erreur EA flamme parasite

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-014 |
| **Catégorie** | Brûleur & Combustion |
| **Système** | Détection flamme Elm Leblanc |
| **Gravité** | Moyenne |
| **Marques** | Elm Leblanc (Egalis, Acleis, Megalis) |

**Symptômes :**
- Code erreur EA affiché
- Message "Flamme parasite" ou "Fausse flamme"
- Chaudière bloquée avant même tentative d'allumage
- Réarmement possible mais erreur réapparaît

**Cause racine probable :**
Électrode ionisation humide ou mal isolée, fuite à la masse du câble ionisation, vanne gaz qui fuit (passage gaz vanne fermée), carte électronique hypersensible ou défectueuse.

**Étapes de résolution :**

1. **Compréhension défaut**
   - EA = détection flamme alors que vanne gaz fermée
   - Sécurité anti-gazage (évite explosion)
   - Ne jamais négliger cette erreur

2. **Contrôle électrode ionisation**
   - Déconnecter l'électrode ionisation
   - Nettoyer et sécher complètement
   - Vérifier l'isolation du câble (multimètre : > 1 MΩ)
   - Contrôler absence d'humidité dans connectique

3. **Test sans électrode**
   - Débrancher électrode ionisation
   - Réarmer la chaudière
   - Si EA persiste : défaut carte électronique
   - Si EA disparaît : problème électrode ou câble

4. **Contrôle vanne gaz**
   - Fermer robinet gaz chaudière
   - Déconnecter tuyau sortie vanne gaz
   - Rouvrir robinet gaz
   - Vanne fermée : aucun gaz ne doit s'échapper
   - Si fuite gaz : vanne défectueuse, remplacement obligatoire

5. **Contrôle étanchéité circuit gaz**
   - Vérifier tous les raccords
   - Test eau savonneuse
   - Resserrer ou remplacer joints si nécessaire

6. **Contrôle carte électronique**
   - Vérifier absence d'humidité sur la carte
   - Sécher si nécessaire (soufflette air sec)
   - Vérifier seuil détection ionisation
   - Remplacer carte si hypersensible

7. **Vérification environnement**
   - Vérifier absence d'humidité excessive local
   - Contrôler étanchéité chaudière (pas de condensation interne)
   - Améliorer ventilation si nécessaire

**Prévention :**
- Vérification annuelle étanchéité vanne gaz
- Contrôle isolation câblage ionisation
- Séchage et nettoyage électrode
- Protection contre humidité local chaufferie

**⚠️ SÉCURITÉ :** Ne jamais neutraliser cette sécurité. Une vanne gaz qui fuit présente un risque d'explosion.

---

## FACT-CHAUD-015: Chaffoteaux - Erreur 501 absence flamme

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-015 |
| **Catégorie** | Brûleur & Combustion |
| **Système** | Allumage Chaffoteaux |
| **Gravité** | **Élevée** |
| **Marques** | Chaffoteaux (Talia, Pigma, Alixia, Urbia) |

**Symptômes :**
- Code erreur 501 affiché
- Blocage après 3 tentatives d'allumage
- Étincelle visible mais pas d'inflammation
- Reset nécessaire

**Cause racine probable :**
Absence gaz, électrode mal positionnée, vanne gaz défectueuse, pression gaz insuffisante, air dans circuit gaz, transformateur HT faible.

**Étapes de résolution :**

1. **Diagnostic initial**
   - Réarmer (bouton Reset 2 secondes)
   - Observer tentatives allumage (nombre, durée)
   - Vérifier présence étincelle
   - Écouter ouverture vanne gaz

2. **Contrôle gaz**
   - Vérifier robinet gaz ouvert
   - Contrôler pression réseau (20-25 mbar)
   - Vérifier compteur gaz (crédit, débit)
   - Purger air si nécessaire

3. **Contrôle électrodes**
   - Déposer ensemble brûleur/électrodes
   - Nettoyer électrode allumage et ionisation
   - Vérifier écartement allumage : 3,5 mm ± 0,5 mm
   - Vérifier position ionisation : dans zone flamme
   - Contrôler câbles HT (pas de craquelure)

4. **Contrôle vanne gaz**
   - Vérifier tension bobines (230V en phase allumage)
   - Mesurer résistance bobines (3-5 kΩ)
   - Contrôler pression sortie vanne
   - Remplacer si défectueuse

5. **Contrôle transformateur HT**
   - Mesurer tension sortie (8-10 kV)
   - Vérifier qualité étincelle (franche, bleue)
   - Remplacer si faible

6. **Contrôle ventilateur et fumées**
   - Vérifier rotation ventilateur
   - Contrôler conduit fumées (pas d'obstruction)
   - Vérifier pressostat air

7. **Contrôle paramètres**
   - Accéder menu installateur
   - Vérifier paramètre type gaz (G20/G25)
   - Contrôler réglage puissance
   - Réinitialiser si nécessaire

**Prévention :**
- Entretien annuel électrodes
- Vérification pression gaz
- Contrôle vanne gaz
- Nettoyage ventilateur

**Spécificités Chaffoteaux :**
- Accès menu : bouton Mode + Reset
- Réglage débit gaz : potentiomètre carte (tournevis)
- Électrode ionisation intégrée brûleur sur certains modèles
- Référence électrode : varie selon modèle (ex: 60001708)

---

## FACT-CHAUD-016: Combustion bruyante - Ronflement anormal

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-016 |
| **Catégorie** | Brûleur & Combustion |
| **Système** | Acoustique combustion |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques (surtout anciennes chaudières) |

**Symptômes :**
- Bruit de ronflement important au brûleur
- Vibrations de la chaudière
- Flamme instable visuellement
- Parfois résonance dans le conduit de fumées
- Gêne sonore importante

**Cause racine probable :**
Mauvais réglage air/gaz, instabilité aérodynamique, résonance acoustique, brûleur encrassé, gicleurs mal adaptés, tirage excessif ou insuffisant.

**Étapes de résolution :**

1. **Analyse du bruit**
   - Caractériser : ronflement grave, sifflement aigu, pulsation
   - Identifier fréquence (grave < 100 Hz, aigu > 1000 Hz)
   - Localiser la source (brûleur, échangeur, conduit)
   - Enregistrer si possible (analyse ultérieure)

2. **Contrôle combustion**
   - Analyser fumées (CO₂, O₂, CO)
   - Vérifier couleur et stabilité flamme
   - Mesurer température fumées
   - Observer le comportement sur différentes puissances

3. **Réglage air/gaz**
   - Ajuster progressivement l'air de combustion
   - Tester effet sur le bruit
   - Optimiser pour silence + combustion correcte
   - Fixer le réglage final

4. **Contrôle mécanique brûleur**
   - Vérifier fixation brûleur (pas de jeu)
   - Contrôler l'état mécanique (déformation)
   - Vérifier les gicleurs (bons diamètres, non bouchés)
   - Nettoyer si encrassé

5. **Contrôle tirage**
   - Mesurer la dépression fumées
   - Un tirage excessif peut créer instabilité
   - Installer/ajuster coupe-tirage si nécessaire
   - Vérifier ventilateur (vitesse adaptée)

6. **Solutions anti-bruit**
   - Installer silencieux sur conduit fumées si nécessaire
   - Améliorer l'isolation phonique de la chaufferie
   - Installer des plots anti-vibrations sous chaudière
   - Vérifier fixations murales (si murale)

7. **Cas particulier condensation**
   - Sur chaudières condensation : bruit peut être lié à condensation
   - Vérifier évacuation condensats (siphon plein)
   - Contrôler température retour (< 54°C pour condenser)

**Prévention :**
- Réglage combustion précis à chaque entretien
- Nettoyage régulier brûleur
- Contrôle fixations
- Vérification tirage

**Note :** Un bruit de combustion peut aussi révéler un défaut de conception (brûleur inadapté). Sur installations anciennes, envisager le remplacement du brûleur par un modèle silencieux.

---

## FACT-CHAUD-017: Odeur de gaz au démarrage

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-017 |
| **Catégorie** | Brûleur & Combustion |
| **Système** | Étanchéité gaz |
| **Gravité** | **Critique** |
| **Marques** | Multi-marques |

**Symptômes :**
- Odeur de gaz perceptible au démarrage ou à l'arrêt
- Odeur dans le local chaufferie
- Parfois odeur persistante
- Détecteur gaz déclenché

**Cause racine probable :**
Fuite vanne gaz, joint défectueux, raccord mal serré, brûleur mal positionné, fissure rampe gaz, allumage retardé (imbrûlés).

**Étapes de résolution :**

1. **SÉCURITÉ IMMÉDIATE**
   - NE PAS allumer/éteindre d'équipement électrique
   - NE PAS fumer ou créer d'étincelle
   - Ouvrir fenêtres et portes pour aérer
   - Fermer le robinet de gaz général
   - Évacuer les personnes si odeur forte
   - Appeler pompiers si odeur importante (18/112)

2. **Détection fuite (après aération)**
   - Utiliser détecteur électronique gaz
   - Ou eau savonneuse sur raccords (bulles = fuite)
   - Tester robinet gaz fermé puis ouvert
   - Localiser précisément la fuite

3. **Contrôle raccords**
   - Vérifier tous les raccords vissés
   - Resserrer si nécessaire (sans forcer)
   - Remplacer joints si défectueux
   - Tester étanchéité après resserrage

4. **Contrôle vanne gaz**
   - Fermer robinet gaz en amont
   - Déposer la vanne si fuite identifiée
   - Remplacer vanne si fuite interne
   - Remonter avec joints neufs
   - Tester étanchéité

5. **Contrôle brûleur et rampe**
   - Vérifier positionnement brûleur (joint étanche)
   - Contrôler rampe gaz (pas de fissure, corrosion)
   - Vérifier gicleurs (bien serrés)
   - Remplacer joints de brûleur

6. **Contrôle allumage**
   - Si odeur liée imbrûlés (allumage retardé)
   - Optimiser séquence allumage
   - Vérifier électrode et transformateur HT
   - Améliorer pré-ventilation

7. **Test étanchéité complet**
   - Rouvrir le gaz progressivement
   - Tester tous les points à l'eau savonneuse
   - Utiliser détecteur électronique
   - Mesurer pression et attendre 15 min (pas de chute pression)
   - Ventiler avant remise en marche

8. **Remise en service**
   - Vérifier absence totale d'odeur gaz
   - Lancer démarrage chaudière
   - Observer allumage (pas d'imbrûlés)
   - Rester présent 30 minutes
   - Installer détecteur CO/gaz si non présent

**Prévention :**
- Contrôle étanchéité annuel obligatoire
- Détecteur gaz dans local chaufferie
- Remplacement préventif joints vieux raccords
- Vérification vanne gaz (durée vie ~10-15 ans)

**⚠️ DANGER MORTEL :**
- Le gaz naturel est explosif (4-14% dans l'air)
- Le CO (combustion incomplète) est mortel
- Toute odeur de gaz = urgence absolue
- Intervention par professionnel qualifié uniquement
- En cas de doute : fermer gaz et appeler pompiers

---

## FACT-CHAUD-018: Température fumées excessive

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-018 |
| **Catégorie** | Brûleur & Combustion |
| **Système** | Échange thermique |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques |

**Symptômes :**
- Température fumées > 200°C (chaudière standard)
- Température fumées > 100°C (chaudière condensation)
- Rendement faible
- Surconsommation gaz
- Échangeur très chaud

**Cause racine probable :**
Échangeur encrassé (suies, tartre), mauvaise circulation eau, débit eau insuffisant, sur-puissance brûleur, mauvais échange thermique.

**Étapes de résolution :**

1. **Mesures et diagnostic**
   - Mesurer température fumées (thermomètre ou analyseur)
   - Mesurer températures eau départ/retour
   - Calculer ΔT eau (départ - retour)
   - Mesurer CO₂, O₂ pour vérifier combustion
   - Calculer rendement instantané

2. **Contrôle échangeur côté fumées**
   - Inspecter visuellement (présence suies)
   - Nettoyer mécaniquement (brosse adaptée)
   - Aspirer dépôts et suies
   - Sur chaudière condensation : vérifier encrassement spécifique
   - Contrôler ailettes échangeur (pas de déformation)

3. **Contrôle échangeur côté eau**
   - Vérifier absence entartrage (eau très dure)
   - Détartrage chimique si nécessaire (produit adapté)
   - Vérifier la qualité eau circuit (dureté, pH)
   - Traiter l'eau si problème récurrent

4. **Contrôle circulation eau**
   - Vérifier fonctionnement pompe (débit suffisant)
   - Contrôler absence obstruction (filtre, vanne)
   - Mesurer ΔT : doit être 15-20°C en chauffage
   - Si ΔT > 25°C : débit insuffisant

5. **Contrôle puissance brûleur**
   - Vérifier que puissance est adaptée à l'installation
   - Mesurer débit gaz (compteur gaz)
   - Réduire puissance maximale si surdimensionnée
   - Ajuster via paramètres carte électronique

6. **Optimisation combustion**
   - Ajuster air/gaz pour CO₂ optimal
   - Vérifier absence excès d'air (< 40%)
   - Contrôler température fumées après optimisation

7. **Contrôle isolation**
   - Vérifier l'isolation du corps de chauffe
   - Contrôler l'étanchéité portes et capots
   - Améliorer isolation si pertes importantes

**Prévention :**
- Nettoyage annuel échangeur
- Analyse combustion annuelle
- Traitement eau selon dureté
- Détartrage périodique si eau dure
- Vérification circulation hydraulique

**Valeurs de référence température fumées :**
- Chaudière standard : 120-180°C
- Chaudière basse température : 100-140°C
- Chaudière condensation : 50-90°C
- Si > valeurs : problème échange thermique

---

## FACT-CHAUD-019: Condensation anormale chaudière non-condensation

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-019 |
| **Catégorie** | Brûleur & Combustion |
| **Système** | Échangeur / Fumées |
| **Gravité** | **Élevée** |
| **Marques** | Multi-marques (chaudières standard/basse température) |

**Symptômes :**
- Présence d'eau/condensation dans la chaudière
- Corrosion échangeur
- Gouttes d'eau dans conduit fumées
- Taches de rouille
- Perforation échangeur possible

**Cause racine probable :**
Température retour trop basse (< 45-50°C), régulation mal paramétrée, surdimensionnement chaudière, réseau déséquilibré, by-pass insuffisant.

**Étapes de résolution :**

1. **Diagnostic situation**
   - Mesurer température retour chaudière
   - Si < 50°C : risque condensation important
   - Observer présence condensats dans chaudière
   - Évaluer l'état de corrosion échangeur

2. **Comprendre le phénomène**
   - Fumées de combustion gaz : ~120-180°C
   - Point de rosée fumées : ~55-60°C
   - Si paroi échangeur < 55°C : condensation acide (pH 3-4)
   - Acide = corrosion rapide échangeur fonte ou acier

3. **Contrôle régulation**
   - Vérifier température minimale départ (réglage)
   - Augmenter consigne minimale à 60-65°C minimum
   - Vérifier courbe de chauffe (pas trop basse)
   - Désactiver mode basse température si présent

4. **Contrôle installation hydraulique**
   - Vérifier présence by-pass anti-condensation
   - Si absent : installation obligatoire (vanne 3 voies ou kit)
   - Régler température retour > 50°C
   - Équilibrer le réseau (éviter retour trop froid)

5. **Solutions techniques**
   - **Vanne 3 voies thermostatique** : mélange départ/retour
   - Température ouverture : 55°C
   - Assure protection chaudière au démarrage
   - **Ballon tampon** : sur installations plancher chauffant
   - Découple température chaudière et réseau
   - Permet fonctionnement chaudière température élevée

6. **Vérification dimensionnement**
   - Calculer puissance réelle nécessaire
   - Si chaudière surdimensionnée (> 150%) : cycles courts, retour froid
   - Envisager remplacement par chaudière adaptée
   - Ou installer régulation climatique

7. **Contrôle dommages**
   - Inspecter échangeur (corrosion, perforation)
   - Si perforation : fuite eau/fumées, remplacement obligatoire
   - Si corrosion importante : anticiper remplacement
   - Documenter l'état (photos)

**Prévention :**
- Température retour chaudière > 50°C obligatoire
- Installation by-pass si température réseau basse
- Régulation adaptée au type émetteurs
- Dimensionnement correct puissance chaudière

**⚠️ IMPORTANT :**
- La condensation sur chaudière non-condensation = DÉFAUT GRAVE
- Destruction rapide échangeur (quelques mois)
- Fumées acides peuvent perforer échangeur
- Ne jamais faire condenser une chaudière standard

**Type émetteurs et risque :**
- Radiateurs haute température : risque faible
- Radiateurs basse température : risque moyen
- Plancher chauffant : risque très élevé (by-pass obligatoire)

---

## FACT-CHAUD-020: Variation puissance incohérente

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-020 |
| **Catégorie** | Brûleur & Combustion |
| **Système** | Régulation puissance |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques (chaudières modulantes) |

**Symptômes :**
- Puissance affichée ne correspond pas au comportement
- Flamme visuelle ne correspond pas à % affiché
- Brûleur en puissance max alors que faible demande
- Température départ dépasse consigne largement
- Ou inversement : puissance insuffisante

**Cause racine probable :**
Sonde température défectueuse (valeur erronée), carte électronique défaillante, paramètres mal configurés, vanne gaz défectueuse, capteur pression différentielle défectueux.

**Étapes de résolution :**

1. **Diagnostic comportement**
   - Observer puissance affichée menu diagnostic
   - Comparer avec flamme réelle (hauteur, bruit)
   - Mesurer température départ réelle vs consigne
   - Noter écart et incohérences

2. **Contrôle sondes température**
   - Mesurer résistance sonde départ (CTN/NTC)
   - Comparer à courbe constructeur (température/résistance)
   - Exemple : 10 kΩ à 25°C, 3,3 kΩ à 50°C (selon modèle)
   - Remplacer sonde si valeur incohérente

3. **Contrôle vanne gaz modulante**
   - Mesurer signal commande vanne (tension 0-10V ou PWM)
   - Vérifier que signal varie selon puissance demandée
   - Contrôler servomoteur vanne (bruit, mouvement)
   - Mesurer pression gaz brûleur (doit varier avec puissance)

4. **Contrôle capteur pression différentielle**
   - Sur certaines chaudières : capteur ΔP pour calcul débit
   - Vérifier valeur capteur cohérente
   - Nettoyer prises de pression si bouchées
   - Remplacer capteur si défectueux

5. **Contrôle carte électronique**
   - Vérifier paramètres puissance (mini/maxi)
   - Régler puissance max selon installation
   - Vérifier algorithme modulation (PID)
   - Réinitialiser paramètres usine
   - Mettre à jour firmware si disponible

6. **Test mode manuel**
   - Accéder mode test (forçage puissance)
   - Imposer différents % de puissance
   - Vérifier cohérence flamme/puissance imposée
   - Identifier défaut (capteur, vanne, carte)

7. **Optimisation paramètres**
   - Ajuster paramètres PID régulation
   - Coefficient proportionnel (P)
   - Temps intégration (I)
   - Temps dérivation (D)
   - Demande expertise si nécessaire

**Prévention :**
- Vérification annuelle sondes température
- Contrôle modulation en mode diagnostic
- Mise à jour firmware régulière
- Test cohérence puissance/flamme

**Menus diagnostic utiles (exemples) :**
- Vaillant : d.40 (vitesse ventilateur), d.41 (% puissance)
- Saunier Duval : P01 (puissance demandée)
- Frisquet : températures et puissance via boîtier Vision
- De Dietrich : menu Service → Diagnostic

---

*Fin du fichier 01_Combustion_Bruleur.md*

**Retour à l'index :** [Knowledge_Base_Chaudieres.md](Knowledge_Base_Chaudieres.md)
