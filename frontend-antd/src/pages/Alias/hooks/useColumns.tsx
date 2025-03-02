import {
  ProColumns,
  ProDescriptionsItemProps,
} from '@ant-design/pro-components';
import { AliasItem } from '../../../services/alias/typings';

export const useColumns = () => {
  const columns: ProColumns<AliasItem>[] = [
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
      title: 'call log Id',
      dataIndex: 'callLogId',
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
      title: '别名',
      dataIndex: 'alias',
      valueType: 'text',
    },
  ];

  const checkColumns: ProDescriptionsItemProps<AliasItem>[] = [
    {
      title: 'ID',
      dataIndex: 'id',
    },
    {
      title: 'token Id',
      dataIndex: 'tokenId',
    },
    {
      title: 'call log Id',
      dataIndex: 'callLogId',
    },
    {
      title: '目标邮箱',
      dataIndex: 'targetEmail',
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
      title: '别名',
      dataIndex: 'alias',
      valueType: 'text',
    },
  ];

  return {
    columns,
    checkColumns,
  };
};
