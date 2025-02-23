import { useCallback, useState } from 'react'
import { headSetting } from '../constant'
import { useModel } from './useModel'

export const useViewModel = () => {
  const [curPage, setCurPage] = useState(0)

  const { data, isLoading } = useModel(curPage)
  const handleChangePage = useCallback((num: number) => {
    setCurPage(num)
  }, [])

  return {
    head: headSetting,
    data: data?.list,
    pageData: {
      page: curPage,
      pageSize: data?.pageSize,
    },
    isLoading,
    onPageChange: handleChangePage,
  }
}
