import useSWR from 'swr'
import { tokenListUrl } from '../../../service'
import { getFetcher } from '../../../utils/swr'

export const useModel = (curPage: number) => {
  const { data, isLoading } = useSWR(
    {
      url: tokenListUrl,
      params: {
        page: curPage + 1,
      },
    },
    getFetcher
  )

  return {
    data,
    isLoading,
  }
}
