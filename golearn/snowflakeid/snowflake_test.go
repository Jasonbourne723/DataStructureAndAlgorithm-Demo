package snowflakeid

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSnowflakeId(t *testing.T) {
	s := &SnowFlaker{}

	id := s.Next()

	fmt.Println(id)

	binaryID := strconv.FormatInt(id, 2)

	fmt.Println(binaryID)
	fmt.Println(strconv.FormatInt(607937840337428488, 2))
}
