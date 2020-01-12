# ce

### usage

```
package ce

import (
	"io"
	"os"
	"testing"
)

func TestCheckError(t *testing.T) {
	CheckError(io.EOF)
}
```