package immuta

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"math/big"
	"reflect"
)

func stringResourceId() schema.StringAttribute {
	return schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Terraform resource identifier",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
}

func numberResourceId() schema.NumberAttribute {
	return schema.NumberAttribute{
		Computed:            true,
		MarkdownDescription: "Terraform resource identifier",
		PlanModifiers: []planmodifier.Number{
			numberplanmodifier.UseStateForUnknown(),
		},
	}
}

func intToNumberValue(i int) types.Number {
	return types.NumberValue(big.NewFloat(float64(i)))
}

func goMapFromTf(ctx context.Context, m types.Map) (map[string]interface{}, diag.Diagnostics) {
	goObject := make(map[string]interface{})
	// read from the terraform data into the map
	if diags := m.ElementsAs(ctx, &goObject, false); diags != nil {
		return nil, diags
	}
	return goObject, nil
}

func tfMapFromGo(ctx context.Context, m map[string]interface{}) (types.Map, diag.Diagnostics) {
	mappedValue, diags := types.MapValueFrom(ctx, types.StringType, m)
	if diags != nil {
		return types.MapNull(nil), diags
	}
	return mappedValue, nil
}

func updateMapIfChanged(ctx context.Context, tfMap types.Map, comparisonMap map[string]interface{}) (types.Map, diag.Diagnostics) {

	goTfMap, diags := goMapFromTf(ctx, tfMap)
	if diags != nil {
		return types.MapNull(nil), diags
	}

	// compare the map to the response from the API and if it's changed update the data object
	if !reflect.DeepEqual(goTfMap, comparisonMap) {
		tfFromComparison, err := tfMapFromGo(ctx, comparisonMap)
		if err != nil {
			return types.MapNull(nil), err
		}
		return tfFromComparison, nil
	}
	//return types.MapNull(nil), nil
	return tfMap, nil
}

func goListFromTf(ctx context.Context, l types.List) ([]interface{}, diag.Diagnostics) {
	goObject := make([]interface{}, 0)
	// read from the terraform data into the map
	if diags := l.ElementsAs(ctx, &goObject, false); diags != nil {
		return nil, diags
	}
	return goObject, nil
}

func tfListFromGo(ctx context.Context, l []interface{}) (types.List, diag.Diagnostics) {
	mappedValue, diags := types.ListValueFrom(ctx, types.StringType, l)
	if diags != nil {
		return types.ListNull(nil), diags
	}
	return mappedValue, nil
}

func updateListIfChanged(ctx context.Context, tfList types.List, comparisonList []interface{}) (types.List, diag.Diagnostics) {

	goTfList, diags := goListFromTf(ctx, tfList)
	if diags != nil {
		return types.ListNull(nil), diags
	}

	// compare two lists to see if they are equal

	listsAreSame := true
	if len(goTfList) != len(comparisonList) {
		listsAreSame = false
	}
	for i := range goTfList {
		if goTfList[i] != comparisonList[i] {
			listsAreSame = false
		}
	}

	if !listsAreSame {
		tfFromComparison, comparisonDiags := tfListFromGo(ctx, comparisonList)
		if comparisonDiags != nil {
			return types.ListNull(nil), comparisonDiags
		}
		return tfFromComparison, nil
	}
	//return types.ListNull(nil), nil
	return tfList, nil
}

// todo turn this into a generic function?
// Purpose specific

func purposeListFromTf(ctx context.Context, l types.List) ([]Purpose, diag.Diagnostics) {
	goObject := make([]Purpose, 0)
	// read from the terraform data into the map
	if diags := l.ElementsAs(ctx, &goObject, false); diags != nil {
		return nil, diags
	}
	return goObject, nil
}

func purposeListFromGo(ctx context.Context, l []Purpose) (types.List, diag.Diagnostics) {

	purposeTypes := map[string]attr.Type{
		"name":            types.StringType,
		"description":     types.StringType,
		"acknowledgement": types.StringType,
	}

	mappedValue, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: purposeTypes}, l)
	if diags != nil {
		return types.ListNull(nil), diags
	}
	return mappedValue, nil
}

func updatePurposeListIfChanged(ctx context.Context, tfList types.List, comparisonList []Purpose) (types.List, diag.Diagnostics) {

	goTfList, diags := purposeListFromTf(ctx, tfList)
	if diags != nil {
		return types.ListNull(nil), diags
	}

	// compare two lists to see if they are equal

	listsAreSame := true
	if len(goTfList) != len(comparisonList) {
		listsAreSame = false
	}
	for i := range goTfList {
		if !reflect.DeepEqual(goTfList[i], comparisonList[i]) {
			listsAreSame = false
		}
	}

	if !listsAreSame {
		tfFromComparison, comparisonDiags := purposeListFromGo(ctx, comparisonList)
		if comparisonDiags != nil {
			return types.ListNull(nil), comparisonDiags
		}
		return tfFromComparison, nil
	}
	//return types.ListNull(nil), nil
	return tfList, nil
}

// String specific
func stringListFromTf(ctx context.Context, l types.List) ([]string, diag.Diagnostics) {
	goObject := make([]string, 0)
	// read from the terraform data into the map
	if diags := l.ElementsAs(ctx, &goObject, false); diags != nil {
		return nil, diags
	}
	return goObject, nil
}

func stringListFromGo(ctx context.Context, l []string) (types.List, diag.Diagnostics) {

	mappedValue, diags := types.ListValueFrom(ctx, types.StringType, l)
	if diags != nil {
		return types.ListNull(nil), diags
	}
	return mappedValue, nil
}

func updateStringListIfChanged(ctx context.Context, tfList types.List, comparisonList []string) (types.List, diag.Diagnostics) {

	goTfList, diags := stringListFromTf(ctx, tfList)
	if diags != nil {
		return types.ListNull(nil), diags
	}

	// compare two lists to see if they are equal

	listsAreSame := true
	if len(goTfList) != len(comparisonList) {
		listsAreSame = false
	}
	for i := range goTfList {
		if !reflect.DeepEqual(goTfList[i], comparisonList[i]) {
			listsAreSame = false
		}
	}

	if !listsAreSame {
		tfFromComparison, comparisonDiags := stringListFromGo(ctx, comparisonList)
		if comparisonDiags != nil {
			return types.ListNull(nil), comparisonDiags
		}
		return tfFromComparison, nil
	}
	//return types.ListNull(nil), nil
	return tfList, nil
}
