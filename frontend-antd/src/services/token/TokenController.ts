import { request } from '@umijs/max';
import { TokenListRequest, TokenListResponse } from './typings';

export async function queryTokenList(params: TokenListRequest) {
  return request<TokenListResponse>(
    'http://127.0.0.1:4523/m1/5903528-5590418-default/api/v1/tokens',
    {
      method: 'GET',
      params: params,
    },
  );
}
