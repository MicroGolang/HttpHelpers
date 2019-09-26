/*******************************************************************************
** @Author:					Thomas Bouder <Tbouder>
** @Email:					Tbouder@protonmail.com
** @Date:					Thursday 26 September 2019 - 13:58:44
** @Filename:				Args.go
**
** @Last modified by:		Tbouder
** @Last modified time:		Thursday 26 September 2019 - 14:00:12
*******************************************************************************/

package		httpHelper

import		"github.com/julienschmidt/httprouter"
import		"encoding/json"

/*GetArg **********************************************************************
*	Decrypt the args to the type of myValue and assign it to myValue
******************************************************************************/
func	GetArg(ps httprouter.Params, argName string, myValue interface{}) {
	argument := ps.ByName(argName)
	json.Unmarshal([]byte(argument), myValue)
}

/*FormatArgs ******************************************************************
*	Transform an argument (no matter the type) to a string for later decryption
******************************************************************************/
func	FormatArgs(myValue interface{}) string {
	str, _ := json.Marshal(myValue)
	return (string(str))
}

/*UnFormatArgs ****************************************************************
*	Decrypt and argument formated with FormatArgs
******************************************************************************/
func	UnFormatArgs(variable string, myValue interface{}) {
	json.Unmarshal([]byte(variable), myValue)
}