# Trinethra

A Go-based web service that analyzes supervisor interview transcripts to evaluate DT Fellows — early-career professionals on short-term client engagements. It uses a local Ollama LLM to score performance across execution quality, systems-building impact, KPI contribution, and change management, then returns structured JSON analysis.

---

## Prerequisites

- [Go](https://golang.org/dl/) 1.21+
- [Ollama](https://ollama.com) running locally with `llama3.2` pulled

---

## Setup

**1. Clone the repository**

```bash
git clone https://github.com/your-username/trinethra.git
cd trinethra
```

**2. Install dependencies**

```bash
go mod tidy
```

**3. Start Ollama and pull the model**

```bash
ollama serve
ollama pull llama3.2
```

Ollama must be running at `http://localhost:11434` before starting the server.

---

## Running the Server

```bash
go run ./cmd/api/main.go
```

Server starts on `http://localhost:8080`.

The static frontend is served at `/` and the analysis endpoint is at `/analyze`.

---

## API

**POST** `/analyze`

Request body:

```json
{
  "transcript": "Paste the supervisor interview transcript here."
}
```

Response:

```json
{
  "score": {
    "value": 7,
    "label": "Strong Performer",
    "band": "6-8",
    "justification": "...",
    "confidence": "High"
  },
  "evidence": [...],
  "kpi": [...],
  "gap": [...],
  "questions": [...]
}
```

---

## Project Structure

```
trinethra/
├── cmd/
│   └── api/
│       └── main.go           # Entry point
├── internal/
│   ├── handler/
│   │   └── analyze.go        # HTTP handler for /analyze
│   ├── ollama/
│   │   └── client.go         # Ollama API client
│   ├── prompt/
│   │   └── prompt.go         # Prompt construction logic
│   └── response/
│       └── response.go       # LLM response parser
├── package/
│   └── models/
│       └── models.go         # Shared data models
└── static/
    └── index.html            # Frontend UI
```
