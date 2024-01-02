import { invalidate, invalidateAll } from '$app/navigation';
import {
  CreateWorksheetRequest,
  UpdateWorksheetRequest,
} from 'pb/dburst/worksheet/v1/service';
import getServices from 'provider/services';

export default {
  createWorksheet: async function (data: CreateWorksheetRequest) {
    const svc = getServices();

    const req = CreateWorksheetRequest.create(data);
    await svc.worksheet.createWorksheet(req).response;
    await invalidateAll();
  },
  updateWorksheet: async (data: UpdateWorksheetRequest) => {
    const svc = getServices();
    console.log(data);

    const req = UpdateWorksheetRequest.create(data);
    await svc.worksheet.updateWorksheet(req).response;
    await invalidate(`/worksheets/${data.id}`);
    await invalidateAll();
  },
};
