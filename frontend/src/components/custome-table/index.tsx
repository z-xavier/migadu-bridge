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

interface CustomTableProps<H extends Record<string, string>> {
  head?: H
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
  const { head, data, pageData, onPageChange } = props

  const keys = head ? Object.keys(head) : []

  return (
    <TableContainer>
      <Table>
        <TableHead>
          <TableRow>
            {keys?.map((key) => (
              <TableCell key={key}>{head?.[key]}</TableCell>
            ))}
          </TableRow>
        </TableHead>
        <TableBody>
          {data?.map((val, index) => (
            <TableRow key={val?.id ?? index} hover>
              {keys?.map((key, keyIdx) => (
                <TableCell key={index + ' ' + keyIdx}>{val?.[key]}</TableCell>
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
