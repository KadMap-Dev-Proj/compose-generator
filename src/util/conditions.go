package util

import (
	"compose-generator/model"
	"encoding/json"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

func EvaluateConditionalSections(
	filePath string,
	templateData map[string][]model.ServiceTemplateConfig,
	varMap map[string]string,
) string {
	dataString := PrepareInputData(templateData, varMap)
	// Execute CCom
	return ExecuteAndWaitWithOutput("ccom", "-d", string(dataString), "-s", filePath)
}

// EvaluateCondition evaluates the given condition to a boolean result
func EvaluateCondition(
	condition string,
	templateData map[string][]model.ServiceTemplateConfig,
	varMap map[string]string,
) bool {
	dataString := PrepareInputData(templateData, varMap)
	// Execute CCom
	result := ExecuteAndWaitWithOutput("ccom", "-m", "-d", string(dataString), "-s", condition)
	return result == "true"
}

func CheckIfCComIsInstalled() {
	if !CommandExists("ccom") {
		Error("CCom is not missing on your system. Please go to https://github.com/compose-generator/compose-generator/releases/latest to download the latest version.", nil, true)
	}
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------

func PrepareInputData(
	templateData map[string][]model.ServiceTemplateConfig,
	varMap map[string]string,
) string {
	data := model.CComDataInput{
		Services: templateData,
		Var:      varMap,
	}
	dataJson, err := json.Marshal(data)
	if err != nil {
		Error("Could not evaluate conditional sections in template. Could be corrupted", err, true)
	}
	return string(dataJson)
}
