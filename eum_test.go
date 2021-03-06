package instana_test

import (
	"testing"

	"github.com/instana/golang-sensor"
	"github.com/stretchr/testify/assert"
)

const eumExpectedResult string = `<script>
  (function(i,s,o,g,r,a,m){i['InstanaEumObject']=r;i[r]=i[r]||function(){
  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
  })(window,document,'script','//eum.instana.io/eum.min.js','ineum');

  ineum('apiKey', 'myApiKey');
  ineum('traceId', 'myTraceId');
  ineum('meta', 'key1', 'value1');
  ineum('meta', 'key2', 'value2');

</script>`

func TestEum(t *testing.T) {

	m := make(map[string]string, 2)
	m["key1"] = "value1"
	m["key2"] = "value2"

	var result string = instana.EumSnippet("myApiKey", "myTraceId", m)
	assert.Equal(t, result, eumExpectedResult, "They should be equal")
}
