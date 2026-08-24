package cli_test

import (
	"os"
	"testing"

	"github.com/gftcode/function-square-api/internal/infrastructure/entrypoint/cli"
)

func TestOutput_CaptureAndStorage(t *testing.T) {
	input := "2\n-4\n6\ntrue\n"
	
	defer mockStdin(t, input)()

	o := cli.NewCli()
	_ = o.SetValues() 

	a, b, c := o.GetValues()
	wantA, wantB, wantC := 2, -4, 6

	if a != wantA || b != wantB || c != wantC {
		t.Errorf("SetValues() = [%d, %d, %d]; esperado GetValues() = [%d, %d, %d]", a, b, c, wantA, wantB, wantC)
	}
}

func mockStdin(t *testing.T, content string) func() {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("falha ao criar pipe: %v", err)
	}

	oldStdin := os.Stdin
	os.Stdin = r

	go func() {
		defer w.Close()
		_, _ = w.WriteString(content)
	}()

	return func() {
		os.Stdin = oldStdin
		r.Close()
	}
}
