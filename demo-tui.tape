# File used to generate misc/demo-tui.gif, the demo of `golings watch`.
#
# Needs vhs (https://github.com/charmbracelet/vhs) and a golings built from this
# checkout on the PATH:
#   go install ./golings && vhs demo-tui.tape

Output misc/demo-tui.gif

Set Shell "bash"
Set FontSize 18
Set Width 1100
Set Height 640
Set Padding 15
Set Framerate 20
Set TypingSpeed 70ms

# Set up a fresh workspace, out of sight. A background job saves a fix for the
# first exercise after a while, the way an editor does: it writes a temporary
# file and renames it over the original.
Hide
Type "cd $(mktemp -d) && golings init > /dev/null && cd golings-exercises"
Enter
Type "(sleep 12; sed -i.bak 's/var = 5/var x = 5/' exercises/variables/variables1/main.go) & disown"
Enter
Type "clear"
Enter
Show

Type "golings watch"
Sleep 600ms
Enter
Sleep 4s

# The hint.
Type "h"
Sleep 3s
Escape
Sleep 1s

# The fix is saved by the background job: the view runs the exercise again.
Sleep 4s

# n takes out the marker of the exercise that works and moves on.
Type "n"
Sleep 3s

# l lists every exercise: filter by name or topic, enter opens one.
Type "l"
Sleep 2s
Type "/"
Sleep 400ms
Type "map"
Sleep 1s
Enter
Sleep 1s
Enter
Sleep 3s

# c checks all the exercises, without changing any file.
Type "c"
Sleep 13s
Escape
Sleep 1s

Type "q"
Sleep 1s
