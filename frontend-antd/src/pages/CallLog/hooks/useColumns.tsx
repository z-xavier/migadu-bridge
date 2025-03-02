import {
  ProColumns,
  ProDescriptionsItemProps,
} from '@ant-design/pro-components';
import { CallLogItem } from '../../../services/call-log/typings';
import { formatDateRange, formatDateTime } from '../../../utils/day';

export const useColumns = () => {
  const columns: ProColumns<CallLogItem>[] = [
    {
      title: 'ID',
      dataIndex: 'id',
      hideInForm: true,
      search: false,
      editable: false,
      hideInTable: true,
    },
    {
      title: 'token Id',
      dataIndex: 'tokenId',
      hideInForm: true,
      search: false,
      editable: false,
    },
    {
      title: '目标邮箱',
      dataIndex: 'targetEmail',
      formItemProps: {
        rules: [
          {
            required: true,
            message: '目标邮箱为必填项',
          },
          {
            pattern: /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/,
            message: '请输入正确的邮箱地址',
          },
        ],
      },
      valueType: 'text',
      editable: false,
    },
    {
      title: '模拟服务商',
      dataIndex: 'mockProvider',
      valueType: 'text',
      valueEnum: {
        addy: { text: 'addy', status: 'addy' },
        sl: { text: 'sl', status: 'sl' },
      },
      formItemProps: {
        rules: [
          {
            required: true,
            message: '模拟种类为必填项',
          },
        ],
      },
      editable: false,
    },
    {
      title: '请求路径',
      dataIndex: 'requestPath',
      valueType: 'text',
    },
    {
      title: '别名',
      dataIndex: 'genAlias',
      valueType: 'text',
    },
    {
      title: 'IP',
      dataIndex: 'requestIp',
      valueType: 'text',
    },
    {
      title: '调用时间',
      dataIndex: 'requestAt',
      valueType: 'dateTime',
      sorter: true,
      search: false,
      hideInForm: true,
      editable: false,
      render: (_dom, record) => {
        return formatDateTime(record.requestAt);
      },
    },
    {
      title: '调用时间',
      dataIndex: 'requestAt',
      valueType: 'dateRange',
      hideInTable: true,
      search: {
        transform: formatDateRange('requestAt'),
      },
      hideInForm: true,
    },
  ];

  const checkColumns: ProDescriptionsItemProps<CallLogItem>[] = [
    {
      title: 'ID',
      dataIndex: 'id',
    },
    {
      title: 'token Id',
      dataIndex: 'tokenId',
    },
    {
      title: '目标邮箱',
      dataIndex: 'targetEmail',
      formItemProps: {
        rules: [
          {
            required: true,
            message: '目标邮箱为必填项',
          },
          {
            pattern: /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/,
            message: '请输入正确的邮箱地址',
          },
        ],
      },
    },
    {
      title: '模拟服务商',
      dataIndex: 'mockProvider',
      valueType: 'text',
      valueEnum: {
        addy: { text: 'addy', status: 'addy' },
        sl: { text: 'sl', status: 'sl' },
      },
    },
    {
      title: '请求路径',
      dataIndex: 'requestPath',
    },
    {
      title: '别名',
      dataIndex: 'genAlias',
    },
    {
      title: 'IP',
      dataIndex: 'requestIp',
    },
    {
      title: '调用时间',
      dataIndex: 'requestAt',
      valueType: 'dateTime',
      render: (_dom, record) => {
        return formatDateTime(record.requestAt);
      },
    },
  ];

  return {
    columns,
    checkColumns,
  };
};
