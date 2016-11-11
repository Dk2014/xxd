import React from 'react';
import Request from 'superagent';
import {Navigation} from 'react-router';
import httpUtil from '../../util/http.js';
import GlobalData from '../../data/global';

var ServerList = React.createClass({
	mixins: [Navigation],

	getInitialState: function() {
		return {list:[], isCloseType: false}
	},

	componentDidMount: function() {
		let self = this;
		let index = GlobalData.App.lastIndexOf('_');
		let app =  GlobalData.App.substring(0, index);
		let type = GlobalData.App.substring(index+1, GlobalData.App.length);
		Request.get('/api/serverall?App=' + app + '&Type=' + type)
			.set('Accept', 'application/json')
			.end(function(err, data) {
				let val = httpUtil.handleHttpData(err, data);
				self.state.list = val.data;
				self.setState(self.state);
			});
		Request.get('/api/isCloseType?App=' + app + '&Type=' + type)
			.set('Accept', 'application/json')
			.end(function(err, data) {
				let val = httpUtil.handleHttpData(err, data);
				self.state.isCloseType = val.data || false;
				self.setState(self.state);
			});
	},

	componentWillUnmount: function() {
		console.log(' ServerList componentWillUnmount');		
	},

	handleEditClick: function(s) {
		GlobalData['server'] = s;
		this.transitionTo('/server/edit');
	},

	handleCloseTypeChange: function(e) {
		this.state.isCloseType = e.target.checked;
		this.setState(this.state);

		let self = this;
		let index = GlobalData.App.lastIndexOf('_');
		let app =  GlobalData.App.substring(0, index);
		let type = GlobalData.App.substring(index+1, GlobalData.App.length);
		let strIsClose = this.state.isCloseType ? "1": "0";
		Request.get('/api/setIsCloseType?app=' + app + '&type=' + type + '&isCloseType=' + strIsClose)
			.set('Accept', 'application/json')
			.end(function(err, data) {
				let val = httpUtil.handleHttpData(err, data);
				self.state.isCloseType = val.data || false;
				self.setState(self.state);
			});
	},

	getStatusFromId: function(id) {
		switch (id) {
			case 0: 
			return '维护';
			case 1: 
			return '畅通';
			case 2: 
			return '繁忙';
			case 3: 
			return '拥挤';
			default: 
			return '未知';
		}
	},

	getHumanType: function(type) {
		switch (type) {
			case 1: 
			return 'ios';
			case 17: 
			return 'android';
			case 252: 
			return 'ios审核';
			case 255: 
			return '调试';
			default: 
			return '未知';
		}
	},

	render: function() {
		let list = this.state.list;
		let self = this;
		return (
			<div>
				<div className="form-inline">
					<label  className="">全服维护？
						<input type="checkbox" className="form-control" checked={this.state.isCloseType} onChange={this.handleCloseTypeChange}/>	
					</label>
				</div>
				<table className="table table-bordered table-responsive">
					<thead>
						<tr className="success">
							<th>serverId</th>
							<th>名称</th>
							<th>类型</th>
							<th>状态</th>
							<th>开服时间</th>
							<th>新服</th>
							<th>推荐服</th>
							<th>游戏服信息</th>
							<th>备注</th>
						</tr>
					</thead>
					<tbody>
					{list.map(function(server) {
						return (
							<tr key={server.Id}>
								<td>{server.Id}</td>
								<td>{server.Name}</td>
								<td>{self.getHumanType(server.Type)}</td>
								<td>{self.getStatusFromId(server.Status)}</td>
								<td>{new Date(server.OpenTime * 1000).toLocaleString()}</td>
								<td>{server.IsNew.toString()}</td>
								<td>{server.IsHot.toString()}</td>
								<td>
									<table className="table-bordered" >
										{server.GServers.map(function(gsserver) {
											return (<tr key={gsserver.GSId}>
												<td>
													{"gsid:" + gsserver.GSId + 
														" addr:" + gsserver.Ip + ":" + gsserver.Port +
														" rpc:" + gsserver.RPCIp + ":" + gsserver.RPCPort}
												</td>
											</tr>);
										})}
									</table>
								</td>
								<td>
									<button className="btn btn-default" name="edit_server" value="编辑" onClick={function(s){
										return function(e) {
											self.handleEditClick(s);
										}
									}(server)}>
									编辑
									</button>
								</td>
							</tr>
						);
					})}
					</tbody>
				</table>
			</div>
		);
	}
});


export default ServerList;