package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Benchmark

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("insun")
	}
}

func BenchmarkHelloWorldmaster(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("Master Good Taichi")
	}
}

// Sub Benchmark
func BenchmarkSubHello(b *testing.B) {
	b.Run("insun", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("insun")
		}
	})

	b.Run("master", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("master")
		}
	})
}

// Table Bench Mark
func BenchmarkTableHello(b *testing.B) {
	benchmark := []struct {
		name    string
		request string
	}{
		{
			name:    "insun",
			request: "insun",
		},
		{
			name:    "inzun",
			request: "inzun",
		},
		{
			name:    "insan",
			request: "insan",
		},
	}
	for _, benchmark := range benchmark {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SayHello(benchmark.request)
			}
		})
	}
}

// Table Test
func TestHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "insun",
			request:  "insun",
			expected: "Hello insun",
		},
		{
			name:     "inzun",
			request:  "inzun",
			expected: "Hello inzun",
		},
		{
			name:     "master",
			request:  "master",
			expected: "Hello master",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SayHello((test.request))
			assert.Equal(t, test.expected, result)
		})
	}
}

// Sub Test
func TestSubTest(t *testing.T) {
	t.Run("insun", func(t *testing.T) {
		result := SayHello(("insun"))
		assert.Equal(t, "Hello insun", result, "Text Must Be Hello insun")
	})

	t.Run("inzun", func(t *testing.T) {
		result := SayHello(("inzun"))
		assert.Equal(t, "Hello inzun", result, "Text Must Be Hello inzun")
	})
}

// Before After Test
func TestMain(m *testing.M) {
	fmt.Println("Before Unit Tes")

	m.Run()

	fmt.Println("After Unit Tes")
}

// Assert Test
func TestHelloCrypto(t *testing.T) {
	result := SayHello(("Crypto"))

	assert.Equal(t, "Hello Crypto", result, "Text Must Be Hello Crypto")
	fmt.Println("Testing Done")
}

// Require Test
func TestHelloMaster(t *testing.T) {
	result := SayHello(("Master"))

	require.Equal(t, "Hello Master", result, "Text Must Be Hello Master")
	fmt.Println("Testing Done")
}

func TestHelloInsun(t *testing.T) {
	result := SayHello("insun")
	// fmt.Println(result)

	if result != "Hello insun" {
		// panic("Result is not Hello insun")
		// t.Fail()
		t.Error("Text Not Same")
	}
	fmt.Println("Tes Has Been Done")
}

func TestHelloInzun(t *testing.T) {
	result := SayHello("inzun")
	if result != "Hello inzun" {
		// panic("Result is not Hello insun") disarankan menggunakan error/ fatal
		// t.FailNow()
		t.Fatal("Text Not Same")
	}
	fmt.Println("Tes Has Been Done")
}

// Os Check
func TestSkipTes(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Your OS Can't Run")
	}
	fmt.Println("Tes Has Been Done")
}
