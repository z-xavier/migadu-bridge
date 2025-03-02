import { request } from '@umijs/max';
import { AliasListRequest, AliasListResponse } from './typings';

export function query(params: AliasListRequest) {
  return request<AliasListResponse>(
    // 'http://127.0.0.1:4523/m1/5903528-5590418-default/api/v1/tokens',
    '/api/v1/aliases',
    {
      method: 'GET',
      params: params,
    },
  );
}
