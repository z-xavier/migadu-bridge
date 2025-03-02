import { request } from '@umijs/max';
import { CallLogListRequest, CallLogListResponse } from './typings';

export function query(params: CallLogListRequest) {
  return request<CallLogListResponse>(
    // 'http://127.0.0.1:4523/m1/5903528-5590418-default/api/v1/tokens',
    '/api/v1/calllogs',
    {
      method: 'GET',
      params: params,
    },
  );
}
