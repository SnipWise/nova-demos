# Guide RAG pour Génération de Diagrammes ASCII avec Petits LLMs

## SECTION 1: INSTRUCTIONS SYSTÈME

### Instruction système minimaliste (pour modèles 4B)
```
Tu es un assistant spécialisé dans la création de diagrammes ASCII.

RÈGLES:
1. Utilise UNIQUEMENT les caractères de boîte: ┌ ┐ └ ┘ ─ │ ├ ┤ ┬ ┴ ┼
2. Utilise les flèches: ▶ ◀ ▲ ▼ → ← ↑ ↓
3. Entoure le diagramme avec ```
4. Copie la STRUCTURE des exemples fournis

EXEMPLES SIMILAIRES:
{retrieved_diagrams}

DEMANDE UTILISATEUR:
{user_query}

Génère le diagramme en suivant la structure des exemples.
```

### Instruction système standard (pour modèles 7B-8B)
```
Tu es un expert en création de diagrammes ASCII techniques.

## Caractères autorisés
Boîtes: ┌ ┐ └ ┘ ─ │ ├ ┤ ┬ ┴ ┼
Flèches: ▶ ◀ ▲ ▼ → ← ↑ ↓ ──▶ ◀──
Décisions: ╱ ╲
Jointures: ═ ╧ ╤ ╪

## Méthode
1. Analyse les exemples similaires fournis
2. Identifie la structure qui correspond le mieux
3. Adapte les labels au contexte demandé
4. Garde des proportions cohérentes
5. Aligne verticalement et horizontalement

## Exemples de référence (récupérés par RAG)
{retrieved_diagrams}

## Demande utilisateur
{user_query}

## Instructions
- Copie la STRUCTURE d'un exemple similaire
- Modifie uniquement les LABELS
- Conserve l'alignement et les espacements
- Encadre le code avec ```
```

### Instruction système détaillée (pour modèles 8B optimisés)
```
<role>
Tu es DiagramBot, un assistant spécialisé dans la génération de diagrammes ASCII professionnels.
</role>

<capabilities>
- Diagrammes de flux (flowcharts)
- Architectures système
- Diagrammes de séquence
- Structures de données
- Schémas réseau
- Pipelines CI/CD
- Diagrammes UML simplifiés
</capabilities>

<character_reference>
## Caractères de boîte
┌───┐  Coin supérieur gauche/droit
│   │  Barre verticale
└───┘  Coin inférieur gauche/droit
├───┤  Jonction T gauche/droite
┬ ┴    Jonction T haut/bas
┼      Croisement

## Flèches
──▶ ◀── ───▶ ◀───  Flèches horizontales
 │      │          Lignes verticales
 ▼      ▲          Flèches verticales

## Décisions (losange simulé)
   ┌───────┐
  ╱         ╲
 ╱ Condition ╲
 ╲    ?      ╱
  ╲         ╱
   └───────┘

## Parallélisme
═══╧═══  Fork/Join
</character_reference>

<methodology>
1. ANALYSE: Comprendre le type de diagramme demandé
2. MATCH: Trouver l'exemple le plus similaire dans les références
3. STRUCTURE: Copier la structure de base
4. ADAPTATION: Modifier les labels pour le contexte
5. VALIDATION: Vérifier l'alignement des caractères
</methodology>

<similar_examples>
{retrieved_diagrams}
</similar_examples>

<user_request>
{user_query}
</user_request>

<output_format>
1. Identifie le type de diagramme (1 ligne)
2. Génère le diagramme dans un bloc ```
3. Explication courte (optionnel, 1-2 lignes)
</output_format>
```

---

## SECTION 2: PROMPTS D'EXEMPLES (FEW-SHOT)

### Exemple 1: Flux simple
```
UTILISATEUR: Crée un diagramme montrant: entrée → validation → traitement → sortie

ASSISTANT: Type: Flux linéaire horizontal

```
┌─────────┐     ┌───────────┐     ┌───────────┐     ┌─────────┐
│ Entrée  │────▶│ Validation│────▶│Traitement │────▶│ Sortie  │
└─────────┘     └───────────┘     └───────────┘     └─────────┘
```
```

### Exemple 2: Décision binaire
```
UTILISATEUR: Fais un flowchart: si l'utilisateur est admin alors accès complet, sinon accès limité

ASSISTANT: Type: Flux avec décision binaire

```
┌─────────────────┐
│  Utilisateur    │
│  se connecte    │
└────────┬────────┘
         │
         ▼
    ┌─────────┐
   ╱  Admin?   ╲
   ╲           ╱
    └────┬────┘
    Oui  │  Non
    ┌────┴────┐
    │         │
    ▼         ▼
┌───────┐ ┌────────┐
│ Accès │ │ Accès  │
│Complet│ │ Limité │
└───────┘ └────────┘
```
```

### Exemple 3: Architecture
```
UTILISATEUR: Architecture avec frontend, API, et base de données

ASSISTANT: Type: Architecture 3-tiers

```
┌─────────────────────────────────────┐
│           Frontend                   │
│  ┌──────────┐    ┌──────────┐       │
│  │   Web    │    │  Mobile  │       │
│  └────┬─────┘    └────┬─────┘       │
└───────┼───────────────┼─────────────┘
        │               │
        └───────┬───────┘
                │
                ▼
┌─────────────────────────────────────┐
│            API Layer                 │
│         ┌──────────┐                │
│         │  REST    │                │
│         │   API    │                │
│         └────┬─────┘                │
└──────────────┼──────────────────────┘
               │
               ▼
┌─────────────────────────────────────┐
│            Database                  │
│         ┌──────────┐                │
│         │PostgreSQL│                │
│         └──────────┘                │
└─────────────────────────────────────┘
```
```

### Exemple 4: Séquence
```
UTILISATEUR: Diagramme de séquence pour login: user envoie credentials, server vérifie, retourne token

ASSISTANT: Type: Diagramme de séquence

```
┌──────┐          ┌──────┐          ┌──────┐
│ User │          │Server│          │  DB  │
└──┬───┘          └──┬───┘          └──┬───┘
   │                 │                 │
   │  credentials    │                 │
   │────────────────▶│                 │
   │                 │  query user     │
   │                 │────────────────▶│
   │                 │                 │
   │                 │  user data      │
   │                 │◀────────────────│
   │                 │                 │
   │  JWT token      │                 │
   │◀────────────────│                 │
   │                 │                 │
```
```

### Exemple 5: Pipeline
```
UTILISATEUR: Pipeline CI/CD avec build, test, deploy

ASSISTANT: Type: Pipeline CI/CD

```
┌────────┐     ┌────────┐     ┌────────┐     ┌────────┐
│  Code  │────▶│  Build │────▶│  Test  │────▶│ Deploy │
│ Commit │     │        │     │        │     │        │
└────────┘     └────────┘     └────────┘     └────────┘
    │              │              │              │
    ▼              ▼              ▼              ▼
┌────────┐     ┌────────┐     ┌────────┐     ┌────────┐
│  Git   │     │ Docker │     │ pytest │     │  K8s   │
└────────┘     └────────┘     └────────┘     └────────┘
```
```

### Exemple 6: Arbre
```
UTILISATEUR: Structure de fichiers: src contient components et utils, components contient Header et Footer

ASSISTANT: Type: Arbre de fichiers

```
         ┌─────┐
         │ src │
         └──┬──┘
      ┌─────┴─────┐
      │           │
┌─────┴─────┐ ┌───┴───┐
│components │ │ utils │
└─────┬─────┘ └───────┘
   ┌──┴──┐
   │     │
┌──┴──┐┌─┴────┐
│Header││Footer│
└─────┘└──────┘
```
```

### Exemple 7: Boucle
```
UTILISATEUR: Boucle de retry: essayer, si échec et retry < 3 alors réessayer, sinon erreur

ASSISTANT: Type: Flux avec boucle

```
┌─────────────────┐
│     Start       │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│    retry = 0    │
└────────┬────────┘
         │
         ▼
┌─────────────────┐◀─────────────┐
│     Try         │              │
│   Request       │              │
└────────┬────────┘              │
         │                       │
         ▼                       │
    ┌─────────┐                  │
   ╱ Success?  ╲   Non     ┌─────┴─────┐
   ╲           ╱──────────▶│ retry++   │
    └────┬────┘            └─────┬─────┘
         │ Oui                   │
         │                       ▼
         │                 ┌───────────┐
         │                ╱ retry < 3? ╲
         │                ╲            ╱
         │                 └─────┬────┘
         │            Oui ───────┘
         │                       │ Non
         ▼                       ▼
┌─────────────────┐     ┌─────────────────┐
│     Success     │     │     Error       │
└─────────────────┘     └─────────────────┘
```
```

### Exemple 8: Microservices
```
UTILISATEUR: 3 microservices (users, orders, products) avec un API gateway

ASSISTANT: Type: Architecture microservices

```
┌───────────────┐
│    Client     │
└───────┬───────┘
        │
        ▼
┌───────────────────────────────────────┐
│           API Gateway                  │
└───────┬───────────┬───────────┬───────┘
        │           │           │
        ▼           ▼           ▼
┌───────────┐ ┌───────────┐ ┌───────────┐
│  Users    │ │  Orders   │ │ Products  │
│  Service  │ │  Service  │ │  Service  │
└─────┬─────┘ └─────┬─────┘ └─────┬─────┘
      │             │             │
      ▼             ▼             ▼
┌─────────┐   ┌─────────┐   ┌─────────┐
│Users DB │   │Orders DB│   │Prod. DB │
└─────────┘   └─────────┘   └─────────┘
```
```

### Exemple 9: Agent LLM
```
UTILISATEUR: Agent avec LLM qui utilise 3 outils: search, calculator, code

ASSISTANT: Type: Agent LLM avec outils

```
┌─────────────────────────────────────────────────────┐
│                   LLM Agent                          │
│  ┌───────────────────────────────────────────────┐  │
│  │                 LLM Core                       │  │
│  │            (reasoning engine)                  │  │
│  └──────────────────────┬────────────────────────┘  │
│                         │                           │
│         ┌───────────────┼───────────────┐           │
│         │               │               │           │
│         ▼               ▼               ▼           │
│    ┌─────────┐    ┌─────────┐    ┌─────────┐       │
│    │ Search  │    │  Calc   │    │  Code   │       │
│    │  Tool   │    │  Tool   │    │  Tool   │       │
│    └─────────┘    └─────────┘    └─────────┘       │
│                                                      │
└─────────────────────────────────────────────────────┘
```
```

### Exemple 10: État machine
```
UTILISATEUR: Machine d'état pour une lampe: off -> on (bouton) -> off (bouton)

ASSISTANT: Type: Machine d'état simple

```
                    bouton
            ┌──────────────────────┐
            │                      │
            ▼                      │
       ┌─────────┐            ┌────┴────┐
       │   OFF   │───bouton──▶│   ON    │
       │         │            │         │
       └─────────┘            └─────────┘
```
```

---

## SECTION 3: MÉTADONNÉES POUR L'INDEXATION RAG

### Format de métadonnées JSON pour chaque diagramme
```json
{
  "id": "001",
  "title": "Flux linéaire simple (3 étapes)",
  "category": "flowchart",
  "subcategory": "linear",
  "keywords": [
    "flux", "linéaire", "simple", "horizontal",
    "entrée", "sortie", "processus", "étapes",
    "input", "output", "process", "flow"
  ],
  "complexity": "simple",
  "elements": ["boxes", "arrows_horizontal"],
  "use_cases": [
    "pipeline simple",
    "transformation de données",
    "étapes séquentielles"
  ],
  "similar_to": ["002", "010", "043"]
}
```

### Index complet des catégories
```json
{
  "categories": {
    "flowchart": {
      "description": "Diagrammes de flux et processus",
      "ids": ["001", "002", "003", "004", "005", "006", "007", "008", "009", "010"]
    },
    "architecture": {
      "description": "Architectures système et logicielle",
      "ids": ["011", "012", "013", "014", "015", "016", "017", "018", "019", "020"]
    },
    "sequence": {
      "description": "Diagrammes de séquence et interactions",
      "ids": ["021", "022", "023", "024", "025", "026", "027", "028"]
    },
    "data_structure": {
      "description": "Structures de données et collections",
      "ids": ["029", "030", "031", "032", "033", "034", "035", "036", "037", "038"]
    },
    "network": {
      "description": "Topologies réseau et sécurité",
      "ids": ["039", "040", "041", "042"]
    },
    "cicd": {
      "description": "Pipelines CI/CD et déploiement",
      "ids": ["043", "044", "045", "046"]
    },
    "ml": {
      "description": "Machine Learning et AI",
      "ids": ["047", "048", "049", "050"]
    },
    "state": {
      "description": "Machines d'état et cycles de vie",
      "ids": ["051", "052", "053"]
    },
    "uml": {
      "description": "Diagrammes UML simplifiés",
      "ids": ["054", "055", "056", "057"]
    },
    "database": {
      "description": "Schémas de base de données",
      "ids": ["058", "059", "060"]
    },
    "kubernetes": {
      "description": "Conteneurs et orchestration",
      "ids": ["061", "062", "063", "064"]
    },
    "agent": {
      "description": "Agents LLM et patterns AI",
      "ids": ["065", "066", "067", "068", "069"]
    },
    "misc": {
      "description": "Diagrammes divers",
      "ids": ["070", "071", "072", "073", "074", "075"]
    }
  }
}
```

---

## SECTION 4: STRATÉGIES DE RETRIEVAL

### Stratégie 1: Par mots-clés simples
```python
# Pour modèles très petits (4B), utiliser une correspondance simple
def simple_keyword_match(query, diagrams_index):
    query_words = query.lower().split()
    scores = {}
    
    for diagram_id, metadata in diagrams_index.items():
        score = sum(1 for word in query_words 
                   if word in metadata['keywords'])
        scores[diagram_id] = score
    
    # Retourner les 3 meilleurs
    return sorted(scores.items(), key=lambda x: x[1], reverse=True)[:3]
```

### Stratégie 2: Par embeddings sémantiques
```python
# Pour un retrieval plus précis
from sentence_transformers import SentenceTransformer

model = SentenceTransformer('paraphrase-multilingual-MiniLM-L12-v2')

def semantic_search(query, diagram_embeddings, top_k=3):
    query_embedding = model.encode(query)
    
    # Calculer similarité cosinus
    similarities = cosine_similarity([query_embedding], diagram_embeddings)[0]
    
    # Retourner les top_k
    top_indices = similarities.argsort()[-top_k:][::-1]
    return top_indices
```

### Stratégie 3: Hybride (recommandée)
```python
def hybrid_retrieval(query, diagrams, alpha=0.5):
    # 1. Keyword matching
    keyword_scores = keyword_match(query, diagrams)
    
    # 2. Semantic similarity
    semantic_scores = semantic_search(query, diagrams)
    
    # 3. Combiner avec pondération
    final_scores = {}
    for diagram_id in diagrams:
        kw = keyword_scores.get(diagram_id, 0)
        sem = semantic_scores.get(diagram_id, 0)
        final_scores[diagram_id] = alpha * kw + (1 - alpha) * sem
    
    return sorted(final_scores.items(), key=lambda x: x[1], reverse=True)[:3]
```

---

## SECTION 5: CONFIGURATION POUR DIFFÉRENTS MODÈLES

### Configuration Phi-3 Mini (3.8B)
```yaml
model_config:
  name: "microsoft/Phi-3-mini-4k-instruct"
  max_tokens: 1024
  temperature: 0.3
  
rag_config:
  num_examples: 2  # Moins d'exemples pour contexte court
  system_prompt: "minimal"  # Utiliser le prompt minimaliste
  
prompt_template: |
  Exemples:
  {examples}
  
  Crée un diagramme pour: {query}
```

### Configuration Llama 3.2 (8B)
```yaml
model_config:
  name: "meta-llama/Llama-3.2-8B-Instruct"
  max_tokens: 2048
  temperature: 0.2
  
rag_config:
  num_examples: 3
  system_prompt: "standard"
  
prompt_template: |
  <|system|>
  {system_prompt}
  <|user|>
  {query}
  <|assistant|>
```

### Configuration Mistral (7B)
```yaml
model_config:
  name: "mistralai/Mistral-7B-Instruct-v0.3"
  max_tokens: 2048
  temperature: 0.1
  
rag_config:
  num_examples: 3
  system_prompt: "standard"
  include_categories: true
  
prompt_template: |
  [INST] {system_prompt}
  
  Catégorie détectée: {category}
  Exemples similaires:
  {examples}
  
  Demande: {query} [/INST]
```

### Configuration Qwen 2.5 (7B)
```yaml
model_config:
  name: "Qwen/Qwen2.5-7B-Instruct"  
  max_tokens: 2048
  temperature: 0.2
  
rag_config:
  num_examples: 4  # Qwen gère bien les contextes plus longs
  system_prompt: "detailed"
  
prompt_template: |
  <|im_start|>system
  {system_prompt}
  <|im_end|>
  <|im_start|>user
  {query}
  <|im_end|>
  <|im_start|>assistant
```

---

## SECTION 6: EXEMPLES DE REQUÊTES ET RÉPONSES ATTENDUES

### Requêtes de test avec réponses attendues

```yaml
tests:
  - query: "Crée un diagramme pour une API REST avec authentification"
    expected_category: "architecture"
    expected_similar: ["011", "012", "022"]
    expected_elements: ["boxes", "arrows", "layers"]
    
  - query: "Pipeline de données ETL"
    expected_category: "flowchart"
    expected_similar: ["010", "043", "047"]
    expected_elements: ["boxes", "horizontal_flow"]
    
  - query: "Architecture multi-agents avec orchestrateur"
    expected_category: "agent"
    expected_similar: ["066", "065", "067"]
    expected_elements: ["boxes", "bidirectional_arrows"]
    
  - query: "Schéma de base de données utilisateurs et commandes"
    expected_category: "database"
    expected_similar: ["058", "059"]
    expected_elements: ["tables", "relations", "keys"]
    
  - query: "Boucle while avec condition de sortie"
    expected_category: "flowchart"
    expected_similar: ["004", "050"]
    expected_elements: ["boxes", "loop", "decision"]
```

---

## SECTION 7: IMPLÉMENTATION COMPLÈTE

### Script Python complet pour RAG de diagrammes
```python
#!/usr/bin/env python3
"""
RAG System for ASCII Diagram Generation with Small LLMs
"""

import json
import re
from pathlib import Path
from typing import List, Dict, Tuple
import numpy as np

# Pour les embeddings (optionnel, fallback sur keywords)
try:
    from sentence_transformers import SentenceTransformer
    HAS_EMBEDDINGS = True
except ImportError:
    HAS_EMBEDDINGS = False


class DiagramRAG:
    def __init__(self, diagrams_path: str, config: dict = None):
        self.diagrams = self._load_diagrams(diagrams_path)
        self.config = config or self._default_config()
        
        if HAS_EMBEDDINGS:
            self.embedder = SentenceTransformer(
                'paraphrase-multilingual-MiniLM-L12-v2'
            )
            self._compute_embeddings()
    
    def _default_config(self) -> dict:
        return {
            "num_examples": 3,
            "use_embeddings": HAS_EMBEDDINGS,
            "hybrid_alpha": 0.5
        }
    
    def _load_diagrams(self, path: str) -> Dict:
        """Parse le fichier markdown et extrait les diagrammes"""
        content = Path(path).read_text(encoding='utf-8')
        diagrams = {}
        
        # Pattern pour extraire les diagrammes
        pattern = r'### (\d+) - (.+?)\n```\n([\s\S]+?)\n```'
        
        for match in re.finditer(pattern, content):
            diagram_id = match.group(1)
            title = match.group(2)
            code = match.group(3)
            
            diagrams[diagram_id] = {
                "id": diagram_id,
                "title": title,
                "code": code,
                "keywords": self._extract_keywords(title, code)
            }
        
        return diagrams
    
    def _extract_keywords(self, title: str, code: str) -> List[str]:
        """Extrait les mots-clés du titre et du contenu"""
        text = f"{title} {code}".lower()
        # Nettoyer les caractères spéciaux
        text = re.sub(r'[┌┐└┘─│├┤┬┴┼▶◀▲▼→←↑↓╱╲═╧╤╪◆◇△]', ' ', text)
        words = re.findall(r'[a-zàâäéèêëïîôùûüç]+', text)
        return list(set(words))
    
    def _compute_embeddings(self):
        """Pré-calcule les embeddings pour tous les diagrammes"""
        texts = [
            f"{d['title']} {' '.join(d['keywords'])}"
            for d in self.diagrams.values()
        ]
        self.embeddings = self.embedder.encode(texts)
        self.diagram_ids = list(self.diagrams.keys())
    
    def _keyword_search(self, query: str) -> Dict[str, float]:
        """Recherche par mots-clés"""
        query_words = set(query.lower().split())
        scores = {}
        
        for diagram_id, diagram in self.diagrams.items():
            keywords = set(diagram['keywords'])
            intersection = query_words & keywords
            scores[diagram_id] = len(intersection) / max(len(query_words), 1)
        
        return scores
    
    def _semantic_search(self, query: str) -> Dict[str, float]:
        """Recherche sémantique par embeddings"""
        if not HAS_EMBEDDINGS:
            return {}
        
        query_emb = self.embedder.encode([query])[0]
        
        # Similarité cosinus
        similarities = np.dot(self.embeddings, query_emb) / (
            np.linalg.norm(self.embeddings, axis=1) * np.linalg.norm(query_emb)
        )
        
        return {
            self.diagram_ids[i]: float(similarities[i])
            for i in range(len(self.diagram_ids))
        }
    
    def retrieve(self, query: str, top_k: int = None) -> List[Dict]:
        """Récupère les diagrammes les plus pertinents"""
        top_k = top_k or self.config["num_examples"]
        
        # Scores par mots-clés
        kw_scores = self._keyword_search(query)
        
        # Scores sémantiques (si disponible)
        if self.config["use_embeddings"] and HAS_EMBEDDINGS:
            sem_scores = self._semantic_search(query)
            alpha = self.config["hybrid_alpha"]
            
            final_scores = {
                did: alpha * kw_scores.get(did, 0) + 
                     (1 - alpha) * sem_scores.get(did, 0)
                for did in self.diagrams
            }
        else:
            final_scores = kw_scores
        
        # Trier et retourner top_k
        sorted_ids = sorted(
            final_scores.items(), 
            key=lambda x: x[1], 
            reverse=True
        )[:top_k]
        
        return [self.diagrams[did] for did, _ in sorted_ids]
    
    def build_prompt(
        self, 
        query: str, 
        system_prompt: str,
        examples: List[Dict] = None
    ) -> str:
        """Construit le prompt complet pour le LLM"""
        if examples is None:
            examples = self.retrieve(query)
        
        # Formater les exemples
        examples_text = "\n\n".join([
            f"### {ex['title']}\n```\n{ex['code']}\n```"
            for ex in examples
        ])
        
        # Remplacer les placeholders
        prompt = system_prompt.replace("{retrieved_diagrams}", examples_text)
        prompt = prompt.replace("{user_query}", query)
        
        return prompt


# Exemple d'utilisation
if __name__ == "__main__":
    rag = DiagramRAG("diagram_library.md")
    
    query = "Crée une architecture microservices avec 3 services"
    
    system_prompt = """Tu es un assistant spécialisé dans la création de diagrammes ASCII.

EXEMPLES SIMILAIRES:
{retrieved_diagrams}

DEMANDE: {user_query}

Génère le diagramme en suivant la structure des exemples."""
    
    prompt = rag.build_prompt(query, system_prompt)
    print(prompt)
```

---

## SECTION 8: CONSEILS POUR OPTIMISER LES RÉSULTATS

### 1. Pré-traitement des requêtes
```python
def preprocess_query(query: str) -> str:
    """Normalise la requête pour un meilleur matching"""
    # Synonymes courants
    synonyms = {
        "flowchart": "flux diagramme",
        "architecture": "système structure",
        "sequence": "séquence interaction",
        "microservice": "microservices service",
        "pipeline": "flux étapes",
        "database": "base données db",
        "api": "interface rest",
        "loop": "boucle while for",
        "condition": "décision if else",
        "tree": "arbre hiérarchie",
    }
    
    query_lower = query.lower()
    for key, expansion in synonyms.items():
        if key in query_lower:
            query_lower = f"{query_lower} {expansion}"
    
    return query_lower
```

### 2. Post-traitement des réponses
```python
def validate_diagram(response: str) -> Tuple[bool, str]:
    """Valide que la réponse contient un diagramme valide"""
    # Extraire le bloc de code
    code_match = re.search(r'```\n?([\s\S]+?)\n?```', response)
    
    if not code_match:
        return False, "Pas de bloc de code trouvé"
    
    code = code_match.group(1)
    
    # Vérifier la présence de caractères de diagramme
    diagram_chars = set('┌┐└┘─│├┤┬┴┼▶◀▲▼')
    if not any(c in code for c in diagram_chars):
        return False, "Pas de caractères de diagramme"
    
    # Vérifier l'alignement basique
    lines = code.split('\n')
    if len(lines) < 2:
        return False, "Diagramme trop court"
    
    return True, code
```

### 3. Fallback et retry
```python
def generate_with_fallback(
    rag: DiagramRAG,
    llm,
    query: str,
    max_retries: int = 3
) -> str:
    """Génère avec retry et augmentation du contexte"""
    
    for attempt in range(max_retries):
        # Augmenter le nombre d'exemples à chaque retry
        num_examples = 2 + attempt
        examples = rag.retrieve(query, top_k=num_examples)
        
        prompt = rag.build_prompt(query, SYSTEM_PROMPT, examples)
        response = llm.generate(prompt)
        
        is_valid, result = validate_diagram(response)
        
        if is_valid:
            return result
        
        print(f"Tentative {attempt + 1} échouée: {result}")
    
    # Fallback: retourner le diagramme le plus similaire tel quel
    best_match = rag.retrieve(query, top_k=1)[0]
    return best_match['code']
```
