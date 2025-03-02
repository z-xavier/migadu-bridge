export interface CallLogListRequest {
  mockProvider?: MockProvider;
  orderBy?: string;
  page?: number;
  pageSize?: number;
  requestAtBegin?: number;
  requestAtEnd?: number;
  requestIp?: string;
  requestPath?: string;
  targetEmail?: string;
}

export enum MockProvider {
  Addy = 'addy',
  Sl = 'sl',
}

export interface CallLogListResponse {
  code: string;
  data: CallLogList;
  message: string;
}

export interface CallLogList {
  list: List[];
  page: number;
  pageSize: number;
  total: number;
}

export interface CallLogItem {
  genAlias?: string;
  id?: string;
  mockProvider?: string;
  requestAt?: number;
  requestIp?: string;
  requestPath?: string;
  targetEmail?: string;
  tokenId?: string;
}
