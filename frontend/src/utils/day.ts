import dayjs from 'dayjs';

export const formatDateTime = (val?: number) => {
  if (val === undefined || val === null) {
    return '';
  }

  const parsedDate = dayjs.unix(val);
  if (!parsedDate.isValid()) {
    throw new Error('Invalid date provided');
  }

  return parsedDate.format('YYYY-MM-DD HH:mm:ss');
};
export const getBeginOfDay = (date: dayjs.ConfigType) => {
  return dayjs(date).startOf('day').unix();
};

export const getEndOfDay = (date: dayjs.ConfigType) => {
  return dayjs(date).endOf('day').unix();
};

export const formatDateRange =
  (fieldName: string) => (val: dayjs.ConfigType[]) => ({
    [`${fieldName}Begin`]: getBeginOfDay(val[0]),
    [`${fieldName}End`]: getEndOfDay(val[1]),
  });
