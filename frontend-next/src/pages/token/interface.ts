export interface TokenListItem {
  createdAt: number
  description: string
  expiryAt: number
  id: string
  lastCalledAt: number
  mockProvider: string
  status: number
  targetEmail: string
  token: string
  updatedAt: number
}

export interface TokenListResponse {
  list: TokenListItem[]
  page: number
  pageSize: number
}
