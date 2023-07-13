package dbops

import
(
	"testing"
)
func TestDbConnection(t *testing.T){
	err:=ConnectToDB()
	if(err!=nil){
		t.Errorf("error is %s\n",err.Error())
	}
}