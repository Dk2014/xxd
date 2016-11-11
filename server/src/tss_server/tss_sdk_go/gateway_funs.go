package tss_sdk_go

/*
#include "tss_sdk_anti.h"

#include <stdio.h>

#if defined(WIN32) || defined(_WIN64)
#include <windows.h>
#include <winsock2.h>
#else
#include <string.h>
#include <sys/time.h>
#endif

// The gateway function
TssSdkProcResult on_send_data_to_client_cgo(TssSdkAntiSendDataInfoEx *anti_data)
{
	//return on_send_data_to_client((*anti_data).openid_.openid_,(*anti_data).openid_.openid_len_,(int)((*anti_data).plat_id_),(*anti_data).anti_data_,(*anti_data).anti_data_len_);
	return on_send_data_to_client((*anti_data).openid_.openid_,(int)((*anti_data).plat_id_),(*anti_data).anti_data_,(int)((*anti_data).anti_data_len_));
}

TssSdkProcResult add_user_cgo(TssSdkAntiInterfEx *anti_interf_, char *openid, int plat_id, char *client_ip, int client_ver,char *role_name)
{
	TssSdkAntiAddUserInfoEx user_info;
	memset(&user_info,0,sizeof(user_info));

	user_info.openid_.openid_ = (unsigned char*)openid;
	user_info.openid_.openid_len_ = strlen(user_info.openid_.openid_);
	user_info.plat_id_ = plat_id;
	user_info.client_ip_ = inet_addr(client_ip);
	user_info.client_ver_ = client_ver;
	user_info.role_name_ = role_name;

	return anti_interf_->add_user_(&user_info);
}

TssSdkProcResult del_user_cgo(TssSdkAntiInterfEx *anti_interf_,char *openid, int plat_id)
{
	TssSdkAntiDelUserInfoEx user_info;

	user_info.openid_.openid_ = (unsigned char*)openid;
	user_info.openid_.openid_len_ = strlen(user_info.openid_.openid_);
	user_info.plat_id_ = plat_id;

	return anti_interf_->del_user_(&user_info);
}

TssSdkProcResult recv_anti_data_cgo(TssSdkAntiInterfEx *anti_interf_,char *openid, int plat_id, unsigned char *anti_data, unsigned int data_len)
{
	TssSdkAntiRecvDataInfoEx pkg_info;

    pkg_info.openid_.openid_ = (unsigned char*)openid;
    pkg_info.openid_.openid_len_ = strlen(pkg_info.openid_.openid_);
    pkg_info.plat_id_ = plat_id;
    pkg_info.anti_data_ = anti_data;
    pkg_info.anti_data_len_ = data_len;

    return anti_interf_->recv_anti_data_(&pkg_info);
}

*/
import "C"
