import { goto, invalidate, invalidateAll } from '$app/navigation';
import {
  DeleteConnectionRequest,
  UpdateConnectionRequest,
} from 'pb/dburst/session/v1/service';
import getServices from 'provider/services';

const actions = {
  delete: async (id: string) => {
    const svc = getServices();

    const req = DeleteConnectionRequest.create({
      connectionId: id,
    });
    await svc.session.deleteConnection(req).response;
    await invalidate('/connections');
    goto('/connections');
  },
  update: async (data: UpdateConnectionRequest) => {
    const svc = getServices();

    const req = UpdateConnectionRequest.create(data);
    await svc.session.updateConnection(req).response;

    await invalidate('/connections');
    goto('/connections');
  },
};

export default actions;
