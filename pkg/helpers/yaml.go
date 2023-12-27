package helper

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	_ "github.com/joho/godotenv/autoload"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
	ctyconvert "github.com/zclconf/go-cty/cty/convert"
)

var (
	awsBucketName = os.Getenv("AWS_BUCKET_NAME")
	pathTempFile  = os.Getenv("PATH_TEMP_FILE")
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateYaml(c *gin.Context, filename string, vars map[string]interface{}) error {

	content, err := ioutil.ReadFile(pathTempFile + filename)
	if err != nil {
		log.Fatal(err)
	}

	// before splitting the document, parse out any template details
	rendered := string(content)
	log.Println("rendered: " + rendered)

	rendered, err = parseTemplate(rendered, vars)
	check(err)

	f, err := os.Create(pathTempFile + "rendered_" + filename)
	check(err)
	defer f.Close()

	n3, err := f.WriteString(rendered)
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	return err
}

// execute parses and executes a template using vars.
func parseTemplate(s string, vars map[string]interface{}) (string, error) {
	expr, diags := hclsyntax.ParseTemplate([]byte(s), "<template_file>", hcl.Pos{Line: 1, Column: 1})
	if expr == nil || (diags != nil && diags.HasErrors()) {
		return "", diags
	}

	ctx := &hcl.EvalContext{
		Variables: map[string]cty.Value{},
	}

	for k, v := range vars {
		// In practice today this is always a string due to limitations of
		// the schema system. In future we'd like to support other types here.
		s, ok := v.(string)
		if !ok {
			return "", fmt.Errorf("unexpected type for variable %q: %T", k, v)
		}
		ctx.Variables[k] = cty.StringVal(s)
	}

	result, diags := expr.Value(ctx)
	if diags != nil && diags.HasErrors() {
		return "", diags
	}

	// Our result must always be a string, so we'll try to convert it.
	var err error
	result, err = ctyconvert.Convert(result, cty.String)
	if err != nil {
		return "", fmt.Errorf("invalid template result: %s", err)
	}
	log.Println("result: " + result.AsString())

	return result.AsString(), nil

}
