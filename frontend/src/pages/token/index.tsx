import { useCallback, useState } from 'react'
import useSWR from 'swr'
import CustomTable from '../../components/custome-table'
import Layout from '../../components/layout'
import { getFetcher } from '../../utils/swr'
import { headSetting } from './constant'

export default function TokenPage() {
  const [curPage, setCurPage] = useState(0)

  const { data, error, isLoading } = useSWR(
    {
      url: 'https://run.mocky.io/v3/71978153-25e9-4ee6-a3e3-f008b6abd82e',
      params: {
        page: curPage + 1,
      },
    },
    getFetcher
  )

  console.log('data', data)
  console.log('error', error)
  console.log('isLoading', isLoading)

  const handleChangePage = useCallback((num: number) => {
    setCurPage(num)
  }, [])

  return (
    <Layout>
      <CustomTable
        head={headSetting}
        data={data?.list}
        pageData={{
          page: curPage,
          pageSize: data?.pageSize,
        }}
        onPageChange={handleChangePage}
      />
    </Layout>
  )
}
