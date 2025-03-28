// generated code - do not edit
package models

func GetReverseFieldOwnerName(
	stage *Stage,
	instance any,
	reverseField *ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *Cell:
		switch reverseField.GongstructName {
		// insertion point
		case "Row":
			switch reverseField.Fieldname {
			case "Cells":
				if _row, ok := stage.Row_Cells_reverseMap[inst]; ok {
					res = _row.Name
				}
			}
		}

	case *CellBoolean:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *CellFloat64:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *CellIcon:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *CellInt:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *CellString:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *CheckBox:
		switch reverseField.GongstructName {
		// insertion point
		case "FormDiv":
			switch reverseField.Fieldname {
			case "CheckBoxs":
				if _formdiv, ok := stage.FormDiv_CheckBoxs_reverseMap[inst]; ok {
					res = _formdiv.Name
				}
			}
		}

	case *DisplayedColumn:
		switch reverseField.GongstructName {
		// insertion point
		case "Table":
			switch reverseField.Fieldname {
			case "DisplayedColumns":
				if _table, ok := stage.Table_DisplayedColumns_reverseMap[inst]; ok {
					res = _table.Name
				}
			}
		}

	case *FormDiv:
		switch reverseField.GongstructName {
		// insertion point
		case "FormGroup":
			switch reverseField.Fieldname {
			case "FormDivs":
				if _formgroup, ok := stage.FormGroup_FormDivs_reverseMap[inst]; ok {
					res = _formgroup.Name
				}
			}
		}

	case *FormEditAssocButton:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *FormField:
		switch reverseField.GongstructName {
		// insertion point
		case "FormDiv":
			switch reverseField.Fieldname {
			case "FormFields":
				if _formdiv, ok := stage.FormDiv_FormFields_reverseMap[inst]; ok {
					res = _formdiv.Name
				}
			}
		}

	case *FormFieldDate:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *FormFieldDateTime:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *FormFieldFloat64:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *FormFieldInt:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *FormFieldSelect:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *FormFieldString:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *FormFieldTime:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *FormGroup:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *FormSortAssocButton:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Option:
		switch reverseField.GongstructName {
		// insertion point
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			case "Options":
				if _formfieldselect, ok := stage.FormFieldSelect_Options_reverseMap[inst]; ok {
					res = _formfieldselect.Name
				}
			}
		}

	case *Row:
		switch reverseField.GongstructName {
		// insertion point
		case "Table":
			switch reverseField.Fieldname {
			case "Rows":
				if _table, ok := stage.Table_Rows_reverseMap[inst]; ok {
					res = _table.Name
				}
			}
		}

	case *Table:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T Gongstruct](
	stage *Stage,
	instance *T,
	reverseField *ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *Cell:
		switch reverseField.GongstructName {
		// insertion point
		case "Row":
			switch reverseField.Fieldname {
			case "Cells":
				res = stage.Row_Cells_reverseMap[inst]
			}
		}

	case *CellBoolean:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *CellFloat64:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *CellIcon:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *CellInt:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *CellString:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *CheckBox:
		switch reverseField.GongstructName {
		// insertion point
		case "FormDiv":
			switch reverseField.Fieldname {
			case "CheckBoxs":
				res = stage.FormDiv_CheckBoxs_reverseMap[inst]
			}
		}

	case *DisplayedColumn:
		switch reverseField.GongstructName {
		// insertion point
		case "Table":
			switch reverseField.Fieldname {
			case "DisplayedColumns":
				res = stage.Table_DisplayedColumns_reverseMap[inst]
			}
		}

	case *FormDiv:
		switch reverseField.GongstructName {
		// insertion point
		case "FormGroup":
			switch reverseField.Fieldname {
			case "FormDivs":
				res = stage.FormGroup_FormDivs_reverseMap[inst]
			}
		}

	case *FormEditAssocButton:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *FormField:
		switch reverseField.GongstructName {
		// insertion point
		case "FormDiv":
			switch reverseField.Fieldname {
			case "FormFields":
				res = stage.FormDiv_FormFields_reverseMap[inst]
			}
		}

	case *FormFieldDate:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *FormFieldDateTime:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *FormFieldFloat64:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *FormFieldInt:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *FormFieldSelect:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *FormFieldString:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *FormFieldTime:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *FormGroup:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *FormSortAssocButton:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Option:
		switch reverseField.GongstructName {
		// insertion point
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			case "Options":
				res = stage.FormFieldSelect_Options_reverseMap[inst]
			}
		}

	case *Row:
		switch reverseField.GongstructName {
		// insertion point
		case "Table":
			switch reverseField.Fieldname {
			case "Rows":
				res = stage.Table_Rows_reverseMap[inst]
			}
		}

	case *Table:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
