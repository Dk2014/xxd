const data = [
  {
    name: '游戏服',
    path: 'server',
    items: [
      { name: '列表', path: 'list'},
      { name: '添加', path: 'add'}
    ]
  },
  {
    name: '静态资源',
    path: 'resource',
    items: [
      { name: 'Patch', path: 'patch'},
      { name: 'Town', path: 'town'}
    ]
  },
  {
    name: '其他',
    path: 'other',
    items: [
      { name: '版本', path: 'version'}
    ]
  }
];

const dataMap = data.reduce(function (map, category) {
  category.itemsMap = category.items.reduce(function (itemsMap, item) {
    itemsMap[item.path] = item;
    return itemsMap;
  }, {});
  map[category.path] = category;
  return map;
}, {});

const Data = {};
Data.getAll = function () {
  return data;
};

Data.lookupCategory = function (name) {
  return dataMap[name];
};

Data.lookupItem = function (categoryPath, itemPath) {
  return dataMap[categoryPath].itemsMap[itemPath];
};

export default Data;