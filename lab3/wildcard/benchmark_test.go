package wildcard

import (
	"strings"
	"testing"
)

var sourceString = `Lorem ipsum dolor sit amet,
consectetuer adipiscing elit. Aenean commodo 
ligula eget dolor. Aenean massa. Cum sociis 
natoque penatibus et magnis dis parturient montes, 
nascetur ridiculus mus. Donec quam felis, 
ultricies nec, pellentesque eu, pretium quis, sem. 
Nulla consequat massa quis enim. Donec pede justo, 
fringilla vel, aliquet nec, vulputate eget, arcu. 
In enim justo, rhoncus ut, imperdiet a, venenatis vitae, 
justo. Nullam dictum felis eu pede mollis pretium. 
Integer tincidunt. Cras dapibus. Vivamus elementum 
semper nisi. Aenean vulputate eleifend tellus. Aenean 
leo ligula, porttitor eu, consequat vitae, eleifend ac, 
enim. Aliquam lorem ante, dapibus in, viverra quis, 
feugiat a, tellus. Phasellus viverra nulla ut metus varius 
laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies 
nisi vel augue. Curabitur ullamcorper ultricies nisi. 
Nam eget dui. Etiam rhoncus. Maecenas tempus, 
tellus eget condimentum rhoncus, sem quam semper libero, 
sit amet adipiscing sem neque sed ipsum. Nam quam nunc, 
blandit vel, luctus pulvinar, hendrerit id, lorem. 
Maecenas nec odio et ante tincidunt tempus. Donec 
vitae sapien ut libero venenatis faucibus. Nullam quis 
ante. Etiam sit amet orci eget eros faucibus tincidunt. 
Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales 
sagittis magna. Sed consequat, leo eget bibendum sodales, 
augue velit cursus nunc.`

var simplePattern = `ma*sa`
var frequentPattern = `a*a`
var absentPattern = `abrakad*abra`
var frequentLetter = `a`

func benchmarkMatch(b *testing.B, pattern string, srcMultiplier int) {
	strA := []string{sourceString}
	for i := 0; i < srcMultiplier; i++ {
		strA = append(strA, sourceString)
	}
	str := strings.Join(strA, "")
	for n := 0; n < b.N; n++ {
		// run the Match function b.N times
		Match(sourceString, str)
	}
}

func BenchmarkSingle(b *testing.B)            { benchmarkMatch(b, simplePattern, 0) }
func BenchmarkFrequent(b *testing.B)          { benchmarkMatch(b, frequentPattern, 0) }
func BenchmarkAbsent(b *testing.B)            { benchmarkMatch(b, absentPattern, 0) }
func BenchmarkFrequentLetter(b *testing.B)    { benchmarkMatch(b, frequentLetter, 0) }
func BenchmarkSingle100(b *testing.B)         { benchmarkMatch(b, simplePattern, 100) }
func BenchmarkFrequent100(b *testing.B)       { benchmarkMatch(b, frequentPattern, 100) }
func BenchmarkAbsent100(b *testing.B)         { benchmarkMatch(b, absentPattern, 100) }
func BenchmarkFrequentLetter100(b *testing.B) { benchmarkMatch(b, frequentLetter, 100) }
