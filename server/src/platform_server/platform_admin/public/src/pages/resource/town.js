import React from 'react';
import Request from 'superagent';
import GlobalData from '../../data/global.js';
import httpUtil from '../../util/http.js';

let TownPage = React.createClass({
	getInitialState: function() {
		return {towns: {}, townUrl:""};
	},

	componentDidMount: function() {
		this.refreshTownUrl();
		this.refreshTowns();
	},

	handleTownUrlChange: function(e) {
		this.setState({townUrl: e.target.value});
	},

	handleEditTown: function(version, name) {
		if (version === '' || name === '') {
			alert('不能为空');
			return;
		}

		var towns = this.state.towns[version]
		if (!towns) {
			console.log("the towns of version:" + version + " is not found");
			return
		}
		var town = towns[name]
		if (!town) {
			console.log("the town of name:" + name + " is not found");
			return
		}
		React.findDOMNode(this.refs.town_version).value = version
		React.findDOMNode(this.refs.town_name).value = name
		React.findDOMNode(this.refs.town_path).value = town.split(",")[0]
		React.findDOMNode(this.refs.town_size).value = town.split(",")[1]
	},

	handleDelTown: function(version, name) {
		if (version === '' || name === '') {
			alert('不能为空');
			return;
		}

		let self = this;
		let index = GlobalData.App.lastIndexOf('_');
		let app =  GlobalData.App.substring(0, index);
		let type = GlobalData.App.substring(index+1, GlobalData.App.length);
		let requestPath = '/api/editTown?app=' + app + '&type=' + type + 
			'&version=' + version + '&name=' + name + '&path=' + '&size=';
		Request.get(requestPath)
			.set('Accept', 'application/json')
			.end(function(err, data) {
				let val = httpUtil.handleHttpData(err, data);
				console.dir(val);
				if (val.code === 0) {
					self.refreshTowns();
				}
			});
	},

	handleAddTown: function() {
		let version = React.findDOMNode(this.refs.town_version).value
		let name = React.findDOMNode(this.refs.town_name).value
		let path = React.findDOMNode(this.refs.town_path).value
		let size = React.findDOMNode(this.refs.town_size).value

		if (version === '' || name === '' || path === '') {
			alert('不能为空');
			return;
		}

		let self = this;
		let index = GlobalData.App.lastIndexOf('_');
		let app =  GlobalData.App.substring(0, index);
		let type = GlobalData.App.substring(index+1, GlobalData.App.length);
		let requestPath = '/api/editTown?app=' + app + '&type=' + type + 
			'&version=' + version + '&name=' + name + '&path=' + 
			path + '&size=' + size;
		Request.get(requestPath)
			.set('Accept', 'application/json')
			.end(function(err, data) {
				let val = httpUtil.handleHttpData(err, data);
				console.dir(val);
				if (val.code === 0) {
					self.refreshTowns();
				}
			});
	},

	refreshTowns: function() {
		let self = this;
		let index = GlobalData.App.lastIndexOf('_');
		let app =  GlobalData.App.substring(0, index);
		let requestPath = '/api/towns?app=' + app;

		Request.get(requestPath)
			.set('Accept', 'application/json')
			.end(function(err, data) {
				let val = httpUtil.handleHttpData(err, data);
				self.setState({towns:val.data});
			});
	},

	refreshTownUrl: function() {
		let self = this;
		let index = GlobalData.App.lastIndexOf('_');
		let app =  GlobalData.App.substring(0, index);
		let requestPath = '/api/getTownUrl?app=' + app;

		Request.get(requestPath)
			.set('Accept', 'application/json')
			.end(function(err, data) {
				let val = httpUtil.handleHttpData(err, data);
				self.setState({townUrl:val.data});
			});
	},

	handleEditTownUrl: function() {
		let self = this;
		let index = GlobalData.App.lastIndexOf('_');
		let app =  GlobalData.App.substring(0, index);
		let requestPath = '/api/editTownUrl?app=' + app + '&townurl=' + this.state.townUrl;

		Request.get(requestPath)
			.set('Accept', 'application/json')
			.end(function(err, data) {
				let val = httpUtil.handleHttpData(err, data);
				console.dir(val);
				if (val.code === 0) {
					self.refreshTownUrl();
				}
			});
	},

	render: function() {
		let self = this;
		return (
			<div>
				<div className="form-inline">
					<label className="label-control col-sm-1">baseurl：</label>
					<input className="form-control col-sm-4" type="text" name="patch_url" value={this.state.townUrl} onChange={this.handleTownUrlChange}/>
					<button className="btn-primary form-control" type="button" onClick={this.handleEditTownUrl}>修改</button>
				</div>
				<br></br>
				<table className="table table-bordered table-responsive">
					<thead>
						<tr className="success">
							<th>版本</th>
							<th>城镇名称</th>
							<th>路径</th>
							<th>文件大小(字节)</th>
							<th>操作</th>
						</tr>
					</thead>
					<tbody>
						{Object.keys(this.state.towns).map(function(version) {
							return Object.keys(self.state.towns[version]).map(function(name) {
									return (
										<tr key={version + name}>
											<td>{version}</td>
											<td>{name}</td>
											<td>{self.state.towns[version][name].split(",")[0]}</td>
											<td>{self.state.towns[version][name].split(",")[1]}</td>
											<td>
											<button className="btn-primary" type="button" onClick={function(e){self.handleEditTown(version, name)}}>编辑</button>
											<span>------</span>
											<button className="btn-primary" type="button" onClick={function(e){self.handleDelTown(version, name)}}>删除</button>
											</td>
										</tr>);
								})}
						)}
					</tbody>
				</table>

				<div className="form-inline">
					<p >配置town：</p>
					<input className="form-control" type="number" name="version" ref="town_version" placeholder="版本号" />
					<input className="form-control" type="number" name="name" ref="town_name" placeholder="城镇包名称" />
					<span className="label-control">:</span>
					<input className="form-control" type="text" name="town_patch" ref="town_path" placeholder="town文件" />
					<input className="form-control" type="number" name="town_size" ref="town_size" placeholder="文件大小" />
					<button className="btn-primary" type="button" onClick={this.handleAddTown}>修改</button>
				</div>
			</div>
			);
	}
});

export default TownPage;