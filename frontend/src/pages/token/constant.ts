import { TokenListItem } from './interface'

export const headSetting: Partial<Record<keyof TokenListItem, string>> = {
  id: 'ID',
  targetEmail: '目标邮箱',
  mockProvider: '模拟种类',
  description: '描述',
  token: '密钥',
  expiryTime: '密钥过期时间',
  lastCallTime: '最近一次调用',
}
