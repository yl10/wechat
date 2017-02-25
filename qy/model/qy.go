package model

//Qy 企业号
type Qy struct {
	CorpID string //CorpID是企业号的标识，每个企业号拥有一个唯一的CorpID；
	Secret string //Secret是管理组凭证密钥
}
