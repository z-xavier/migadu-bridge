import { request } from '@umijs/max';
import {
  TokenAddRequest,
  TokenAddResponse,
  TokenListRequest,
  TokenListResponse,
  TokenUpdateRequest,
} from './typings';

export async function queryTokenList(params: TokenListRequest) {
  return request<TokenListResponse>(
    // 'http://127.0.0.1:4523/m1/5903528-5590418-default/api/v1/tokens',
    '/api/v1/tokens',
    {
      method: 'GET',
      params: params,
    },
  );
}

export async function addToken(
  body?: TokenAddRequest,
  options?: { [key: string]: any },
) {
  return request<TokenAddResponse>('/api/v1/tokens', {
    method: 'POST',
    data: body,
    ...(options || {}),
  });
}

export async function updateToken(
  id: string,
  body?: Omit<TokenUpdateRequest, 'id'>,
  options?: { [key: string]: any },
) {
  return request<TokenAddResponse>(`/api/v1/tokens/${id}`, {
    method: 'PUT',
    data: body,
    ...(options || {}),
  });
}

export async function updateTokenStatus(
  id: string,
  status: number,
  options?: { [key: string]: any },
) {
  return request(`/api/v1/tokens/${id}`, {
    method: 'PATCH',
    data: {
      status,
    },
    ...(options || {}),
  });
}

export async function deleteToken(
  id: string,
  options?: { [key: string]: any },
) {
  return request(`/api/v1/tokens/${id}`, {
    method: 'DELETE',
    ...(options || {}),
  });
}
