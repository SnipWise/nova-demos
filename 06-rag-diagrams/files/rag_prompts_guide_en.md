# RAG Guide for ASCII Diagram Generation with Small LLMs

## SECTION 1: SYSTEM INSTRUCTIONS

### Minimal system instruction (for 4B models)
```
You are an assistant specialized in creating ASCII diagrams.

RULES:
1. Use ONLY box characters: ┌ ┐ └ ┘ ─ │ ├ ┤ ┬ ┴ ┼
2. Use arrows: ▶ ◀ ▲ ▼ → ← ↑ ↓
3. Surround the diagram with ```
4. Copy the STRUCTURE from provided examples

SIMILAR EXAMPLES:
{retrieved_diagrams}

USER REQUEST:
{user_query}

Generate the diagram following the structure of the examples.
```

### Standard system instruction (for 7B-8B models)
```
You are an expert in creating technical ASCII diagrams.

## Allowed characters
Boxes: ┌ ┐ └ ┘ ─ │ ├ ┤ ┬ ┴ ┼
Arrows: ▶ ◀ ▲ ▼ → ← ↑ ↓ ──▶ ◀──
Decisions: ╱ ╲
Jointures: ═ ╧ ╤ ╪

## Method
1. Analyze the provided similar examples
2. Identify the structure that best matches
3. Adapt labels to the requested context
4. Keep consistent proportions
5. Align vertically and horizontally

## Reference examples (retrieved by RAG)
{retrieved_diagrams}

## User request
{user_query}

## Instructions
- Copy the STRUCTURE from a similar example
- Modify ONLY the LABELS
- Preserve alignment and spacing
- Wrap code with ```
```

### Detailed system instruction (for optimized 8B models)
```
<role>
You are DiagramBot, an assistant specialized in generating professional ASCII diagrams.
</role>

<capabilities>
- Flowcharts
- System architectures
- Sequence diagrams
- Data structures
- Network schemas
- CI/CD pipelines
- Simplified UML diagrams
</capabilities>

<character_reference>
## Box characters
┌───┐  Top left/right corner
│   │  Vertical bar
└───┘  Bottom left/right corner
├───┤  T-junction left/right
┬ ┴    T-junction top/bottom
┼      Crossroads

## Arrows
──▶ ◀── ───▶ ◀───  Horizontal arrows
 │      │          Vertical lines
 ▼      ▲          Vertical arrows

## Decisions (diamond simulated)
   ┌───────┐
  ╱         ╲
 ╱ Condition ╲
 ╲    ?      ╱
  ╲         ╱
   └───────┘

## Parallelism
═══╧═══  Fork/Join
</character_reference>

<methodology>
1. ANALYZE: Understand the type of diagram requested
2. MATCH: Find the most similar example in references
3. STRUCTURE: Copy the base structure
4. ADAPTATION: Modify labels for context
5. VALIDATION: Check character alignment
</methodology>

<alignment_rules>
CRITICAL: All lines in a box MUST have EXACTLY the same length.

Step-by-step alignment process:
1. Count characters in each line of the example
2. For top/bottom borders: copy EXACTLY (e.g., ┌────────────────────────┐ = 24 chars)
3. For middle lines: │ + content + spaces + │ = SAME total length
4. If label is shorter: add spaces to RIGHT before │
5. If label is longer: ABBREVIATE it, don't expand the box

Example - CORRECT alignment (all lines = 24 chars):
┌────────────────────────┐
│ Person                 │  (│ + " Person" + 17 spaces + │ = 24)
├────────────────────────┤
│ - name: string         │  (│ + " - name: string" + 7 spaces + │ = 24)
└────────────────────────┘

Example - WRONG alignment (inconsistent lengths):
┌────────────────────────┐  (24 chars)
│ Person │                   (10 chars) ❌ WRONG
├────────────────────────┤  (24 chars)

## Validation before output
Before sending response:
1. Count characters of EACH line
2. Verify all lines in same box have SAME length
3. If not: add/remove spaces until aligned
4. Only then: output the diagram
</alignment_rules>

<similar_examples>
{retrieved_diagrams}
</similar_examples>

<user_request>
{user_query}
</user_request>

<output_format>
1. Identify diagram type (1 line)
2. Generate diagram in ``` block
3. Short explanation (optional, 1-2 lines)
</output_format>
```

---

## SECTION 2: FEW-SHOT EXAMPLES

### Example 1: Simple flow
```
USER: Create a diagram showing: input → validation → processing → output

ASSISTANT: Type: Horizontal linear flow

```
┌─────────┐     ┌───────────┐     ┌───────────┐     ┌─────────┐
│ Input   │────▶│ Validation│────▶│ Processing│────▶│ Output  │
└─────────┘     └───────────┘     └───────────┘     └─────────┘
```
```

### Example 2: Binary decision
```
USER: Make a flowchart: if user is admin then full access, else limited access

ASSISTANT: Type: Flow with binary decision

```
┌─────────────────┐
│  User           │
│  logs in        │
└────────┬────────┘
         │
         ▼
    ┌─────────┐
   ╱  Admin?   ╲
   ╲           ╱
    └────┬────┘
    Yes  │  No
    ┌────┴────┐
    │         │
    ▼         ▼
┌───────┐ ┌────────┐
│ Full  │ │Limited │
│Access │ │ Access │
└───────┘ └────────┘
```
```

### Example 3: Architecture
```
USER: Architecture with frontend, API, and database

ASSISTANT: Type: 3-tier architecture

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

### Example 4: UML Class with Inheritance
```
USER: Create a Warrior class that inherits from Character class, Warrior has an attack() method, Character has name and life attributes

ASSISTANT: Type: UML Class Inheritance

```
         ┌────────────────────────┐
         │      Character         │
         ├────────────────────────┤
         │ - name: string         │
         │ - life: int            │
         ├────────────────────────┤
         │ + getName(): string    │
         │ + setName(n: string)   │
         │ + getLife(): int       │
         │ + setLife(l: int)      │
         └───────────┬────────────┘
                     △
                     │ inherits
                     │
         ┌───────────┴────────────┐
         │       Warrior          │
         ├────────────────────────┤
         │ - weapon: string       │
         ├────────────────────────┤
         │ + attack(): void       │
         │ + setWeapon(w: string) │
         └────────────────────────┘
```
```

---

## SECTION 3: ALIGNMENT TECHNIQUES FOR LLMS

### Technique 1: Character counting in prompt
```
IMPORTANT ALIGNMENT RULE:
Every line in a box must have EXACTLY the same number of characters.

To ensure this:
1. The top border ┌────────────────────────┐ has N characters
2. EVERY middle line must have N characters: │ + text + spaces + │
3. The bottom border └────────────────────────┘ has N characters

Example with 24 characters per line:
┌────────────────────────┐  ← count: 24
│ Person                 │  ← count: 24 (│ + " Person" + 17 spaces + │)
├────────────────────────┤  ← count: 24
│ - name: string         │  ← count: 24 (│ + " - name: string" + 7 spaces + │)
└────────────────────────┘  ← count: 24

If the text "Person" is 6 characters, you need:
- 1 char for │
- 1 space before text
- 6 chars for "Person"
- 15 spaces after text
- 1 char for │
= 24 total
```

### Technique 2: Template-based generation
```
Provide this template in the prompt:

TEMPLATE FOR CLASS DIAGRAM (each line = 24 chars):
┌────────────────────────┐
│ CLASSNAME              │
├────────────────────────┤
│ - attr1                │
│ - attr2                │
├────────────────────────┤
│ + method1()            │
│ + method2()            │
└────────────────────────┘

INSTRUCTIONS:
1. Replace CLASSNAME with actual class name
2. Add spaces after each line to reach 24 chars total
3. DO NOT change the border characters
4. DO NOT add or remove ─ characters from borders
```

### Technique 3: Post-generation validation
```python
def validate_and_fix_alignment(diagram: str) -> str:
    """Validate and auto-fix alignment issues"""
    lines = diagram.split('\n')

    # Find max line length
    max_length = max(len(line) for line in lines if line.strip())

    fixed_lines = []
    for line in lines:
        if not line.strip():
            fixed_lines.append('')
            continue

        # If line starts with box char
        if line.lstrip()[0] in '┌├└│':
            # Pad to max_length
            padding_needed = max_length - len(line)
            if line.rstrip().endswith('│'):
                # Insert spaces before closing │
                fixed_line = line.rstrip()[:-1] + ' ' * padding_needed + '│'
            elif line.rstrip().endswith(('┐', '┤', '┘')):
                # Extend border
                border_char = '─' if line.rstrip()[-1] in '┐┤┘' else line.rstrip()[-1]
                fixed_line = line.rstrip()[:-1] + border_char * padding_needed + line.rstrip()[-1]
            else:
                fixed_line = line
            fixed_lines.append(fixed_line)
        else:
            fixed_lines.append(line)

    return '\n'.join(fixed_lines)
```

### Technique 4: Explicit padding instructions
```
When generating a UML class diagram:

1. Determine the longest text line in the class
2. Set box width = longest_text + 4 (for │ + space + text + space + │)
3. For ALL other lines: add spaces to match this width

Step-by-step example:
- Class name: "Warrior" (7 chars)
- Longest attribute: "- weapon: string" (16 chars)
- Longest method: "+ setWeapon(w: string)" (21 chars)

Longest line = 21 chars
Box width = 21 + 4 = 25 chars total

Now generate:
┌───────────────────────┐  (25 chars: ┌ + 23×─ + ┐)
│ Warrior               │  (25 chars: │ + " Warrior" + 15 spaces + │)
├───────────────────────┤  (25 chars)
│ - weapon: string      │  (25 chars: │ + " - weapon: string" + 4 spaces + │)
├───────────────────────┤  (25 chars)
│ + attack(): void      │  (25 chars: │ + " + attack(): void" + 4 spaces + │)
│ + setWeapon(w: string)│  (25 chars: │ + " + setWeapon(w: string)" + 0 spaces + │)
└───────────────────────┘  (25 chars)
```

---

## SECTION 4: IMPLEMENTATION - IMPROVED RAG SYSTEM

### Enhanced DiagramRAG with alignment checking
```python
class DiagramRAG:
    # ... (previous code)

    def validate_alignment(self, diagram: str) -> Tuple[bool, List[str]]:
        """Check if diagram is properly aligned"""
        lines = diagram.split('\n')
        errors = []

        # Group lines by boxes
        current_box = []
        boxes = []

        for line in lines:
            if line.strip() and line.lstrip()[0] in '┌├└│':
                current_box.append(line)
                if line.lstrip()[0] in '└':  # End of box
                    boxes.append(current_box)
                    current_box = []

        # Check each box
        for i, box in enumerate(boxes):
            lengths = [len(line) for line in box]
            if len(set(lengths)) > 1:
                errors.append(f"Box {i+1}: inconsistent line lengths {set(lengths)}")

        return len(errors) == 0, errors

    def fix_alignment(self, diagram: str) -> str:
        """Auto-fix alignment issues"""
        lines = diagram.split('\n')
        fixed_lines = []

        i = 0
        while i < len(lines):
            line = lines[i]

            # Detect start of box
            if line.lstrip().startswith('┌'):
                # Find end of box
                box_lines = [line]
                i += 1
                while i < len(lines) and not lines[i].lstrip().startswith('└'):
                    box_lines.append(lines[i])
                    i += 1
                if i < len(lines):
                    box_lines.append(lines[i])

                # Fix this box
                max_len = max(len(l) for l in box_lines)
                for bl in box_lines:
                    if bl.rstrip().endswith('│'):
                        # Pad content line
                        spaces_needed = max_len - len(bl)
                        fixed_line = bl.rstrip()[:-1] + ' ' * spaces_needed + '│'
                    elif bl.rstrip().endswith(('┐', '┤', '┘')):
                        # Extend border
                        spaces_needed = max_len - len(bl)
                        border_end = bl.rstrip()[-1]
                        fixed_line = bl.rstrip()[:-1] + '─' * spaces_needed + border_end
                    else:
                        fixed_line = bl
                    fixed_lines.append(fixed_line)
            else:
                fixed_lines.append(line)

            i += 1

        return '\n'.join(fixed_lines)

    def build_prompt_with_alignment(
        self,
        query: str,
        system_prompt: str,
        examples: List[Dict] = None
    ) -> str:
        """Build prompt with enhanced alignment instructions"""
        if examples is None:
            examples = self.retrieve(query)

        # Add alignment validation to examples
        validated_examples = []
        for ex in examples:
            is_valid, errors = self.validate_alignment(ex['code'])
            if not is_valid:
                # Fix alignment before using as example
                ex['code'] = self.fix_alignment(ex['code'])
            validated_examples.append(ex)

        # Format examples
        examples_text = "\n\n".join([
            f"### {ex['title']}\n```\n{ex['code']}\n```"
            for ex in validated_examples
        ])

        # Enhanced system prompt with alignment rules
        alignment_rules = """
## CRITICAL ALIGNMENT RULES

1. Every line in a box MUST have the EXACT same character count
2. To achieve this:
   - Count characters in the top border (e.g., ┌────────┐ = 10 chars)
   - ALL middle lines must be: │ + content + padding spaces + │ = 10 chars
   - Bottom border must be: └────────┘ = 10 chars

3. Process for each line:
   a) Start with │ (1 char)
   b) Add 1 space
   c) Add your text
   d) Calculate: remaining = total_width - 2 - 1 - len(text)
   e) Add 'remaining' spaces
   f) End with │ (1 char)

4. DO NOT change border character counts
5. If text is too long: ABBREVIATE it, don't expand the box
6. If text is too short: ADD SPACES before the closing │

VALIDATION CHECK:
Before outputting, mentally verify:
- Line 1 length = Line 2 length = Line 3 length = ... = Line N length
- If not equal, fix by adjusting spaces
"""

        # Replace placeholders
        prompt = system_prompt + "\n\n" + alignment_rules
        prompt = prompt.replace("{retrieved_diagrams}", examples_text)
        prompt = prompt.replace("{user_query}", query)

        return prompt
```

---

## SECTION 5: RECOMMENDED CONFIGURATIONS

### For Qwen2.5 (7B-8B) - Best alignment performance
```yaml
model_config:
  name: "ai/qwen2.5:latest"
  max_tokens: 2048
  temperature: 0.0  # Zero temperature for deterministic alignment

rag_config:
  num_examples: 2  # Fewer examples = more room for alignment instructions
  use_alignment_validation: true
  auto_fix_examples: true

prompt_strategy: "detailed_with_alignment"
```

### For smaller models (4B)
```yaml
model_config:
  name: "hf.co/menlo/jan-nano-gguf:q4_k_m"
  max_tokens: 1024
  temperature: 0.1

rag_config:
  num_examples: 1  # Only 1 example due to context limits
  use_template_based: true  # Provide explicit templates

prompt_strategy: "template_based"
```

---

## SECTION 6: BEST PRACTICES

### 1. Use lower temperature for better alignment
```python
# Bad: high temperature causes random spacing
model.temperature = 0.7  ❌

# Good: zero temperature for consistent spacing
model.temperature = 0.0  ✅
```

### 2. Provide character count examples in prompt
```
Example with character counts:

┌────────────────────────┐  (24 characters total)
│ Person                 │  (24 = 1 + 1 + 6 + 15 + 1)
├────────────────────────┤  (24 characters total)

Breakdown of second line:
- │ = 1 char
- space = 1 char
- "Person" = 6 chars
- padding spaces = 15 chars
- │ = 1 char
Total = 24 chars
```

### 3. Post-process with alignment fixer
```python
def generate_diagram(rag, llm, query):
    # Generate
    prompt = rag.build_prompt_with_alignment(query, SYSTEM_PROMPT)
    raw_output = llm.generate(prompt)

    # Extract diagram
    diagram = extract_code_block(raw_output)

    # Validate and fix
    is_valid, errors = rag.validate_alignment(diagram)
    if not is_valid:
        print(f"Alignment issues detected: {errors}")
        diagram = rag.fix_alignment(diagram)

    return diagram
```

### 4. Use monospace font verification
```python
def verify_monospace_rendering(diagram: str):
    """Check if diagram renders correctly in terminal"""
    print("\nRaw character lengths:")
    for i, line in enumerate(diagram.split('\n'), 1):
        print(f"Line {i}: {len(line)} chars")

    print("\nRendered diagram:")
    print(diagram)
```

---

## SECTION 7: TROUBLESHOOTING

### Problem 1: Lines have different lengths
**Cause**: LLM adds variable spacing
**Solution**: Use character count instructions + temperature = 0.0

### Problem 2: Boxes appear misaligned
**Cause**: Unicode characters have different widths in some terminals
**Solution**: Test in standard terminal with monospace font

### Problem 3: LLM changes border structure
**Cause**: Prompt doesn't emphasize structure preservation
**Solution**: Add explicit "COPY STRUCTURE EXACTLY" instruction

### Problem 4: Inheritance symbol missing
**Cause**: Example doesn't show inheritance
**Solution**: Ensure retrieved examples include inheritance diagrams (055)
