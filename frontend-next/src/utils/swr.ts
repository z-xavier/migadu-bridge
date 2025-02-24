import axios from 'axios'

export const getFetcher = ({
  url,
  params,
}: {
  url: string
  params?: Record<string, unknown>
}) => {
  return axios.get(url, { params }).then((res) => res.data)
}
