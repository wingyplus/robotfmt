# Robot Framework Formatter

The command `robotfmt` print formatted code by given filename.
By default this command print the result to standard output. When
use option `-w` it write the result to filename.

## Install

```go
$ go get github.com/wingyplus/robotfmt
```

## Example

*Before*

```robot
# example.robot
*** Settings ***
Library  Selenium2Library
Library  CollectionLibrary
Test Template	Work Around


*** Test Cases ***	fieldA	fieldB
Case 1	1	2


*** Keywords ***
Work Around
     Capture Page Screenshot
```

*After*

```robot
# example.robot
*** Settings ***
Library          Selenium2Library
Library          CollectionLibrary
Test Template    Work Around


*** Test Cases ***    fieldA    fieldB
Case 1                1         2


*** Keywords ***
Work Around
    Capture Page Screenshot
```
