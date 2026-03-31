/*
Dice Roller
30 Mar 2026
*/

package "main"

import (
	"fmt"
	"math/rand"
	"time"
	"flag"
)

func init() {
    rand.Seed(int64(time.Now().Nanosecond()))
}

type Dice str