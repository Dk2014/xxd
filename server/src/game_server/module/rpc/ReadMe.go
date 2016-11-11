package rpc

/*

   本rpc模块实现游戏服（简称gs）与gs；gs与互动服（简称hd）的异步通信逻辑；

   文件命名规则如下：
   		- 由gs处理的rpc逻辑统一放在rpc_moduleName格式的文件中
   		- 由hd处理的rpc逻辑统一放在rpc_global_moduleName格式文件中
   		- 由idip处理的rpc逻辑统一放在idip_rpc_moduleName格式文件中
*/
