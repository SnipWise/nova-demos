# RAG-Powered ASCII Diagram Generator

A Go application using the Nova SDK that generates ASCII diagrams using RAG (Retrieval-Augmented Generation) with automatic alignment validation and correction.

## Overview

This project demonstrates how to combine RAG and chat agents to generate properly formatted ASCII diagrams. The system:
1. Searches a vector store for similar diagram examples
2. Uses retrieved examples as reference to generate new diagrams
3. Validates alignment of the generated output
4. Auto-fixes any alignment issues

## Project Structure

```
06-rag-diagrams/
├── main.go                      # Main application entry point
├── go.mod                       # Go module definition
├── go.sum                       # Go dependencies checksums
├── config/
│   └── config.go                # Configuration and system instructions
├── diagram/
│   ├── extractor.go             # Extracts diagrams from markdown code blocks
│   └── fixer.go                 # Auto-fixes alignment issues in diagrams
├── validator/
│   └── validator.go             # Validates diagram alignment
├── data/
│   └── diagram_library_en.md    # Source diagrams for RAG (327 lines)
├── store/
│   └── diagrams.json            # Persistent vector store
└── files/
    ├── rag_prompts_guide.md     # RAG prompting guide (French)
    └── rag_prompts_guide_en.md  # RAG prompting guide (English)
```

## How It Works

### 1. RAG Agent Initialization

The application uses a RAG agent with the Nova SDK to:
- Load or create a vector store from markdown files in [data/](data/)
- Chunk markdown content by sections using `chunks.SplitMarkdownBySections()`
- Generate embeddings using the configured embedding model
- Persist the vector store to [store/diagrams.json](store/diagrams.json)

**Configuration** ([config/config.go:14-23](config/config.go#L14-L23)):
- **RAG Model**: `huggingface.co/mixedbread-ai/mxbai-embed-large-v1:f16` (default)
- **Store Path**: `./store/diagrams.json`
- **Data Source**: `./data` directory

### 2. Similarity Search

When a diagram request is made, the RAG agent ([main.go:134-150](main.go#L134-L150)):
- Searches for similar diagrams using semantic similarity (threshold: 0.5)
- Returns relevant examples with their similarity scores
- Builds a context string with the most similar diagrams

### 3. Chat Agent Generation

A chat agent generates the new diagram ([main.go:153-167](main.go#L153-L167)):
- **Model**: `hf.co/menlo/jan-nano-gguf:q4_k_m` (default)
- **Temperature**: 0.0 (for consistent alignment)
- **System Instructions**: Defined in [config/config.go:26-82](config/config.go#L26-L82)
- **Streaming**: Enabled for real-time output

**Key prompting strategy** ([main.go:170-246](main.go#L170-L246)):
1. Provides similar diagrams as reference examples
2. Instructs the model to copy the structure identically
3. Only allows modification of text labels (not borders/structure)
4. Enforces character count validation before output

### 4. Validation and Auto-Fix

After generation ([main.go:258-303](main.go#L258-L303)):

**Extraction** ([diagram/extractor.go:6-22](diagram/extractor.go#L6-L22)):
- Extracts diagram from markdown code blocks

**Validation** ([validator/validator.go:14-53](validator/validator.go#L14-L53)):
- Checks if all lines in each box have the same visual character count (using runes)
- Identifies alignment issues by box number

**Auto-Fix** ([diagram/fixer.go:15-165](diagram/fixer.go#L15-L165)):
- Detects box structures (lines starting with `┌`)
- Calculates target width based on longest content
- Fixes border lines to match target width
- Pads content lines with spaces to achieve uniform length
- Handles UTF-8 characters correctly using rune counting

## Configuration

The application uses environment variables or defaults ([config/config.go:15-23](config/config.go#L15-L23)):

```bash
# Engine configuration
ENGINE_URL="http://localhost:12434/engines/llama.cpp/v1"

# Models
CHAT_MODEL="hf.co/menlo/jan-nano-gguf:q4_k_m"
RAG_MODEL="huggingface.co/mixedbread-ai/mxbai-embed-large-v1:f16"
```

## ASCII Box Characters

The system supports the following characters ([config/config.go:31-35](config/config.go#L31-L35)):

- **Boxes**: `┌ ┐ └ ┘ ─ │ ├ ┤ ┬ ┴ ┼`
- **Arrows**: `▶ ◀ ▲ ▼ → ← ↑ ↓ ──▶ ◀──`
- **Decisions**: `╱ ╲`
- **Jointures**: `═ ╧ ╤ ╪`
- **Inheritance**: `△` (triangle pointing up)

## Alignment Rules

Every line in a box MUST have exactly the same character count ([config/config.go:39-58](config/config.go#L39-L58)).

**Example** (all lines = 24 characters):
```
┌────────────────────────┐  (24 chars)
│ Person                 │  (24 chars: │ + space + "Person" + 15 spaces + │)
├────────────────────────┤  (24 chars)
│ - name: string         │  (24 chars)
└────────────────────────┘  (24 chars)
```

## Running the Application

1. **Install dependencies**:
   ```bash
   go mod download
   ```

2. **Run the application**:
   ```bash
   go run main.go
   ```

3. **First run**: Creates the vector store from [data/diagram_library_en.md](data/diagram_library_en.md)

4. **Subsequent runs**: Loads the existing vector store from [store/diagrams.json](store/diagrams.json)

## Example Query

The default query in [main.go:47-51](main.go#L47-L51):

```go
query := `
Create a class diagram of a warrior that inherits from a character class.
The warrior has attributes strength and weapon, and methods attack() and defend().
The character class has attributes name and level, and method move().
`
```

## Output Flow

1. **RAG Search**: Displays similar diagrams with similarity scores
2. **Generation**: Streams the generated diagram in real-time
3. **Raw Analysis**: Shows character count for each line
4. **Validation**: Reports alignment issues (if any)
5. **Auto-Fix**: Corrects alignment problems
6. **Final Output**: Displays the corrected diagram

## Key Features

- **RAG-Powered**: Uses semantic search to find similar diagram examples
- **Streaming Generation**: Real-time output using Nova SDK streaming
- **Automatic Validation**: Checks alignment using UTF-8 aware rune counting
- **Auto-Correction**: Fixes alignment issues automatically
- **Persistent Store**: Vector embeddings saved to JSON for fast reloading
- **Configurable Models**: Supports different LLM and embedding models

## Dependencies

- **Nova SDK**: `github.com/snipwise/nova` (v1.0.6)
- **Go**: 1.25.4

## Notes

- The project is marked as work in progress ([this.is.a.work.in.progress.todo](this.is.a.work.in.progress.todo))
- Logging is enabled at INFO level ([main.go:28-30](main.go#L28-L30))
- Conversation history is disabled for diagram generation ([main.go:160](main.go#L160))
- Temperature is set to 0.0 for consistent alignment ([main.go:164](main.go#L164))
