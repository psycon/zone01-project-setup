# zone01-project-setup

A small CLI tool for Zone01 students that automatically fetches project and audit instructions from the [01-edu public repository](https://github.com/01-edu/public) and populates the local template files, ready to use with an LLM.

---

## What it does

1. Asks for the **project name** → fetches the `README.md` from `subjects/<project>/` and writes it to `Project Instructions.md`
2. Asks whether it is a **sub-project** (e.g. `filters` inside `groupie-tracker`) → adjusts the path accordingly
3. Fetches `subjects/<project>/audit/README.md` and writes it to `Audit Instructions.md`

---

## <span style="color: #CCFF99;">How it is supposed to be used by the students</span>

This program should just save you time from creating the final "Audit Instructions.md" and "Project Instructions.md" files by hand. Then go to your favourite LLM chat copy/paste the Master Prompt (context) and attach the 4 files ("Project Instructions.md, Audit Instructions.md, Good Practises.pdf, Unit Testing Instructions.pdf").

The preferred AI should be web-based rather than IDE-based or agent-based, because in the first phase you don't need to write code in the IDE (VS Code) — you need the advantages that web-based chat offers, such as complex reasoning, architecture discussions, large context windows, long chat history, and raw capabilities, in order to have a back-and-forth conversation with the LLM before it asks you to write the code/pseudocode for every gate (see the Master Prompt for context).

If you use that Master Prompt, you will NOT finish any project in 1 hour, not even in 1 day, but you will gain significant benefits that aims to:

1. Escaping the "Junior Trap" (Architecture vs. Typing)
The most common pitfall for junior developers is the urge to immediately start writing syntax the moment they read a prompt. This "execute first, think later" approach inevitably leads to brittle, spaghetti code and endless, frustrating debugging cycles. This prompt acts as a mandatory circuit breaker. It forces the AI to halt your typing momentum and ask fundamental questions like, "Before we even write the main() function, how are we managing the memory allocation and lifecycle for this slice?" It completely rewires your brain, training you to think systematically like a Software Architect rather than just acting as a typist.

2. Forging Enterprise-Grade Engineering Habits
Concepts like maintaining an ai_changelog.md, enforcing "Quality Gates," and doing Spec-Driven Development aren't just theoretical exercises—they are the exact operational standards utilized by top-tier tech giants (like Google, Meta, and Amazon). By integrating these practices into your daily workflow right now, you are doing much more than just solving a school assignment. You are learning how to collaborate within massive codebases, document architectural context, and build auditable, maintainable software. You are essentially building senior-level professional habits from day one, making you highly employable and team-ready.

3. Mastering Golang "Under the Hood"
Go is a deceptively simple language that hides immense power—and potential pitfalls—in its concurrency models and memory management. Because this prompt strictly mandates that the AI explains mechanics at the OS and RAM level, you won't just learn how to use a Goroutine, a pointer, or an interface. You will understand exactly what the Go compiler, the operating system, and the garbage collector are doing behind the scenes. This deep-dive approach will elevate you from someone who just "makes Go code compile" to a master of performance optimization, memory safety, and concurrent programming.

4. Turning Restrictions into a Superpower (The Ultimate Synergy)
At first glance, strict Standard Operating Procedures (SOPs) can feel like annoying bureaucratic constraints that slow down your coding. However, this prompt reframes those rules entirely. Instead of acting as a rigid wall, the SOP becomes the very scaffolding your AI Mentor uses to safely guide your education. These guardrails actually reduce your cognitive load; they free your mind to focus on pure engineering logic and problem-solving, knowing that your Mentor will catch structural flaws and edge cases before they cascade into unfixable technical debt.

---

## Requirements

- [Go 1.21+](https://go.dev/dl/)

---

## Installation

```bash
git clone https://github.com/<your-username>/zone01-project-setup.git
cd zone01-project-setup
go build -o cmd/setup cmd/setup.go
```

---

## Usage

Always run the tool from the **repository root**:

```bash
# using the compiled binary
./cmd/setup

# or without compiling
go run cmd/setup.go
```

---

## Notes

- Project names must match exactly the folder names used in the [01-edu subjects directory](https://github.com/01-edu/public/tree/master/subjects).
- `Master_Prompt.md` contains the LLM communication guidelines and should **not** be modified.
- `Project Instructions.md` and `Audit Instructions.md` are overwritten every time you run the tool.

---

Copy the files to your project folder:
cp "Project Instructions.md" "Master_Prompt.md" "Unit Testing Instructions.pdf" "Good Practises.pdf" "Audit Instructions.md" "LICENSE" /Users/con/zone01/push-swap/

or

mkdir -p /Users/con/zone01/push-swap && cp "Project Instructions.md" "Master_Prompt.md" "Unit Testing Instructions.pdf" "Good Practises.pdf" "Audit Instructions.md" "LICENSE" /Users/con/zone01/push-swap/

---

*A project from Constantine E. Papadopoulos*
