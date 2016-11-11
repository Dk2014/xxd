import React from 'react';
import {State} from'react-router';
import ItemData from '../data/item.js';
import NotFound from '../pages/misc/notfound.js';
import PageHeader from '../components/pageHeader.js';
import ServerList from '../pages/server/list.js';
import EditServer from '../pages/server/edit.js';
import PatchPage from '../pages/resource/patch.js';
import TownPage from '../pages/resource/town.js';
import GlobalData from '../data/global.js';
import VersionPage from '../pages/misc/version.js';

let PageMap={};
PageMap['/server/list'] = ServerList;
PageMap['/server/add'] = EditServer;
PageMap['/server/edit']  = EditServer;
PageMap['/resource/patch']  = PatchPage;
PageMap['/resource/town']  = TownPage;
PageMap['/other/version']  = VersionPage;

//style
let pageStyle = {
  marginTop: '15',
};

let itemStyle = {
  margin: '15'
};

let Item = React.createClass({

  contextTypes: {
    router: React.PropTypes.func.isRequired
  },

  render: function () {
    let params = this.context.router.getCurrentParams();
    let path = '/' + params.categoryPath + '/' + params.itemPath;

    let page = PageMap[path];
    if (!page) {
      page = React.createElement(NotFound);
    }else {
      page = React.createElement(PageMap[path]);
    }

    let category = ItemData.lookupCategory(params.categoryPath);
    let item = ItemData.lookupItem(category.path, params.itemPath);
    let categoryName = category ? category.name: params.categoryPath;
    let itemName = item ? item.name: params.itemPath;
    params = {category:categoryName, item:itemName};

    //清除缓存服务器信息
    if (path != "/server/edit") {
      GlobalData.server = {};
    }

    return (
      <div style={itemStyle}>
        <PageHeader params={params}>
          <h>test</h>
        </PageHeader>
        <div style={pageStyle}>
          {page}
        </div>
      </div>
    );
  }
});

export default Item;