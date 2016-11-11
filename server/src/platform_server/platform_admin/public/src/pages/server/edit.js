import React from 'react';
import GlobalData from '../../data/global.js';
import Request from 'superagent';
import httpUtil from '../../util/http.js';
import {Navigation} from 'react-router';

let EditPage = React.createClass({
	mixins: [Navigation],

	getInitialState: function() {
		let state = {
			server:{Type:1},
			toAddGServer: {}
		};

		for(let key in GlobalData.server) {
			state.server[key] = GlobalData.server[key];
		}

		//清空缓存的数据库信息
		GlobalData.server = {};

		state.server.OpenTime = state.server.OpenTime || 2402531200;
		state.server.GServers = state.server.GServers || [];

		state.isAdd = !!state.server.Id ? false: true;
		return state;
	},

	handleTypeChange: function(e) {
		this.state.server.Type = Number(e.target.value);
		this.setState(this.state);
	},

	handleAddGServer: function(e) {
		let gsid = Number(React.findDOMNode(this.refs.tx_gserver_gsid).value);
		let hd = React.findDOMNode(this.refs.tx_gserver_hd).value === "true";
		let ip = React.findDOMNode(this.refs.tx_gserver_ip).value;
		let port = React.findDOMNode(this.refs.tx_gserver_port).value;
		let rpcip = React.findDOMNode(this.refs.tx_gserver_rpcip).value;
		let rpcport = React.findDOMNode(this.refs.tx_gserver_rpcport).value;

		if (!gsid) {
			alert('gsid is empty!');
			return;
		}

		for(let i = 0; i < this.state.server.GServers.length; i++) {
			if (gsid === this.state.server.GServers[i].GSId) {
				alert('重复的gsid!!')
				return;
			}
		}

		this.state.server.GServers.push({
			GSId: gsid,
			Ip: ip,
			HD: hd,
			Port: port,
			RPCIp: rpcip,
			RPCPort: rpcport,
		});

		this.setState(this.state);
	},

	handleDelGServer: function(gsid) {
		for (let i = 0; i < this.state.server.GServers.length; i++) {
			if(this.state.server.GServers[i].GSId === gsid) {
				this.state.server.GServers.splice(i,1);
				this.setState(this.state);
				break;
			}
		}
	},

	handleIsAddGServer: function(e) {
		this.setState({isToAddGServer: !this.state.isToAddGServer});
	},

	handleIdChange: function(e) {
		this.state.server.Id = Number(e.target.value);
		this.setState(this.state);
	},

	handleNameChange: function(e) {
		this.state.server.Name = e.target.value;
		this.setState(this.state);
	},

	handleStatusChange: function(e) {
		this.state.server.Status = Number(e.target.value);
		this.setState(this.state);
	},

	handleOpenTimeChange: function(e) {
		this.state.server.OpenTime = Number(e.target.value);
		this.setState(this.state);
	},

	handleIsNewChange: function(e) {
		this.state.server.IsNew = e.target.value === "true";
		this.setState(this.state);
	},

	handleIsHotChange: function(e) {
		this.state.server.IsHot = e.target.value === "true";
		this.setState(this.state);
	},

	handleSaveServer: function() {
		if (!this.state.server.Id || !this.state.server.Name || !this.state.server.OpenTime) {
			return alert('id or name or opentime should not be empty');
		}

		console.dir(this.state.server);
		let self = this;
		let app =  GlobalData.App.substring(0, GlobalData.App.lastIndexOf('_'));
		Request.post('/api/editServer?App=' + app)
			.send(JSON.stringify(this.state.server))
			.set('Accept', 'application/json')
			.end(function(err, data) {
				let val = httpUtil.handleHttpData(err, data);
				if (!!val) {
					self.transitionTo('/server/list');
				}
			});
	},

	//add gserver ui section
	renderAddGServer: function () {
		return this.state.isToAddGServer ? (
			<div className="form-horizontal">
				<div className="form-group">
					<label  className="col-sm-1 control-label">GSId</label>
					<div className="col-sm-2">
						<input type="number" className="form-control" name="GSId" ref="tx_gserver_gsid"/>	
					</div>
				</div>
				<div className="form-group">
					<label  className="col-sm-1 control-label">互动服</label>
					<div className="col-sm-2">
						<select className="form-control" name="HD" ref="tx_gserver_hd">
							<option value="false">否</option>
							<option value="true">是</option>
						</select>
					</div>
				</div>
				<div className="form-group">
					<label  className="col-sm-1 control-label">IP</label>
					<div className="col-sm-2">
						<input type="text" className="form-control" name="ip" ref="tx_gserver_ip"/>
					</div>
				</div>
				<div className="form-group">
					<label className="col-sm-1 control-label">Port</label>
					<div className="col-sm-2">
						<input type="number" className="form-control" name="port" ref="tx_gserver_port"/>
					</div>
				</div>
				<div className="form-group">
					<label className="col-sm-1 control-label">RPCIp</label>
					<div className="col-sm-2">
						<input type="text" className="form-control" name="rpcip" ref="tx_gserver_rpcip"/>
					</div>
				</div>
				<div className="form-group">
					<label className="col-sm-1 control-label">RPCPort</label>
					<div className="col-sm-2">
						<input type="number" className="form-control" name="rpcport" ref="tx_gserver_rpcport"/>
					</div>
				</div>
				<div className="col-sm-offset-1 col-sm-2">
					<button className="btn btn-default" onClick={this.handleAddGServer}>添加</button>
				</div>
			</div>): null
	},

	render: function() {
		let self = this;
		return (
			<div>
				<form className="form-horizontal">
					<div className="form-group">
						<label className="col-sm-1 control-label">Id</label>
						<div className="col-sm-2">
							<input type="number" disabled={this.state.isAdd?null:"disabled"} className="form-control" name="Id" value={this.state.server.Id} onChange={this.handleIdChange}/>	
						</div>
					</div>
					<div className="form-group">
						<label className="col-sm-1 control-label">名称</label>
						<div className="col-sm-2">
							<input type="text" className="form-control" name="Name"  ref="tx_server_name" value={this.state.server.Name} onChange={this.handleNameChange}/>
						</div>
					</div>
					<div className="form-group">
						<div>
							<label className="col-sm-1 control-label">类型 </label>
						</div>
						<div className="col-sm-4">
							<div className="radio-inline">
							  	<label>
							    	<input type="radio" name="android" id="typeRadio" value="17" checked={this.state.server.Type===17?true:false} onChange={this.handleTypeChange}/>
							    	安卓
							  	</label>
							</div>
							<div className="radio-inline">
							  	<label>
							    	<input type="radio" name="iOS" id="typeRadio" value="1" checked={this.state.server.Type===1?true:false} onChange={this.handleTypeChange}/>
							    	iOS
							  	</label>
							</div>
							<div className="radio-inline">
							  	<label>
							    	<input type="radio" name="iOS审核" id="typeRadio" value="252" checked={this.state.server.Type===252?true:false} onChange={this.handleTypeChange}/>
							    	iOS审核
							  	</label>
							</div>
							<div className="radio-inline">
							  	<label>
							    	<input type="radio" name="iOS" id="typeRadio" value="255" checked={this.state.server.Type===255?true:false} onChange={this.handleTypeChange}/>
							    	调试
							  	</label>
							</div>
						</div>
					</div>
					<div className="form-group">
						<label className="col-sm-1 control-label">状态</label>
						<div className="col-sm-2">
							<select className="form-control" name="status" value={this.state.server.Status} onChange={this.handleStatusChange}>
								<option value="0">维护</option>
								<option value="1">畅通</option>
								<option value="2">繁忙</option>
								<option value="3">拥挤</option>
							</select>						
						</div>
					</div>
					<div className="form-group">
						<label className="col-sm-1 control-label">新服</label>
						<div className="col-sm-2">
							<select className="form-control" name="isNew" value={this.state.server.IsNew} onChange={this.handleIsNewChange}>
								<option value="false">否</option>
								<option value="true">是</option>
							</select>						</div>
					</div>
					<div className="form-group">
						<label className="col-sm-1 control-label">推荐服</label>
						<div className="col-sm-2">
							<select className="form-control" name="isHot" value={this.state.server.IsHot} onChange={this.handleIsHotChange}>
								<option value="false">否</option>
								<option value="true">是</option>
							</select>
						</div>
					</div>
					<div className="form-group">
						<label className="col-sm-1 control-label">开服时间(秒为单位的时间戳)</label>
						<div className="col-sm-2">
							<input type="number" className="form-control" name="OpenTime" onChange={this.handleOpenTimeChange} value={this.state.server.OpenTime}/>
						</div>
					</div>
				</form>
				<label>GServer信息</label>
				<table className="table table-bordered table-responsive" >
					<tr className="success">
						<th>GSId</th>
						<th>互动服</th>
						<th>Ip</th>
						<th>Port</th>
						<th>RPCIp</th>
						<th>PRCPort</th>
						<th>备注</th>
					</tr>
					<tbody>
					{this.state.server.GServers.map(function(gserver) {
							return (
								<tr key={gserver.GSId}>
									<td>{gserver.GSId}</td>
									<td>{gserver.HD.toString()}</td>
									<td>{gserver.Ip}</td>
									<td>{gserver.Port}</td>
									<td>{gserver.RPCIp}</td>
									<td>{gserver.RPCPort}</td>
									<td>
										<button className="btn btn-default" onClick={(e)=>self.handleDelGServer(gserver.GSId)}>删除</button>
									</td>
								</tr>
							);
						})
					}
					</tbody>
				</table>
				<div>
					<button type="submit" className="btn btn-default" onClick={this.handleIsAddGServer}>{this.state.isToAddGServer? "^":"+"}</button>
				</div>
				{this.renderAddGServer()}
				<br></br>
				<br></br>
				<br></br>
				<div>
					<button type="submit" className="btn btn-default" onClick={this.handleSaveServer}>保存修改</button>
				</div>
			</div>
		)
	}
});

export default EditPage;