package module

type MoneyState struct {
	Openid       string //从手Q登录态中获取的openid的值
	Openkey      string //从手Q登录态中获取的access_token的值
	PayToken     string //从手Q登录态中获取的pay_token的值
	Pfkey        string //跟平台来源和openkey根据规则生成的一个密钥串
	Zoneid       int    //区服ID，sid
	Pf           string //平台来源，平台-注册渠道-系统运行平台-安装渠道-业务自定义。从MSDK的getpf接口获取。例如： qq_m_qq-2001-android-2011-xxxx
	PlatformType int8   //服务器类型17-安卓微信，18-安卓手q
}
