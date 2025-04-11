package game

import "testing"

func TestLoadItems(t *testing.T) {
	items, err := LoadItems()
	if err != nil {
		t.Fatalf("failed to load items: %v", err)
	}
	if len(items) == 0 {
		t.Errorf("no items loaded")
	}
	if _, ok := items["healing_potion"]; !ok {
		t.Errorf("expected item 'healing_potion' in item list")
	}
}

func TestLoadBackgrounds(t *testing.T) {
	bgs, err := LoadBackgrounds()
	if err != nil {
		t.Fatalf("failed to load backgrounds: %v", err)
	}
	if len(bgs) == 0 {
		t.Errorf("no backgrounds loaded")
	}
	if len(bgs[0].StartingItemIDs) == 0 {
		t.Errorf("expected to have starting items")
	}
}

func TestLoadEnemies(t *testing.T) {
	enemies, err := LoadEnemies()
	if err != nil {
		t.Fatalf("failed to load enemies: %v", err)
	}
	if len(enemies) == 0 {
		t.Errorf("no enemies loaded")
	}
}
