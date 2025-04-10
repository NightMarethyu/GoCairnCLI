# Cairn Console RPG – Development Roadmap

A portfolio-focused, feature-complete text-based RPG inspired by _Cairn 2e_.

---

## 📦 Phase 1: Foundation & Setup

- [x] Initialize Git repo and run `go mod init`
- [ ] Create basic file structure
- [x] Write `README.md` with Cairn license + project goals
- [ ] Create basic `main.go` with placeholder input loop
- [ ] Add `backgrounds.json` to `/data`
- [ ] Write basic tests for loading backgrounds
- [ ] Implement `utils/dice.go` for dice rolling

---

## 🧙 Phase 2: Character Creation

- [ ] Define `Character`, `Traits`, `Inventory` structs
- [ ] Load backgrounds and randomly select one
- [ ] Prompt user to roll or choose background + traits
- [ ] Roll attributes (STR, DEX, WIL) and HP
- [ ] Roll Bonds and Omens
- [ ] Store and display character info
- [ ] Add unit tests for character generator

---

## 🌍 Phase 3: World & Region Generation

- [ ] Define `Region`, `Faction`, and `POI` structs
- [ ] Generate terrain, weather, and paths
- [ ] Add seed-based world generation
- [ ] Populate map with POIs and Factions
- [ ] Display region and POI info in CLI
- [ ] Write unit tests for worldgen

---

## ⚔️ Phase 4: Exploration & Loop

- [ ] Create main game loop and command parser
- [ ] Allow movement between POIs
- [ ] Implement time/watch tracking
- [ ] Add event/encounter randomizer
- [ ] CLI shows options, feedback, narration
- [ ] Add test coverage for loop actions

---

## 💥 Phase 5: Combat

- [ ] Define `Enemy` and `Combat` structs
- [ ] Add turn-based combat resolution
- [ ] Track HP, damage, death, scars
- [ ] Implement basic enemy generator
- [ ] Add combat test coverage

---

## 🏰 Phase 6: Dungeons & Events

- [ ] Generate and explore dungeon POIs
- [ ] Implement basic puzzle, trap, or secret system
- [ ] Include NPCs or factions within dungeon
- [ ] Write tests for dungeon traversal

---

## 🧽 Phase 7: Polish & Presentation

- [ ] Implement save/load system (JSON)
- [ ] Add ASCII/logo, spacing, or colors (optional)
- [ ] Update README with demo, instructions, license
- [ ] Record and upload short gameplay video/GIF
- [ ] Perform final test run and bug sweep

---

## 🧪 Testing Focus

- [ ] Dice rolling and edge cases
- [ ] Data loading and JSON format
- [ ] Player actions (travel, combat, use item)
- [ ] World and POI generation
- [ ] CLI responses and error handling
