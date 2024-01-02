import {
  GetWorksheetRequest,
  UpdateWorksheetRequest,
} from 'pb/dburst/worksheet/v1/service';
import type { PageLoad } from './$types';
import type { Worksheet } from 'pb/dburst/worksheet/v1/worksheet';
import getServices from 'provider/services';

export const load: PageLoad = async (e) => {
  const { api } = await e.parent();

  const req = GetWorksheetRequest.create({
    worksheetId: e.params.worksheetid,
  });

  return api?.worksheet.getWorksheet(req).response;
};

export const _actions = {
  update: async (values: Worksheet) => {
    const api = getServices();

    const req = UpdateWorksheetRequest.create(values);
    await api.worksheet.updateWorksheet(req);
  },
};

export const _debounce = (callback: Function, waitTime: number) => {
  let timeout: number;
  return (...args: any[]) => {
    clearTimeout(timeout);
    timeout = setTimeout(() => callback(...args), waitTime);
  };
};
