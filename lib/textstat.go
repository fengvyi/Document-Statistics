package lib

// TODO: Document functions.
// TODO: Add analysis of sentences.

import (
	"bufio"
	"io"
	"os"
	"strings"
      "sort"
      "strconv"  
      "fmt"
)

// Textstat ...
type Textstat struct {
	total     int
	histogram Histogram
      character Character
      length    int
}

// New ...
func New() Textstat {
	return Textstat{histogram: make(map[string]int),character:make(map[byte]int) }
}

// FromReader ...
func FromReader(reader io.Reader) (Textstat, error) {
	t := New()
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		t.add(scanner.Text())
	}
	return t, scanner.Err()
}

// FromFile ...
func FromFile(path string) (Textstat, error) {
	file, err := os.Open(path)
	if err != nil {
		return New(), err
	}
	return FromReader(file)
}

// Parse ...
func (t *Textstat) Parse(text string) {
	words := strings.Fields(text)
	for _, word := range words {
		t.add(word)
	}
}

// TotalWords ...
func (t Textstat) TotalWords() int {
	return t.total
}

// UniqueWords ...
func (t Textstat) UniqueWords() int {
	return len(t.histogram)
}

// AverageWordLength ...
func (t Textstat) AverageWordLength() float64 {
	return round2(divideInt(t.length, t.total))
}

// MostUsedWords ...
func (t Textstat) MostUsedWords() []string {
	max := minInt(len(t.histogram), 10)
	words := make([]string, max)
      temp :=t.histogram
	for i := 0; i < max; i++ {
		s,c := temp.RemoveMax()
      words[i] =s+":"+ strconv.Itoa(c)
	}
	return words
}

func (t Textstat) TheTopTen() {
     s := t.MostUsedWords()
     fmt.Printf("\n")
     fmt.Printf("\n")
     fmt.Printf("  %s","The Top Ten: ") 
     for i:=0;i<len(s);i++{
     fmt.Printf("%s, ",s[i])
}
}


func (t Textstat) ListOfWords() []string {
	words := make([]string, len(t.histogram))
      i:=0
      	for key,_ :=range t.histogram{
		words[i] = key
             i++
	}
      sort.Strings(words)
	return words
}

func (t Textstat) HistogramOfWords(){
      
      i:=0

     fmt.Printf("\n")
     fmt.Printf("  %s","Histogram Of Words:   ") 
     fmt.Printf("\n\t\t")

      for key,val := range t.histogram{
      i++
      fmt.Printf(" %s: %d, ", key,val)
      if i%6==0{
      fmt.Printf("\n")
      fmt.Print("\t\t")
}

}

} 


func (t Textstat) CharactersCount() {
      i:=0

     fmt.Printf("\n")
     fmt.Printf("  %s","Characters:   ") 

      for key,val:= range t.character{
           i++
           fmt.Printf(" %q:%d, ",key,val)
           if i%10==0{
      fmt.Printf("\n")
      fmt.Print("\t\t")
}
}
}



//
// Private:
//

// TODO: Remove special chars from words
func (t *Textstat) add(word string) {
	t.total++
	t.histogram[strings.ToLower(word)]++
       for i:=0;i<len([]rune(word));i++{
       t.character[strings.ToLower(word)[i]]++
}
	t.length += len([]rune(word))
}
