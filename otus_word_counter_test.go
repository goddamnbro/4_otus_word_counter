package otus_word_counter

import (
	"strings"
	"testing"
)

var samples = []map[string]string{
	{
		"input":    "a b b c c c d d d d",
		"expected": "d, c, b, a",
	},
	{
		"input":    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque eu leo a urna semper molestie non eget ex. Maecenas posuere eleifend risus, ut sagittis metus congue eu. Donec suscipit, lorem aliquet pretium vulputate, velit leo auctor risus, at finibus lorem purus et orci. Nullam quis enim facilisis, ultricies velit ut, volutpat risus. Mauris malesuada facilisis tempor. Aenean blandit sodales neque fringilla tempus. Praesent viverra porttitor consectetur. Fusce ac cursus turpis. Donec tempus accumsan velit nec tristique. In condimentum, diam et gravida volutpat, justo ante tincidunt dui, vel aliquam metus orci in metus. Curabitur pretium faucibus euismod.",
		"expected": "velit, risus, metus, lorem, volutpat, ut, tempus, pretium, orci, leo",
	},
}

func TestCompress(t *testing.T) {

	for _, sample := range samples {
		input := sample["input"]
		expected := sample["expected"]

		actual := strings.Join(getTop10FrequencyWords(input), ", ")
		if actual != expected {
			t.Errorf("Word counter works incorrect!\n "+
				"got:      %s \n "+
				"expected: %s",
				actual, expected)
		}
	}
}

