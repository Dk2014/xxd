
let util = {};

util.handleHttpData = function(err, data) {
	if(err) {
		alert(err);
		return {};
	}
	let body = data.body;
	if (body.code !== 0) {
		console.log(body.msg);
		return {};
	}
	return body;
}

export default util;