export enum MockProvider {
  addy = 'addy',
  sl = 'sl',
}

export interface TokenListRequest {
  description?: string;
  expiryAtBegin?: number;
  expiryAtEnd?: number;
  lastCalledAtBegin?: number;
  lastCalledAtEnd?: number;
  mockProvider?: MockProvider;
  /**
   * orderBy=mockProvider:desc&orderBy=expiryAt:desc
   */
  orderBy?: null | string;
  page?: number;
  pageSize?: number;
  status?: number;
  targetEmail?: string;
  updatedAtBegin?: number;
  updatedAtEnd?: number;
}

export interface TokenListResponse {
  code: string;
  data: TokenListResponseData;
  message: string;
}

export interface TokenListResponseData {
  list: List[];
  page: number;
  pageSize: number;
  total: number;
}
export interface TokenListItem {
  createdAt: number;
  description: string;
  expiryAt: number;
  id: string;
  lastCalledAt: number;
  mockProvider: string;
  status: number;
  targetEmail: string;
  token: string;
  updatedAt: number;
}

export interface TokenAddRequest {
  description?: string;
  expiryAt: number;
  mockProvider: MockProvider;
  targetEmail: string;
}

export interface TokenAddResponse {
  code: string;
  data: TokenListItem;
  message: string;
}

export interface TokenUpdateRequest {
  id: string;
  description?: string;
  expiryAt: number;
}

export interface TokenUpdateStatusRequest {
  id: string;
  status: number;
}
