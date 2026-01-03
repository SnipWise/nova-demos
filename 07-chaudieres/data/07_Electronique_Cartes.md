# Problèmes Électronique & Cartes de Commande

**Catégorie :** Électronique & Cartes de Commande
**Nombre de Facts :** 10
**Retour à l'index :** [Knowledge_Base_Chaudieres.md](Knowledge_Base_Chaudieres.md)

---

## FACT-CHAUD-091: Carte électronique défaillante - Symptômes généraux

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-091 |
| **Catégorie** | Électronique & Cartes de Commande |
| **Système** | Carte électronique principale |
| **Gravité** | **Élevée** |
| **Marques** | Multi-marques |

**Symptômes :**
- Chaudière complètement inerte (pas de réaction)
- Affichage erratique ou codes erreurs multiples
- Dysfonctionnements aléatoires et intermittents
- Réarmages fréquents nécessaires
- Bruits de relais qui claquent de manière répétée
- Odeur de composants brûlés

**Cause racine probable :**
Surtension électrique, foudre indirecte, vieillissement composants, humidité/condensation, surchauffe carte, défaut alimentation, composant CMS dessoudé.

**Étapes de résolution :**

1. **Sécurité électrique**
   - Couper l'alimentation 230V au disjoncteur
   - Attendre 5 minutes (décharge condensateurs)
   - Vérifier absence de tension avant intervention
   - Prendre photo de la carte avant démontage

2. **Inspection visuelle carte**
   - Retirer la carte électronique selon procédure constructeur
   - Observer à la lumière : composants brûlés, noircis, gonflés
   - Vérifier condensateurs : pas de bombement, fuite électrolyte
   - Contrôler soudures : pas de fissures, dessoudages
   - Chercher traces d'humidité, corrosion, oxydation
   - Vérifier absence de poussière conductrice

3. **Contrôle alimentation électrique**
   - Mesurer tension secteur : 230V ± 10% (207-253V)
   - Vérifier qualité terre (résistance < 100 Ω)
   - Contrôler protection parafoudre si présente
   - Tester stabilité tension (pas de micro-coupures)
   - Vérifier transformateur carte : tensions secondaires correctes

4. **Test fusibles et protections**
   - Vérifier tous les fusibles carte (visuel + continuité)
   - Contrôler état fusibles : pas noircis
   - Si fusible grillé : chercher cause avant remplacement
   - Vérifier varistances protection (composants ronds jaunes/bleus)

5. **Diagnostic connectique**
   - Débrancher et rebrancher tous les connecteurs
   - Vérifier état des cosses (oxydation, desserrage)
   - Contrôler câblage : pas de fils dénudés, coupés
   - Tester continuité des nappes et fils
   - Nettoyer contacts avec bombe contact

6. **Test par substitution**
   - Si doute sur carte : test avec carte identique (prêt SAV)
   - Attention : certaines cartes nécessitent paramétrage
   - Noter références exactes carte (étiquette, sérigraphie)
   - Vérifier compatibilité version firmware

7. **Remplacement carte**
   - Commander carte selon références constructeur
   - Certaines marques : carte codée (paramétrage usine nécessaire)
   - Transférer paramètres ancienne carte si possible
   - Programmer type chaudière, puissance, options
   - Effectuer auto-diagnostic après remplacement

8. **Prévention surtensions**
   - Installer parafoudre modulaire tableau électrique
   - Ajouter onduleur si zone à risque foudre
   - Vérifier qualité mise à terre installation
   - Protéger contre micro-coupures réseau

**Prévention :**
- Parafoudre et protection surtensions obligatoires
- Vérification annuelle serrage connecteurs
- Nettoyage carte (soufflette air sec) si environnement poussiéreux
- Protection contre humidité local chaufferie
- Remplacement préventif condensateurs après 10-12 ans

**Coût indicatif :**
- Carte électronique : 150-600€ selon marque/modèle
- Main d'œuvre remplacement : 1-2h
- Programmation/paramétrage : 30min-1h
- Parafoudre préventif : 50-150€

**Spécificités par marque :**
- **Vaillant** : cartes souvent codées, paramétrage via logiciel
- **Saunier Duval** : cartes interchangeables sans codage
- **Frisquet** : système Vision nécessite appairage
- **De Dietrich** : code installation à saisir
- **Elm Leblanc** : auto-reconnaissance sur certains modèles

---

## FACT-CHAUD-092: Fusible carte grillé

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-092 |
| **Catégorie** | Électronique & Cartes de Commande |
| **Système** | Protection électrique carte |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques |

**Symptômes :**
- Chaudière totalement inerte, aucun affichage
- Pas de LED, pas de réaction aux boutons
- Ou : fonction spécifique HS (ventilateur, pompe, etc.)
- Fusible noirci ou filament rompu visible

**Cause racine probable :**
Court-circuit sur composant protégé, surtension secteur, composant défaillant en aval, humidité, vieillissement fusible, erreur câblage.

**Étapes de résolution :**

1. **Identification fusible défectueux**
   - Couper alimentation 230V
   - Localiser les fusibles sur la carte (cylindres verre ou CMS)
   - Types courants : 5x20mm verre, fusibles CMS, porte-fusible
   - Tester continuité au multimètre (0 Ω si bon)
   - Déposer et inspecter visuellement

2. **Identification type fusible**
   - Noter valeur marquée : ex. T2A (temporisé 2 ampères)
   - F = rapide (Fast), T = temporisé (Time-lag)
   - Tensions : 250V ou 125V marquées
   - Exemples courants chaudières :
     - Alimentation générale : T3.15A ou T5A (250V)
     - Ventilateur : T2A (250V)
     - Pompe : T2A ou T3.15A
     - Circuits basse tension : T500mA ou T1A (125V)

3. **Recherche cause du défaut**
   - **IMPORTANT** : Ne jamais remplacer sans chercher la cause
   - Fusible = protection, pas pièce d'usure
   - Débrancher la charge protégée (pompe, ventilateur, vanne)
   - Tester isolement charge (résistance > 1 MΩ vers masse)
   - Mesurer résistance/bobine (doit correspondre valeurs nominales)
   - Vérifier absence court-circuit évident

4. **Contrôle composants aval**
   - Si fusible pompe : tester pompe isolée en direct
   - Si fusible ventilateur : tester ventilateur
   - Si fusible vanne gaz : tester bobines vanne
   - Si fusible alimentation : chercher court-circuit carte
   - Utiliser multimètre mode résistance

5. **Vérification absence humidité**
   - Inspecter carte : traces d'humidité, condensation
   - Sécher carte complètement si humide (air sec tiède)
   - Vérifier joint carter étanche
   - Contrôler absence infiltration eau

6. **Remplacement fusible**
   - Utiliser fusible **exactement identique** (calibre, type, tension)
   - JAMAIS de valeur supérieure (risque incendie)
   - JAMAIS de fil ou "bricolage" à la place
   - Enficher fermement dans porte-fusible
   - Sur fusible CMS : soudure avec fer température contrôlée

7. **Test remise en service**
   - Reconnecter les charges une par une
   - Remettre tension et tester
   - Observer comportement : fusible doit tenir
   - Si re-grille immédiatement : défaut persistant
   - Mesurer consommation si possible (pince ampèremétrique)

8. **Cas fusible re-grille immédiatement**
   - Court-circuit franc : chercher composant HS sur carte
   - Vérifier condensateurs gonflés
   - Tester diodes, transistors, relais
   - Débrancher charges suspectes une par une
   - Si pas de solution : remplacement carte nécessaire

**Prévention :**
- Contrôle annuel serrage connecteurs (mauvais contact = échauffement)
- Protection parafoudre secteur
- Éviter humidité locale chaufferie
- Vérification fonctionnement charges (pompe, ventilateur)

**Stock préventif recommandé :**
- Fusibles T3.15A 250V (x2)
- Fusibles T2A 250V (x2)
- Conserver dans boîte étanche

**Références fusibles courants :**
- Standard 5x20mm temporisé trouvable commerce électronique
- Marques fiables : Littelfuse, Schurter, Eaton Bussmann
- Éviter fusibles "premier prix" sous-calibrés

---

## FACT-CHAUD-093: Défaut alimentation 230V

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-093 |
| **Catégorie** | Électronique & Cartes de Commande |
| **Système** | Alimentation électrique |
| **Gravité** | **Élevée** |
| **Marques** | Multi-marques |

**Symptômes :**
- Chaudière complètement morte, pas d'affichage
- Disjoncteur qui saute au démarrage
- Affichage faible ou clignotant
- Réinitialisation horloge fréquente
- Fonctionnement intermittent

**Cause racine probable :**
Disjoncteur défectueux, câblage défaillant, mauvaise connexion bornier, tension secteur anormale, absence de neutre, inversion phase/neutre, défaut terre.

**Étapes de résolution :**

1. **Vérification tableau électrique**
   - Vérifier disjoncteur chaudière en position ON
   - Tester disjoncteur : déclencher et réenclencher
   - Vérifier calibre disjoncteur (généralement 10A ou 16A courbe C)
   - Contrôler état disjoncteur : pas de traces brûlure
   - Si vieux disjoncteur : tester par substitution

2. **Mesure tension secteur**
   - Mesurer tension au tableau : 230V ± 10% (207-253V)
   - Mesurer tension au bornier chaudière (peut différer)
   - Vérifier phase-neutre : ~230V
   - Vérifier phase-terre : ~230V
   - Vérifier neutre-terre : < 5V (sinon problème neutre)

3. **Contrôle câblage alimentation**
   - Vérifier section câble : minimum 1.5mm² (conseillé 2.5mm²)
   - Contrôler état câble (gaine non dégradée)
   - Vérifier longueur : si > 20m, chute tension possible
   - Mesurer résistance câble (doit être ~0 Ω)
   - Contrôler absence de câble sectionné, pincé

4. **Contrôle bornier chaudière**
   - Couper alimentation tableau
   - Ouvrir bornier électrique chaudière
   - Vérifier serrage vis : doivent être fermes
   - Contrôler état connexions : pas d'oxydation, noircissement
   - Vérifier identification : L (phase), N (neutre), T (terre)
   - Resserrer toutes les connexions

5. **Vérification phase/neutre/terre**
   - Identifier phase (fil marron ou noir, ou rouge)
   - Identifier neutre (fil bleu)
   - Identifier terre (fil vert/jaune)
   - ATTENTION inversion phase/neutre : certains appareils ne fonctionnent pas
   - Vérifier avec testeur de polarité ou tournevis testeur

6. **Test continuité terre**
   - Mesurer résistance terre : < 100 Ω (norme < 100 Ω)
   - Si > 100 Ω : problème installation terre
   - Vérifier continuité terre tableau → chaudière
   - Contrôler raccordement terre sur chaudière (cosse serrée)

7. **Diagnostic disjoncteur qui saute**
   - **Immédiatement au démarrage** : court-circuit franc
     - Débrancher charges une par une (pompe, ventilateur, vanne)
     - Identifier composant en court-circuit
   - **Après quelques secondes** : surconsommation
     - Mesurer courant total chaudière (< 10A normalement)
     - Vérifier pas de charge anormale
   - **Aléatoirement** : défaut terre (différentiel) ou disjoncteur défectueux
     - Tester avec autre disjoncteur
     - Vérifier isolement terre tous composants

8. **Contrôle transformateur alimentation**
   - Sur carte : localiser transformateur (composant gros rectangulaire)
   - Mesurer tensions secondaires (selon schéma : 12V, 24V)
   - Si tensions absentes : transformateur HS ou fusible primaire grillé
   - Remplacer transformateur si défectueux (selon modèle : soudé ou enfichable)

9. **Cas micro-coupures réseau**
   - Symptôme : réinitialisation horloge, pertes paramètres
   - Vérifier qualité alimentation secteur (voltmètre enregistreur)
   - Installer onduleur si micro-coupures fréquentes
   - Ou système batterie backup pour maintien mémoire

**Prévention :**
- Vérification annuelle serrage bornier
- Contrôle calibre et état disjoncteur
- Test terre annuel
- Resserrage connexions tableau électrique
- Protection parafoudre + différentiel 30mA obligatoire

**Conformité électrique :**
- Disjoncteur dédié chaudière obligatoire
- Différentiel 30mA type A obligatoire
- Section câble : 1.5mm² mini (2.5mm² conseillé)
- Couleurs normalisées respectées
- Mise à terre < 100 Ω obligatoire

**Outils nécessaires :**
- Multimètre (tension AC, résistance)
- Tournevis testeur phase
- Pince ampèremétrique (diagnostic)
- Testeur différentiel (si disponible)

---

## FACT-CHAUD-094: Relais carte défectueux

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-094 |
| **Catégorie** | Électronique & Cartes de Commande |
| **Système** | Relais commande charges |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques |

**Symptômes :**
- Pompe ne démarre pas ou tourne en permanence
- Ventilateur ne s'active pas ou reste actif
- Vanne gaz ne s'ouvre pas
- Bruit de claquement répété sans effet
- Brûleur ne démarre pas malgré demande

**Cause racine probable :**
Relais collé (contacts soudés), bobine relais grillée, contacts usés/oxydés, surtension ayant endommagé relais, vieillissement mécanique.

**Étapes de résolution :**

1. **Identification relais défectueux**
   - Localiser les relais sur carte (composants rectangulaires transparents/bleus)
   - Identification marquage : K1, K2, K3, RLY1, etc.
   - Fonctions courantes :
     - Relais pompe (marche/arrêt circulateur)
     - Relais ventilateur
     - Relais vanne gaz (1 ou 2 relais)
     - Relais brûleur/allumage
   - Observer LED associée si présente

2. **Test auditif relais**
   - Remettre chaudière sous tension
   - Lancer demande chauffage
   - Écouter "clic" relais à chaque activation
   - Absence de "clic" : bobine relais HS ou commande absente
   - "Clic" présent mais pas d'effet : contacts défectueux

3. **Test visuel relais**
   - Couper alimentation
   - Déposer carte pour accès relais
   - Vérifier état visuel relais :
     - Traces de brûlure, noircissement
     - Déformation plastique
     - Traces d'arc électrique
   - Si doute : dessouder et tester hors carte

4. **Test électrique relais**
   - **Test bobine** (relais déposé ou in-situ) :
     - Mesurer résistance bobine : généralement 50-500 Ω
     - Si infini (∞) : bobine coupée, relais HS
     - Si 0 Ω : bobine court-circuitée, relais HS
   - **Test contacts** (relais déposé) :
     - Identifier bornes : commun (C), NO (normalement ouvert), NC (normalement fermé)
     - Mesurer continuité au repos : C-NC fermé, C-NO ouvert
     - Alimenter bobine (12V ou 24V selon specs) : inversion
     - Si pas d'inversion : contacts grippés/soudés

5. **Diagnostic contacts collés**
   - Symptôme : charge fonctionne en permanence
   - Cause : arc électrique a soudé les contacts
   - Test : couper alimentation, mesurer continuité contacts NO (doit être ouvert)
   - Si fermé au repos : relais collé, remplacement obligatoire

6. **Diagnostic bobine grillée**
   - Symptôme : pas de "clic", charge ne démarre jamais
   - Mesurer tension commande bobine (carte sous tension, demande active)
   - Si tension présente (12V ou 24V) mais pas de "clic" : bobine HS
   - Si pas de tension : problème carte électronique (transistor commande)

7. **Remplacement relais**
   - Noter références exactes relais :
     - Tension bobine : 12VDC, 24VDC, etc.
     - Courant contacts : 10A, 16A, etc.
     - Configuration : SPDT (1RT), DPDT (2RT)
     - Exemple : OMRON G5LE, FINDER 40.52, SONGLE SRD
   - Dessouder ancien relais (pompe à dessouder ou tresse)
   - Souder nouveau relais (orientation correcte !)
   - Nettoyer flux soudure

8. **Alternative réparation temporaire**
   - **Contacts oxydés/sales** (si relais accessible) :
     - Démonter capot relais (clipsé)
     - Nettoyer contacts avec papier abrasif très fin ou contact cleaner
     - Remonter et tester
     - Solution temporaire : prévoir remplacement
   - **Relais collé** :
     - Tapoter légèrement sur relais (peut décrocher contacts)
     - Solution très temporaire, remplacement urgent

9. **Test après remplacement**
   - Remonter carte
   - Reconnecter charges
   - Mettre sous tension
   - Tester activation relais : écouter "clic"
   - Vérifier fonctionnement charge (pompe, ventilateur)
   - Mesurer tension sortie relais

**Prévention :**
- Remplacement préventif relais après 8-10 ans (selon sollicitation)
- Éviter surtensions (parafoudre)
- Vérifier charges (pompe, ventilateur) pour éviter surcourant
- Protéger relais avec varistance ou RC snubber si charges inductives

**Relais courants chaudières :**
- Relais 12VDC 10A SPDT : pompe, ventilateur
- Relais 24VDC 16A SPDT : charges puissantes
- Montage sur circuit imprimé (PCB)
- Disponibles en commerce électronique

**Compétences requises :**
- Soudure composants traversants
- Lecture schéma électrique
- Mesures électriques multimètre
- Si pas compétent : remplacement carte complète

---

## FACT-CHAUD-095: Afficheur LCD HS ou illisible

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-095 |
| **Catégorie** | Électronique & Cartes de Commande |
| **Système** | Interface utilisateur |
| **Gravité** | Faible |
| **Marques** | Multi-marques |

**Symptômes :**
- Écran complètement noir ou blanc
- Affichage partiel (segments manquants)
- Affichage fantôme (tous segments allumés)
- Contraste trop faible ou trop fort
- Affichage inversé (négatif)
- LCD cassé, fissuré

**Cause racine probable :**
Nappe LCD dessoudée/déconnectée, LCD cassé (choc), vieillissement cristaux liquides, température extrême, défaut contraste, défaut rétroéclairage, carte électronique défaillante.

**Étapes de résolution :**

1. **Diagnostic type affichage**
   - **Écran complètement noir** :
     - Vérifier rétroéclairage (visible dans le noir)
     - Si rétroéclairage OK mais pas de caractères : problème contraste ou LCD
   - **Écran blanc** :
     - Problème contraste (réglage ou composant)
     - Nappe déconnectée
   - **Segments manquants** :
     - Nappe partiellement déconnectée
     - Pistes coupées sur nappe
     - LCD défectueux
   - **Tous segments allumés** :
     - Court-circuit commande LCD
     - Carte électronique défaillante

2. **Vérification nappe LCD**
   - Couper alimentation
   - Localiser connexion LCD ↔ carte électronique
   - Types connexion :
     - Nappe élastomère (zébra) : bande caoutchouc conductrice
     - Nappe flexible soudée
     - Connecteur enfichable
   - Vérifier état nappe : pas de coupure, pliure excessive
   - Nettoyer contacts (alcool isopropylique + coton-tige)
   - Remettre en place en appuyant fermement (zébra)

3. **Réglage contraste LCD**
   - Certains LCD ont potentiomètre contraste sur carte
   - Localiser (petit composant bleu à vis cruciforme)
   - Ajuster délicatement (1/4 tour max à la fois)
   - Tester affichage après chaque ajustement
   - Position optimale : affichage net, noir/blanc contrasté

4. **Test rétroéclairage**
   - Regarder LCD dans le noir
   - Si éclairage visible mais pas de caractères : problème LCD ou nappe
   - Si pas d'éclairage :
     - Vérifier LED rétroéclairage (arrière LCD)
     - Tester alimentation LED (généralement 12V ou 3.3V)
     - Remplacer LED si grillée (soudure CMS)

5. **Contrôle température**
   - LCD sensibles températures extrêmes :
     - < 0°C : cristaux liquides figés, affichage lent/absent
     - > 60°C : dégradation cristaux, affichage fantôme
   - Vérifier température locale chaufferie
   - Isoler/ventiler si nécessaire
   - LCD peut récupérer après retour température normale

6. **Test LCD avec autre carte**
   - Si possible : tester LCD sur carte identique fonctionnelle
   - Permet identifier : problème LCD ou problème carte
   - Attention compatibilité : certains LCD spécifiques

7. **Remplacement LCD**
   - **LCD zébra** (bande élastomère) :
     - Commander LCD + bande élastomère
     - Déposer ancien LCD (clips ou vis)
     - Positionner bande zébra (alignement précis)
     - Clip nouveau LCD
   - **LCD nappe soudée** :
     - Dessouder nappe (fer fine pointe, température contrôlée)
     - Nettoyer pads
     - Souder nouveau LCD (flux, soudure fine)
     - Attention : facile d'endommager pistes
   - **LCD enfichable** :
     - Retirer connecteur
     - Enficher nouveau LCD

8. **Solution temporaire**
   - Si affichage défectueux mais chaudière fonctionne :
     - Utiliser codes LED (selon marque) pour diagnostic
     - Installer thermostat d'ambiance programmable (compense)
     - Planifier remplacement LCD ou carte

9. **Alternative : remplacement carte complète**
   - Si LCD intégré carte (soudé CMS complexe) :
     - Remplacement LCD seul difficile
     - Préférer remplacement carte complète
   - Comparer coûts : LCD seul vs carte complète

**Prévention :**
- Protection contre températures extrêmes
- Éviter chocs sur afficheur
- Nettoyage doux (chiffon microfibre légèrement humide)
- Ne jamais appuyer fortement sur LCD

**Compatibilité :**
- LCD souvent spécifiques à chaque modèle chaudière
- Noter références exactes avant commande
- Vérifier nombre de caractères, segments
- Photos utiles pour identification

**Fonctionnement sans afficheur :**
- Chaudière peut fonctionner sans affichage
- Utiliser codes LED ou thermostat externe
- Diagnostic limité mais fonctionnel

---

## FACT-CHAUD-096: Boutons commande non fonctionnels

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-096 |
| **Catégorie** | Électronique & Cartes de Commande |
| **Système** | Interface boutons poussoirs |
| **Gravité** | Faible à Moyenne |
| **Marques** | Multi-marques |

**Symptômes :**
- Boutons ne répondent pas à la pression
- Bouton enfoncé en permanence (collé)
- Réponse intermittente (il faut appuyer plusieurs fois)
- Mauvais bouton activé (crosstalk)
- Impression tactile altérée (pas de "clic")

**Cause racine probable :**
Bouton mécanique usé, dôme conducteur cassé, oxydation contacts, encrassement, humidité, nappe boutons déconnectée, piste carte coupée.

**Étapes de résolution :**

1. **Identification type boutons**
   - **Boutons poussoirs mécaniques** (switches) :
     - Composant monté sur carte
     - Tactile "clic" mécanique
   - **Boutons à dôme** (dome switch) :
     - Dôme métallique sous membrane
     - Contact sur pistes PCB
   - **Boutons capacitifs** (tactile) :
     - Pas de mouvement mécanique
     - Détection capacitive doigt

2. **Test boutons individuels**
   - Accéder mode diagnostic/menu si possible
   - Tester chaque bouton séquentiellement
   - Noter quels boutons répondent/ne répondent pas
   - Vérifier absence de bouton bloqué enfoncé

3. **Nettoyage boutons à dôme**
   - Démonter façade chaudière
   - Retirer membrane silicone/plastique
   - Nettoyer dômes métalliques :
     - Alcool isopropylique + coton-tige
     - Enlever oxydation, saleté
   - Nettoyer pistes PCB sous dômes :
     - Alcool isopropylique
     - Gomme douce si oxydation importante
   - Nettoyer membrane par l'intérieur
   - Sécher complètement avant remontage

4. **Contrôle dômes métalliques**
   - Vérifier état dômes :
     - Bombement correct (pas écrasés)
     - Pas de déchirure, déformation
     - Élasticité conservée (ressort)
   - Remplacer dômes défectueux :
     - Disponibles commerce électronique
     - Tailles standards : 8mm, 12mm diamètre
     - Hauteur et force variable

5. **Test boutons poussoirs mécaniques**
   - Couper alimentation
   - Mesurer continuité bouton au repos : ouvert (∞)
   - Appuyer sur bouton : fermé (0 Ω)
   - Si pas de commutation : bouton HS
   - Dessouder et remplacer bouton

6. **Contrôle nappe boutons**
   - Si boutons déportés (nappe vers carte) :
     - Vérifier connexion nappe ↔ carte
     - Nettoyer connecteur
     - Tester continuité pistes nappe
     - Remplacer nappe si coupée

7. **Vérification pistes PCB**
   - Inspecter pistes sous boutons (loupe)
   - Vérifier absence coupure piste
   - Tester continuité bouton → composant carte
   - Réparer piste si coupée (fil émaillé fin + soudure)

8. **Remplacement boutons poussoirs**
   - Noter référence bouton (hauteur, force, type)
   - Types courants :
     - Tactile switch 6x6mm hauteur 4.3mm à 13mm
     - Force : 160gf, 260gf, 520gf
   - Dessouder ancien bouton (pompe à dessouder)
   - Souder nouveau bouton (alignement correct)

9. **Solution temporaire**
   - Si bouton unique défectueux :
     - Utiliser thermostat d'ambiance pour contrôle
     - Ou télécommande filaire (si compatible)
   - Planifier réparation

10. **Boutons capacitifs défectueux**
    - Plus complexe à diagnostiquer
    - Vérifier circuits capacitifs (condensateurs, IC)
    - Nettoyer surface tactile (traces de doigt)
    - Recalibration possible sur certains modèles
    - Sinon : remplacement carte ou module tactile

**Prévention :**
- Éviter humidité (mains mouillées)
- Nettoyage régulier façade
- Ne pas appuyer avec objets durs (risque casse)
- Protection contre vapeur cuisine (si chaudière en cuisine)

**Pièces de rechange :**
- Dômes conducteurs : lot 100 pièces ~5-10€
- Boutons tactiles 6x6mm : lot 50 pièces ~3-5€
- Disponibles sur sites électronique (AliExpress, Amazon, Mouser)

**Fonctionnement sans boutons :**
- Installer thermostat d'ambiance filaire
- Utilise uniquement consigne thermostat
- Chaudière suit demande externe
- Perd accès réglages avancés

---

## FACT-CHAUD-097: Reset permanent / Réinitialisation intempestive

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-097 |
| **Catégorie** | Électronique & Cartes de Commande |
| **Système** | Alimentation / Mémoire carte |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques |

**Symptômes :**
- Chaudière redémarre aléatoirement
- Horloge se réinitialise fréquemment (00:00)
- Paramètres perdus après chaque coupure
- Écran affiche "Initialisation" de façon répétée
- Erreurs mémoire affichées
- Code erreur watchdog / reset

**Cause racine probable :**
Micro-coupures secteur, tension instable, condensateur mémoire défaillant, pile CMOS/RTC déchargée, défaut alimentation carte, watchdog intempestif, défaut firmware.

**Étapes de résolution :**

1. **Analyse fréquence réinitialisations**
   - Noter quand elles surviennent :
     - À chaque coupure secteur → problème mémoire/pile
     - Aléatoire en fonctionnement → problème alimentation ou watchdog
     - À chaque démarrage pompe/ventilateur → chute tension
   - Observer codes erreur associés

2. **Contrôle alimentation secteur**
   - Mesurer tension secteur : 230V stable
   - Installer enregistreur tension (si disponible)
   - Vérifier absence micro-coupures
   - Contrôler qualité neutre (neutre-terre < 5V)
   - Tester avec autre circuit électrique (isolation)

3. **Vérification pile RTC (Real-Time Clock)**
   - Localiser pile sur carte :
     - Type CR2032 (lithium 3V) souvent
     - Ou pile rechargeable ML2032
     - Parfois super-condensateur
   - Mesurer tension pile : doit être > 2.7V
   - Si < 2.5V : pile déchargée, remplacement nécessaire
   - Pile durée vie : 3-5 ans

4. **Remplacement pile RTC**
   - Couper alimentation chaudière
   - Retirer pile (clip ou porte-pile)
   - Remplacer par pile identique (CR2032 3V)
   - **Attention** : certaines piles sont soudées (ML2032)
     - Dessouder ancienne
     - Souder nouvelle (rapidement, chaleur limitée)
   - Remettre alimentation
   - Régler horloge et paramètres

5. **Contrôle condensateurs mémoire**
   - Localiser condensateurs alimentation (gros cylindres)
   - Inspecter visuellement :
     - Pas de bombement dessus
     - Pas de fuite électrolyte (liquide)
     - Pas de traces brunâtres
   - Mesurer capacité si testeur disponible
   - Remplacer si défectueux (ESR élevé ou capacité faible)

6. **Diagnostic watchdog**
   - Watchdog = circuit surveillance qui reset si blocage détecté
   - Reset watchdog intempestif → problème firmware ou carte
   - Codes erreur possibles : "Watchdog", "WD Reset", "Internal Error"
   - Solutions :
     - Mise à jour firmware
     - Réinitialisation complète paramètres usine
     - Remplacement carte si persiste

7. **Contrôle chutes tension au démarrage charges**
   - Symptôme : reset quand pompe ou ventilateur démarre
   - Cause : chute tension excessive (démarrage moteur)
   - Mesurer tension 230V au démarrage pompe (oscilloscope ou voltmètre rapide)
   - Solutions :
     - Vérifier câblage alimentation (section suffisante)
     - Installer condensateur stabilisation
     - Vérifier qualité transformateur alimentation carte

8. **Réinitialisation complète paramètres**
   - Accéder menu service/installateur
   - Réinitialisation factory reset
   - Reprogrammer tous paramètres :
     - Type chaudière, puissance
     - Type gaz (G20/G25)
     - Paramètres spécifiques installation
   - Tester stabilité après réinit

9. **Mise à jour firmware**
   - Vérifier version firmware actuelle (menu diagnostic)
   - Consulter site constructeur : firmware plus récent disponible ?
   - Mise à jour via :
     - Interface USB (selon marque)
     - Remplacement EPROM (ancien modèles)
     - Outil SAV constructeur
   - Attention : mise à jour ratée peut bricker la carte

10. **Protection contre micro-coupures**
    - Installer onduleur (UPS) :
      - Puissance 300-500VA suffisante (carte électronique seule)
      - Maintient alimentation lors coupures brèves
    - Ou stabilisateur/régulateur tension
    - Protège aussi contre surtensions

**Prévention :**
- Remplacement pile RTC préventif tous les 5 ans
- Protection secteur (parafoudre + onduleur)
- Vérification qualité alimentation électrique
- Mise à jour firmware régulière

**Coût solutions :**
- Pile CR2032 : 1-3€
- Onduleur 500VA : 50-100€
- Mise à jour firmware : gratuite (si matériel compatible)
- Remplacement carte : 150-600€

**Paramètres à sauvegarder avant réinit :**
- Photographier tous écrans paramètres
- Noter sur papier : puissance, type gaz, courbe chauffe, etc.
- Permet reprogrammation rapide après reset

---

## FACT-CHAUD-098: Connectique oxydée ou dessoudée

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-098 |
| **Catégorie** | Électronique & Cartes de Commande |
| **Système** | Connecteurs carte électronique |
| **Gravité** | Moyenne |
| **Marques** | Multi-marques |

**Symptômes :**
- Fonctionnement intermittent aléatoire
- Perte de fonction spécifique (pompe, ventilateur, sonde)
- Faux contacts au démarrage
- Besoin de "tapoter" la chaudière pour redémarrer
- Sonde température avec valeurs erratiques
- Erreurs aléatoires non reproductibles

**Cause racine probable :**
Oxydation contacts (humidité), vibrations desserrant connexions, échauffement excessif (mauvais contact), corrosion (condensation), soudures froides, vieillissement.

**Étapes de résolution :**

1. **Identification connecteurs problématiques**
   - Observer quelles fonctions sont intermittentes
   - Localiser connecteurs associés sur carte :
     - Sondes température (CTN/NTC)
     - Pompe, ventilateur, vanne gaz
     - Alimentation 230V
     - Thermostat d'ambiance
   - Inspecter visuellement

2. **Diagnostic visuel connecteurs**
   - Retirer carte électronique
   - Inspecter chaque connecteur :
     - **Oxydation** : dépôt vert/blanc, contacts ternis
     - **Surchauffe** : plastique noirci, déformé
     - **Dessoudure** : soudure fissurée, composant mobile
     - **Corrosion** : dépôts brunâtres, poudre verte (cuivre oxydé)

3. **Nettoyage connecteurs oxydés**
   - **Connecteurs enfichables** (cosses Faston, connecteurs à clip) :
     - Débrancher connecteur
     - Nettoyer contacts mâles et femelles :
       - Bombe contact électrique (spray)
       - Ou alcool isopropylique + brosse douce
       - Ou gomme douce type Caig DeoxIT
     - Sécher complètement
     - Rebrancher fermement
   - **Bornier à vis** :
     - Desserrer vis
     - Nettoyer cosse fil (papier abrasif fin si oxydée)
     - Nettoyer borne (brosse laiton)
     - Revisser fermement

4. **Traitement anti-oxydation**
   - Après nettoyage, appliquer :
     - Graisse conductrice (vaseline technique)
     - Ou spray contact (CRC Contact, WD40 Contact Cleaner)
   - Protège contre oxydation future
   - Améliore conductivité

5. **Réparation soudures froides**
   - Identifier soudures fissurées (loupe) :
     - Aspect terne, granuleux
     - Fissure visible autour pin
     - Composant qui bouge légèrement
   - Refaire soudure :
     - Chauffer soudure existante + ajouter flux
     - Ajouter un peu étain neuf
     - Soudure doit être brillante, conique
   - Zones critiques :
     - Connecteurs puissance (230V, pompe, ventilateur)
     - Transformateur (composant lourd)
     - Relais (composant soumis vibrations)

6. **Remplacement connecteur défectueux**
   - Si connecteur plastique cassé/fondu :
     - Dessouder ancien connecteur
     - Identifier référence (pas, nombre broches)
     - Souder nouveau connecteur identique
   - Types courants :
     - Borniers à vis KF2EDGK (pas 5.08mm)
     - Connecteurs XH, JST (pas 2.54mm)
     - Cosses Faston 6.3mm

7. **Renforcement mécanique**
   - Si dessoudure due vibrations :
     - Ajouter colle thermofusible (hot glue) pour fixation mécanique
     - Ou mousse isolante entre composant et capot
   - Évite contraintes mécaniques répétées

8. **Contrôle continuité fils**
   - Mesurer continuité fil de bout en bout
   - Si coupure interne fil :
     - Couper fil défectueux
     - Dénuder
     - Raccorder avec domino ou soudure + gaine thermo
   - Vérifier absence faux contact dans gaine

9. **Protection contre humidité**
   - Si oxydation due humidité :
     - Vérifier étanchéité capot chaudière
     - Améliorer ventilation local
     - Appliquer vernis tropicalisation sur carte (précaution)
   - Installer absorbeur humidité dans capot si nécessaire

10. **Vérification serrage**
    - Resserrer TOUS les connecteurs à vis
    - Vérifier que cosses enfichables sont bien clipsées
    - Noter couples de serrage si spécifié (borniers puissance)

**Prévention :**
- Vérification annuelle serrage connecteurs
- Nettoyage préventif contacts tous les 2-3 ans
- Protection contre humidité locale chaufferie
- Application spray contact après chaque intervention

**Outils et produits :**
- Bombe contact électrique (CRC, WD40)
- Alcool isopropylique 99%
- Graisse conductrice (Electrolube, Caig DeoxIT)
- Fer à souder température contrôlée
- Flux décapant (pas acide)
- Multimètre continuité

**Zones sensibles :**
- Sondes température (courant faible, sensibles oxydation)
- Connecteurs puissance (échauffement si mauvais contact)
- Connecteurs thermostat (basse tension, oxydation)

**Symptômes typiques par connecteur :**
- **Sonde température** : valeurs erratiques, erreur sonde
- **Pompe** : démarrage aléatoire, arrêts intempestifs
- **Ventilateur** : F33 intermittent (Saunier Duval)
- **Thermostat** : chaudière ne suit pas demande

---

## FACT-CHAUD-099: Condensation/humidité sur carte

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-099 |
| **Catégorie** | Électronique & Cartes de Commande |
| **Système** | Étanchéité / Protection carte |
| **Gravité** | **Élevée** |
| **Marques** | Multi-marques |

**Symptômes :**
- Traces d'humidité visibles sur carte
- Gouttelettes d'eau sous capot
- Corrosion verte (oxydation cuivre) sur composants
- Dysfonctionnements erratiques
- Courts-circuits aléatoires
- Fusibles qui grillent de manière répétée
- Odeur de moisi dans chaudière

**Cause racine probable :**
Joint capot défectueux, condensation interne chaudière, fuite eau, humidité excessive local, mauvaise ventilation, température extérieure très basse (condensation).

**Étapes de résolution :**

1. **SÉCURITÉ IMMÉDIATE**
   - Couper alimentation 230V IMMÉDIATEMENT
   - Risque court-circuit et incendie
   - Laisser sécher avant toute manipulation
   - Ne pas remettre sous tension avant séchage complet

2. **Diagnostic origine humidité**
   - **Condensation interne chaudière** :
     - Température extérieure très basse
     - Chaudière condensation : fuite siphon, échangeur
   - **Fuite eau** :
     - Vérifier absence fuite sur corps chaudière
     - Contrôler siphon condensats
     - Vérifier échangeur (pas de perforation)
   - **Humidité ambiante** :
     - Local chaufferie mal ventilé
     - Infiltration eau pluie
     - Remontée humidité murs

3. **Séchage carte électronique**
   - Retirer carte de la chaudière
   - Séchage doux :
     - **Méthode 1** : air libre, 24-48h (plus sûr)
     - **Méthode 2** : air chaud (sèche-cheveux position tiède, distance 30cm)
     - **Méthode 3** : déshumidificateur ou absorbeur humidité (sachets silice)
   - **JAMAIS** :
     - Four (risque fusion composants)
     - Chaleur excessive (> 60°C)
     - Soufflette air comprimé (propulse eau dans composants)

4. **Nettoyage carte humide**
   - Une fois sèche, nettoyer la carte :
     - Alcool isopropylique 99% + pinceau doux
     - Brosser délicatement composants et pistes
     - Insister sur zones oxydées
     - Laisser évaporer complètement (alcool s'évapore vite)
   - Si oxydation importante :
     - Utiliser produit nettoyant spécial électronique (Cramolin, Contact Cleaner)
     - Brosser avec brosse douce (brosse à dents souple)

5. **Contrôle dommages**
   - Inspecter composants :
     - **Oxydation cuivre** (vert) : sur pistes, pattes composants
     - **Corrosion** : composants gonflés, déformés
     - **Pistes coupées** : oxydation ayant coupé connexion
   - Tester continuité pistes critiques
   - Remplacer composants endommagés si possible

6. **Réparation oxydation pistes**
   - Si piste coupée par oxydation :
     - Gratter oxydation (cutter, grattoir)
     - Étamer piste (flux + soudure)
     - Si piste détruite : ponter avec fil émaillé fin
   - Travail minutieux, loupe nécessaire

7. **Protection carte (tropicalisation)**
   - Appliquer vernis de protection :
     - Vernis acrylique ou polyuréthane
     - Spray "Conformal Coating" (Electrolube, MG Chemicals)
     - Protège contre humidité future
   - Application :
     - Masquer connecteurs (ruban adhésif)
     - Pulvériser fine couche uniforme
     - Sécher 24h avant remontage
     - Ne pas tropicaliser : relais, fusibles, potentiomètres

8. **Réparation étanchéité chaudière**
   - **Joint capot** :
     - Vérifier état joint mousse
     - Remplacer si comprimé, déchiré
     - Appliquer joint mousse adhésif si absent
   - **Presse-étoupes** :
     - Vérifier étanchéité passages câbles
     - Resserrer ou remplacer joints
   - **Siphon condensats** (chaudière condensation) :
     - Vérifier absence fuite siphon
     - Vérifier niveau eau siphon (barrière vapeur)
     - Nettoyer et remplir si nécessaire

9. **Amélioration ventilation**
   - Vérifier ventilation local chaufferie :
     - Grilles aération non obstruées
     - Aération basse/haute conforme
   - Si humidité excessive :
     - Installer VMC ou ventilation mécanique
     - Déshumidificateur local
     - Traiter infiltrations eau

10. **Test remise en service**
    - Remonter carte une fois COMPLÈTEMENT sèche
    - Vérifier étanchéité capot (joint bien positionné)
    - Remettre alimentation
    - Observer comportement :
      - Pas de fumée, odeur
      - Affichage normal
      - Pas de court-circuit (fusible tient)
    - Surveiller première heure fonctionnement

**Prévention :**
- Contrôle annuel joint capot
- Vérification étanchéité siphon (condensation)
- Ventilation correcte local chaufferie
- Éviter installation en local très humide (cave humide)
- Tropicalisation préventive carte si zone humide

**Signes précurseurs :**
- Buée sur vitre afficheur
- Odeur moisi à l'ouverture capot
- Corrosion visible vis, composants métalliques
- Fusibles grillent fréquemment

**Cas irréparable :**
- Si oxydation massive, composants CMS corrodés : remplacement carte nécessaire
- Si pistes multicouches coupées (PCB > 2 couches) : réparation très difficile
- Évaluer coût réparation vs remplacement carte

**Produits recommandés :**
- Alcool isopropylique 99% (nettoyage)
- Vernis tropicalisation (MG Chemicals 422B, Electrolube APL)
- Nettoyant contact (CRC Contact Cleaner)
- Absorbeur humidité (sachets silice gel)

---

## FACT-CHAUD-100: Mise à jour firmware nécessaire

| Champ | Valeur |
|-------|--------|
| **ID** | FACT-CHAUD-100 |
| **Catégorie** | Électronique & Cartes de Commande |
| **Système** | Logiciel embarqué (firmware) |
| **Gravité** | Faible à Moyenne |
| **Marques** | Multi-marques (surtout récentes) |

**Symptômes :**
- Bugs répétés, comportements illogiques
- Erreurs non justifiées (composants OK)
- Incompatibilité avec nouveau thermostat/régulation
- Fonctionnalités manquantes documentées notice
- Code erreur "Software Error" ou similaire
- Problèmes résolus par mise à jour selon constructeur

**Cause racine probable :**
Firmware obsolète, bugs connus corrigés version ultérieure, évolution normes (communication, sécurité), nouvelles fonctionnalités ajoutées, incompatibilité matérielle.

**Étapes de résolution :**

1. **Identification version firmware actuelle**
   - Accéder menu diagnostic/service
   - Chercher information version :
     - "Software version", "Firmware", "SW Ver"
     - Exemple : "v2.15", "FW 03.04.02"
   - Noter version exacte
   - Noter aussi : modèle chaudière, référence carte

2. **Vérification disponibilité mise à jour**
   - Consulter site web constructeur (espace pro SAV)
   - Rechercher firmware pour modèle exact
   - Vérifier notes de version (changelog) :
     - Corrections de bugs
     - Nouvelles fonctionnalités
     - Améliorations performances
   - Comparer version installée vs disponible

3. **Types de mise à jour selon marque**
   - **Mise à jour USB** (Vaillant, Saunier Duval, De Dietrich récents) :
     - Télécharger fichier firmware (.bin, .hex, .upd)
     - Copier sur clé USB (FAT32, racine)
     - Insérer clé USB dans port chaudière
     - Suivre procédure affichée ou menu service
   - **Mise à jour via outil SAV** (Frisquet Vision, Atlantic) :
     - Nécessite outil/logiciel propriétaire constructeur
     - Connexion PC ↔ chaudière (câble spécifique)
     - Lancer logiciel SAV
     - Suivre procédure guidée
   - **Remplacement EPROM** (anciennes chaudières) :
     - Composant mémoire (puce) à dessouder/remplacer
     - Commander EPROM avec nouveau firmware
     - Dessouder ancienne, souder nouvelle
     - Ou support EPROM (socket) si présent

4. **Préparation mise à jour**
   - **Sauvegarder paramètres** :
     - Photographier tous écrans paramètres
     - Noter sur papier : puissance, type gaz, courbe chauffe, etc.
     - Certaines MAJ effacent paramètres
   - **Vérifier alimentation stable** :
     - Chaudière connectée secteur stable
     - Pas de risque coupure pendant MAJ (risque brick)
   - **Charge batterie** (si applicable) :
     - Chaudière chaude, batterie auxiliaire chargée

5. **Procédure mise à jour USB (type Vaillant)**
   - Télécharger firmware exact modèle chaudière
   - Décompresser si fichier .zip
   - Formater clé USB en FAT32
   - Copier fichier firmware à la racine (pas dans dossier)
   - Éteindre chaudière
   - Insérer clé USB dans port (généralement sous capot, sur carte)
   - Allumer chaudière
   - Menu affiche "Update available" ou similaire
   - Valider mise à jour (bouton OK ou séquence boutons)
   - Patienter (5-15 minutes, ne PAS couper alimentation)
   - Redémarrage automatique après MAJ
   - Retirer clé USB

6. **Procédure mise à jour logiciel SAV**
   - Installer logiciel SAV sur PC (fourni constructeur)
   - Connecter câble PC ↔ chaudière :
     - USB vers connecteur carte
     - Ou adaptateur série/USB
   - Lancer logiciel
   - Détecter chaudière (communication établie)
   - Lire version actuelle firmware
   - Charger fichier firmware (.hex, .bin)
   - Lancer mise à jour (bouton "Update", "Flash")
   - Suivre progression (barre %)
   - Attendre fin (ne pas déconnecter)
   - Redémarrage chaudière

7. **Vérification après mise à jour**
   - Vérifier nouvelle version firmware (menu diagnostic)
   - Contrôler paramètres :
     - Si effacés : reprogrammer (voir notes sauvegarde)
     - Type chaudière, puissance, type gaz, etc.
   - Lancer auto-diagnostic si disponible
   - Tester fonctionnement :
     - Demande chauffage : OK
     - Demande ECS : OK
     - Modulation : OK
     - Sondes : valeurs cohérentes

8. **Résolution problèmes MAJ**
   - **MAJ échoue / erreur** :
     - Vérifier fichier firmware correspond exactement au modèle
     - Reformater clé USB en FAT32
     - Essayer autre clé USB (compatibilité)
     - Vérifier connexion câble (logiciel SAV)
   - **Chaudière brickée** (ne démarre plus) :
     - Tenter re-flash avec ancien firmware
     - Ou mode recovery (selon marque, séquence boutons spéciale)
     - Contacter SAV constructeur (intervention technicien)
   - **Paramètres effacés** :
     - Reprogrammer selon notes sauvegarde
     - Lancer assistant première mise en service

9. **Cas particuliers par marque**
   - **Vaillant ecoTEC** :
     - MAJ via clé USB ou logiciel DIALOGplus
     - Port USB sous capot avant
   - **Saunier Duval** :
     - MAJ via clé USB (modèles récents)
     - Fichier .upd sur clé FAT32
   - **Frisquet Eco Radio Vision** :
     - MAJ via logiciel Frisquet SAV
     - Connexion USB → module Vision
   - **De Dietrich** :
     - MAJ via logiciel Diematic SAV
     - Certains modèles : carte SD
   - **Elm Leblanc** :
     - MAJ rare, souvent remplacement carte
   - **Atlantic** :
     - Logiciel Cozytouch Pro (modèles connectés)

10. **Quand est-ce nécessaire ?**
    - Bugs répétés non résolus par interventions matérielles
    - Code erreur "Update required" affiché
    - Installation nouveau thermostat incompatible avec firmware ancien
    - Constructeur publie MAJ corrective pour problème connu
    - Évolution normes (ex: ErP, éco-conception)

**Prévention :**
- Vérifier firmware à chaque entretien annuel
- S'abonner newsletter SAV constructeur (alertes MAJ)
- Conserver clé USB + fichiers firmware (stock préventif)
- Documenter versions firmware installations

**Risques :**
- **Brick** (carte inutilisable) si :
  - Coupure alimentation pendant MAJ
  - Mauvais fichier firmware
  - Erreur procédure
- Toujours sauvegarder paramètres avant MAJ
- Ne jamais interrompre mise à jour en cours

**Coût :**
- Mise à jour firmware : gratuite (fichier)
- Logiciel SAV : gratuit (constructeurs) ou payant (selon marque)
- Câble interface : 20-100€ (si non fourni)
- Intervention technicien SAV : 100-200€ (si nécessaire assistance)

**Outils nécessaires :**
- Clé USB (FAT32, 1-4 Go suffisant)
- Ou PC + logiciel SAV + câble interface
- Documentation procédure (notice SAV)
- Notes paramètres chaudière

**Ressources :**
- Sites constructeurs (espace pro/SAV)
- Forums chauffagistes (partage expériences MAJ)
- SAV constructeur (assistance téléphonique)

---

*Fin du fichier 07_Electronique_Cartes.md*

**Retour à l'index :** [Knowledge_Base_Chaudieres.md](Knowledge_Base_Chaudieres.md)
