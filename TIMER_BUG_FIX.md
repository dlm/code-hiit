# Timer Bug Fix - td-00a44e

## Problem
Timer countdown was not advancing during HIIT workout phases. Timer appeared frozen until user typed characters, preventing automatic phase transitions.

## Root Cause
The `transitionPhase()` function at line 270 was returning `nil` as the command:
```go
return m, nil
```

This stopped the Bubble Tea tick loop, preventing further timer updates until a keyboard event occurred.

## The Bug Chain
This single bug caused **three** issues:
1. **td-00a44e**: Timer doesn't advance - tick loop stopped after first transition
2. **td-707ec2**: Phase doesn't auto-transition - timer not running to reach 0:00
3. **td-f0d990**: Work → Recovery fails - same root cause

## Solution
Changed `transitionPhase()` to return `tickCmd()` to continue the timer:
```go
// Continue the timer tick loop
return m, tickCmd()
```

**File:** `main.go` line 270-271

## How It Works

### Bubble Tea Event Loop
1. `tickCmd()` schedules a message every 100ms
2. `Update()` receives `tickMsg`
3. Checks if timer expired (`RemainingTime() <= 0`)
4. If expired, calls `transitionPhase()`
5. **Critical:** Must return `tickCmd()` to continue loop

### Before Fix
```
Init() → tickCmd() → tickMsg → Update()
  → transitionPhase() → return nil ❌
  → Timer stops, no more ticks
```

### After Fix
```
Init() → tickCmd() → tickMsg → Update()
  → transitionPhase() → return tickCmd() ✅
  → Timer continues → more ticks → smooth countdown
```

## Testing

### Manual Verification
```bash
./code-hiit --demo-hiit
```

**Expected behavior:**
- ✅ Timer counts down continuously (every 100ms)
- ✅ Display updates without typing
- ✅ Phase auto-transitions at 0:00
- ✅ Warmup (15s) → Work (20s) → Recovery (10s) → Next set
- ✅ All 3 sets complete automatically if no typing

### Test Cases

**1. Timer Advances Without Typing**
- Start workout
- Don't type anything
- Verify: Timer counts down from 00:15 → 00:00
- Expected: Smooth countdown, transitions at 0:00

**2. Phase Auto-Transition**
- Let warmup timer expire
- Expected: Automatically transitions to work phase
- Verify: New snippet loads, timer resets to 00:20

**3. Work → Recovery Transition**
- Complete warmup
- Let work phase timer expire
- Expected: Transitions to recovery (quote snippet)
- Verify: Timer shows 00:10

**4. Complete Set → Next Set**
- Let all 3 phases expire in set 1
- Expected: Transitions to Set 2, Warmup phase
- Verify: "Set 2/3" displays

**5. Complete Workout**
- Let all 3 sets complete
- Expected: Shows workout summary screen
- Verify: "WORKOUT COMPLETE!" displays

## Impact

### Bugs Fixed
- ✅ **td-00a44e**: Timer advances continuously
- ✅ **td-707ec2**: Phases auto-transition at expiry
- ✅ **td-f0d990**: Work → Recovery works correctly

### Side Effects
- None - fix only affects timer loop continuity
- Typing still works normally
- Pause/resume unaffected
- Manual phase completion (typing snippet) still works

## Files Modified
- `main.go`: Line 270-271 (1 line change)

## Related Code

### Where Timer Starts
- Line 137: `Init()` starts timer for HIIT mode
- Line 473: `startHIITWorkout()` returns `tickCmd()`

### Where Timer Continues
- Line 166: `Update()` returns `tickCmd()` for each tick
- Line 271: **Fixed** `transitionPhase()` returns `tickCmd()`

### Timer Loop Flow
```
startHIITWorkout()
  → sets isHIITMode = true
  → returns tickCmd()
    → 100ms passes
    → tickMsg received
    → Update() processes tick
      → checks RemainingTime()
      → returns tickCmd() to continue
        → loop repeats ♻️
```

## Prevention
To avoid similar bugs:
1. Any function that modifies workout state during active timer must return `tickCmd()`
2. Only return `nil` when intentionally stopping the timer (e.g., workout complete, quit)
3. Pattern: `return m, tickCmd()` for phase transitions

## Verification Checklist
- [x] Build succeeds
- [x] Timer counts down without user input
- [x] Phase transitions happen automatically
- [x] Work → Recovery transition works
- [x] Recovery → Warmup (next set) works
- [x] Complete workout shows summary
- [x] Pause/resume still functional
- [x] Typing during countdown works

## Notes
This was a **critical** bug that broke the core HIIT experience. The fix is minimal (1 line) but fixes three reported issues. Ready for user testing.
