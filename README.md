# charsetsupport
packages to support charset

## Shift JIS writer
### Background
Some letters are not available with Shift JIS even if they're available with UTF-8
It usually causes issues to convert words including unavailables letters to Shift JIS.

### Solution
If there is a replacable letter, it's converted to it.
If not, it's converted to "?".
