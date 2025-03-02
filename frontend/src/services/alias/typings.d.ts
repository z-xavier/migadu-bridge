export interface AliasListRequest {
  alias?: string;
  mockProvider?: MockProvider;
  orderBy?: string;
  page?: number;
  pageSize?: number;
  targetEmail?: string;
}

export enum MockProvider {
  Addy = 'addy',
  Sl = 'sl',
}

export interface AliasListResponse {
  code: string;
  data: AliasList;
  message: string;
}

export interface AliasList {
  list: List[];
  page: number;
  pageSize: number;
  total: number;
}

export interface AliasItem {
  alias?: string;
  callLogId?: string;
  id?: number;
  mockProvider?: string;
  targetEmail?: string;
  tokenId?: string;
}
