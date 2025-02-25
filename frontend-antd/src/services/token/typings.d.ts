export interface TokenListRequest {
  description?: string;
  expiryAtBegin?: number;
  expiryAtEnd?: number;
  lastCalledAtBegin?: number;
  lastCalledAtEnd?: number;
  mockProvider?: string;
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
