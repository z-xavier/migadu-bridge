import {
  ProColumns,
  ProDescriptionsItemProps,
} from '@ant-design/pro-components';
import { Button, ConfigProvider, Flex, Modal } from 'antd';
import { TokenListItem } from '../../../services/token/typings';
import { formatDateRange, formatDateTime } from '../../../utils/day';

interface UseTokenColumnsProps {
  setUpdateModalVisible: (visible: boolean) => void;
  setRow: (row: TokenListItem) => void;
  setStepFormValues: (values: TokenListItem) => void;
  handleUpdateStatus: (values: { id: string; status: number }) => void;
  handleRemove: (id: string) => Promise<boolean>;
}
export const useTokenColumns = ({
  setUpdateModalVisible,
  setRow,
  setStepFormValues,
  handleUpdateStatus,
  handleRemove,
}: UseTokenColumnsProps) => {
  const editColumns: ProColumns<TokenListItem>[] = [
    {
      title: '描述',
      dataIndex: 'description',
      valueType: 'textarea',
      ellipsis: true,
    },
    {
      title: '过期时间',
      dataIndex: 'expiryAt',
      valueType: 'dateTime',
      sorter: true,
      search: false,
      formItemProps: {
        rules: [
          {
            required: true,
            message: '过期时间为必填项',
          },
        ],
      },
    },
  ];

  const columns: ProColumns<TokenListItem>[] = [
    {
      title: 'ID',
      dataIndex: 'id',
      hideInForm: true,
      search: false,
      editable: false,
      hideInTable: true,
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
      title: '模拟种类',
      dataIndex: 'mockProvider',
      valueType: 'text',
      // filters: true,
      // onFilter: true,
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
      title: '描述',
      dataIndex: 'description',
      valueType: 'textarea',
      ellipsis: true,
      hideInTable: true,
    },
    {
      title: '密钥',
      dataIndex: 'token',
      valueType: 'text',
      hideInForm: true,
      copyable: true,
      editable: false,
    },
    {
      title: '状态',
      dataIndex: 'status',
      valueType: 'select',
      width: 80,
      valueEnum: {
        1: { text: '未使用', status: 'Error' },
        2: { text: '使用中', status: 'Success' },
        3: { text: '暂停中', status: 'Processing' },
      },
      hideInForm: true,
      editable: false,
    },
    {
      title: '过期时间',
      dataIndex: 'expiryAt',
      valueType: 'dateTime',
      sorter: true,
      search: false,
      formItemProps: {
        rules: [
          {
            required: true,
            message: '过期时间为必填项',
          },
        ],
      },
      render: (_dom, record) => {
        return formatDateTime(record.expiryAt);
      },
    },
    {
      title: '过期时间',
      dataIndex: 'expiryAt',
      valueType: 'dateRange',
      hideInTable: true,
      search: {
        transform: formatDateRange('expiryAt'),
      },
      hideInForm: true,
    },
    {
      title: '最近一次调用',
      dataIndex: 'lastCalledAt',
      valueType: 'dateTime',
      sorter: true,
      search: false,
      hideInForm: true,
      editable: false,
      render: (_dom, record) => {
        return formatDateTime(record.lastCalledAt);
      },
    },
    {
      title: '最近一次调用',
      dataIndex: 'lastCalledAt',
      valueType: 'dateRange',
      hideInTable: true,
      search: {
        transform: formatDateRange('lastCalledAt'),
      },
      hideInForm: true,
    },
    {
      title: '创建时间',
      dataIndex: 'createdAt',
      valueType: 'dateTime',
      sorter: true,
      search: false,
      hideInForm: true,
      editable: false,
      render: (_dom, record) => {
        return formatDateTime(record.createdAt);
      },
      hideInTable: true,
    },
    {
      title: '创建时间',
      dataIndex: 'createdAt',
      valueType: 'dateRange',
      hideInTable: true,
      search: {
        transform: formatDateRange('createdAt'),
      },
      hideInForm: true,
    },
    {
      title: '更新时间',
      dataIndex: 'updatedAt',
      valueType: 'dateTime',
      sorter: true,
      search: false,
      hideInForm: true,
      editable: false,
      render: (_dom, record) => {
        return formatDateTime(record.updatedAt);
      },
    },
    {
      title: '更新时间',
      dataIndex: 'updatedAt',
      valueType: 'dateRange',
      hideInTable: true,
      search: {
        transform: formatDateRange('updatedAt'),
      },
      hideInForm: true,
    },
    {
      title: '操作',
      dataIndex: 'option',
      valueType: 'option',
      width: 240,
      fixed: 'right',
      render: (_node, record, _number, action) => (
        <ConfigProvider componentSize="small">
          <Flex gap="small" wrap>
            {record?.status === 1 && (
              <Button
                variant="text"
                onClick={() => {
                  handleUpdateStatus({ id: record.id, status: 3 });
                }}
              >
                激活
              </Button>
            )}
            {record?.status === 2 && (
              <Button
                variant="text"
                onClick={() => {
                  handleUpdateStatus({ id: record.id, status: 3 });
                }}
              >
                暂停
              </Button>
            )}
            {record?.status === 3 && (
              <Button
                variant="text"
                onClick={() => {
                  handleUpdateStatus({ id: record.id, status: 2 });
                }}
              >
                使用
              </Button>
            )}
            <Button
              variant="text"
              onClick={() => {
                setRow(record);
              }}
            >
              详情
            </Button>
            <Button
              variant="text"
              onClick={() => {
                setUpdateModalVisible(true);
                setStepFormValues(record);
              }}
            >
              编辑
            </Button>
            <Button
              variant="text"
              onClick={() => {
                Modal.confirm({
                  title: '确定要删除这项纪录吗？',
                  okText: '确认',
                  cancelText: '取消',
                  onOk: async () => {
                    const success = await handleRemove(record.id);
                    if (success) {
                      if (action) {
                        action?.reload();
                      }
                    }
                  },
                });
              }}
            >
              删除
            </Button>
          </Flex>
        </ConfigProvider>
      ),
    },
  ];

  const checkColumns: ProDescriptionsItemProps<TokenListItem>[] = [
    {
      title: 'ID',
      dataIndex: 'id',
    },
    {
      title: '目标邮箱',
      dataIndex: 'targetEmail',
    },
    {
      title: '模拟种类',
      dataIndex: 'mockProvider',
      valueEnum: {
        addy: { text: 'addy', status: 'addy' },
        sl: { text: 'sl', status: 'sl' },
      },
    },
    {
      title: '描述',
      dataIndex: 'description',
      valueType: 'textarea',
    },
    {
      title: '密钥',
      dataIndex: 'token',
      valueType: 'text',
      copyable: true,
    },
    {
      title: '状态',
      dataIndex: 'status',
      valueType: 'select',
      valueEnum: {
        1: { text: '未使用', status: 'Error' },
        2: { text: '使用中', status: 'Success' },
        3: { text: '暂停中', status: 'Processing' },
      },
    },
    {
      title: '过期时间',
      dataIndex: 'expiryAt',
      valueType: 'dateTime',
      render: (_dom, record) => {
        return formatDateTime(record.expiryAt);
      },
    },
    {
      title: '最近一次调用',
      dataIndex: 'lastCalledAt',
      valueType: 'dateTime',
      render: (_dom, record) => {
        return formatDateTime(record.lastCalledAt);
      },
    },
    {
      title: '创建时间',
      dataIndex: 'createdAt',
      valueType: 'dateTime',
      render: (_dom, record) => {
        return formatDateTime(record.createdAt);
      },
    },
    {
      title: '更新时间',
      dataIndex: 'updatedAt',
      valueType: 'dateTime',
      render: (_dom, record) => {
        return formatDateTime(record.updatedAt);
      },
    },
  ];

  return {
    columns,
    checkColumns,
    editColumns,
  };
};
