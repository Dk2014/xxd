#include "pal/tos.h"
#include "tdr/tdr.h"
#include "tlog/tlog.h"
#include "tloghelp/tlogload.h"

//#include <Winsock2.h>
//#pragma comment(lib, "WS2_32.lib")



int main(int argc, char* argv[])
{
	int iRet = 0;

	LPTLOGCTX			pstLogCtx;
	LPTLOGCATEGORYINST	pstCat;

	// write file log
	pstLogCtx = tlog_init_from_file("../../test_file.xml");

	pstCat = tlog_get_category(pstLogCtx, "test");
	if (NULL == pstCat)
	{
		printf("tlog_get_category is error!\n");
		return -1;
	}

	tlog_info(pstCat, 0, 0, "%s|%s|%d|%d", "rolelogin", "2012-07-12", 245335694, 1123);
	tlog_fini_ctx(&pstLogCtx);


	//====================================================
	// write net log
	
	pstLogCtx = tlog_init_from_file("../../test_net.xml");
	if(NULL == pstLogCtx)
	{
		printf("tlog_init_from_file is error!\n");
		return -1;		
	}
	pstCat = tlog_get_category(pstLogCtx, "test");
	if (NULL == pstCat)
	{
		printf("tlog_get_category is error!\n");
		return -1;
	}

  tlog_info(pstCat, 0, 0, "%s|%s|%d|%d", "rolelogin", "2012-07-12", 245335694, 1123);
	tlog_fini_ctx(&pstLogCtx);

	//=====================================================
	pstLogCtx = tlog_init_from_file("../../test_vec.xml");
	if(NULL == pstLogCtx)
	{
		printf("tlog_init_from_file is error!\n");
		return -1;		
	}
	pstCat = tlog_get_category(pstLogCtx, "test");
	if (NULL == pstCat)
	{
		printf("tlog_get_category is error!\n");
		return -1;
	}

	tlog_info(pstCat, 0, 0, "%s|%s|%d|%d", "rolelogin", "2012-07-12", 245335694, 1123);
	tlog_fini_ctx(&pstLogCtx);
	return 0;
}

