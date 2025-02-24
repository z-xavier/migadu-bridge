import {
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableFooter,
  TableHead,
  TablePagination,
  TableRow,
} from '@mui/material'

export type SettingTypes = {
  field: string
  headName: string
  render?: (val?: string | number) => React.ReactNode
}

interface CustomTableProps<H extends Record<string, string>> {
  setting?: SettingTypes[]
  data?: H[]
  pageData?: {
    page: number
    pageSize: number
  }
  onPageChange: (page: number) => void
}

export default function CustomTable<H extends Record<string, string>>(
  props: Readonly<CustomTableProps<H>>
) {
  const { setting, data, pageData, onPageChange } = props

  return (
    <TableContainer>
      <Table>
        <TableHead>
          <TableRow>
            {setting?.map((val) => (
              <TableCell key={`${val.field}-head`}>{val?.headName}</TableCell>
            ))}
          </TableRow>
        </TableHead>
        <TableBody>
          {data?.map((item, itemIdx) => (
            <TableRow key={item?.id ?? itemIdx} hover>
              {setting?.map((val, idx) => (
                <TableCell key={`${val.field}-head-${idx}`}>
                  {val.render
                    ? val.render(item?.[val?.field])
                    : item?.[val?.field]}
                </TableCell>
              ))}
            </TableRow>
          ))}
        </TableBody>
        <TableFooter></TableFooter>
      </Table>
      {pageData && (
        <TablePagination
          component="div"
          rowsPerPage={10}
          rowsPerPageOptions={[-1]}
          page={pageData?.page}
          count={pageData?.pageSize}
          onPageChange={(e, num) => {
            onPageChange?.(num)
          }}
        />
      )}
    </TableContainer>
  )
}
