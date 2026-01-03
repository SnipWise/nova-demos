# Problèmes Régulation & Sondes

**Catégorie :** Régulation & Sondes
**Nombre de Facts :** 15
**Retour à l'index :** [Knowledge_Base_Chaudieres.md](Knowledge_Base_Chaudieres.md)

---

## FACT-CHAUD-066: Sonde température départ défectueuse

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-066 |
| **Catégorie** | Régulation & Sondes |
| **Système** | Sonde température départ |
| **Gravité** | **Élevée** |
| **Marques** | Multi-marques |

**Symptômes :**
- Affichage température incohérente ou absente
- Code erreur sonde départ (ex: Saunier Duval F73, Vaillant F75)
- Température affichée bloquée ou erratique
- Chaudière en sécurité ou fonctionne en mode dégradé
- Pas de modulation de puissance

**Cause racine probable :**
Sonde CTN/NTC défectueuse, oxydation connectique, câble coupé ou court-circuité, connecteur déconnecté, corrosion doigt de gant.

**Étapes de résolution :**

1. **Diagnostic initial**
   - Relever température affichée et comparer au réel (thermomètre contact)
   - Vérifier code erreur spécifique dans menu diagnostic
   - Noter si valeur bloquée ou variable anormalement
   - Tester en mode chauffage et ECS si applicable

2. **Mesure résistance sonde**
   - Couper alimentation électrique chaudière
   - Déconnecter sonde au niveau carte électronique
   - Mesurer résistance avec multimètre (ohmmètre)
   - Valeurs courantes CTN 10k : 10 kΩ à 25°C, 3,3 kΩ à 50°C, 1,5 kΩ à 70°C
   - Comparer avec courbe constructeur (notice technique)

3. **Test variation résistance**
   - Chauffer légèrement la sonde (chaleur main ou eau chaude)
   - La résistance doit diminuer progressivement (CTN)
   - Si résistance constante : sonde morte
   - Si résistance infinie : sonde coupée ou câble sectionné
   - Si résistance nulle : court-circuit

4. **Contrôle câblage**
   - Vérifier continuité câble entre sonde et carte
   - Contrôler isolation câble (pas de court-circuit à la masse)
   - Inspecter connectique (oxydation, corrosion)
   - Nettoyer contacts avec bombe contact électronique
   - Vérifier fixation connecteur (bien clipsé)

5. **Contrôle doigt de gant**
   - Déposer la sonde du doigt de gant
   - Vérifier absence d'eau dans doigt de gant (corrosion)
   - Nettoyer et sécher si humidité
   - Appliquer pâte thermique pour bon contact thermique
   - Vérifier fixation doigt de gant (pas de jeu)

6. **Remplacement sonde**
   - Identifier référence exacte selon modèle chaudière
   - Déposer ancienne sonde (noter position, longueur immergée)
   - Installer sonde neuve avec pâte thermique
   - Reconnecter en respectant polarité si applicable
   - Isoler connexion de l'humidité

7. **Test et calibration**
   - Remettre en service la chaudière
   - Vérifier température affichée cohérente
   - Comparer affichage avec thermomètre de référence
   - Laisser chauffer : vérifier évolution température
   - Tester modulation puissance
   - Calibrer si fonction disponible (menu installateur)

**Prévention :**
- Vérification annuelle valeur sonde (menu diagnostic)
- Contrôle connectique (oxydation)
- Protection connexions contre humidité (graisse silicone)
- Remplacement préventif si valeurs dérivent (± 3°C)
- Application pâte thermique lors entretien

**Valeurs typiques sondes CTN :**
- 10 kΩ à 25°C
- 5,8 kΩ à 40°C
- 3,3 kΩ à 50°C
- 2,0 kΩ à 60°C
- 1,5 kΩ à 70°C
- 0,8 kΩ à 90°C

**Note :** En cas de panne sonde, certaines chaudières passent en mode sécurité (arrêt), d'autres en mode dégradé (puissance fixe, température estimée).

---

## FACT-CHAUD-067: Sonde température retour HS

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-067 |
| **Catégorie** | Régulation & Sondes |
| **Système** | Sonde température retour |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques (chaudières avec régulation avancée) |

**Symptômes :**
- Code erreur sonde retour (moins fréquent que sonde départ)
- Calcul ΔT impossible ou erroné
- Protection anti-condensation non fonctionnelle (chaudières standard)
- Modulation sous-optimale
- Régulation hydraulique dégradée

**Cause racine probable :**
Sonde CTN défectueuse, mauvais contact thermique, câblage défectueux, connecteur oxydé, absence pâte thermique.

**Étapes de résolution :**

1. **Vérification présence sonde**
   - Toutes les chaudières n'ont pas de sonde retour
   - Vérifier schéma technique de la chaudière
   - Localiser la sonde (généralement sur tube retour primaire)
   - Identifier le connecteur sur carte électronique

2. **Diagnostic différentiel**
   - Comparer température retour affichée avec départ
   - ΔT normal chauffage : 15-25°C (retour < départ)
   - Si retour > départ : incohérence, sonde défectueuse
   - Si retour = départ : sonde mal placée ou défectueuse

3. **Mesure résistance sonde**
   - Couper alimentation électrique
   - Déconnecter sonde au niveau carte
   - Mesurer résistance (mêmes valeurs que sonde départ)
   - Comparer avec courbe constructeur
   - Mesurer température réelle eau retour (thermomètre contact)

4. **Contrôle positionnement sonde**
   - Vérifier que sonde est sur retour (pas départ)
   - Contrôler immersion dans doigt de gant
   - Vérifier contact thermique (pâte thermique)
   - S'assurer que doigt de gant est dans flux eau
   - Vérifier absence air dans doigt de gant

5. **Test croisé avec sonde départ**
   - Échanger temporairement sondes départ/retour (si identiques)
   - Si défaut suit la sonde : sonde HS
   - Si défaut reste sur même entrée carte : défaut carte
   - Remettre en position d'origine

6. **Remplacement et installation**
   - Remplacer par sonde identique (référence constructeur)
   - Appliquer pâte thermique généreusement
   - Enfoncer complètement dans doigt de gant
   - Fixer solidement
   - Protéger connexion de l'humidité

7. **Vérification fonctionnelle**
   - Contrôler températures affichées départ et retour
   - Vérifier calcul ΔT cohérent (menu diagnostic)
   - Tester régulation (adaptation puissance selon ΔT)
   - Vérifier anti-condensation si applicable
   - Observer modulation sur cycle complet

**Prévention :**
- Vérification annuelle cohérence départ/retour
- Contrôle ΔT lors entretien (valeur typique 15-20°C)
- Inspection connectique
- Renouvellement pâte thermique

**Impact panne sonde retour :**
- Perte régulation optimale (modulation moins précise)
- Protection anti-condensation inopérante (chaudières standard)
- Calcul rendement impossible
- Diagnostic hydraulique difficile
- Consommation potentiellement accrue

**Fonction sonde retour selon type chaudière :**
- **Chaudière condensation :** optimisation condensation, modulation
- **Chaudière standard :** protection anti-condensation (T retour > 50°C)
- **Chaudière modulante :** calcul ΔT pour ajuster puissance
- **Système hydraulique complexe :** gestion circuits multiples

---

## FACT-CHAUD-068: Sonde extérieure défaillante

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-068 |
| **Catégorie** | Régulation & Sondes |
| **Système** | Régulation climatique |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques (sur installations avec régulation climatique) |

**Symptômes :**
- Température extérieure affichée aberrante (ex: -50°C ou +80°C)
- Régulation climatique inopérante (retour loi d'eau fixe)
- Température départ ne s'adapte pas aux conditions extérieures
- Surchauffe ou sous-chauffe locaux
- Code erreur sonde extérieure (certaines marques)

**Cause racine probable :**
Sonde CTN défectueuse, câble sectionné (longueur importante), oxydation connexion, infiltration eau dans sonde, sonde mal positionnée (soleil direct, source chaleur).

**Étapes de résolution :**

1. **Vérification activation régulation climatique**
   - S'assurer que fonction est activée (menu installateur)
   - Vérifier paramétrage courbe de chauffe
   - Contrôler que sonde est bien déclarée dans système
   - Noter température extérieure affichée

2. **Diagnostic valeur affichée**
   - Comparer température affichée avec météo locale
   - Si écart > 5°C : problème sonde ou positionnement
   - Si valeur bloquée : sonde ou câble défectueux
   - Si valeur erratique : mauvaise connexion

3. **Contrôle positionnement sonde**
   - **Emplacement correct :**
     - Face nord ou nord-ouest du bâtiment
     - À l'ombre permanente (pas de soleil direct)
     - Éloignée fenêtres, bouches ventilation, cheminées
     - Hauteur 2-3 mètres
     - Protégée vent dominant mais bien ventilée
   - **Vérifier :**
     - Pas d'exposition soleil (fausse température +10 à +20°C)
     - Pas de source chaleur proximité
     - Pas d'eau stagnante (corrosion)

4. **Mesure résistance sonde**
   - Déconnecter sonde au niveau carte régulation
   - Mesurer résistance (CTN typique 10 kΩ à 25°C)
   - Comparer avec température réelle extérieure et courbe
   - Exemples : 15 kΩ à 15°C, 10 kΩ à 25°C, 28 kΩ à 0°C
   - Résistance infinie : câble coupé ou sonde HS

5. **Contrôle câblage**
   - Vérifier continuité sur toute longueur (peut être > 20m)
   - Contrôler isolation (pas de court-circuit)
   - Inspecter passage gaines, traversées murs
   - Vérifier absence écrasement, coupure
   - Tester connexions intermédiaires si présentes

6. **Contrôle boîtier sonde extérieure**
   - Ouvrir boîtier sonde (vis)
   - Vérifier étanchéité (joint, presse-étoupe câble)
   - Contrôler absence infiltration eau
   - Nettoyer et sécher si humidité
   - Vérifier fixation sonde dans boîtier

7. **Remplacement sonde**
   - Utiliser référence constructeur ou compatible
   - Installer en position optimale (voir point 3)
   - Assurer étanchéité boîtier et passage câble
   - Fixer solidement (pas de vibration vent)
   - Protéger connexions

8. **Paramétrage et test**
   - Vérifier température affichée cohérente
   - Activer régulation climatique
   - Ajuster courbe de chauffe selon besoin (pente + décalage)
   - Tester sur 24-48h : température départ doit évoluer
   - Affiner courbe selon retour occupants

**Prévention :**
- Vérification annuelle température affichée vs météo
- Contrôle étanchéité boîtier sonde
- Vérification positionnement (végétation, modification bâtiment)
- Test câble (continuité, isolation)

**Impact panne sonde extérieure :**
- Perte régulation climatique (retour thermostat seul)
- Inconfort thermique (température intérieure instable)
- Surconsommation (surchauffe par sécurité)
- Usure prématurée chaudière (cycles fréquents)

**Avantages régulation climatique fonctionnelle :**
- Anticipation besoins selon météo
- Confort amélioré (température stable)
- Économies 10-25% selon installation
- Moins de cycles marche/arrêt
- Adaptation automatique saison

**Courbe de chauffe typique :**
- Pente : relation entre T°ext et T°départ
- Exemple : T°ext -10°C → T°départ 70°C, T°ext +15°C → T°départ 35°C
- Ajustement selon isolation, émetteurs, occupation

---

## FACT-CHAUD-069: Sonde ballon ECS défectueuse

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-069 |
| **Catégorie** | Régulation & Sondes |
| **Système** | Production ECS accumulée |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques (chaudières avec ballon ECS) |

**Symptômes :**
- Température ECS affichée incohérente
- Eau tiède ou froide alors que chaudière chauffe
- Eau brûlante (surchauffe ballon)
- Cycles de chauffe incessants ou absents
- Code erreur sonde ballon (selon marque)

**Cause racine probable :**
Sonde CTN défectueuse, mauvais contact thermique avec ballon, sonde mal positionnée (hors eau), connectique oxydée, câble endommagé.

**Étapes de résolution :**

1. **Diagnostic symptômes**
   - Relever température ECS affichée
   - Mesurer température réelle eau (thermomètre robinet)
   - Noter fréquence cycles chauffe ballon
   - Vérifier consigne température ECS (menu)
   - Observer déclenchements chauffage ballon

2. **Localisation sonde ballon**
   - Sonde généralement dans doigt de gant sur ballon
   - Position idéale : tiers supérieur ballon
   - Parfois sonde sur départ ECS ou bouclage
   - Identifier câble et connecteur carte chaudière

3. **Mesure résistance sonde**
   - Couper alimentation électrique
   - Déconnecter sonde au niveau carte
   - Mesurer résistance avec multimètre
   - Comparer avec température réelle eau ballon
   - Valeurs attendues CTN : 2 kΩ à 60°C, 1,5 kΩ à 70°C
   - Incohérence : sonde défectueuse

4. **Contrôle positionnement sonde**
   - Vérifier immersion complète dans doigt de gant
   - S'assurer que doigt de gant baigne dans l'eau
   - Vérifier pâte thermique (contact thermique)
   - Contrôler absence air dans doigt de gant
   - Vérifier que doigt de gant n'est pas obstrué (tartre)

5. **Test comportement**
   - Lancer cycle chauffe ECS manuellement
   - Observer évolution température affichée
   - Température doit monter progressivement
   - Arrêt chauffe doit se produire à consigne
   - Si pas d'arrêt : sonde ne détecte pas montée température

6. **Contrôle installation**
   - Sur ballon indirect : vérifier circulation primaire (vanne 3 voies)
   - Contrôler pompe ballon si présente
   - Vérifier absence entartrage échangeur ballon
   - S'assurer stratification correcte ballon

7. **Remplacement sonde**
   - Vidanger partiellement ballon si nécessaire
   - Déposer ancienne sonde (dévisser doigt de gant)
   - Nettoyer doigt de gant (détartrage si besoin)
   - Appliquer pâte thermique
   - Installer sonde neuve, enfoncer complètement
   - Resserrer doigt de gant (joint neuf)
   - Reconnecter électriquement

8. **Paramétrage et test**
   - Régler consigne ECS (généralement 55-60°C)
   - Lancer cycle chauffe complet
   - Vérifier arrêt à consigne
   - Tester puisage : température cohérente
   - Vérifier absence relance intempestive
   - Ajuster consigne selon confort et légionelle

**Prévention :**
- Vérification annuelle température ballon (affichée vs réelle)
- Contrôle cycles chauffe ECS
- Détartrage ballon selon dureté eau (tous les 3-5 ans)
- Contrôle connectique sonde

**Risques sonde ballon défectueuse :**
- **Sonde indiquant trop chaud :** eau tiède, inconfort, risque légionelle
- **Sonde indiquant trop froid :** surchauffe, ébouillantage, surconsommation, entartrage accéléré
- **Sonde HS :** pas de production ECS ou chauffe permanente

**Réglage température ECS :**
- **Minimum :** 55°C (prévention légionelle)
- **Recommandé :** 55-60°C (confort + sécurité sanitaire)
- **Maximum :** 65°C (limite entartrage, sécurité)
- **Avec mitigeur thermostatique :** possible 60-65°C (distribution 50°C)

**Cycle anti-légionelle :**
- Certaines chaudières proposent montée 65-70°C hebdomadaire
- Désinfection thermique ballon
- Vérifier fonction active et opérationnelle

---

## FACT-CHAUD-070: Régulation climatique mal paramétrée

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-070 |
| **Catégorie** | Régulation & Sondes |
| **Système** | Régulation climatique |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques |

**Symptômes :**
- Température intérieure instable ou inconfortable
- Surchauffe locaux (trop chaud malgré thermostat)
- Sous-chauffe locaux (pas assez chaud)
- Température départ inadaptée aux conditions
- Surconsommation énergétique

**Cause racine probable :**
Courbe de chauffe inadaptée (pente incorrecte), décalage parallèle mal réglé, sonde extérieure mal positionnée, absence optimisation lors mise en service, modification isolation non prise en compte.

**Étapes de résolution :**

1. **Compréhension régulation climatique**
   - Principe : température départ varie selon température extérieure
   - Plus il fait froid dehors, plus eau départ est chaude
   - Loi d'eau = courbe de chauffe (graphique T°ext vs T°départ)
   - Objectif : température intérieure stable quelle que soit météo

2. **Diagnostic situation actuelle**
   - Relever paramètres courbe chauffe actuelle (menu installateur)
   - Noter pente et décalage parallèle
   - Mesurer température départ en fonction T°ext
   - Interroger occupants sur confort (trop chaud/froid, quand)
   - Relever température intérieure réelle (plusieurs pièces)

3. **Vérification sonde extérieure**
   - S'assurer température extérieure affichée cohérente
   - Vérifier positionnement correct sonde (voir FACT-CHAUD-068)
   - Pas de soleil direct (fausserait mesure)
   - Comparer avec météo locale

4. **Analyse courbe de chauffe**
   - **Pente (inclinaison courbe) :**
     - Pente faible (0,3-0,8) : plancher chauffant, forte isolation
     - Pente moyenne (1,0-1,5) : radiateurs BT, isolation moyenne
     - Pente forte (1,5-2,5) : radiateurs HT, faible isolation
     - Pente excessive : radiateurs anciens, pas d'isolation
   - **Décalage parallèle :**
     - Décale toute la courbe vers haut ou bas
     - Ajustement fin du confort
     - Plage typique : -10 à +10°C

5. **Ajustement courbe - Méthode progressive**
   - **Si trop froid en permanence :**
     - Augmenter décalage parallèle (+2 à +5°C)
     - Attendre 24-48h, évaluer
     - Si insuffisant : augmenter pente (0,1 à 0,2)
   - **Si trop chaud en permanence :**
     - Réduire décalage parallèle (-2 à -5°C)
     - Attendre 24-48h, évaluer
     - Si insuffisant : réduire pente (0,1 à 0,2)
   - **Si trop froid seulement quand très froid dehors :**
     - Augmenter pente uniquement
   - **Si trop chaud seulement en mi-saison :**
     - Réduire décalage parallèle

6. **Paramétrage type émetteurs**
   - **Plancher chauffant :**
     - Pente : 0,3 à 0,6
     - T°départ max : 35-45°C
     - Régulation lente (inertie importante)
   - **Radiateurs basse température :**
     - Pente : 0,8 à 1,2
     - T°départ max : 45-55°C
   - **Radiateurs moyenne température :**
     - Pente : 1,2 à 1,8
     - T°départ max : 55-65°C
   - **Radiateurs haute température :**
     - Pente : 1,8 à 2,5
     - T°départ max : 70-80°C

7. **Optimisations complémentaires**
   - Activer optimisation démarrage (anticipation)
   - Paramétrer température réduit nuit (économies)
   - Régler temporisations (éviter cycles courts)
   - Configurer influence thermostat ambiance (pondération)

8. **Test et affinage**
   - Laisser fonctionner 48-72h entre chaque modification
   - Interroger occupants régulièrement
   - Mesurer températures intérieures (plusieurs pièces)
   - Noter consommation avant/après
   - Affiner par petites touches (patience nécessaire)

**Prévention :**
- Révision paramètres après travaux isolation
- Ajustement saisonnier (automne/hiver)
- Réévaluation annuelle confort occupants
- Formation utilisateurs (ne pas modifier n'importe comment)

**Erreurs fréquentes :**
- Pente trop forte : surchauffe mi-saison, cycles courts
- Pente trop faible : sous-chauffe grand froid
- Modifications trop fréquentes (ne pas laisser temps stabilisation)
- Ignorer l'inertie du bâtiment
- Ne pas tenir compte des apports gratuits (soleil, occupation)

**Gains régulation climatique bien réglée :**
- Économies : 10-25% selon installation
- Confort : température stable ±0,5°C
- Longévité chaudière : moins de cycles
- Réduction empreinte carbone

**Exemple courbe type (radiateurs MT, T°int consigne 20°C) :**
| T°ext | T°départ (pente 1,4) |
|-------|----------------------|
| -10°C | 70°C |
| -5°C  | 63°C |
| 0°C   | 56°C |
| 5°C   | 49°C |
| 10°C  | 42°C |
| 15°C  | 35°C |
| 20°C  | Arrêt chauffage |

---

## FACT-CHAUD-071: Courbe de chauffe inadaptée

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-071 |
| **Catégorie** | Régulation & Sondes |
| **Système** | Loi d'eau |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques (régulation climatique) |

**Symptômes :**
- Inconfort variable selon température extérieure
- Locaux froids lors grand froid, surchauffe mi-saison
- Thermostat ambiance sollicité en permanence
- Consommation énergétique excessive
- Cycles marche/arrêt trop fréquents en mi-saison

**Cause racine probable :**
Pente courbe incorrecte pour type émetteurs ou isolation, mise en service bâclée, modification bâtiment non prise en compte, réglages usine non adaptés.

**Étapes de résolution :**

1. **Identification problème spécifique**
   - **Symptôme A :** Froid quand T°ext basse, OK quand doux
     - → Pente insuffisante (courbe trop "plate")
   - **Symptôme B :** OK quand froid, trop chaud quand doux
     - → Pente excessive (courbe trop "raide")
   - **Symptôme C :** Toujours trop froid
     - → Décalage parallèle insuffisant ou pente trop faible
   - **Symptôme D :** Toujours trop chaud
     - → Décalage parallèle excessif ou pente trop forte

2. **Détermination pente théorique**
   - **Méthode calcul :**
     - Pente = (T°départ grand froid - T°départ mi-saison) / (T°mi-saison - T°grand froid)
     - Exemple : (70°C - 40°C) / (10°C - (-10°C)) = 30/20 = 1,5
   - **Méthode empirique selon émetteurs :**
     - Plancher chauffant : 0,3-0,6
     - Ventilo-convecteurs BT : 0,6-0,9
     - Radiateurs alu BT : 0,8-1,2
     - Radiateurs fonte MT : 1,2-1,8
     - Radiateurs fonte HT : 1,8-2,5

3. **Calcul température départ nécessaire**
   - Identifier température dimensionnement (T°ext base)
   - Exemple : T°base = -7°C dans région (donnée climatique)
   - Identifier puissance/température émetteurs
   - Exemple radiateurs : 20°C ambiance à 60°C départ
   - Déduire T°départ max nécessaire

4. **Traçage courbe adaptée**
   - **Point 1 (grand froid) :**
     - T°ext = température base région (-5 à -15°C selon zone)
     - T°départ = température max émetteurs (35-80°C)
   - **Point 2 (arrêt chauffage) :**
     - T°ext = 15-18°C (seuil arrêt chauffage)
     - T°départ = arrêt ou mini (30-35°C)
   - **Pente = (T°départ P1 - T°départ P2) / (T°ext P2 - T°ext P1)**

5. **Ajustement progressif**
   - Partir de pente calculée ou recommandée
   - Programmer dans régulation chaudière
   - Laisser fonctionner 48h minimum
   - Relever températures intérieures matin/soir
   - Ajuster par incrément 0,1 pente
   - Répéter jusqu'à confort optimal

6. **Optimisation décalage parallèle**
   - Une fois pente correcte, ajuster décalage
   - Décalage = ajustement fin ±2 à ±5°C
   - Ne compense pas une pente incorrecte
   - Utiliser pour petits ajustements saisonniers

7. **Prise en compte spécificités**
   - **Inertie bâtiment :**
     - Bâtiment lourd : inertie forte, anticipation nécessaire
     - Bâtiment léger : réaction rapide, pente peut être plus faible
   - **Apports gratuits :**
     - Forte occupation : réduire légèrement pente
     - Apports solaires importants : idem
   - **Altitude :**
     - Augmenter pente en altitude (T°base plus basse)

8. **Validation sur cycle complet**
   - Tester sur période froide (-5°C ou moins)
   - Tester sur période mi-saison (5-15°C)
   - Vérifier confort stable
   - Mesurer consommation (comparaison avant/après)
   - Ajustement final si nécessaire

**Prévention :**
- Mise en service rigoureuse par professionnel qualifié
- Réévaluation après travaux isolation
- Ajustement après changement émetteurs
- Documentation paramètres (traçabilité)

**Cas particuliers :**

**Installation mixte (radiateurs + plancher) :**
- Créer 2 zones avec courbes différentes
- Ou compromis pente intermédiaire + mélange
- Préférer séparation hydraulique (2 départs)

**Rénovation énergétique :**
- Isolation renforcée → réduire pente (besoin moins important)
- Remplacement fenêtres → idem
- Après travaux : reprendre réglage à zéro

**Régulation pièce par pièce :**
- Têtes thermostatiques radiateurs
- Ajustent localement selon besoins
- Courbe chauffe = base, têtes = ajustement fin

**Outils aide réglage :**
- Applications constructeur (Vaillant, Viessmann, etc.)
- Abaques courbes chauffe (documentation technique)
- Régulateur Lernen/Auto-adapt (apprentissage automatique)

**Formule simplifiée pente :**
Pente ≈ (ΔT eau émetteur × 1,3) / 30
- Exemple radiateurs MT (ΔT = 50°C) : (50 × 1,3) / 30 ≈ 2,2
- Exemple plancher chauffant (ΔT = 15°C) : (15 × 1,3) / 30 ≈ 0,65

---

## FACT-CHAUD-072: Thermostat d'ambiance sans effet

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-072 |
| **Catégorie** | Régulation & Sondes |
| **Système** | Thermostat d'ambiance |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques |

**Symptômes :**
- Modification consigne thermostat sans effet sur chaudière
- Chaudière fonctionne en continu malgré température atteinte
- Ou chaudière ne démarre pas malgré demande thermostat
- Température intérieure ne correspond pas à consigne
- Thermostat affiche correct mais régulation inopérante

**Cause racine probable :**
Panne thermostat (piles, capteur), problème câblage, mauvais paramétrage chaudière (influence thermostat), thermostat mal positionné, conflit régulation climatique/thermostat.

**Étapes de résolution :**

1. **Vérification type thermostat**
   - **Thermostat filaire TOR (Tout Ou Rien) :**
     - Contact sec marche/arrêt
     - 2 fils vers chaudière
   - **Thermostat modulant (OpenTherm, eBUS) :**
     - Communication bidirectionnelle
     - 2 fils bus de communication
   - **Thermostat radio/WiFi :**
     - Récepteur sur/dans chaudière
     - Communication sans fil

2. **Test thermostat TOR filaire**
   - Retirer thermostat du socle mural
   - Court-circuiter les 2 fils au niveau chaudière
   - Si chaudière démarre : thermostat défectueux
   - Si chaudière ne démarre pas : problème chaudière ou câblage
   - Mesurer continuité contact thermostat (multimètre)
   - Contact fermé = demande chauffe (0 Ω)
   - Contact ouvert = pas de demande (infini)

3. **Contrôle câblage**
   - Vérifier continuité câble sur toute longueur
   - Contrôler serrage bornes (thermostat et chaudière)
   - Vérifier polarité si thermostat modulant
   - Contrôler absence court-circuit
   - Tester avec câble provisoire si doute

4. **Vérification piles thermostat**
   - Remplacer piles (même si affichage OK)
   - Piles faibles : affichage OK mais relais ne commute pas
   - Utiliser piles alcalines qualité (durée 1-2 ans)
   - Certains thermostats signalent piles faibles

5. **Contrôle positionnement thermostat**
   - **Position correcte :**
     - Pièce de vie principale (séjour)
     - Mur intérieur (pas façade)
     - Hauteur 1,50 m
     - Loin sources chaleur (radiateur, soleil, TV, lampe)
     - Loin courants d'air (porte, fenêtre)
     - Circulation air normale (pas derrière rideau/meuble)
   - **Mauvais positionnement :**
     - → Mesure température non représentative
     - → Régulation inadaptée

6. **Paramétrage chaudière**
   - Vérifier activation entrée thermostat (menu installateur)
   - Contrôler type thermostat déclaré (TOR/modulant)
   - Vérifier influence thermostat sur régulation
   - Paramètres selon marque :
     - Vaillant : paramètre 700 (fonction thermostat)
     - Saunier Duval : activation contact TA
     - Frisquet : déclaration Eco-Radio ou contact sec

7. **Conflit régulation climatique**
   - Si régulation climatique + thermostat :
     - Régler pondération (priorité climatique ou thermostat)
     - Mode cascade : climatique principal, thermostat limiteur
     - Éviter consignes contradictoires
   - Vérifier température départ (ne doit pas être trop basse)
   - Autoriser chaudière à chauffer selon météo

8. **Test thermostat modulant**
   - Vérifier communication (LED sur récepteur)
   - Contrôler messages erreur thermostat
   - Tester association thermostat/récepteur
   - Réinitialiser et ré-appairer si nécessaire
   - Vérifier compatibilité (OpenTherm, eBUS selon marque)

9. **Diagnostic avancé**
   - Consulter menu diagnostic chaudière
   - Vérifier état entrée thermostat (ouvert/fermé)
   - Observer changement état lors modification consigne
   - Si état change mais pas d'effet : paramétrage chaudière
   - Si état ne change pas : thermostat ou câblage

**Prévention :**
- Remplacement piles annuel (préventif)
- Vérification câblage lors entretien
- Test fonctionnement début saison chauffe
- Dépoussiérage thermostat
- Vérification positionnement (pas de nouveau meuble, rideau)

**Solutions selon diagnostic :**
- Thermostat HS : remplacement
- Câblage coupé : réparation ou remplacement câble
- Mauvais paramétrage : reconfiguration
- Mauvais positionnement : déplacement

**Choix thermostat :**
- **Simple TOR :** économique, fiable, basique
- **Programmable :** confort + économies (plages horaires)
- **Modulant OpenTherm :** optimisation rendement chaudière
- **Connecté WiFi :** pilotage distance, statistiques
- **Multi-zones :** confort pièce par pièce

**Économies thermostat programmable :**
- Réduction nuit : 15-20% économies
- Réduction absences : 10-15% économies
- Total potentiel : 20-30% vs thermostat fixe

---

## FACT-CHAUD-073: Communication thermostat/chaudière coupée

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-073 |
| **Catégorie** | Régulation & Sondes |
| **Système** | Communication bus (OpenTherm, eBUS, radio) |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques (systèmes modulants) |

**Symptômes :**
- Message erreur communication sur thermostat ou chaudière
- Thermostat connecté affiche "perte connexion"
- Chaudière fonctionne en mode dégradé (température fixe)
- Pas de retour info chaudière sur thermostat (température, pression)
- Pictogramme communication absent/barré

**Cause racine probable :**
Problème câblage bus, mauvaise polarité, perturbation électromagnétique, défaut alimentation, incompatibilité protocole, défaut module communication.

**Étapes de résolution :**

1. **Identification type communication**
   - **OpenTherm :** standard européen, 2 fils non polarisés
   - **eBUS :** Vaillant/Saunier Duval, 2 fils polarisés
   - **Propriétaire :** Frisquet Eco-Radio, Atlantic, etc.
   - **Radio 868 MHz :** sans fil, récepteur chaudière
   - **WiFi :** connecté internet, box/passerelle

2. **Diagnostic radio/WiFi**
   - **Thermostat radio :**
     - Vérifier piles émetteur (remplacer)
     - Contrôler LED récepteur (doit clignoter réception)
     - Vérifier distance émetteur/récepteur (< 30m général)
     - Contrôler obstacles (murs béton, métal)
     - Ré-appairer émetteur et récepteur (procédure constructeur)
   - **Thermostat WiFi :**
     - Vérifier connexion WiFi thermostat (SSID, mot de passe)
     - Contrôler signal WiFi emplacement thermostat
     - Vérifier connexion internet box
     - Redémarrer box internet et thermostat
     - Reconnecter application mobile

3. **Diagnostic bus filaire (OpenTherm/eBUS)**
   - Vérifier continuité 2 fils bus (multimètre)
   - Contrôler absence court-circuit fils entre eux ou masse
   - Vérifier serrage bornes thermostat et chaudière
   - Mesurer tension bus (doit varier : OpenTherm 0-40V, eBUS ~15-24V)
   - Longueur max câble : généralement 50m (vérifier notice)

4. **Contrôle polarité (eBUS)**
   - eBUS est polarisé : + et - à respecter
   - Inverser polarité si pas de communication
   - Repérer fils : généralement rouge (+) et noir/bleu (-)
   - Vérifier schéma notice chaudière et thermostat

5. **Perturbations électromagnétiques**
   - Câble bus éloigné câbles puissance (> 10 cm)
   - Pas de câble bus dans même gaine que 230V
   - Éloigner transformateurs, variateurs lumière
   - Utiliser câble blindé si environnement perturbé
   - Raccorder blindage à masse unique (éviter boucles)

6. **Contrôle alimentation**
   - Vérifier alimentation chaudière (secteur OK)
   - Contrôler alimentation piles thermostat
   - Sur thermostat alimenté par bus : vérifier tension fournie
   - Certains thermostats nécessitent alimentation séparée

7. **Compatibilité protocole**
   - Vérifier compatibilité thermostat/chaudière (notice)
   - OpenTherm : standard mais versions différentes (v2.2, v3.0)
   - eBUS : spécifique Vaillant/Saunier Duval
   - Certains thermostats multi-protocoles (réglage à effectuer)
   - Consulter liste compatibilité constructeur

8. **Réinitialisation et association**
   - **Réinitialiser thermostat :** procédure constructeur (reset usine)
   - **Réinitialiser module communication chaudière**
   - **Ré-appairer :**
     - Thermostat radio : mode appairage simultané
     - Thermostat WiFi : reconfiguration réseau
     - Bus filaire : généralement auto-détection
   - Patienter 2-5 minutes après association (initialisation)

9. **Diagnostic module communication**
   - Vérifier LED module récepteur chaudière
   - LED verte fixe : OK
   - LED rouge ou clignotante : erreur
   - Tester avec autre thermostat compatible si disponible
   - Remplacer module si défectueux

10. **Mode dégradé**
    - En attendant réparation : chaudière en mode standalone
    - Paramétrer température départ fixe (menu chaudière)
    - Ou installer thermostat TOR filaire temporaire

**Prévention :**
- Vérification annuelle communication (test)
- Remplacement préventif piles
- Contrôle câblage lors travaux
- Mise à jour firmware si disponible
- Protection câble bus (gaine, chemin protégé)

**Dépannage rapide selon symptôme :**

| Symptôme | Cause probable | Action |
|----------|----------------|--------|
| Perte liaison radio intermittente | Piles faibles | Remplacer piles |
| Perte liaison radio permanente | Désappairage | Ré-appairer |
| Pas de liaison bus filaire | Câble coupé ou polarité | Vérifier continuité et polarité |
| Communication OK puis perdue | Perturbation EMI | Éloigner câbles puissance |
| Liaison intermittente bus | Mauvais contact | Resserrer bornes |

**Avantages communication modulante (vs TOR) :**
- Régulation précise (modulation puissance)
- Retour infos chaudière (température, pression, erreurs)
- Optimisation rendement
- Programmation avancée
- Statistiques consommation

---

## FACT-CHAUD-074: Programmation horaire non respectée

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-074 |
| **Catégorie** | Régulation & Sondes |
| **Système** | Programmateur / Horloge |
| **Gravité** | Faible |
| **Marques** | Multi-marques |

**Symptômes :**
- Chauffage actif hors plages horaires programmées
- Chauffage ne démarre pas aux heures prévues
- Décalage horaire (heure été/hiver non prise en compte)
- Programme effacé ou modifié sans intervention
- Mode forcé activé en permanence

**Cause racine probable :**
Horloge mal réglée, programme incorrect, mode manuel activé, perte mémoire (pile horloge), conflit entre plusieurs dispositifs (thermostat + chaudière).

**Étapes de résolution :**

1. **Vérification heure et date**
   - Contrôler heure affichée chaudière/thermostat
   - Vérifier date (jour semaine important pour programme)
   - Corriger si décalage
   - Vérifier passage heure été/hiver automatique (activation)
   - Certains systèmes : mise à jour auto via internet (WiFi)

2. **Contrôle programmation**
   - Afficher programme horaire configuré
   - Vérifier plages confort/éco/hors-gel
   - Exemple type :
     - Lundi-Vendredi : 6h-8h et 18h-22h → Confort (20°C)
     - Lundi-Vendredi : 8h-18h et 22h-6h → Éco (17°C)
     - Samedi-Dimanche : 8h-23h → Confort
   - Corriger si incohérent

3. **Vérification mode fonctionnement**
   - Identifier mode actif :
     - **Auto/Programme :** suit programmation horaire
     - **Manuel/Permanent Confort :** température confort 24/7
     - **Manuel/Permanent Éco :** température réduite 24/7
     - **Hors-gel :** antigel uniquement (7-10°C)
     - **Arrêt :** pas de chauffage
   - Si mode manuel actif : repasser en mode Auto
   - Certains systèmes : dérogation temporaire (24-48h puis retour auto)

4. **Contrôle pile horloge**
   - Certaines chaudières : pile CR2032 pour horloge interne
   - Pile faible : perte heure lors coupure secteur
   - Remplacer pile si > 3-5 ans ou symptômes perte heure
   - Après remplacement : reprogrammer heure et date

5. **Conflit multi-dispositifs**
   - Si thermostat programmable + chaudière programmable :
     - **Configurer UN SEUL** programmateur (l'autre en mode permanent)
     - Recommandé : programmer thermostat, chaudière en mode chauffage permanent
     - Éviter double programmation (conflits)
   - Si régulation climatique : vérifier interaction avec programme

6. **Vérification fonction vacances**
   - Certains systèmes : mode vacances/absence
   - Maintient hors-gel ou température réduite période définie
   - Vérifier si activé par erreur
   - Désactiver si non souhaité

7. **Contrôle fonction optimisation**
   - Fonction "optimisation démarrage" : anticipe chauffe
   - Démarre avant heure programmée pour atteindre consigne à l'heure
   - Normal si chaudière démarre 30-60 min avant
   - Désactiver si non souhaité (perte efficacité)

8. **Réinitialisation programme**
   - Si programme corrompu : effacer et reprogrammer
   - Sauvegarder programme si fonction disponible
   - Certains thermostats : sauvegarde cloud (WiFi)
   - Procédure reset : voir notice constructeur

9. **Vérification fonctions intelligentes**
   - Thermostats connectés : détection présence/absence
   - Géolocalisation smartphone (mode away auto)
   - Apprentissage automatique habitudes
   - Vérifier paramétrage de ces fonctions
   - Désactiver si comportement non souhaité

**Prévention :**
- Vérification heure lors changement été/hiver
- Remplacement préventif pile horloge (tous les 5 ans)
- Sauvegarde programme (photo écran ou notice)
- Formation utilisateur (ne pas changer mode par erreur)
- Simplifier programmation (plus c'est simple, moins d'erreurs)

**Programmation type recommandée :**

**Occupation standard (travail journée) :**
- **6h-8h :** Confort (réveil, petit-déjeuner)
- **8h-17h :** Éco (absence)
- **17h-23h :** Confort (soirée)
- **23h-6h :** Éco/Nuit (sommeil)

**Optimisations :**
- Weekend : adapter selon habitudes
- Nuit : température nuit 16-17°C (économies + confort sommeil)
- Absences > 2 jours : hors-gel ou éco
- Mi-saison : réduire plages confort

**Températures recommandées :**
- **Confort jour :** 19-20°C (réglementation RT2012 : 19°C)
- **Réduit nuit :** 16-17°C
- **Éco absence courte :** 17-18°C
- **Absence longue :** 12-15°C
- **Hors-gel :** 7-10°C

**Économies programmation horaire :**
- Réduction 1°C : 7% économies
- Nuit 16°C vs 20°C : ~20% économies
- Programmation vs permanent : 10-30% économies selon occupation

**Systèmes programmation :**
- **Horloge chaudière :** basique, limité
- **Thermostat programmable :** plus flexible, interface meilleure
- **Thermostat connecté :** pilotage distance, géolocalisation, optimisation
- **Régulation multi-zones :** programmation pièce par pièce

---

## FACT-CHAUD-075: Mode été/hiver défaillant

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-075 |
| **Catégorie** | Régulation & Sondes |
| **Système** | Commutation été/hiver |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques (chaudières mixtes chauffage+ECS) |

**Symptômes :**
- Chauffage actif en été (hors demande)
- Pas de production ECS en mode hiver
- Commutation automatique ne fonctionne pas
- Mode été/hiver bloqué sur une position
- Chaudière ne suit pas sélection mode

**Cause racine probable :**
Vanne 3 voies bloquée, paramètre mode mal configuré, sonde extérieure défaillante (mode auto), servomoteur vanne défectueux, carte électronique.

**Étapes de résolution :**

1. **Compréhension modes**
   - **Mode Hiver :** chauffage + ECS actifs
   - **Mode Été :** ECS seule, chauffage désactivé
   - **Mode Auto :** commutation selon T°ext (si sonde extérieure)
   - Objectif été : économiser en n'activant pas chauffage

2. **Vérification sélection mode**
   - Contrôler mode sélectionné (afficheur chaudière/thermostat)
   - Tester changement manuel été ↔ hiver
   - Observer réaction chaudière (affichage, vanne 3 voies)
   - Vérifier cohérence saison/mode

3. **Contrôle vanne 3 voies (si présente)**
   - Localiser vanne 3 voies (généralement sortie chaudière)
   - Vérifier servomoteur (bruit rotation lors commutation)
   - Contrôler position vanne (indicateur mécanique)
   - Tester manuellement rotation (levier secours sur certains modèles)
   - Positions typiques :
     - Chauffage seul : flux vers circuit chauffage
     - ECS seule : flux vers échangeur ballon
     - Possible position intermédiaire (mixte)

4. **Test servomoteur vanne**
   - Forcer mode chauffage : servomoteur doit tourner
   - Forcer mode ECS : servomoteur doit tourner sens inverse
   - Écouter bruit moteur (ronronnement)
   - Si pas de mouvement : servomoteur HS ou câblage
   - Vérifier alimentation électrique servomoteur (230V ou 24V)

5. **Contrôle câblage servomoteur**
   - Vérifier connectique servomoteur (bien enfichée)
   - Mesurer tensions commande (multimètre)
   - Contrôler continuité câbles
   - Vérifier commutation tension lors changement mode

6. **Diagnostic vanne mécanique**
   - Si servomoteur OK mais pas d'effet : vanne grippée
   - Causes : entartrage, vieillissement joints, corrosion
   - Tenter rotation manuelle (avec précaution)
   - Démonter et nettoyer si accessible
   - Remplacer si bloquée définitivement

7. **Contrôle mode automatique**
   - Si mode auto été/hiver selon T°ext :
     - Vérifier sonde extérieure fonctionnelle (voir FACT-CHAUD-068)
     - Contrôler seuil commutation (menu installateur)
     - Seuil typique : 15-18°C extérieur
     - Si T°ext > seuil : mode été
     - Si T°ext < seuil : mode hiver
   - Ajuster seuil selon climat local

8. **Vérification priorité ECS**
   - En mode hiver : priorité ECS normale
   - Demande ECS → arrêt chauffage temporaire
   - Vanne 3 voies bascule vers ballon
   - Après satisfaction ECS → retour chauffage
   - Vérifier ce cycle fonctionne correctement

9. **Contrôle carte électronique**
   - Vérifier menu diagnostic (position vanne, mode actif)
   - Tester commande manuelle vanne si disponible
   - Contrôler relais carte (écouter clic commutation)
   - Remplacer carte si commande absente

10. **Paramétrage chaudière**
    - Vérifier activation fonction été/hiver (menu installateur)
    - Contrôler paramètres :
      - Type installation (chauffage seul, mixte, ECS seule)
      - Présence ballon ECS (oui/non)
      - Type vanne (3 voies, 2 vannes séparées)
    - Corriger si mal configuré

**Prévention :**
- Exercice vanne 3 voies hors saison (test été/hiver)
- Contrôle servomoteur lors entretien
- Vérification sonde extérieure (mode auto)
- Détartrage installation selon dureté eau

**Solutions selon diagnostic :**
- Servomoteur HS : remplacement (100-200€)
- Vanne grippée : nettoyage ou remplacement (200-400€)
- Carte électronique : remplacement module ou carte (150-500€)
- Paramétrage : reconfiguration (gratuit)

**Alternatives si panne :**
- Mode été : désactiver manuellement chauffage (commutateur/menu)
- Mode hiver : forcer chauffage actif en permanence
- Installer vanne manuelle temporaire (by-pass)

**Systèmes sans vanne 3 voies :**
- Chaudière chauffage seul + ballon électrique ECS : pas de commutation
- Chaudière ECS instantanée : pas de vanne, gestion électronique
- Systèmes séparés : 2 circulateurs (chauffage + ECS)

**Optimisation été :**
- Arrêt complet chauffage (économies)
- Réduction température ECS possible (55°C mini légionelle)
- Certaines chaudières : arrêt ventilateur été (économie électrique)
- Entretien chaudière profiter de l'été (plus disponible)

**Consommation mode été :**
- Chaudière gaz ECS : ~30-50 m³/mois pour 4 personnes
- Préchauffage chaudière avant puisage : pertes 10-15%
- Isolation ballon importante (pertes statiques)

---

## FACT-CHAUD-076: Anticipation surchauffe

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-076 |
| **Catégorie** | Régulation & Sondes |
| **Système** | Régulation température |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques |

**Symptômes :**
- Température départ dépasse largement consigne
- Température intérieure atteint consigne puis continue à monter
- Chaudière ne s'arrête pas à temps
- Surchauffe de 2-5°C après arrêt chaudière
- Inconfort thermique (trop chaud)

**Cause racine probable :**
Inertie système non compensée, régulation trop lente, paramètres PID inadaptés, absence d'anticipation, sur-dimensionnement chaudière, mauvaise régulation hydraulique.

**Étapes de résolution :**

1. **Compréhension phénomène**
   - **Inertie thermique :** système continue à chauffer après arrêt
   - Sources inertie :
     - Masse eau circuit (volume important)
     - Masse émetteurs (radiateurs fonte)
     - Masse bâtiment (murs, dalles)
   - Chaleur stockée est restituée → dépassement consigne

2. **Mesure dépassement**
   - Noter température consigne (ex: 20°C)
   - Mesurer température max atteinte (ex: 23°C)
   - Calculer dépassement (ici: +3°C)
   - Observer temps retour à consigne après arrêt
   - Identifier si problème ponctuel ou systématique

3. **Contrôle type régulation**
   - **Régulation TOR simple :** pas d'anticipation
     - Marche/arrêt basique
     - Dépassement normal (1-2°C acceptable)
     - Si > 2°C : problème autre
   - **Régulation PID :** anticipation possible
     - Paramètres P, I, D à ajuster
     - Peut limiter dépassement à 0,5°C
   - **Régulation climatique :** anticipation météo
     - Devrait éviter dépassements importants

4. **Optimisation régulation TOR**
   - **Différentiel/hystérésis :**
     - Écart entre seuil marche et arrêt
     - Exemple : arrêt à 20°C, redémarrage à 19°C (différentiel 1°C)
     - Si dépassement : réduire différentiel (attention cycles courts)
   - **Anticipation arrêt :**
     - Certains thermostats : fonction anticipation
     - Arrêt avant consigne (ex: -0,5°C) pour compenser inertie
     - Activer/ajuster si disponible

5. **Ajustement régulation PID**
   - **Paramètre P (Proportionnel) :**
     - Réduit puissance à l'approche consigne
     - Augmenter P si dépassement (action plus forte)
   - **Paramètre I (Intégral) :**
     - Corrige erreur résiduelle
     - Augmenter I si oscillations
   - **Paramètre D (Dérivé) :**
     - Anticipe évolution
     - Augmenter D pour limiter dépassement
   - **Attention :** ajustements par petits pas, observer 24-48h entre modifications

6. **Contrôle dimensionnement chaudière**
   - Chaudière surdimensionnée : montée rapide température
   - Vérifier puissance chaudière vs besoins
   - Si > 150% besoins : risque surchauffe
   - Solutions :
     - Réduire puissance max chaudière (menu)
     - Améliorer régulation (modulante)
     - Installer ballon tampon (stockage, amortissement)

7. **Optimisation hydraulique**
   - **Débit circulation :**
     - Débit excessif : transport rapide chaleur, dépassement
     - Réduire vitesse pompe (si réglable)
     - Viser ΔT départ-retour 15-20°C
   - **Bypass :**
     - Vérifier réglage bypass
     - Bypass trop ouvert : débit réduit émetteurs, accumulation chaleur chaudière

8. **Solutions émetteurs**
   - **Vannes thermostatiques radiateurs :**
     - Installer si absentes
     - Fermeture progressive à l'approche consigne
     - Limite dépassement local
   - **Têtes thermostatiques anticipation :**
     - Détectent montée température
     - Ferment avant atteinte consigne
   - **Régulation pièce par pièce :**
     - Évite surchauffe globale

9. **Inertie bâtiment**
   - **Bâtiment lourd (béton, pierre) :**
     - Forte inertie : dépassement normal
     - Augmenter anticipation régulation
     - Programmer arrêt plus tôt
   - **Bâtiment léger (ossature bois) :**
     - Faible inertie : réaction rapide
     - Dépassement anormal si se produit

10. **Fonction optimisation**
    - Certains thermostats : apprentissage automatique
    - Analysent comportement thermique
    - Ajustent automatiquement anticipation
    - Activer si disponible (ex: Nest, Netatmo)

**Prévention :**
- Dimensionnement correct chaudière dès installation
- Régulation adaptée (PID préférable à TOR)
- Vannes thermostatiques radiateurs
- Ajustement paramètres après travaux isolation (inertie modifiée)

**Dépassement acceptable selon système :**
| Système | Dépassement acceptable |
|---------|------------------------|
| TOR simple | 1-2°C |
| PID bien réglé | 0,5-1°C |
| Régulation climatique | < 0,5°C |
| Plancher chauffant | 0,5°C (forte inertie mais régulation lente) |
| Radiateurs fonte | 1-2°C (inertie émetteurs) |
| Radiateurs alu | 0,5-1°C (faible inertie) |

**Impact surchauffe :**
- Inconfort occupants (transpiration, ouverture fenêtres)
- Surconsommation énergétique (gaspillage)
- Usure installation (cycles thermiques)
- Air intérieur sec (inconfort, santé)

**Solutions court terme :**
- Réduire consigne -1°C (compenser dépassement)
- Programmer arrêt anticipé
- Régler têtes thermostatiques plus bas

**Solutions long terme :**
- Régulation performante (PID, OpenTherm)
- Chaudière modulante correctement dimensionnée
- Ballon tampon si nécessaire
- Régulation pièce par pièce

---

## FACT-CHAUD-077: Cycles marche/arrêt trop fréquents

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-077 |
| **Catégorie** | Régulation & Sondes |
| **Système** | Régulation cycles |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques |

**Symptômes :**
- Chaudière démarre et s'arrête toutes les 2-5 minutes
- Bruit cycles allumage répétés
- Brûleur ne tient pas en fonctionnement
- Code erreur cycles courts (certaines marques)
- Usure prématurée composants

**Cause racine probable :**
Sur-dimensionnement chaudière, absence anti-cycles courts, régulation inadaptée, débit hydraulique insuffisant, thermostat mal positionné, installation non équilibrée.

**Étapes de résolution :**

1. **Mesure fréquence cycles**
   - Chronométrer cycles marche/arrêt
   - Noter durée marche et durée arrêt
   - **Cycles courts = marche < 3 minutes**
   - Compter nombre cycles/heure
   - > 6 cycles/heure = problématique

2. **Identification cause principale**
   - **Sur-dimensionnement :**
     - Chaudière trop puissante pour besoins
     - Atteint température rapidement
     - S'arrête, refroidit, redémarre
   - **Différentiel trop faible :**
     - Écart marche/arrêt insuffisant
     - Oscillation rapide autour consigne
   - **Débit insuffisant :**
     - Chaleur non évacuée
     - Surchauffe rapide chaudière

3. **Contrôle dimensionnement**
   - Calculer besoins réels (déperditions bâtiment)
   - Comparer puissance chaudière
   - Exemple problématique : chaudière 25 kW pour besoin 10 kW
   - Sur-dimensionnement > 150% : problème cycles courts

4. **Réduction puissance chaudière**
   - Régler puissance maximale dans menu installateur
   - Paramètre "Puissance max chauffage"
   - Réduire progressivement (ex: de 100% à 60%)
   - Tester et ajuster jusqu'à cycles > 5 minutes
   - Vérifier capacité chauffage grand froid

5. **Ajustement différentiel/hystérésis**
   - Augmenter différentiel thermostat
   - Exemple : 0,5°C → 1,5°C
   - Compromis : cycles plus longs mais variation température
   - Différentiel 1-2°C : bon équilibre général

6. **Temporisation anti-cycles courts**
   - Paramètre chaudière : temps mini arrêt
   - Empêche redémarrage immédiat
   - Régler à 3-5 minutes minimum
   - Certaines chaudières : fonction "anti-tact" intégrée

7. **Contrôle débit hydraulique**
   - Vérifier vitesse pompe circulation (pas trop lente)
   - Mesurer ΔT départ-retour :
     - ΔT faible (< 10°C) : débit excessif ou puissance faible
     - ΔT élevé (> 25°C) : débit insuffisant (cycles courts)
   - Ajuster vitesse pompe
   - Objectif : ΔT 15-20°C

8. **Équilibrage installation**
   - Vérifier ouverture vannes radiateurs
   - Équilibrer réseau (débit homogène)
   - Fermer bypass si trop ouvert
   - Purger air circuit (air = débit réduit)

9. **Vérification thermostat**
   - Contrôler positionnement (voir FACT-CHAUD-072)
   - Thermostat sur radiateur ou source chaleur : cycles courts
   - Déplacer en position neutre
   - Vérifier différentiel thermostat

10. **Solution ballon tampon**
    - Si impossible autre solution : installer ballon tampon
    - Volume 50-100L selon puissance chaudière
    - Stocke chaleur produite
    - Lisse cycles, allonge durée fonctionnement
    - Particulièrement utile chaudière bois, PAC

11. **Optimisation régulation**
    - Passer en régulation climatique si possible
    - Modulation puissance (chaudière modulante)
    - Régulation anticipative
    - Évite cycles tout ou rien

12. **Vérification paramètres avancés**
    - Temps minimal fonctionnement brûleur
    - Temps minimal arrêt
    - Gradient montée puissance (rampe)
    - Hystérésis températures

**Prévention :**
- Dimensionnement correct dès installation
- Calcul déperditions sérieux (bureau étude)
- Régulation adaptée
- Entretien régulier (débit optimal)

**Conséquences cycles courts :**
- **Usure prématurée :**
  - Électrode allumage
  - Vanne gaz
  - Ventilateur
  - Échangeur (chocs thermiques)
- **Surconsommation :**
  - Rendement dégradé (phase démarrage)
  - Pertes balayage pré-ventilation
- **Nuisance sonore :** "tac-tac-tac" permanent
- **Inconfort :** température instable

**Cycles normaux selon système :**
| Type installation | Durée cycle marche | Cycles/heure |
|-------------------|-------------------|--------------|
| Idéal | 10-15 min | 2-4 |
| Acceptable | 5-10 min | 4-6 |
| Problématique | < 5 min | > 6 |
| Critique | < 3 min | > 10 |

**Solutions selon cause :**
| Cause | Solution |
|-------|----------|
| Sur-dimensionnement | Réduire puissance max, ballon tampon |
| Différentiel faible | Augmenter hystérésis |
| Débit insuffisant | Augmenter vitesse pompe, purger |
| Thermostat mal placé | Déplacer en position neutre |
| Absence anti-cycles | Activer temporisation |

**Chaudière modulante vs TOR :**
- **TOR (Tout Ou Rien) :** plus sensible cycles courts
- **Modulante :** adapte puissance, moins de cycles
- Modulation 30-100% : amélioration notable
- Modulation 10-100% : cycles quasi éliminés

---

## FACT-CHAUD-078: Température départ trop élevée/faible

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-078 |
| **Catégorie** | Régulation & Sondes |
| **Système** | Régulation température départ |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques |

**Symptômes :**
- **Température trop élevée :** surchauffe locaux, inconfort, surconsommation, cycles courts
- **Température trop faible :** locaux froids, inconfort, plainte occupants
- Température départ ne correspond pas à consigne
- Température départ inadaptée aux besoins

**Cause racine probable :**
Consigne mal réglée, courbe chauffe inadaptée, sonde température défectueuse, régulation défaillante, sur/sous-dimensionnement émetteurs.

**Étapes de résolution - Température TROP ÉLEVÉE :**

1. **Diagnostic situation**
   - Mesurer température départ réelle (thermomètre)
   - Comparer avec consigne affichée
   - Mesurer température intérieure
   - Relever température extérieure
   - Identifier si problème permanent ou contextuel

2. **Vérification consigne départ**
   - Consulter consigne température départ (menu chaudière)
   - **Valeurs normales selon émetteurs :**
     - Plancher chauffant : 35-45°C
     - Radiateurs BT : 45-55°C
     - Radiateurs MT : 55-65°C
     - Radiateurs HT : 70-80°C
   - Réduire consigne si trop élevée

3. **Contrôle courbe de chauffe**
   - Si régulation climatique active
   - Vérifier pente et décalage (voir FACT-CHAUD-070, 071)
   - Pente trop forte → température départ excessive
   - Réduire pente ou décalage parallèle

4. **Vérification sonde départ**
   - Température affichée vs réelle
   - Si sonde indique trop froid → chaudière chauffe trop
   - Tester sonde (voir FACT-CHAUD-066)
   - Remplacer si défectueuse

5. **Contrôle modulation**
   - Chaudière doit moduler selon besoin
   - Vérifier modulation active (menu diagnostic)
   - Si bloquée puissance max : problème vanne gaz ou carte

6. **Solutions température trop élevée**
   - Réduire consigne température départ
   - Ajuster courbe de chauffe (pente, décalage)
   - Activer régulation climatique si disponible
   - Installer vannes thermostatiques radiateurs
   - Réduire puissance max chaudière si surdimensionnée

**Étapes de résolution - Température TROP FAIBLE :**

1. **Diagnostic situation**
   - Mesurer température départ
   - Comparer avec consigne
   - Vérifier température intérieure (toutes pièces)
   - Noter température extérieure
   - Identifier radiateurs tièdes ou froids

2. **Vérification consigne départ**
   - Consulter consigne
   - Augmenter si inférieure aux besoins
   - Vérifier consigne max autorisée (limite protection)

3. **Contrôle courbe de chauffe**
   - Pente trop faible → température départ insuffisante froid
   - Augmenter pente ou décalage parallèle
   - Vérifier température extérieure cohérente (sonde OK)

4. **Vérification sonde départ**
   - Si sonde indique trop chaud → chaudière ne chauffe pas assez
   - Exemple : 70°C affiché mais 50°C réel
   - Tester et remplacer sonde si défectueuse

5. **Contrôle puissance chaudière**
   - Vérifier que chaudière atteint puissance max
   - Consulter menu diagnostic (% puissance)
   - Si limitée : problème gaz, vanne, modulation

6. **Vérification hydraulique**
   - Pompe circulation fonctionne (vitesse correcte)
   - Débit suffisant (pas d'air, pas d'obstruction)
   - Vanne mélangeuse ouverte (si présente)
   - By-pass pas trop ouvert

7. **Contrôle émetteurs**
   - Radiateurs correctement dimensionnés
   - Tous radiateurs chauffent
   - Vannes ouvertes
   - Radiateurs non entartrés (si eau dure)
   - Pas de poche d'air

8. **Solutions température trop faible**
   - Augmenter consigne température départ
   - Ajuster courbe de chauffe (pente, décalage)
   - Augmenter puissance max chaudière
   - Vérifier dimensionnement installation
   - Améliorer isolation (réduire besoins)
   - Envisager émetteurs supplémentaires

**Adaptation température selon conditions :**

**Grand froid (T°ext < -5°C) :**
- Température départ maximale nécessaire
- Vérifier que chaudière atteint consigne

**Mi-saison (T°ext 5-15°C) :**
- Température départ réduite
- Éviter surchauffe
- Régulation climatique optimale

**Nuit/réduit :**
- Réduction température départ possible
- Ou arrêt complet selon stratégie

**ECS prioritaire :**
- Température départ peut baisser temporairement (vanne 3 voies)
- Normal, retour chauffage après production ECS

**Prévention :**
- Réglage minutieux mise en service
- Vérification saisonnière (automne)
- Contrôle sondes température annuel
- Ajustement après travaux isolation/émetteurs

**Tableau températures recommandées :**
| Émetteurs | T°ext -10°C | T°ext 0°C | T°ext +10°C |
|-----------|-------------|-----------|-------------|
| Plancher chauffant | 40°C | 35°C | 30°C |
| Radiateurs BT | 50°C | 45°C | 35°C |
| Radiateurs MT | 65°C | 55°C | 45°C |
| Radiateurs HT | 75°C | 65°C | 50°C |

**Impact température départ :**
- **Trop élevée :**
  - Surconsommation (+10-30%)
  - Inconfort (surchauffe)
  - Cycles courts
  - Perte condensation (chaudières condensation)
- **Trop faible :**
  - Inconfort (froid)
  - Insatisfaction occupants
  - Humidité (pas assez chaud pour sécher)

**Optimisation condensation :**
- Chaudière condensation : T°retour < 54°C
- Baisser T°départ favorise condensation
- Rendement optimal : T°départ < 55°C
- Compromis confort/rendement à trouver

---

## FACT-CHAUD-079: Différentiel température mal réglé

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-079 |
| **Catégorie** | Régulation & Sondes |
| **Système** | Différentiel régulation |
| **Gravité** | Faible |
| **Marques** | Multi-marques |

**Symptômes :**
- Cycles marche/arrêt trop fréquents (différentiel trop faible)
- Variation température intérieure excessive (différentiel trop élevé)
- Inconfort thermique (oscillations température)
- Chaudière ne redémarre pas assez vite (différentiel trop grand)

**Cause racine probable :**
Paramètre différentiel/hystérésis mal configuré, valeur usine inadaptée installation, absence réglage lors mise en service.

**Étapes de résolution :**

1. **Compréhension différentiel**
   - **Différentiel = écart entre seuil marche et seuil arrêt**
   - **Exemple :**
     - Consigne 20°C, différentiel 1°C
     - Arrêt chauffage à 20°C
     - Redémarrage chauffage à 19°C
   - **Hystérésis = autre terme pour différentiel**

2. **Identification problème**
   - **Différentiel trop faible (< 0,5°C) :**
     - Cycles marche/arrêt très fréquents
     - Usure composants
     - Nuisance sonore
     - Surconsommation
   - **Différentiel trop élevé (> 2°C) :**
     - Variation température ressentie
     - Inconfort (trop chaud puis trop froid)
     - Cycles longs (économique mais moins confortable)

3. **Mesure différentiel actuel**
   - Observer température départ ou intérieure
   - Noter température arrêt chaudière (T_stop)
   - Noter température redémarrage (T_start)
   - Calculer : Différentiel = T_stop - T_start
   - Exemple : arrêt 20°C, redémarrage 18,5°C → différentiel 1,5°C

4. **Localisation réglage différentiel**
   - **Thermostat d'ambiance :**
     - Menu paramètres/réglages
     - Souvent nommé "Hystérésis", "Différentiel", "Écart"
   - **Chaudière :**
     - Menu installateur
     - Paramètre "Hystérésis température"
   - **Régulation climatique :**
     - Parfois différentiel fixe (non modifiable)

5. **Réglage optimal différentiel**
   - **Valeurs recommandées :**
     - **Thermostat ambiance classique :** 0,5-1,5°C
     - **Thermostat précis/modulant :** 0,3-0,5°C
     - **Régulation climatique :** 0,5-1°C
     - **Plancher chauffant :** 0,5-1°C (forte inertie)
     - **Radiateurs :** 1-2°C (acceptable)
   - **Compromis :**
     - Faible différentiel : confort max, risque cycles courts
     - Différentiel élevé : cycles longs, confort réduit

6. **Ajustement progressif**
   - Partir valeur actuelle
   - Modifier par incrément 0,2-0,5°C
   - Observer 24-48h
   - Mesurer :
     - Fréquence cycles
     - Variation température ressentie
     - Confort occupants
   - Réajuster si nécessaire

7. **Adaptation selon inertie**
   - **Faible inertie (bâtiment léger, radiateurs alu) :**
     - Réaction rapide
     - Différentiel faible possible (0,5-1°C)
   - **Forte inertie (bâtiment lourd, plancher chauffant) :**
     - Réaction lente
     - Différentiel plus élevé acceptable (1-2°C)
     - Inertie compense variations

8. **Différentiel sur température départ**
   - Certaines chaudières : différentiel sur T°départ (pas T°ambiance)
   - Valeurs plus élevées : 5-10°C typique
   - Exemple : consigne 60°C, arrêt à 60°C, redémarrage à 50°C
   - Évite cycles courts sur production chaleur

9. **Cas particulier régulation modulante**
   - Thermostat modulant (OpenTherm) : différentiel moins critique
   - Modulation progressive puissance
   - Évite marche/arrêt franc
   - Différentiel peut être faible (0,3-0,5°C)

10. **Vérification anti-cycles courts**
    - En complément différentiel : temporisation
    - Temps minimum arrêt avant redémarrage
    - Paramètre chaudière "Temps anti-tact"
    - Typique : 3-5 minutes
    - Combine différentiel + temporisation = régulation optimale

**Prévention :**
- Réglage lors mise en service (professionnel)
- Réévaluation après modification installation
- Ajustement saisonnier si nécessaire
- Documentation valeur retenue

**Exemples configuration :**

**Configuration confort max (faibles variations) :**
- Différentiel : 0,5°C
- Anti-cycles : 5 minutes
- Régulation modulante recommandée

**Configuration économique (cycles longs) :**
- Différentiel : 1,5-2°C
- Anti-cycles : 3 minutes
- Acceptable si inertie bâtiment

**Configuration équilibrée :**
- Différentiel : 1°C
- Anti-cycles : 3-4 minutes
- Bon compromis général

**Interaction autres paramètres :**
- Différentiel + Anticipation sonde = régulation fine
- Différentiel + Courbe chauffe = optimisation globale
- Différentiel + Programmation horaire = confort programmé

**Mesure confort :**
- Variation température < 1°C : excellent
- Variation 1-2°C : bon
- Variation 2-3°C : acceptable
- Variation > 3°C : inconfortable, revoir réglages

**Relation différentiel/cycles :**
| Différentiel | Cycles/heure (typique) | Confort |
|--------------|------------------------|---------|
| 0,3°C | 6-8 | Excellent |
| 0,5°C | 4-6 | Très bon |
| 1°C | 3-4 | Bon |
| 1,5°C | 2-3 | Correct |
| 2°C | 1-2 | Acceptable |
| > 2°C | < 1 | Variations perceptibles |

---

## FACT-CHAUD-080: Capteur pression différentielle défectueux

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-080 |
| **Catégorie** | Régulation & Sondes |
| **Système** | Capteur ΔP |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques (chaudières récentes, haute performance) |

**Symptômes :**
- Code erreur débit/pression différentielle
- Chaudière ne démarre pas (pas de débit détecté)
- Modulation puissance incohérente
- Régulation hydraulique défaillante
- Protection manque d'eau activée sans raison

**Cause racine probable :**
Capteur ΔP défectueux, prises de pression bouchées, membrane capteur encrassée, câblage défectueux, défaut carte électronique.

**Étapes de résolution :**

1. **Compréhension rôle capteur ΔP**
   - **ΔP = différence pression entre 2 points**
   - Mesure généralement : pression départ - pression retour
   - Fonction :
     - Détection circulation eau (débit)
     - Calcul débit volumique
     - Protection manque d'eau
     - Optimisation modulation selon débit
   - Remplace pressostat eau sur chaudières récentes

2. **Localisation capteur**
   - Généralement sur corps de chauffe
   - 2 tubes de prise pression (départ et retour)
   - Capteur électronique (signal 0-10V ou 4-20mA)
   - Connecté à carte électronique

3. **Diagnostic initial**
   - Consulter menu diagnostic chaudière
   - Relever valeur ΔP affichée
   - Comparer avec valeur attendue
   - Valeur typique fonctionnement : 100-300 mbar (varie selon modèle)
   - Pompe arrêtée : ΔP ≈ 0
   - Pompe en marche : ΔP doit augmenter

4. **Test fonctionnement pompe**
   - Vérifier pompe circulation fonctionne
   - Écouter bruit pompe
   - Vérifier rotation (rotor non bloqué)
   - Sans pompe : pas de ΔP (normal)

5. **Contrôle prises de pression**
   - Localiser 2 tubes prise pression (petits tubes 4-6mm)
   - Vérifier non bouchés :
     - Déconnecter du capteur
     - Souffler légèrement (doivent être dégagés)
     - Nettoyer si bouchés (aiguille fine, air comprimé)
   - Vérifier non pincés ou percés
   - Reconnecter solidement

6. **Contrôle membrane capteur**
   - Certains capteurs : membrane mesure ΔP
   - Membrane peut s'encrasser (tartre, boues)
   - Démonter capteur (couper pression, vidanger localement)
   - Nettoyer membrane délicatement (eau, brosse douce)
   - Rincer abondamment
   - Remonter avec joints neufs

7. **Test électrique capteur**
   - Mesurer alimentation capteur (multimètre)
   - Typique : 5V ou 12V DC
   - Mesurer signal sortie capteur (tension ou courant)
   - Exemple : 0-10V proportionnel à ΔP
   - Vérifier variation signal si on modifie débit (vitesse pompe)
   - Si signal fixe ou absent : capteur HS

8. **Contrôle câblage**
   - Vérifier connectique capteur (bien enfichée)
   - Contrôler continuité câbles (pas de coupure)
   - Vérifier absence court-circuit
   - Contrôler connexions carte électronique

9. **Test croisé**
   - Si possible, tester avec capteur identique (autre chaudière)
   - Ou simuler signal capteur (générateur tension/courant)
   - Si chaudière fonctionne avec autre capteur : capteur HS
   - Si problème persiste : carte électronique

10. **Remplacement capteur**
    - Identifier référence exacte (notice technique)
    - Commander pièce origine constructeur
    - Procédure remplacement :
      - Couper électricité et eau
      - Vidanger localement (circuit chaudière)
      - Déconnecter tubes pression
      - Déconnecter câble électrique
      - Dévisser capteur
      - Monter nouveau capteur (joints neufs, pâte d'étanchéité)
      - Reconnecter tubes et câble
      - Remplir et purger
      - Tester fonctionnement

11. **Paramétrage après remplacement**
    - Certains systèmes : calibration nécessaire
    - Procédure menu installateur
    - Vérifier valeur ΔP cohérente
    - Tester modulation puissance
    - Vérifier sécurités (arrêt si pompe HS)

12. **Contrôle installation hydraulique**
    - Vérifier pression circuit (1-1,5 bar)
    - Purger air (air = faux débit)
    - Contrôler filtre (pas bouché)
    - Vérifier vase expansion (pas de membrane percée)

**Prévention :**
- Vérification annuelle valeur ΔP (menu diagnostic)
- Nettoyage prises pression tous les 2-3 ans
- Traitement eau circuit (éviter encrassement)
- Contrôle câblage lors entretien

**Valeurs ΔP typiques :**
| Situation | ΔP attendu |
|-----------|------------|
| Pompe arrêtée | 0-20 mbar |
| Pompe vitesse 1 | 50-150 mbar |
| Pompe vitesse 2 | 100-250 mbar |
| Pompe vitesse 3 | 150-400 mbar |
| Filtre bouché | ΔP excessive (> 500 mbar) |

**Diagnostic selon valeur ΔP :**
- **ΔP = 0 en permanence :** capteur HS, pompe arrêtée, tubes bouchés
- **ΔP excessive :** filtre bouché, obstruction circuit, vanne fermée
- **ΔP erratique :** air dans circuit, capteur défectueux
- **ΔP ne varie pas :** capteur bloqué, membrane encrassée

**Codes erreur fréquents :**
- Vaillant F22 : manque d'eau (ΔP insuffisant)
- Saunier Duval F34 : défaut pression différentielle
- Frisquet : sécurité manque d'eau (ΔP)

**Fonction avancée capteur ΔP :**
- **Calcul débit :** Q = k × √ΔP (k = constante installation)
- **Optimisation modulation :** adapter puissance au débit
- **Détection anomalies :** filtre bouché, vanne fermée
- **Protection hydraulique :** arrêt si débit nul
- **Efficacité énergétique :** meilleur rendement

**Coût remplacement :**
- Capteur ΔP : 80-200€ selon modèle
- Main d'œuvre : 1-2h
- Total intervention : 150-400€

**Alternative si panne :**
- Certaines chaudières : mode dégradé sans capteur ΔP
- Ou by-pass temporaire (shunt sécurité, avec précautions)
- **DANGER :** ne pas neutraliser sécurité définitivement (risque surchauffe)

---

*Fin du fichier 05_Regulation_Sondes.md*

**Retour à l'index :** [Knowledge_Base_Chaudieres.md](Knowledge_Base_Chaudieres.md)
