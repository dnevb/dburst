import { ListWorksheetRequest } from 'pb/dburst/worksheet/v1/service';
import type { LayoutLoad } from './$types';

export const load: LayoutLoad = async (e) => {
  const { api } = await e.parent();

  const req = ListWorksheetRequest.create();
  return api?.worksheet.listWorksheet(req).response;
};
