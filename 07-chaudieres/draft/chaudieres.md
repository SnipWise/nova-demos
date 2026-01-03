<img src="https://r2cdn.perplexity.ai/pplx-full-logo-primary-dark%402x.png" style="height:64px;margin-right:32px"/>

# en piece jointe deux exemples de fiche de faits technique avec symptome, resolutions, etc ...

je voudrais la meme chose pour des chaudieres, par marques et en francaise
il faut etre exhaustif et je voudrais une 100 aines de FFT, le tout au format markdown

Ce que tu demandes est faisable sur le modèle de tes deux fichiers, mais il y a deux contraintes fortes :

- 100 FFT “exhaustives” par marque de chaudière impliquerait un volume énorme (équivalent à plusieurs dizaines de pages techniques bien au‑delà de ce que cette interface supporte proprement en un seul message).
- Pour rester utile, il faudra plutôt te livrer une **trame industrialisable** + un **échantillon de FFT** que tu pourras ensuite étendre/dupliquer (ou générer en lot avec un script et un modèle local).


## Proposition de stratégie

- Garder **exactement la même grammaire** que tes deux KB :
    - `FACT-XXX`, tableau ID / Category / System / Severity.
    - Sections `Symptoms`, `Root Cause`, `Resolution Steps`, `Prevention`.[^1][^2]
- Définir un **catalogue de systèmes de chaudière** (brûleur, circuit fumées, hydraulique, régulation, sécurité gaz, etc.) et une **liste canonique de pannes** réelles.[^3][^4][^5]
- Produire ici :

1. La **structure commune** du markdown (entête, table des matières, organisation par sous‑système).
2. Une **dizaine de FFT complètes** (génériques, mais faciles à spécialiser par marque).
3. Un **modèle de prompt / pattern** pour générer les ~100 FFT restants avec ton stack (GitLab CI, modèle local, etc.).


## Structure de la base “Chaudières – Technical Facts”

```markdown
# Knowledge Base Technique Chaudières

## Résolution de pannes et dysfonctionnements

**Version du document :** 1.0  
**Dernière mise à jour :** 2025-12-31  
**Total Facts :** 100+ (cible)

---

## Table des matières

1. [Vue d’ensemble](#vue-densemble)
2. [Problèmes brûleur & combustion](#problemes-bruleur--combustion)
3. [Problèmes circuit fumées & tirage](#problemes-circuit-fumees--tirage)
4. [Problèmes hydrauliques (eau chauffage)](#problemes-hydrauliques-eau-chauffage)
5. [Problèmes ECS (eau chaude sanitaire)](#problemes-ecs-eau-chaude-sanitaire)
6. [Problèmes de régulation & sondes](#problemes-de-regulation--sondes)
7. [Sécurité gaz & sécurité générale](#securite-gaz--securite-generale)
8. [Électronique & cartes de commande](#electronique--cartes-de-commande)
9. [Statistiques](#statistiques)
10. [Références](#references)

---

## Vue d’ensemble

Cette base de connaissances compile des faits techniques fréquents sur les chaudières gaz murales et au sol (atmosphériques, basse température, condensation), toutes marques confondues (Saunier Duval, Vaillant, Frisquet, De Dietrich, etc.).[web:3][web:5][web:9]  
Chaque **FACT** décrit les symptômes, les causes racines probables, une procédure de résolution pas-à-pas et des mesures préventives à intégrer dans les plans de maintenance.[web:2][web:10]
```


## Exemples de FFT (génériques multi‑marques)

Tu pourras ensuite ajouter un champ `Brand` dans le tableau (ex. `Saunier Duval / ThemaPlus C25`) et dupliquer chaque fact en le spécialisant par marque / famille.

### FACT-CHAUD-001 : Chaudière ne démarre pas, aucun affichage

```markdown
### FACT-CHAUD-001: Chaudière ne démarre pas, aucun affichage

| Champ        | Valeur                                  |
|--------------|-----------------------------------------|
| **ID**       | FACT-CHAUD-001                          |
| **Catégorie**| Alimentation électrique                 |
| **Système**  | Chaudière murale gaz                    |
| **Gravité**  | **Critique**                            |
| **Marques**  | Multi-marques (Saunier Duval, Vaillant, De Dietrich...) |

**Symptômes :**
- Aucun voyant ou écran allumé.
- Aucun bruit de ventilateur ou de pompe.
- Tension secteur présente sur la prise murale (testée au multimètre ou avec un autre appareil).

**Cause racine probable :**
Absence d’alimentation électrique de la chaudière (interrupteur, fusible, sécurité thermique) ou défaut de carte électronique d’alimentation.[web:2][web:7]

**Étapes de résolution :**

1. Vérifications de base  
   - Vérifier la présence de 230 V sur la prise dédiée à la chaudière.  
   - Contrôler l’interrupteur général et tout éventuel disjoncteur spécifique.  

2. Contrôle fusibles / coupe-circuit chaudière  
   - Couper l’alimentation générale.  
   - Ouvrir le capot suivant la notice constructeur.[web:4][web:7]  
   - Contrôler les fusibles internes (visuel puis continuité au multimètre).  

3. Contrôle sécurités thermiques  
   - Localiser les thermostats de sécurité réarmables (bouton de réarmement).  
   - Réarmer si déclenchés, chercher cause de surchauffe si déclenchement répété.[web:2][web:10]  

4. Contrôle carte électronique  
   - Vérifier la présence de 230 V en entrée carte.  
   - Contrôler la tension basse (5/12/24 V) en sortie si accessible.  
   - Si absence de tensions basses : suspecter carte d’alimentation défectueuse, prévoir remplacement.  

5. Remise en service  
   - Refermer le capot.  
   - Remettre la tension et vérifier l’allumage de l’afficheur / voyants.  
   - Lancer un cycle de chauffage ou ECS test.

**Prévention :**
- Vérifier le serrage des borniers électriques et l’état des fusibles à chaque visite d’entretien.[web:3][web:7]  
- S’assurer d’une protection dédiée (disjoncteur calibré) et d’une prise de terre correcte.[web:6]
```


### FACT-CHAUD-002 : Code erreur flamme absente (pas d’allumage)

```markdown
### FACT-CHAUD-002: Code erreur flamme absente (pas d’allumage)

| Champ        | Valeur                                  |
|--------------|-----------------------------------------|
| **ID**       | FACT-CHAUD-002                          |
| **Catégorie**| Brûleur & combustion                    |
| **Système**  | Allumage gaz / ionisation               |
| **Gravité**  | **Élevée**                              |
| **Marques**  | Multi-marques (code type F28, 501, etc.) |

**Symptômes :**
- Tentatives d’allumage visibles (étincelle) mais pas de flamme.
- Code erreur “absence de flamme” ou “défaut ionisation”.
- La chaudière se met en sécurité après plusieurs essais.[web:2][web:5]

**Cause racine probable :**
Absence de gaz au brûleur, pression gaz insuffisante, électrode d’allumage/ionisation mal positionnée ou encrassée, mélange air/gaz incorrect.[web:2][web:9]

**Étapes de résolution :**

1. Sécurité et environnement  
   - Couper l’alimentation électrique et fermer le robinet de gaz si odeur suspecte.  
   - S’assurer de la ventilation du local conformément à la réglementation gaz.[web:6][web:9]  

2. Contrôle alimentation gaz  
   - Vérifier que le robinet de gaz chaudière est ouvert.  
   - Vérifier la pression gaz amont (manomètre / prise de pression selon constructeur).  

3. Contrôle électrode d’allumage / ionisation  
   - Couper l’alimentation, déposer le brûleur.  
   - Nettoyer l’électrode (non abrasif) et vérifier l’écartement et la position selon la notice.[web:4][web:7]  
   - Contrôler l’état du câble haute tension et des connectiques.  

4. Contrôle du brûleur et de la rampe gaz  
   - Nettoyer le brûleur si encrassement (poussières, oxydation).  
   - Vérifier la propreté des orifices de la rampe gaz, souffler si bouchés (selon recommandations constructeur).  

5. Contrôle mélange air/gaz  
   - Sur chaudière à ventilateur : vérifier que le ventilateur démarre correctement.  
   - Contrôler les conduits d’air/fumées (obstruction possible).  

6. Remise en service et tests  
   - Remonter l’ensemble, remettre gaz et électricité.  
   - Lancer un cycle de démarrage, contrôler la stabilité de la flamme et les valeurs de combustion (CO₂, O₂).[web:3][web:10]  

**Prévention :**
- Nettoyage régulier du brûleur et contrôle de l’électrode à chaque entretien annuel.[web:3][web:5]  
- Vérification de la pression gaz et des conduits d’air/fumées conformément aux préconisations constructeur.[web:6][web:9]
```


### FACT-CHAUD-003 : Pression d’eau instable, mise en sécurité

```markdown
### FACT-CHAUD-003: Pression d’eau instable, mise en sécurité

| Champ        | Valeur                                  |
|--------------|-----------------------------------------|
| **ID**       | FACT-CHAUD-003                          |
| **Catégorie**| Circuit hydraulique chauffage           |
| **Système**  | Circuit primaire, vase d’expansion      |
| **Gravité**  | Moyenne                                 |
| **Marques**  | Multi-marques                           |

**Symptômes :**
- Pression à froid très basse (< 1 bar), alarme manque d’eau.  
- Montée rapide > 3 bar en chauffe, ouverture soupape de sécurité.  
- Remplissages fréquents par l’utilisateur.[web:2][web:5]

**Cause racine probable :**
Vase d’expansion dégonflé ou hors service, fuite sur le circuit, erreur de remplissage, manomètre ou pressostat défectueux.[web:2][web:10]

**Étapes de résolution :**

1. Contrôle visuel et recherche de fuites  
   - Inspecter radiateurs, raccords, purgeurs automatiques, soupape 3 bar.  
   - Vérifier absence d’écoulement permanent en évacuation de soupape.  

2. Contrôle vase d’expansion  
   - Isoler la chaudière du réseau et vider la pression côté eau.  
   - Mesurer la pression d’air du vase (valve type pneu).  
   - Regonfler à la pression constructeur (souvent 0,8–1,0 bar) ou remplacer si membrane HS.[web:2][web:5]  

3. Contrôle capteur de pression / pressostat  
   - Vérifier la cohérence entre manomètre mécanique et valeur affichée.  
   - Si incohérent : suspecter capteur défectueux, prévoir remplacement.  

4. Remplissage et purge  
   - Remplir lentement le circuit jusqu’à 1,2–1,5 bar à froid (selon notice).  
   - Purger l’air sur les radiateurs et purgeurs automatiques.  
   - Vérifier la stabilité de la pression en chauffe et à l’arrêt.  

**Prévention :**
- Contrôle systématique du vase d’expansion et de la pression de gonflage lors de l’entretien annuel.[web:3][web:5]  
- Sensibilisation de l’utilisateur à ne pas dépasser la pression recommandée au remplissage.[web:5]
```


### FACT-CHAUD-004 : Surchauffe fréquente, coupure sécurité

```markdown
### FACT-CHAUD-004: Surchauffe fréquente, coupure sécurité

| Champ        | Valeur                                  |
|--------------|-----------------------------------------|
| **ID**       | FACT-CHAUD-004                          |
| **Catégorie**| Régulation & sécurité température       |
| **Système**  | Capteurs température, circulation       |
| **Gravité**  | Élevée                                  |
| **Marques**  | Multi-marques                           |

**Symptômes :**
- Arrêts fréquents pour surchauffe, code type “surchauffe primaire”.  
- Bruits de bouillonnement dans la chaudière.  
- Radiateurs peu ou mal chauffés malgré surchauffe chaudière.[web:2][web:5]

**Cause racine probable :**
Mauvaise circulation d’eau (pompe bloquée, filtres bouchés, vannes fermées), entartrage important échangeur, ou sonde de température défaillante.[web:2][web:10]

**Étapes de résolution :**

1. Contrôle circulation  
   - Vérifier que les vannes départ/retour sont ouvertes.  
   - Contrôler le fonctionnement de la pompe (bruit, vibration, intensité).  

2. Nettoyage filtre / pot à boues  
   - Couper la chaudière et isoler.  
   - Nettoyer le filtre retour / pot de décantation si présent.  

3. Contrôle échangeur  
   - Sur chaudières instantanées : vérifier entartrage possible de l’échangeur.  
   - Prévoir détartrage chimique si fortement encrassé (suivant protocole constructeur).[web:3][web:10]  

4. Contrôle sondes  
   - Vérifier les valeurs lues par la régulation (menus diagnostics).  
   - Mesurer la résistance des sondes NTC et comparer aux courbes constructeur.[web:4][web:7]  

5. Réglages de régulation  
   - Vérifier la loi d’eau et la température maximale départ.  
   - Adapter si surdimensionnement / surchauffe systématique.  

**Prévention :**
- Mise en place de filtration et traitement d’eau adaptés (désembouage, inhibiteur). [web:3][web:5]  
- Vérification et nettoyage réguliers de la pompe et des filtres lors de l’entretien.
```


### FACT-CHAUD-005 : Rendement faible, consommation gaz élevée

```markdown
### FACT-CHAUD-005: Rendement faible, consommation gaz élevée

| Champ        | Valeur                                  |
|--------------|-----------------------------------------|
| **ID**       | FACT-CHAUD-005                          |
| **Catégorie**| Performance & rendement                 |
| **Système**  | Combustion, régulation, hydraulique     |
| **Gravité**  | Moyenne                                 |
| **Marques**  | Multi-marques                           |

**Symptômes :**
- Facture de gaz anormalement élevée.  
- Température ambiante peu stable.  
- Température de fumées élevée sur chaudière non condensation, ou mauvaise condensation sur chaudière à condensation.[web:3][web:8]

**Cause racine probable :**
Réglages de combustion non optimisés, loi d’eau mal paramétrée, absence de régulation pièce, circulation mal équilibrée, ou entretien insuffisant (échangeur encrassé).[web:3][web:5]

**Étapes de résolution :**

1. Vérification combustion  
   - Mesurer CO₂/O₂ et T fumées, comparer aux valeurs constructeur.[web:2][web:8]  
   - Ajuster réglages brûleur (air/gaz) si hors plage.  

2. Optimisation régulation  
   - Paramétrer la loi d’eau et les consignes selon l’isolation et l’émetteur (radiateurs vs plancher).  
   - Activer / configurer la sonde extérieure si disponible.[web:5][web:9]  

3. Hydraulique  
   - Équilibrer le réseau (débimètres, robinets d’équilibrage).  
   - Vérifier l’ouverture des robinets thermostatiques, éviter les fermetures complètes.  

4. Entretien  
   - Nettoyer échangeur, conduit de fumées et ventilateur.  
   - S’assurer de l’absence de boues importantes dans le réseau.  

**Prévention :**
- Entretien annuel avec mesure systématique de combustion, réglages consignés dans la fiche d’entretien.[web:2][web:3][web:10]  
- Revue périodique des réglages de régulation après travaux d’isolation ou modification du réseau.[web:5]
```


## Comment passer à ~100 FFT et par marque

Pour rester gérable, la suite est plus efficace en mode “génération assistée” que tout faire à la main ici :

1. **Axes de déclinaison**
    - Par système : brûleur, fumées, hydraulique, ECS, régulation, sécurité gaz, électronique.
    - Par type de chaudière : atmosphérique, ventouse, condensation, murale/sol, avec/ sans ballon.
    - Par marque : Saunier Duval, Vaillant, Frisquet, De Dietrich, Atlantic, Chaffoteaux…[^6][^7][^8]
2. **Pattern de prompt pour ton modèle local**
Tu peux nourrir un modèle local avec une FACT existante et lui demander de générer, par exemple :

> “En utilisant ce format de FACT (ID, tableau, Symptoms, Root Cause, Resolution Steps, Prevention), génère une nouvelle fiche pour une panne typique de chaudière gaz murale **Saunier Duval ThemaPlus C25**, catégorie ‘Brûleur \& combustion’, avec un code erreur d’allumage, en français.”
3. **Organisation GitLab**
    - Un repo `chaudieres-technical-facts` avec un fichier par domaine (`01_combustion.md`, `02_hydraulique.md`, etc.) sur le modèle de tes deux KB existantes.[^2][^1]
    - Un job CI qui valide la numérotation des FACT, la présence des sections, etc.

Si tu veux, prochain message : génération d’un **fichier complet** (par ex. “Problèmes brûleur \& combustion” avec 10–15 FACT) prêt à être commit, et tu pourras ensuite cloner ce pattern pour atteindre les ~100 FFT.

<div align="center">⁂</div>

[^1]: GitLab_Technical_Facts_Knowledge_Base.md

[^2]: USS_Enterprise_Technical_Troubleshooting_Manual.md

[^3]: https://energieplus-lesite.be/evaluer/chauffage4/evaluer-efficacite-energetique-production-chaleur/interpreter-la-fiche-d-entretien-d-une-chaudiere-ou-une-mesure-des-parametres-de-combustion/

[^4]: https://www.alec-saint-brieuc.org/wp-content/uploads/2018/09/Fiches-Bonnes-pratiques-et-reglages.pdf

[^5]: https://cegibat.grdf.fr/sites/default/files/assets/Dossier_technique_-_Entretien_annuel_des_chaudieres_gaz_individuelles.pdf

[^6]: https://appli.savchaudiere.com/pdf/2/notices/2N11095.pdf

[^7]: https://www.saunierduval.fr/particulier/nous-vous-accompagnons/je-m-informe/sur-la-chaudiere/chaudiere-normes-d-installation-et-reglementation-gaz/

[^8]: https://cegibat.grdf.fr/fiches-pratiques

