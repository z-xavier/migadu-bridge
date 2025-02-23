export interface TokenListItem {
  id?: string
  targetEmail?: string
  mockProvider?: string
  description?: string
  token?: string
  expiryTime?: number
  lastCallTime?: number
  status?: number // 1: 开启 2: 暂停使用
}
