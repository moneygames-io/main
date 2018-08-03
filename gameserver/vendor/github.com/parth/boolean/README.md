# Go Boolean to Integer conversions

Once in a while, I end up writing some logic that relies on being able to convert booleans to ints and vice versa. 

# Usage - boolean to integer:

```
package main

import (
	"fmt"
	"github.com/parth/boolean"
)

func main() {
	fmt.Println(boolean.BtoI(false)); // will print 0
}
```

# Usage - integer to boolean

```
package main

import (
	"fmt"
	"github.com/parth/boolean"
)

func main() {
	fmt.Println(boolean.ItoB(1)); // will print true
}
```
