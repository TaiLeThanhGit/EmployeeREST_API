package Employee

import (
	Lib "employee/common/lib"
	Messages "employee/common/messages"
	Types "employee/common/types"
)

var employees = make([]Types.Employee, 0)

func GetEmployees() []Types.Employee {

	Lib.ShowDebugInfo("GetEmployees ", employees)

	return employees
}

func GetEmployeeById(id int) Types.Employee {
	var employee Types.Employee

	for _, item := range employees {
		if item.ID == id {
			employee = item
			break
		}
	}

	Lib.ShowDebugInfo("GetEmployeeById", employee)

	return employee
}

func DeleteEmployeeById(id int) (bool, string) {
	result := false
	message := Messages.EMPLOYEE_NOT_FOUND

	for index, item := range employees {
		if item.ID == id {
			employees = append(employees[:index], employees[index+1:]...)

			result = true
			message = Messages.DELETE_EMPLOYEE_SUCCESSFULLY
			break
		}
	}

	Lib.ShowDebugInfo("DeleteEmployeeById", result)

	return result, message
}

func UpdateEmployeeById(employee Types.Employee) (bool, string) {
	result := false
	message := Messages.EMPLOYEE_NOT_FOUND

	for index, item := range employees {
		if item.ID == employee.ID {
			employees[index] = employee

			result = true
			message = Messages.UPDATE_EMPLOYEE_SUCCESSFULLY
			break
		}
	}

	Lib.ShowDebugInfo("UpdateEmployeeById", result)

	return result, message
}
func AddEmployee(employee Types.Employee) (bool, string) {
	result := true
	message := Messages.ADD_EMPLOYEE_SUCCESSFULLY

	employees = append(employees, employee)
	/* check more errors when adding employee */

	Lib.ShowDebugInfo("AddEmployee", employees)

	return result, message

}
