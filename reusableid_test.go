package reusableid

import (
	"fmt"
	"testing"
)

func checkResult(t *testing.T, id1 uint64, id2 uint64) {
	if id1 != id2 {
		t.Fatal("Ids are not equal")
	}
}

func TestIdGenerator(t *testing.T) {
	// case 1
	{
		idGen := NewIdGenerator()
		id0 := idGen.Next()
		fmt.Printf("idGen.Next: %v\n", id0)
		checkResult(t, id0, 0)
		id1 := idGen.Next()
		fmt.Printf("idGen.Next: %v\n", id1)
		checkResult(t, id1, 1)

		idGen.Recycle(id1)
		id1Recycled := idGen.Next()
		fmt.Printf("id1Recycled: %v\n", id1Recycled)
		checkResult(t, id1Recycled, 1)
	}
}
