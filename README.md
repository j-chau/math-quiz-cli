# math-quiz-cli

A math quiz CLI that poses math questions to the user with a time limit. 

## Usage

```bash
math -csv="problems.csv" -limit=30
```
- `problems.csv` is the list of questions and answers to ask the user; should be written in `question,answer` format
- limit is the time limit in seconds (default 30s)
