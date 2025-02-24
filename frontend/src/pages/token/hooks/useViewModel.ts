import { useCallback, useState } from 'react'
import { SettingTypes } from '../../../components/custome-table'
import { formatDateTime } from '../../../utils/day'
import { useModel } from './useModel'

const tableSetting: SettingTypes[] = [
  {
    field: 'id',
    headName: 'ID',
  },
  {
    field: 'targetEmail',
    headName: '目标邮箱',
  },
  {
    field: 'mockProvider',
    headName: '模拟种类',
  },
  {
    field: 'description',
    headName: '描述',
  },
  {
    field: 'token',
    headName: '密钥',
  },
  {
    field: 'createdAt',
    headName: '创建时间',
    render: formatDateTime,
  },
  {
    field: 'expiryAt',
    headName: '过期时间',
    render: formatDateTime,
  },
  {
    field: 'updatedAt',
    headName: '更新时间',
    render: formatDateTime,
  },
  {
    field: 'lastCalledAt',
    headName: '最近一次调用',
    render: formatDateTime,
  },
  {
    field: 'status',
    headName: '状态',
  },
]

export const useViewModel = () => {
  const [curPage, setCurPage] = useState(0)

  const { data, isLoading } = useModel(curPage)
  const handleChangePage = useCallback((num: number) => {
    setCurPage(num)
  }, [])

  return {
    setting: tableSetting,
    data: data?.list,
    pageData: {
      page: curPage,
      pageSize: data?.pageSize,
    },
    isLoading,
    onPageChange: handleChangePage,
  }
}
