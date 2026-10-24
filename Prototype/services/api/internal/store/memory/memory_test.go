package memory_test

import (
	"testing"

	"study/api/internal/store/memory"
	"study/api/internal/store/storetest"
)

func TestConformance(t *testing.T) { storetest.Run(t, memory.New(), "mem") }
