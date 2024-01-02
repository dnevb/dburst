import { goto } from '$app/navigation';
import { CreateConnectionRequest } from 'pb/dburst/session/v1/service';
import getServices from 'provider/services';
import type FormSubmitHandler from 'utils/handleSubmit';

const handleSubmit: FormSubmitHandler = async (values: any) => {
  const req = CreateConnectionRequest.create(values);

  const svc = getServices();

  const { response } = svc.session.createConnection(req);
  await response;
  goto('/connections');
};

export default handleSubmit;
