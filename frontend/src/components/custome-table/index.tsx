import {
  Stack,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableFooter,
  TableHead,
  TablePagination,
  TableRow,
} from '@mui/material'
import { LocalizationProvider } from '@mui/x-date-pickers'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'
import { DatePicker } from '@mui/x-date-pickers/DatePicker'
import { Dayjs } from 'dayjs'
import { useState } from 'react'

export type ConfigTypes = {
  field: string
  headName: string
  render?: (val?: string | number) => React.ReactNode
}

interface CustomTableProps<H extends Record<string, string>> {
  config?: ConfigTypes[]
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
  const { config, data, pageData, onPageChange } = props
  const [fromDate, setFromDate] = useState<Dayjs | null>(null)

  return (
    <Stack direction="column" spacing={2} padding={8}>
      <LocalizationProvider dateAdapter={AdapterDayjs}>
        <Stack direction="row">
          <DatePicker
            value={fromDate}
            onChange={(newValue) => setFromDate(newValue)}
          />
        </Stack>
      </LocalizationProvider>
      <TableContainer>
        <Table>
          <TableHead>
            <TableRow>
              {config?.map((val) => (
                <TableCell key={`${val.field}-head`}>{val?.headName}</TableCell>
              ))}
            </TableRow>
          </TableHead>
          <TableBody>
            {data?.map((item, itemIdx) => (
              <TableRow key={item?.id ?? itemIdx} hover>
                {config?.map((val, idx) => (
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
    </Stack>
  )
}
