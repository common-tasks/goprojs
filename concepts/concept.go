package concepts

type Cloud interface {
	Connect(url string) bool
	Disconnect() bool
}
type AzureCloud struct {
	ClientID string
	TenantID string
}

type AWS struct {
	UserName string
	Pwd      string
	ID       string
}

func (azure AzureCloud) Connect(url string) bool {
	return true
}
func (azure AzureCloud) Disconnect() bool {
	return true
}
func (aws AWS) Connect(url string) bool {
	return true
}
func (aws AWS) Disconnect() bool {
	return true
}
func TestIFace() {
	azure:=AzureCloud{ClientID: "99",TenantID: "12"}
	ConnectToCloud(azure)
	DisConnectCloud(azure)
	aws:=AWS{UserName: "kk",Pwd: "**",ID: "xyz"}
	ConnectToCloud(aws)
	DisConnectCloud(aws)
}
func ConnectToCloud(cloud Cloud){
	(cloud).Connect("")
}
func DisConnectCloud(cloud Cloud){
	(cloud).Disconnect()
}