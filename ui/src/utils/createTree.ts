import _ from 'lodash';

export type WithTree<T> = T & { _children: WithTree<T>[] };

export default function createTree<T>(
  items: T[],
  parentCol: keyof T,
): WithTree<T>[] {
  const data: any = items;
  var groupedByParents = _.groupBy(data, parentCol);
  var catsById = _.keyBy(data, 'id');
  _.each(_.omit(groupedByParents, ''), function (children, parentId) {
    _.set(catsById[parentId], '_children', children);
  });
  return _.orderBy(groupedByParents['']);
}
