import CustomTable from '../../components/custome-table'
import Layout from '../../components/layout'
import { useViewModel } from './hooks/useViewModel'

export default function TokenPage() {
  const props = useViewModel()

  return (
    <Layout>
      <CustomTable {...props} />
    </Layout>
  )
}
