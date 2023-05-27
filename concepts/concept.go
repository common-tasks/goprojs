package concepts

type Cloud interface{
	Connect(url string)bool
	Disconnect() bool

}
type AzureCloud struct{
	ClientID string
	TenantID string
}
func TestIFace(){

}