import React from 'react';
import Request from 'superagent';
import GlobalData from '../../data/global.js';
import httpUtil from '../../util/http.js';

let TownPage = React.createClass({
	getInitialState: function() {
		return {};
	},

	componentDidMount: function() {
		this.requestVersionInfo();
		this.requestDarkInfo();
	},

	//event
	handleMinVersionChange: function(e) {
		this.state.MinVersion = e.target.value;
		this.setState(this.state);
	},

	handleAuditVersionChange: function(e) {
		this.state.AuditVersion = e.target.value;
		this.setState(this.state);
	},

	handleUpgradeUrlChange: function(e) {
		this.state.UpgradeUrl = e.target.value;
		this.setState(this.state);
	},

	handleDarkSidsChange: function(e) {
		this.state.DarkSids = e.target.value;
		this.setState(this.state);
	},

	handleDarkVersionChange:function(e) {
		this.state.DarkVersion = e.target.value;
		this.setState(this.state);
	},

	handleIsCompatibleChange: function(e) {
		this.state.DarkIsCompatible = e.target.value === "1";
		this.setState(this.state);
	},

	handleDarkUpgradeUrlChange: function(e) {
		this.state.DarkUpgradeUrl = e.target.value;
		this.setState(this.state);
	},

	//action
	handleSaveVersions: function() {
		this.state.MinVersion = this.state.MinVersion || "0"
		this.state.AuditVersion = this.state.AuditVersion || "0"

		let self = this;
		let index = GlobalData.App.lastIndexOf('_');
		let app =  GlobalData.App.substring(0, index);
		let type = GlobalData.App.substring(index+1, GlobalData.App.length);
		Request.get('/api/setVersionInfo?app=' + app + '&type=' + type +
				'&minVersion=' + this.state.MinVersion +
				'&auditVersion=' + this.state.AuditVersion + 
				'&upgradeUrl=' + this.state.UpgradeUrl)
			.set('Accept', 'application/json')
			.end(function(err, data) {
				let val = httpUtil.handleHttpData(err, data);
				if (val.code === 0) {
					self.requestVersionInfo();
				}
			});
	},

	handleSaveDarkInfo: function() {
		let self = this;
		let index = GlobalData.App.lastIndexOf('_');
		let app =  GlobalData.App.substring(0, index);
		let type = GlobalData.App.substring(index+1, GlobalData.App.length);
		Request.get('/api/setDarkInfo?app=' + app + '&type=' + type +
				'&sids=' + this.state.DarkSids + 
				'&version=' + this.state.DarkVersion +
				'&isCompatible=' + this.state.DarkIsCompatible + 
				'&upgradeUrl=' + this.state.DarkUpgradeUrl)
			.set('Accept', 'application/json')
			.end(function(err, data) {
				let val = httpUtil.handleHttpData(err, data);
				console.log(val);
				if (val.code === 0) {
					self.requestDarkInfo();
				}
				alert(val.msg);
			});
	},

	handleDelDarkInfo: function() {
		let self = this;
		let index = GlobalData.App.lastIndexOf('_');
		let app =  GlobalData.App.substring(0, index);
		let type = GlobalData.App.substring(index+1, GlobalData.App.length);
		Request.get('/api/delDarkInfo?app=' + app + '&type=' + type)
			.set('Accept', 'application/json')
			.end(function(err, data) {
				let val = httpUtil.handleHttpData(err, data);
				if (val.code === 0) {
					self.setState({
						DarkSids: "",
						DarkVersion: "",
						DarkIsCompatible: "0",
						DarkUpgradeUrl: "",
					});
				}
			});
	},

	requestVersionInfo: function() {
		let self = this;
		let index = GlobalData.App.lastIndexOf('_');
		let app =  GlobalData.App.substring(0, index);
		let type = GlobalData.App.substring(index+1, GlobalData.App.length);
		Request.get('/api/getVersionInfo?app=' + app + '&type=' + type)
			.set('Accept', 'application/json')
			.end(function(err, data) {
				let val = httpUtil.handleHttpData(err, data);
				self.setState({MinVersion: val.data.MinVersion, 
								AuditVersion: val.data.AuditVersion,
								UpgradeUrl: val.data.UpgradeUrl});
			});
	},

	requestDarkInfo: function() {
		let self = this;
		let index = GlobalData.App.lastIndexOf('_');
		let app =  GlobalData.App.substring(0, index);
		let type = GlobalData.App.substring(index+1, GlobalData.App.length);
		Request.get('/api/getDarkInfo?app=' + app + '&type=' + type)
			.set('Accept', 'application/json')
			.end(function(err, data) {
				let val = httpUtil.handleHttpData(err, data);
				if(val.data) {
					self.setState({
						DarkSids: val.data.Sids,
						DarkVersion: val.data.Version,
						DarkIsCompatible: val.data.IsCompatible,
						DarkUpgradeUrl: val.data.UpgradeUrl,
					});
				}
			});
	},

	render: function() {
		return (
			<div>
				<div>
					<h4>版本配置：</h4>
				</div>
				<form className="form-horizontal" >
					<div className="form-group">
						<label  className="col-sm-1 control-label">最小版本</label>
						<div className="col-sm-2">
							<input type="number" className="form-control" value={this.state.MinVersion} onChange={this.handleMinVersionChange}/>	
						</div>
					</div>
					<div className="form-group">
						<label  className="col-sm-1 control-label">最新客户端地址</label>
						<div className="col-sm-4">
							<input type="text" className="form-control" value={this.state.UpgradeUrl} onChange={this.handleUpgradeUrlChange}/>
						</div>
					</div>
					<div className="form-group">
						<label  className="col-sm-1 control-label">审核版本</label>
						<div className="col-sm-2">
							<input type="number" className="form-control"  value={this.state.AuditVersion} onChange={this.handleAuditVersionChange}/>	
						</div>
					</div>
				</form>
				<div>
					<button className="btn btn-default" onClick={this.handleSaveVersions}>修改</button>
				</div>
				{/*灰度服*/}
				<br></br>
				<div>
					<h4>灰度服配置：</h4>
				</div>
				<form className="form-horizontal" >
					<div className="form-group">
						<label  className="col-sm-1 control-label">服务器ID</label>
						<div className="col-sm-2">
							<input type="text" className="form-control" name="sids" placeholder="类似11001:11002" onChange={this.handleDarkSidsChange} value={this.state.DarkSids}/>	
						</div>
					</div>
					<div className="form-group">
						<label  className="col-sm-1 control-label">版本号</label>
						<div className="col-sm-2">
							<input type="number" className="form-control" placeholder="灰度客户端版本号" value={this.state.DarkVersion} onChange={this.handleDarkVersionChange}/>
						</div>
					</div>
					<div className="form-group">
						<label className="col-sm-1 control-label">是否兼容非灰度服</label>
						<div className="col-sm-2">
							<select className="form-control" name="isCompatible" value={this.state.DarkIsCompatible} onChange={this.handleIsCompatibleChange}>
								<option value="0">否</option>
								<option value="1">是</option>
							</select>
						</div>
					</div>
					<div className="form-group">
						<label  className="col-sm-1 control-label">灰度客户端地址</label>
						<div className="col-sm-2">
							<input type="number" className="form-control" placeholder="若没有请空着" value={this.state.DarkUpgradeUrl} onChange={this.handleDarkUpgradeUrlChange}/>
						</div>
					</div>
				</form>
				<div>
					<button className="btn btn-default" onClick={this.handleSaveDarkInfo}>修改</button>
				</div>
				<br></br>
				<div>
					<button className="btn btn-default" onClick={this.handleDelDarkInfo}>删除灰度服配置</button>
				</div>
			</div>
		);
	}

});

export default TownPage;