# Problèmes Eau Chaude Sanitaire (ECS)

**Catégorie :** Eau Chaude Sanitaire (ECS)
**Nombre de Facts :** 15
**Retour à l'index :** [Knowledge_Base_Chaudieres.md](Knowledge_Base_Chaudieres.md)

---

## FACT-CHAUD-051: Pas d'eau chaude sanitaire

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-051 |
| **Catégorie** | Eau Chaude Sanitaire (ECS) |
| **Système** | Production ECS |
| **Gravité** | **Élevée** |
| **Marques** | Multi-marques |

**Symptômes :**
- Aucune production d'eau chaude au robinet
- Eau reste froide même après plusieurs minutes
- Chaudière ne démarre pas sur demande ECS
- Mode chauffage fonctionne normalement
- Pas de bruit brûleur lors ouverture robinet ECS

**Cause racine probable :**
Vanne 3 voies bloquée en position chauffage, débitmètre ECS défectueux ou encrassé, carte électronique ne détecte pas demande ECS, brûleur ne s'allume pas sur demande sanitaire, ballon ECS non alimenté.

**Étapes de résolution :**

1. **Diagnostic initial**
   - Ouvrir robinet eau chaude et maintenir ouvert
   - Écouter si brûleur démarre (tentative allumage)
   - Vérifier température départ chauffage (évolution)
   - Consulter afficheur chaudière (mode ECS actif ?)

2. **Contrôle débitmètre ECS**
   - Localiser le débitmètre (généralement sur circuit sanitaire)
   - Démonter et vérifier rotation turbine
   - Nettoyer si encrassement (calcaire, impuretés)
   - Vérifier aimant turbine (présent, non cassé)
   - Tester fonctionnement : souffler dedans → turbine doit tourner
   - Remplacer si défectueux

3. **Contrôle vanne 3 voies sanitaire**
   - Localiser vanne 3 voies (corps chaudière)
   - Écouter commutation au démarrage ECS (clic moteur)
   - Vérifier alimentation électrique moteur vanne
   - Tester commutation manuelle si possible
   - Démonter et débloquer si nécessaire
   - Remplacer si moteur grillé

4. **Contrôle carte électronique**
   - Vérifier détection débit ECS (menu diagnostic)
   - Contrôler signal débitmètre (impulsions)
   - Vérifier commande vanne 3 voies (tension sortie)
   - Tester mode forcé ECS si disponible
   - Consulter codes erreur éventuels

5. **Contrôle échangeur sanitaire**
   - Sur chaudière instantanée : vérifier échangeur à plaques
   - Vérifier circulation eau dans échangeur
   - Contrôler absence obstruction totale (calcaire)
   - Mesurer pression eau froide entrée (2-3 bars mini)

6. **Contrôle ballon ECS (si présent)**
   - Vérifier alimentation électrique ballon
   - Contrôler thermostat ballon
   - Vérifier sonde température ballon
   - Tester résistance électrique ballon si présent
   - Contrôler vanne motorisée ballon

7. **Tests et remise en service**
   - Purger l'air circuit sanitaire
   - Ouvrir robinet ECS et maintenir
   - Vérifier démarrage brûleur
   - Contrôler montée en température
   - Mesurer température ECS au robinet (45-60°C)

**Prévention :**
- Détartrage préventif échangeur sanitaire annuel (eau dure)
- Vérification fonctionnement vanne 3 voies
- Nettoyage débitmètre à chaque entretien
- Installation adoucisseur si eau très dure (> 25°F)
- Test cycle ECS lors entretien annuel

---

## FACT-CHAUD-052: ECS tiède ou insuffisamment chaude

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-052 |
| **Catégorie** | Eau Chaude Sanitaire (ECS) |
| **Système** | Température ECS |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques |

**Symptômes :**
- Eau chaude produite mais température insuffisante
- Température ECS < 40°C au robinet
- Nécessité ouvrir en grand côté chaud pour avoir tiède
- Température diminue si débit élevé
- Ballon ECS ne monte pas assez en température

**Cause racine probable :**
Consigne température ECS trop basse, échangeur sanitaire entartré (perte rendement), puissance chaudière insuffisante pour débit demandé, vanne 3 voies en position intermédiaire, mélangeur thermostatique mal réglé, ballon sous-dimensionné.

**Étapes de résolution :**

1. **Vérification consigne température**
   - Consulter température consigne ECS (afficheur)
   - Consigne normale : 50-60°C
   - Augmenter consigne progressivement (+5°C)
   - Tester après stabilisation (10-15 min)

2. **Mesure températures**
   - Mesurer température départ chaudière en mode ECS
   - Doit être 60-65°C minimum
   - Mesurer température au robinet ECS (après 2 min écoulement)
   - Calculer perte thermique réseau (départ - robinet)
   - Si > 10°C : problème circuit ou échangeur

3. **Contrôle échangeur sanitaire**
   - Échangeur à plaques : suspecter entartrage
   - Démonter et inspecter visuellement
   - Plaques entartrées = passage réduit, échange dégradé
   - Détartrer chimiquement (acide citrique ou produit spécifique)
   - Rincer abondamment avant remontage
   - Remplacer échangeur si trop dégradé

4. **Contrôle vanne 3 voies**
   - Vérifier position vanne en mode ECS
   - Doit être complètement basculée sur sanitaire
   - Si position intermédiaire : débit insuffisant
   - Ajuster ou remplacer vanne

5. **Contrôle puissance et modulation**
   - Vérifier puissance chaudière en mode ECS (menu diagnostic)
   - Doit être proche puissance maximale
   - Si puissance bridée : vérifier paramètres
   - Augmenter puissance max ECS si paramétrable

6. **Contrôle débit ECS**
   - Mesurer débit au robinet (litres/minute)
   - Débit excessif : puissance chaudière insuffisante
   - Calculer : puissance nécessaire = débit × ΔT × 1,16
   - Exemple : 10 L/min, ΔT 40°C → 27 kW nécessaires
   - Limiter débit si puissance insuffisante (réducteur pression)

7. **Contrôle ballon ECS (si présent)**
   - Vérifier température ballon (sonde)
   - Contrôler temps de chauffe ballon
   - Vérifier isolation ballon (pertes thermiques)
   - Augmenter consigne température ballon
   - Vérifier stratification (eau chaude en haut)

8. **Contrôle mitigeur thermostatique**
   - Si mitigeur installé : vérifier réglage
   - Mitigeur bloqué en position tiède
   - Tester en retirant mitigeur temporairement
   - Détartrer ou remplacer mitigeur

**Prévention :**
- Détartrage annuel échangeur sanitaire (eau dure)
- Réglage température ECS 55-60°C
- Installation adoucisseur (dureté > 25°F)
- Vérification annuelle vanne 3 voies
- Dimensionnement correct puissance chaudière
- Isolation tuyauteries ECS

**Note :** Température ECS recommandée : 50-60°C (confort + anti-légionelle). Température < 50°C : risque légionellose.

---

## FACT-CHAUD-053: Température ECS instable

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-053 |
| **Catégorie** | Eau Chaude Sanitaire (ECS) |
| **Système** | Régulation ECS |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques (surtout instantané) |

**Symptômes :**
- Température ECS varie brutalement (chaud/froid)
- Eau brûlante puis tiède en alternance
- Variations synchrones avec débit
- Inconfort important à la douche
- Température stable puis chutes brutales

**Cause racine probable :**
Sonde température ECS défectueuse, régulation PID mal paramétrée, débitmètre encrassé (signal erratique), vanne 3 voies instable, pression eau fluctuante, échangeur partiellement entartré.

**Étapes de résolution :**

1. **Caractérisation instabilité**
   - Observer le cycle : périodicité variations
   - Variations liées débit (ouvrir/fermer robinet)
   - Variations liées température chauffage
   - Consulter température affichée vs ressentie

2. **Contrôle sonde température ECS**
   - Localiser sonde ECS (sortie échangeur sanitaire)
   - Mesurer résistance sonde (CTN/NTC)
   - Comparer à courbe constructeur
   - Vérifier contact thermique sonde (doigt de gant)
   - Remplacer si valeurs erratiques

3. **Contrôle débitmètre**
   - Démonter et nettoyer débitmètre
   - Vérifier rotation fluide turbine
   - Tester signal électrique (oscilloscope si possible)
   - Signal doit être proportionnel au débit
   - Remplacer si signal instable

4. **Optimisation régulation ECS**
   - Accéder paramètres régulation ECS
   - Ajuster paramètres PID (Proportionnel, Intégral, Dérivé)
   - Augmenter temps intégration (stabilité)
   - Réduire gain proportionnel (éviter oscillations)
   - Tester et ajuster progressivement

5. **Contrôle pression eau**
   - Mesurer pression eau froide sanitaire
   - Doit être stable 2-3 bars
   - Si fluctuations importantes : problème réseau
   - Installer réducteur pression si nécessaire
   - Vérifier vase expansion sanitaire (si présent)

6. **Contrôle vanne 3 voies**
   - Vérifier stabilité position vanne en mode ECS
   - Écouter vibrations ou bruits anormaux
   - Lubrifier mécanisme si prévu constructeur
   - Remplacer si oscillations mécaniques

7. **Contrôle échangeur sanitaire**
   - Entartrage partiel : passages bouchés irrégulièrement
   - Crée turbulences et variations température
   - Détartrer chimiquement
   - Vérifier homogénéité après détartrage

8. **Solutions complémentaires**
   - Installer ballon tampon sanitaire (stabilisation)
   - Installer mitigeur thermostatique (sécurité utilisateur)
   - Régler température ECS légèrement supérieure
   - Installer régulateur débit (stabilisation)

**Prévention :**
- Vérification annuelle sonde ECS
- Nettoyage débitmètre
- Contrôle pression eau sanitaire
- Détartrage préventif échangeur
- Réglage régulation optimisé
- Installation mitigeur thermostatique recommandée

**Note :** Mitigeur thermostatique = sécurité anti-brûlure + confort. Obligatoire dans certains ERP (établissements recevant du public).

---

## FACT-CHAUD-054: Débit ECS faible

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-054 |
| **Catégorie** | Eau Chaude Sanitaire (ECS) |
| **Système** | Hydraulique ECS |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques |

**Symptômes :**
- Débit eau chaude très faible au robinet
- Eau froide : débit normal / Eau chaude : filet
- Débit diminue progressivement
- Débit faible sur tous les points de puisage
- Parfois sifflements dans tuyauterie

**Cause racine probable :**
Échangeur sanitaire entartré (obstruction passages), filtre ECS bouché, clapet anti-retour bloqué, réducteur de pression mal réglé, vanne 3 voies partiellement ouverte, tuyauterie obstruée (calcaire).

**Étapes de résolution :**

1. **Localisation problème**
   - Tester débit sur plusieurs robinets
   - Si tous robinets : problème chaudière/réseau
   - Si un seul robinet : problème local (brise-jet, mitigeur)
   - Comparer débit eau froide vs eau chaude

2. **Contrôle pression eau**
   - Mesurer pression eau froide sanitaire (manomètre)
   - Pression normale : 2-4 bars
   - Si < 2 bars : pression insuffisante (réducteur, réseau)
   - Ajuster réducteur pression si présent

3. **Contrôle filtres**
   - Localiser filtre entrée eau froide chaudière
   - Démonter et nettoyer (grillage, cartouche)
   - Vérifier filtre amont compteur (si accessible)
   - Vérifier filtres robinets (mousseurs, brise-jets)

4. **Contrôle échangeur sanitaire**
   - Échangeur à plaques : très sensible entartrage
   - Passages étroits se bouchent rapidement
   - Démonter échangeur
   - Inspection visuelle : plaques entartrées
   - Détartrage chimique obligatoire :
     - Solution acide citrique 15% ou produit spécifique
     - Faire circuler en boucle 2-4 heures
     - Rincer abondamment à l'eau claire
   - Remplacer échangeur si obstruction totale

5. **Contrôle vanne 3 voies**
   - Vérifier ouverture complète en mode ECS
   - Vanne partiellement ouverte : débit réduit
   - Démonter et contrôler mécanisme
   - Détartrer ou remplacer

6. **Contrôle clapet anti-retour**
   - Localiser clapet (circuit sanitaire)
   - Clapet bloqué ou entartré : perte charge importante
   - Démonter, nettoyer ou remplacer

7. **Contrôle tuyauterie ECS**
   - Sur installation ancienne : entartrage tuyauteries
   - Réduction progressive section (tartre)
   - Vérifier diamètre tuyauterie (mini 12 mm)
   - Si entartrage général : remplacement tuyauteries nécessaire

8. **Mesures curatives**
   - Installer adoucisseur eau si dureté > 25°F
   - Installer filtre anti-tartre magnétique/polyphosphates
   - Programmer détartrage préventif annuel
   - Régler température ECS < 60°C (limite entartrage)

**Prévention :**
- Détartrage annuel échangeur sanitaire (eau dure)
- Installation adoucisseur (dureté > 25°F)
- Température ECS 55°C (compromis confort/entartrage)
- Nettoyage filtres régulier
- Traitement anti-tartre
- Éviter stagnation eau (purges régulières)

**Dureté eau et risque entartrage :**
- < 15°F : eau douce, peu de risque
- 15-25°F : eau moyennement dure, détartrage annuel
- 25-40°F : eau dure, adoucisseur recommandé
- > 40°F : eau très dure, adoucisseur obligatoire

---

## FACT-CHAUD-055: Temps de chauffe ECS trop long

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-055 |
| **Catégorie** | Eau Chaude Sanitaire (ECS) |
| **Système** | Performance ECS |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques |

**Symptômes :**
- Attente prolongée avant eau chaude au robinet (> 30 secondes)
- Gaspillage eau en attendant température
- Temps d'attente augmente progressivement
- Problème pire sur points éloignés chaudière
- Ballon met très longtemps à chauffer

**Cause racine probable :**
Distance chaudière/robinet importante (volume tuyauterie), échangeur entartré (puissance réduite), tuyauteries non calorifugées (pertes thermiques), puissance chaudière insuffisante, ballon ECS sous-dimensionné ou défectueux.

**Étapes de résolution :**

1. **Diagnostic problème**
   - Mesurer temps attente eau chaude (chronomètre)
   - À 10 cm robinet : doit être < 30 secondes
   - Calculer volume tuyauterie à purger
   - Exemple : 10 m en Ø16 mm = 2 litres
   - Évaluer si temps cohérent avec volume

2. **Contrôle chaudière instantanée**
   - Vérifier puissance démarrage ECS (menu diagnostic)
   - Mesurer temps montée température départ (0→60°C)
   - Doit être < 10 secondes
   - Si lent : échangeur entartré ou puissance insuffisante
   - Détartrer échangeur sanitaire

3. **Contrôle ballon ECS**
   - Mesurer temps chauffe ballon (froid → 60°C)
   - Ballon 150L : 2-3 heures maximum
   - Si > 4 heures : problème puissance ou échangeur
   - Vérifier sonde température ballon (détection)
   - Contrôler isolation ballon (pertes thermiques)
   - Vérifier résistance électrique ballon (si présente)

4. **Optimisation installation**
   - **Solution 1 : Calorifugeage tuyauteries**
     - Isoler tous tuyaux ECS (mousse, manchons)
     - Épaisseur isolant : 10-20 mm minimum
     - Réduit refroidissement entre puisages
   - **Solution 2 : Boucle de recyclage**
     - Sur installation complexe : boucle bouclée + circulateur
     - Eau chaude en permanence dans réseau
     - Confort immédiat mais surconsommation
   - **Solution 3 : Production décentralisée**
     - Installer chauffe-eau locaux (points éloignés)
     - Ballon électrique ou chauffe-eau instantané gaz

5. **Contrôle échangeur et puissance**
   - Détartrer échangeur sanitaire
   - Vérifier puissance chaudière adaptée
   - Puissance nécessaire (instantané) : débit × ΔT × 1,16
   - Exemple : 12 L/min, ΔT 40°C → 28 kW
   - Si puissance insuffisante : bridage ou sous-dimensionnement

6. **Solutions techniques**
   - Installer ballon tampon sanitaire (petite capacité)
   - Ballon 10-20L proche robinets : réserve immédiate
   - Installer système pré-chauffage (thermosiphon solaire)
   - Optimiser tracé tuyauteries (réduire longueur)

7. **Réglage chaudière**
   - Augmenter température consigne ECS
   - Activer mode "Confort" si disponible (maintien température)
   - Certaines chaudières : maintien T° en permanence
   - Attention surconsommation

**Prévention :**
- Calorifugeage obligatoire tuyauteries ECS
- Détartrage régulier échangeur
- Dimensionnement correct installation
- Limitation longueur tuyauteries (< 10 m si possible)
- Ballon tampon si distance importante

**Réglementation :**
- RT 2012 / RE 2020 : calorifugeage obligatoire
- Pertes réseau limitées (distance, isolation)
- Objectif : < 3 litres purge avant eau chaude

**Gain calorifugeage :**
- Réduction pertes thermiques : 70-80%
- Économie énergie : 5-10% sur ECS
- Confort amélioré (eau chaude plus rapide)

---

## FACT-CHAUD-056: Ballon ECS qui ne chauffe pas

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-056 |
| **Catégorie** | Eau Chaude Sanitaire (ECS) |
| **Système** | Ballon ECS |
| **Gravité** | **Élevée** |
| **Marques** | Multi-marques (chaudière + ballon) |

**Symptômes :**
- Ballon ECS reste froid
- Pas de production eau chaude (stock ballon épuisé)
- Chaudière ne chauffe pas le ballon
- Température ballon stagne (affichage)
- Mode chauffage fonctionne normalement

**Cause racine probable :**
Vanne motorisée ballon défectueuse, circulateur ballon ne démarre pas, sonde température ballon HS, programmation horaire incorrecte, aquastat ballon défaillant, échangeur ballon entartré.

**Étapes de résolution :**

1. **Diagnostic système**
   - Identifier type ballon :
     - Ballon avec serpentin (échangeur interne)
     - Ballon avec échangeur externe (à plaques)
     - Ballon électrique + appoint chaudière
   - Vérifier programmation (plages horaires chauffe)
   - Consulter afficheur chaudière (demande ballon active ?)

2. **Contrôle sonde température ballon**
   - Localiser sonde (doigt de gant sur ballon)
   - Mesurer résistance sonde (multimètre)
   - Comparer à courbe constructeur
   - Sonde défectueuse = chaudière pense ballon chaud
   - Remplacer sonde si défectueuse

3. **Contrôle vanne motorisée ballon**
   - Localiser vanne 3/4 voies ballon
   - Vérifier alimentation électrique moteur vanne
   - Écouter commutation vanne (demande ballon)
   - Tester position vanne manuellement
   - Vérifier câblage et contacts
   - Remplacer vanne si moteur grillé

4. **Contrôle circulateur ballon**
   - Vérifier fonctionnement circulateur ballon
   - Circulateur doit tourner en mode charge ballon
   - Contrôler alimentation électrique
   - Débloquer rotor si grippé (vis déblocage)
   - Vérifier vitesse circulateur (adaptée)
   - Remplacer circulateur si défectueux

5. **Contrôle aquastat ballon**
   - Si ballon ancien : aquastat mécanique
   - Tester continuité contacts
   - Ajuster température consigne aquastat
   - Remplacer si défectueux

6. **Contrôle hydraulique ballon**
   - Vérifier circulation eau primaire dans échangeur
   - Purger air circuit primaire ballon
   - Vérifier absence obstruction (vannes fermées)
   - Contrôler pression circuit primaire
   - Détartrer échangeur ballon si nécessaire

7. **Contrôle programmation**
   - Vérifier plages horaires chauffe ballon
   - Mode absence/vacances désactivé
   - Température consigne ballon (55-60°C)
   - Forcer charge ballon (mode manuel)

8. **Contrôle résistance électrique (si présente)**
   - Sur ballon mixte (gaz + électrique)
   - Vérifier alimentation résistance
   - Tester continuité résistance (multimètre)
   - Contrôler thermostat sécurité (réarmable)
   - Détartrer résistance si entartrée
   - Remplacer résistance si coupée

9. **Contrôle priorité ECS**
   - Vérifier paramètre priorité ECS activé
   - En mode chauffe ballon : arrêt chauffage normal
   - Si chauffage + ballon simultané : vanne défectueuse

**Prévention :**
- Vérification annuelle circuit ballon
- Test charge ballon lors entretien
- Contrôle sonde et aquastat
- Détartrage ballon selon dureté eau (tous les 2-5 ans)
- Vérification circulateur et vanne motorisée

**Dimensionnement ballon :**
- 1 personne : 50-75 litres
- 2 personnes : 100-150 litres
- 3 personnes : 150-200 litres
- 4 personnes : 200-300 litres
- Température stockage : 55-60°C

---

## FACT-CHAUD-057: Échangeur sanitaire entartré

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-057 |
| **Catégorie** | Eau Chaude Sanitaire (ECS) |
| **Système** | Échangeur sanitaire |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques (surtout instantané) |

**Symptômes :**
- Débit ECS progressivement réduit
- Température ECS insuffisante ou instable
- Bruit anormal (sifflements) circuit sanitaire
- Chaudière monte en température (surchauffe)
- Performances dégradées progressivement

**Cause racine probable :**
Eau dure (forte teneur calcaire), température ECS élevée (> 60°C favorise dépôt), stagnation eau dans échangeur, absence traitement eau, vieillissement installation.

**Étapes de résolution :**

1. **Diagnostic entartrage**
   - Mesurer dureté eau (TH en °F)
   - Analyser historique : performances initiales vs actuelles
   - Vérifier débit ECS (comparaison avant/après)
   - Consulter carnet entretien (dernier détartrage)

2. **Démontage échangeur à plaques**
   - Fermer vannes isolement entrée/sortie
   - Vidanger échangeur
   - Démonter échangeur (boulons, connexions)
   - Photographier état avant nettoyage
   - Inspecter plaques (dépôts blancs/gris = calcaire)

3. **Détartrage chimique échangeur**
   - **Méthode 1 : Trempage**
     - Préparer bac avec solution détartrante
     - Acide citrique 15% ou vinaigre blanc 14° (économique)
     - Ou produit détartrant spécifique chaudière
     - Immerger plaques 2-4 heures
     - Brosser doucement (brosse douce, non métallique)
     - Rincer abondamment eau claire
   - **Méthode 2 : Circulation (échangeur en place)**
     - Utiliser pompe détartrage
     - Faire circuler solution détartrante en circuit fermé
     - Température 40-50°C (efficacité accrue)
     - Durée : 2-4 heures selon encrassement
     - Rincer en circulation (eau claire jusqu'à pH neutre)

4. **Produits détartrants**
   - Acide citrique : écologique, efficace, pas dangereux
   - Acide chlorhydrique dilué : très efficace mais corrosif (précautions)
   - Produits spécifiques chaudière (ex: Sentinel X400, Cillit)
   - Respecter dosages fabricant
   - Porter EPI : gants, lunettes, protection

5. **Contrôle état plaques**
   - Après détartrage : inspecter plaques
   - Vérifier absence perforation, corrosion
   - Joints : remplacer systématiquement
   - Si plaques très dégradées : remplacement échangeur

6. **Remontage échangeur**
   - Nettoyer surfaces contact
   - Installer joints neufs (référence constructeur)
   - Respecter ordre et sens plaques (marquage)
   - Serrer boulons progressivement en croix (couple)
   - Vérifier étanchéité avant remise en service

7. **Test après détartrage**
   - Remplir circuit, purger air
   - Ouvrir vannes progressivement
   - Vérifier absence fuite
   - Tester débit ECS (avant/après)
   - Mesurer température ECS (performances)
   - Gain attendu : +30 à 100% débit/puissance

8. **Prévention entartrage**
   - **Traitement eau :**
     - Adoucisseur si TH > 25°F (recommandé > 20°F)
     - Filtre polyphosphates (petites installations)
     - Système anti-tartre magnétique/électronique
   - **Réglages :**
     - Température ECS 55°C (limite entartrage vs confort)
     - Éviter surchauffe (> 60°C favorise dépôt)
   - **Entretien :**
     - Détartrage préventif annuel (eau dure)
     - Purges régulières (éviter stagnation)

**Prévention :**
- Installation adoucisseur (TH > 25°F)
- Réglage température ECS ≤ 60°C
- Détartrage préventif selon dureté :
  - TH < 15°F : tous les 4-5 ans
  - TH 15-25°F : tous les 2-3 ans
  - TH > 25°F : annuel
- Traitement anti-tartre

**Dureté eau (TH) :**
- Très douce : < 8°F
- Douce : 8-15°F
- Moyennement dure : 15-25°F
- Dure : 25-40°F
- Très dure : > 40°F

---

## FACT-CHAUD-058: Vanne 3 voies sanitaire bloquée

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-058 |
| **Catégorie** | Eau Chaude Sanitaire (ECS) |
| **Système** | Vanne 3 voies |
| **Gravité** | **Élevée** |
| **Marques** | Multi-marques |

**Symptômes :**
- Pas d'eau chaude sanitaire (vanne bloquée chauffage)
- Ou pas de chauffage (vanne bloquée sanitaire)
- Bruit moteur vanne mais pas de mouvement
- Température départ ne change pas entre modes
- Chaudière chauffe en continu sans effet

**Cause racine probable :**
Vanne bloquée mécaniquement (calcaire, boues), moteur vanne grillé, condensateur moteur défectueux (si moteur 230V), came ou engrenage cassé, vanne mal positionnée ou montée incorrectement.

**Étapes de résolution :**

1. **Diagnostic blocage**
   - Provoquer demande ECS (ouvrir robinet)
   - Écouter moteur vanne (bruit commutation)
   - Toucher moteur (vibration = tentative mouvement)
   - Observer tige vanne (mouvement visible ?)
   - Identifier type vanne :
     - Vanne 3 voies chauffage/ECS
     - Vanne 4 voies (avec ballon)

2. **Test moteur vanne**
   - Vérifier alimentation électrique moteur (multimètre)
   - Tension normale : 24V ou 230V selon modèle
   - Si pas de tension : problème carte électronique
   - Si tension OK mais pas mouvement : moteur ou mécanique

3. **Déblocage manuel vanne**
   - Localiser levier déblocage manuel (sur moteur)
   - Débrayer moteur (vis ou levier)
   - Actionner vanne manuellement
   - Si très dur : blocage mécanique (tartre, corrosion)
   - Si fluide : problème moteur

4. **Démontage et nettoyage vanne**
   - **Préparation :**
     - Vidanger chaudière (ou isoler vanne si possible)
     - Déconnecter moteur (câblage)
     - Démonter moteur vanne
   - **Nettoyage :**
     - Nettoyer tige vanne (graisse, calcaire)
     - Actionner vanne manuellement (lubrifier)
     - Nettoyer came et engrenages moteur
     - Éliminer boues et dépôts
     - Vérifier absence casse (came, tige)
   - **Si blocage interne :**
     - Démonter corps vanne (selon modèle)
     - Détartrer mécanisme interne
     - Remplacer joints

5. **Contrôle moteur vanne**
   - Tester moteur hors installation (alimentation directe)
   - Si moteur 230V : vérifier condensateur
     - Condensateur gonflé ou défectueux : remplacement
     - Condensateur 2-4 µF typiquement
   - Si moteur ne tourne pas : remplacement moteur
   - Couples moteur/vanne spécifiques selon marque

6. **Contrôle came et positionnement**
   - Vérifier position came (repères moteur/vanne)
   - Came mal positionnée : vanne pas en bonne position
   - Régler position mécanique (repères constructeur)
   - Position repos : généralement chauffage (hiver)

7. **Remontage et test**
   - Lubrifier mécanisme (graisse adaptée haute température)
   - Remonter moteur avec repérage correct
   - Reconnecter câblage (respecter polarité si DC)
   - Tester commutation :
     - Mode chauffage → mode ECS → mode chauffage
     - Vérifier mouvement complet vanne
     - Écouter fin de course (arrêt moteur)

8. **Calibration (si nécessaire)**
   - Certaines vannes : procédure calibration auto
   - Consulter notice constructeur
   - Forcer plusieurs cycles complets
   - Vérifier positions extrêmes

**Prévention :**
- Exercice vanne mensuel (cycle chauffage/ECS)
- Vérification annuelle lors entretien
- Traitement eau (limite tartre/boues)
- Remplacement préventif moteur (durée vie ~10 ans)
- Lubrification mécanisme selon fabricant

**Pannes fréquentes par marque :**
- Saunier Duval : moteur Honeywell VC4012/VC4013
- Vaillant : moteur VR 30 (24V)
- Frisquet : vanne spécifique (difficile à trouver)
- De Dietrich : vanne Siemens VDN

**Remplacement vanne complète :**
- Si casse mécanique : vanne complète
- Coût : 100-300€ selon modèle
- Prévoir joints et graisse

---

## FACT-CHAUD-059: Débitmètre ECS défectueux

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-059 |
| **Catégorie** | Eau Chaude Sanitaire (ECS) |
| **Système** | Détection débit ECS |
| **Gravité** | **Élevée** |
| **Marques** | Multi-marques (chaudières instantanées) |

**Symptômes :**
- Chaudière ne détecte pas demande ECS
- Pas de démarrage brûleur à l'ouverture robinet
- Démarrage aléatoire ou tardif
- Code erreur débit ECS insuffisant
- Brûleur s'arrête en cours de puisage

**Cause racine probable :**
Turbine débitmètre bloquée (calcaire, impuretés), aimant turbine manquant ou cassé, capteur Hall défectueux, câblage débitmètre coupé, corps débitmètre fissuré (fuite).

**Étapes de résolution :**

1. **Compréhension fonctionnement**
   - Débitmètre = turbine + aimant + capteur Hall
   - Eau fait tourner turbine
   - Aimant génère impulsions électriques (capteur Hall)
   - Carte électronique compte impulsions → calcul débit
   - Seuil déclenchement : 2-4 L/min selon modèles

2. **Test fonctionnement débitmètre**
   - Ouvrir robinet ECS progressivement
   - Observer afficheur chaudière :
     - Symbole robinet qui s'affiche = détection OK
     - Ou affichage débit (L/min) si disponible
   - Si aucune détection : débitmètre HS

3. **Démontage débitmètre**
   - Fermer vanne entrée eau froide chaudière
   - Purger pression (ouvrir robinet ECS)
   - Déconnecter câble électrique débitmètre
   - Dévisser débitmètre (clé, selon modèle)
   - Attention : eau résiduelle (prévoir récipient)

4. **Inspection visuelle**
   - Extraire turbine du corps débitmètre
   - Vérifier rotation turbine (souffler ou tourner doigt)
   - Turbine bloquée : calcaire, impuretés, cheveu
   - Vérifier présence aimant sur turbine
   - Aimant cassé/absent : turbine inefficace

5. **Nettoyage débitmètre**
   - Démonter turbine complètement
   - Nettoyer corps débitmètre (brosse, eau)
   - Détartrer si calcaire (vinaigre blanc)
   - Rincer abondamment
   - Nettoyer turbine (aimant, pales)
   - Vérifier axe turbine (pas de jeu excessif)
   - Remonter turbine : doit tourner librement

6. **Test électrique capteur**
   - Mesurer résistance capteur Hall (si accessible)
   - Valeur typique : 500-1500 Ω (selon modèle)
   - Ou tester signal électrique :
     - Reconnecter câble électrique
     - Souffler dans débitmètre (turbine tourne)
     - Mesurer impulsions (oscilloscope ou multimètre AC)
   - Signal absent : capteur HS → remplacement débitmètre

7. **Contrôle câblage**
   - Vérifier continuité câble débitmètre → carte
   - Contrôler connecteur (oxydation, mauvais contact)
   - Vérifier alimentation capteur (5-12V DC selon modèle)

8. **Remontage et test**
   - Remonter débitmètre avec joint neuf (obligatoire)
   - Serrer modérément (pas trop : risque casse plastique)
   - Reconnecter câble électrique
   - Ouvrir vanne eau froide progressivement
   - Vérifier absence fuite
   - Tester démarrage ECS (ouvrir robinet)
   - Vérifier affichage débit ou symbole

9. **Calibration débit (si paramétrable)**
   - Certaines chaudières : réglage seuil déclenchement
   - Ajuster selon installation (pression, tuyauterie)
   - Seuil trop élevé : pas de détection petits débits
   - Seuil trop bas : déclenchements intempestifs

**Prévention :**
- Nettoyage annuel débitmètre (eau dure)
- Installation filtre entrée eau froide (protection)
- Vérification test ECS lors entretien
- Traitement anti-tartre
- Remplacement préventif (durée vie ~10 ans)

**Modèles débitmètres courants :**
- Saunier Duval : 0020133280
- Vaillant : 0020039057
- Elm Leblanc : 87167808640
- Débitmètres souvent interchangeables (vérifier raccords)

**Coût remplacement :**
- Débitmètre seul : 30-80€
- Main d'œuvre : 1 heure

---

## FACT-CHAUD-060: Sonde ECS défectueuse

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-060 |
| **Catégorie** | Eau Chaude Sanitaire (ECS) |
| **Système** | Régulation température ECS |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques |

**Symptômes :**
- Température ECS incohérente (affichage vs réalité)
- Code erreur sonde ECS
- Température ECS instable ou excessive
- Chaudière ne régule pas température ECS
- Affichage température aberrante (ex: 120°C)

**Cause racine probable :**
Sonde CTN/NTC défectueuse (dérive, coupure), mauvais contact thermique sonde, câblage coupé ou court-circuit, connecteur oxydé, sonde mal positionnée (hors flux).

**Étapes de résolution :**

1. **Identification symptômes**
   - Code erreur sonde affiché (ex: Vaillant F75)
   - Température affichée aberrante
   - Température affichée fixe (ne varie pas)
   - ECS trop chaude ou trop froide

2. **Localisation sonde ECS**
   - Sonde généralement sur sortie échangeur sanitaire
   - Montage :
     - Doigt de gant (immergé dans flux)
     - Collier sur tuyauterie (contact externe)
     - Intégré échangeur (selon modèle)
   - Vérifier repérage (sonde ECS vs sonde chauffage)

3. **Test électrique sonde**
   - Déconnecter sonde (connecteur chaudière)
   - Mesurer résistance sonde (multimètre Ω)
   - **Sonde CTN/NTC typique :**
     - 10 kΩ à 25°C
     - 3,3 kΩ à 50°C
     - 1,5 kΩ à 70°C
   - Comparer à courbe constructeur (notice technique)
   - Valeurs aberrantes :
     - Infini (∞) : sonde coupée
     - 0 Ω : court-circuit
     - Valeur fixe qui ne varie pas : sonde HS

4. **Test variation résistance**
   - Chauffer sonde (eau chaude, sèche-cheveux)
   - Résistance doit diminuer avec température (CTN)
   - Variation normale : divisée par 2-3 tous les 20°C
   - Si pas de variation : sonde défectueuse

5. **Contrôle contact thermique**
   - Sonde doigt de gant : vérifier présence pâte thermique
   - Doigt de gant vide : mauvais contact → mesure erronée
   - Appliquer pâte thermique (conductivité thermique)
   - Sonde collier : vérifier serrage correct (contact métal/métal)

6. **Contrôle câblage**
   - Vérifier continuité câble sonde → carte
   - Contrôler isolation (résistance câble/masse > 1 MΩ)
   - Inspecter câble (coupure, dénudage, brûlure)
   - Vérifier connecteur (oxydation, humidité)
   - Nettoyer contacts (bombe contact électronique)

7. **Test avec sonde de rechange**
   - Si doute : tester avec sonde neuve
   - Ou permuter temporairement sonde chauffage/ECS (test)
   - Attention : sondes parfois différentes (courbes)

8. **Remplacement sonde**
   - Commander sonde référence constructeur
   - Ou sonde universelle (vérifier courbe compatible)
   - Vidanger circuit si nécessaire (doigt de gant)
   - Installer sonde neuve :
     - Doigt de gant : pâte thermique obligatoire
     - Collier : serrage correct, contact tuyau
   - Reconnecter câblage (respecter polarité si marquée)

9. **Vérification après remplacement**
   - Vérifier température affichée cohérente
   - Lancer production ECS
   - Observer régulation température
   - Vérifier température robinet (thermomètre)
   - Écart affiché/réel doit être < 5°C

**Prévention :**
- Vérification annuelle valeur sonde (menu diagnostic)
- Contrôle contact thermique
- Protection câblage (pas de frottement, chaleur excessive)
- Nettoyage connecteurs
- Remplacement préventif (durée vie ~15 ans)

**Codes erreur sonde ECS par marque :**
- Vaillant : F75 (défaut sonde ECS)
- Saunier Duval : F73 (sonde ECS court-circuit), F74 (sonde ECS coupée)
- Elm Leblanc : E03 (sonde ECS)
- De Dietrich : E30 (sonde sanitaire)
- Frisquet : 203 (sonde sanitaire)

**Courbe typique CTN 10kΩ (exemple) :**
- 0°C : 32,7 kΩ
- 25°C : 10 kΩ
- 50°C : 3,3 kΩ
- 80°C : 1,1 kΩ

---

## FACT-CHAUD-061: Légionellose - Cycle anti-légionelle

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-061 |
| **Catégorie** | Eau Chaude Sanitaire (ECS) |
| **Système** | Sécurité sanitaire |
| **Gravité** | **Critique** (santé publique) |
| **Marques** | Multi-marques (fonction anti-légionellose) |

**Symptômes :**
- Pas de symptôme visible sur chaudière
- Risque sanitaire (légionellose = pneumopathie grave)
- Température ECS < 50°C (zone à risque)
- Stagnation eau dans ballon ou réseau
- Eau tiède stockée longtemps

**Cause racine probable :**
Température ECS insuffisante (< 50°C), absence cycle anti-légionelle, stagnation eau (points morts, peu utilisés), ballon mal dimensionné (surdimensionné), fonction anti-légionelle désactivée.

**Étapes de résolution :**

1. **Compréhension risque légionellose**
   - **Légionelle :** bactérie pathogène eau chaude
   - Développement optimal : 25-45°C
   - Multiplication rapide si stagnation
   - Transmission : inhalation gouttelettes (douche, vapeur)
   - Maladie : légionellose (pneumopathie grave, potentiellement mortelle)
   - **Zones à risque :**
     - Température 25-50°C
     - Stagnation eau > 3 jours
     - Réseau complexe (bras morts)
     - Ballon surdimensionné

2. **Prévention température**
   - **Température stockage ECS : 55-60°C minimum**
   - À 55°C : légionelle ne se multiplie pas
   - À 60°C : destruction progressive légionelle
   - À 70°C : destruction rapide (90% en 2 minutes)
   - NE JAMAIS stocker eau < 50°C (sauf instantané)

3. **Activation cycle anti-légionelle**
   - **Fonction chaudière moderne :** cycle hebdomadaire automatique
   - Principe :
     - 1 fois/semaine (souvent dimanche 2h du matin)
     - Chauffe ECS à 65-70°C pendant 30-60 minutes
     - Détruit légionelles dans ballon et réseau
   - **Vérifier activation :**
     - Menu chaudière : fonction anti-légionelle ON
     - Consulter notice : procédure activation
     - Planification : jour et heure cycle

4. **Configuration cycle anti-légionelle**
   - **Réglages typiques :**
     - Température cycle : 65-70°C
     - Durée : 30-60 minutes
     - Fréquence : hebdomadaire
     - Jour : programmable (ex: dimanche)
     - Heure : nuit (2h) pour éviter risque brûlure
   - **Vérification fonctionnement :**
     - Consulter historique (si disponible)
     - Provoquer cycle manuel (test)
     - Vérifier montée température (affichage)

5. **Cas chaudière instantanée**
   - Pas de stockage = pas de légionelle
   - Eau chauffée instantanément à chaque puisage
   - Température > 50°C généralement
   - **Risque résiduel :** tuyauteries en aval (stagnation)
   - Solution : purge hebdomadaire points peu utilisés

6. **Traitement ballon ECS existant**
   - Si suspicion contamination légionelle :
     - Choc thermique : 70°C minimum 30 minutes
     - Purge complète ballon (évacuation eau contaminée)
     - Nettoyage mécanique ballon si possible
     - Désinfection chimique (eau de Javel, chlore)
     - Rinçage abondant avant remise en service
   - Faire appel professionnel si installation complexe

7. **Mesures complémentaires**
   - **Conception installation :**
     - Limiter longueur tuyauteries (< 10m)
     - Éviter bras morts (points non utilisés)
     - Calorifuger tuyauteries (maintien température)
     - Boucle bouclée + circulateur (bâtiments collectifs)
   - **Dimensionnement ballon :**
     - Ballon adapté aux besoins (pas surdimensionné)
     - Rotation eau complète quotidienne
   - **Entretien :**
     - Détartrage régulier ballon
     - Purge points peu utilisés (hebdomadaire)
     - Vérification température régulière

8. **Réglementation (France)**
   - **ERP (Établissements Recevant du Public) :**
     - Température stockage : 55-60°C
     - Température distribution : ≥ 50°C
     - Contrôle légionelle : analyse eau selon arrêté
   - **Logements collectifs :**
     - Température ballon : 55-60°C
     - Température réseau : ≥ 50°C au point le plus éloigné
   - **Particuliers :** recommandations (non obligatoire mais fortement conseillé)

9. **Signaux d'alerte**
   - Température ECS < 50°C régulièrement
   - Eau tiède qui stagne
   - Odeur désagréable eau chaude
   - Points d'eau peu utilisés
   - Ballon ancien (> 15 ans, entartré)

**Prévention :**
- Température ECS : 55-60°C minimum
- Activation cycle anti-légionelle (vérifier annuellement)
- Purge hebdomadaire points peu utilisés
- Dimensionnement correct ballon
- Détartrage régulier
- Calorifugeage tuyauteries
- Éviter stagnation (rotation eau)

**⚠️ SANTÉ PUBLIQUE :**
- Légionellose = maladie déclaration obligatoire
- Taux mortalité : 10-15% (personnes fragiles)
- Prévention simple : température + circulation
- Ne jamais négliger température ECS

**Température et légionelle :**
- < 20°C : survie, pas de multiplication
- 25-45°C : multiplication rapide (optimal 35°C)
- 50-55°C : arrêt multiplication
- 60°C : destruction 90% en 32 minutes
- 66°C : destruction 90% en 2 minutes
- 70°C : destruction quasi-instantanée

---

## FACT-CHAUD-062: Fuite échangeur sanitaire (mélange eau chauffage/ECS)

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-062 |
| **Catégorie** | Eau Chaude Sanitaire (ECS) |
| **Système** | Échangeur sanitaire |
| **Gravité** | **Critique** |
| **Marques** | Multi-marques (échangeur à plaques) |

**Symptômes :**
- Pression circuit chauffage chute anormalement
- Eau ECS colorée (boues, inhibiteur chauffage)
- Goût métallique eau chaude
- Pression chauffage monte à l'ouverture robinet ECS
- Présence additif chauffage dans eau sanitaire

**Cause racine probable :**
Perforation échangeur à plaques (corrosion, surpression), joint échangeur défectueux, fissure plaque échangeur, entartrage excessif (contraintes mécaniques), gel échangeur (hiver).

**Étapes de résolution :**

1. **Diagnostic fuite échangeur**
   - **Test pression chauffage :**
     - Noter pression circuit chauffage (robinets ECS fermés)
     - Ouvrir robinet ECS et maintenir
     - Observer pression chauffage
     - Si pression monte : fuite échangeur sanitaire → chauffage
   - **Test visuel eau ECS :**
     - Eau colorée (rouge, brun) : inhibiteur/boues chauffage
     - Eau trouble : contamination
   - **Test chimique :**
     - Prélever eau ECS
     - Tester présence glycol ou inhibiteur (kit test)

2. **Confirmation fuite**
   - Isoler échangeur sanitaire (vannes)
   - Mettre circuit chauffage en pression (2 bars)
   - Fermer robinet remplissage chauffage
   - Noter pression et attendre 1 heure
   - Si pression chute : fuite interne échangeur confirmée

3. **Comprendre gravité**
   - **DANGER SANITAIRE :**
     - Eau chauffage = additifs chimiques (inhibiteur corrosion, glycol)
     - NON POTABLE
     - Risque intoxication si ingestion
   - **Contamination réseau eau potable :**
     - Retour eau chauffage vers réseau eau froide (si pression inversée)
     - Réglementation : disconnecteur obligatoire (protection réseau)

4. **Arrêt immédiat installation**
   - Fermer robinet eau froide chaudière
   - Couper chaudière
   - Vidanger circuit sanitaire
   - NE PLUS CONSOMMER eau chaude
   - Prévenir occupants (risque sanitaire)

5. **Démontage et inspection échangeur**
   - Vidanger circuit chauffage
   - Démonter échangeur à plaques
   - Inspecter plaques :
     - Recherche perforation (trous, fissures)
     - Corrosion (piqûres, zones fragilisées)
     - Déformation plaques
   - Contrôler joints (déchirés, écrasés)

6. **Causes fuite échangeur**
   - **Corrosion :**
     - Eau agressive (pH acide)
     - Électrolyse (métaux différents)
     - Eau non traitée circuit chauffage
   - **Surpression :**
     - Coup de bélier
     - Surchauffe (vapeur)
     - Gel (dilatation glace)
   - **Vieillissement :**
     - Fatigue matériau
     - Cycles thermiques répétés
     - Entartrage → contraintes

7. **Réparation / Remplacement**
   - **Échangeur à plaques :**
     - Si perforation plaques : remplacement échangeur obligatoire
     - Plaques non réparables (soudure impossible, dangereuse)
     - Si joint défectueux seul : remplacement joints possible
   - **Coût :**
     - Échangeur neuf : 100-400€ selon modèle
     - Main d'œuvre : 2-3 heures
   - **Référence :**
     - Commander échangeur référence constructeur
     - Fournir n° série chaudière (compatibilité)

8. **Remise en service sécurisée**
   - Installer échangeur neuf avec joints neufs
   - Rincer abondamment circuit sanitaire (purge contamination)
   - Remplir circuit chauffage (eau + additifs neufs)
   - Mettre en pression et tester étanchéité
   - Vérifier absence contamination :
     - Ouvrir ECS : eau claire, inodore
     - Test chimique : absence additif chauffage
   - Purger plusieurs fois avant consommation

9. **Prévention contamination réseau**
   - **Disconnecteur obligatoire :**
     - BA type BA (protection réseau eau potable)
     - Empêche retour eau chauffage → eau froide
     - Vérification annuelle obligatoire (réglementation)
   - Contrôle pression différentielle chauffage/sanitaire

**Prévention :**
- Traitement eau chauffage (inhibiteur, pH neutre)
- Protection antigel (hors gel local)
- Détartrage régulier (éviter contraintes)
- Contrôle pression (éviter surpression)
- Remplacement préventif échangeur (durée vie ~15 ans)
- Disconnecteur vérifié annuellement

**⚠️ DANGER SANITAIRE :**
- Eau chauffage = NON POTABLE
- Additifs chimiques toxiques
- En cas contamination ECS :
  - Arrêt consommation immédiat
  - Remplacement échangeur obligatoire
  - Rinçage complet réseau sanitaire
  - Purges multiples avant consommation

**Réglementation :**
- Disconnecteur BA obligatoire (protection réseau)
- DTU 65.10 : installations chauffage
- Arrêté 30/11/2005 : protection réseau eau potable

---

## FACT-CHAUD-063: Pression ECS excessive

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-063 |
| **Catégorie** | Eau Chaude Sanitaire (ECS) |
| **Système** | Pression sanitaire |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques |

**Symptômes :**
- Pression ECS trop élevée (> 5 bars)
- Coups de bélier au robinet
- Soupape sécurité sanitaire qui fuit
- Bruit fort à la fermeture robinet
- Groupe sécurité goutte en continu

**Cause racine probable :**
Réducteur pression défectueux (trop haute pression réseau), absence réducteur pression, dilatation eau chaude (pas de vase expansion sanitaire), clapet anti-retour bloqué, groupe sécurité sous-dimensionné.

**Étapes de résolution :**

1. **Mesure pression ECS**
   - Installer manomètre sur robinet purge
   - Mesurer pression eau froide statique
   - Mesurer pression eau chaude (robinets fermés)
   - **Pressions normales :**
     - Réseau eau froide : 2-4 bars (ville)
     - Maximum recommandé : 5 bars
     - Maximum admissible : 6-7 bars (risque)

2. **Contrôle réducteur pression**
   - Localiser réducteur pression (entrée installation)
   - Mesurer pression aval réducteur
   - Réglage réducteur :
     - Pression cible : 3 bars
     - Ajuster vis réglage (rotation)
     - Tester après ajustement
   - Si réglage impossible : réducteur HS (membrane, ressort)
   - Remplacer réducteur si défectueux

3. **Contrôle dilatation eau chaude**
   - **Problème :**
     - Eau chaude se dilate (+3% de 20 à 60°C)
     - Si circuit fermé (clapet anti-retour) : surpression
     - Pression peut monter à 10-15 bars (danger)
   - **Solution 1 : Vase expansion sanitaire**
     - Vase 2-4 litres sur circuit ECS
     - Absorbe dilatation
     - Prégonflage : 80% pression eau froide
   - **Solution 2 : Groupe sécurité**
     - Groupe sécurité ECS (soupape + disconnecteur)
     - Évacue surpression par goutte-à-goutte
     - Raccordement évacuation obligatoire

4. **Contrôle groupe sécurité**
   - Groupe sécurité = soupape taré 7 bars
   - Fonctionnement normal : goutte occasionnelle (dilatation)
   - Fonctionnement anormal : fuite continue
   - **Causes fuite continue :**
     - Soupape entartrée (reste ouverte)
     - Pression excessive (> 7 bars)
     - Soupape HS (ressort fatigué)
   - **Test groupe sécurité :**
     - Actionner manette vidange
     - Eau doit s'écouler franchement
     - Relâcher : écoulement doit s'arrêter
     - Si fuite persiste : détartrer ou remplacer

5. **Détartrage groupe sécurité**
   - Démonter groupe sécurité
   - Tremper dans vinaigre blanc 12h
   - Actionner mécanisme plusieurs fois
   - Rincer abondamment
   - Si fuite persiste après nettoyage : remplacement

6. **Installation vase expansion sanitaire**
   - **Si absent :**
     - Vase 2-4L selon volume ECS
     - Montage sur départ eau chaude
     - Prégonflage côté air : 2,5 bars (80% pression eau froide)
   - **Avantages :**
     - Absorbe dilatation
     - Réduit cycles groupe sécurité
     - Protège installation
     - Réduit consommation eau (moins pertes)

7. **Contrôle clapet anti-retour**
   - Clapet anti-retour : empêche retour ECS vers eau froide
   - Si bloqué fermé : circuit fermé → surpression
   - Démonter, nettoyer ou remplacer

8. **Cas pression réseau excessive**
   - Si pression réseau > 5 bars :
     - Installation réducteur pression OBLIGATOIRE
     - Protection installation (robinets, flexibles)
     - Confort (éviter coups de bélier)
   - Réducteur à membrane : plus fiable
   - Réglage : 3 bars recommandé

**Prévention :**
- Installation réducteur pression (pression réseau > 4 bars)
- Vase expansion sanitaire (circuit ECS fermé)
- Groupe sécurité vérifié annuellement
- Détartrage groupe sécurité si eau dure
- Contrôle pression annuel

**Réglementation :**
- Groupe sécurité obligatoire (chauffe-eau, ballon)
- Pression max admissible : 7 bars (appareils domestiques)
- Évacuation groupe sécurité : visible, raccordée évacuation

**Dimensionnement vase expansion sanitaire :**
- Formule : V = (Ve × Ce × ΔT) / (1 - (P1/P2))
  - Ve : volume eau ECS (litres)
  - Ce : coefficient expansion eau (0,04 pour ΔT 50°C)
  - ΔT : variation température
  - P1 : pression gonflage vase (bars absolus)
  - P2 : pression tarage soupape (bars absolus)
- Exemple : ballon 200L → vase 3-4L

---

## FACT-CHAUD-064: Brûleur ne démarre pas sur demande ECS

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-064 |
| **Catégorie** | Eau Chaude Sanitaire (ECS) |
| **Système** | Démarrage ECS |
| **Gravité** | **Élevée** |
| **Marques** | Multi-marques |

**Symptômes :**
- Ouverture robinet ECS : aucune réaction chaudière
- Brûleur ne s'allume pas en demande sanitaire
- Mode chauffage fonctionne normalement
- Pas de bruit, pas de ventilateur
- Afficheur ne passe pas en mode ECS

**Cause racine probable :**
Débitmètre ne détecte pas débit, carte électronique ne reçoit pas signal, priorité ECS désactivée, programmation bloquée, capteur débit défectueux, pression eau insuffisante.

**Étapes de résolution :**

1. **Vérification basique**
   - Ouvrir robinet ECS en grand (débit maximum)
   - Maintenir ouvert 30 secondes
   - Observer afficheur chaudière :
     - Symbole robinet/goutte doit apparaître
     - Ou passage mode ECS
   - Écouter brûleur (tentative démarrage)

2. **Contrôle débit et pression**
   - Mesurer débit robinet (litres/minute)
   - Débit minimum déclenchement : 2-4 L/min (selon modèle)
   - Si débit faible : problème hydraulique amont
   - Mesurer pression eau froide (manomètre)
   - Pression minimum : 1,5-2 bars
   - Si pression insuffisante : pas de détection débit

3. **Contrôle débitmètre**
   - Voir FACT-CHAUD-059 (détails complets)
   - Vérifier rotation turbine débitmètre
   - Nettoyer si encrassé
   - Tester signal électrique
   - Remplacer si défectueux

4. **Contrôle paramètres chaudière**
   - **Mode chaudière :**
     - Vérifier mode : Été, Hiver, ECS seule
     - Mode "Hiver" ou "Auto" : chauffage + ECS
     - Mode "Été" : ECS seule
     - Mode "Chauffage seul" : pas d'ECS
   - **Priorité ECS :**
     - Paramètre priorité ECS activé
     - Si désactivé : ECS ne démarre pas
   - **Programmation :**
     - Vérifier plages horaires ECS autorisées
     - Mode absence/vacances : ECS parfois bloquée
     - Débloquer programmation

5. **Contrôle carte électronique**
   - Vérifier réception signal débitmètre :
     - Menu diagnostic : affichage débit ECS
     - Ouvrir robinet : débit doit s'afficher
     - Si pas d'affichage : câblage ou carte
   - Contrôler sortie commande :
     - Vanne 3 voies doit commuter
     - Brûleur doit recevoir ordre démarrage
   - Tester mode forcé ECS si disponible

6. **Contrôle vanne 3 voies**
   - Voir FACT-CHAUD-058 (détails complets)
   - Vanne doit commuter en position ECS
   - Écouter bruit moteur vanne
   - Vérifier position vanne (manuelle si possible)
   - Débloquer ou remplacer si nécessaire

7. **Contrôle sécurités**
   - Vérifier absence blocage sécurité :
     - Surchauffe
     - Pression chauffage trop basse
     - Défaut combustion
   - Consulter codes erreur historique
   - Réarmer sécurités si nécessaire

8. **Contrôle contacteur priorité (si présent)**
   - Certaines installations : contacteur externe priorité ECS
   - Contacteur ballons tampons, bouilleurs
   - Vérifier fonctionnement contacteur
   - Tester continuité (multimètre)

9. **Test diagnostic carte**
   - Accéder menu diagnostic/installateur
   - Forcer demande ECS (mode test)
   - Observer comportement :
     - Vanne 3 voies commute ?
     - Ventilateur démarre ?
     - Brûleur s'allume ?
   - Identifier étape défaillante

**Prévention :**
- Vérification annuelle fonctionnement ECS (entretien)
- Test débitmètre (nettoyage)
- Contrôle paramètres après coupure courant
- Vérification vanne 3 voies
- Maintien pression eau suffisante

**Diagnostic par étapes :**
1. Débitmètre détecte débit ? → Si non : nettoyer/remplacer
2. Carte reçoit signal ? → Si non : câblage
3. Vanne 3 voies commute ? → Si non : vanne HS
4. Ventilateur démarre ? → Si non : sécurité/ventilateur
5. Brûleur s'allume ? → Si non : voir combustion

**Seuil déclenchement ECS (exemples) :**
- Saunier Duval : 2,5 L/min
- Vaillant : 2,8 L/min
- Elm Leblanc : 2,5 L/min
- Réglable sur certains modèles (menu installateur)

---

## FACT-CHAUD-065: Mitigeur thermostatique défectueux

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-065 |
| **Catégorie** | Eau Chaude Sanitaire (ECS) |
| **Système** | Distribution ECS |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques (mitigeurs tous fabricants) |

**Symptômes :**
- Température eau mitigée instable
- Impossible régler température souhaitée
- Eau brûlante ou froide malgré réglage
- Débit très faible au mitigeur
- Mitigeur bloqué (poignée dure)

**Cause racine probable :**
Cartouche thermostatique entartrée, clapet anti-retour bloqué, filtres mitigeur bouchés, cartouche usée (joint, ressort), inversion eau chaude/froide, pression différentielle excessive.

**Étapes de résolution :**

1. **Compréhension mitigeur thermostatique**
   - **Fonction :**
     - Mélange automatique eau chaude/froide
     - Maintien température constante (± 2°C)
     - Protection anti-brûlure (sécurité enfants)
   - **Composition :**
     - Cartouche thermostatique (bilame, cire)
     - Filtres entrée eau chaude + froide
     - Clapets anti-retour
     - Limiteur température (38°C généralement)

2. **Diagnostic panne**
   - **Test température :**
     - Régler mitigeur température médiane (38°C)
     - Mesurer température eau sortie (thermomètre)
     - Écart > 5°C : cartouche défectueuse
   - **Test débit :**
     - Ouvrir complètement débit
     - Débit faible : filtres bouchés ou cartouche
   - **Test mécanique :**
     - Tourner poignée température
     - Blocage ou dureté excessive : entartrage

3. **Nettoyage filtres mitigeur**
   - Fermer vannes d'arrêt eau chaude + froide
   - Purger pression (ouvrir mitigeur)
   - Dévisser raccords entrée eau (clé)
   - Extraire filtres (petits tamis)
   - Nettoyer filtres :
     - Brosser sous eau
     - Détartrer vinaigre blanc si calcaire
     - Rincer abondamment
   - Remonter filtres (sens correct)

4. **Contrôle clapets anti-retour**
   - Clapets intégrés raccords mitigeur
   - Fonction : empêche mélange eau chaude/froide en amont
   - Démonter et nettoyer clapets
   - Vérifier mobilité clapet (ressort)
   - Remplacer joints si usés

5. **Démontage et nettoyage cartouche**
   - **Préparation :**
     - Fermer eau, purger pression
     - Retirer enjoliveur (cache central)
     - Dévisser vis fixation poignée
     - Retirer poignée température
   - **Extraction cartouche :**
     - Dévisser écrou fixation cartouche
     - Extraire cartouche (parfois dur : entartrage)
     - Utiliser pince si nécessaire (délicatement)
   - **Nettoyage cartouche :**
     - Tremper cartouche vinaigre blanc 12-24h
     - Brosser délicatement extérieur
     - Rincer abondamment eau claire
     - Actionner mécanisme (vérifier mobilité)
     - Si bloqué après nettoyage : remplacement

6. **Remplacement cartouche thermostatique**
   - **Identification :**
     - Noter marque et modèle mitigeur
     - Cartouches spécifiques par fabricant :
       - Grohe : cartouche 47450
       - Hansgrohe : cartouche 94282000
       - Porcher : cartouche spécifique modèle
       - Etc.
   - **Installation cartouche neuve :**
     - Nettoyer logement cartouche (corps mitigeur)
     - Installer cartouche (repérage crans)
     - Serrer écrou fixation (modérément)
     - Remonter poignée (position repère)
     - Vérifier jeu axial (pas trop serré)

7. **Calibration température**
   - Ouvrir vannes eau chaude + froide
   - Positionner poignée sur 38°C (repère)
   - Ouvrir mitigeur, attendre stabilisation
   - Mesurer température (thermomètre)
   - **Ajustement :**
     - Retirer poignée
     - Tourner axe cartouche (recalibrage)
     - Remonter poignée position ajustée
     - Retester jusqu'à température correcte (38°C)

8. **Contrôle pression différentielle**
   - **Problème :**
     - Mitigeur thermostatique sensible écart pression
     - Écart > 2 bars : dysfonctionnement
   - **Mesure :**
     - Pression eau froide vs eau chaude
     - Écart normal : < 1 bar
   - **Solution si écart excessif :**
     - Réducteur pression sur circuit haute pression
     - Égalisation pressions eau chaude/froide

9. **Vérification installation**
   - **Raccordement :**
     - Eau chaude : GAUCHE (norme)
     - Eau froide : DROITE
     - Inversion = dysfonctionnement total
   - **Clapets anti-retour :**
     - Obligatoires (protection réseau)
     - Vérifier présence et fonctionnement

**Prévention :**
- Détartrage annuel mitigeur (eau dure)
- Nettoyage filtres tous les 6 mois
- Installation adoucisseur (TH > 25°F)
- Température ECS ≤ 60°C (limite entartrage)
- Remplacement préventif cartouche (5-10 ans)

**Avantages mitigeur thermostatique :**
- Sécurité anti-brûlure (enfants, personnes âgées)
- Confort (température stable)
- Économie eau (pas de réglage manuel)
- Obligatoire certains ERP

**Durée vie :**
- Cartouche : 5-10 ans (selon dureté eau, usage)
- Mitigeur complet : 15-20 ans

**Coût remplacement :**
- Cartouche seule : 30-80€
- Mitigeur complet : 80-300€ (selon gamme)
- Main d'œuvre : 1h

---

*Fin du fichier 04_ECS.md*

**Retour à l'index :** [Knowledge_Base_Chaudieres.md](Knowledge_Base_Chaudieres.md)
