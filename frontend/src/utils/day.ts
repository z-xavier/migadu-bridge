import dayjs from 'dayjs'

export const formatDateTime = (val?: string | number) => {
  if (val === undefined || val === null) {
    return dayjs().format('YYYY-MM-DD HH:mm:ss') // 使用当前时间作为默认值
  }

  const parsedDate = dayjs(val)
  if (!parsedDate.isValid()) {
    throw new Error('Invalid date provided')
  }

  return parsedDate.format('YYYY-MM-DD HH:mm:ss')
}
